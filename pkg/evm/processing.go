package evm

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rs/zerolog/log"
	chains "github.com/scalarorg/data-models/chains"
	evmAbi "github.com/scalarorg/evms-indexer/pkg/evm/abi"
)

const (
	BATCH_SIZE = 1000
	// PostgreSQL has a limit of 65535 parameters per query
	// We'll use a conservative limit to account for field overhead
	MAX_PARAMS_PER_QUERY = 65535
)

// func (ec *EvmClient) ProcessLogFromSubscription(ctx context.Context, mapEvents map[string]*abi.Event,
// 	logChan <-chan types.Log, blockHeightsChan chan<- map[uint64]uint8) error {
// 	mapHeights := make(map[uint64]uint8)
// 	mapRecords := make(map[string][]any)
// 	for {
// 		select {
// 		case <-ctx.Done():
// 			return ctx.Err()
// 		case txLog := <-logChan:
// 			mapHeights[txLog.BlockNumber] = 1
// 			topic := txLog.Topics[0].String()
// 			event, ok := mapEvents[topic]
// 			if !ok {
// 				log.Error().Str("topic", topic).Any("txLog", txLog).Msg("[EvmClient] [ProcessMissingLogs] event not found")
// 				continue
// 			}
// 			_, model, err := parseEventLog(ec.EvmConfig.GetId(), event, txLog)
// 			if err != nil {
// 				log.Error().Err(err).Msg("[EvmClient] [ProcessMissingLogs] failed to parse event log")
// 				continue
// 			}
// 			if model != nil {
// 				mapRecords[topic] = append(mapRecords[topic], model)
// 			}
// 		}
// 	}
// }

func (ec *EvmClient) ProcessLogsFromFetcher(ctx context.Context, mapEvents map[string]*abi.Event,
	logsChan <-chan []types.Log,
	blockHeightsChan chan<- map[uint64]uint8) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case logs, ok := <-logsChan:
			if !ok {
				time.Sleep(100 * time.Millisecond)
				continue
			}

			mapRecords := make(map[string][]any)
			mapBlockHeights := make(map[uint64]uint8)
			for _, txLog := range logs {
				mapBlockHeights[txLog.BlockNumber] = 1
				topic := txLog.Topics[0].String()
				event, ok := mapEvents[topic]
				if !ok {
					log.Error().Str("topic", topic).Any("txLog", txLog).Msg("[EvmClient] [ProcessMissingLogs] event not found")
					continue
				}
				_, model, err := parseEventLog(ec.EvmConfig.GetId(), event, txLog)
				if err != nil {
					log.Error().Err(err).Msg("[EvmClient] [ProcessMissingLogs] failed to parse event log")
					continue
				}
				if model != nil {
					mapRecords[topic] = append(mapRecords[topic], model)
				}
			}
			blockHeightsChan <- mapBlockHeights
			for topic, records := range mapRecords {
				event := mapEvents[topic]
				if event == nil {
					log.Error().Str("topic", topic).Msg("[EvmClient] [ProcessRecords] event not found")
					continue
				}
				err := ec.SaveRecords(ctx, event, records)
				if err != nil {
					log.Error().Err(err).Msg("[EvmClient] [ProcessRecords] failed to save records")
					continue
				}
			}

		}
	}
}

// CalculateOptimalBatchSize calculates the optimal batch size based on the record type
// to avoid PostgreSQL's extended protocol parameter limit (65535)
func CalculateOptimalBatchSize(recordType reflect.Type) int {
	// Handle pointer types
	for recordType.Kind() == reflect.Ptr {
		recordType = recordType.Elem()
	}

	// Ensure we have a struct type
	if recordType.Kind() != reflect.Struct {
		// If it's not a struct, use a conservative default
		return BATCH_SIZE
	}

	// Count the number of fields that would be inserted
	fieldCount := recordType.NumField()

	// Calculate how many records we can safely insert in one batch
	// Each record will have fieldCount parameters, plus some overhead for the query itself
	maxRecordsPerBatch := MAX_PARAMS_PER_QUERY / fieldCount

	// Ensure we don't exceed our BATCH_SIZE constant
	if maxRecordsPerBatch > BATCH_SIZE {
		maxRecordsPerBatch = BATCH_SIZE
	}

	// Ensure we have at least 1 record per batch
	if maxRecordsPerBatch < 1 {
		maxRecordsPerBatch = 1
	}

	return maxRecordsPerBatch
}

