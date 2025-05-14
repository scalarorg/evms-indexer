package evm_test

import (
	"context"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	contracts "github.com/scalarorg/evms-indexer/pkg/evm/contracts/generated"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestContractCallsListener(t *testing.T) {
	TEST_RPC_ENDPOINT := "ws://localhost:8546/"
	TEST_CONTRACT_ADDRESS := "0x2bb588d7bb6faAA93f656C3C78fFc1bEAfd1813D"

	// Setup
	ctx, cancel := context.WithTimeout(context.Background(), 29*time.Minute)
	defer cancel()

	// Connect to a test network
	client, err := ethclient.Dial(TEST_RPC_ENDPOINT)
	require.NoError(t, err)
	defer client.Close()

	// Initialize the contract
	contractAddress := common.HexToAddress(TEST_CONTRACT_ADDRESS)
	gateway, err := contracts.NewIScalarGateway(contractAddress, client)
	require.NoError(t, err)

	// Create channels for events and errors
	eventCh := make(chan *contracts.IScalarGatewayContractCall)
	errorCh := make(chan error)

	// Start listening for ContractCall events
	sub, err := gateway.WatchContractCall(
		&bind.WatchOpts{Context: ctx},
		eventCh,
		nil, // Empty slice for all addresses
		nil, // Empty slice for all command IDs
	)
	t.Log("Subscribed to ContractCall events")
	require.NoError(t, err)
	defer sub.Unsubscribe()

	// Replace timeout-based approach with a 5-minute duration
	timeout := time.After(30 * time.Minute)
	eventCount := 0

	for {
		select {
		case err := <-errorCh:
			t.Fatalf("Error during test: %v", err)
		case event := <-eventCh:
			eventCount++
			t.Logf("Received event %d: DestinationChain=%s, DestinationContractAddress=%s",
				eventCount, event.DestinationChain, event.DestinationContractAddress)
			// Add more specific assertions about the event
			assert.NotEmpty(t, event.DestinationChain)
			assert.NotEmpty(t, event.DestinationContractAddress)
		case <-timeout:
			t.Logf("Test completed after 5 minutes. Total events received: %d", eventCount)
			return
		case <-ctx.Done():
			t.Log("Context cancelled")
			return
		}
	}
}
