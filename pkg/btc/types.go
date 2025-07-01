package btc

import (
	"fmt"
)

// Constants for hash sizes
const (
	SCALAR_TAG_SIZE          = 6
	VERSION_SIZE             = 1
	NETWORK_ID_SIZE          = 1
	FLAGS_SIZE               = 1
	SERVICE_TAG_SIZE         = 5
	SESSION_SEQUENCE_SIZE    = 8
	CUSTODIAN_QUORUM_SIZE    = 1
	HASH_SIZE                = 32
	DEST_CHAIN_SIZE          = 8
	DEST_TOKEN_ADDR_SIZE     = 20
	DEST_RECIPIENT_ADDR_SIZE = 20
	CUSTODIAN_GROUP_UID_SIZE = 32
	SCRIPT_PUBKEY_SIZE       = 32
)

// VaultTransactionType represents the type of vault transaction
type VaultTransactionType uint8

const (
	VaultTxTypeStaking VaultTransactionType = iota + 1
	VaultTxTypeRedeem
)

func (v VaultTransactionType) String() string {
	switch v {
	case VaultTxTypeStaking:
		return "staking"
	case VaultTxTypeRedeem:
		return "redeem"
	default:
		return "unknown"
	}
}

// VaultReturnTxOutputType represents the type of vault return transaction output
type VaultReturnTxOutputType uint8

const (
	VaultReturnTxOutputTypeRedeem VaultReturnTxOutputType = iota
	VaultReturnTxOutputTypeStaking
)

func (v VaultReturnTxOutputType) String() string {
	switch v {
	case VaultReturnTxOutputTypeStaking:
		return "staking"
	case VaultReturnTxOutputTypeRedeem:
		return "redeem"
	default:
		return "unknown"
	}
}

// DestinationChain represents a destination chain
type DestinationChain struct {
	ChainType uint8
	ChainID   uint64
}

// DestinationTokenAddress represents a destination token address (20 bytes for EVM)
type DestinationTokenAddress [20]byte

func (d DestinationTokenAddress) String() string {
	return fmt.Sprintf("0x%x", d[:])
}

// DestinationRecipientAddress represents a destination recipient address (20 bytes for EVM)
type DestinationRecipientAddress [20]byte

func (d DestinationRecipientAddress) String() string {
	return fmt.Sprintf("0x%x", d[:])
}

// VaultReturnTxOutput represents a return transaction output matching the Rust struct
type VaultReturnTxOutput struct {
	Tag                         [SCALAR_TAG_SIZE]byte
	Version                     uint8
	NetworkID                   uint8
	Flags                       uint8
	ServiceTag                  [SERVICE_TAG_SIZE]byte
	TransactionType             VaultReturnTxOutputType
	CustodianQuorum             uint8
	DestChainType               uint8
	DestChainID                 uint64
	DestinationTokenAddress     DestinationTokenAddress
	DestinationRecipientAddress DestinationRecipientAddress
	ScriptPubkey                []byte
	SessionSequence             uint64
	CustodianGroupUID           [HASH_SIZE]byte
}

type VaultChangeTxOutput struct {
	Amount  uint64
	Address string
}

// VaultTransactionData represents the parsed data from a vault transaction
type VaultTransactionData struct {
	Version                     uint8
	TxType                      VaultTransactionType
	Amount                      uint64
	StakerAddress               string
	StakerPubkey                string
	ChangeAmount                uint64
	ChangeAddress               string
	CovenantQuorum              uint8
	DestinationChain            uint64
	DestinationTokenAddress     string
	DestinationRecipientAddress string
	SessionSequence             uint64
	CustodianGroupUID           string
	ScriptPubkey                string
	//FeeOptions                  BTCFeeOpts
	RBF     bool
	TxIds   []string
	Vouts   []uint32
	Amounts []uint64
	ReqId   [32]byte
	Psbt    []byte
}

// BTCFeeOpts represents Bitcoin fee options
// type BTCFeeOpts uint8

// const (
// 	BTCFeeOptsMinimum BTCFeeOpts = iota
// 	BTCFeeOptsEconomy
// 	BTCFeeOptsHour
// 	BTCFeeOptsHalfHour
// 	BTCFeeOptsFastest
// )

// func (b BTCFeeOpts) String() string {
// 	switch b {
// 	case BTCFeeOptsMinimum:
// 		return "minimum"
// 	case BTCFeeOptsEconomy:
// 		return "economy"
// 	case BTCFeeOptsHour:
// 		return "hour"
// 	case BTCFeeOptsHalfHour:
// 		return "half_hour"
// 	case BTCFeeOptsFastest:
// 		return "fastest"
// 	default:
// 		return "unknown"
// 	}
// }
