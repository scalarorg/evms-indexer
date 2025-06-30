package btc

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"

	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/wire"
	"github.com/rs/zerolog/log"
	"github.com/scalarorg/data-models/chains"
	"github.com/scalarorg/go-common/chain"
)

// ============================================================================
// Main Entry Points for Transaction Parsing
// ============================================================================

// ParseVaultTxRawResult parses a btcjson.TxRawResult transaction
func (c *BtcClient) ParseVaultTxRawResult(tx *btcjson.TxRawResult, block *btcjson.GetBlockVerboseResult, txPosition int) (*chains.VaultTransaction, error) {
	// Convert btcjson.TxRawResult to wire.MsgTx
	msgTx, err := c.convertTxRawResultToMsgTx(tx)
	if err != nil {
		return nil, fmt.Errorf("failed to convert TxRawResult to MsgTx: %w", err)
	}

	// Use the existing ParseVaultMsgTx logic
	vaultTx, err := c.ParseVaultMsgTx(msgTx, txPosition, block.Height, block.Hash, block.Time)
	if err != nil {
		return nil, fmt.Errorf("failed to parse vault transaction: %w", err)
	}
	return vaultTx, nil
}

// ParseVaultMsgTx parses a wire.MsgTx transaction to check if it's a VaultTransaction
func (c *BtcClient) ParseVaultMsgTx(tx *wire.MsgTx, txPosition int, blockHeight int64, blockHash string, blockTime int64) (*chains.VaultTransaction, error) {
	// Check if transaction has at least 2 outputs (OP_RETURN + at least one other output)
	if len(tx.TxOut) < 2 {
		return nil, fmt.Errorf("transaction has less than 2 outputs")
	}

	// Check if first output is OP_RETURN
	firstOutput := tx.TxOut[0]
	if len(firstOutput.PkScript) < 2 || firstOutput.PkScript[0] != 0x6a {
		return nil, nil // Not a vault transaction
	}

	// Extract OP_RETURN data
	opReturnData := firstOutput.PkScript[2:] // Skip OP_RETURN and length byte
	vaultReturnTxOutput, err := c.ParseVaultReturnTxOutput(opReturnData)
	if err != nil {
		return nil, fmt.Errorf("failed to parse vault return transaction output: %w", err)
	}

	// Parse the vault transaction data
	vaultTx, err := c.parseVaultTransactionData(tx, vaultReturnTxOutput, txPosition, blockHeight, blockHash, blockTime)
	if err != nil {
		log.Warn().Err(err).Msg("Failed to parse vault transaction data")
		return nil, err
	}

	return vaultTx, nil
}

// ============================================================================
// Vault Transaction Data Parsing
// ============================================================================

// parseVaultTransactionData parses the vault transaction data from OP_RETURN
func (c *BtcClient) parseVaultTransactionData(tx *wire.MsgTx, vaultReturnTxOutput *VaultReturnTxOutput,
	txPosition int, blockHeight int64, blockHash string, blockTime int64) (*chains.VaultTransaction, error) {
	// Parse the vault transaction data using the new parser
	// Serialize the transaction for RawTx field
	rawTx, err := c.serializeTx(tx)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize transaction: %w", err)
	}

	// Convert to chains.VaultTransaction
	vaultTx := &chains.VaultTransaction{
		Chain:                       c.config.SourceChain,
		BlockNumber:                 uint64(blockHeight),
		BlockHash:                   blockHash,
		TxHash:                      tx.TxHash().String(),
		TxPosition:                  uint(txPosition),
		Timestamp:                   uint64(blockTime),
		RawTx:                       rawTx, // Will be set below
		ServiceTag:                  string(vaultReturnTxOutput.Tag[:]),
		VaultTxType:                 uint8(vaultReturnTxOutput.TransactionType),
		CovenantQuorum:              vaultReturnTxOutput.CustodianQuorum,
		DestinationChain:            vaultReturnTxOutput.DestChainID,
		DestinationTokenAddress:     vaultReturnTxOutput.DestinationTokenAddress.String(),
		DestinationRecipientAddress: vaultReturnTxOutput.DestinationRecipientAddress.String(),
		SessionSequence:             vaultReturnTxOutput.SessionSequence,
		CustodianGroupUID:           hex.EncodeToString(vaultReturnTxOutput.CustodianGroupUID[:]),
		ScriptPubkey:                hex.EncodeToString(vaultReturnTxOutput.ScriptPubkey),
	}

	if vaultReturnTxOutput.TransactionType == VaultReturnTxOutputTypeStaking {
		if len(tx.TxOut) >= 2 {
			vaultTx.Amount = uint64(tx.TxOut[1].Value)
		}
		if len(tx.TxOut) >= 3 {
			vaultTx.ChangeAmount = uint64(tx.TxOut[2].Value)
			vaultTx.ChangeAddress = hex.EncodeToString(tx.TxOut[2].PkScript)
		}
		if len(tx.TxIn) >= 1 {
			previousTxout, err := c.getPreviousTxout(&tx.TxIn[0].PreviousOutPoint)
			if err != nil {
				return nil, fmt.Errorf("failed to get previous transaction output: %w", err)
			}
			vaultTx.StakerScriptPubkey = hex.EncodeToString(previousTxout.PkScript)
		}
	}

	return vaultTx, nil
}