// ChunkRecords divides records into chunks based on the optimal batch size
func ChunkRecords(records []any, batchSize int) [][]any {
	if len(records) == 0 {
		return nil
	}

	var chunks [][]any
	for i := 0; i < len(records); i += batchSize {
		end := i + batchSize
		if end > len(records) {
			end = len(records)
		}
		chunks = append(chunks, records[i:end])
	}
	return chunks
}

func (ec *EvmClient) SaveRecords(ctx context.Context, event *abi.Event, records []any) error {
	if len(records) == 0 {
		return nil
	}

	// Use reflection to determine the optimal batch size
	recordType := reflect.TypeOf(records[0])

	// Log the type information for debugging
	// log.Debug().
	// 	Str("eventName", event.Name).
	// 	Str("recordType", recordType.String()).
	// 	Str("recordKind", recordType.Kind().String()).
	// 	Msg("[EvmClient] [SaveRecords] Analyzing record type")

	maxRecordsPerBatch := CalculateOptimalBatchSize(recordType)

	// Get field count safely for logging
	fieldCount := 0
	if recordType.Kind() == reflect.Ptr {
		elemType := recordType.Elem()
		if elemType.Kind() == reflect.Struct {
			fieldCount = elemType.NumField()
		}
	} else if recordType.Kind() == reflect.Struct {
		fieldCount = recordType.NumField()
	}

	log.Debug().
		Str("eventName", event.Name).
		Int("totalRecords", len(records)).
		Int("fieldCount", fieldCount).
		Int("maxRecordsPerBatch", maxRecordsPerBatch).
		Msg("[EvmClient] [SaveRecords] Processing records in batches")

	// Process records in chunks
	chunks := ChunkRecords(records, maxRecordsPerBatch)
	for i, chunk := range chunks {
		log.Debug().
			Str("eventName", event.Name).
			Int("batchIndex", i).
			Int("batchSize", len(chunk)).
			Msg("[EvmClient] [SaveRecords] Processing batch")

		err := ec.SaveBatchRecords(ctx, event, chunk)
		if err != nil {
			log.Error().
				Err(err).
				Str("eventName", event.Name).
				Int("batchIndex", i).
				Msg("[EvmClient] [SaveRecords] Failed to save batch")
			return fmt.Errorf("failed to save batch %d: %w", i, err)
		}
	}

	return nil
}

