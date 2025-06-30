package btc

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net"
	"time"

	"github.com/btcsuite/btcd/wire"
	"github.com/rs/zerolog/log"
	"github.com/scalarorg/data-models/chains"
)

// Connect establishes connection to the electrum server
func (c *BtcClient) Connect() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.isConnected {
		return nil
	}

	endpoint := fmt.Sprintf("%s:%d", c.config.ElectrumHost, c.config.ElectrumPort)
	conn, err := net.DialTimeout("tcp", endpoint, c.config.DialTimeout)
	if err != nil {
		return fmt.Errorf("failed to connect to electrum server %s: %w", endpoint, err)
	}

	c.conn = conn
	c.isConnected = true

	log.Info().Str("endpoint", endpoint).Msg("Connected to electrum server")
	return nil
}

// Disconnect closes the connection to the electrum server
func (c *BtcClient) Disconnect() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if !c.isConnected {
		return nil
	}

	if c.conn != nil {
		err := c.conn.Close()
		c.conn = nil
		c.isConnected = false
		return err
	}

	return nil
}

// Reconnect attempts to reconnect to the electrum server with exponential backoff
func (c *BtcClient) Reconnect() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	// Close existing connection if any
	if c.conn != nil {
		c.conn.Close()
		c.conn = nil
		c.isConnected = false
	}

	// Calculate delay with exponential backoff
	delay := c.baseReconnectDelay * time.Duration(1<<c.reconnectAttempts)
	if delay > c.maxReconnectDelay {
		delay = c.maxReconnectDelay
	}

	log.Info().
		Str("host", c.config.ElectrumHost).
		Int("port", c.config.ElectrumPort).
		Int("attempt", c.reconnectAttempts+1).
		Dur("delay", delay).
		Msg("Attempting to reconnect to electrum server")

	// Wait before attempting reconnection
	time.Sleep(delay)

	// Attempt to connect
	endpoint := fmt.Sprintf("%s:%d", c.config.ElectrumHost, c.config.ElectrumPort)
	conn, err := net.DialTimeout("tcp", endpoint, c.config.DialTimeout)
	if err != nil {
		c.reconnectAttempts++
		return fmt.Errorf("failed to reconnect to electrum server %s: %w", endpoint, err)
	}

	// Connection successful
	c.conn = conn
	c.isConnected = true
	c.reconnectAttempts = 0 // Reset attempts on successful connection

	log.Info().
		Str("endpoint", endpoint).
		Msg("Successfully reconnected to electrum server")

	return nil
}

// ConnectWithRetry continuously attempts to maintain connection with automatic reconnection
func (c *BtcClient) ConnectWithRetry(ctx context.Context) {
	var retryInterval = c.baseReconnectDelay
	maxRetryInterval := c.maxReconnectDelay

	for {
		select {
		case <-ctx.Done():
			log.Info().Str("host", c.config.ElectrumHost).Msg("Context cancelled, stopping reconnection")
			return
		case <-c.stopChan:
			log.Info().Str("host", c.config.ElectrumHost).Msg("Stop signal received, stopping reconnection")
			return
		default:
			// Try to connect
			err := c.Connect()
			if err != nil {
				log.Error().Err(err).Str("host", c.config.ElectrumHost).Msg("Failed to connect to electrum server")
			} else {
				// Connection successful, start monitoring
				go c.monitorConnection(ctx)
				return
			}

			// If context is cancelled, stop retrying
			if ctx.Err() != nil {
				log.Info().Str("host", c.config.ElectrumHost).Msg("Context cancelled, stopping reconnection")
				return
			}

			// Wait before retrying
			log.Info().Str("host", c.config.ElectrumHost).Dur("retryInterval", retryInterval).Msg("Reconnecting...")
			time.Sleep(retryInterval)

			// Exponential backoff with cap
			if retryInterval < maxRetryInterval {
				retryInterval *= 2
			}
		}
	}
}

// monitorConnection monitors the connection and triggers reconnection if needed
func (c *BtcClient) monitorConnection(ctx context.Context) {
	ticker := time.NewTicker(c.config.PingInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-c.stopChan:
			return
		case <-ticker.C:
			// Check if connection is still alive
			if !c.isConnectionAlive() {
				log.Warn().Str("host", c.config.ElectrumHost).Msg("Connection lost, attempting reconnection")

				// Trigger reconnection
				select {
				case c.reconnectChan <- struct{}{}:
				default:
					// Channel is full, skip this reconnection attempt
				}
			}
		}
	}
}

// isConnectionAlive checks if the connection is still alive
func (c *BtcClient) isConnectionAlive() bool {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if !c.isConnected || c.conn == nil {
		return false
	}

	// Try to ping the server
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := c.callRPCMethod(ctx, "server.ping")
	return err == nil
}

