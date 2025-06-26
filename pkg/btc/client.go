package btc

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/btcsuite/btcd/rpcclient"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

// BtcClient represents an electrum indexer that directly connects to electrum servers
type BtcClient struct {
	config        *BtcConfig
	rpcClient     *rpcclient.Client
	dbAdapter     *gorm.DB // Separate DB connection for electrum data
	conn          net.Conn
	mu            sync.RWMutex
	isConnected   bool
	lastHeight    int64
	reconnectChan chan struct{}
	stopChan      chan struct{}
	// Reconnection fields
	reconnectAttempts  int
	baseReconnectDelay time.Duration
	maxReconnectDelay  time.Duration
	reconnectTicker    *time.Ticker
}

// Connect establishes connection to the electrum server
func (ei *BtcClient) Connect() error {
	ei.mu.Lock()
	defer ei.mu.Unlock()

	if ei.isConnected {
		return nil
	}

	endpoint := fmt.Sprintf("%s:%d", ei.config.ElectrumHost, ei.config.ElectrumPort)
	conn, err := net.DialTimeout("tcp", endpoint, ei.config.DialTimeout)
	if err != nil {
		return fmt.Errorf("failed to connect to electrum server %s: %w", endpoint, err)
	}

	ei.conn = conn
	ei.isConnected = true

	log.Info().Str("endpoint", endpoint).Msg("Connected to electrum server")
	return nil
}

// Disconnect closes the connection to the electrum server
func (ei *BtcClient) Disconnect() error {
	ei.mu.Lock()
	defer ei.mu.Unlock()

	if !ei.isConnected {
		return nil
	}

	if ei.conn != nil {
		err := ei.conn.Close()
		ei.conn = nil
		ei.isConnected = false
		return err
	}

	return nil
}

// Reconnect attempts to reconnect to the electrum server with exponential backoff
func (ei *BtcClient) Reconnect() error {
	ei.mu.Lock()
	defer ei.mu.Unlock()

	// Close existing connection if any
	if ei.conn != nil {
		ei.conn.Close()
		ei.conn = nil
		ei.isConnected = false
	}

	// Calculate delay with exponential backoff
	delay := ei.baseReconnectDelay * time.Duration(1<<ei.reconnectAttempts)
	if delay > ei.maxReconnectDelay {
		delay = ei.maxReconnectDelay
	}

	log.Info().
		Str("host", ei.config.ElectrumHost).
		Int("port", ei.config.ElectrumPort).
		Int("attempt", ei.reconnectAttempts+1).
		Dur("delay", delay).
		Msg("Attempting to reconnect to electrum server")

	// Wait before attempting reconnection
	time.Sleep(delay)

	// Attempt to connect
	endpoint := fmt.Sprintf("%s:%d", ei.config.ElectrumHost, ei.config.ElectrumPort)
	conn, err := net.DialTimeout("tcp", endpoint, ei.config.DialTimeout)
	if err != nil {
		ei.reconnectAttempts++
		return fmt.Errorf("failed to reconnect to electrum server %s: %w", endpoint, err)
	}

	// Connection successful
	ei.conn = conn
	ei.isConnected = true
	ei.reconnectAttempts = 0 // Reset attempts on successful connection

	log.Info().
		Str("endpoint", endpoint).
		Msg("Successfully reconnected to electrum server")

	return nil
}

// ConnectWithRetry continuously attempts to maintain connection with automatic reconnection
func (ei *BtcClient) ConnectWithRetry(ctx context.Context) {
	var retryInterval = ei.baseReconnectDelay
	maxRetryInterval := ei.maxReconnectDelay

	for {
		select {
		case <-ctx.Done():
			log.Info().Str("host", ei.config.ElectrumHost).Msg("Context cancelled, stopping reconnection")
			return
		case <-ei.stopChan:
			log.Info().Str("host", ei.config.ElectrumHost).Msg("Stop signal received, stopping reconnection")
			return
		default:
			// Try to connect
			err := ei.Connect()
			if err != nil {
				log.Error().Err(err).Str("host", ei.config.ElectrumHost).Msg("Failed to connect to electrum server")
			} else {
				// Connection successful, start monitoring
				go ei.monitorConnection(ctx)
				return
			}

			// If context is cancelled, stop retrying
			if ctx.Err() != nil {
				log.Info().Str("host", ei.config.ElectrumHost).Msg("Context cancelled, stopping reconnection")
				return
			}

			// Wait before retrying
			log.Info().Str("host", ei.config.ElectrumHost).Dur("retryInterval", retryInterval).Msg("Reconnecting...")
			time.Sleep(retryInterval)

			// Exponential backoff with cap
			if retryInterval < maxRetryInterval {
				retryInterval *= 2
			}
		}
	}
}

// monitorConnection monitors the connection and triggers reconnection if needed
func (ei *BtcClient) monitorConnection(ctx context.Context) {
	ticker := time.NewTicker(ei.config.PingInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ei.stopChan:
			return
		case <-ticker.C:
			// Check if connection is still alive
			if !ei.isConnectionAlive() {
				log.Warn().Str("host", ei.config.ElectrumHost).Msg("Connection lost, attempting reconnection")

				// Trigger reconnection
				select {
				case ei.reconnectChan <- struct{}{}:
				default:
					// Channel is full, skip this reconnection attempt
				}
			}
		}
	}
}

// isConnectionAlive checks if the connection is still alive
func (ei *BtcClient) isConnectionAlive() bool {
	ei.mu.RLock()
	defer ei.mu.RUnlock()

	if !ei.isConnected || ei.conn == nil {
		return false
	}

	// Try to ping the server
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := ei.callRPCMethod(ctx, "server.ping")
	return err == nil
}

// callRPCMethod makes an RPC call to the electrum server
func (ei *BtcClient) callRPCMethod(ctx context.Context, method string, params ...interface{}) ([]byte, error) {
	ei.mu.RLock()
	if !ei.isConnected {
		ei.mu.RUnlock()
		return nil, fmt.Errorf("not connected to electrum server")
	}
	ei.mu.RUnlock()

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
	ei.mu.Lock()
	if ei.conn == nil {
		ei.mu.Unlock()
		// Trigger reconnection
		select {
		case ei.reconnectChan <- struct{}{}:
		default:
		}
		return nil, fmt.Errorf("connection is nil")
	}

	// Set write deadline
	err = ei.conn.SetWriteDeadline(time.Now().Add(ei.config.MethodTimeout))
	if err != nil {
		ei.mu.Unlock()
		// Trigger reconnection
		select {
		case ei.reconnectChan <- struct{}{}:
		default:
		}
		return nil, fmt.Errorf("failed to set write deadline: %w", err)
	}

	_, err = ei.conn.Write(requestBytes)
	ei.mu.Unlock()
	if err != nil {
		// Trigger reconnection
		select {
		case ei.reconnectChan <- struct{}{}:
		default:
		}
		return nil, fmt.Errorf("failed to write request: %w", err)
	}

	// Read response
	ei.mu.Lock()
	err = ei.conn.SetReadDeadline(time.Now().Add(ei.config.MethodTimeout))
	if err != nil {
		ei.mu.Unlock()
		// Trigger reconnection
		select {
		case ei.reconnectChan <- struct{}{}:
		default:
		}
		return nil, fmt.Errorf("failed to set read deadline: %w", err)
	}

	buffer := make([]byte, 4096)
	n, err := ei.conn.Read(buffer)
	ei.mu.Unlock()
	if err != nil {
		// Trigger reconnection
		select {
		case ei.reconnectChan <- struct{}{}:
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
