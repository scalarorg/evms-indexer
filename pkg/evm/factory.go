package evm

import (
	"fmt"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/scalarorg/evms-indexer/config"
	"github.com/scalarorg/evms-indexer/pkg/db"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewEvmClientsWithSeparateDB creates EVM clients with separate database connections
func NewEvmClientsWithSeparateDB(globalConfig *config.Config, dbAdapter *db.DatabaseAdapter) ([]*EvmClient, error) {
	if globalConfig == nil || globalConfig.ConfigPath == "" {
		return nil, fmt.Errorf("config path is not set")
	}

	evmCfgPath := fmt.Sprintf("%s/evm-separate-db.json", globalConfig.ConfigPath)
	configs, err := config.ReadJsonArrayConfig[EvmNetworkConfig](evmCfgPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read evm configs: %w", err)
	}

	evmClients := make([]*EvmClient, 0, len(configs))
	for _, evmConfig := range configs {
		// Create separate database connection for this EVM client
		evmDB, err := gorm.Open(postgres.Open(evmConfig.DatabaseURL), &gorm.Config{})
		if err != nil {
			return nil, fmt.Errorf("failed to connect to evm database for %s: %w", evmConfig.Name, err)
		}

		//Set default value for block time if is not set
		if evmConfig.BlockTime == 0 {
			evmConfig.BlockTime = 12 * time.Second
		} else {
			evmConfig.BlockTime = evmConfig.BlockTime * time.Millisecond
		}
		if evmConfig.RecoverRange == 0 {
			evmConfig.RecoverRange = 1000000
		}

		client, err := NewEvmClient(globalConfig.ConfigPath, &evmConfig, dbAdapter)
		if err != nil {
			log.Warn().Msgf("Failed to create evm client for %s: %v", evmConfig.GetName(), err)
			continue
		}

		// Set the separate database connection
		client.separateDB = evmDB

		evmClients = append(evmClients, client)
	}

	return evmClients, nil
}
