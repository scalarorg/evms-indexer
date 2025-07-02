package db

import (
	"fmt"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/scalarorg/data-models/chains"
	"gorm.io/gorm"
)

func (db *DatabaseAdapter) UpdateContractCallApproved(messageID string, executeHash string) error {
	updateData := map[string]interface{}{
		"execute_hash": executeHash,
		"status":       chains.ContractCallStatusApproved,
	}
	record := db.PostgresClient.Model(&chains.ContractCall{}).Where("id = ?", messageID).Updates(updateData)
	if record.Error != nil {
		return record.Error
	}
	log.Info().Msgf("[DatabaseAdapter] [UpdateContractCallApproved]: RelayData[%s]", messageID)
	return nil
}
func (db *DatabaseAdapter) FindContractCallByCommnadId(commandId string) (*chains.ContractCall, error) {
	var contractCall chains.ContractCall
	result := db.PostgresClient.Where("command_id = ?", commandId).First(&contractCall)
	if result.Error != nil {
		return nil, result.Error
	}
	return &contractCall, nil
}
func (db *DatabaseAdapter) FindContractCallWithTokenPayloadByEventId(eventId string) (*chains.ContractCallWithToken, error) {
	var contractCallWithToken chains.ContractCallWithToken
	result := db.PostgresClient.Where("event_id = ?", eventId).First(&contractCallWithToken)
	if result.Error != nil {
		log.Error().Err(result.Error).Str("EventId", eventId).Msg("[DatabaseAdapter] [FindContractCallWithTokenPayloadByEventId] failed to find contract call with token by event id")
		return nil, result.Error
	}
	return &contractCallWithToken, nil
}

func (db *DatabaseAdapter) FindContractCallPayloadByHash(payloadHash string) ([]byte, error) {
	var contractCall chains.ContractCall
	log.Debug().Str("payloadHash", payloadHash).Msg("[DatabaseAdapter] [FindPayloadByHash] finding payload by hash")
	result := db.PostgresClient.
		Where("payload_hash = ?", strings.ToLower(payloadHash)).
		First(&contractCall)

	if result.Error != nil {
		return nil, fmt.Errorf("failed to find payload by hash: %w", result.Error)
	}

	return contractCall.Payload, nil
}

// Find Realaydata by ContractAddress, SourceAddress, PayloadHash
func (db *DatabaseAdapter) FindContractCallByParams(sourceAddress string, destContractAddress string, payloadHash string) ([]chains.ContractCall, error) {
	var contractCalls []chains.ContractCall
	result := db.PostgresClient.
		Joins("CallContract").
		Where("contract_address = ? AND source_address = ? AND payload_hash = ?",
			strings.ToLower(destContractAddress),
			strings.ToLower(sourceAddress),
			strings.ToLower(payloadHash)).
		Where("status IN (?)", []chains.ContractCallStatus{chains.ContractCallStatusPending, chains.ContractCallStatusApproved}).
		Preload("CallContract").
		Find(&contractCalls)

	if result.Error != nil {
		return contractCalls, fmt.Errorf("find relaydatas by contract call with error: %w", result.Error)
	}
	if len(contractCalls) == 0 {
		log.Warn().Str("contractAddress", destContractAddress).Str("sourceAddress", sourceAddress).Str("payloadHash", payloadHash).Msg("[DatabaseAdapter] [FindRelayDataByContractCall] no relaydata found")
	}
	return contractCalls, nil

}

// func (db *DatabaseAdapter) UpdateCallContractWithExecuteHash(eventId string, status chains.ContractCallStatus, executeHash *string) error {
// 	data := chains.ContractCall{
// 		Status:      status,
// 		ExecuteHash: executeHash,
// 	}
// 	updateResult := db.updateCallContract(eventId, data)
// 	if updateResult.Error != nil {
// 		return updateResult.Error
// 	}
// 	return nil
// }

func (db *DatabaseAdapter) UpdateCallContractWithTokenExecuteHash(eventId string, status chains.ContractCallStatus, executeHash string) error {
	data := chains.ContractCallWithToken{
		ContractCall: chains.ContractCall{
			Status:      status,
			ExecuteHash: executeHash,
		},
	}
	updateResult := db.updateCallContractWithToken(eventId, data)
	if updateResult.Error != nil {
		return updateResult.Error
	}
	return nil
}

