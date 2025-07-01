package evm

import (
	"fmt"
	"math/big"
	"reflect"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	eth_types "github.com/ethereum/go-ethereum/core/types"
	evmAbi "github.com/scalarorg/evms-indexer/pkg/evm/abi"
	contracts "github.com/scalarorg/evms-indexer/pkg/evm/contracts/generated"
	"github.com/scalarorg/evms-indexer/pkg/evm/parser"
)

type ValidEvmEvent interface {
	*contracts.IScalarGatewayContractCallApproved |
		*contracts.IScalarGatewayContractCall |
		*contracts.IScalarGatewayContractCallWithToken |
		*contracts.IScalarGatewayRedeemToken |
		*contracts.IScalarGatewayExecuted |
		*contracts.IScalarGatewayTokenSent |
		*contracts.IScalarGatewaySwitchPhase |
		*contracts.IScalarGatewayTokenDeployed
}

// Return contracts.{EVENT} and gorm.Model
func parseEventLog(sourceChain string, event *abi.Event, txLog types.Log) (any, any, error) {
	switch event.Name {
	case evmAbi.EVENT_EVM_TOKEN_SENT:
		tokenSent := &contracts.IScalarGatewayTokenSent{
			Raw: txLog,
		}
		err := parser.ParseEventData(&txLog, event, tokenSent)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to parse event %s: %w", event.Name, err)
		}
		model := parser.TokenSentEvent2Model(sourceChain, tokenSent)
		return tokenSent, model, nil
	case evmAbi.EVENT_EVM_CONTRACT_CALL_WITH_TOKEN:
		contractCallWithToken := &contracts.IScalarGatewayContractCallWithToken{
			Raw: txLog,
		}
		err := parser.ParseEventData(&txLog, event, contractCallWithToken)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to parse event %s: %w", event.Name, err)
		}
		model := parser.ContractCallWithTokenEvent2Model(sourceChain, contractCallWithToken)
		return contractCallWithToken, model, nil
	case evmAbi.EVENT_EVM_CONTRACT_CALL:
		contractCall := &contracts.IScalarGatewayContractCall{
			Raw: txLog,
		}
		err := parser.ParseEventData(&txLog, event, contractCall)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to parse event %s: %w", event.Name, err)
		}
		model := parser.ContractCallEvent2Model(sourceChain, contractCall)
		return contractCall, model, nil
	case evmAbi.EVENT_EVM_CONTRACT_CALL_APPROVED:
		contractCallApproved := &contracts.IScalarGatewayContractCallApproved{
			Raw: txLog,
		}
		err := parser.ParseEventData(&txLog, event, contractCallApproved)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to parse event %s: %w", event.Name, err)
		}
		model := parser.ContractCallApprovedEvent2Model(sourceChain, contractCallApproved)
		return contractCallApproved, model, nil
	case evmAbi.EVENT_EVM_REDEEM_TOKEN:
		redeemToken := &contracts.IScalarGatewayRedeemToken{
			Raw: txLog,
		}
		err := parser.ParseEventData(&txLog, event, redeemToken)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to parse event %s: %w", event.Name, err)
		}
		model := parser.RedeemTokenEvent2Model(sourceChain, redeemToken)
		return redeemToken, model, nil
	case evmAbi.EVENT_EVM_COMMAND_EXECUTED:
		executed := &contracts.IScalarGatewayExecuted{
			Raw: txLog,
		}
		err := parser.ParseEventData(&txLog, event, executed)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to parse event %s: %w", event.Name, err)
		}
		model := parser.CommandExecutedEvent2Model(sourceChain, executed)
		return executed, model, nil
	case evmAbi.EVENT_EVM_TOKEN_DEPLOYED:
		tokenDeployed := &contracts.IScalarGatewayTokenDeployed{
			Raw: txLog,
		}
		err := parser.ParseEventData(&txLog, event, tokenDeployed)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to parse event %s: %w", event.Name, err)
		}
		model := parser.TokenDeployedEvent2Model(sourceChain, tokenDeployed)
		return tokenDeployed, model, nil
	case evmAbi.EVENT_EVM_SWITCHED_PHASE:
		switchedPhase := &contracts.IScalarGatewaySwitchPhase{
			Raw: txLog,
		}
		err := parser.ParseEventData(&txLog, event, switchedPhase)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to parse event %s: %w", event.Name, err)
		}
		model := parser.SwitchPhaseEvent2Model(sourceChain, switchedPhase)
		return switchedPhase, model, nil
	default:
		return nil, nil, fmt.Errorf("invalid event type for %s: %T", event.Name, txLog)
	}
}

