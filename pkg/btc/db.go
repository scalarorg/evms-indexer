package btc

import (
	"context"

	"github.com/btcsuite/btcd/wire"
	"github.com/scalarorg/data-models/chains"
)

func (c *BtcClient) GetLatestIndexedHeight(ctx context.Context) (int64, error) {
	return c.dbAdapter.GetLatestIndexedHeight(c.config.SourceChain)
}

func (c *BtcClient) StoreBlockHeader(ctx context.Context, header *wire.BlockHeader, height int64) error {
	blockHeader := &chains.BtcBlockHeader{
		Version:       header.Version,
		PrevBlockhash: header.PrevBlock.CloneBytes(),
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

func (c *BtcClient) StoreVaultTransactions(vaultTxs []*chains.VaultTransaction) error {
	return c.dbAdapter.PostgresClient.Create(vaultTxs).Error
}
