package evm

import (
	"context"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/zerolog/log"
	"github.com/scalarorg/evms-indexer/config"
	"github.com/scalarorg/evms-indexer/pkg/db"
	evmAbi "github.com/scalarorg/evms-indexer/pkg/evm/abi"
	contracts "github.com/scalarorg/evms-indexer/pkg/evm/contracts/generated"
)

type EvmClient struct {
	EvmConfig      *EvmNetworkConfig
	Client         *ethclient.Client
	ChainName      string
	GatewayAddress common.Address
	Gateway        *contracts.IScalarGateway
	dbAdapter      *db.DatabaseAdapter
	subscriptions  ethereum.Subscription
	TokenAddresses map[string]string //Map token address by symbol
}

// SimpleCheckpoint represents a minimal checkpoint for recovery
type SimpleCheckpoint struct {
	BlockNumber uint64
	TxHash      string
	LogIndex    uint
}

// This function is used to adjust the rpc url to the ws prefix
// format: ws:// -> http://
// format: wss:// -> https://
// Todo: Improve this implementation

func NewEvmClients(configPath string) ([]*EvmClient, error) {
	if configPath == "" {
		return nil, fmt.Errorf("config path is not set")
	}
	evmCfgPath := fmt.Sprintf("%s/evms.json", configPath)
	configs, err := config.ReadJsonArrayConfig[EvmNetworkConfig](evmCfgPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read electrum configs: %w", err)
	}

	evmClients := []*EvmClient{}
	for _, evmConfig := range configs {
		if !evmConfig.Enable {
			continue
		}
		//Set default value for block time if is not set
		if evmConfig.BlockTime == 0 {
			evmConfig.BlockTime = 12 * time.Second
		} else {
			evmConfig.BlockTime = evmConfig.BlockTime * time.Millisecond
		}
		if evmConfig.RecoverRange == 0 {
			evmConfig.RecoverRange = 1000000
		}
		client, err := NewEvmClient(configPath, &evmConfig)
		if err != nil {
			log.Warn().Msgf("Failed to create evm client for %s: %v", evmConfig.GetName(), err)
			continue
		}
		client.TokenAddresses = make(map[string]string)
		evmClients = append(evmClients, client)
	}

	return evmClients, nil
}

func NewEvmClient(configPath string, evmConfig *EvmNetworkConfig) (*EvmClient, error) {
	// Create database adapter if DatabaseURL is set
	var dbAdapter *db.DatabaseAdapter
	if evmConfig.DatabaseURL != "" {
		var err error
		dbAdapter, err = db.NewDatabaseAdapter(evmConfig.DatabaseURL)
		if err != nil {
			return nil, fmt.Errorf("failed to create database adapter: %w", err)
		}
	}

	// Create Ethereum client
	client, err := ethclient.Dial(evmConfig.RPCUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Ethereum client: %w", err)
	}

	// Get gateway address
	gatewayAddress := common.HexToAddress(evmConfig.Gateway)

	// Create gateway contract instance
	gateway, err := contracts.NewIScalarGateway(gatewayAddress, client)
	if err != nil {
		return nil, fmt.Errorf("failed to create gateway contract instance: %w", err)
	}

	log.Info().
		Str("chainId", evmConfig.GetId()).
		Uint64("configStartBlock", evmConfig.StartBlock).
		Msg("EVM client starting block determined")

	// Create EVM client
	evmClient := &EvmClient{
		EvmConfig:      evmConfig,
		Client:         client,
		Gateway:        gateway,
		GatewayAddress: gatewayAddress,
		dbAdapter:      dbAdapter,
		TokenAddresses: make(map[string]string),
	}

	return evmClient, nil
}
func CreateGateway(networName string, gwAddr string, client *ethclient.Client) (*contracts.IScalarGateway, *common.Address, error) {
	if gwAddr == "" {
		return nil, nil, fmt.Errorf("gateway address is not set for network %s", networName)
	}
	gatewayAddress := common.HexToAddress(gwAddr)
	gateway, err := contracts.NewIScalarGateway(gatewayAddress, client)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to initialize gateway contract for network %s: %w", networName, err)
	}
	return gateway, &gatewayAddress, nil
}

