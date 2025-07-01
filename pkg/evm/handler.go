package evm

import (
	"context"
	"encoding/hex"
	"fmt"
	"reflect"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rs/zerolog/log"
	"github.com/scalarorg/data-models/chains"
	evmAbi "github.com/scalarorg/evms-indexer/pkg/evm/abi"
	contracts "github.com/scalarorg/evms-indexer/pkg/evm/contracts/generated"
	"github.com/scalarorg/evms-indexer/pkg/evm/parser"
)

const (
	BATCH_SIZE = 1000
	// PostgreSQL has a limit of 65535 parameters per query
	// We'll use a conservative limit to account for field overhead
	MAX_PARAMS_PER_QUERY = 65535
)

func (ec *EvmClient) ProcessLogs(ctx context.Context, mapEvents map[string]*abi.Event,
	logsChan <-chan []types.Log) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case logs, ok := <-logsChan:
			if !ok {
				return ctx.Err()
			}

			mapRecords := make(map[string][]any)
			for _, txLog := range logs {
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
				mapRecords[topic] = append(mapRecords[topic], model)
			}
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
	log.Debug().
		Str("eventName", event.Name).
		Str("recordType", recordType.String()).
		Str("recordKind", recordType.Kind().String()).
		Msg("[EvmClient] [SaveRecords] Analyzing record type")

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
		tokenSents := make([]*chains.TokenSent, 0, len(records))
		for _, record := range records {
			if tokenSent, ok := record.(*chains.TokenSent); ok {
				tokenSents = append(tokenSents, tokenSent)
			} else {
				return fmt.Errorf("invalid record type for TokenSent: expected *chains.TokenSent, got %T", record)
			}
		}
		return ec.dbAdapter.SaveTokenSents(tokenSents)
	case evmAbi.EVENT_EVM_CONTRACT_CALL:
		contractCalls := make([]*chains.ContractCall, 0, len(records))
		for _, record := range records {
			if contractCall, ok := record.(*chains.ContractCall); ok {
				contractCalls = append(contractCalls, contractCall)
			} else {
				return fmt.Errorf("invalid record type for ContractCall: expected *chains.ContractCall, got %T", record)
			}
		}
		return ec.dbAdapter.SaveContractCalls(contractCalls)
	case evmAbi.EVENT_EVM_CONTRACT_CALL_WITH_TOKEN:
		contractCallWithTokens := make([]*chains.ContractCallWithToken, 0, len(records))
		for _, record := range records {
			if contractCallWithToken, ok := record.(*chains.ContractCallWithToken); ok {
				contractCallWithTokens = append(contractCallWithTokens, contractCallWithToken)
			} else {
				return fmt.Errorf("invalid record type for ContractCallWithToken: expected *chains.ContractCallWithToken, got %T", record)
			}
		}
		return ec.dbAdapter.SaveContractCallWithTokens(contractCallWithTokens)
	case evmAbi.EVENT_EVM_CONTRACT_CALL_APPROVED:
		contractCallApproveds := make([]*chains.ContractCallApproved, 0, len(records))
		for _, record := range records {
			if contractCallApproved, ok := record.(*chains.ContractCallApproved); ok {
				contractCallApproveds = append(contractCallApproveds, contractCallApproved)
			} else {
				return fmt.Errorf("invalid record type for ContractCallApproved: expected *chains.ContractCallApproved, got %T", record)
			}
		}
		return ec.dbAdapter.SaveContractCallApproveds(contractCallApproveds)
	case evmAbi.EVENT_EVM_REDEEM_TOKEN:
		redeemTxs := make([]*chains.EvmRedeemTx, 0, len(records))
		for _, record := range records {
			if redeemTx, ok := record.(*chains.EvmRedeemTx); ok {
				redeemTxs = append(redeemTxs, redeemTx)
			} else {
				return fmt.Errorf("invalid record type for EvmRedeemTx: expected *chains.EvmRedeemTx, got %T", record)
			}
		}
		return ec.dbAdapter.SaveEvmRedeemTxs(redeemTxs)
	case evmAbi.EVENT_EVM_COMMAND_EXECUTED:
		executedTxs := make([]*chains.CommandExecuted, 0, len(records))
		for _, record := range records {
			if executedTx, ok := record.(*chains.CommandExecuted); ok {
				executedTxs = append(executedTxs, executedTx)
			} else {
				return fmt.Errorf("invalid record type for CommandExecuted: expected *chains.CommandExecuted, got %T", record)
			}
		}
		return ec.dbAdapter.SaveCommandExecuteds(executedTxs)
	case evmAbi.EVENT_EVM_TOKEN_DEPLOYED:
		tokenDeployed := make([]*chains.TokenDeployed, 0, len(records))
		for _, record := range records {
			if tokenDeployedRecord, ok := record.(*chains.TokenDeployed); ok {
				tokenDeployed = append(tokenDeployed, tokenDeployedRecord)
			} else {
				return fmt.Errorf("invalid record type for TokenDeployed: expected *chains.TokenDeployed, got %T", record)
			}
		}
		return ec.dbAdapter.SaveTokenDeployeds(tokenDeployed)
	case evmAbi.EVENT_EVM_SWITCHED_PHASE:
		switchedPhase := make([]*chains.SwitchedPhase, 0, len(records))
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

