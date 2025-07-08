package types

import "gorm.io/gorm"

type LogEventCheckPoint struct {
	gorm.Model
	ChainID   string `json:"chain_id" gorm:"uniqueIndex:idx_log_event_checkpoint_chain"`
	LastBlock uint64 `json:"last_block"`
}
