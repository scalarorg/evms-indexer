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
	"github.com/scalarorg/data-models/scalarnet"
	"github.com/scalarorg/evms-indexer/config"
	"github.com/scalarorg/evms-indexer/pkg/evm"
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
		RPCUrl:  "https://eth-sepolia.g.alchemy.com/v2/nNbspp-yjKP9GtAcdKi8xcLnBTptR2Zx",
		//Gateway:    "0x842C080EE1399addb76830CFe21D41e47aaaf57e",
		//Gateway:    "0x78eE3111ab44078FB32D7E7A7bCf99cf3664415B", //Version Mar 27, 2025
		//Gateway:      "0xD2B76Ce7Bf49c8C0965e25B9d76c9cb0c550D7a7", //Version Mar 31, 2025
		Gateway:      "0xCd60852A48fc101304C603A9b1Bbd1E40d35E8c8", //Version Apr 15, 2025
		Finality:     1,
		BlockTime:    time.Second * 12,
		StartBlock:   8095740,
		RecoverRange: 500,
	}
	bnbConfig *evm.EvmNetworkConfig = &evm.EvmNetworkConfig{
		ChainID: 97,
		ID:      CHAIN_ID_BNB,
		Name:    "Ethereum bnb",
		RPCUrl:  "https://bnb-testnet.g.alchemy.com/v2/DpCscOiv_evEPscGYARI3cOVeJ59CRo8",
		//RPCUrl:       "https://data-seed-prebsc-2-s1.binance.org:8545/",
		Gateway:      "0x930C3c4f7d26f18830318115DaD97E0179DA55f0",
		Finality:     1,
		BlockTime:    time.Second * 12,
		StartBlock:   50180903,
		RecoverRange: 1000000,
	}
	bnbClient     *evm.EvmClient
	sepoliaClient *evm.EvmClient
)

func TestMain(m *testing.M) {
	// Load .env file
	err := godotenv.Load("../../../.env.test")
	if err != nil {
		log.Error().Err(err).Msg("Error loading .env.test file: %v")
	}
	evmUserAddress = os.Getenv("EVM_USER_ADDRESS")
	sepoliaEthClient, _ = createEVMClient("RPC_SEPOLIA")
	bnbEthClient, _ = createEVMClient("RPC_BNB")

	log.Info().Msgf("Creating evm client with config: %v", sepoliaConfig)
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
	fnCreateEventData := func(log types.Log) *contracts.IScalarGatewayContractCall {
		return &contracts.IScalarGatewayContractCall{
			Raw: log,
		}
	}
	err := evm.RecoverEvent[*contracts.IScalarGatewayContractCall](sepoliaClient, context.Background(), evm.EVENT_EVM_CONTRACT_CALL, fnCreateEventData)
	require.NoError(t, err)
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
func TestRecoverEventTokenSent(t *testing.T) {
	bnbClient, err := evm.NewEvmClient(configPath, bnbConfig)
	require.NoError(t, err)
	//Get current block number
	blockNumber, err := bnbClient.Client.BlockNumber(context.Background())
	require.NoError(t, err)
	lastCheckpoint := scalarnet.EventCheckPoint{
		ChainName:   bnbConfig.ID,
		EventName:   evm.EVENT_EVM_TOKEN_SENT,
		BlockNumber: blockNumber - 10000,
		TxHash:      "",
		LogIndex:    0,
		EventKey:    "",
	}
	missingEvents, err := evm.GetMissingEvents[*contracts.IScalarGatewayTokenSent](bnbClient, evm.EVENT_EVM_TOKEN_SENT,
		&lastCheckpoint, func(log types.Log) *contracts.IScalarGatewayTokenSent {
			return &contracts.IScalarGatewayTokenSent{
				Raw: log,
			}
		})
	require.NoError(t, err)
	fmt.Printf("missingEvents %v\n", missingEvents)
}
func TestRecoverEventContractCallWithToken(t *testing.T) {
	bnbClient, err := evm.NewEvmClient(configPath, bnbConfig)
	require.NoError(t, err)
	//Get current block number
	blockNumber, err := bnbClient.Client.BlockNumber(context.Background())
	fmt.Printf("blockNumber %v\n", blockNumber)
	require.NoError(t, err)
	lastCheckpoint := scalarnet.EventCheckPoint{
		ChainName:   bnbConfig.ID,
		EventName:   evm.EVENT_EVM_CONTRACT_CALL_WITH_TOKEN,
		BlockNumber: bnbConfig.StartBlock,
		TxHash:      "",
		LogIndex:    0,
		EventKey:    "",
	}
	missingEvents, err := evm.GetMissingEvents[*contracts.IScalarGatewayContractCallWithToken](bnbClient, evm.EVENT_EVM_CONTRACT_CALL_WITH_TOKEN,
		&lastCheckpoint, func(log types.Log) *contracts.IScalarGatewayContractCallWithToken {
			return &contracts.IScalarGatewayContractCallWithToken{
				Raw: log,
			}
		})
	require.NoError(t, err)
	fmt.Printf("%d missing events found\n", len(missingEvents))

	txHash := "0xc7d4fac102169c129a4f04ccd4e3fa17fcd962f137e4928fb2462c52da039899"
	tx, isPending, err := bnbClient.Client.TransactionByHash(context.Background(), common.HexToHash(txHash))
	fmt.Printf("tx %v\n", tx)
	fmt.Printf("isPending %v\n", isPending)
	fmt.Printf("err %v\n", err)
	txHash = "1c9623e21b55e9c4767a12b27f9f68578c167284651efb2b87a51ce438e9fa53"
	tx, isPending, err = bnbClient.Client.TransactionByHash(context.Background(), common.HexToHash(txHash))
	require.NoError(t, err)
	fmt.Printf("tx %v\n", tx)
	fmt.Printf("isPending %v\n", isPending)
	fmt.Printf("err %v\n", err)

	log.Info().Str("txHash", txHash).Any("tx", tx).Msgf("ContractCallWithToken")
	for _, event := range missingEvents {
		receipt, err := bnbClient.Client.TransactionReceipt(context.Background(), common.HexToHash(event.Hash))
		require.NoError(t, err)
		log.Info().Str("txHash", event.Hash).Any("receipt", receipt).Msgf("ContractCallWithToken")
	}
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