// func (ec *EvmClient) handleContractCall(event *contracts.IScalarGatewayContractCall) error {
// 	//0. Preprocess the event
// 	ec.preprocessContractCall(event)
// 	//1. Convert into a RelayData instance then store to the db
// 	contractCall, err := ec.ContractCallEvent2Model(event)
// 	if err != nil {
// 		return fmt.Errorf("failed to convert ContractCallEvent to RelayData: %w", err)
// 	}
// 	//2. update last checkpoint
// 	lastCheckpoint, err := ec.dbAdapter.GetLastEventCheckPoint(ec.EvmConfig.GetId(), EVENT_EVM_CONTRACT_CALL)
// 	if err != nil {
// 		log.Debug().Str("chainId", ec.EvmConfig.GetId()).
// 			Str("eventName", EVENT_EVM_CONTRACT_CALL).
// 			Msg("[EvmClient] [handleContractCall] Get event from begining")
// 	}
// 	if event.Raw.BlockNumber > lastCheckpoint.BlockNumber ||
// 		(event.Raw.BlockNumber == lastCheckpoint.BlockNumber && event.Raw.TxIndex > lastCheckpoint.LogIndex) {
// 		lastCheckpoint.BlockNumber = event.Raw.BlockNumber
// 		lastCheckpoint.TxHash = event.Raw.TxHash.String()
// 		lastCheckpoint.LogIndex = event.Raw.Index
// 		lastCheckpoint.EventKey = fmt.Sprintf("%s-%d-%d", event.Raw.TxHash.String(), event.Raw.BlockNumber, event.Raw.Index)
// 	}
// 	//3. store relay data to the db, update last checkpoint
// 	err = ec.dbAdapter.CreateContractCall(contractCall, lastCheckpoint)
// 	if err != nil {
// 		return fmt.Errorf("failed to create evm contract call: %w", err)
// 	}
// 	//2. Send to the bus
// 	confirmTxs := events.ConfirmTxsRequest{
// 		ChainName: ec.EvmConfig.GetId(),
// 		TxHashs:   map[string]string{contractCall.TxHash: contractCall.DestinationChain},
// 	}
// 	if ec.eventBus != nil {
// 		ec.eventBus.BroadcastEvent(&events.EventEnvelope{
// 			EventType:        EVENT_EVM_CONTRACT_CALL,
// 			DestinationChain: events.SCALAR_NETWORK_NAME,
// 			Data:             confirmTxs,
// 		})
// 	} else {
// 		log.Warn().Msg("[EvmClient] [handleContractCall] event bus is undefined")
// 	}
// 	return nil
// }
// func (ec *EvmClient) preprocessContractCall(event *contracts.IScalarGatewayContractCall) error {
// 	log.Info().
// 		Str("sender", event.Sender.Hex()).
// 		Str("destinationChain", event.DestinationChain).
// 		Str("destinationContractAddress", event.DestinationContractAddress).
// 		Str("payloadHash", hex.EncodeToString(event.PayloadHash[:])).
// 		Str("txHash", event.Raw.TxHash.String()).
// 		Uint("logIndex", event.Raw.Index).
// 		Uint("txIndex", event.Raw.TxIndex).
// 		Str("logData", hex.EncodeToString(event.Raw.Data)).
// 		Msg("[EvmClient] [preprocessContractCall] Start handle Contract call")
// 	//Todo: validate the event
// 	return nil
// }

