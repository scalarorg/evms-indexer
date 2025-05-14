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
