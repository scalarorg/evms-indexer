package pending

import (
	"time"
)

const PENDING_CHECK_INTERVAL = 1 * time.Second

// func (c *PendingTxs) AddTx(txHash string, timestamp time.Time) {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()
// 	c.items = append(c.items, PendingTx{TxHash: txHash, Timestamp: timestamp, Retries: 0})
// }

// func (c *PendingTxs) GetTxs(timeout time.Duration) ([]PendingTx, bool) {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()
// 	txs := []PendingTx{}
// 	for _, item := range c.items {
// 		if time.Since(item.Timestamp) > timeout {
// 			txs = append(txs, item)
// 		}
// 	}
// 	return txs, len(txs) > 0
// }

// // Remove processed pending tx, return true if removed, false if not found
// func (c *PendingTxs) RemoveTx(txHash string) bool {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()
// 	for i, item := range c.items {
// 		if item.TxHash == txHash {
// 			c.items = append(c.items[:i], c.items[i+1:]...)
// 			return true
// 		}
// 	}
// 	return false
// }
