package abi

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/rs/zerolog/log"
	contracts "github.com/scalarorg/evms-indexer/pkg/evm/contracts/generated"
)

var (
	uint8Type, _       = abi.NewType("uint8", "uint8", nil)
	uint64Type, _      = abi.NewType("uint64", "uint64", nil)
	boolType, _        = abi.NewType("bool", "bool", nil)
	bytesType, _       = abi.NewType("bytes", "bytes", nil)
	bytes32Type, _     = abi.NewType("bytes32", "bytes32", nil)
	stringArrayType, _ = abi.NewType("string[]", "string[]", nil)
	uint32ArrayType, _ = abi.NewType("uint32[]", "uint32[]", nil)
	uint64ArrayType, _ = abi.NewType("uint64[]", "uint64[]", nil)

	ContractCallWithTokenCustodianOnly = abi.Arguments{{Type: uint64Type}, {Type: bytesType}, {Type: stringArrayType}, {Type: uint32ArrayType}, {Type: uint64ArrayType}, {Type: bytes32Type}}
	ContractCallWithTokenUPC           = abi.Arguments{{Type: bytesType}}
	scalarGatewayAbi                   *abi.ABI
	mapEvents                          = map[string]*abi.Event{}
)

func init() {
	log.Info().Msg("Initializing ABI")
	scalarGatewayAbi, _ = contracts.IScalarGatewayMetaData.GetAbi()
	for _, event := range scalarGatewayAbi.Events {
		mapEvents[event.Name] = &event
	}
}

func GetScalarGatewayAbi() (*abi.ABI, error) {
	if scalarGatewayAbi == nil {
		var err error
		scalarGatewayAbi, err = contracts.IScalarGatewayMetaData.GetAbi()
		if err != nil {
			return nil, err
		}
	}
	return scalarGatewayAbi, nil
}

func GetMapEvents() map[string]*abi.Event {
	return mapEvents
}

func GetEventByName(name string) (*abi.Event, bool) {
	event, ok := mapEvents[name]
	return event, ok
}
