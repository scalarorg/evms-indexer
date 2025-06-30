package btc

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseVaultReturnTxOutput_Staking(t *testing.T) {
	// Test data for staking transaction
	// SCALAR tag (6 bytes) + version (1) + network_id (1) + flags (1) + service_tag (5) + custodian_quorum (1) + dest_chain (9) + dest_token_addr (20) + dest_recipient_addr (20)
	testData := []byte{
		// SCALAR tag: 0x5343414c4152
		0x53, 0x43, 0x41, 0x4c, 0x41, 0x52,
		// Version: 1
		0x01,
		// Network ID: 1 (mainnet)
		0x01,
		// Flags: 0 (staking)
		0x00,
		// Service tag: "VAULT" (5 bytes)
		0x56, 0x41, 0x55, 0x4c, 0x54,
		// Custodian quorum: 3
		0x03,
		// Destination chain: EVM (1) + Chain ID 1 (Ethereum mainnet)
		0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01,
		// Destination token address: 0x1234567890123456789012345678901234567890
		0x12, 0x34, 0x56, 0x78, 0x90, 0x12, 0x34, 0x56, 0x78, 0x90,
		0x12, 0x34, 0x56, 0x78, 0x90, 0x12, 0x34, 0x56, 0x78, 0x90,
		// Destination recipient address: 0xabcdefabcdefabcdefabcdefabcdefabcdefabcd
		0xab, 0xcd, 0xef, 0xab, 0xcd, 0xef, 0xab, 0xcd, 0xef, 0xab,
		0xcd, 0xef, 0xab, 0xcd, 0xef, 0xab, 0xcd, 0xef, 0xab, 0xcd,
	}

	client := &BtcClient{
		config: &BtcConfig{
			VaultTag:     "SCALAR",
			VaultVersion: 1,
		},
	}

	result, err := client.ParseVaultReturnTxOutput(testData)
	require.NoError(t, err)
	require.NotNil(t, result)

	// Verify the parsed data
	assert.Equal(t, [6]byte{0x53, 0x43, 0x41, 0x4c, 0x41, 0x52}, result.Tag)
	assert.Equal(t, uint8(1), result.Version)
	assert.Equal(t, uint8(1), result.NetworkID)
	assert.Equal(t, uint8(0), result.Flags)
	assert.Equal(t, [5]byte{0x56, 0x41, 0x55, 0x4c, 0x54}, result.ServiceTag)
	assert.Equal(t, VaultReturnTxOutputTypeStaking, result.TransactionType)
	assert.Equal(t, uint8(3), result.CustodianQuorum)
	assert.Equal(t, uint8(1), result.DestChainType)
	assert.Equal(t, uint64(1), result.DestChainID)
	assert.Equal(t, "0x1234567890123456789012345678901234567890", result.DestinationTokenAddress.String())
	assert.Equal(t, "0xabcdefabcdefabcdefabcdefabcdefabcdefabcd", result.DestinationRecipientAddress.String())
}

func TestParseVaultReturnTxOutput_PoolRedeem(t *testing.T) {
	// Test data for pool redeem transaction
	// SCALAR tag (6) + version (1) + network_id (1) + flags (1) + service_tag (5) + session_sequence (8) + custodian_group_uid (32)
	testData := []byte{
		// SCALAR tag: 0x5343414c4152
		0x53, 0x43, 0x41, 0x4c, 0x41, 0x52,
		// Version: 1
		0x01,
		// Network ID: 1 (mainnet)
		0x01,
		// Flags: 0b01000001 (pool redeem)
		0x41,
		// Service tag: "VAULT" (5 bytes)
		0x56, 0x41, 0x55, 0x4c, 0x54,
		// Session sequence: 12345
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x30, 0x39,
		// Custodian group UID: 32 bytes
		0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88,
		0x99, 0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff, 0x00,
		0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88,
		0x99, 0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff, 0x00,
	}

	client := &BtcClient{
		config: &BtcConfig{
			VaultTag:     "SCALAR",
			VaultVersion: 1,
		},
	}

	result, err := client.ParseVaultReturnTxOutput(testData)
	require.NoError(t, err)
	require.NotNil(t, result)

	// Verify the parsed data
	assert.Equal(t, [6]byte{0x53, 0x43, 0x41, 0x4c, 0x41, 0x52}, result.Tag)
	assert.Equal(t, uint8(1), result.Version)
	assert.Equal(t, uint8(1), result.NetworkID)
	assert.Equal(t, uint8(0x41), result.Flags)
	assert.Equal(t, [5]byte{0x56, 0x41, 0x55, 0x4c, 0x54}, result.ServiceTag)
	assert.Equal(t, VaultReturnTxOutputTypeRedeem, result.TransactionType)
	assert.Equal(t, uint64(12345), result.SessionSequence)

	expectedUID := [32]byte{
		0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88,
		0x99, 0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff, 0x00,
		0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88,
		0x99, 0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff, 0x00,
	}
	assert.Equal(t, expectedUID, result.CustodianGroupUID)
}