func (ec *EvmClient) HandleContractCallWithToken(event *contracts.IScalarGatewayContractCallWithToken) error {
	// Get block header
	err := ec.fetchBlockHeader(event.Raw.BlockNumber)
	if err != nil {
		log.Error().Err(err).Msgf("[EvmClient] [HandleContractCallWithToken] failed to fetch block header %d", event.Raw.BlockNumber)
	}

	// Convert into a RelayData instance then store to the db
	contractCallWithToken, err := ec.ContractCallWithToken2Model(event)
	if err != nil {
		return fmt.Errorf("failed to convert ContractCallEvent to ContractCallWithToken: %w", err)
	}

	// Store relay data to the db
	err = ec.dbAdapter.CreateContractCallWithToken(contractCallWithToken)
	if err != nil {
		return fmt.Errorf("failed to create evm contract call: %w", err)
	}
	return nil
}

func (ec *EvmClient) HandleRedeemToken(event *contracts.IScalarGatewayRedeemToken) error {
	log.Info().Str("Chain", ec.EvmConfig.ID).Any("event", event).Msg("[EvmClient] [HandleRedeemToken] Start processing evm redeem token")
	err := ec.fetchBlockHeader(event.Raw.BlockNumber)
	if err != nil {
		log.Error().Err(err).Msgf("[EvmClient] [HandleRedeemToken] failed to fetch block header %d", event.Raw.BlockNumber)
	}

	// Convert into a RelayData instance then store to the db
	redeemToken, err := ec.RedeemTokenEvent2Model(event)
	if err != nil {
		return fmt.Errorf("failed to convert ContractCallEvent to ContractCallWithToken: %w", err)
	}

	// Store relay data to the db
	err = ec.dbAdapter.CreateContractCallWithToken(redeemToken)
	if err != nil {
		return fmt.Errorf("failed to create evm contract call: %w", err)
	}
	return nil
}

func (ec *EvmClient) HandleTokenSent(event *contracts.IScalarGatewayTokenSent) error {
	err := ec.fetchBlockHeader(event.Raw.BlockNumber)
	if err != nil {
		log.Error().Err(err).Msgf("[EvmClient] [HandleTokenSent] failed to fetch block header %d", event.Raw.BlockNumber)
	}

	// Convert into a RelayData instance then store to the db
	tokenSent := parser.TokenSentEvent2Model(ec.EvmConfig.GetId(), event)

	// For evm, the token sent is verified immediately by the scalarnet
	tokenSent.Status = chains.TokenSentStatusVerifying

	// Store relay data to the db
	err = ec.dbAdapter.SaveTokenSent(tokenSent)
	if err != nil {
		return fmt.Errorf("failed to create evm token send: %w", err)
	}
	return nil
}

func (ec *EvmClient) preprocessContractCallWithToken(event *contracts.IScalarGatewayContractCallWithToken) error {
	log.Info().
		Str("sender", event.Sender.Hex()).
		Str("destinationChain", event.DestinationChain).
		Str("destinationContractAddress", event.DestinationContractAddress).
		Str("payloadHash", hex.EncodeToString(event.PayloadHash[:])).
		Str("Symbol", event.Symbol).
		Uint64("Amount", event.Amount.Uint64()).
		Str("txHash", event.Raw.TxHash.String()).
		Uint("logIndex", event.Raw.Index).
		Uint("txIndex", event.Raw.TxIndex).
		Str("logData", hex.EncodeToString(event.Raw.Data)).
		Msg("[EvmClient] [preprocessContractCallWithToken] Start handle Contract call with token")
	//Todo: validate the event
	return nil
}

