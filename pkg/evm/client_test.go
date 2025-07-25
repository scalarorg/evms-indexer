package evm_test

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	chains "github.com/scalarorg/data-models/chains"
	"github.com/scalarorg/evms-indexer/config"
	"github.com/scalarorg/evms-indexer/pkg/evm"
	evmAbi "github.com/scalarorg/evms-indexer/pkg/evm/abi"
	contracts "github.com/scalarorg/evms-indexer/pkg/evm/contracts/generated"
	"github.com/stretchr/testify/require"
)

const (
	CHAIN_ID_SEPOLIA = "evm|11155111"
	CHAIN_ID_BNB     = "evm|97"
	TOKEN_SYMBOL     = "pBtc"
)

var (
	configPath                 = "../../../example/config"
	globalConfig config.Config = config.Config{
		ConfigPath: configPath,
	}
	sepoliaEthClient *ethclient.Client
	bnbEthClient     *ethclient.Client
	evmUserAddress   string
	sepoliaConfig    *evm.EvmNetworkConfig = &evm.EvmNetworkConfig{
		ChainID: 11155111,
		ID:      CHAIN_ID_SEPOLIA,
		Name:    "Ethereum sepolia",
		RPCUrl:  "https://eth-sepolia.g.alchemy.com/v2/xxxxx-xxxxx-xxxxx",
		//Gateway:    "0x842C080EE1399addb76830CFe21D41e47aaaf57e",
		//Gateway:    "0x78eE3111ab44078FB32D7E7A7bCf99cf3664415B", //Version Mar 27, 2025
		//Gateway:      "0xD2B76Ce7Bf49c8C0965e25B9d76c9cb0c550D7a7", //Version Mar 31, 2025
		//Gateway:      "0xCd60852A48fc101304C603A9b1Bbd1E40d35E8c8", //Version Apr 15, 2025
		Gateway:      "0x6f91bbb2e3b61B466Ba9f3426DcCAcd20e212f40", //Version Apr 15, 2025
		Finality:     1,
		BlockTime:    time.Second * 12,
		StartBlock:   8316916,
		RecoverRange: 100000,
	}
	bnbConfig *evm.EvmNetworkConfig = &evm.EvmNetworkConfig{
		ChainID: 97,
		ID:      CHAIN_ID_BNB,
		Name:    "Ethereum bnb",
		RPCUrl:  "https://bnb-testnet.g.alchemy.com/v2/xxxxx-xxxxx-xxxxx",
		//RPCUrl:       "https://data-seed-prebsc-2-s1.binance.org:8545/",
		Gateway:      "0x6f91bbb2e3b61B466Ba9f3426DcCAcd20e212f40",
		Finality:     1,
		BlockTime:    time.Second * 12,
		StartBlock:   51807731,
		RecoverRange: 1000000,
	}
	bnbClient     *evm.EvmClient
	sepoliaClient *evm.EvmClient
)

func TestMain(m *testing.M) {
	// Load .env file
	err := godotenv.Load("../../.env.test")
	if err != nil {
		log.Error().Err(err).Msg("Error loading .env.test file: %v")
	}
	evmUserAddress = os.Getenv("EVM_USER_ADDRESS")
	sepoliaEthClient, _ = createEVMClient("RPC_SEPOLIA")
	bnbEthClient, _ = createEVMClient("RPC_BNB")
	sepoliaConfig.RPCUrl = os.Getenv("RPC_SEPOLIA")
	log.Info().Msgf("Creating evm client with config: %+v", sepoliaConfig)
	sepoliaClient, err = evm.NewEvmClient(globalConfig.ConfigPath, sepoliaConfig)
	if err != nil {
		log.Error().Msgf("failed to create evm client: %v", err)
	}
	bnbClient, err = evm.NewEvmClient(configPath, bnbConfig)
	if err != nil {
		log.Error().Msgf("failed to create evm client: %v", err)
	}
	os.Exit(m.Run())
}

func createEVMClient(key string) (*ethclient.Client, error) {
	rpcEndpoint := os.Getenv(key)
	if rpcEndpoint == "" {
		return nil, fmt.Errorf("rpc endpoint is empty")
	}
	rpcSepolia, err := rpc.DialContext(context.Background(), rpcEndpoint)
	if err != nil {
		fmt.Printf("failed to connect to sepolia with rpc %s: %v", rpcEndpoint, err)
		return nil, err
	}
	return ethclient.NewClient(rpcSepolia), nil
}
func TestGetBlockHeader(t *testing.T) {
	sepoliaClient, err := evm.NewEvmClient(configPath, sepoliaConfig)
	if err != nil {
		log.Error().Msgf("failed to create evm client: %v", err)
	}
	block, err := sepoliaClient.Client.BlockByNumber(context.Background(), big.NewInt(8279879))
	require.NoError(t, err)
	blockHeader := &chains.BlockHeader{
		Chain:       sepoliaConfig.GetId(),
		BlockNumber: block.NumberU64(),
		BlockHash:   hex.EncodeToString(block.Hash().Bytes()),
		BlockTime:   block.Time(),
	}
	t.Logf("blockHeader %++v", blockHeader)
}

