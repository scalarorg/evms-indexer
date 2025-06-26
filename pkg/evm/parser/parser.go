package parser

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"
	eth_types "github.com/ethereum/go-ethereum/core/types"
	"github.com/rs/zerolog/log"
	contracts "github.com/scalarorg/evms-indexer/pkg/evm/contracts/generated"
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

var (
	scalarGatewayAbi *abi.ABI
)

func init() {
	var err error
	scalarGatewayAbi, err = contracts.IScalarGatewayMetaData.GetAbi()
	if err != nil {
		log.Error().Msgf("failed to get scalar gateway abi: %v", err)
	}
}

func GetScalarGatewayAbi() *abi.ABI {
	return scalarGatewayAbi
}

func getEventIndexedArguments(event *abi.Event) abi.Arguments {
	var args abi.Arguments
	if event != nil {
		for _, arg := range event.Inputs {
			if arg.Indexed {
				//Cast to non-indexed
				args = append(args, abi.Argument{
					Name: arg.Name,
					Type: arg.Type,
				})
			}
		}
	}
	return args
}

// TODO: change argument eventName to event
func ParseEventData(receiptLog *eth_types.Log, event *abi.Event, eventData any) error {
	gatewayAbi := GetScalarGatewayAbi()
	// Unpack non-indexed arguments
	if err := gatewayAbi.UnpackIntoInterface(eventData, event.Name, receiptLog.Data); err != nil {
		return fmt.Errorf("failed to unpack event: %w", err)
	}
	// Unpack indexed arguments
	// concat all topic data from second element into single buffer
	var buffer []byte
	for i := 1; i < len(receiptLog.Topics); i++ {
		buffer = append(buffer, receiptLog.Topics[i].Bytes()...)
	}
	indexedArgs := getEventIndexedArguments(event)
	if len(buffer) > 0 && len(indexedArgs) > 0 {
		unpacked, err := indexedArgs.Unpack(buffer)
		if err == nil {
			indexedArgs.Copy(eventData, unpacked)
		}
	}
	return nil
}

func CreateEvmEventFromArgs[T ValidEvmEvent](eventArgs T, log *eth_types.Log) *EvmEvent[T] {
	// Get the value of eventArgs using reflection
	// v := reflect.ValueOf(eventArgs).Elem()
	// sourceChain := currentChainName
	// if f := v.FieldByName("SourceChain"); f.IsValid() {
	// 	sourceChain = f.String()
	// }
	// destinationChain := currentChainName
	// if f := v.FieldByName("DestinationChain"); f.IsValid() {
	// 	destinationChain = f.String()
	// }

	return &EvmEvent[T]{
		Hash:        log.TxHash.Hex(),
		BlockNumber: log.BlockNumber,
		TxIndex:     log.TxIndex,
		LogIndex:    log.Index,
		Args:        eventArgs,
	}
}