// Todo: Check if this is the correct way to extract the destination chain
// Maybe add destination chain to the event.Log
func extractDestChainFromEvmGwContractCallApproved(event *contracts.IScalarGatewayContractCallApproved) string {
	return event.SourceChain
}
func parseLogIntoEventArgs(log eth_types.Log) (any, error) {
	// Try parsing as ContractCallApproved
	if eventArgs, err := parseContractCallApproved(log); err == nil {
		return eventArgs, nil
	}

	// Try parsing as ContractCall
	if eventArgs, err := parseContractCall(log); err == nil {
		return eventArgs, nil
	}

	// Try parsing as Execute
	if eventArgs, err := parseExecute(log); err == nil {
		return eventArgs, nil
	}

	return nil, fmt.Errorf("failed to parse log into any known event type")
}

// func parseEventIntoEnvelope(currentChainName string, eventArgs any, log eth_types.Log) (types.EventEnvelope, error) {
// 	switch args := eventArgs.(type) {
// 	case *contracts.IScalarGatewayContractCallApproved:
// 		event, err := parseEventArgsIntoEvent[*contracts.IScalarGatewayContractCallApproved](args, currentChainName, log)
// 		if err != nil {
// 			return types.EventEnvelope{}, err
// 		}
// 		return types.EventEnvelope{
// 			Component:        "DbAdapter",
// 			SenderClientName: currentChainName,
// 			Handler:          "FindCosmosToEvmCallContractApproved",
// 			Data:             event,
// 		}, nil

// 	case *contracts.IScalarGatewayContractCall:
// 		event, err := parseEventArgsIntoEvent[*contracts.IScalarGatewayContractCall](args, currentChainName, log)
// 		if err != nil {
// 			return types.EventEnvelope{}, err
// 		}
// 		return types.EventEnvelope{
// 			Component:        "DbAdapter",
// 			SenderClientName: currentChainName,
// 			Handler:          "CreateEvmCallContractEvent",
// 			Data:             event,
// 		}, nil

// 	case *contracts.IScalarGatewayExecuted:
// 		event, err := parseEventArgsIntoEvent[*contracts.IScalarGatewayExecuted](args, currentChainName, log)
// 		if err != nil {
// 			return types.EventEnvelope{}, err
// 		}
// 		return types.EventEnvelope{
// 			Component:        "DbAdapter",
// 			SenderClientName: currentChainName,
// 			Handler:          "CreateEvmExecutedEvent",
// 			Data:             event,
// 		}, nil

// 	default:
// 		return types.EventEnvelope{}, fmt.Errorf("unknown event type: %T", eventArgs)
// 	}
// }

func parseEventArgsIntoEvent[T ValidEvmEvent](eventArgs T, currentChainName string, log eth_types.Log) (*parser.EvmEvent[T], error) {
	// Get the value of eventArgs using reflection
	v := reflect.ValueOf(eventArgs).Elem()
	sourceChain := currentChainName
	if f := v.FieldByName("SourceChain"); f.IsValid() {
		sourceChain = f.String()
	}
	destinationChain := currentChainName
	if f := v.FieldByName("DestinationChain"); f.IsValid() {
		destinationChain = f.String()
	}

	return &parser.EvmEvent[T]{
		Hash:             log.TxHash.Hex(),
		BlockNumber:      log.BlockNumber,
		LogIndex:         log.Index,
		SourceChain:      sourceChain,
		DestinationChain: destinationChain,
		Args:             eventArgs,
	}, nil
}

// parseAnyEvent is a generic function that parses any EVM event into a standardized EvmEvent structure
func parseEvmEventContractCallApproved[T *contracts.IScalarGatewayContractCallApproved](
	currentChainName string,
	log eth_types.Log,
) (*parser.EvmEvent[T], error) {
	eventArgs, err := parseContractCallApproved(log)
	if err != nil {
		return nil, err
	}

	event, err := parseEventArgsIntoEvent[T](eventArgs, currentChainName, log)
	if err != nil {
		return nil, err
	}

	return event, nil
}

