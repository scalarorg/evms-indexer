package btc

import (
	"context"
	"time"

	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/wire"
	"github.com/rs/zerolog/log"
	"github.com/scalarorg/data-models/chains"
)

type BlockWithHeight struct {
	Block  *wire.MsgBlock
	Height int32
}

type BlockHeaderWithHeight struct {
	BlockHeader *wire.BlockHeader
	Height      int64
}

func (c *BtcClient) StartBtcIndexer(ctx context.Context) error {

	blockHeaderChan := make(chan *BlockHeaderWithHeight, 10240)
	blockChan := make(chan *btcjson.GetBlockVerboseResult, 10240)
	log.Info().Msg("Starting BTC indexer")
	// Goroutine 1: Periodically fetch new BTC blocks and send to channel
	go func() {
		defer close(blockChan)
		err := c.fetchData(ctx, blockHeaderChan, blockChan)
		if err != nil {
			log.Warn().Err(err).Msg("Failed to fetch data")
		}
	}()
	// Index block headers
	go func() {
		err := c.indexBlockHeader(ctx, blockHeaderChan)
		if err != nil {
			log.Warn().Err(err).Msg("Failed to index block headers")
		}
	}()
	// Goroutine 2: Receive block data, parse VaultTransactions, write to DB
	go func() {
		err := c.indexBlock(ctx, blockChan)
		if err != nil {
			log.Warn().Err(err).Msg("Failed to index blocks")
		}
	}()

	return nil
}
func (c *BtcClient) fetchData(ctx context.Context,
	blockHeaderChan chan<- *BlockHeaderWithHeight,
	blockChan chan<- *btcjson.GetBlockVerboseResult) error {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	lastHeight, err := c.GetLatestIndexedHeight(ctx)
	if err != nil {
		log.Warn().Err(err).Msg("Failed to get latest indexed height, starting from 0")
		lastHeight = 0
	}

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			// Get latest height from network
			blockHash, err := c.rpcClient.GetBestBlockHash()
			if err != nil {
				log.Warn().Err(err).Msg("Failed to get BTC chain info")
				continue
			}
			blockHeaderVerbose, err := c.rpcClient.GetBlockHeaderVerbose(blockHash)
			latestHeight := int64(blockHeaderVerbose.Height)
			log.Info().Int64("lastHeight", lastHeight).Int64("latestHeight", latestHeight).Msg("Fetching BTC blocks")
			for h := lastHeight + 1; h <= latestHeight; h++ {
				blockHash, err := c.rpcClient.GetBlockHash(int64(h))
				if err != nil {
					log.Warn().Err(err).Int64("height", h).Msg("Failed to get BTC block hash")
					continue
				}
				blockHeader, err := c.rpcClient.GetBlockHeader(blockHash)
				if err != nil {
					log.Warn().Err(err).Int64("height", h).Msg("Failed to get BTC block header")
					continue
				}
				//Send block header to channel
				blockHeaderChan <- &BlockHeaderWithHeight{BlockHeader: blockHeader, Height: h}
				//Get block
				blockVerbose, err := c.rpcClient.GetBlockVerbose(blockHash)
				if err != nil {
					log.Warn().Err(err).Int64("height", h).Msg("Failed to fetch BTC block")
					continue
				}
				//Send block to channel
				blockChan <- blockVerbose
			}
		}
	}
}

func (c *BtcClient) indexBlockHeader(ctx context.Context, blockHeaderChan <-chan *BlockHeaderWithHeight) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case bh, ok := <-blockHeaderChan:
			if !ok {
				return ctx.Err()
			}
			log.Info().Int64("height", bh.Height).Str("hash", bh.BlockHeader.BlockHash().String()).Msg("Storing block header")
			err := c.StoreBlockHeader(ctx, bh.BlockHeader, bh.Height)
			if err != nil {
				log.Warn().Err(err).Msg("Failed to store block header")
			}
		}
	}
}

func (c *BtcClient) indexBlock(ctx context.Context, blockChan <-chan *btcjson.GetBlockVerboseResult) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case block, ok := <-blockChan:
			if !ok {
				return ctx.Err()
			}
			vaultTxs := []*chains.VaultTransaction{}
			for i, tx := range block.RawTx {
				vaultTx, err := c.ParseVaultTxRawResult(&tx, block, i)
				if err != nil {
					log.Warn().Err(err).Int("txIndex", i).Msg("Failed to parse transaction")
					continue
				}
				if vaultTx != nil {
					vaultTxs = append(vaultTxs, vaultTx)

				}
			}
			if len(vaultTxs) == 0 {
				log.Info().Int64("height", block.Height).Msg("No vault transactions found in block")
				continue
			} else {
				log.Info().Int64("height", block.Height).Int("vaultTxs", len(vaultTxs)).Msg("Parsed and saved vault transactions from block")
				err := c.StoreVaultTransactions(vaultTxs)
				if err != nil {
					log.Warn().Err(err).Msg("Failed to save vault transaction")
				}
			}
		}
	}
}

// IndexBlocks indexes blocks by height
// func (c *BtcClient) IndexBlocks(ctx context.Context, height int64) error {
// 	// Get full block from BTC node
// 	block, err := c.GetBlock(ctx, height)
// 	if err != nil {
// 		return fmt.Errorf("failed to get block for height %d: %w", height, err)
// 	}

// 	// Parse each transaction for VaultTransaction
// 	vaultTxs := []*chains.VaultTransaction{}
// 	for i, tx := range block.Transactions {
// 		// Note: block hash is not available from block header here, so we use block.BlockHash().String()
// 		blockHash := block.BlockHash().String()
// 		vaultTx, err := c.ParseVaultTransaction(tx, height, blockHash, i)
// 		if err != nil {
// 			log.Warn().Err(err).Int("txIndex", i).Msg("Failed to parse transaction")
// 			continue
// 		}
// 		if vaultTx != nil {
// 			vaultTxs = append(vaultTxs, vaultTx)
// 			// Save vault transaction to database
// 			err = c.dbAdapter.CreateVaultTransaction(vaultTx)
// 			if err != nil {
// 				log.Warn().Err(err).Msg("Failed to save vault transaction")
// 			}
// 		}
// 	}

// 	log.Info().Int64("height", height).Int("vaultTxs", len(vaultTxs)).Msg("Indexed block (transactions only)")
// 	return nil
// }