// callRPCMethod makes an RPC call to the electrum server
func (c *BtcClient) callRPCMethod(ctx context.Context, method string, params ...interface{}) ([]byte, error) {
	c.mu.RLock()
	if !c.isConnected {
		c.mu.RUnlock()
		return nil, fmt.Errorf("not connected to electrum server")
	}
	c.mu.RUnlock()

	// Create JSON-RPC request
	request := map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      1,
		"method":  method,
		"params":  params,
	}

	requestBytes, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	// Send request
	c.mu.Lock()
	if c.conn == nil {
		c.mu.Unlock()
		// Trigger reconnection
		select {
		case c.reconnectChan <- struct{}{}:
		default:
		}
		return nil, fmt.Errorf("connection is nil")
	}

	// Set write deadline
	err = c.conn.SetWriteDeadline(time.Now().Add(c.config.MethodTimeout))
	if err != nil {
		c.mu.Unlock()
		// Trigger reconnection
		select {
		case c.reconnectChan <- struct{}{}:
		default:
		}
		return nil, fmt.Errorf("failed to set write deadline: %w", err)
	}

	_, err = c.conn.Write(requestBytes)
	c.mu.Unlock()
	if err != nil {
		// Trigger reconnection
		select {
		case c.reconnectChan <- struct{}{}:
		default:
		}
		return nil, fmt.Errorf("failed to write request: %w", err)
	}

	// Read response
	c.mu.Lock()
	err = c.conn.SetReadDeadline(time.Now().Add(c.config.MethodTimeout))
	if err != nil {
		c.mu.Unlock()
		// Trigger reconnection
		select {
		case c.reconnectChan <- struct{}{}:
		default:
		}
		return nil, fmt.Errorf("failed to set read deadline: %w", err)
	}

	buffer := make([]byte, 4096)
	n, err := c.conn.Read(buffer)
	c.mu.Unlock()
	if err != nil {
		// Trigger reconnection
		select {
		case c.reconnectChan <- struct{}{}:
		default:
		}
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	response := buffer[:n]

	// Parse JSON-RPC response
	var jsonResponse map[string]interface{}
	err = json.Unmarshal(response, &jsonResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	// Check for error
	if errorObj, exists := jsonResponse["error"]; exists && errorObj != nil {
		return nil, fmt.Errorf("RPC error: %v", errorObj)
	}

	// Extract result
	result, exists := jsonResponse["result"]
	if !exists {
		return nil, fmt.Errorf("no result in response")
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal result: %w", err)
	}

	return resultBytes, nil
}

// GetBlockHeader retrieves block header information
func (c *BtcClient) GetBlockHeader(ctx context.Context, height int64) (*chains.BtcBlockHeader, error) {
	result, err := c.callRPCMethod(ctx, "blockchain.block.header", height)
	if err != nil {
		return nil, fmt.Errorf("failed to get block header: %w", err)
	}

	var headerEntry chains.HeaderEntry
	err = json.Unmarshal(result, &headerEntry)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal header hex: %w", err)
	}

	blockHeader := chains.BtcBlockHeader{}
	err = blockHeader.ParseHeaderEntry(&headerEntry)
	if err != nil {
		return nil, fmt.Errorf("failed to parse header entry: %w", err)
	}

	return &blockHeader, nil
}

// GetBlock retrieves a full block with transactions
func (c *BtcClient) GetBlockFromElectrum(ctx context.Context, height int64) (*wire.MsgBlock, error) {
	result, err := c.callRPCMethod(ctx, "blockchain.block.get", height)
	if err != nil {
		return nil, fmt.Errorf("failed to get block: %w", err)
	}

	var blockHex string
	err = json.Unmarshal(result, &blockHex)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal block hex: %w", err)
	}

	blockBytes, err := hex.DecodeString(blockHex)
	if err != nil {
		return nil, fmt.Errorf("failed to decode block hex: %w", err)
	}

	// Parse Bitcoin block
	var block wire.MsgBlock
	err = block.Deserialize(bytes.NewReader(blockBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to deserialize block: %w", err)
	}

	return &block, nil
}

// IndexBlock indexes a block by height
func (c *BtcClient) IndexBlockHeader(ctx context.Context, height int64) error {
	// Get block header
	header, err := c.GetBlockHeader(ctx, height)
	if err != nil {
		return fmt.Errorf("failed to get block header for height %d: %w", height, err)
	}

	// Save block header to database
	err = c.dbAdapter.CreateBtcBlockHeader(header)
	if err != nil {
		return fmt.Errorf("failed to save block header: %w", err)
	}

	// Get full block
	block, err := c.GetBlockFromElectrum(ctx, height)
	if err != nil {
		return fmt.Errorf("failed to get block for height %d: %w", height, err)
	}

	// Process transactions
	vaultTxs := []*chains.VaultTransaction{}
	for i, tx := range block.Transactions {
		vaultTx, err := c.ParseVaultMsgTx(tx, i, height, header.Hash, int64(header.Time))
		if err != nil {
			log.Warn().Err(err).Int("txIndex", i).Msg("Failed to parse transaction")
			continue
		}

		if vaultTx != nil {
			// Set timestamp from block header
			vaultTx.Timestamp = uint64(header.Time)

			vaultTxs = append(vaultTxs, vaultTx)

			// Save vault transaction to database
			err = c.StoreVaultTransaction(vaultTx)
			if err != nil {
				log.Warn().Err(err).Msg("Failed to save vault transaction")
			}
		}
	}

	log.Info().Int64("height", height).Int("vaultTxs", len(vaultTxs)).Msg("Indexed block")
	return nil
}