// 2025 July error:
// client_test.go:132: failed to subscribe to events use watch: notifications not supported
// client_test.go:137: failed to watch executed: notifications not supported
func TestSubscribeAllEvents(t *testing.T) {
	logsChan := make(chan types.Log, 1024)
	_, topics, err := evm.PrepareEvents()
	if err != nil {
		t.Fatalf("failed to prepare events: %v", err)
	}
	err = sepoliaClient.SubscribeAllEvents(context.Background(), topics, logsChan)
	if err != nil {
		t.Logf("failed to subscribe to events use watch: %v", err)
		watchOpts := bind.WatchOpts{Start: &sepoliaConfig.StartBlock, Context: context.Background()}
		sink := make(chan *contracts.IScalarGatewayExecuted)
		subExecuted, err := sepoliaClient.Gateway.WatchExecuted(&watchOpts, sink, nil)
		if err != nil {
			t.Logf("failed to watch executed: %v", err)
		}
		if subExecuted != nil {
			t.Logf("subscribed to executed")
			go func() {
				for event := range sink {
					t.Logf("event %v\n", event)
				}
			}()
			go func() {
				errChan := subExecuted.Err()
				if err := <-errChan; err != nil {
					t.Logf("received error: %v", err)
				}
			}()
		}
	}
	// require.NoError(t, err)
	for {
		select {
		case log := <-logsChan:
			t.Logf("log %v\n", log)

		case <-time.After(10 * time.Second):
			t.Logf("Timeout waiting for logs")
		}
	}
}
func TestEvmClientListenContractCallEvent(t *testing.T) {
	watchOpts := bind.WatchOpts{Start: &sepoliaConfig.StartBlock, Context: context.Background()}
	sink := make(chan *contracts.IScalarGatewayContractCall)

	subContractCall, err := sepoliaClient.Gateway.WatchContractCall(&watchOpts, sink, nil, nil)
	if err != nil {
		log.Error().Err(err).Msg("ContractCallEvent")
	}
	if subContractCall != nil {
		log.Info().Msg("Subscribed to ContractCallEvent successfully.")
		go func() {
			log.Info().Msg("Waiting for events...")
			for event := range sink {
				log.Info().Any("event", event).Msgf("ContractCall")
			}
		}()
		go func() {
			errChan := subContractCall.Err()
			if err := <-errChan; err != nil {
				log.Error().Err(err).Msg("Received error")
			}
		}()
	}
	select {}
}

