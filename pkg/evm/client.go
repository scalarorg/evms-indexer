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
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/rs/zerolog/log"
	chains "github.com/scalarorg/data-models/chains"
	"github.com/scalarorg/data-models/scalarnet"
	"github.com/scalarorg/evms-indexer/config"
	"github.com/scalarorg/evms-indexer/pkg/db"
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
	MissingLogs       MissingLogs
	retryInterval     time.Duration
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

	evmClients := make([]*EvmClient, 0, len(configs))
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
	// Setup
	ctx := context.Background()
	log.Info().Any("evmConfig", evmConfig).Msgf("[EvmClient] [NewEvmClient] connecting to EVM network")
	// Connect to a test network
	rpc, err := rpc.DialContext(ctx, evmConfig.RPCUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to EVM network %s: %w", evmConfig.Name, err)
	}
	client := ethclient.NewClient(rpc)
	gateway, gatewayAddress, err := CreateGateway(evmConfig.Name, evmConfig.Gateway, client)
	if err != nil {
		return nil, fmt.Errorf("failed to create gateway for network %s: %w", evmConfig.Name, err)
	}

	// Initialize database adapter
	var dbAdapter *db.DatabaseAdapter
	if evmConfig.DatabaseURL != "" {
		dbAdapter, err = db.NewDatabaseAdapter(evmConfig.DatabaseURL)
		if err != nil {
			return nil, fmt.Errorf("failed to create database adapter for network %s: %w", evmConfig.Name, err)
		}
	}

	evmClient := &EvmClient{
		EvmConfig:         evmConfig,
		Client:            client,
		GatewayAddress:    *gatewayAddress,
		Gateway:           gateway,
		dbAdapter:         dbAdapter,
		separateDB:        nil,
		ChnlReceivedBlock: make(chan uint64),
		MissingLogs: MissingLogs{
			chainId: evmConfig.GetId(),
			logs:    make([]types.Log, 0),
		},
		retryInterval: RETRY_INTERVAL,
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
	c.startFetchBlock()
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

/*
 * Recover initiated events
 */
func (c *EvmClient) RecoverInitiatedEvents(ctx context.Context) error {
	//Recover ContractCall events
	if err := RecoverEvent[*contracts.IScalarGatewayContractCall](c, ctx,
		EVENT_EVM_CONTRACT_CALL, func(log types.Log) *contracts.IScalarGatewayContractCall {
			return &contracts.IScalarGatewayContractCall{
				Raw: log,
			}
		}); err != nil {
		log.Error().Err(err).Str("eventName", EVENT_EVM_CONTRACT_CALL).Msg("failed to recover missing events")
	}
	if err := RecoverEvent[*contracts.IScalarGatewayContractCallWithToken](c, ctx,
		EVENT_EVM_CONTRACT_CALL_WITH_TOKEN, func(log types.Log) *contracts.IScalarGatewayContractCallWithToken {
			return &contracts.IScalarGatewayContractCallWithToken{
				Raw: log,
			}
		}); err != nil {
		log.Error().Err(err).Str("eventName", EVENT_EVM_CONTRACT_CALL_WITH_TOKEN).Msg("failed to recover missing events")
	}
	if err := RecoverEvent[*contracts.IScalarGatewayTokenSent](c, ctx,
		EVENT_EVM_TOKEN_SENT, func(log types.Log) *contracts.IScalarGatewayTokenSent {
			return &contracts.IScalarGatewayTokenSent{
				Raw: log,
			}
		}); err != nil {
		log.Error().Err(err).Str("eventName", EVENT_EVM_TOKEN_SENT).Msg("failed to recover missing events")
	}
	return nil
}
func (c *EvmClient) RecoverApprovedEvents(ctx context.Context) error {
	err := RecoverEvent[*contracts.IScalarGatewayContractCallApproved](c, ctx,
		EVENT_EVM_CONTRACT_CALL_APPROVED, func(log types.Log) *contracts.IScalarGatewayContractCallApproved {
			return &contracts.IScalarGatewayContractCallApproved{
				Raw: log,
			}
		})
	if err != nil {
		log.Error().Err(err).Str("eventName", EVENT_EVM_CONTRACT_CALL_APPROVED).Msg("failed to recover missing events")
	}
	return err
}

func (c *EvmClient) RecoverExecutedEvents(ctx context.Context) error {
	err := RecoverEvent[*contracts.IScalarGatewayExecuted](c, ctx,
		EVENT_EVM_COMMAND_EXECUTED, func(log types.Log) *contracts.IScalarGatewayExecuted {
			return &contracts.IScalarGatewayExecuted{
				Raw: log,
			}
		})
	if err != nil {
		log.Error().Err(err).Str("eventName", EVENT_EVM_COMMAND_EXECUTED).Msg("failed to recover missing events")
	}
	return err
}

// Try to recover missing events from the last checkpoint block number to the current block number
func RecoverEvent[T ValidEvmEvent](c *EvmClient, ctx context.Context, eventName string, fnCreateEventData func(types.Log) T) error {
	// Get the latest block number from all event tables instead of using checkpoints
	latestBlockNumber, err := c.dbAdapter.GetLatestBlockFromAllEvents(c.EvmConfig.GetId())
	if err != nil {
		log.Warn().Str("chainId", c.EvmConfig.GetId()).
			Str("eventName", eventName).
			Msg("[EvmClient] [getLatestBlockFromAllEvents] using start block as fallback")
		latestBlockNumber = c.EvmConfig.StartBlock
	} else if latestBlockNumber > 0 {
		// Get the block before the maximum one
		latestBlockNumber = latestBlockNumber - 1
	}

	// Get current block number
	blockNumber, err := c.Client.BlockNumber(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get current block number: %w", err)
	}
	log.Info().Str("Chain", c.EvmConfig.ID).Uint64("Latest BlockNumber", latestBlockNumber).Uint64("Current BlockNumber", blockNumber).Msg("[EvmClient] [RecoverEvent]")

	// Create a checkpoint starting from the latest block in our database
	lastCheckpoint := &scalarnet.EventCheckPoint{
		ChainName:   c.EvmConfig.GetId(),
		EventName:   eventName,
		BlockNumber: latestBlockNumber,
		TxHash:      "",
		LogIndex:    0,
		EventKey:    "",
	}

	// Recover missing events from the latest block number to the current block number
	for {
		missingEvents, err := GetMissingEvents[T](c, eventName, lastCheckpoint, fnCreateEventData)
		if err != nil {
			return fmt.Errorf("failed to get missing events: %w", err)
		}
		if len(missingEvents) == 0 {
			log.Info().Str("eventName", eventName).Msg("[EvmClient] [RecoverEvent] no more missing events")
			break
		}
		// Process the missing events
		for _, event := range missingEvents {
			log.Debug().
				Str("chainId", c.EvmConfig.GetId()).
				Str("eventName", eventName).
				Str("txHash", event.Hash).
				Msg("[EvmClient] [RecoverEvent] start handling missing event")
			err := handleEvent(c, eventName, event.Args)
			// Update the last checkpoint value for next iteration
			lastCheckpoint.BlockNumber = event.BlockNumber
			lastCheckpoint.LogIndex = event.LogIndex
			lastCheckpoint.TxHash = event.Hash
			// If handleEvent success, the last checkpoint is updated within the function
			// So we need to update the last checkpoint only if handleEvent failed
			if err != nil {
				log.Error().Err(err).Msg("[EvmClient] [RecoverEvent] failed to handle event")
			}
			// Store the last checkpoint value into db,
			// this can be performed only once when we finish recover, but some thing can break recover process, so we store state immediately
			err = c.dbAdapter.UpdateLastEventCheckPoint(lastCheckpoint)
			if err != nil {
				log.Error().Err(err).Msg("[EvmClient] [RecoverEvent] update last checkpoint failed")
			}
		}
		// Try to get more missing events in the next iteration
	}
	return nil
}

// RecoverAllEvents recovers all events from the latest block in the database
func (c *EvmClient) RecoverAllEvents(ctx context.Context) error {
	log.Info().Str("chainId", c.EvmConfig.GetId()).Msg("[EvmClient] [RecoverAllEvents] starting recovery of all events")

	// Recover all event types
	events := []struct {
		name    string
		recover func(context.Context) error
	}{
		{EVENT_EVM_CONTRACT_CALL, func(ctx context.Context) error {
			return RecoverEvent[*contracts.IScalarGatewayContractCall](c, ctx,
				EVENT_EVM_CONTRACT_CALL, func(log types.Log) *contracts.IScalarGatewayContractCall {
					return &contracts.IScalarGatewayContractCall{
						Raw: log,
					}
				})
		}},
		{EVENT_EVM_CONTRACT_CALL_WITH_TOKEN, func(ctx context.Context) error {
			return RecoverEvent[*contracts.IScalarGatewayContractCallWithToken](c, ctx,
				EVENT_EVM_CONTRACT_CALL_WITH_TOKEN, func(log types.Log) *contracts.IScalarGatewayContractCallWithToken {
					return &contracts.IScalarGatewayContractCallWithToken{
						Raw: log,
					}
				})
		}},
		{EVENT_EVM_TOKEN_SENT, func(ctx context.Context) error {
			return RecoverEvent[*contracts.IScalarGatewayTokenSent](c, ctx,
				EVENT_EVM_TOKEN_SENT, func(log types.Log) *contracts.IScalarGatewayTokenSent {
					return &contracts.IScalarGatewayTokenSent{
						Raw: log,
					}
				})
		}},
		{EVENT_EVM_CONTRACT_CALL_APPROVED, func(ctx context.Context) error {
			return RecoverEvent[*contracts.IScalarGatewayContractCallApproved](c, ctx,
				EVENT_EVM_CONTRACT_CALL_APPROVED, func(log types.Log) *contracts.IScalarGatewayContractCallApproved {
					return &contracts.IScalarGatewayContractCallApproved{
						Raw: log,
					}
				})
		}},
		{EVENT_EVM_COMMAND_EXECUTED, func(ctx context.Context) error {
			return RecoverEvent[*contracts.IScalarGatewayExecuted](c, ctx,
				EVENT_EVM_COMMAND_EXECUTED, func(log types.Log) *contracts.IScalarGatewayExecuted {
					return &contracts.IScalarGatewayExecuted{
						Raw: log,
					}
				})
		}},
		{EVENT_EVM_TOKEN_DEPLOYED, func(ctx context.Context) error {
			return RecoverEvent[*contracts.IScalarGatewayTokenDeployed](c, ctx,
				EVENT_EVM_TOKEN_DEPLOYED, func(log types.Log) *contracts.IScalarGatewayTokenDeployed {
					return &contracts.IScalarGatewayTokenDeployed{
						Raw: log,
					}
				})
		}},
		{EVENT_EVM_SWITCHED_PHASE, func(ctx context.Context) error {
			return RecoverEvent[*contracts.IScalarGatewaySwitchPhase](c, ctx,
				EVENT_EVM_SWITCHED_PHASE, func(log types.Log) *contracts.IScalarGatewaySwitchPhase {
					return &contracts.IScalarGatewaySwitchPhase{
						Raw: log,
					}
				})
		}},
		{EVENT_EVM_REDEEM_TOKEN, func(ctx context.Context) error {
			return RecoverEvent[*contracts.IScalarGatewayRedeemToken](c, ctx,
				EVENT_EVM_REDEEM_TOKEN, func(log types.Log) *contracts.IScalarGatewayRedeemToken {
					return &contracts.IScalarGatewayRedeemToken{
						Raw: log,
					}
				})
		}},
	}

	for _, event := range events {
		if err := event.recover(ctx); err != nil {
			log.Error().Err(err).Str("eventName", event.name).Msg("[EvmClient] [RecoverAllEvents] failed to recover event")
		}
	}

	log.Info().Str("chainId", c.EvmConfig.GetId()).Msg("[EvmClient] [RecoverAllEvents] completed recovery of all events")
	return nil
}

// Get missing events from the last checkpoint block number to the current block number
// In query we filter out the event with index equal to the last checkpoint log index
func GetMissingEvents[T ValidEvmEvent](c *EvmClient, eventName string, lastCheckpoint *scalarnet.EventCheckPoint, fnCreateEventData func(types.Log) T) ([]*parser.EvmEvent[T], error) {
	event, ok := scalarGatewayAbi.Events[eventName]
	if !ok {
		return nil, fmt.Errorf("event %s not found", eventName)
	}

	// Set up a query for logs
	// Todo: config default last checkpoint block number to make sure we don't recover too old events
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(lastCheckpoint.BlockNumber)),
		Addresses: []common.Address{c.GatewayAddress},
		Topics:    [][]common.Hash{{event.ID}}, // Filter by event signature
	}
	// // Fetch the logs
	logs, err := c.Client.FilterLogs(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch logs: %w", err)
	}
	log.Debug().Int("logsCount", len(logs)).Any("logs", logs).Msg("[EvmClient] [GetMissingEvents] fetched logs")
	result := []*parser.EvmEvent[T]{}
	// Parse the logs
	for _, receiptLog := range logs {
		if receiptLog.BlockNumber < lastCheckpoint.BlockNumber ||
			(receiptLog.BlockNumber == lastCheckpoint.BlockNumber && receiptLog.Index <= lastCheckpoint.LogIndex) {
			log.Info().Uint64("receiptLogBlockNumber", receiptLog.BlockNumber).
				Uint("receiptLogIndex", receiptLog.Index).
				Msg("[EvmClient] [GetMissingEvents] skip log")
			continue
		}
		var eventData = fnCreateEventData(receiptLog)
		err := parser.ParseEventData(&receiptLog, &event, eventData)
		if err == nil {
			evmEvent := parser.CreateEvmEventFromArgs[T](eventData, &receiptLog)
			result = append(result, evmEvent)
			log.Info().
				Str("eventName", eventName).
				Uint64("blockNumber", evmEvent.BlockNumber).
				Str("txHash", evmEvent.Hash).
				Uint("txIndex", evmEvent.TxIndex).
				Uint("logIndex", evmEvent.LogIndex).
				Msg("[EvmClient] [GetMissingEvents] parsing successfully.")
		} else {
			log.Error().Err(err).Msg("[EvmClient] [GetMissingEvents] failed to unpack log data")
		}
	}
	log.Info().Int("eventCount", len(logs)).
		Str("eventName", eventName).
		Any("query", query).
		Any("lastCheckpoint", lastCheckpoint).
		Any("missingEventsCount", len(result)).
		Msg("[EvmClient] [GetMissingEvents]")
	return result, nil
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
		{EVENT_EVM_TOKEN_SENT, func(ctx context.Context) error {
			return WatchForEvent[*contracts.IScalarGatewayTokenSent](c, ctx, EVENT_EVM_TOKEN_SENT)
		}},
		{EVENT_EVM_CONTRACT_CALL_WITH_TOKEN, func(ctx context.Context) error {
			return WatchForEvent[*contracts.IScalarGatewayContractCallWithToken](c, ctx, EVENT_EVM_CONTRACT_CALL_WITH_TOKEN)
		}},
		{EVENT_EVM_CONTRACT_CALL_APPROVED, func(ctx context.Context) error {
			return WatchForEvent[*contracts.IScalarGatewayContractCallApproved](c, ctx, EVENT_EVM_CONTRACT_CALL_APPROVED)
		}},
		{EVENT_EVM_COMMAND_EXECUTED, func(ctx context.Context) error {
			return WatchForEvent[*contracts.IScalarGatewayExecuted](c, ctx, EVENT_EVM_COMMAND_EXECUTED)
		}},
		{EVENT_EVM_TOKEN_DEPLOYED, func(ctx context.Context) error {
			return WatchForEvent[*contracts.IScalarGatewayTokenDeployed](c, ctx, EVENT_EVM_TOKEN_DEPLOYED)
		}},
		{EVENT_EVM_SWITCHED_PHASE, func(ctx context.Context) error {
			return WatchForEvent[*contracts.IScalarGatewaySwitchPhase](c, ctx, EVENT_EVM_SWITCHED_PHASE)
		}},
		{EVENT_EVM_REDEEM_TOKEN, func(ctx context.Context) error {
			return WatchForEvent[*contracts.IScalarGatewayRedeemToken](c, ctx, EVENT_EVM_REDEEM_TOKEN)
		}},
	}

	for _, event := range events {
		go func(e struct {
			name  string
			watch func(context.Context) error
		}) {
			if err := e.watch(ctx); err != nil {
				log.Error().Err(err).Any("Config", c.EvmConfig).Msgf("[EvmClient] [ListenToEvents] failed to watch for event: %s", e.name)
			}
		}(event)
	}

	<-ctx.Done()
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

	lastCheckpoint, err := c.dbAdapter.GetLastEventCheckPoint(c.EvmConfig.GetId(), eventName, c.EvmConfig.StartBlock)
	if err != nil {
		log.Warn().Str("chainId", c.EvmConfig.GetId()).
			Str("eventName", eventName).
			Msg("[EvmClient] [getLastCheckpoint] using default value")
	}

	watchOpts := bind.WatchOpts{Start: &lastCheckpoint.BlockNumber, Context: ctx}

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
	case EVENT_EVM_TOKEN_SENT:
		return c.Gateway.WatchTokenSent(watchOpts, sinkInterface.(chan *contracts.IScalarGatewayTokenSent), nil)
	case EVENT_EVM_CONTRACT_CALL_WITH_TOKEN:
		return c.Gateway.WatchContractCallWithToken(watchOpts, sinkInterface.(chan *contracts.IScalarGatewayContractCallWithToken), nil, nil)
	case EVENT_EVM_CONTRACT_CALL_APPROVED:
		return c.Gateway.WatchContractCallApproved(watchOpts, sinkInterface.(chan *contracts.IScalarGatewayContractCallApproved), nil, nil, nil)
	case EVENT_EVM_COMMAND_EXECUTED:
		return c.Gateway.WatchExecuted(watchOpts, sinkInterface.(chan *contracts.IScalarGatewayExecuted), nil)
	case EVENT_EVM_SWITCHED_PHASE:
		return c.Gateway.WatchSwitchPhase(watchOpts, sinkInterface.(chan *contracts.IScalarGatewaySwitchPhase), nil, nil)
	case EVENT_EVM_TOKEN_DEPLOYED:
		return c.Gateway.WatchTokenDeployed(watchOpts, sinkInterface.(chan *contracts.IScalarGatewayTokenDeployed))
	case EVENT_EVM_REDEEM_TOKEN:
		return c.Gateway.WatchRedeemToken(watchOpts, sinkInterface.(chan *contracts.IScalarGatewayRedeemToken), nil, nil, nil)
	default:
		return nil, fmt.Errorf("[EvmClient] [setupSubscription] unsupported event type for %s, event: %v", eventName, (*T)(nil))
	}
}

