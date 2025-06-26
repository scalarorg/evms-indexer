package indexer

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/scalarorg/evms-indexer/config"
	"github.com/scalarorg/evms-indexer/pkg/db"
	"github.com/scalarorg/evms-indexer/pkg/evm"
)

type Service struct {
	dbAdapter  *db.DatabaseAdapter
	EvmClients []*evm.EvmClient
}

func NewService(config *config.Config, dbAdapter *db.DatabaseAdapter) (*Service, error) {
	// Initialize EVM clients
	evmClients, err := evm.NewEvmClients(config.ConfigPath, dbAdapter)
	if err != nil {
		return nil, fmt.Errorf("failed to create evm clients: %w", err)
	}
	return &Service{
		dbAdapter:  dbAdapter,
		EvmClients: evmClients,
	}, nil
}

func (s *Service) Start(ctx context.Context) error {
	// Start EVM clients
	for _, client := range s.EvmClients {
		// Process recovered logs in dependent go routine
		go client.ProcessMissingLogs()

		// Start listening to new events immediately
		go func(c *evm.EvmClient) {
			c.Start(ctx)
		}(client)

		// Recover all events in parallel
		go func(c *evm.EvmClient) {
			err := c.RecoverAllEvents(ctx)
			if err != nil {
				log.Warn().Err(err).Msgf("[Indexer] [Start] cannot recover events for evm client %s", c.EvmConfig.GetId())
			} else {
				log.Info().Msgf("[Indexer] [Start] recovered missing events for evm client %s", c.EvmConfig.GetId())
			}
		}(client)
	}
	return nil
}

func (s *Service) Stop() {
	// Stop EVM clients
	for _, client := range s.EvmClients {
		client.Stop()
	}
}
