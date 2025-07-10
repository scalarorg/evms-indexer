package evm

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/scalarorg/evms-indexer/pkg/types"
)

const (
	COMPONENT_NAME = "EvmClient"
	RETRY_INTERVAL = time.Second * 12 // Initial retry interval
)

type Byte32 [32]uint8
type Bytes []byte
type EvmNetworkConfig struct {
	ChainID                uint64        `mapstructure:"chain_id" json:"chainId"`
	Enable                 bool          `mapstructure:"enable" json:"enable"`
	ID                     string        `mapstructure:"id" json:"id"`
	Name                   string        `mapstructure:"name" json:"name"`
	RPCUrl                 string        `mapstructure:"rpc_url" json:"rpcUrl"`
	Gateway                string        `mapstructure:"gateway" json:"gateway"`
	Finality               int           `mapstructure:"finality" json:"finality"`
	FetchIntervalInMinutes int           `mapstructure:"fetch_interval_in_minutes" json:"fetchIntervalInMinutes"`
	StartBlock             uint64        `mapstructure:"start_block" json:"startBlock"`
	BlockTime              time.Duration `mapstructure:"blockTime" json:"blockTime"` //Timeout im ms for pending txs
	MaxRetry               int           `mapstructure:"max_retry" json:"maxRetry"`
	FetchRange             uint64        `mapstructure:"fetch_range" json:"fetchRange"`     //Max block range to fetch events in single query
	RecoverRange           uint64        `mapstructure:"recover_range" json:"recoverRange"` //Max block range to recover events in single query
	FetchThread            int           `mapstructure:"fetch_thread" json:"fetchThread"`   //Number of threads to fetch blocks
	RetryDelay             time.Duration `mapstructure:"retry_delay" json:"retryDelay"`
	DatabaseURL            string        `mapstructure:"database_url" json:"database_url"` //Separate database URL for this EVM chain
}

func (c *EvmNetworkConfig) GetChainId() uint64 {
	return c.ChainID
}
func (c *EvmNetworkConfig) GetId() string {
	return c.ID
}
func (c *EvmNetworkConfig) GetName() string {
	return c.Name
}
func (c *EvmNetworkConfig) GetFamily() string {
	return types.ChainTypeEVM.String()
}

type DecodedExecuteData struct {
	//Data
	ChainId    uint64
	CommandIds [][32]byte
	Commands   []string
	Params     [][]byte
	//Proof
	Operators  []common.Address
	Weights    []uint64
	Threshold  uint64
	Signatures []string
	//Input
	Input []byte
}

type ExecuteData[T any] struct {
	//Data
	Data T
	//Proof
	Operators  []common.Address
	Weights    []uint64
	Threshold  uint64
	Signatures []string
	//Input
	Input []byte
}
type ApproveContractCall struct {
	ChainId    uint64
	CommandIds [][32]byte
	Commands   []string
	Params     [][]byte
}
type DeployToken struct {
	//Data
	Name         string
	Symbol       string
	Decimals     uint8
	Cap          uint64
	TokenAddress common.Address
	MintLimit    uint64
}

type RedeemPhase struct {
	Sequence uint64
	Phase    uint8
}