func handleEvent(c *EvmClient, eventName string, event any) error {
	switch eventName {
	case EVENT_EVM_TOKEN_SENT:
		if evt, ok := event.(*contracts.IScalarGatewayTokenSent); ok {
			return c.HandleTokenSent(evt)
		}
		return fmt.Errorf("cannot parse event %s: %T to %T", eventName, event, (*contracts.IScalarGatewayTokenSent)(nil))
	case EVENT_EVM_CONTRACT_CALL_WITH_TOKEN:
		if evt, ok := event.(*contracts.IScalarGatewayContractCallWithToken); ok {
			return c.HandleContractCallWithToken(evt)
		}
		return fmt.Errorf("cannot parse event %s: %T to %T", eventName, event, (*contracts.IScalarGatewayContractCallWithToken)(nil))
	case EVENT_EVM_CONTRACT_CALL_APPROVED:
		if evt, ok := event.(*contracts.IScalarGatewayContractCallApproved); ok {
			return c.HandleContractCallApproved(evt)
		}
		return fmt.Errorf("cannot parse event %s: %T to %T", eventName, event, (*contracts.IScalarGatewayContractCallApproved)(nil))
	case EVENT_EVM_COMMAND_EXECUTED:
		if evt, ok := event.(*contracts.IScalarGatewayExecuted); ok {
			return c.HandleCommandExecuted(evt)
		}
		return fmt.Errorf("cannot parse event %s: %T to %T", eventName, event, (*contracts.IScalarGatewayExecuted)(nil))
	case EVENT_EVM_TOKEN_DEPLOYED:
		if evt, ok := event.(*contracts.IScalarGatewayTokenDeployed); ok {
			return c.HandleTokenDeployed(evt)
		}
		return fmt.Errorf("cannot parse event %s: %T to %T", eventName, event, (*contracts.IScalarGatewayTokenDeployed)(nil))
	case EVENT_EVM_SWITCHED_PHASE:
		if evt, ok := event.(*contracts.IScalarGatewaySwitchPhase); ok {
			return c.HandleSwitchPhase(evt)
		}
		return fmt.Errorf("cannot parse event %s: %T to %T", eventName, event, (*contracts.IScalarGatewaySwitchPhase)(nil))
	case EVENT_EVM_REDEEM_TOKEN:
		if evt, ok := event.(*contracts.IScalarGatewayRedeemToken); ok {
			return c.HandleRedeemToken(evt)
		}
		return fmt.Errorf("cannot parse event %s: %T to %T", eventName, event, (*contracts.IScalarGatewayRedeemToken)(nil))
	}
	return fmt.Errorf("invalid event type for %s: %T", eventName, event)
}

