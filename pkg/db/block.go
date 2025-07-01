package db

import (
	"context"
	"encoding/hex"
	"fmt"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/scalarorg/data-models/chains"
	"gorm.io/gorm/clause"
)

// BlockHeaderLite is a minimal block header for reorg checks
type BlockHeaderLite struct {
	Height   int64
	Hash     *chainhash.Hash
	PrevHash *chainhash.Hash
}

func (db *DatabaseAdapter) FindBlockHeader(chainId string, blockNumber uint64) (*chains.BlockHeader, error) {
	var blockHeader chains.BlockHeader
	result := db.PostgresClient.Where("chain = ? AND block_number = ?", chainId, blockNumber).First(&blockHeader)
	if result.Error != nil {
		return nil, result.Error
	}
	return &blockHeader, nil
}

func (db *DatabaseAdapter) CreateBlockHeader(blockHeader *chains.BlockHeader) error {
	return db.PostgresClient.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "chain"}, {Name: "block_number"}},
		DoNothing: true,
	}).Create(blockHeader).Error
}

func (db *DatabaseAdapter) CreateBtcBlockHeader(blockHeader *chains.BtcBlockHeader) error {
	return db.PostgresClient.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "hash"}, {Name: "height"}},
		UpdateAll: true,
	}).Create(blockHeader).Error
}

func (db *DatabaseAdapter) GetBlockTime(chainId string, blockNumbers []uint64) (map[uint64]uint64, error) {
	var blockHeaders []*chains.BlockHeader
	result := db.PostgresClient.Where("chain = ? AND block_number IN ?", chainId, blockNumbers).Find(&blockHeaders)
	if result.Error != nil {
		return nil, result.Error
	}
	blockTimeMap := make(map[uint64]uint64)
	for _, blockHeader := range blockHeaders {
		blockTimeMap[blockHeader.BlockNumber] = blockHeader.BlockTime
	}
	return blockTimeMap, nil
}

// GetLatestBtcIndexedHeight returns the latest block height indexed in the database for BTC chains
func (db *DatabaseAdapter) GetLatestIndexedHeight(chainId string) (int64, error) {
	var header chains.BtcBlockHeader
	result := db.PostgresClient.Order("height DESC").First(&header)
	if result.Error != nil {
		return 0, result.Error
	}
	return int64(header.Height), nil
}

// GetLatestBlockFromAllEvents returns the latest block number from all event tables
func (db *DatabaseAdapter) GetLatestBlockFromAllEvents(chainId string) (uint64, error) {
	var maxBlock uint64 = 0
	if db.PostgresClient == nil {
		return 0, fmt.Errorf("database client is nil")
	}
	// Check token_sents table
	var tokenSent chains.TokenSent
	result := db.PostgresClient.Where("source_chain = ?", chainId).Order("block_number DESC").First(&tokenSent)
	if result.Error == nil && tokenSent.BlockNumber > maxBlock {
		maxBlock = tokenSent.BlockNumber
	}

	// Check contract_calls table
	var contractCall chains.ContractCall
	result = db.PostgresClient.Where("source_chain = ?", chainId).Order("block_number DESC").First(&contractCall)
	if result.Error == nil && contractCall.BlockNumber > maxBlock {
		maxBlock = contractCall.BlockNumber
	}

	// Check contract_calls_with_token table
	var contractCallWithToken chains.ContractCallWithToken
	result = db.PostgresClient.Where("source_chain = ?", chainId).Order("block_number DESC").First(&contractCallWithToken)
	if result.Error == nil && contractCallWithToken.BlockNumber > maxBlock {
		maxBlock = contractCallWithToken.BlockNumber
	}

	// Check command_executed table
	var commandExecuted chains.CommandExecuted
	result = db.PostgresClient.Where("source_chain = ?", chainId).Order("block_number DESC").First(&commandExecuted)
	if result.Error == nil && commandExecuted.BlockNumber > maxBlock {
		maxBlock = commandExecuted.BlockNumber
	}

	// Check token_deployed table
	var tokenDeployed chains.TokenDeployed
	result = db.PostgresClient.Where("source_chain = ?", chainId).Order("block_number DESC").First(&tokenDeployed)
	if result.Error == nil && tokenDeployed.BlockNumber > maxBlock {
		maxBlock = tokenDeployed.BlockNumber
	}

	// Check evm_redeem_txs table
	var evmRedeemTx chains.EvmRedeemTx
	result = db.PostgresClient.Where("source_chain = ?", chainId).Order("block_number DESC").First(&evmRedeemTx)
	if result.Error == nil && evmRedeemTx.BlockNumber > maxBlock {
		maxBlock = evmRedeemTx.BlockNumber
	}

	// Check switched_phases table
	var switchedPhase chains.SwitchedPhase
	result = db.PostgresClient.Where("source_chain = ?", chainId).Order("block_number DESC").First(&switchedPhase)
	if result.Error == nil && switchedPhase.BlockNumber > maxBlock {
		maxBlock = switchedPhase.BlockNumber
	}

	return maxBlock, nil
}

// GetBlockHashByHeight returns the block hash for a given height
func (db *DatabaseAdapter) GetBlockHashByHeight(ctx context.Context, height int64) (*chainhash.Hash, error) {
	var header chains.BtcBlockHeader
	result := db.PostgresClient.Where("height = ?", height).First(&header)
	if result.Error != nil {
		return nil, result.Error
	}

	hash, err := chainhash.NewHashFromStr(header.Hash)
	if err != nil {
		return nil, err
	}
	return hash, nil
}

// GetBlockHeaderByHeight returns a minimal block header for reorg checks
func (db *DatabaseAdapter) GetBlockHeaderByHeight(ctx context.Context, height int64) (*BlockHeaderLite, error) {
	var header chains.BtcBlockHeader
	result := db.PostgresClient.Where("height = ?", height).First(&header)
	if result.Error != nil {
		return nil, result.Error
	}

	hash, err := chainhash.NewHashFromStr(header.Hash)
	if err != nil {
		return nil, err
	}

	// Convert byte slice to hex string for chainhash.NewHashFromStr
	prevHashHex := hex.EncodeToString(header.PrevBlockhash)
	prevHash, err := chainhash.NewHashFromStr(prevHashHex)
	if err != nil {
		return nil, err
	}

	return &BlockHeaderLite{
		Height:   int64(header.Height),
		Hash:     hash,
		PrevHash: prevHash,
	}, nil
}

// DeleteBlockAndTxsFromHeight deletes blocks and their transactions from the given height onwards
func (db *DatabaseAdapter) DeleteBlockAndTxsFromHeight(ctx context.Context, height int64) error {
	// Delete vault transactions from the given height onwards
	if err := db.PostgresClient.Where("block_number >= ?", uint64(height)).Delete(&chains.VaultTransaction{}).Error; err != nil {
		return err
	}

	// Delete block headers from the given height onwards
	if err := db.PostgresClient.Where("height >= ?", height).Delete(&chains.BtcBlockHeader{}).Error; err != nil {
		return err
	}

	return nil
}