func (c *BtcClient) getPreviousTxout(output *wire.OutPoint) (*wire.TxOut, error) {
	tx, err := c.rpcClient.GetRawTransaction(&output.Hash)
	if err != nil {
		return nil, fmt.Errorf("failed to get raw transaction: %w", err)
	}
	msgTx := tx.MsgTx()
	if len(msgTx.TxOut) <= int(output.Index) {
		return nil, fmt.Errorf("output index out of range")
	}
	return msgTx.TxOut[output.Index], nil
}

// ============================================================================
// Vault Return Transaction Output Parsing
// ============================================================================

// ParseVaultReturnTxOutput parses a VaultReturnTxOutput from bytes
func (c *BtcClient) ParseVaultReturnTxOutput(data []byte) (*VaultReturnTxOutput, error) {
	// if len(data) < 32+1+1+1+6+1+1+1+8+20+20+1+8+32 {
	// 	return nil, fmt.Errorf("insufficient data for VaultReturnTxOutput")
	// }
	vaultReturnTxOutput := &VaultReturnTxOutput{}
	offset := 0

	// Parse tag (32 bytes)
	var tag [SCALAR_TAG_SIZE]byte
	copy(tag[:], data[offset:offset+SCALAR_TAG_SIZE])
	//if !bytes.Equal(tag[:], []byte{0x53, 0x43, 0x41, 0x4c, 0x41, 0x52}) {
	if !bytes.Equal(tag[:], []byte(c.config.VaultTag)) {
		return nil, fmt.Errorf("vault tag mismatch: expected %s, got %s", c.config.VaultTag, tag)
	}
	offset += SCALAR_TAG_SIZE
	vaultReturnTxOutput.Tag = tag
	// Parse version (1 byte)
	vaultReturnTxOutput.Version = data[offset]
	offset++
	if vaultReturnTxOutput.Version != c.config.VaultVersion {
		return nil, fmt.Errorf("vault version mismatch: expected %d, got %d", c.config.VaultVersion, vaultReturnTxOutput.Version)
	}
	// Parse network_id (1 byte)
	vaultReturnTxOutput.NetworkID = data[offset]
	offset++

	// Parse flags (1 byte)
	vaultReturnTxOutput.Flags = data[offset]
	offset++
	//Parse vault output based on flags
	switch vaultReturnTxOutput.Flags {
	case 0b01000001:
		return c.parsePoolRedeemVaultReturnTxOutput(data[offset:], vaultReturnTxOutput)
	case 0b10000001:
		return c.parseUpcRedeemVaultReturnTxOutput(data[offset:], vaultReturnTxOutput)
	default:
		return c.parseStakingVaultReturnTxOuput(data[offset:], vaultReturnTxOutput)
	}
}

func (c *BtcClient) parsePoolRedeemVaultReturnTxOutput(data []byte, vaultReturnTxOutput *VaultReturnTxOutput) (*VaultReturnTxOutput, error) {
	vaultReturnTxOutput.TransactionType = VaultReturnTxOutputTypeRedeem
	if len(data) < SERVICE_TAG_SIZE+SESSION_SEQUENCE_SIZE+HASH_SIZE {
		return nil, fmt.Errorf("insufficient data for pool redeem vault return transaction output")
	}
	offset := 0
	// Parse service_tag (6 bytes)
	var serviceTag [SERVICE_TAG_SIZE]byte
	copy(serviceTag[:], data[offset:offset+SERVICE_TAG_SIZE])
	offset += SERVICE_TAG_SIZE
	vaultReturnTxOutput.ServiceTag = serviceTag

	vaultReturnTxOutput.SessionSequence = binary.BigEndian.Uint64(data[offset : offset+SESSION_SEQUENCE_SIZE])
	offset += SESSION_SEQUENCE_SIZE

	vaultReturnTxOutput.CustodianGroupUID = [HASH_SIZE]byte{}
	copy(vaultReturnTxOutput.CustodianGroupUID[:], data[offset:offset+HASH_SIZE])
	offset += HASH_SIZE

	return vaultReturnTxOutput, nil
}

