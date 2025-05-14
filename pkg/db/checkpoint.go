package db

import (
	"fmt"

	"github.com/scalarorg/data-models/scalarnet"
	"gorm.io/gorm"
)

// Get earliest event check point for all eventNames
func (db *DatabaseAdapter) GetLastCheckPoint(chainName string) (*scalarnet.EventCheckPoint, error) {
	//Default value
	lastBlock := scalarnet.EventCheckPoint{
		ChainName:   chainName,
		EventName:   "",
		BlockNumber: 0,
		TxHash:      "",
		LogIndex:    0,
		EventKey:    "",
	}
	result := db.PostgresClient.Where("chain_name = ?", chainName).Order("block_number ASC").First(&lastBlock)
	return &lastBlock, result.Error
}

func (db *DatabaseAdapter) GetLastEventCheckPoint(chainName, eventName string, fromBlock uint64) (*scalarnet.EventCheckPoint, error) {
	//Default value
	lastBlock := scalarnet.EventCheckPoint{
		ChainName:   chainName,
		EventName:   eventName,
		BlockNumber: fromBlock,
		TxHash:      "",
		LogIndex:    0,
		EventKey:    "",
	}
	result := db.PostgresClient.Where("chain_name = ? AND event_name = ?", chainName, eventName).First(&lastBlock)
	return &lastBlock, result.Error
}
func (db *DatabaseAdapter) UpdateLastEventCheckPoints(checkpoints map[string]scalarnet.EventCheckPoint) error {
	return db.PostgresClient.Transaction(func(tx *gorm.DB) error {
		for _, checkpoint := range checkpoints {
			err := UpdateLastEventCheckPoint(tx, &checkpoint)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func (db *DatabaseAdapter) UpdateLastEventCheckPoint(value *scalarnet.EventCheckPoint) error {
	return UpdateLastEventCheckPoint(db.PostgresClient, value)
}

// For transactional update
func UpdateLastEventCheckPoint(db *gorm.DB, value *scalarnet.EventCheckPoint) error {
	//1. Get event check point from db
	storedEventCheckPoint := scalarnet.EventCheckPoint{}
	err := db.Where("chain_name = ? AND event_name = ?", value.ChainName, value.EventName).First(&storedEventCheckPoint).Error
	if err != nil {
		err = db.Create(value).Error
	} else {
		err = db.Model(&scalarnet.EventCheckPoint{}).Where("chain_name = ? AND event_name = ?", value.ChainName, value.EventName).Updates(map[string]interface{}{
			"block_number": value.BlockNumber,
			"tx_hash":      value.TxHash,
			"log_index":    value.LogIndex,
			"event_key":    value.EventKey,
		}).Error
	}

	// err := db.Clauses(
	// 	clause.OnConflict{
	// 		Columns: []clause.Column{{Name: "chain_name"}, {Name: "event_name"}},
	// 		DoUpdates: clause.Assignments(map[string]interface{}{
	// 			"block_number": value.BlockNumber,
	// 			"tx_hash":      utils.NormalizeHash(value.TxHash),
	// 			"log_index":    value.LogIndex,
	// 			"event_key":    value.EventKey,
	// 		}),
	// 	},
	// ).Create(value).Error
	if err != nil {
		return fmt.Errorf("failed to update last event check point: %w", err)
	}

	return nil
}
