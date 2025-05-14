package evm_test

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	contracts "github.com/scalarorg/evms-indexer/pkg/evm/contracts/generated"
	"github.com/stretchr/testify/require"
)

func TestEvmExecutedHandler(t *testing.T) {
	topic := []common.Hash{
		common.HexToHash("0xa74c8847d513feba22a0f0cb38d53081abf97562cdb293926ba243689e7c41ca"),
		common.HexToHash("0x20f2d1ed02585aa8e7f9f49672c93818f2cdae0e72442cd663a4d5dd5ec88213"),
	}
	raw := types.Log{
		Address:     common.HexToAddress("0xc9c5EC5975070a5CF225656e36C53e77eEa318b5"),
		BlockHash:   common.HexToHash("0x23dcc72fa65a311399139274b9f53f6393b319e01b4f98be1316581cd5eb226a"),
		BlockNumber: 7123460,
		Topics:      topic,
		Data:        []byte{},
		Index:       65,
		Removed:     false,
		TxHash:      common.HexToHash("0x4142c36d1e81a6f98de51f685d4588a95ede1c19edd6c44444e66627ccce1e99"),
		TxIndex:     102,
	}
	event := &contracts.IScalarGatewayExecuted{
		CommandId: [32]byte{32, 242, 209, 237, 2, 88, 90, 168, 231, 249, 244, 150, 114, 201, 56, 24, 242, 205, 174, 14, 114, 68, 44, 214, 99, 164, 213, 221, 94, 200, 130, 19},
		Raw:       raw,
	}
	err := sepoliaClient.HandleCommandExecuted(event)
	require.NoError(t, err)
}
