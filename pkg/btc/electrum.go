package btc

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"

	"github.com/btcsuite/btcd/rpcclient"
	"github.com/btcsuite/btcd/wire"
	"github.com/rs/zerolog/log"
	"github.com/scalarorg/data-models/chains"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewElectrumIndexer creates a new electrum indexer
func NewBtcClient(config *BtcConfig) (*BtcClient, error) {
	// Set default values
	if config.DialTimeout == 0 {
		config.DialTimeout = 30 * time.Second
	}
	if config.MethodTimeout == 0 {
		config.MethodTimeout = 60 * time.Second
	}
	if config.PingInterval == 0 {
		config.PingInterval = 30 * time.Second
	}
	if config.MaxReconnectAttempts == 0 {
		config.MaxReconnectAttempts = 120
	}
	if config.ReconnectDelay == 0 {
		config.ReconnectDelay = 5 * time.Second
	}
	if config.BatchSize == 0 {
		config.BatchSize = 1
	}
	if config.Confirmations == 0 {
		config.Confirmations = 1
	}

	// Create separate database connection for electrum data
	dbAdapter, err := gorm.Open(postgres.Open(config.DatabaseURL), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to electrum database: %w", err)
	}

	// Run migrations for electrum tables
	err = runElectrumMigrations(dbAdapter)
	if err != nil {
		return nil, fmt.Errorf("failed to run electrum migrations: %w", err)
	}

	// Configure connection
	connCfg := &rpcclient.ConnConfig{
		Host:         fmt.Sprintf("%s:%d", config.BtcHost, config.BtcPort),
		User:         config.BtcUser,
		Pass:         config.BtcPassword,
		HTTPPostMode: true,
		DisableTLS:   config.BtcSSL == nil || !*config.BtcSSL,
	}

	// Create new client
	rpcClient, err := rpcclient.New(connCfg, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create BTC client for network %s: %w", config.BtcNetwork, err)
	}

	indexer := &BtcClient{
		config:             config,
		dbAdapter:          dbAdapter,
		rpcClient:          rpcClient,
		reconnectChan:      make(chan struct{}, 1),
		stopChan:           make(chan struct{}),
		reconnectAttempts:  0,
		baseReconnectDelay: config.ReconnectDelay,
		maxReconnectDelay:  2 * time.Minute, // Maximum 2 minutes between reconnection attempts
	}

	return indexer, nil
}

// runElectrumMigrations creates the necessary tables for electrum data
func runElectrumMigrations(db *gorm.DB) error {
	return db.AutoMigrate(
		&chains.BtcBlockHeader{},
		&chains.VaultTransaction{},
	)
}