func TestParseVaultReturnTxOutput_UpcRedeem(t *testing.T) {
	// Test data for UPC redeem transaction
	// SCALAR tag (6) + version (1) + network_id (1) + flags (1) + service_tag (5)
	testData := []byte{
		// SCALAR tag: 0x5343414c4152
		0x53, 0x43, 0x41, 0x4c, 0x41, 0x52,
		// Version: 1
		0x01,
		// Network ID: 1 (mainnet)
		0x01,
		// Flags: 0b10000001 (UPC redeem)
		0x81,
		// Service tag: "VAULT" (5 bytes)
		0x56, 0x41, 0x55, 0x4c, 0x54,
	}

	client := &BtcClient{
		config: &BtcConfig{
			VaultTag:     "SCALAR",
			VaultVersion: 1,
		},
	}

	result, err := client.ParseVaultReturnTxOutput(testData)
	require.NoError(t, err)
	require.NotNil(t, result)

	// Verify the parsed data
	assert.Equal(t, [6]byte{0x53, 0x43, 0x41, 0x4c, 0x41, 0x52}, result.Tag)
	assert.Equal(t, uint8(1), result.Version)
	assert.Equal(t, uint8(1), result.NetworkID)
	assert.Equal(t, uint8(0x81), result.Flags)
	assert.Equal(t, [5]byte{0x56, 0x41, 0x55, 0x4c, 0x54}, result.ServiceTag)
	assert.Equal(t, VaultReturnTxOutputTypeRedeem, result.TransactionType)
}

func TestParseVaultReturnTxOutput_InvalidTag(t *testing.T) {
	// Test with invalid tag
	testData := []byte{
		// Invalid tag: "INVALID"
		0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49,
		// Version: 1
		0x01,
		// Network ID: 1
		0x01,
		// Flags: 0
		0x00,
	}

	client := &BtcClient{
		config: &BtcConfig{
			VaultTag:     "SCALAR",
			VaultVersion: 1,
		},
	}

	result, err := client.ParseVaultReturnTxOutput(testData)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "vault tag mismatch")
}

func TestParseVaultReturnTxOutput_InvalidVersion(t *testing.T) {
	// Test with invalid version
	testData := []byte{
		// SCALAR tag: 0x5343414c4152
		0x53, 0x43, 0x41, 0x4c, 0x41, 0x52,
		// Version: 2 (invalid)
		0x02,
		// Network ID: 1
		0x01,
		// Flags: 0
		0x00,
	}

	client := &BtcClient{
		config: &BtcConfig{
			VaultTag:     "SCALAR",
			VaultVersion: 1,
		},
	}

	result, err := client.ParseVaultReturnTxOutput(testData)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "vault version mismatch")
}

func TestParseVaultReturnTxOutput_InsufficientData(t *testing.T) {
	// Test with insufficient data
	testData := []byte{
		// SCALAR tag: 0x5343414c4152
		0x53, 0x43, 0x41, 0x4c, 0x41, 0x52,
		// Version: 1
		0x01,
		// Network ID: 1
		0x01,
		// Flags: 0
		0x00,
		// Missing service tag and other required data
	}

	client := &BtcClient{
		config: &BtcConfig{
			VaultTag:     "SCALAR",
			VaultVersion: 1,
		},
	}

	result, err := client.ParseVaultReturnTxOutput(testData)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "insufficient data")
}

func TestDestinationTokenAddress_String(t *testing.T) {
	addr := DestinationTokenAddress{
		0x12, 0x34, 0x56, 0x78, 0x90, 0x12, 0x34, 0x56, 0x78, 0x90,
		0x12, 0x34, 0x56, 0x78, 0x90, 0x12, 0x34, 0x56, 0x78, 0x90,
	}

	expected := "0x1234567890123456789012345678901234567890"
	assert.Equal(t, expected, addr.String())
}

func TestDestinationRecipientAddress_String(t *testing.T) {
	addr := DestinationRecipientAddress{
		0xab, 0xcd, 0xef, 0xab, 0xcd, 0xef, 0xab, 0xcd, 0xef, 0xab,
		0xcd, 0xef, 0xab, 0xcd, 0xef, 0xab, 0xcd, 0xef, 0xab, 0xcd,
	}

	expected := "0xabcdefabcdefabcdefabcdefabcdefabcdefabcd"
	assert.Equal(t, expected, addr.String())
}

func TestScriptBuf_String(t *testing.T) {
	script := ScriptBuf{0x12, 0x34, 0x56, 0x78, 0x9a, 0xbc, 0xde, 0xf0}
	expected := "123456789abcdef0"
	assert.Equal(t, expected, script.String())
}

func TestVaultTransactionType_String(t *testing.T) {
	assert.Equal(t, "staking", VaultTxTypeStaking.String())
	assert.Equal(t, "redeem", VaultTxTypeRedeem.String())
	assert.Equal(t, "unknown", VaultTransactionType(99).String())
}

func TestVaultReturnTxOutputType_String(t *testing.T) {
	assert.Equal(t, "redeem", VaultReturnTxOutputTypeRedeem.String())
	assert.Equal(t, "staking", VaultReturnTxOutputTypeStaking.String())
	assert.Equal(t, "unknown", VaultReturnTxOutputType(99).String())
}
