package indexer

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/scalarorg/evms-indexer/config"
	"github.com/scalarorg/evms-indexer/pkg/btc"
	"github.com/scalarorg/evms-indexer/pkg/db"
	"github.com/scalarorg/evms-indexer/pkg/evm"
)

type Service struct {
	dbAdapter  *db.DatabaseAdapter
	EvmClients []*evm.EvmClient
	BtcClients []*btc.BtcClient
}

func NewService(config *config.Config) (*Service, error) {
	// Initialize EVM clients (with shared database)
	evmClients, err := evm.NewEvmClients(config.ConfigPath)
	if err != nil {
		return nil, fmt.Errorf("failed to create evm clients: %w", err)
	}

	// Initialize Electrum Indexers
	btcClients, err := btc.NewBtcClients(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create btc indexers: %w", err)
	}
	log.Info().Int("Number of BTC indexers", len(btcClients)).
		Int("Number of EVM indexers", len(evmClients)).
		Msg("Indexers initialized")
	return &Service{
		EvmClients: evmClients,
		BtcClients: btcClients,
	}, nil
}

func (s *Service) Start(ctx context.Context) error {
	// Start EVM clients (shared database)
	for _, client := range s.EvmClients {
		// Start listening to new events immediately
		go func(c *evm.EvmClient) {
			c.Start(ctx)
		}(client)
	}

	// Start Btc Clients
	for _, client := range s.BtcClients {
		go func(idx *btc.BtcClient) {
			err := idx.StartBtcIndexer(ctx)
			if err != nil {
				log.Error().Err(err).Msg("Failed to start electrum client")
			}
		}(client)
	}
	log.Info().Msg("Indexers started")
	return nil
}

func (s *Service) Stop() {
	// Stop EVM clients (shared database)
	for _, client := range s.EvmClients {
		client.Stop()
	}

	// Stop Electrum Clients
	for _, client := range s.BtcClients {
		client.Stop()
	}
}