func (ec *EvmClient) SaveBatchRecords(ctx context.Context, event *abi.Event, records []any) error {
	if len(records) == 0 {
		return nil
	}

	// Validate that all records are of the same type
	firstRecordType := reflect.TypeOf(records[0])
	for i, record := range records {
		recordType := reflect.TypeOf(record)
		if recordType != firstRecordType {
			return fmt.Errorf("inconsistent record types: expected %v, got %v at index %d", firstRecordType, recordType, i)
		}
	}

	switch event.Name {
	case evmAbi.EVENT_EVM_TOKEN_SENT:
		tokenSents := []*chains.TokenSent{}
		for _, record := range records {
			if tokenSent, ok := record.(*chains.TokenSent); ok {
				tokenSents = append(tokenSents, tokenSent)
			} else {
				return fmt.Errorf("invalid record type for TokenSent: expected *chains.TokenSent, got %T", record)
			}
		}
		return ec.dbAdapter.SaveTokenSents(tokenSents)
	case evmAbi.EVENT_EVM_CONTRACT_CALL:
		contractCalls := []*chains.ContractCall{}
		for _, record := range records {
			if contractCall, ok := record.(*chains.ContractCall); ok {
				contractCalls = append(contractCalls, contractCall)
			} else {
				return fmt.Errorf("invalid record type for ContractCall: expected *chains.ContractCall, got %T", record)
			}
		}
		return ec.dbAdapter.SaveContractCalls(contractCalls)
	case evmAbi.EVENT_EVM_CONTRACT_CALL_WITH_TOKEN:
		contractCallWithTokens := []*chains.ContractCallWithToken{}
		for _, record := range records {
			if contractCallWithToken, ok := record.(*chains.ContractCallWithToken); ok {
				contractCallWithTokens = append(contractCallWithTokens, contractCallWithToken)
			} else {
				return fmt.Errorf("invalid record type for ContractCallWithToken: expected *chains.ContractCallWithToken, got %T", record)
			}
		}
		return ec.dbAdapter.SaveContractCallWithTokens(contractCallWithTokens)
	case evmAbi.EVENT_EVM_CONTRACT_CALL_APPROVED:
		contractCallApproveds := []*chains.ContractCallApproved{}
		for _, record := range records {
			if contractCallApproved, ok := record.(*chains.ContractCallApproved); ok {
				contractCallApproveds = append(contractCallApproveds, contractCallApproved)
			} else {
				return fmt.Errorf("invalid record type for ContractCallApproved: expected *chains.ContractCallApproved, got %T", record)
			}
		}
		return ec.dbAdapter.SaveContractCallApproveds(contractCallApproveds)
	case evmAbi.EVENT_EVM_REDEEM_TOKEN:
		redeemTxs := []*chains.EvmRedeemTx{}
		for _, record := range records {
			if redeemTx, ok := record.(*chains.EvmRedeemTx); ok {
				redeemTxs = append(redeemTxs, redeemTx)
			} else {
				return fmt.Errorf("invalid record type for EvmRedeemTx: expected *chains.EvmRedeemTx, got %T", record)
			}
		}
		return ec.dbAdapter.SaveEvmRedeemTxs(redeemTxs)
	case evmAbi.EVENT_EVM_COMMAND_EXECUTED:
		executedTxs := []*chains.CommandExecuted{}
		for _, record := range records {
			if executedTx, ok := record.(*chains.CommandExecuted); ok {
				executedTxs = append(executedTxs, executedTx)
			} else {
				return fmt.Errorf("invalid record type for CommandExecuted: expected *chains.CommandExecuted, got %T", record)
			}
		}
		return ec.dbAdapter.SaveCommandExecuteds(executedTxs)
	case evmAbi.EVENT_EVM_TOKEN_DEPLOYED:
		tokenDeployed := []*chains.TokenDeployed{}
		for _, record := range records {
			if tokenDeployedRecord, ok := record.(*chains.TokenDeployed); ok {
				tokenDeployed = append(tokenDeployed, tokenDeployedRecord)
			} else {
				return fmt.Errorf("invalid record type for TokenDeployed: expected *chains.TokenDeployed, got %T", record)
			}
		}
		return ec.dbAdapter.SaveTokenDeployeds(tokenDeployed)
	case evmAbi.EVENT_EVM_SWITCHED_PHASE:
		switchedPhase := []*chains.SwitchedPhase{}
		for _, record := range records {
			if switchedPhaseRecord, ok := record.(*chains.SwitchedPhase); ok {
				switchedPhase = append(switchedPhase, switchedPhaseRecord)
			} else {
				return fmt.Errorf("invalid record type for SwitchedPhase: expected *chains.SwitchedPhase, got %T", record)
			}
		}
		return ec.dbAdapter.SaveSwitchedPhases(switchedPhase)
	default:
		return fmt.Errorf("unknown event: %s", event.Name)
	}
}
