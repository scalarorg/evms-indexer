package types

import (
	"math"

	"github.com/ethereum/go-ethereum/common"
	contracts "github.com/scalarorg/evms-indexer/pkg/evm/contracts/generated"
)

type ExecuteParams struct {
	SourceChain      string
	SourceAddress    string
	ContractAddress  common.Address
	PayloadHash      [32]byte
	SourceTxHash     [32]byte
	SourceEventIndex uint64
}

// AxelarListenerEvent represents an event with generic type T for parsed events
type ScalarListenerEvent[T any] struct {
	TopicID    string
	Type       string
	ParseEvent func(events map[string][]string) (T, error)
}

// IBCEvent represents a generic IBC event with generic type T for Args
type IBCEvent[T any] struct {
	Hash        string `json:"hash"`
	SrcChannel  string `json:"srcChannel,omitempty"`
	DestChannel string `json:"destChannel,omitempty"`
	Args        T      `json:"args"`
}

// IBCPacketEvent represents an IBC packet event
type IBCPacketEvent struct {
	Hash        string      `json:"hash"`
	SrcChannel  string      `json:"srcChannel"`
	DestChannel string      `json:"destChannel"`
	Denom       string      `json:"denom"`
	Amount      string      `json:"amount"`
	Sequence    int         `json:"sequence"`
	Memo        interface{} `json:"memo"`
}

// ------ Payloads ------
// TODO: USING COSMOS SDK TO DEFINE THESE TYPES LATER
const (
	ChainsProtobufPackage = "scalar.chains.v1beta1"
	AxelarProtobufPackage = "scalar.scalarnet.v1beta1"
)

// ConfirmGatewayTxRequest represents a request to confirm a gateway transaction
type ConfirmGatewayTxRequest struct {
	Sender []byte `json:"sender"`
	Chain  string `json:"chain"`
	TxID   []byte `json:"txId"`
}

// RouteMessageRequest represents a request to route a message
type RouteMessageRequest struct {
	Sender  []byte `json:"sender"`
	ID      string `json:"id"`
	Payload []byte `json:"payload"`
}

// SignCommandsRequest represents a request to sign commands
type SignCommandsRequest struct {
	Sender []byte `json:"sender"`
	Chain  string `json:"chain"`
}

//	type RedeemTx struct {
//		BlockHeight uint64
//		TxHash      string
//		LogIndex    uint
//	}
type Session struct {
	Sequence uint64
	Phase    uint8
}

func (s *Session) Cmp(other *Session) int64 {
	if other == nil {
		return math.MaxInt64
	}
	var diffSeq, diffPhase int64
	if s.Sequence >= other.Sequence {
		diffSeq = int64(s.Sequence - other.Sequence)
	} else {
		diffSeq = -int64(other.Sequence - s.Sequence)
	}

	if s.Phase >= other.Phase {
		diffPhase = int64(s.Phase - other.Phase)
	} else {
		diffPhase = -int64(other.Phase - s.Phase)
	}

	return diffSeq*2 + diffPhase
}

// Eache chainId container an array of SwitchPhaseEvents,
// with the first element is switch to Preparing phase
type GroupRedeemSessions struct {
	GroupUid          string
	MaxSession        Session
	MinSession        Session
	SwitchPhaseEvents map[string][]*contracts.IScalarGatewaySwitchPhase //Map by chainId
	RedeemTokenEvents map[string][]*contracts.IScalarGatewayRedeemToken
}

// Each chain store switch phase events array with 1 or 2 elements of the form:
// 1. [Preparing]
// 2. [Preparing, Executing] in the same sequence
type ChainRedeemSessions struct {
	SwitchPhaseEvents map[string][]*contracts.IScalarGatewaySwitchPhase //Map by custodian group uid
	RedeemTokenEvents map[string][]*contracts.IScalarGatewayRedeemToken
}
