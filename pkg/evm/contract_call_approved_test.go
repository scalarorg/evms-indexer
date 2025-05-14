package evm_test

import (
	"encoding/hex"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	contracts "github.com/scalarorg/evms-indexer/pkg/evm/contracts/generated"
	"github.com/stretchr/testify/require"
)

func TestContractCallApprovedEventHandler(t *testing.T) {
	data, _ := hex.DecodeString("000000000000000000000000000000000000000000000000000000000000008000000000000000000000000000000000000000000000000000000000000000c084381bf297d2687ac4dc55896c5439713539ecc05929809c4d0260e669f65c4800000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000010626974636f696e2d746573746e65743400000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002a30783133304334383130443537313430653145363239363763424637343243614561453931623665634500000000000000000000000000000000000000000000")
	topic := []common.Hash{
		common.HexToHash("0x44e4f8f6bd682c5a3aeba93601ab07cb4d1f21b2aab1ae4880d9577919309aa4"),
		common.HexToHash("0x20f2d1ed02585aa8e7f9f49672c93818f2cdae0e72442cd663a4d5dd5ec88213"),
		common.HexToHash("0x0000000000000000000000001f98c06d8734d5a9ff0b53e3294626e62e4d232c"),
		common.HexToHash("0x6c9cd958fb924c8f75a1edbcc8093e842fb51aee16b449accc62f49796fc5994"),
	}
	raw := types.Log{
		Address:     common.HexToAddress("0xc9c5ec5975070a5cf225656e36c53e77eea318b5"),
		BlockHash:   common.HexToHash("0x23dcc72fa65a311399139274b9f53f6393b319e01b4f98be1316581cd5eb226a"),
		BlockNumber: 0x6cb204,
		Topics:      topic,
		Data:        data,
		Index:       0x40,
		Removed:     false,
		TxHash:      common.HexToHash("0x4142c36d1e81a6f98de51f685d4588a95ede1c19edd6c44444e66627ccce1e99"),
		TxIndex:     0x66,
	}
	event := &contracts.IScalarGatewayContractCallApproved{
		CommandId:     [32]byte{32, 242, 209, 237, 2, 88, 90, 168, 231, 249, 244, 150, 114, 201, 56, 24, 242, 205, 174, 14, 114, 68, 44, 214, 99, 164, 213, 221, 94, 200, 130, 19},
		SourceChain:   "bitcoin-testnet4",
		SourceAddress: "0x130C4810D57140e1E62967cBF742CaEaE91b6ecE",
		//ContractAddress:  common.HexToAddress("0x1f98c06d8734d5a9ff0b53e3294626e62e4d232c"),
		ContractAddress:  common.HexToAddress("0x954690a705742Acc832d9A4244F58Fb1ba323949"),
		PayloadHash:      [32]byte{108, 156, 217, 88, 251, 146, 76, 143, 117, 161, 237, 188, 200, 9, 62, 132, 47, 181, 26, 238, 22, 180, 73, 172, 204, 98, 244, 151, 150, 252, 89, 148},
		SourceTxHash:     [32]byte{132, 56, 27, 242, 151, 210, 104, 122, 196, 220, 85, 137, 108, 84, 57, 113, 53, 57, 236, 192, 89, 41, 128, 156, 77, 2, 96, 230, 105, 246, 92, 72},
		SourceEventIndex: big.NewInt(0),
		Raw:              raw,
	}
	err := sepoliaClient.HandleContractCallApproved(event)
	require.NoError(t, err)
}
