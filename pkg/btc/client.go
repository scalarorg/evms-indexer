package btc

import (
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/btcsuite/btcd/rpcclient"
	"github.com/rs/zerolog/log"
	"github.com/scalarorg/evms-indexer/config"
	"github.com/scalarorg/evms-indexer/pkg/db"
)

// BtcClient represents an electrum indexer that directly connects to electrum servers
type BtcClient struct {
	config        *BtcConfig
	rpcClient     *rpcclient.Client
	dbAdapter     *db.DatabaseAdapter // Database adapter for electrum data
	reorgHandler  *ReorgHandler       // Handler for reorg detection and rollback
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

	// Create database adapter for electrum data
	dbAdapter, err := db.NewDatabaseAdapter(config.DatabaseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to electrum database: %w", err)
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
		reorgHandler:       &ReorgHandler{DB: dbAdapter},
		reconnectChan:      make(chan struct{}, 1),
		stopChan:           make(chan struct{}),
		reconnectAttempts:  0,
		baseReconnectDelay: config.ReconnectDelay,
		maxReconnectDelay:  2 * time.Minute, // Maximum 2 minutes between reconnection attempts
	}

	return indexer, nil
}

// NewElectrumIndexers creates multiple electrum indexers from configuration
func NewBtcClients(globalConfig *config.Config) ([]*BtcClient, error) {
	if globalConfig == nil || globalConfig.ConfigPath == "" {
		return nil, fmt.Errorf("config path is required")
	}

	btcCfgPath := fmt.Sprintf("%s/btcs.json", globalConfig.ConfigPath)
	configs, err := config.ReadJsonArrayConfig[BtcConfig](btcCfgPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read electrum indexer configs: %w", err)
	}

	indexers := []*BtcClient{}
	for _, cfg := range configs {
		if !cfg.Enable {
			log.Info().Msgf("BTC indexer %s is disabled", cfg.BtcHost)
			continue
		}
		indexer, err := NewBtcClient(&cfg)
		if err != nil {
			log.Error().Err(err).Msgf("failed to create btc client %s: %w", cfg.BtcHost, err)
		} else if indexer != nil {
			log.Info().Msgf("BTC indexer %s is enabled", cfg.BtcHost)
			indexers = append(indexers, indexer)
		}
	}

	return indexers, nil
}
