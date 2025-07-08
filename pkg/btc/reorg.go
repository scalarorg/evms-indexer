package btc

import (
	"context"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/rs/zerolog/log"
	"github.com/scalarorg/evms-indexer/pkg/db"
)

// ReorgHandler provides methods to detect and handle Bitcoin chain reorgs.
type ReorgHandler struct {
	DB DBAdapter // Interface to interact with block storage
}

// DBAdapter abstracts DB operations needed for reorg handling.
type DBAdapter interface {
	GetBlockHashByHeight(ctx context.Context, height int64) (*chainhash.Hash, error)
	GetBlockHeaderByHeight(ctx context.Context, height int64) (*db.BlockHeaderLite, error)
	DeleteBlockAndTxsFromHeight(ctx context.Context, height int64) error
}

// DetectAndHandleReorg checks if the new block connects to the current tip. If not, it rolls back to the fork point.
func (r *ReorgHandler) DetectAndHandleReorg(ctx context.Context, newBlockHeader *db.BlockHeaderLite) error {
	if newBlockHeader.Height == 0 {
		// Genesis block, nothing to check
		return nil
	}

	// Get the last indexed block header
	tipHeader, err := r.DB.GetBlockHeaderByHeight(ctx, newBlockHeader.Height-1)
	if err != nil {
		return err
	}

	if tipHeader.Hash.IsEqual(newBlockHeader.PrevHash) {
		// No reorg, chain continues
		return nil
	}

	// Reorg detected: walk back to find the fork point
	log.Warn().Int64("tipHeader", newBlockHeader.Height-1).
		Str("tipHash", tipHeader.Hash.String()).
		Str("blockPrevHash", newBlockHeader.PrevHash.String()).
		Msg("Reorg detected, rolling back")
	forkHeight, err := r.findForkPoint(ctx, tipHeader, newBlockHeader.PrevHash)
	if err != nil {
		return err
	}
	log.Info().Int64("forkHeight", forkHeight).Msg("Fork point found, rolling back blocks")
	return r.DB.DeleteBlockAndTxsFromHeight(ctx, forkHeight+1)
}

// findForkPoint walks back from the current tip to find the fork point.
func (r *ReorgHandler) findForkPoint(ctx context.Context, tipHeader *db.BlockHeaderLite, targetPrevHash *chainhash.Hash) (int64, error) {
	height := tipHeader.Height
	current := tipHeader
	for height >= 0 {
		if current.Hash.IsEqual(targetPrevHash) {
			return height, nil
		}
		next, err := r.DB.GetBlockHeaderByHeight(ctx, height-1)
		if err != nil {
			return 0, err
		}
		current = next
		height--
	}
	return 0, nil // fallback to genesis
}
