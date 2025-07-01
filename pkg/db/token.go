package db

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/scalarorg/data-models/chains"
	"github.com/scalarorg/data-models/scalarnet"
	"gorm.io/gorm"
)

// find relay datas by token sent attributes
func (db *DatabaseAdapter) FindPendingBtcTokenSent(sourceChain string, height int) ([]*chains.TokenSent, error) {
	var tokenSents []*chains.TokenSent
	result := db.PostgresClient.
		Where("source_chain = ? AND block_number <= ?",
			sourceChain,
			height).
		Where("status = ?", string(chains.TokenSentStatusPending)).
		Find(&tokenSents)

	if result.Error != nil {
		return tokenSents, fmt.Errorf("FindPendingBtcTokenSent with error: %w", result.Error)
	}
	if len(tokenSents) == 0 {
		log.Warn().
			Str("sourceChain", sourceChain).
			Msgf("[DatabaseAdapter] [FindPendingBtcTokenSent] no token sent with block height before %d found", height)
	}
	return tokenSents, nil
}

func (db *DatabaseAdapter) SaveTokenSentsAndRemoveDupplicates(tokenSents []*chains.TokenSent) error {
	tx := db.PostgresClient.Begin()
	if tx == nil {
		return fmt.Errorf("failed to begin transaction")
	}
	defer tx.Rollback() // Will be ignored if transaction is committed

	// Delete existing verifying entries with matching tx_hashes
	txHashes := make([]string, 0, len(tokenSents))
	for _, tokenSent := range tokenSents {
		txHashes = append(txHashes, tokenSent.TxHash)
	}

	// Currently too much places using of ChainToken without filtering delete_at != NULL => so we need to hard delete, consider using soft delete

	// if err := tx.Where("tx_hash IN ? AND status = ?", txHashes, chains.TokenSentStatusPending).
	// 	Updates(map[string]interface{}{"status": chains.TokenSentStatusDeleted, "deleted_at": time.Now()}).Error; err != nil {
	// 	return fmt.Errorf("failed to remove duplicate token sents: %w", err)
	// }

	if err := tx.Where("tx_hash IN ? AND status = ?", txHashes, chains.TokenSentStatusPending).
		Delete(&chains.TokenSent{}).Error; err != nil {
		return fmt.Errorf("failed to remove duplicate token sents: %w", err)
	}

	// Save new token sents
	if err := tx.Save(tokenSents).Error; err != nil {
		return fmt.Errorf("failed to save new token sents: %w", err)
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (db *DatabaseAdapter) SaveTokenSents(tokenSents []*chains.TokenSent) error {
	result := db.PostgresClient.Save(tokenSents)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (db *DatabaseAdapter) SaveTokenSent(tokenSent chains.TokenSent) error {
	result := db.PostgresClient.Save(&tokenSent)
	if result.Error != nil {
		return fmt.Errorf("failed to create evm token send: %w", result.Error)
	}
	return nil
}

func (db *DatabaseAdapter) UpdateTokenSentsStatus(ctx context.Context, cmdIds []string, status chains.TokenSentStatus) error {
	log.Debug().Any("cmdIds", cmdIds).Msg("[DatabaseAdapter] UpdateTokenSentsStatus")
	err := db.PostgresClient.Transaction(func(tx *gorm.DB) error {
		eventIds := tx.Model(&scalarnet.TokenSentApproved{}).Select("event_id").Where("command_id IN (?)", cmdIds)
		tokenSents := []*chains.TokenSent{}
		//only update the token sent that is not success
		result := tx.Model(&chains.TokenSent{}).Where("event_id IN (?) and status != ?", eventIds, chains.TokenSentStatusSuccess).Find(&tokenSents)
		if result.Error == nil {
			ids := []string{}
			for _, tokenSent := range tokenSents {
				ids = append(ids, tokenSent.EventID)
			}
			result = tx.Model(&chains.TokenSent{}).Where("event_id IN (?)", ids).Update("status", status)
		}
		return result.Error
	})
	return err
}

func (db *DatabaseAdapter) SaveTokenDeployed(tokenDeployed *chains.TokenDeployed) error {
	result := db.PostgresClient.Save(tokenDeployed)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// BatchSaveTokenSents saves multiple token sent records in a single transaction
func (db *DatabaseAdapter) BatchSaveTokenSents(tokenSents []chains.TokenSent) error {
	if len(tokenSents) == 0 {
		return nil
	}

	result := db.PostgresClient.Create(&tokenSents)
	if result.Error != nil {
		return fmt.Errorf("failed to batch save token sents: %w", result.Error)
	}
	return nil
}

// BatchSaveTokenDeployed saves multiple token deployed records in a single transaction
func (db *DatabaseAdapter) BatchSaveTokenDeployed(tokenDeployed []*chains.TokenDeployed) error {
	if len(tokenDeployed) == 0 {
		return nil
	}

	result := db.PostgresClient.Create(&tokenDeployed)
	if result.Error != nil {
		return fmt.Errorf("failed to batch save token deployed: %w", result.Error)
	}
	return nil
}

// BatchSaveCommandExecuted saves multiple command executed records in a single transaction
func (db *DatabaseAdapter) BatchSaveCommandExecuted(commandExecuted []chains.CommandExecuted) error {
	if len(commandExecuted) == 0 {
		return nil
	}

	result := db.PostgresClient.Create(&commandExecuted)
	if result.Error != nil {
		return fmt.Errorf("failed to batch save command executed: %w", result.Error)
	}
	return nil
}

// BatchSaveSwitchedPhase saves multiple switched phase records in a single transaction
func (db *DatabaseAdapter) BatchSaveSwitchedPhase(switchedPhase []chains.SwitchedPhase) error {
	if len(switchedPhase) == 0 {
		return nil
	}

	result := db.PostgresClient.Create(&switchedPhase)
	if result.Error != nil {
		return fmt.Errorf("failed to batch save switched phase: %w", result.Error)
	}
	return nil
}

// BatchSaveEvmRedeemTx saves multiple EVM redeem transaction records in a single transaction
func (db *DatabaseAdapter) BatchSaveEvmRedeemTx(evmRedeemTx []chains.EvmRedeemTx) error {
	if len(evmRedeemTx) == 0 {
		return nil
	}

	result := db.PostgresClient.Create(&evmRedeemTx)
	if result.Error != nil {
		return fmt.Errorf("failed to batch save EVM redeem tx: %w", result.Error)
	}
	return nil
}