func TestEvmClientListenContractCallApprovedEvent(t *testing.T) {
	watchOpts := bind.WatchOpts{Start: &sepoliaConfig.StartBlock, Context: context.Background()}
	sink := make(chan *contracts.IScalarGatewayContractCallApproved)

	subContractCallApproved, err := sepoliaClient.Gateway.WatchContractCallApproved(&watchOpts, sink, nil, nil, nil)
	if err != nil {
		log.Error().Err(err).Msg("ContractCallApprovedEvent")
	}
	if subContractCallApproved != nil {
		log.Info().Msg("Subscribed to ContractCallApprovedEvent successfully.")
		go func() {
			log.Info().Msg("Waiting for events...")
			for event := range sink {
				log.Info().Any("event", event).Msgf("ContractCallApproved")
			}
		}()
		go func() {
			errChan := subContractCallApproved.Err()
			if err := <-errChan; err != nil {
				log.Error().Err(err).Msg("Received error")
			}
			subContractCallApproved.Unsubscribe()
		}()
	}
	select {}
}
func TestEvmClientListenEVMExecutedEvent(t *testing.T) {
	watchOpts := bind.WatchOpts{Start: &sepoliaConfig.StartBlock, Context: context.Background()}
	sink := make(chan *contracts.IScalarGatewayExecuted)

	subExecuted, err := sepoliaClient.Gateway.WatchExecuted(&watchOpts, sink, nil)
	if err != nil {
		log.Error().Err(err).Msg("ExecutedEvent")
	}
	if subExecuted != nil {
		log.Info().Msg("Subscribed to ExecutedEvent successfully. Waiting for events...")
		go func() {
			for event := range sink {
				log.Info().Any("event", event).Msgf("Executed")
			}
		}()
		go func() {
			errChan := subExecuted.Err()
			if err := <-errChan; err != nil {
				log.Error().Err(err).Msg("Received error")
			}
		}()
	}
	select {}
}
func TestRecoverEvent(t *testing.T) {
	logsChan := make(chan []types.Log)
	gatewayAbi, err := evmAbi.GetScalarGatewayAbi()
	require.NoError(t, err)
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
	for _, eventName := range eventNames {
		event, ok := gatewayAbi.Events[eventName]
		if !ok {
			log.Warn().Str("eventName", eventName).Msg("Event not found in ABI")
			continue
		}
		topics = append(topics, event.ID)
	}

	go sepoliaClient.LoopFetchLogs(context.Background(), topics, logsChan)
	for {
		t.Logf("Waiting for logs")
		select {
		case logs := <-logsChan:
			for _, log := range logs {
				t.Logf("log %v\n", log)
			}
		case <-time.After(10 * time.Second):
			t.Logf("Timeout waiting for logs")
		}
	}
}
func TestEvmSubscribe(t *testing.T) {
	fmt.Println("Test evm client")

	// Connect to Ethereum client
	client, err := ethclient.Dial(sepoliaConfig.RPCUrl)
	require.NoError(t, err)
	if err != nil {
		fmt.Printf("failed to connect to the Ethereum client: %v", err)
	}

	// Get current block
	currentBlock, err := client.BlockNumber(context.Background())
	require.NoError(t, err)
	if err != nil {
		fmt.Printf("failed to get current block: %v", err)
	}
	fmt.Printf("Current block %d\n", currentBlock)

	// Create the event signature
	contractCallSig := []byte("ContractCall(address,string,string,bytes32,bytes)")

	// Create the filter query
	query := ethereum.FilterQuery{
		Addresses: []common.Address{common.HexToAddress(sepoliaConfig.Gateway)},
		Topics: [][]common.Hash{{
			crypto.Keccak256Hash(contractCallSig),
		}},
	}

	// Subscribe to events
	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	require.NoError(t, err)
	if err != nil {
		fmt.Printf("failed to subscribe to logs: %v", err)
	}

	// Handle events in a separate goroutine
	go func() {
		for {
			select {
			case err := <-sub.Err():
				fmt.Printf("Received error: %v", err)
			case vLog := <-logs:
				fmt.Println("Log:", vLog)
			}
		}
	}()

	// Keep the program running
	select {}
}

func TestEvmClientWatchTokenSent(t *testing.T) {
	watchOpts := bind.WatchOpts{Start: &sepoliaConfig.StartBlock, Context: context.Background()}
	sink := make(chan *contracts.IScalarGatewayTokenSent)
	bnbClient, err := evm.NewEvmClient(configPath, bnbConfig)
	if err != nil {
		log.Error().Msgf("failed to create evm client: %v", err)
	}
	subscription, err := bnbClient.Gateway.WatchTokenSent(&watchOpts, sink, nil)
	require.NoError(t, err)
	defer subscription.Unsubscribe()
	log.Info().Msgf("[EvmClient] [watchEVMTokenSent] success. Listening to TokenSent")

	for {
		select {
		case err := <-subscription.Err():
			log.Error().Msgf("[EvmClient] [watchEVMTokenSent] error: %v", err)
		case event := <-sink:
			log.Info().Any("event", event).Msgf("EvmClient] [watchEVMTokenSent]")
		}
	}
}
func TestReconnectWithWatchTokenSent(t *testing.T) {
	watchOpts := bind.WatchOpts{Start: &sepoliaConfig.StartBlock, Context: context.Background()}
	sink := make(chan *contracts.IScalarGatewayTokenSent)
	bnbClient, err := evm.NewEvmClient(configPath, bnbConfig)
	if err != nil {
		log.Error().Msgf("failed to create evm client: %v", err)
	}
	subscription, err := bnbClient.Gateway.WatchTokenSent(&watchOpts, sink, nil)
	require.NoError(t, err)
	defer subscription.Unsubscribe()
	log.Info().Msgf("[EvmClient] [watchEVMTokenSent] success. Listening to TokenSent")
	done := false
	for !done {
		select {
		case err := <-subscription.Err():
			log.Error().Err(err).Msg("[EvmClient] [watchEVMTokenSent] error with subscription, perform reconnect")
			subscription, err = bnbClient.Gateway.WatchTokenSent(&watchOpts, sink, nil)
			require.NoError(t, err)
		case <-watchOpts.Context.Done():
			log.Info().Msgf("[EvmClient] [watchEVMTokenSent] context done")
			done = true
		case event := <-sink:
			log.Info().Any("event", event).Msgf("EvmClient] [watchEVMTokenSent]")
		}
	}
}
func createErc20ProxyContract(proxyAddress string, client *ethclient.Client) (*contracts.IScalarERC20CrossChain, error) {
	proxy, err := contracts.NewIScalarERC20CrossChain(common.HexToAddress(proxyAddress), client)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize proxy contract from address: %s", proxyAddress)
	}
	return proxy, nil
}
