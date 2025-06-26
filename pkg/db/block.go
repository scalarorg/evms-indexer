package db

import (
	"github.com/scalarorg/data-models/chains"
)

func (db *DatabaseAdapter) FindBlockHeader(chainId string, blockNumber uint64) (*chains.BlockHeader, error) {
	var blockHeader chains.BlockHeader
	result := db.PostgresClient.Where("chain = ? AND block_number = ?", chainId, blockNumber).First(&blockHeader)
	if result.Error != nil {
		return nil, result.Error
	}
	return &blockHeader, nil
}

func (db *DatabaseAdapter) CreateBlockHeader(blockHeader *chains.BlockHeader) error {
	return db.PostgresClient.Create(blockHeader).Error
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

// GetLatestBlockNumber returns the latest block number for a given chain
func (db *DatabaseAdapter) GetLatestBlockNumber(chainId string) (uint64, error) {
	var blockHeader chains.BlockHeader
	result := db.PostgresClient.Where("chain = ?", chainId).Order("block_number DESC").First(&blockHeader)
	if result.Error != nil {
		return 0, result.Error
	}
	return blockHeader.BlockNumber, nil
}

// GetLatestBlockFromAllEvents returns the latest block number from all event tables
func (db *DatabaseAdapter) GetLatestBlockFromAllEvents(chainId string) (uint64, error) {
	var maxBlock uint64 = 0

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