// GetBlockHeader retrieves block header information
func (ei *BtcClient) GetBlockHeader(ctx context.Context, height int64) (*chains.BtcBlockHeader, error) {
	result, err := ei.callRPCMethod(ctx, "blockchain.block.header", height)
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

// serializeTx serializes a transaction to hex string
func (ei *BtcClient) serializeTx(tx *wire.MsgTx) (string, error) {
	var buf bytes.Buffer
	err := tx.Serialize(&buf)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(buf.Bytes()), nil
}

// IndexBlock indexes a block by height
func (ei *BtcClient) IndexBlockHeader(ctx context.Context, height int64) error {
	// Get block header
	header, err := ei.GetBlockHeader(ctx, height)
	if err != nil {
		return fmt.Errorf("failed to get block header for height %d: %w", height, err)
	}

	// Save block header to database
	err = ei.dbAdapter.Create(header).Error
	if err != nil {
		return fmt.Errorf("failed to save block header: %w", err)
	}

	// Get full block
	block, err := ei.GetBlock(ctx, height)
	if err != nil {
		return fmt.Errorf("failed to get block for height %d: %w", height, err)
	}

	// Process transactions
	vaultTxs := []*chains.VaultTransaction{}
	for i, tx := range block.Transactions {
		vaultTx, err := ei.ParseVaultTransaction(tx, height, header.Hash, i)
		if err != nil {
			log.Warn().Err(err).Int("txIndex", i).Msg("Failed to parse transaction")
			continue
		}

		if vaultTx != nil {
			// Set timestamp from block header
			vaultTx.Timestamp = uint64(header.Time)

			vaultTxs = append(vaultTxs, vaultTx)

			// Save vault transaction to database
			err = ei.dbAdapter.Create(vaultTx).Error
			if err != nil {
				log.Warn().Err(err).Msg("Failed to save vault transaction")
			}
		}
	}

	log.Info().Int64("height", height).Int("vaultTxs", len(vaultTxs)).Msg("Indexed block")
	return nil
}

// GetLatestHeight gets the latest block height from the electrum server
func (ei *BtcClient) GetLatestHeight(ctx context.Context) (int64, error) {
	result, err := ei.callRPCMethod(ctx, "blockchain.numblocks.subscribe")
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
func (ei *BtcClient) StartElectrumIndexer(ctx context.Context) error {
	// Start connection with retry
	go ei.ConnectWithRetry(ctx)

	// Wait for initial connection
	timeout := time.After(30 * time.Second)
	for {
		select {
		case <-timeout:
			return fmt.Errorf("timeout waiting for initial connection to electrum server")
		case <-ctx.Done():
			return ctx.Err()
		default:
			ei.mu.RLock()
			connected := ei.isConnected
			ei.mu.RUnlock()

			if connected {
				break
			}
			time.Sleep(100 * time.Millisecond)
		}
		if ei.isConnected {
			break
		}
	}

	// Get latest indexed height from DB
	dbLatest, err := ei.GetLatestIndexedHeight(ctx)
	if err != nil {
		return fmt.Errorf("failed to get latest indexed height from DB: %w", err)
	}

	// Get latest height from Electrum
	electrumLatest, err := ei.GetLatestHeight(ctx)
	if err != nil {
		return fmt.Errorf("failed to get latest height from electrum: %w", err)
	}

	ei.lastHeight = dbLatest
	log.Info().Int64("dbLatest", dbLatest).Int64("electrumLatest", electrumLatest).Msg("Electrum indexer starting catch-up and live listeners")

	// Catch-up goroutine: index all missing blocks from DB up to Electrum tip
	go func() {
		for height := dbLatest + 1; height <= electrumLatest; height++ {
			if ctx.Err() != nil {
				return
			}
			err := ei.IndexBlockHeader(ctx, height)
			if err != nil {
				log.Warn().Err(err).Int64("height", height).Msg("Failed to index catch-up block")
				continue
			}
			ei.lastHeight = height
		}
		log.Info().Int64("catchup_to", electrumLatest).Msg("Catch-up complete, switching to live indexing")
	}()

	// Live goroutine: poll for new blocks and index as they appear
	go ei.indexBlockHeaders(ctx)

	// Start reconnection handler
	go ei.handleReconnection(ctx)

	return nil
}

// handleReconnection handles reconnection events
func (ei *BtcClient) handleReconnection(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-ei.stopChan:
			return
		case <-ei.reconnectChan:
			log.Info().Str("host", ei.config.ElectrumHost).Msg("Reconnection triggered")

			// Attempt to reconnect
			err := ei.Reconnect()
			if err != nil {
				log.Error().Err(err).Str("host", ei.config.ElectrumHost).Msg("Failed to reconnect")
				// Continue monitoring, will retry on next reconnection event
			} else {
				log.Info().Str("host", ei.config.ElectrumHost).Msg("Successfully reconnected")
			}
		}
	}
}

// indexBlocks continuously indexes new blocks
func (ei *BtcClient) indexBlockHeaders(ctx context.Context) {
	ticker := time.NewTicker(10 * time.Second) // Check for new blocks every 10 seconds
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ei.stopChan:
			return
		case <-ticker.C:
			// Check if we're connected
			ei.mu.RLock()
			connected := ei.isConnected
			ei.mu.RUnlock()

			if !connected {
				log.Warn().Str("host", ei.config.ElectrumHost).Msg("Not connected, skipping block indexing")
				continue
			}

			// Get latest height
			latestHeight, err := ei.GetLatestHeight(ctx)
			if err != nil {
				log.Warn().Err(err).Str("host", ei.config.ElectrumHost).Msg("Failed to get latest height")
				continue
			}

			// Index new blocks
			for height := ei.lastHeight + 1; height <= latestHeight; height++ {
				err := ei.IndexBlockHeader(ctx, height)
				if err != nil {
					log.Warn().Err(err).Int64("height", height).Str("host", ei.config.ElectrumHost).Msg("Failed to index block")
					continue
				}
				ei.lastHeight = height
			}
		}
	}
}

// Stop stops the electrum indexer
func (ei *BtcClient) Stop() {
	close(ei.stopChan)

	// Stop reconnection ticker if it exists
	if ei.reconnectTicker != nil {
		ei.reconnectTicker.Stop()
	}

	ei.Disconnect()
	log.Info().Str("host", ei.config.ElectrumHost).Msg("Electrum indexer stopped")
}

// GetLatestIndexedHeight returns the latest block height indexed in the database
func (ei *BtcClient) GetLatestIndexedHeight(ctx context.Context) (int64, error) {
	var header chains.BtcBlockHeader
	err := ei.dbAdapter.WithContext(ctx).
		Order("block_number DESC").
		Limit(1).
		Find(&header).Error
	if err != nil {
		return 0, err
	}
	return int64(header.Height), nil
}
