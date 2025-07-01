package evm

import (
	"context"
	"encoding/hex"
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/zerolog/log"
	chains "github.com/scalarorg/data-models/chains"
	"github.com/scalarorg/evms-indexer/config"
	"github.com/scalarorg/evms-indexer/pkg/db"
	evmAbi "github.com/scalarorg/evms-indexer/pkg/evm/abi"
	contracts "github.com/scalarorg/evms-indexer/pkg/evm/contracts/generated"
	"github.com/scalarorg/evms-indexer/pkg/evm/parser"
	"gorm.io/gorm"
)

type EvmClient struct {
	EvmConfig         *EvmNetworkConfig
	Client            *ethclient.Client
	ChainName         string
	GatewayAddress    common.Address
	Gateway           *contracts.IScalarGateway
	dbAdapter         *db.DatabaseAdapter
	separateDB        *gorm.DB          // Separate database connection for this EVM client
	TokenAddresses    map[string]string //Map token address by symbol
	ChnlReceivedBlock chan uint64
	//MissingLogs       MissingLogs
	retryInterval time.Duration
	startingBlock uint64
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

	// Find starting block as the max block number across indexed events
	var startingBlock uint64
	if dbAdapter != nil {
		latestBlock, err := dbAdapter.GetLatestBlockFromAllEvents(evmConfig.GetId())
		if err != nil {
			log.Warn().Str("chainId", evmConfig.GetId()).
				Msg("Failed to get latest block from database, using config start block")
			startingBlock = evmConfig.StartBlock
		} else if latestBlock > 0 {
			// Start from the next block after the latest indexed block
			startingBlock = latestBlock + 1
		} else {
			startingBlock = evmConfig.StartBlock
		}
	} else {
		startingBlock = evmConfig.StartBlock
	}

	log.Info().
		Str("chainId", evmConfig.GetId()).
		Uint64("startingBlock", startingBlock).
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
		// MissingLogs: MissingLogs{
		// 	chainId: evmConfig.GetId(),
		// 	logs:    make([]types.Log, 0),
		// },
		startingBlock: startingBlock,
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

func (c *EvmClient) Start(ctx context.Context) error {
	logsChan := make(chan []types.Log, 1024)
	gatewayAbi, err := evmAbi.GetScalarGatewayAbi()
	if err != nil {
		return fmt.Errorf("failed to get scalar gateway abi: %w", err)
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
	go func() {
		err := c.RecoverAllEvents(ctx, topics, logsChan)
		if err != nil {
			log.Warn().Err(err).Msgf("[Indexer] [Start] cannot recover events for evm client %s", c.EvmConfig.GetId())
		} else {
			log.Info().Msgf("[Indexer] [Start] recovered missing events for evm client %s", c.EvmConfig.GetId())
		}
	}()
	// Process recovered logs in dependent go routine
	go c.ProcessLogs(ctx, eventMap, logsChan)
	go c.startFetchBlock()
	c.ConnectWithRetry(ctx)
	return fmt.Errorf("context cancelled")
}

func (c *EvmClient) Stop() {
	log.Info().Msgf("[EvmClient] [Stop] stopping evm client %s", c.EvmConfig.GetId())
}

func (c *EvmClient) ConnectWithRetry(ctx context.Context) {
	var retryInterval = time.Second * 12 // Initial retry interval
	maxRetryInterval := time.Minute * 5  // Maximum retry interval

	for {
		select {
		case <-ctx.Done():
			log.Info().Str("chainId", c.EvmConfig.GetId()).Msg("[EvmClient] [ConnectWithRetry] context cancelled, stopping reconnection")
			return
		default:
			// Listen to new events
			err := c.ListenToEvents(ctx)
			if err != nil {
				log.Error().Err(err).Str("chainId", c.EvmConfig.GetId()).Msg("[EvmClient] [ConnectWithRetry] error listening to events")
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
func (c *EvmClient) VerifyDeployTokens(ctx context.Context) error {
	return nil
}

func setQueryRange(query *ethereum.FilterQuery, fromBlock uint64, recoverRange uint64, currentBlockNumber uint64) {
	toBlock := fromBlock + recoverRange - 1
	if toBlock > currentBlockNumber {
		toBlock = currentBlockNumber
	}
	query.FromBlock = big.NewInt(int64(fromBlock))
	query.ToBlock = big.NewInt(int64(toBlock))
}

// RecoverAllEvents recovers all events from the latest block in the database
func (c *EvmClient) RecoverAllEvents(ctx context.Context, topics []common.Hash, logsChan chan<- []types.Log) error {
	currentBlockNumber, err := c.Client.BlockNumber(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get current block number: %w", err)
	}
	recoverRange := uint64(1000000)
	if c.EvmConfig.RecoverRange > 0 && c.EvmConfig.RecoverRange < 1000000 {
		recoverRange = c.EvmConfig.RecoverRange
	}
	fromBlock := uint64(0)
	if c.dbAdapter != nil {
		fromBlock, err = c.dbAdapter.GetLatestBlockFromAllEvents(c.EvmConfig.GetId())
		if err != nil {
			return fmt.Errorf("failed to get latest block number: %w", err)
		}
	}
	if fromBlock < c.EvmConfig.StartBlock {
		fromBlock = c.EvmConfig.StartBlock
	}
	// Set up a query for logs
	query := ethereum.FilterQuery{
		Addresses: []common.Address{c.GatewayAddress},
		Topics:    [][]common.Hash{topics}, // Filter by multiple event signatures
	}
	logCounter := 0
	log.Info().Str("Chain", c.EvmConfig.ID).
		Uint64("FromBlock", fromBlock).
		Uint64("RecoverRange", recoverRange).
		Uint64("CurrentBlockNumber", currentBlockNumber).
		Msg("[EvmClient] [RecoverAllEvents] starting recover all events")
	for fromBlock < currentBlockNumber {
		setQueryRange(&query, fromBlock, recoverRange, currentBlockNumber)
		start := time.Now()
		logs, err := c.Client.FilterLogs(context.Background(), query)
		if err != nil {
			log.Error().Err(err).Msgf("[EvmClient] [RecoverEvents] failed to filter logs")
			recoverRange, err = extractRecoverRange(err.Error())
			if err != nil {
				log.Error().Err(err).Msgf("[EvmClient] [RecoverEvents] failed to extract recover range from error message: %s", err.Error())
				return err
			}
			log.Info().Str("Chain", c.EvmConfig.ID).Uint64("Adjusted RecoverRange", recoverRange).
				Msgf("[EvmClient] [RecoverEvents] recover range extracted from error message: %d", recoverRange)
			setQueryRange(&query, fromBlock, recoverRange, currentBlockNumber)
			continue
		}
		if len(logs) > 0 {
			log.Info().Str("Chain", c.EvmConfig.ID).Msgf("[EvmClient] [RecoverEvents] found %d logs, [%d, %d]/%d, time: %s",
				len(logs), fromBlock, query.ToBlock, currentBlockNumber, time.Since(start))
			logsChan <- logs
			logCounter += len(logs)
		}
		//Set fromBlock to the next block number for next iteration
		fromBlock = query.ToBlock.Uint64() + 1
	}
	log.Info().
		Str("Chain", c.EvmConfig.ID).
		Uint64("CurrentBlockNumber", currentBlockNumber).
		Int("TotalLogs", logCounter).
		Msg("[EvmClient] [FinishRecover] recovered all events")
	return nil
}

// GetMissingEvents fetches logs for multiple event types in a single query and returns parsed events
func GetMissingEvents[T ValidEvmEvent](c *EvmClient, eventNames []string, lastCheckpoint *SimpleCheckpoint, fnCreateEventData func(types.Log) T) (map[string][]*parser.EvmEvent[T], error) {
	// Collect all event topics and create event mapping
	topics := []common.Hash{}
	eventMap := make(map[string]*abi.Event)

	for _, eventName := range eventNames {
		event, ok := evmAbi.GetEventByName(eventName)
		if !ok {
			log.Warn().Str("eventName", eventName).Msg("Event not found in ABI")
			continue
		}
		topics = append(topics, event.ID)
		eventMap[event.ID.String()] = event
	}

	if len(topics) == 0 {
		return nil, fmt.Errorf("no valid events found")
	}

	// Set up a query for logs
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(lastCheckpoint.BlockNumber)),
		Addresses: []common.Address{c.GatewayAddress},
		Topics:    [][]common.Hash{topics}, // Filter by multiple event signatures
	}

	// Fetch the logs
	logs, err := c.Client.FilterLogs(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch logs: %w", err)
	}

	log.Debug().Int("logsCount", len(logs)).Any("logs", logs).Msg("[EvmClient] [GetMissingEvents] fetched logs")

	// Group and parse events by event type
	result := make(map[string][]*parser.EvmEvent[T])

	for _, receiptLog := range logs {
		if receiptLog.BlockNumber < lastCheckpoint.BlockNumber ||
			(receiptLog.BlockNumber == lastCheckpoint.BlockNumber && receiptLog.Index <= lastCheckpoint.LogIndex) {
			log.Info().Uint64("receiptLogBlockNumber", receiptLog.BlockNumber).
				Uint("receiptLogIndex", receiptLog.Index).
				Msg("[EvmClient] [GetMissingEvents] skip log")
			continue
		}

		// Find the event type for this log
		if len(receiptLog.Topics) == 0 {
			continue
		}

		topic := receiptLog.Topics[0].String()
		event, exists := eventMap[topic]
		if !exists {
			log.Warn().Str("topic", topic).Msg("Unknown event topic")
			continue
		}

		// Parse the event data
		var eventData = fnCreateEventData(receiptLog)
		err := parser.ParseEventData(&receiptLog, event, eventData)
		if err == nil {
			evmEvent := parser.CreateEvmEventFromArgs[T](eventData, &receiptLog)
			result[event.Name] = append(result[event.Name], evmEvent)
			log.Info().
				Str("eventName", event.Name).
				Uint64("blockNumber", evmEvent.BlockNumber).
				Str("txHash", evmEvent.Hash).
				Uint("txIndex", evmEvent.TxIndex).
				Uint("logIndex", evmEvent.LogIndex).
				Msg("[EvmClient] [GetMissingEvents] parsing successfully.")
		} else {
			log.Error().Err(err).Str("eventName", event.Name).Msg("[EvmClient] [GetMissingEvents] failed to unpack log data")
		}
	}

	// Log summary
	totalEvents := 0
	for eventName, events := range result {
		totalEvents += len(events)
		log.Info().
			Str("eventName", eventName).
			Int("eventCount", len(events)).
			Msg("[EvmClient] [GetMissingEvents] parsed events")
	}

	log.Info().
		Int("totalLogs", len(logs)).
		Int("totalEvents", totalEvents).
		Any("query", query).
		Any("lastCheckpoint", lastCheckpoint).
		Msg("[EvmClient] [GetMissingEvents] completed")

	return result, nil
}

// GetMissingEventsSingle is a convenience function for backward compatibility
// It fetches logs for a single event type
func GetMissingEventsSingle[T ValidEvmEvent](c *EvmClient, eventName string, lastCheckpoint *SimpleCheckpoint, fnCreateEventData func(types.Log) T) ([]*parser.EvmEvent[T], error) {
	events, err := GetMissingEvents(c, []string{eventName}, lastCheckpoint, fnCreateEventData)
	if err != nil {
		return nil, err
	}

	if eventList, exists := events[eventName]; exists {
		return eventList, nil
	}

	return []*parser.EvmEvent[T]{}, nil
}

func (c *EvmClient) ListenToEvents(ctx context.Context) error {
	c.retryInterval = RETRY_INTERVAL

	events := []struct {
		name  string
		watch func(context.Context) error
	}{
		// {EVENT_EVM_CONTRACT_CALL, func(ctx context.Context) error {
		// 	return WatchForEvent[*contracts.IScalarGatewayContractCall](c, ctx, EVENT_EVM_CONTRACT_CALL)
		// }},
		{evmAbi.EVENT_EVM_TOKEN_SENT, func(ctx context.Context) error {
			return WatchForEvent[*contracts.IScalarGatewayTokenSent](c, ctx, evmAbi.EVENT_EVM_TOKEN_SENT)
		}},
		{evmAbi.EVENT_EVM_CONTRACT_CALL_WITH_TOKEN, func(ctx context.Context) error {
			return WatchForEvent[*contracts.IScalarGatewayContractCallWithToken](c, ctx, evmAbi.EVENT_EVM_CONTRACT_CALL_WITH_TOKEN)
		}},
		{evmAbi.EVENT_EVM_CONTRACT_CALL_APPROVED, func(ctx context.Context) error {
			return WatchForEvent[*contracts.IScalarGatewayContractCallApproved](c, ctx, evmAbi.EVENT_EVM_CONTRACT_CALL_APPROVED)
		}},
		{evmAbi.EVENT_EVM_COMMAND_EXECUTED, func(ctx context.Context) error {
			return WatchForEvent[*contracts.IScalarGatewayExecuted](c, ctx, evmAbi.EVENT_EVM_COMMAND_EXECUTED)
		}},
		{evmAbi.EVENT_EVM_TOKEN_DEPLOYED, func(ctx context.Context) error {
			return WatchForEvent[*contracts.IScalarGatewayTokenDeployed](c, ctx, evmAbi.EVENT_EVM_TOKEN_DEPLOYED)
		}},
		{evmAbi.EVENT_EVM_SWITCHED_PHASE, func(ctx context.Context) error {
			return WatchForEvent[*contracts.IScalarGatewaySwitchPhase](c, ctx, evmAbi.EVENT_EVM_SWITCHED_PHASE)
		}},
		{evmAbi.EVENT_EVM_REDEEM_TOKEN, func(ctx context.Context) error {
			return WatchForEvent[*contracts.IScalarGatewayRedeemToken](c, ctx, evmAbi.EVENT_EVM_REDEEM_TOKEN)
		}},
	}

	for _, event := range events {
		go func(event struct {
			name  string
			watch func(context.Context) error
		}) {
			for {
				err := event.watch(ctx)
				if err != nil {
					log.Error().Err(err).Str("eventName", event.name).Msg("[EvmClient] [ListenToEvents] error watching event")
				}
				select {
				case <-ctx.Done():
					return
				default:
					time.Sleep(c.retryInterval)
				}
			}
		}(event)
	}

	return nil
}

type ValidWatchEvent interface {
	*contracts.IScalarGatewayTokenSent |
		*contracts.IScalarGatewayContractCallWithToken |
		// *contracts.IScalarGatewayContractCall |
		*contracts.IScalarGatewayContractCallApproved |
		*contracts.IScalarGatewayExecuted |
		*contracts.IScalarGatewayTokenDeployed |
		*contracts.IScalarGatewaySwitchPhase |
		*contracts.IScalarGatewayRedeemToken
}

const (
	baseDelay    = 5 * time.Second
	maxDelay     = 2 * time.Minute
	maxAttempts  = math.MaxUint64
	jitterFactor = 0.2 // 20% jitter
)

func WatchForEvent[T ValidWatchEvent](c *EvmClient, ctx context.Context, eventName string) error {
	// Use the client's starting block that was determined during initialization
	watchOpts := bind.WatchOpts{Start: &c.startingBlock, Context: ctx}

	sink := make(chan T)

	subscription, err := setupSubscription(c, &watchOpts, sink, eventName)
	if err != nil {
		return err
	}
	defer subscription.Unsubscribe()

	log.Info().Str("eventName", eventName).Str("chainId", c.EvmConfig.GetId()).Msgf("[EvmClient] [watchEVMTokenSent] success")

	for {
		select {
		case err := <-subscription.Err():
			log.Error().Err(err).Msgf("[EvmClient] [WatchForEvent] error with subscription for %s, attempting reconnect", eventName)

			subscription, err = reconnectWithBackoff(c, &watchOpts, sink, eventName)
			if err != nil {
				return fmt.Errorf("failed to reconnect: %w", err)
			}

		case <-watchOpts.Context.Done():
			return nil

		case event := <-sink:
			err := handleEvent(c, eventName, event)
			if err != nil {
				log.Error().Err(err).Msgf("[EvmClient] [WatchForEvent] error handling %s event", eventName)
			} else {
				log.Info().Any("event", event).Msgf("[EvmClient] [WatchForEvent] handled %s event", eventName)
			}
		}
	}
}

func reconnectWithBackoff[T ValidWatchEvent](c *EvmClient, watchOpts *bind.WatchOpts, sink chan T, eventName string) (ethereum.Subscription, error) {
	delay := baseDelay
	for attempt := uint64(1); attempt <= maxAttempts; attempt++ {
		jitter := time.Duration(rand.Float64() * float64(delay) * jitterFactor)
		retryDelay := delay + jitter

		log.Info().
			Uint64("attempt", attempt).
			Dur("delay", retryDelay).
			Msgf("[EvmClient] [WatchForEvent] attempting reconnection for %s", eventName)

		select {
		case <-watchOpts.Context.Done():
			return nil, watchOpts.Context.Err()
		case <-time.After(retryDelay):
			subscription, err := setupSubscription(c, watchOpts, sink, eventName)
			if err == nil {
				log.Info().Msgf("[EvmClient] [WatchForEvent] successfully reconnected for %s", eventName)
				return subscription, nil
			}

			log.Error().
				Err(err).
				Uint64("attempt", attempt).
				Msgf("[EvmClient] [WatchForEvent] reconnection failed for %s", eventName)

			delay = time.Duration(float64(delay) * 2)
			if delay > maxDelay {
				delay = maxDelay
			}
		}
	}

	return nil, fmt.Errorf("failed to reconnect after %d attempts", uint64(maxAttempts))
}

func setupSubscription[T ValidWatchEvent](c *EvmClient, watchOpts *bind.WatchOpts, sink chan T, eventName string) (ethereum.Subscription, error) {
	sinkInterface := any(sink)

	switch eventName {
	case evmAbi.EVENT_EVM_TOKEN_SENT:
		return c.Gateway.WatchTokenSent(watchOpts, sinkInterface.(chan *contracts.IScalarGatewayTokenSent), nil)
	case evmAbi.EVENT_EVM_CONTRACT_CALL_WITH_TOKEN:
		return c.Gateway.WatchContractCallWithToken(watchOpts, sinkInterface.(chan *contracts.IScalarGatewayContractCallWithToken), nil, nil)
	case evmAbi.EVENT_EVM_CONTRACT_CALL_APPROVED:
		return c.Gateway.WatchContractCallApproved(watchOpts, sinkInterface.(chan *contracts.IScalarGatewayContractCallApproved), nil, nil, nil)
	case evmAbi.EVENT_EVM_COMMAND_EXECUTED:
		return c.Gateway.WatchExecuted(watchOpts, sinkInterface.(chan *contracts.IScalarGatewayExecuted), nil)
	case evmAbi.EVENT_EVM_SWITCHED_PHASE:
		return c.Gateway.WatchSwitchPhase(watchOpts, sinkInterface.(chan *contracts.IScalarGatewaySwitchPhase), nil, nil)
	case evmAbi.EVENT_EVM_TOKEN_DEPLOYED:
		return c.Gateway.WatchTokenDeployed(watchOpts, sinkInterface.(chan *contracts.IScalarGatewayTokenDeployed))
	case evmAbi.EVENT_EVM_REDEEM_TOKEN:
		return c.Gateway.WatchRedeemToken(watchOpts, sinkInterface.(chan *contracts.IScalarGatewayRedeemToken), nil, nil, nil)
	default:
		return nil, fmt.Errorf("[EvmClient] [setupSubscription] unsupported event type for %s, event: %v", eventName, (*T)(nil))
	}
}

func handleEvent(c *EvmClient, eventName string, event any) error {
	switch eventName {
	case evmAbi.EVENT_EVM_TOKEN_SENT:
		if evt, ok := event.(*contracts.IScalarGatewayTokenSent); ok {
			return c.HandleTokenSent(evt)
		}
		return fmt.Errorf("cannot parse event %s: %T to %T", eventName, event, (*contracts.IScalarGatewayTokenSent)(nil))
	case evmAbi.EVENT_EVM_CONTRACT_CALL_WITH_TOKEN:
		if evt, ok := event.(*contracts.IScalarGatewayContractCallWithToken); ok {
			return c.HandleContractCallWithToken(evt)
		}
		return fmt.Errorf("cannot parse event %s: %T to %T", eventName, event, (*contracts.IScalarGatewayContractCallWithToken)(nil))
	case evmAbi.EVENT_EVM_CONTRACT_CALL_APPROVED:
		if evt, ok := event.(*contracts.IScalarGatewayContractCallApproved); ok {
			return c.HandleContractCallApproved(evt)
		}
		return fmt.Errorf("cannot parse event %s: %T to %T", eventName, event, (*contracts.IScalarGatewayContractCallApproved)(nil))
	case evmAbi.EVENT_EVM_COMMAND_EXECUTED:
		if evt, ok := event.(*contracts.IScalarGatewayExecuted); ok {
			return c.HandleCommandExecuted(evt)
		}
		return fmt.Errorf("cannot parse event %s: %T to %T", eventName, event, (*contracts.IScalarGatewayExecuted)(nil))
	case evmAbi.EVENT_EVM_TOKEN_DEPLOYED:
		if evt, ok := event.(*contracts.IScalarGatewayTokenDeployed); ok {
			return c.HandleTokenDeployed(evt)
		}
		return fmt.Errorf("cannot parse event %s: %T to %T", eventName, event, (*contracts.IScalarGatewayTokenDeployed)(nil))
	case evmAbi.EVENT_EVM_SWITCHED_PHASE:
		if evt, ok := event.(*contracts.IScalarGatewaySwitchPhase); ok {
			return c.HandleSwitchPhase(evt)
		}
		return fmt.Errorf("cannot parse event %s: %T to %T", eventName, event, (*contracts.IScalarGatewaySwitchPhase)(nil))
	case evmAbi.EVENT_EVM_REDEEM_TOKEN:
		if evt, ok := event.(*contracts.IScalarGatewayRedeemToken); ok {
			return c.HandleRedeemToken(evt)
		}
		return fmt.Errorf("cannot parse event %s: %T to %T", eventName, event, (*contracts.IScalarGatewayRedeemToken)(nil))
	}
	return fmt.Errorf("invalid event type for %s: %T", eventName, event)
}

func (c *EvmClient) handleEventLog(event *abi.Event, txLog types.Log) error {
	switch event.Name {
	case evmAbi.EVENT_EVM_TOKEN_SENT:
		tokenSent := &contracts.IScalarGatewayTokenSent{
			Raw: txLog,
		}
		err := parser.ParseEventData(&txLog, event, tokenSent)
		if err != nil {
			return fmt.Errorf("failed to parse event %s: %w", event.Name, err)
		}
		return c.HandleTokenSent(tokenSent)
	case evmAbi.EVENT_EVM_CONTRACT_CALL_WITH_TOKEN:
		contractCallWithToken := &contracts.IScalarGatewayContractCallWithToken{
			Raw: txLog,
		}
		err := parser.ParseEventData(&txLog, event, contractCallWithToken)
		if err != nil {
			return fmt.Errorf("failed to parse event %s: %w", event.Name, err)
		}
		return c.HandleContractCallWithToken(contractCallWithToken)
	case evmAbi.EVENT_EVM_REDEEM_TOKEN:
		redeemToken := &contracts.IScalarGatewayRedeemToken{
			Raw: txLog,
		}
		err := parser.ParseEventData(&txLog, event, redeemToken)
		if err != nil {
			return fmt.Errorf("failed to parse event %s: %w", event.Name, err)
		}
		return c.HandleRedeemToken(redeemToken)
	case evmAbi.EVENT_EVM_CONTRACT_CALL:
		//return c.HandleContractCall(txLog)
		return fmt.Errorf("not implemented")
	case evmAbi.EVENT_EVM_CONTRACT_CALL_APPROVED:
		contractCallApproved := &contracts.IScalarGatewayContractCallApproved{
			Raw: txLog,
		}
		err := parser.ParseEventData(&txLog, event, contractCallApproved)
		if err != nil {
			return fmt.Errorf("failed to parse event %s: %w", event.Name, err)
		}
		return c.HandleContractCallApproved(contractCallApproved)
	case evmAbi.EVENT_EVM_COMMAND_EXECUTED:
		executed := &contracts.IScalarGatewayExecuted{
			Raw: txLog,
		}
		err := parser.ParseEventData(&txLog, event, executed)
		if err != nil {
			return fmt.Errorf("failed to parse event %s: %w", event.Name, err)
		}
		return c.HandleCommandExecuted(executed)
	case evmAbi.EVENT_EVM_TOKEN_DEPLOYED:
		tokenDeployed := &contracts.IScalarGatewayTokenDeployed{
			Raw: txLog,
		}
		err := parser.ParseEventData(&txLog, event, tokenDeployed)
		if err != nil {
			return fmt.Errorf("failed to parse event %s: %w", event.Name, err)
		}
		return c.HandleTokenDeployed(tokenDeployed)
	case evmAbi.EVENT_EVM_SWITCHED_PHASE:
		switchedPhase := &contracts.IScalarGatewaySwitchPhase{
			Raw: txLog,
		}
		err := parser.ParseEventData(&txLog, event, switchedPhase)
		if err != nil {
			return fmt.Errorf("failed to parse event %s: %w", event.Name, err)
		}
		return c.HandleSwitchPhase(switchedPhase)
	default:
		return fmt.Errorf("invalid event type for %s: %T", event.Name, txLog)
	}
}

func (c *EvmClient) startFetchBlock() {
	for blockNumber := range c.ChnlReceivedBlock {
		if c.dbAdapter == nil {
			log.Error().Msgf("[EvmClient] [startFetchBlock] db adapter is not set")
			continue
		}
		log.Info().Str("ChainId", c.EvmConfig.ID).Uint64("BlockNumber", blockNumber).Msg("[EvmClient] Fetch block by number")
		block, err := c.Client.BlockByNumber(context.Background(), big.NewInt(int64(blockNumber)))
		if err != nil {
			log.Error().Err(err).Msgf("[EvmClient] [startFetchBlock] failed to fetch block %d", blockNumber)
		} else if block == nil {
			log.Error().Msgf("[EvmClient] [startFetchBlock] block %d not found", blockNumber)
		} else {
			log.Info().Uint64("BlockNumber", block.NumberU64()).
				Str("BlockHash", hex.EncodeToString(block.Hash().Bytes())).
				Uint64("BlockTime", block.Time()).
				Msgf("[EvmClient] [startFetchBlock] found block")
			blockNumber := block.NumberU64()
			blockHeader := &chains.BlockHeader{
				Chain:       c.EvmConfig.GetId(),
				BlockNumber: blockNumber,
				BlockHash:   hex.EncodeToString(block.Hash().Bytes()),
				BlockTime:   block.Time(),
			}
			err = c.dbAdapter.CreateBlockHeader(blockHeader)
			if err != nil {
				log.Error().Err(err).Msgf("[EvmClient] [startFetchBlock] failed to save block header %d", blockNumber)
			}
		}
	}
}
func (c *EvmClient) fetchBlockHeader(blockNumber uint64) error {
	blockHeader, err := c.dbAdapter.FindBlockHeader(c.EvmConfig.GetId(), blockNumber)
	if err == nil && blockHeader != nil {
		//log.Info().Any("blockHeader", blockHeader).Msgf("[EvmClient] [startFetchBlock] block header already exists")
		return nil
	}
	c.ChnlReceivedBlock <- blockNumber
	return nil
}

// GetDatabase returns the appropriate database connection
// If separateDB is set, it returns that, otherwise returns the shared dbAdapter
func (c *EvmClient) GetDatabase() *gorm.DB {
	if c.separateDB != nil {
		return c.separateDB
	}
	return c.dbAdapter.PostgresClient
}

// GetDatabaseAdapter returns the appropriate database adapter
// If separateDB is set, it returns a new adapter with that connection, otherwise returns the shared dbAdapter
func (c *EvmClient) GetDatabaseAdapter() *db.DatabaseAdapter {
	if c.separateDB != nil {
		return &db.DatabaseAdapter{
			PostgresClient: c.separateDB,
		}
	}
	return c.dbAdapter
}

// GetMissingEventsForMultipleEvents fetches logs for multiple event types in a single query
func (c *EvmClient) GetMissingEventsForMultipleEvents(eventNames []string, fromBlock, toBlock uint64) (map[string][]types.Log, error) {
	// Collect all event topics
	topics := []common.Hash{}
	eventMap := make(map[string]*abi.Event)

	for _, eventName := range eventNames {
		event, ok := evmAbi.GetEventByName(eventName)
		if !ok {
			log.Warn().Str("eventName", eventName).Msg("Event not found in ABI")
			continue
		}
		topics = append(topics, event.ID)
		eventMap[event.ID.String()] = event
	}

	if len(topics) == 0 {
		return nil, fmt.Errorf("no valid events found")
	}

	// Set up a query for logs
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(fromBlock)),
		ToBlock:   big.NewInt(int64(toBlock)),
		Addresses: []common.Address{c.GatewayAddress},
		Topics:    [][]common.Hash{topics}, // Filter by multiple event signatures
	}

	// Fetch the logs
	logs, err := c.Client.FilterLogs(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch logs: %w", err)
	}

	log.Info().
		Str("chainId", c.EvmConfig.GetId()).
		Int("totalLogs", len(logs)).
		Uint64("fromBlock", fromBlock).
		Uint64("toBlock", toBlock).
		Msg("Fetched logs for multiple events")

	// Group logs by event type
	groupedLogs := make(map[string][]types.Log)
	for _, log := range logs {
		if len(log.Topics) == 0 {
			continue
		}

		topic := log.Topics[0].String()
		if event, exists := eventMap[topic]; exists {
			groupedLogs[event.Name] = append(groupedLogs[event.Name], log)
		}
	}

	return groupedLogs, nil
}

// RecoverAllEventsBatch recovers all missing events in a single batch operation
func (c *EvmClient) RecoverAllEventsBatch(ctx context.Context) error {
	log.Info().Msg("[EvmClient] [RecoverAllEventsBatch] Starting batch recovery of all events")

	// Get the latest indexed block number
	latestBlock, err := c.dbAdapter.GetLatestBlockFromAllEvents(c.EvmConfig.GetId())
	if err != nil {
		return fmt.Errorf("failed to get latest block number: %w", err)
	}

	// Create checkpoint from the latest indexed block
	lastCheckpoint := &SimpleCheckpoint{
		BlockNumber: latestBlock,
		TxHash:      "",
		LogIndex:    0,
	}

	// Process TokenSent events
	if err := c.recoverTokenSentEvents(lastCheckpoint); err != nil {
		return fmt.Errorf("failed to recover token sent events: %w", err)
	}

	// Process ContractCallWithToken events
	if err := c.recoverContractCallWithTokenEvents(lastCheckpoint); err != nil {
		return fmt.Errorf("failed to recover contract call with token events: %w", err)
	}

	// Process ContractCallApproved events
	if err := c.recoverContractCallApprovedEvents(lastCheckpoint); err != nil {
		return fmt.Errorf("failed to recover contract call approved events: %w", err)
	}

	// Process Executed events
	if err := c.recoverExecutedEvents(lastCheckpoint); err != nil {
		return fmt.Errorf("failed to recover executed events: %w", err)
	}

	log.Info().Msg("[EvmClient] [RecoverAllEventsBatch] Batch recovery completed successfully")
	return nil
}

// recoverTokenSentEvents recovers TokenSent events
func (c *EvmClient) recoverTokenSentEvents(lastCheckpoint *SimpleCheckpoint) error {
	events, err := GetMissingEventsSingle[*contracts.IScalarGatewayTokenSent](c, evmAbi.EVENT_EVM_TOKEN_SENT, lastCheckpoint, func(log types.Log) *contracts.IScalarGatewayTokenSent {
		return &contracts.IScalarGatewayTokenSent{Raw: log}
	})
	if err != nil {
		return err
	}

	if len(events) == 0 {
		log.Info().Str("eventName", evmAbi.EVENT_EVM_TOKEN_SENT).Msg("No missing events found")
		return nil
	}

	return c.processTokenSentEventsBatch(events)
}

// recoverContractCallWithTokenEvents recovers ContractCallWithToken events
func (c *EvmClient) recoverContractCallWithTokenEvents(lastCheckpoint *SimpleCheckpoint) error {
	events, err := GetMissingEventsSingle[*contracts.IScalarGatewayContractCallWithToken](c, evmAbi.EVENT_EVM_CONTRACT_CALL_WITH_TOKEN, lastCheckpoint, func(log types.Log) *contracts.IScalarGatewayContractCallWithToken {
		return &contracts.IScalarGatewayContractCallWithToken{Raw: log}
	})
	if err != nil {
		return err
	}

	if len(events) == 0 {
		log.Info().Str("eventName", evmAbi.EVENT_EVM_CONTRACT_CALL_WITH_TOKEN).Msg("No missing events found")
		return nil
	}

	return c.processContractCallWithTokenEventsBatch(events)
}

// recoverContractCallApprovedEvents recovers ContractCallApproved events
func (c *EvmClient) recoverContractCallApprovedEvents(lastCheckpoint *SimpleCheckpoint) error {
	events, err := GetMissingEventsSingle[*contracts.IScalarGatewayContractCallApproved](c, evmAbi.EVENT_EVM_CONTRACT_CALL_APPROVED, lastCheckpoint, func(log types.Log) *contracts.IScalarGatewayContractCallApproved {
		return &contracts.IScalarGatewayContractCallApproved{Raw: log}
	})
	if err != nil {
		return err
	}

	if len(events) == 0 {
		log.Info().Str("eventName", evmAbi.EVENT_EVM_CONTRACT_CALL_APPROVED).Msg("No missing events found")
		return nil
	}

	return c.processContractCallApprovedEventsBatch(events)
}

// recoverExecutedEvents recovers Executed events
func (c *EvmClient) recoverExecutedEvents(lastCheckpoint *SimpleCheckpoint) error {
	events, err := GetMissingEventsSingle[*contracts.IScalarGatewayExecuted](c, evmAbi.EVENT_EVM_COMMAND_EXECUTED, lastCheckpoint, func(log types.Log) *contracts.IScalarGatewayExecuted {
		return &contracts.IScalarGatewayExecuted{Raw: log}
	})
	if err != nil {
		return err
	}

	if len(events) == 0 {
		log.Info().Str("eventName", evmAbi.EVENT_EVM_COMMAND_EXECUTED).Msg("No missing events found")
		return nil
	}

	return c.processExecutedEventsBatch(events)
}

// processTokenSentEventsBatch processes TokenSent events in batch
func (c *EvmClient) processTokenSentEventsBatch(events []*parser.EvmEvent[*contracts.IScalarGatewayTokenSent]) error {
	var tokenSents []*chains.TokenSent

	for _, evt := range events {
		tokenSent := parser.TokenSentEvent2Model(c.EvmConfig.GetId(), evt.Args)
		tokenSents = append(tokenSents, tokenSent)
	}

	if len(tokenSents) > 0 {
		return c.dbAdapter.BatchSaveTokenSents(tokenSents)
	}
	return nil
}

// processContractCallWithTokenEventsBatch processes ContractCallWithToken events in batch
func (c *EvmClient) processContractCallWithTokenEventsBatch(events []*parser.EvmEvent[*contracts.IScalarGatewayContractCallWithToken]) error {
	var contractCalls []*chains.ContractCallWithToken

	for _, evt := range events {
		contractCall, err := c.ContractCallWithToken2Model(evt.Args)
		if err != nil {
			log.Error().Err(err).Msg("Failed to convert ContractCallWithToken event to model")
			continue
		}
		contractCalls = append(contractCalls, contractCall)
	}

	if len(contractCalls) > 0 {
		return c.dbAdapter.BatchCreateContractCallsWithToken(contractCalls)
	}
	return nil
}

// processContractCallApprovedEventsBatch processes ContractCallApproved events in batch
func (c *EvmClient) processContractCallApprovedEventsBatch(events []*parser.EvmEvent[*contracts.IScalarGatewayContractCallApproved]) error {
	for _, evt := range events {
		contractCallApproved := parser.ContractCallApprovedEvent2Model(c.EvmConfig.GetId(), evt.Args)

		// Save each ContractCallApproved individually since there's no batch method
		if err := c.dbAdapter.SaveSingleValue(&contractCallApproved); err != nil {
			log.Error().Err(err).Msg("Failed to save ContractCallApproved")
			continue
		}
	}

	return nil
}

// processExecutedEventsBatch processes Executed events in batch
func (c *EvmClient) processExecutedEventsBatch(events []*parser.EvmEvent[*contracts.IScalarGatewayExecuted]) error {
	var commandExecuteds []*chains.CommandExecuted

	for _, evt := range events {
		cmdExecuted := parser.CommandExecutedEvent2Model(c.EvmConfig.GetId(), evt.Args)
		commandExecuteds = append(commandExecuteds, cmdExecuted)
	}

	if len(commandExecuteds) > 0 {
		return c.dbAdapter.BatchSaveCommandExecuted(commandExecuteds)
	}
	return nil
}
