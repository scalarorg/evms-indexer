package evm

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/rs/zerolog/log"
	chains "github.com/scalarorg/data-models/chains"
	"github.com/scalarorg/evms-indexer/pkg/types"
)

// RecoverAllEvents recovers all events from the latest block in the database
func (c *EvmClient) FetchRangeLogs(ctx context.Context, query *ethereum.FilterQuery, logsChan chan<- []ethTypes.Log,
	fromBlock uint64, currentBlockNumber uint64, recoverRange *uint64) (int, error) {

	logCounter := 0
	// log.Info().Str("Chain", c.EvmConfig.ID).
	// 	Uint64("FromBlock", fromBlock).
	// 	Uint64("RecoverRange", *recoverRange).
	// 	Uint64("CurrentBlockNumber", currentBlockNumber).
	// 	Msg("[EvmClient] start fetch logs")
	for fromBlock <= currentBlockNumber {
		setQueryRange(query, fromBlock, recoverRange, currentBlockNumber)
		start := time.Now()
		logs, err := c.Client.FilterLogs(context.Background(), *query)
		if err != nil {
			log.Error().Err(err).Msgf("[EvmClient] [FetchRangeLogs] failed to filter logs")
			*recoverRange, err = extractRecoverRange(err.Error())
			if err != nil {
				log.Error().Err(err).Msg("[EvmClient] [FetchRangeLogs] failed to extract recover range from error message")
			} else {
				log.Info().Str("Chain", c.EvmConfig.ID).Uint64("Adjusted RecoverRange", *recoverRange).
					Msgf("[EvmClient] [FetchRangeLogs] recover range extracted from error message: %d", recoverRange)
				setQueryRange(query, fromBlock, recoverRange, currentBlockNumber)
			}
			time.Sleep(time.Second)
		}
		if len(logs) > 0 {
			log.Info().Str("Chain", c.EvmConfig.ID).Msgf("[EvmClient] [RecoverEvents] found %d logs, [%d, %d] /%d, time: %s",
				len(logs), fromBlock, query.ToBlock, currentBlockNumber, time.Since(start))
			logsChan <- logs
			logCounter += len(logs)
			logEventCheckPoint := &types.LogEventCheckPoint{
				ChainID:   c.EvmConfig.GetId(),
				LastBlock: query.ToBlock.Uint64(),
			}
			err = c.dbAdapter.UpdateLatestFetchedBlock(c.EvmConfig.GetId(), logEventCheckPoint)
			if err != nil {
				log.Error().Err(err).Msgf("[EvmClient] [FetchRangeLogs] failed to update latest fetched block")
			}
		}
		//Sleep until 100ms after start to avoid rate limit
		if time.Since(start) > time.Millisecond*100 {
			time.Sleep(time.Millisecond*100 - time.Since(start))
		}
		//Set fromBlock to the next block number for next iteration
		fromBlock = query.ToBlock.Uint64() + 1
	}
	return logCounter, nil
}

func setQueryRange(query *ethereum.FilterQuery, fromBlock uint64, recoverRange *uint64, currentBlockNumber uint64) {
	toBlock := fromBlock + *recoverRange - 1
	if toBlock > currentBlockNumber {
		toBlock = currentBlockNumber
	}
	query.FromBlock = big.NewInt(int64(fromBlock))
	query.ToBlock = big.NewInt(int64(toBlock))
}

func extractRecoverRange(errMsg string) (uint64, error) {
	re := regexp.MustCompile(`You can make eth_getLogs requests with up to a ([\d,","]+) block range`)
	match := re.FindStringSubmatch(errMsg)
	log.Info().Str("errMsg", errMsg).Strs("match", match).Msg("[EvmClient] [extractRecoverRange] match")
	if len(match) > 1 {
		match[1] = strings.ReplaceAll(match[1], `,`, "")
		value, err := strconv.ParseUint(match[1], 10, 64)
		if err != nil {
			return 0, fmt.Errorf("failed to parse number: %w", err)
		}
		return value, nil
	}
	return 0, fmt.Errorf("no block range match")
}

func (c *EvmClient) fetchBlocks(ctx context.Context, blockHeightsChan <-chan map[uint64]uint8) {
	fetchThread := c.EvmConfig.FetchThread
	if fetchThread == 0 {
		fetchThread = 10
	}

	// Create a worker pool for fetching blocks
	blockQueue := types.NewOrderedQueue(func(a, b uint64) int {
		if a < b {
			return -1
		} else if a > b {
			return 1
		}
		return 0
	})
	var wg sync.WaitGroup

	// Start fixed number of worker goroutines
	for i := 0; i < fetchThread; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				default:
					blockNumber, ok := blockQueue.Pop()
					if !ok {
						// Queue is empty, wait a bit before checking again
						time.Sleep(10 * time.Millisecond)
						continue
					}

					// Fetch block
					block, err := c.Client.BlockByNumber(context.Background(), big.NewInt(int64(blockNumber)))
					if err != nil {
						log.Error().Err(err).Msgf("[EvmClient] [worker-%d] failed to fetch block %d", workerID, blockNumber)
						continue
					} else if block == nil {
						log.Error().Msgf("[EvmClient] [worker-%d] block %d not found", workerID, blockNumber)
						continue
					} else {
						log.Info().Uint64("BlockNumber", block.NumberU64()).
							Str("BlockHash", hex.EncodeToString(block.Hash().Bytes())).
							Uint64("BlockTime", block.Time()).
							Int("WorkerID", workerID).
							Msgf("[EvmClient] [worker-%d] found block", workerID)

						blockHeader := &chains.BlockHeader{
							Chain:       c.EvmConfig.GetId(),
							BlockNumber: block.NumberU64(),
							ParentHash:  hex.EncodeToString(block.ParentHash().Bytes()),
							BlockHash:   hex.EncodeToString(block.Hash().Bytes()),
							BlockTime:   block.Time(),
							TxHash:      hex.EncodeToString(block.TxHash().Bytes()), //TransactionRoot, figure out is this merkle tree root or not
							Root:        hex.EncodeToString(block.Root().Bytes()),
							BeaconRoot:  hex.EncodeToString(block.BeaconRoot().Bytes()),
						}
						err = c.dbAdapter.CreateBlockHeader(blockHeader)
						if err != nil {
							log.Error().Err(err).Msgf("[EvmClient] [worker-%d] failed to save block header %d", workerID, blockNumber)
						}
					}
				}
			}
		}(i)
	}

	// Process incoming block heights
	for {
		select {
		case <-ctx.Done():
			wg.Wait()
			return
		case blockHeights, ok := <-blockHeightsChan:
			if !ok {
				wg.Wait()
				return
			}
			for blockNumber := range blockHeights {
				blockQueue.Push(blockNumber)
			}
		}
	}
}