// GetLatestHeightFromElectrum gets the latest block height from the electrum server
func (c *BtcClient) GetLatestHeightFromElectrum(ctx context.Context) (int64, error) {
	result, err := c.callRPCMethod(ctx, "blockchain.numblocks.subscribe")
	if err != nil {
		return 0, fmt.Errorf("failed to get latest height: %w", err)
	}

	var height int64
	err = json.Unmarshal(result, &height)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal height: %w", err)
	}

	return height, nil
}

// Start starts the electrum indexer
func (c *BtcClient) StartElectrumIndexer(ctx context.Context) error {
	// Start connection with retry
	go c.ConnectWithRetry(ctx)

	// Wait for initial connection
	timeout := time.After(30 * time.Second)
	for {
		select {
		case <-timeout:
			return fmt.Errorf("timeout waiting for initial connection to electrum server")
		case <-ctx.Done():
			return ctx.Err()
		default:
			c.mu.RLock()
			connected := c.isConnected
			c.mu.RUnlock()

			if connected {
				break
			}
			time.Sleep(100 * time.Millisecond)
		}
		if c.isConnected {
			break
		}
	}

	// Get latest indexed height from DB
	dbLatest, err := c.GetLatestIndexedHeight(ctx)
	if err != nil {
		return fmt.Errorf("failed to get latest indexed height from DB: %w", err)
	}

	// Get latest height from Electrum
	electrumLatest, err := c.GetLatestHeightFromElectrum(ctx)
	if err != nil {
		return fmt.Errorf("failed to get latest height from electrum: %w", err)
	}

	c.lastHeight = dbLatest
	log.Info().Int64("dbLatest", dbLatest).Int64("electrumLatest", electrumLatest).Msg("Electrum indexer starting catch-up and live listeners")

	// Catch-up goroutine: index all missing blocks from DB up to Electrum tip
	go func() {
		for height := dbLatest + 1; height <= electrumLatest; height++ {
			if ctx.Err() != nil {
				return
			}
			err := c.IndexBlockHeader(ctx, height)
			if err != nil {
				log.Warn().Err(err).Int64("height", height).Msg("Failed to index catch-up block")
				continue
			}
			c.lastHeight = height
		}
		log.Info().Int64("catchup_to", electrumLatest).Msg("Catch-up complete, switching to live indexing")
	}()

	// Live goroutine: poll for new blocks and index as they appear
	go c.indexBlockHeaders(ctx)

	// Start reconnection handler
	go c.handleReconnection(ctx)

	return nil
}

// handleReconnection handles reconnection events
func (c *BtcClient) handleReconnection(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-c.stopChan:
			return
		case <-c.reconnectChan:
			log.Info().Str("host", c.config.ElectrumHost).Msg("Reconnection triggered")

			// Attempt to reconnect
			err := c.Reconnect()
			if err != nil {
				log.Error().Err(err).Str("host", c.config.ElectrumHost).Msg("Failed to reconnect")
				// Continue monitoring, will retry on next reconnection event
			} else {
				log.Info().Str("host", c.config.ElectrumHost).Msg("Successfully reconnected")
			}
		}
	}
}

// indexBlocks continuously indexes new blocks
func (c *BtcClient) indexBlockHeaders(ctx context.Context) {
	ticker := time.NewTicker(10 * time.Second) // Check for new blocks every 10 seconds
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-c.stopChan:
			return
		case <-ticker.C:
			// Check if we're connected
			c.mu.RLock()
			connected := c.isConnected
			c.mu.RUnlock()

			if !connected {
				log.Warn().Str("host", c.config.ElectrumHost).Msg("Not connected, skipping block indexing")
				continue
			}

			// Get latest height
			latestHeight, err := c.GetLatestHeightFromElectrum(ctx)
			if err != nil {
				log.Warn().Err(err).Str("host", c.config.ElectrumHost).Msg("Failed to get latest height")
				continue
			}

			// Index new blocks
			for height := c.lastHeight + 1; height <= latestHeight; height++ {
				err := c.IndexBlockHeader(ctx, height)
				if err != nil {
					log.Warn().Err(err).Int64("height", height).Str("host", c.config.ElectrumHost).Msg("Failed to index block")
					continue
				}
				c.lastHeight = height
			}
		}
	}
}

// Stop stops the electrum indexer
func (c *BtcClient) Stop() {
	close(c.stopChan)

	// Stop reconnection ticker if it exists
	if c.reconnectTicker != nil {
		c.reconnectTicker.Stop()
	}

	c.Disconnect()
	log.Info().Str("host", c.config.ElectrumHost).Msg("BTC indexer stopped")
}
