package parser

import (
	"github.com/ethereum/go-ethereum/core/types"
	contracts "github.com/scalarorg/evms-indexer/pkg/evm/contracts/generated"
)

type EvmEvent[T any] struct {
	Hash             string //TxHash
	BlockNumber      uint64
	TxIndex          uint
	LogIndex         uint
	WaitForFinality  func() (*types.Receipt, error)
	SourceChain      string
	DestinationChain string
	EventName        string
	Args             T
}

type AllEvmEvents struct {
	ContractCallApproved *EvmEvent[*contracts.IScalarGatewayContractCallApproved]
	ContractCall         *EvmEvent[*contracts.IScalarGatewayContractCall]
	Executed             *EvmEvent[*contracts.IScalarGatewayExecuted]
	TokenSent            *EvmEvent[*contracts.IScalarGatewayTokenSent]
}
