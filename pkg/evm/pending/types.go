package pending

import (
	"sync"
	"time"
)

type PendingTx struct {
	TxHash    string
	Timestamp time.Time
	Retries   int
}

// Keep all transactions sent to Gateway for approval, waiting for event from EVM chain.
type PendingTxs struct {
	mu    sync.Mutex
	items []PendingTx //txHash -> PendingTx
}
