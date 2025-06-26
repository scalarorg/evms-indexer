package btc

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"

	"github.com/btcsuite/btcd/wire"
	"github.com/rs/zerolog/log"
	"github.com/scalarorg/data-models/chains"
)

func (ei *BtcClient) StartBtcIndexer(ctx context.Context) error {
	type blockWithHeight struct {
		Block  *wire.MsgBlock
		Height int64
	}
	blockChan := make(chan blockWithHeight, 10)

	// Goroutine 1: Periodically fetch new BTC blocks and send to channel
	go func() {
		defer close(blockChan)
		ticker := time.NewTicker(10 * time.Second)
		defer ticker.Stop()

		lastHeight, err := ei.GetLatestIndexedHeight(ctx)
		if err != nil {
			log.Warn().Err(err).Msg("Failed to get latest indexed height, starting from 0")
			lastHeight = 0
		}

		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				// Get latest height from network
				latestHeight, err := ei.GetLatestHeight(ctx)
				if err != nil {
					log.Warn().Err(err).Msg("Failed to get latest BTC height")
					continue
				}
				for h := lastHeight + 1; h <= latestHeight; h++ {
					block, err := ei.GetBlock(ctx, h)
					if err != nil {
						log.Warn().Err(err).Int64("height", h).Msg("Failed to fetch BTC block")
						continue
					}
					select {
					case blockChan <- blockWithHeight{Block: block, Height: h}:
						lastHeight = h
					case <-ctx.Done():
						return
					}
				}
			}
		}
	}()

	// Goroutine 2: Receive block data, parse VaultTransactions, write to DB
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case bwh, ok := <-blockChan:
				if !ok {
					return
				}
				block := bwh.Block
				blockHeight := bwh.Height
				blockHash := block.BlockHash().String()
				vaultTxs := []*chains.VaultTransaction{}
				for i, tx := range block.Transactions {
					vaultTx, err := ei.ParseVaultTransaction(tx, blockHeight, blockHash, i)
					if err != nil {
						log.Warn().Err(err).Int("txIndex", i).Msg("Failed to parse transaction")
						continue
					}
					if vaultTx != nil {
						vaultTxs = append(vaultTxs, vaultTx)
						err = ei.dbAdapter.Create(vaultTx).Error
						if err != nil {
							log.Warn().Err(err).Msg("Failed to save vault transaction")
						}
					}
				}
				log.Info().Int64("height", blockHeight).Int("vaultTxs", len(vaultTxs)).Msg("Parsed and saved vault transactions from block")
			}
		}
	}()

	return nil
}

// IndexBlocks indexes blocks by height
func (ei *BtcClient) IndexBlocks(ctx context.Context, height int64) error {
	// Get full block from BTC node
	block, err := ei.GetBlock(ctx, height)
	if err != nil {
		return fmt.Errorf("failed to get block for height %d: %w", height, err)
	}

	// Parse each transaction for VaultTransaction
	vaultTxs := []*chains.VaultTransaction{}
	for i, tx := range block.Transactions {
		// Note: block hash is not available from block header here, so we use block.BlockHash().String()
		blockHash := block.BlockHash().String()
		vaultTx, err := ei.ParseVaultTransaction(tx, height, blockHash, i)
		if err != nil {
			log.Warn().Err(err).Int("txIndex", i).Msg("Failed to parse transaction")
			continue
		}
		if vaultTx != nil {
			vaultTxs = append(vaultTxs, vaultTx)
			// Save vault transaction to database
			err = ei.dbAdapter.Create(vaultTx).Error
			if err != nil {
				log.Warn().Err(err).Msg("Failed to save vault transaction")
			}
		}
	}

	log.Info().Int64("height", height).Int("vaultTxs", len(vaultTxs)).Msg("Indexed block (transactions only)")
	return nil
}

// GetBlock retrieves a full block with transactions
func (ei *BtcClient) GetBlock(ctx context.Context, height int64) (*wire.MsgBlock, error) {
	result, err := ei.callRPCMethod(ctx, "blockchain.block.get", height)
	if err != nil {
		return nil, fmt.Errorf("failed to get block: %w", err)
	}

	var blockHex string
	err = json.Unmarshal(result, &blockHex)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal block hex: %w", err)
	}

	blockBytes, err := hex.DecodeString(blockHex)
	if err != nil {
		return nil, fmt.Errorf("failed to decode block hex: %w", err)
	}

	// Parse Bitcoin block
	var block wire.MsgBlock
	err = block.Deserialize(bytes.NewReader(blockBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to deserialize block: %w", err)
	}

	return &block, nil
}

