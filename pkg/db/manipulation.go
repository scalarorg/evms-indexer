package db

import (
	"github.com/rs/zerolog/log"
	"github.com/scalarorg/data-models/chains"
	"gorm.io/gorm/clause"
)

func (db *DatabaseAdapter) SaveSingleValue(value any) error {
	result := db.PostgresClient.Save(value)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// func (db *DatabaseAdapter) SaveMultipleValues(event string, values []any) error {
// 	switch event {
// 	case evmAbi.EVENT_EVM_TOKEN_SENT:
// 		result := db.PostgresClient.Model(&chains.TokenSent{}).Save(values)
// 		return result.Error
// 	case evmAbi.EVENT_EVM_CONTRACT_CALL:
// 		result := db.PostgresClient.Model(&chains.ContractCall{}).Save(values)
// 		return result.Error
// 	case evmAbi.EVENT_EVM_CONTRACT_CALL_WITH_TOKEN:
// 		result := db.PostgresClient.Model(&chains.ContractCallWithToken{}).Save(values)
// 		return result.Error
// 	case evmAbi.EVENT_EVM_CONTRACT_CALL_APPROVED:
// 		result := db.PostgresClient.Model(&chains.ContractCallApproved{}).Save(values)
// 		return result.Error
// 	case evmAbi.EVENT_EVM_REDEEM_TOKEN:
// 		result := db.PostgresClient.Model(&chains.EvmRedeemTx{}).Save(values)
// 		return result.Error
// 	case evmAbi.EVENT_EVM_COMMAND_EXECUTED:
// 		result := db.PostgresClient.Model(&chains.CommandExecuted{}).Save(values)
// 		return result.Error
// 	case evmAbi.EVENT_EVM_TOKEN_DEPLOYED:
// 		result := db.PostgresClient.Model(&chains.TokenDeployed{}).Save(values)
// 		return result.Error
// 	case evmAbi.EVENT_EVM_SWITCHED_PHASE:
// 		result := db.PostgresClient.Model(&chains.SwitchedPhase{}).Save(values)
// 		return result.Error
// 	default:
// 		return fmt.Errorf("unknown event: %s", event)
// 	}
// }

func (db *DatabaseAdapter) SaveContractCalls(values []*chains.ContractCall) error {
	result := db.PostgresClient.Clauses(
		clause.OnConflict{
			Columns:   []clause.Column{{Name: "event_id"}},
			DoNothing: true,
		},
	).Create(values)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (db *DatabaseAdapter) SaveContractCallWithTokens(values []*chains.ContractCallWithToken) error {
	result := db.PostgresClient.Clauses(
		clause.OnConflict{
			Columns:   []clause.Column{{Name: "event_id"}},
			DoNothing: true,
		},
	).Create(values)
	if result.Error != nil {
		log.Info().Any("values", values).Msg("SaveContractCallWithTokens")
		return result.Error
	}
	return nil
}

func (db *DatabaseAdapter) SaveContractCallApproveds(values []*chains.ContractCallApproved) error {
	result := db.PostgresClient.Clauses(
		clause.OnConflict{
			Columns:   []clause.Column{{Name: "event_id"}},
			DoNothing: true,
		},
	).Create(values)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (db *DatabaseAdapter) SaveEvmRedeemTxs(values []*chains.EvmRedeemTx) error {
	result := db.PostgresClient.Clauses(
		clause.OnConflict{
			Columns:   []clause.Column{{Name: "event_id"}},
			DoNothing: true,
		},
	).Create(values)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (db *DatabaseAdapter) SaveTokenDeployeds(values []*chains.TokenDeployed) error {
	result := db.PostgresClient.Clauses(
		clause.OnConflict{
			Columns:   []clause.Column{{Name: "source_chain"}, {Name: "token_address"}},
			DoNothing: true,
		},
	).Create(values)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (db *DatabaseAdapter) SaveSwitchedPhases(values []*chains.SwitchedPhase) error {
	result := db.PostgresClient.Clauses(
		clause.OnConflict{
			Columns:   []clause.Column{{Name: "source_chain"}, {Name: "tx_hash"}, {Name: "custodian_group_uid"}, {Name: "session_sequence"}},
			DoNothing: true,
		},
	).Create(values)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
