package btc

import (
	"context"
	"fmt"

	"github.com/btcsuite/btcd/wire"
	"github.com/scalarorg/data-models/chains"
	"gorm.io/gorm"
)

func (c *BtcClient) GetLatestIndexedHeight(ctx context.Context) (int64, error) {
	if c.dbAdapter == nil {
		return 0, fmt.Errorf("db adapter is nil")
	}
	return c.dbAdapter.GetLatestIndexedHeight(c.config.SourceChain)
}

func (c *BtcClient) StoreBlockHeader(ctx context.Context, header *wire.BlockHeader, height int64) error {
	blockHeader := &chains.BtcBlockHeader{
		Version:       header.Version,
		PrevBlockHash: header.PrevBlock.String(),
		MerkleRoot:    header.MerkleRoot.CloneBytes(),
		Time:          uint32(header.Timestamp.Unix()),
		CompactTarget: header.Bits,
		Nonce:         header.Nonce,
		Hash:          header.BlockHash().String(),
		Height:        int(height),
	}
	return c.dbAdapter.CreateBtcBlockHeader(blockHeader)
}

// CreateVaultTransaction creates a new vault transaction in the database
func (c *BtcClient) StoreVaultTransaction(vaultTx *chains.VaultTransaction) error {
	return c.dbAdapter.PostgresClient.Create(vaultTx).Error
}

// Number of fields in VaultTransaction is 21, so we use a chunk size of 3000
// This is to avoid exceeding PostgreSQL's max number of parameters (65535)
func (c *BtcClient) StoreVaultTransactions(vaultTxs []*chains.VaultTransaction) error {
	if len(vaultTxs) == 0 {
		return nil
	}

	// Chunk size to avoid exceeding PostgreSQL's max number of parameters (65535)
	// Assuming each VaultTransaction has around 10-15 fields, we use a conservative chunk size
	const chunkSize = 1000

	for i := 0; i < len(vaultTxs); i += chunkSize {
		end := i + chunkSize
		if end > len(vaultTxs) {
			end = len(vaultTxs)
		}

		chunk := vaultTxs[i:end]

		// Use Clauses to add conflict resolution
		err := c.dbAdapter.PostgresClient.Clauses(
			gorm.Expr("ON CONFLICT (chain, tx_hash) DO NOTHING"),
		).Create(chunk).Error

		if err != nil {
			return fmt.Errorf("failed to store vault transactions chunk %d-%d: %w", i, end-1, err)
		}
	}

	return nil
}