func (c *BtcClient) parseUpcRedeemVaultReturnTxOutput(data []byte, vaultReturnTxOutput *VaultReturnTxOutput) (*VaultReturnTxOutput, error) {
	if len(data) < SERVICE_TAG_SIZE {
		return nil, fmt.Errorf("insufficient data for upc redeem vault return transaction output")
	}
	vaultReturnTxOutput.TransactionType = VaultReturnTxOutputTypeRedeem
	offset := 0
	// Parse service_tag (6 bytes)
	vaultReturnTxOutput.ServiceTag = [SERVICE_TAG_SIZE]byte{}
	copy(vaultReturnTxOutput.ServiceTag[:], data[offset:offset+SERVICE_TAG_SIZE])
	offset += SERVICE_TAG_SIZE
	return vaultReturnTxOutput, nil
}

func (c *BtcClient) parseStakingVaultReturnTxOuput(data []byte, vaultReturnTxOutput *VaultReturnTxOutput) (*VaultReturnTxOutput, error) {
	vaultReturnTxOutput.TransactionType = VaultReturnTxOutputTypeStaking
	offset := 0
	if len(data) < SERVICE_TAG_SIZE+CUSTODIAN_QUORUM_SIZE+DEST_CHAIN_SIZE+DEST_TOKEN_ADDR_SIZE+DEST_RECIPIENT_ADDR_SIZE {
		return nil, fmt.Errorf("insufficient data for staking vault return transaction output")
	}
	vaultReturnTxOutput.ServiceTag = [SERVICE_TAG_SIZE]byte{}
	copy(vaultReturnTxOutput.ServiceTag[:], data[offset:offset+SERVICE_TAG_SIZE])
	offset += SERVICE_TAG_SIZE

	// Parse custodian_quorum (1 byte)
	vaultReturnTxOutput.CustodianQuorum = data[offset]
	offset++

	chainInfo := chain.NewChainInfoFromBytes(data[offset : offset+DEST_CHAIN_SIZE])
	if chainInfo == nil {
		return nil, fmt.Errorf("failed to parse chain info")
	}
	vaultReturnTxOutput.DestChainType = uint8(chainInfo.ChainType)
	vaultReturnTxOutput.DestChainID = chainInfo.ChainID
	offset += 9

	// Parse destination_token_address (20 bytes)
	vaultReturnTxOutput.DestinationTokenAddress = DestinationTokenAddress{}
	copy(vaultReturnTxOutput.DestinationTokenAddress[:], data[offset:offset+DEST_TOKEN_ADDR_SIZE])
	offset += DEST_TOKEN_ADDR_SIZE

	// Parse destination_recipient_address (20 bytes)
	vaultReturnTxOutput.DestinationRecipientAddress = DestinationRecipientAddress{}
	copy(vaultReturnTxOutput.DestinationRecipientAddress[:], data[offset:offset+DEST_RECIPIENT_ADDR_SIZE])
	offset += DEST_RECIPIENT_ADDR_SIZE

	return vaultReturnTxOutput, nil
}

// ============================================================================
// Utility Functions
// ============================================================================

// convertTxRawResultToMsgTx converts btcjson.TxRawResult to wire.MsgTx
func (c *BtcClient) convertTxRawResultToMsgTx(tx *btcjson.TxRawResult) (*wire.MsgTx, error) {
	// Decode the hex transaction
	txBytes, err := hex.DecodeString(tx.Hex)
	if err != nil {
		return nil, fmt.Errorf("failed to decode transaction hex: %w", err)
	}

	// Deserialize into wire.MsgTx
	var msgTx wire.MsgTx
	err = msgTx.Deserialize(bytes.NewReader(txBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to deserialize transaction: %w", err)
	}

	return &msgTx, nil
}

// serializeTx serializes a transaction to hex string
func (c *BtcClient) serializeTx(tx *wire.MsgTx) (string, error) {
	var buf bytes.Buffer
	err := tx.Serialize(&buf)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(buf.Bytes()), nil
}