func (c *EvmClient) handleEventLog(event *abi.Event, txLog types.Log) error {
	switch event.Name {
	case EVENT_EVM_TOKEN_SENT:
		tokenSent := &contracts.IScalarGatewayTokenSent{
			Raw: txLog,
		}
		err := parser.ParseEventData(&txLog, event, tokenSent)
		if err != nil {
			return fmt.Errorf("failed to parse event %s: %w", event.Name, err)
		}
		return c.HandleTokenSent(tokenSent)
	case EVENT_EVM_CONTRACT_CALL_WITH_TOKEN:
		contractCallWithToken := &contracts.IScalarGatewayContractCallWithToken{
			Raw: txLog,
		}
		err := parser.ParseEventData(&txLog, event, contractCallWithToken)
		if err != nil {
			return fmt.Errorf("failed to parse event %s: %w", event.Name, err)
		}
		return c.HandleContractCallWithToken(contractCallWithToken)
	case EVENT_EVM_REDEEM_TOKEN:
		redeemToken := &contracts.IScalarGatewayRedeemToken{
			Raw: txLog,
		}
		err := parser.ParseEventData(&txLog, event, redeemToken)
		if err != nil {
			return fmt.Errorf("failed to parse event %s: %w", event.Name, err)
		}
		return c.HandleRedeemToken(redeemToken)
	case EVENT_EVM_CONTRACT_CALL:
		//return c.HandleContractCall(txLog)
		return fmt.Errorf("not implemented")
	case EVENT_EVM_CONTRACT_CALL_APPROVED:
		contractCallApproved := &contracts.IScalarGatewayContractCallApproved{
			Raw: txLog,
		}
		err := parser.ParseEventData(&txLog, event, contractCallApproved)
		if err != nil {
			return fmt.Errorf("failed to parse event %s: %w", event.Name, err)
		}
		return c.HandleContractCallApproved(contractCallApproved)
	case EVENT_EVM_COMMAND_EXECUTED:
		executed := &contracts.IScalarGatewayExecuted{
			Raw: txLog,
		}
		err := parser.ParseEventData(&txLog, event, executed)
		if err != nil {
			return fmt.Errorf("failed to parse event %s: %w", event.Name, err)
		}
		return c.HandleCommandExecuted(executed)
	case EVENT_EVM_TOKEN_DEPLOYED:
		tokenDeployed := &contracts.IScalarGatewayTokenDeployed{
			Raw: txLog,
		}
		err := parser.ParseEventData(&txLog, event, tokenDeployed)
		if err != nil {
			return fmt.Errorf("failed to parse event %s: %w", event.Name, err)
		}
		return c.HandleTokenDeployed(tokenDeployed)
	case EVENT_EVM_SWITCHED_PHASE:
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
	go func() {
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
	}()
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