func (db *DatabaseAdapter) CreateContractCall(contractCall chains.ContractCall) error {
	result := db.PostgresClient.Save(&contractCall)
	if result.Error != nil {
		return fmt.Errorf("failed to create evm contract call: %w", result.Error)
	}
	return nil
}

func (db *DatabaseAdapter) CreateContractCallWithToken(contractCallWithToken *chains.ContractCallWithToken) error {
	result := db.PostgresClient.Save(contractCallWithToken)
	if result.Error != nil {
		return fmt.Errorf("failed to create evm contract call with token: %w", result.Error)
	}
	return nil
}

func (db *DatabaseAdapter) updateCallContract(eventId string, data interface{}) (tx *gorm.DB) {
	result := db.PostgresClient.Model(&chains.ContractCall{}).Where("event_id = ?", eventId).Updates(data)
	if result.Error != nil {
		log.Error().Msgf("[DatabaseAdapter] [updateCallContract] %s: %v", eventId, result.Error)
	}
	log.Debug().Msgf("[DatabaseAdapter] [updateCallContract] %s: %v", eventId, data)
	return result
}

func (db *DatabaseAdapter) updateCallContractWithToken(eventId string, data interface{}) (tx *gorm.DB) {
	result := db.PostgresClient.Model(&chains.ContractCallWithToken{}).Where("event_id = ?", eventId).Updates(data)
	if result.Error != nil {
		log.Error().Msgf("[DatabaseAdapter] [updateCallContractWithToken] %s: %v", eventId, result.Error)
	}
	return result
}

// TODO: Find any better way to update batch relay data status
func (db *DatabaseAdapter) UpdateBatchContractCallStatus(data []ContractCallExecuteResult, batchSize int) error {
	// Handle empty data case
	if len(data) == 0 {
		return nil
	}

	// Process updates in batches
	return db.PostgresClient.Transaction(func(tx *gorm.DB) error {
		for i := 0; i < len(data); i += batchSize {
			end := i + batchSize
			if end > len(data) {
				end = len(data)
			}

			batch := data[i:end]
			for _, item := range batch {
				updates := chains.ContractCall{
					Status: item.Status,
				}

				result := tx.Model(&chains.ContractCall{}).
					Where("event_id = ?", item.EventId).
					Updates(updates)

				if result.Error != nil {
					return fmt.Errorf("failed to update contract call batch: %w", result.Error)
				}
			}
		}
		return nil
	})
}

// func (db *DatabaseAdapter) UpdateContractCallWithMintsStatus(ctx context.Context, cmdIds []string, status chains.ContractCallStatus) error {
// 	log.Debug().Any("cmdIds", cmdIds).Msg("[DatabaseAdapter] UpdateContractCallWithMintsStatus")
// 	err := db.PostgresClient.Transaction(func(tx *gorm.DB) error {
// 		eventIds := tx.Model(&scalarnet.ContractCallApprovedWithMint{}).Select("event_id").Where("command_id IN (?)", cmdIds)
// 		//only update the token sent that is not success
// 		result := tx.Model(&chains.ContractCallWithToken{}).Where("event_id IN (?) and status != ?", eventIds, chains.ContractCallStatusSuccess).Update("status", status)
// 		return result.Error
// 	})
// 	return err
// }

// BatchCreateContractCalls saves multiple contract calls in a single transaction
func (db *DatabaseAdapter) BatchCreateContractCalls(contractCalls []chains.ContractCall) error {
	if len(contractCalls) == 0 {
		return nil
	}

	result := db.PostgresClient.Create(&contractCalls)
	if result.Error != nil {
		return fmt.Errorf("failed to batch create contract calls: %w", result.Error)
	}
	return nil
}

// BatchCreateContractCallsWithToken saves multiple contract calls with token in a single transaction
func (db *DatabaseAdapter) BatchCreateContractCallsWithToken(contractCallsWithToken []*chains.ContractCallWithToken) error {
	if len(contractCallsWithToken) == 0 {
		return nil
	}

	result := db.PostgresClient.Create(&contractCallsWithToken)
	if result.Error != nil {
		return fmt.Errorf("failed to batch create contract calls with token: %w", result.Error)
	}
	return nil
}