func (ec *EvmClient) HandleContractCallApproved(event *contracts.IScalarGatewayContractCallApproved) error {
	//0. Preprocess the event
	err := ec.preprocessContractCallApproved(event)
	if err != nil {
		return fmt.Errorf("failed to preprocess contract call approved: %w", err)
	}
	err = ec.fetchBlockHeader(event.Raw.BlockNumber)
	if err != nil {
		log.Error().Err(err).Msgf("[EvmClient] [HandleRedeemToken] failed to fetch block header %d", event.Raw.BlockNumber)
	}
	//1. Convert into a RelayData instance then store to the db
	contractCallApproved := parser.ContractCallApprovedEvent2Model(ec.EvmConfig.GetId(), event)
	err = ec.dbAdapter.SaveSingleValue(&contractCallApproved)
	if err != nil {
		return fmt.Errorf("failed to create contract call approved: %w", err)
	}
	// Find relayData from the db by combination (contractAddress, sourceAddress, payloadHash)
	// This contract call (initiated by the user call to the source chain) is approved by EVM network
	// So anyone can execute it on the EVM by broadcast the corresponding payload to protocol's smart contract on the destination chain
	destContractAddress := strings.TrimLeft(event.ContractAddress.Hex(), "0x")
	sourceAddress := strings.TrimLeft(event.SourceAddress, "0x")
	payloadHash := strings.TrimLeft(hex.EncodeToString(event.PayloadHash[:]), "0x")
	relayDatas, err := ec.dbAdapter.FindContractCallByParams(sourceAddress, destContractAddress, payloadHash)
	if err != nil {
		log.Error().Err(err).Msg("[EvmClient] [handleContractCallApproved] find relay data")
		return err
	}
	log.Debug().Str("contractAddress", event.ContractAddress.String()).
		Str("sourceAddress", event.SourceAddress).
		Str("payloadHash", hex.EncodeToString(event.PayloadHash[:])).
		Any("relayDatas count", len(relayDatas)).
		Msg("[EvmClient] [handleContractCallApproved] query relaydata by ContractCall")
	return nil
}

func (ec *EvmClient) preprocessContractCallApproved(event *contracts.IScalarGatewayContractCallApproved) error {
	log.Info().Any("event", event).Msgf("[EvmClient] [handleContractCallApproved]")
	//Todo: validate the event
	return nil
}

func (ec *EvmClient) HandleCommandExecuted(event *contracts.IScalarGatewayExecuted) error {
	//0. Preprocess the event
	//ec.preprocessCommandExecuted(event)
	err := ec.fetchBlockHeader(event.Raw.BlockNumber)
	if err != nil {
		log.Error().Err(err).Msgf("[EvmClient] [HandleCommandExecuted] failed to fetch block header %d", event.Raw.BlockNumber)
	}
	//1. Convert into a RelayData instance then store to the db
	cmdExecuted := parser.CommandExecutedEvent2Model(ec.EvmConfig.GetId(), event)
	if ec.dbAdapter != nil {
		err = ec.dbAdapter.SaveCommandExecuted(cmdExecuted)
		if err != nil {
			log.Error().Err(err).Msg("[EvmClient] [HandleCommandExecuted] failed to save evm executed to the db")
			return fmt.Errorf("failed to create evm executed: %w", err)
		}
	}
	return nil
}

func (ec *EvmClient) preprocessCommandExecuted(event *contracts.IScalarGatewayExecuted) error {
	log.Info().Any("event", event).Msg("[EvmClient] [ExecutedHandler] Start processing evm command executed")
	//Todo: validate the event
	return nil
}

func (ec *EvmClient) HandleTokenDeployed(event *contracts.IScalarGatewayTokenDeployed) error {
	//0. Preprocess the event
	log.Info().Any("event", event).Msg("[EvmClient] [HandleTokenDeployed] Start processing evm token deployed")
	//1. Convert into a RelayData instance then store to the db
	tokenDeployed := parser.TokenDeployedEvent2Model(ec.EvmConfig.GetId(), event)
	if ec.dbAdapter != nil {
		err := ec.dbAdapter.SaveTokenDeployed(tokenDeployed)
		if err != nil {
			return fmt.Errorf("failed to create evm token deployed: %w", err)
		}
	}
	return nil
}

func (ec *EvmClient) HandleSwitchPhase(event *contracts.IScalarGatewaySwitchPhase) error {
	//0. Preprocess the event
	log.Info().Str("Chain", ec.EvmConfig.ID).Any("event", event).Msg("[EvmClient] [HandleSwitchPhase] Start processing evm switch phase")
	//1. Convert into a RelayData instance then store to the db
	switchPhase := ec.SwitchPhaseEvent2Model(event)
	if ec.dbAdapter != nil {
		err := ec.dbAdapter.SaveSingleValue(&switchPhase)
		if err != nil {
			return fmt.Errorf("failed to create evm switch phase: %w", err)
		}
	}
	return nil
}
