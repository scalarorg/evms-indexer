package utils

import (
	"strings"

	"github.com/scalarorg/evms-indexer/pkg/types"
)

func NormalizeHash(hash string) string {
	return strings.ToLower(strings.TrimPrefix(hash, "0x"))
}

func NormalizeAddress(address string, chainType types.ChainType) string {
	switch chainType {
	case types.ChainTypeEVM:
		address = strings.ToLower(address)
		if !strings.HasPrefix(address, "0x") {
			address = "0x" + address
		}
		return address
	default:
		return address
	}
}

// func ConvertUint64ToChainInfo(n uint64) (*chain.ChainInfo, error) {
// 	buf := make([]byte, 8)
// 	binary.BigEndian.PutUint64(buf, n)
// 	chainInfo := chain.NewChainInfoFromBytes(buf)
// 	if chainInfo == nil {
// 		return nil, fmt.Errorf("invalid destination chain: %d", n)
// 	}
// 	return chainInfo, nil
// }
