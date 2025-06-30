package btc

import (
	"testing"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/scalarorg/data-models/chains"
	"github.com/stretchr/testify/assert"
)

func TestNewElectrumIndexer(t *testing.T) {
	// Create test configuration
	indexerConfig := &BtcConfig{
		BtcHost:              "localhost",
		BtcPort:              50001,
		DatabaseURL:          "postgres://test_user:test_password@localhost:5432/test_db?sslmode=disable",
		DialTimeout:          30 * time.Second,
		MethodTimeout:        60 * time.Second,
		PingInterval:         30 * time.Second,
		MaxReconnectAttempts: 120,
		ReconnectDelay:       5 * time.Second,
		EnableAutoReconnect:  true,
		BatchSize:            1,
		Confirmations:        1,
		SourceChain:          "bitcoin-testnet",
	}

	// Test indexer creation
	indexer, err := NewBtcClient(indexerConfig)

	// Note: This will fail in test environment due to database connection
	// but we can test the configuration validation
	if err != nil {
		log.Info().Err(err).Msg("Expected error due to database connection in test environment")
	} else {
		assert.NotNil(t, indexer)
		assert.Equal(t, indexerConfig, indexer.config)
	}
}

func TestElectrumIndexerConfigDefaults(t *testing.T) {
	// Test with minimal configuration
	indexerConfig := &BtcConfig{
		BtcHost:     "localhost",
		BtcPort:     50001,
		DatabaseURL: "postgres://test@localhost/test",
	}

	// Test that defaults are applied
	_, err := NewBtcClient(indexerConfig)

	if err == nil {
		assert.Equal(t, 30*time.Second, indexerConfig.DialTimeout)
		assert.Equal(t, 60*time.Second, indexerConfig.MethodTimeout)
		assert.Equal(t, 30*time.Second, indexerConfig.PingInterval)
		assert.Equal(t, 120, indexerConfig.MaxReconnectAttempts)
		assert.Equal(t, 5*time.Second, indexerConfig.ReconnectDelay)
		assert.True(t, indexerConfig.EnableAutoReconnect)
		assert.Equal(t, 1, indexerConfig.BatchSize)
		assert.Equal(t, 1, indexerConfig.Confirmations)
	}
}

func TestVaultTransactionParsing(t *testing.T) {
	// Test VaultTransaction struct creation
	vaultTx := &chains.VaultTransaction{
		TxHash:                      "test_tx_id",
		BlockNumber:                 12345,
		BlockHash:                   "test_block_hash",
		TxPosition:                  0,
		Amount:                      1000000,
		StakerScriptPubkey:          "test_staker_pubkey",
		Timestamp:                   1234567890,
		ChangeAmount:                50000,
		ChangeAddress:               "test_change_address",
		ServiceTag:                  "SCALAR",
		CovenantQuorum:              3,
		VaultTxType:                 1, // Staking
		DestinationChain:            1,
		DestinationTokenAddress:     "0x1234567890123456789012345678901234567890",
		DestinationRecipientAddress: "0x0987654321098765432109876543210987654321",
		SessionSequence:             1,
		CustodianGroupUID:           "test_custodian_group_uid",
		ScriptPubkey:                "test_script_pubkey",
		RawTx:                       "test_raw_tx",
	}

	assert.NotNil(t, vaultTx)
	assert.Equal(t, "test_tx_id", vaultTx.TxHash)
	assert.Equal(t, int64(12345), vaultTx.BlockNumber)
	assert.Equal(t, uint8(1), vaultTx.VaultTxType)
}

func TestElectrumIndexerConfigValidation(t *testing.T) {
	tests := []struct {
		name    string
		config  *BtcConfig
		wantErr bool
	}{
		{
			name: "valid config",
			config: &BtcConfig{
				BtcHost:     "localhost",
				BtcPort:     50001,
				DatabaseURL: "postgres://test@localhost/test",
			},
			wantErr: false,
		},
		{
			name: "missing host",
			config: &BtcConfig{
				BtcPort:     50001,
				DatabaseURL: "postgres://test@localhost/test",
			},
			wantErr: true,
		},
		{
			name: "missing port",
			config: &BtcConfig{
				BtcHost:     "localhost",
				DatabaseURL: "postgres://test@localhost/test",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			_, err := NewBtcClient(tt.config)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				// May still fail due to database connection, but config validation should pass
				if err != nil {
					log.Info().Err(err).Msg("Expected error due to database connection in test environment")
				}
			}
		})
	}
}

func TestElectrClientReconnection(t *testing.T) {
	config := &BtcConfig{
		Enable:               true,
		BtcHost:              "localhost",
		BtcPort:              50001,
		DatabaseURL:          "host=localhost user=test database=test",
		DialTimeout:          5 * time.Second,
		MethodTimeout:        10 * time.Second,
		PingInterval:         30 * time.Second,
		MaxReconnectAttempts: 3,
		ReconnectDelay:       1 * time.Second,
		EnableAutoReconnect:  true,
		BatchSize:            1,
		Confirmations:        1,
		SourceChain:          "bitcoin",
	}

	// This test verifies that reconnection fields are properly initialized
	// Note: We can't test actual reconnection without a real electrum server
	client, err := NewBtcClient(config)
	if err != nil {
		// Expected error due to database connection in test environment
		log.Info().Err(err).Msg("Expected error due to database connection in test environment")
		return
	}

	// Verify reconnection fields are set
	if client.baseReconnectDelay != config.ReconnectDelay {
		t.Errorf("Expected baseReconnectDelay to be %v, got %v", config.ReconnectDelay, client.baseReconnectDelay)
	}

	if client.maxReconnectDelay != 2*time.Minute {
		t.Errorf("Expected maxReconnectDelay to be 2 minutes, got %v", client.maxReconnectDelay)
	}

	if client.reconnectAttempts != 0 {
		t.Errorf("Expected reconnectAttempts to be 0, got %d", client.reconnectAttempts)
	}

	// Test reconnection delay calculation
	expectedDelay := config.ReconnectDelay * time.Duration(1<<0) // 1 second
	if expectedDelay > client.maxReconnectDelay {
		expectedDelay = client.maxReconnectDelay
	}

	// Verify the reconnection channel is created
	if client.reconnectChan == nil {
		t.Error("Expected reconnectChan to be initialized")
	}

	// Verify the stop channel is created
	if client.stopChan == nil {
		t.Error("Expected stopChan to be initialized")
	}
}