func parseContractCallApproved(
	log eth_types.Log,
) (*contracts.IScalarGatewayContractCallApproved, error) {
	event := struct {
		CommandId        [32]byte
		SourceChain      string
		SourceAddress    string
		ContractAddress  common.Address
		PayloadHash      [32]byte
		SourceTxHash     [32]byte
		SourceEventIndex *big.Int
	}{}

	abi, err := evmAbi.GetScalarGatewayAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to parse ABI: %w", err)
	}
	if err := abi.UnpackIntoInterface(&event, "ContractCallApproved", log.Data); err != nil {
		return nil, fmt.Errorf("failed to unpack event: %w", err)
	}

	// Add validation for required fields
	if len(event.SourceChain) == 0 || !isValidUTF8(event.SourceChain) {
		return nil, fmt.Errorf("invalid source chain value")
	}

	if len(event.SourceAddress) == 0 || !isValidUTF8(event.SourceAddress) {
		return nil, fmt.Errorf("invalid source address value")
	}

	var eventArgs contracts.IScalarGatewayContractCallApproved = contracts.IScalarGatewayContractCallApproved{
		CommandId:        event.CommandId,
		SourceChain:      event.SourceChain,
		SourceAddress:    event.SourceAddress,
		ContractAddress:  event.ContractAddress,
		PayloadHash:      event.PayloadHash,
		SourceTxHash:     event.SourceTxHash,
		SourceEventIndex: event.SourceEventIndex,
		Raw:              log,
	}

	fmt.Printf("[EVMListener] [parseContractCallApproved] eventArgs: %+v\n", eventArgs)

	return &eventArgs, nil
}

func parseEvmEventContractCall[T *contracts.IScalarGatewayContractCall](
	currentChainName string,
	log eth_types.Log,
) (*parser.EvmEvent[T], error) {
	eventArgs, err := parseContractCall(log)
	if err != nil {
		return nil, err
	}

	event, err := parseEventArgsIntoEvent[T](eventArgs, currentChainName, log)
	if err != nil {
		return nil, err
	}

	return event, nil
}

func parseContractCall(
	log eth_types.Log,
) (*contracts.IScalarGatewayContractCall, error) {
	event := struct {
		Sender                     common.Address
		DestinationChain           string
		DestinationContractAddress string
		PayloadHash                [32]byte
		Payload                    []byte
	}{}

	abi, err := evmAbi.GetScalarGatewayAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to parse ABI: %w", err)
	}
	if err := abi.UnpackIntoInterface(&event, "ContractCall", log.Data); err != nil {
		return nil, fmt.Errorf("failed to unpack event: %w", err)
	}

	// Add validation for required fields
	if len(event.DestinationChain) == 0 || !isValidUTF8(event.DestinationChain) {
		return nil, fmt.Errorf("invalid destination chain value")
	}

	if len(event.DestinationContractAddress) == 0 || !isValidUTF8(event.DestinationContractAddress) {
		return nil, fmt.Errorf("invalid destination address value")
	}

	var eventArgs contracts.IScalarGatewayContractCall = contracts.IScalarGatewayContractCall{
		Sender:                     event.Sender,
		DestinationChain:           event.DestinationChain,
		DestinationContractAddress: event.DestinationContractAddress,
		PayloadHash:                event.PayloadHash,
		Payload:                    event.Payload,
		Raw:                        log,
	}

	fmt.Printf("[EVMListener] [parseContractCall] eventArgs: %+v\n", eventArgs)

	return &eventArgs, nil
}

func parseEvmEventExecute[T *contracts.IScalarGatewayExecuted](
	currentChainName string,
	log eth_types.Log,
) (*parser.EvmEvent[T], error) {
	eventArgs, err := parseExecute(log)
	if err != nil {
		return nil, err
	}

	event, err := parseEventArgsIntoEvent[T](eventArgs, currentChainName, log)
	if err != nil {
		return nil, err
	}

	return event, nil
}

func parseExecute(
	log eth_types.Log,
) (*contracts.IScalarGatewayExecuted, error) {
	event := struct {
		CommandId [32]byte
	}{}
	abi, err := evmAbi.GetScalarGatewayAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to parse ABI: %w", err)
	}

	// Check if log data size matches exactly what we expect (32 bytes for CommandId)
	if len(log.Data) != 32 {
		return nil, fmt.Errorf("unexpected log data size: got %d bytes, want 32 bytes", len(log.Data))
	}

	if err := abi.UnpackIntoInterface(&event, "Executed", log.Data); err != nil {
		return nil, fmt.Errorf("failed to unpack event: %w", err)
	}

	// Add validation for required fields
	if len(event.CommandId) == 0 {
		return nil, fmt.Errorf("invalid command id value")
	}

	var eventArgs contracts.IScalarGatewayExecuted = contracts.IScalarGatewayExecuted{
		CommandId: event.CommandId,
		Raw:       log,
	}

	fmt.Printf("[EVMListener] [parseExecute] eventArgs: %+v\n", eventArgs)

	return &eventArgs, nil
}

// Add helper function to validate UTF-8 strings
func isValidUTF8(s string) bool {
	return strings.ToValidUTF8(s, "") == s
}