// ParseVaultTransaction parses a transaction to check if it's a VaultTransaction
func (ei *BtcClient) ParseVaultTransaction(tx *wire.MsgTx, blockHeight int64, blockHash string, txPosition int) (*chains.VaultTransaction, error) {
	// Check if transaction has at least 2 outputs (OP_RETURN + at least one other output)
	if len(tx.TxOut) < 2 {
		return nil, nil // Not a vault transaction
	}

	// Check if first output is OP_RETURN
	firstOutput := tx.TxOut[0]
	if len(firstOutput.PkScript) < 2 || firstOutput.PkScript[0] != 0x6a {
		return nil, nil // Not a vault transaction
	}

	// Extract OP_RETURN data
	opReturnData := firstOutput.PkScript[2:] // Skip OP_RETURN and length byte

	// Check for SCALAR service tag
	if len(opReturnData) < 6 {
		return nil, nil
	}

	// Look for "SCALAR" tag (0x5343414c4152)
	scalarTag := []byte{0x53, 0x43, 0x41, 0x4c, 0x41, 0x52}
	if len(opReturnData) < len(scalarTag) {
		return nil, nil
	}

	// Find SCALAR tag in the data
	tagIndex := -1
	for i := 0; i <= len(opReturnData)-len(scalarTag); i++ {
		if bytes.Equal(opReturnData[i:i+len(scalarTag)], scalarTag) {
			tagIndex = i
			break
		}
	}

	if tagIndex == -1 {
		return nil, nil // Not a SCALAR transaction
	}

	// Parse the vault transaction data
	vaultTx, err := ei.parseVaultTransactionData(tx, opReturnData[tagIndex:], blockHeight, blockHash, txPosition)
	if err != nil {
		log.Warn().Err(err).Msg("Failed to parse vault transaction data")
		return nil, err
	}

	return vaultTx, nil
}

// parseVaultTransactionData parses the vault transaction data from OP_RETURN
func (ei *BtcClient) parseVaultTransactionData(tx *wire.MsgTx, data []byte, blockHeight int64, blockHash string, txPosition int) (*chains.VaultTransaction, error) {
	if len(data) < 6 {
		return nil, fmt.Errorf("insufficient data for vault transaction")
	}

	// Skip SCALAR tag (6 bytes)
	data = data[6:]

	if len(data) < 1 {
		return nil, fmt.Errorf("missing version byte")
	}

	version := data[0]
	data = data[1:]

	// Parse based on version
	switch version {
	case 1:
		return ei.parseVaultTransactionV1(tx, data, blockHeight, blockHash, txPosition)
	case 2:
		return ei.parseVaultTransactionV2(tx, data, blockHeight, blockHash, txPosition)
	case 3:
		return ei.parseVaultTransactionV3(tx, data, blockHeight, blockHash, txPosition)
	default:
		return nil, fmt.Errorf("unsupported vault transaction version: %d", version)
	}
}

// parseVaultTransactionV1 parses version 1 vault transaction
func (ei *BtcClient) parseVaultTransactionV1(tx *wire.MsgTx, data []byte, blockHeight int64, blockHash string, txPosition int) (*chains.VaultTransaction, error) {
	txid := tx.TxHash().String()
	rawTx, _ := ei.serializeTx(tx)

	vaultTx := &chains.VaultTransaction{
		Chain:       ei.config.SourceChain,
		BlockNumber: uint64(blockHeight),
		BlockHash:   blockHash,
		TxHash:      txid,
		TxPosition:  uint(txPosition),
		Amount:      0, // TODO: parse from outputs
		Timestamp:   0, // TODO: set from block header if needed
		ServiceTag:  "SCALAR",
		VaultTxType: 1, // Default to staking
		RawTx:       rawTx,
	}
	return vaultTx, nil
}

// parseVaultTransactionV2 parses version 2 vault transaction
func (ei *BtcClient) parseVaultTransactionV2(tx *wire.MsgTx, data []byte, blockHeight int64, blockHash string, txPosition int) (*chains.VaultTransaction, error) {
	// V2 parsing logic - similar to V1 but with additional fields
	return ei.parseVaultTransactionV1(tx, data, blockHeight, blockHash, txPosition)
}

// parseVaultTransactionV3 parses version 3 vault transaction
func (ei *BtcClient) parseVaultTransactionV3(tx *wire.MsgTx, data []byte, blockHeight int64, blockHash string, txPosition int) (*chains.VaultTransaction, error) {
	txid := tx.TxHash().String()
	rawTx, _ := ei.serializeTx(tx)

	vaultTx := &chains.VaultTransaction{
		Chain:       ei.config.SourceChain,
		BlockNumber: uint64(blockHeight),
		BlockHash:   blockHash,
		TxHash:      txid,
		TxPosition:  uint(txPosition),
		Amount:      0, // TODO: parse from outputs
		Timestamp:   0, // TODO: set from block header if needed
		ServiceTag:  "SCALAR",
		VaultTxType: 1, // Default to staking
		RawTx:       rawTx,
	}
	return vaultTx, nil
}
