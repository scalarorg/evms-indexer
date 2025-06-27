package btc

import (
	"fmt"

	"github.com/scalarorg/evms-indexer/config"
)

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

	indexers := make([]*BtcClient, len(configs))
	for i, cfg := range configs {
		if !cfg.Enable {
			continue
		}
		indexer, err := NewBtcClient(&cfg)
		if err != nil {
			return nil, fmt.Errorf("failed to create electrum indexer %d: %w", i, err)
		}
		indexers[i] = indexer
	}

	return indexers, nil
}