func PrepareEvents() (map[string]*abi.Event, []common.Hash, error) {
	gatewayAbi, err := evmAbi.GetScalarGatewayAbi()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get scalar gateway abi: %w", err)
	}
	eventNames := []string{
		evmAbi.EVENT_EVM_CONTRACT_CALL,
		evmAbi.EVENT_EVM_CONTRACT_CALL_WITH_TOKEN,
		evmAbi.EVENT_EVM_TOKEN_SENT,
		evmAbi.EVENT_EVM_CONTRACT_CALL_APPROVED,
		evmAbi.EVENT_EVM_COMMAND_EXECUTED,
		evmAbi.EVENT_EVM_TOKEN_DEPLOYED,
		evmAbi.EVENT_EVM_SWITCHED_PHASE,
		evmAbi.EVENT_EVM_REDEEM_TOKEN,
	}
	topics := []common.Hash{}
	eventMap := make(map[string]*abi.Event)
	for _, eventName := range eventNames {
		event, ok := gatewayAbi.Events[eventName]
		if !ok {
			log.Warn().Str("eventName", eventName).Msg("Event not found in ABI")
			continue
		}
		topics = append(topics, event.ID)
		eventMap[event.ID.String()] = &event
	}
	return eventMap, topics, nil
}

func (c *EvmClient) Start(ctx context.Context) error {
	logsChan := make(chan []types.Log, 1024) //For recovery
	//logChan := make(chan types.Log, 1024)    //For subscription
	blockHeightsChan := make(chan map[uint64]uint8, 1024)
	eventMap, topics, err := PrepareEvents()
	if err != nil {
		return fmt.Errorf("failed to prepare events: %w", err)
	}
	// Process recovered logs in dependent go routine
	go c.ProcessLogsFromFetcher(ctx, eventMap, logsChan, blockHeightsChan)
	go c.fetchBlocks(ctx, blockHeightsChan)
	// Main loop, fetch logs from network
	// Each iteration, fetch log until the latest block number
	c.LoopFetchLogs(ctx, topics, logsChan)
	return nil
}

func (c *EvmClient) LoopFetchLogs(ctx context.Context, topics []common.Hash, logsChan chan<- []types.Log) {
	query := ethereum.FilterQuery{
		Addresses: []common.Address{c.GatewayAddress},
		Topics:    [][]common.Hash{topics}, // Filter by multiple event signatures
	}
	currentBlockNumber, err := c.Client.BlockNumber(context.Background())
	if err != nil {
		log.Warn().Err(err).Msg("[EvmClient] [LoopFetchLogs] cannot get current block number.")
	}
	// recover logs in dependent go routine
	go c.recoverLogs(ctx, &query, currentBlockNumber, logsChan)
	// fetch realtime logs in current routine
	c.fetchRealtimeLogs(ctx, &query, currentBlockNumber, logsChan)
}

// Get logs from current block at starting time
func (c *EvmClient) fetchRealtimeLogs(ctx context.Context, query *ethereum.FilterQuery,
	startBlock uint64, logsChan chan<- []types.Log) {
	fromBlock := startBlock
	fetchRange := uint64(1000)
	if c.EvmConfig.RecoverRange > 0 && c.EvmConfig.RecoverRange < fetchRange {
		fetchRange = c.EvmConfig.RecoverRange
	}
	currentBlockNumber, err := c.Client.BlockNumber(context.Background())
	if err != nil {
		log.Warn().Err(err).Msg("[EvmClient] [fetchRealtimeLogs] cannot get current block number.")
	}
	lastExecutionTime := time.Now()
	for {
		sleepTime := c.calculateSleepTime(lastExecutionTime, fromBlock, currentBlockNumber, fetchRange)
		if sleepTime > 0 {
			time.Sleep(sleepTime)
			currentBlockNumber, err = c.Client.BlockNumber(context.Background())
			if err != nil {
				log.Warn().Err(err).Msg("[EvmClient] [fetchRealtimeLogs] cannot get current block number.")
			}
		}

		if fromBlock <= currentBlockNumber {
			lastExecutionTime = time.Now()
			logCounter, err := c.FetchRangeLogs(ctx, query, logsChan, fromBlock, currentBlockNumber, &fetchRange, false)
			if err != nil {
				log.Warn().Err(err).Msgf("[Indexer] [Start] cannot recover events for evm client %s", c.EvmConfig.GetId())
			} else {
				log.Info().
					Str("Chain", c.EvmConfig.ID).
					Uint64("FromBlock", fromBlock).
					Uint64("FetchRange", fetchRange).
					Uint64("CurrentBlockNumber", currentBlockNumber).
					Int("TotalLogs", logCounter).
					Msgf("[EvmClient] Recover logs finished in %s", time.Since(lastExecutionTime).String())
			}
		}
		fromBlock = currentBlockNumber + 1
	}
}
func (c *EvmClient) calculateSleepTime(lastExecutionTime time.Time, fromBlock, currentBlockNumber uint64, fetchRange uint64) time.Duration {
	// There are too many new blocks, fetch logs immediately
	if fromBlock+fetchRange <= currentBlockNumber {
		return time.Duration(0)
	}
	fetchInterval := time.Duration(c.EvmConfig.FetchIntervalInMinutes) * time.Minute
	if fetchInterval == 0 {
		//Minimum fetch interval is 1 minute
		fetchInterval = time.Minute
	}
	elapsedTime := time.Since(lastExecutionTime)
	if elapsedTime < fetchInterval {
		//Sleep until the next fetch
		return fetchInterval - elapsedTime
	}
	return time.Duration(0)
}

