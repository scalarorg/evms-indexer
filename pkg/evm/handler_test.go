package evm_test

import (
	"context"
	"encoding/hex"
	"math/big"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/scalarorg/data-models/chains"
	"github.com/scalarorg/evms-indexer/pkg/evm"
	contracts "github.com/scalarorg/evms-indexer/pkg/evm/contracts/generated"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConvertContractCallWithToken2Model(t *testing.T) {
	sepoliaClient, err := evm.NewEvmClient(globalConfig.ConfigPath, sepoliaConfig)
	require.NoError(t, err)
	gatewayAddress := common.HexToAddress("0x320B307AF11918C752F5ddF415679499BC880F74")
	scalarGatewayAbi, err := contracts.IScalarGatewayMetaData.GetAbi()
	require.NoError(t, err)
	topics := []common.Hash{scalarGatewayAbi.Events["ContractCallWithToken"].ID}
	t.Logf("topics %v", topics)
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(8077640)),
		ToBlock:   big.NewInt(int64(8077640)),
		Addresses: []common.Address{gatewayAddress},
		Topics:    [][]common.Hash{topics},
	}
	logs, err := sepoliaClient.Client.FilterLogs(context.Background(), query)
	require.NoError(t, err)
	require.Equal(t, 1, len(logs))
	payload, err := hex.DecodeString("00000000000000000000000000000000000000000000000000000000000000128500000000000000000000000000000000000000000000000000000000000000C00000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000044000000000000000000000000000000000000000000000000000000000000005008F55BD955E1DDAD9E9C0825F1E06F3E396B8D263C47392EF02DE1032E18A2640000000000000000000000000000000000000000000000000000000000000001600147C4917362E25568858BB812751107946A3C769B600000000000000000000000000000000000000000000000000000000000000000000000000000000000500000000000000000000000000000000000000000000000000000000000000A0000000000000000000000000000000000000000000000000000000000000012000000000000000000000000000000000000000000000000000000000000001A0000000000000000000000000000000000000000000000000000000000000022000000000000000000000000000000000000000000000000000000000000002A000000000000000000000000000000000000000000000000000000000000000423078373337376362613333366366663665306433393437383732323436633939666438303935303533663138646162323961613763366330363233643137633035320000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000042307833323465346437323638623062363562613666643731333539623634363061383261396661653137616637343334623438616436333265323033313536623735000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004230783764326233326537666438326436633831393662303161323166616266353765353431666637636234623966353062323738353964386337363639623566616100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000423078653439393237666532313066626237333961626266333539343538356534376361346362623830323165656363363732616535303961336434366562396639340000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000042307831326530626662393962306138643530626138313461383464623234653935333530393331316331343139636166646464383631326566633839393462346336000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000500000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000500000000000000000000000000000000000000000000000000000000000003E800000000000000000000000000000000000000000000000000000000000003E80000000000000000000000000000000000000000000000000000000000000458000000000000000000000000000000000000000000000000000000000000053300000000000000000000000000000000000000000000000000000000000003E8")
	require.NoError(t, err)
	// t.Logf("payload %v", payload)
	contractCallWithTokenEvent := contracts.IScalarGatewayContractCallWithToken{
		Sender:                     common.HexToAddress("0x4FAb6cB4c6E8b72F1529EDA3E71f45127a85D444"),
		DestinationChain:           "bitcoin|4",
		DestinationContractAddress: "0xD3c404704bbD80902831D99f42df12C47B3B9887",
		PayloadHash:                logs[0].Topics[2],
		Payload:                    payload,
		Symbol:                     "sBtc",
		Amount:                     big.NewInt(4741),
		Raw:                        logs[0],
	}
	contractCallWithToken, err := sepoliaClient.ContractCallWithToken2Model(&contractCallWithTokenEvent)
	require.NoError(t, err)
	t.Logf("contractCallWithToken %v", contractCallWithToken)
	assert.Equal(t, "tb1q03y3wd3wy4tgsk9msyn4zyreg63uw6dk0fndvk", contractCallWithToken.DestinationAddress)

}

func TestCalculateOptimalBatchSize(t *testing.T) {
	// Test with TokenSent struct
	tokenSentType := reflect.TypeOf(&chains.TokenSent{})
	batchSize := evm.CalculateOptimalBatchSize(tokenSentType)

	// TokenSent has many fields, so batch size should be smaller
	assert.True(t, batchSize > 0, "Batch size should be positive")
	assert.True(t, batchSize <= evm.BATCH_SIZE, "Batch size should not exceed BATCH_SIZE")

	t.Logf("TokenSent field count: %d, optimal batch size: %d", tokenSentType.Elem().NumField(), batchSize)

	// Test with TokenDeployed struct (fewer fields)
	tokenDeployedType := reflect.TypeOf(&chains.TokenDeployed{})
	batchSizeDeployed := evm.CalculateOptimalBatchSize(tokenDeployedType)

	// TokenDeployed has fewer fields, so batch size should be larger
	assert.True(t, batchSizeDeployed > 0, "Batch size should be positive")
	assert.True(t, batchSizeDeployed <= evm.BATCH_SIZE, "Batch size should not exceed BATCH_SIZE")

	t.Logf("TokenDeployed field count: %d, optimal batch size: %d", tokenDeployedType.Elem().NumField(), batchSizeDeployed)

	// TokenDeployed should have larger batch size than TokenSent due to fewer fields
	assert.True(t, batchSizeDeployed >= batchSize, "TokenDeployed should have larger or equal batch size due to fewer fields")
}

func TestChunkRecords(t *testing.T) {
	// Create test records
	records := make([]any, 2500)
	for i := range records {
		records[i] = &chains.TokenDeployed{}
	}

	// Test chunking with different batch sizes
	testCases := []struct {
		name      string
		batchSize int
		expected  int
	}{
		{"small batch", 100, 25},
		{"medium batch", 500, 5},
		{"large batch", 1000, 3},
		{"very large batch", 3000, 1},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			chunks := evm.ChunkRecords(records, tc.batchSize)
			assert.Equal(t, tc.expected, len(chunks), "Number of chunks should match expected")

			// Verify all records are included
			totalRecords := 0
			for _, chunk := range chunks {
				totalRecords += len(chunk)
			}
			assert.Equal(t, len(records), totalRecords, "All records should be included in chunks")

			// Verify no chunk exceeds batch size (except possibly the last one)
			for i, chunk := range chunks {
				if i < len(chunks)-1 {
					assert.Equal(t, tc.batchSize, len(chunk), "All chunks except the last should have exactly batchSize records")
				} else {
					assert.LessOrEqual(t, len(chunk), tc.batchSize, "Last chunk should not exceed batch size")
				}
			}
		})
	}
}

func TestChunkRecordsEmpty(t *testing.T) {
	// Test with empty records
	chunks := evm.ChunkRecords([]any{}, 100)
	assert.Nil(t, chunks, "Empty records should return nil chunks")

	// Test with nil records
	chunks = evm.ChunkRecords(nil, 100)
	assert.Nil(t, chunks, "Nil records should return nil chunks")
}

func TestChunkRecordsSmallerThanBatch(t *testing.T) {
	// Test with fewer records than batch size
	records := make([]any, 50)
	for i := range records {
		records[i] = &chains.TokenDeployed{}
	}

	chunks := evm.ChunkRecords(records, 100)
	assert.Equal(t, 1, len(chunks), "Should have exactly one chunk")
	assert.Equal(t, 50, len(chunks[0]), "Chunk should contain all records")
}
