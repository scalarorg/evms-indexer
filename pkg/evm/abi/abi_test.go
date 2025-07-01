package abi_test

import (
	"fmt"
	"testing"

	"github.com/scalarorg/evms-indexer/pkg/evm/abi"
	"github.com/stretchr/testify/require"
)

func TestGetEventByName(t *testing.T) {
	redeemEvent, ok := abi.GetEventByName(abi.EVENT_EVM_REDEEM_TOKEN)
	require.True(t, ok)
	require.NotNil(t, redeemEvent)
	fmt.Printf("redeemEvent %v\n", redeemEvent)
}