// Get missing logs from config startBlock to current block at starting time
func (c *EvmClient) recoverLogs(ctx context.Context, query *ethereum.FilterQuery,
	currentBlockNumber uint64, logsChan chan<- []types.Log) {
	fromBlock := uint64(0)
	lastFetchedBlock := uint64(0)
	var err error
	if c.dbAdapter != nil {
		lastFetchedBlock, err = c.dbAdapter.GetLatestFetchedBlock(c.EvmConfig.GetId())
		if err != nil {
			log.Warn().Err(err).Msg("[EvmClient] [LoopFetchLogs] cannot get latest block number.")
		}
	}
	if lastFetchedBlock > 0 {
		fromBlock = lastFetchedBlock + 1
	}
	if fromBlock < c.EvmConfig.StartBlock {
		fromBlock = c.EvmConfig.StartBlock
	} else {
		fromBlock = lastFetchedBlock + 1
	}
	recoverRange := uint64(1000000)
	if c.EvmConfig.RecoverRange > 0 && c.EvmConfig.RecoverRange < 1000000 {
		recoverRange = c.EvmConfig.RecoverRange
	}

	var logCounter int
	// pass recoverRange as reference for adjustment
	if fromBlock <= currentBlockNumber {
		start := time.Now()
		logCounter, err = c.FetchRangeLogs(ctx, query, logsChan, fromBlock, currentBlockNumber, &recoverRange, true)
		if err != nil {
			log.Warn().Err(err).Msgf("[Indexer] [Start] cannot recover events for evm client %s", c.EvmConfig.GetId())
		} else {
			log.Info().
				Str("Chain", c.EvmConfig.ID).
				Uint64("FromBlock", fromBlock).
				Uint64("RecoverRange", recoverRange).
				Uint64("CurrentBlockNumber", currentBlockNumber).
				Int("TotalLogs", logCounter).
				Msgf("[EvmClient] Recover logs finished in %s", time.Since(start).String())
		}
	} else {
		log.Info().
			Str("Chain", c.EvmConfig.ID).
			Uint64("LastFetchedBlock", lastFetchedBlock).
			Uint64("StartBlock", c.EvmConfig.StartBlock).
			Uint64("CurrentBlockNumber", currentBlockNumber).
			Msgf("[EvmClient] [LoopFetchLogs] current block number is fetched. Waiting for next blocks")
	}
}

func (c *EvmClient) Stop() {
	log.Info().Msgf("[EvmClient] [Stop] stopping evm client %s", c.EvmConfig.GetId())
}

func (c *EvmClient) SubscribeWithRetry(ctx context.Context, topics []common.Hash, logsChan chan<- types.Log) {
	var retryInterval = time.Second * 12 // Initial retry interval
	maxRetryInterval := time.Minute * 5  // Maximum retry interval

	for {
		select {
		case <-ctx.Done():
			log.Info().Str("chainId", c.EvmConfig.GetId()).Msg("[EvmClient] [ConnectWithRetry] context cancelled, stopping reconnection")
			return
		default:
			// Listen to new events
			err := c.SubscribeAllEvents(ctx, topics, logsChan)
			if err != nil {
				log.Error().Err(err).Str("chainId", c.EvmConfig.GetId()).Msg("[EvmClient] [ConnectWithRetry] error subscribing to events")
			}
			// If context is cancelled, stop retrying
			if ctx.Err() != nil {
				log.Info().Str("chainId", c.EvmConfig.GetId()).Msg("[EvmClient] [ConnectWithRetry] context cancelled, stopping reconnection")
				return
			}

			// Wait before retrying
			log.Info().Str("chainId", c.EvmConfig.GetId()).Dur("retryInterval", retryInterval).Msg("[EvmClient] [ConnectWithRetry] reconnecting...")
			time.Sleep(retryInterval)

			// Exponential backoff with cap
			if retryInterval < maxRetryInterval {
				retryInterval *= 2
			}
		}
	}
}
