package btc

import (
	"context"
	"strconv"
	"sync"
	"time"

	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
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

// TODO: handle reorg
func (c *BtcClient) StartBtcIndexer(ctx context.Context) error {

	blockChan := make(chan *btcjson.GetBlockVerboseTxResult, 1024)
	blockHashesChan := make(chan map[int64]*chainhash.Hash, 1024)
	log.Info().Msg("Starting BTC indexer")
	// Goroutine 1: Periodically fetch new BTC blocks and send to channel
	go func() {
		defer close(blockHashesChan)
		err := c.orchestrateFetching(ctx, blockHashesChan)
		if err != nil {
			log.Warn().Err(err).Msg("Failed to fetch data")
		}
	}()
	go func() {
		defer close(blockChan)
		err := c.fetchBlockData(ctx, blockHashesChan, blockChan)
		if err != nil {
			log.Warn().Err(err).Msg("Failed to fetch block hashes")
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

func (c *BtcClient) orchestrateFetching(ctx context.Context,
	blockHashesChan chan<- map[int64]*chainhash.Hash) error {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	lastHeight, err := c.GetLatestIndexedHeight(ctx)
	if err != nil {
		log.Warn().Err(err).Msg("Failed to get latest indexed height, starting from 0")
		lastHeight = 0
	}
	startHeight := c.config.StartHeight
	if startHeight < lastHeight+1 {
		startHeight = lastHeight + 1
	}
	log.Info().Msgf("Starting from height %d", startHeight)

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
			endHeight := int64(blockHeaderVerbose.Height)
			if endHeight > c.config.EndHeight && c.config.EndHeight > 0 {
				endHeight = c.config.EndHeight
			}
			log.Info().Int64("startHeight", startHeight).Int64("endHeight", endHeight).Msg("Fetching BTC blocks")

			if startHeight > endHeight {
				continue
			}

			// Step 1: Define chunks from lastHeight to latestHeight
			chunks := c.defineChunks(startHeight, endHeight)
			log.Debug().Int("numChunks", len(chunks)).Msg("Defined chunks for processing")

			// Step 2: Process each chunk
			for i, chunk := range chunks {
				start := time.Now()
				blockHashes, err := c.fetchBlockHashes(ctx, chunk)
				if err != nil {
					log.Warn().Err(err).Int("chunkIndex", i).Msg("Failed to process chunk")
					continue
				}
				log.Debug().Int("chunkIndex", i).Int64("startHeight", chunk.Start).Int64("endHeight", chunk.End).
					Msgf("Fetched block hashes for chunk in %f seconds", time.Since(start).Seconds())
				blockHashesChan <- blockHashes
			}

			// Update lastHeight to latestHeight
			startHeight = endHeight + 1
		}
	}
}

// Chunk represents a range of block heights to process
type Chunk struct {
	Start int64
	End   int64
}

// defineChunks divides the range [startHeight, endHeight] into chunks
func (c *BtcClient) defineChunks(startHeight, endHeight int64) []Chunk {
	batchSize := c.config.BatchSize
	if batchSize <= 0 {
		batchSize = 10 // default batch size
	}

	var chunks []Chunk

	// Create chunks
	for start := startHeight; start <= endHeight; start += int64(batchSize) {
		end := start + int64(batchSize) - 1
		if end > endHeight {
			end = endHeight
		}

		chunk := Chunk{
			Start: start,
			End:   end,
		}
		chunks = append(chunks, chunk)
	}

	return chunks
}

// fetchBlockHashes processes a single chunk of block heights
func (c *BtcClient) fetchBlockHashes(ctx context.Context, chunk Chunk) (map[int64]*chainhash.Hash, error) {
	// Use async group to collect block hashes from futures
	blockHashes := make(map[int64]*chainhash.Hash)
	var wg sync.WaitGroup
	var mu sync.Mutex
	for height := chunk.Start; height <= chunk.End; height++ {
		wg.Add(1)
		go func(h int64) {
			defer wg.Done()
			blockHash, err := c.rpcClient.GetBlockHash(h)
			if err != nil {
				log.Warn().Err(err).Int64("height", h).Msg("Failed to get BTC block hash")
				return
			}

			// Safely add to the map
			mu.Lock()
			blockHashes[h] = blockHash
			mu.Unlock()
		}(height)
	}
	// Wait for all block hash futures to complete
	wg.Wait()

	log.Debug().Int64("startHeight", chunk.Start).Int64("endHeight", chunk.End).Int("fetchedBlockHashes", len(blockHashes)).Msg("Completed processing chunk")
	return blockHashes, nil
}
func (c *BtcClient) fetchBlockData(ctx context.Context, blockHashesChan <-chan map[int64]*chainhash.Hash,
	blockChan chan<- *btcjson.GetBlockVerboseTxResult) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case blockHashes, ok := <-blockHashesChan:
			if !ok {
				return ctx.Err()
			}
			for height, blockHash := range blockHashes {
				go func(h int64, hash *chainhash.Hash) {
					log.Info().Msgf("Fetching block: height %d, hash %s", h, hash.String())
					block, err := c.rpcClient.GetBlockVerboseTx(hash)
					if err != nil {
						log.Warn().Err(err).Int64("height", h).Msg("Failed to get BTC block")
						return
					}
					blockChan <- block
				}(height, blockHash)
			}
		}
	}
}

func extractBlockHeader(block *btcjson.GetBlockVerboseTxResult) (*wire.BlockHeader, error) {
	prevHash, err := chainhash.NewHashFromStr(block.PreviousHash)
	if err != nil {
		return nil, err
	}
	merkleRootHash, err := chainhash.NewHashFromStr(block.MerkleRoot)
	if err != nil {
		return nil, err
	}
	bits, err := strconv.ParseUint(block.Bits, 10, 32)
	return wire.NewBlockHeader(block.Version, prevHash, merkleRootHash, uint32(bits), block.Nonce), nil
}
func (c *BtcClient) indexBlock(ctx context.Context, blockChan <-chan *btcjson.GetBlockVerboseTxResult) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case block, ok := <-blockChan:
			if !ok {
				return ctx.Err()
			}
			blockHeader, err := extractBlockHeader(block)
			if err != nil {
				log.Warn().Err(err).Msg("Failed to extract block header")
				continue
			}
			c.StoreBlockHeader(ctx, blockHeader, block.Height)
			vaultTxs := []*chains.VaultTransaction{}
			for i, tx := range block.Tx {
				vaultTx, err := c.ParseVaultTxRawResult(&tx, block, i)
				if err != nil {
					continue
				}
				if vaultTx != nil {
					vaultTxs = append(vaultTxs, vaultTx)

				}
			}
			if len(vaultTxs) > 0 {
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
