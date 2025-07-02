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
	"github.com/scalarorg/evms-indexer/pkg/db"
)

type BlockWithHeight struct {
	Block  *wire.MsgBlock
	Height int32
}
type BockHashWithHeight struct {
	blockHash *chainhash.Hash
	height    int64
}
type BlockHeaderWithHeight struct {
	BlockHeader *wire.BlockHeader
	Height      int64
}

// TODO: handle reorg
func (c *BtcClient) StartBtcIndexer(ctx context.Context) error {
	blockChan := make(chan *btcjson.GetBlockVerboseTxResult, 1024)
	blockHashesChan := make(chan BockHashWithHeight, 1024)
	blockHeightChan := make(chan int64, 1024)
	log.Info().Msg("Starting BTC indexer")
	// Goroutine 1: Periodically fetch new BTC blocks and send to channel
	go func() {
		defer close(blockChan)
		err := c.fetchBlockData(ctx, blockHashesChan, blockChan)
		if err != nil {
			log.Warn().Err(err).Msg("Failed to fetch block hashes")
		}
	}()
	//We saparate fetching block hashes from fetching block data for future usel
	go c.fetchBlockHashes(ctx, blockHeightChan, blockHashesChan)
	// Goroutine 2: Receive block data, parse VaultTransactions, write to DB
	go func() {
		err := c.indexBlock(ctx, blockChan)
		if err != nil {
			log.Warn().Err(err).Msg("Failed to index blocks")
		}
	}()
	err := c.orchestrateFetching(ctx, blockHeightChan)
	if err != nil {
		log.Warn().Err(err).Msg("Failed to fetch data")
	}
	return err
}

func (c *BtcClient) orchestrateFetching(ctx context.Context,
	blockHeightChan chan<- int64) error {
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
			if err != nil {
				log.Warn().Err(err).Msg("Failed to get BTC block header verbose")
				continue
			}
			c.SetLastBlockHeight(uint64(blockHeaderVerbose.Height))
			endHeight := int64(blockHeaderVerbose.Height)
			if endHeight > c.config.EndHeight && c.config.EndHeight > 0 {
				endHeight = c.config.EndHeight
			}
			log.Info().Int64("startHeight", startHeight).Int64("endHeight", endHeight).Msg("Fetching BTC blocks")

			if startHeight > endHeight {
				continue
			}
			for height := startHeight; height <= endHeight; height++ {
				blockHeightChan <- height
			}
			// Step 1: Define chunks from lastHeight to latestHeight
			// chunks := c.defineChunks(startHeight, endHeight)
			// log.Debug().Int("numChunks", len(chunks)).Msg("Defined chunks for processing")

			// // Step 2: Process each chunk
			// for _, chunk := range chunks {
			// 	chunkChan <- chunk
			// }

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
		batchSize = 64 // default batch size
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

func (c *BtcClient) fetchBlockData(ctx context.Context, blockHashesChan <-chan BockHashWithHeight,
	blockChan chan<- *btcjson.GetBlockVerboseTxResult) error {

	// Create a worker pool for fetching full blocks
	numWorkers := 10 // Default number of workers for fetching full blocks
	if c.config.FetchThread > 0 {
		numWorkers = c.config.FetchThread
	}
	var wg sync.WaitGroup
	// Start workers
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case blockRequest, ok := <-blockHashesChan:
					if !ok {
						// Queue is empty, wait a bit before checking again
						time.Sleep(10 * time.Millisecond)
						continue
					}
					// Fetch full block
					start := time.Now()
					block, err := c.rpcClient.GetBlockVerboseTx(blockRequest.blockHash)
					if err != nil {
						log.Warn().Err(err).Int64("height", blockRequest.height).Int("workerID", workerID).Msg("Failed to get BTC block")
						continue
					} else {
						log.Info().Int("workerID", workerID).
							Msgf("[BtcClient] Fetched block: height %d, hash %s in %fs",
								blockRequest.height, blockRequest.blockHash.String(), time.Since(start).Seconds())
						blockChan <- block
					}
				}
			}
		}(i)
	}
	wg.Wait()
	return nil
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
		case block := <-blockChan:
			start := time.Now()
			blockHeader, err := extractBlockHeader(block)
			if err != nil {
				log.Warn().Err(err).Msg("Failed to extract block header")
				continue
			}
			// Check for reorg before processing the block
			// Create BlockHeaderLite for reorg detection
			currHeight := c.GetLastBlockHeight()
			if block.Height > int64(currHeight)-FINAL_CONFIRMATIONS {
				log.Warn().Int64("BlockHeight", block.Height).
					Uint64("currHeight", currHeight).
					Msg("Checking for reorg")
				blockHash := blockHeader.BlockHash()
				newBlockHeader := &db.BlockHeaderLite{
					Height:   block.Height,
					Hash:     &blockHash,
					PrevHash: &blockHeader.PrevBlock,
				}

				// Check for reorg and handle if necessary
				if err := c.reorgHandler.DetectAndHandleReorg(ctx, newBlockHeader); err != nil {
					log.Warn().Err(err).Int64("height", block.Height).Msg("Error handling reorg. Continuing...")
				}
			}

			// Store block header
			c.StoreBlockHeader(ctx, blockHeader, block.Height)
			start = time.Now()
			// Process transactions
			vaultTxs := []*chains.VaultTransaction{}
			// In each vault transaction, we need to get previous txout to get the staker script pubkey
			// This request is expensive, so we need to cache the previous txout -> TODO: cache previous txout
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
				err := c.StoreVaultTransactions(vaultTxs)
				if err != nil {
					log.Warn().Err(err).Msg("Failed to save vault transaction")
				}
				log.Info().Int64("height", block.Height).Int("vaultTxs", len(vaultTxs)).
					Msgf("Saved vault transactions from block in %fs", time.Since(start).Seconds())
			}
		}
	}
}

// fetchBlockHashes processes a single chunk of block heights using a worker pool
func (c *BtcClient) fetchBlockHashes(ctx context.Context, blockHeightChan <-chan int64, blockHashesChan chan<- BockHashWithHeight) error {
	// Determine number of workers (use config or default)
	numWorkers := 10 // Default number of workers
	if c.config.FetchThread > 0 {
		numWorkers = c.config.FetchThread
	} else if c.config.BatchSize > 0 && c.config.BatchSize < 50 {
		numWorkers = c.config.BatchSize
	}
	log.Info().Int("numWorkers", numWorkers).Msg("[BtcClient] Fetching block hashes")
	for i := 0; i < numWorkers; i++ {
		go func(workerID int) {
			for {
				select {
				case <-ctx.Done():
					return
				case height := <-blockHeightChan:
					// Fetch block hash
					blockHash, err := c.rpcClient.GetBlockHash(height)
					if err != nil {
						log.Warn().Err(err).Int64("height", height).Int("workerID", workerID).Msg("[BtcClient] Failed to get BTC block hash")
						continue
					} else {
						log.Info().Int64("height", height).Int("workerID", workerID).Msg("[BtcClient] Fetched block hash")
						blockHashesChan <- BockHashWithHeight{blockHash: blockHash, height: height}
					}
				}
			}
		}(i)
	}
	return nil
}
