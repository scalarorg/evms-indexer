package evm_test

import (
	"fmt"
	"testing"

	"github.com/scalarorg/evms-indexer/pkg/evm"
	"github.com/stretchr/testify/require"
)

func TestGetEventByName(t *testing.T) {
	redeemEvent, ok := evm.GetEventByName(evm.EVENT_EVM_REDEEM_TOKEN)
	require.True(t, ok)
	require.NotNil(t, redeemEvent)
	fmt.Printf("redeemEvent %v\n", redeemEvent)
}
