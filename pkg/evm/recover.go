package evm

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rs/zerolog/log"
	chains "github.com/scalarorg/data-models/chains"
)

// RecoverAllEvents recovers all events from the latest block in the database
func (c *EvmClient) RecoverAllEvents(ctx context.Context, topics []common.Hash, logsChan chan<- []types.Log) error {
	currentBlockNumber, err := c.Client.BlockNumber(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get current block number: %w", err)
	}
	recoverRange := uint64(1000000)
	if c.EvmConfig.RecoverRange > 0 && c.EvmConfig.RecoverRange < 1000000 {
		recoverRange = c.EvmConfig.RecoverRange
	}
	fromBlock := uint64(0)
	if c.dbAdapter != nil {
		fromBlock, err = c.dbAdapter.GetLatestBlockFromAllEvents(c.EvmConfig.GetId())
		if err != nil {
			return fmt.Errorf("failed to get latest block number: %w", err)
		}
	}
	if fromBlock < c.EvmConfig.StartBlock {
		fromBlock = c.EvmConfig.StartBlock
	}
	// Set up a query for logs
	query := ethereum.FilterQuery{
		Addresses: []common.Address{c.GatewayAddress},
		Topics:    [][]common.Hash{topics}, // Filter by multiple event signatures
	}
	logCounter := 0
	log.Info().Str("Chain", c.EvmConfig.ID).
		Uint64("FromBlock", fromBlock).
		Uint64("RecoverRange", recoverRange).
		Uint64("CurrentBlockNumber", currentBlockNumber).
		Msg("[EvmClient] [RecoverAllEvents] starting recover all events")
	for fromBlock < currentBlockNumber {
		setQueryRange(&query, fromBlock, recoverRange, currentBlockNumber)
		start := time.Now()
		logs, err := c.Client.FilterLogs(context.Background(), query)
		if err != nil {
			log.Error().Err(err).Msgf("[EvmClient] [RecoverEvents] failed to filter logs")
			recoverRange, err = extractRecoverRange(err.Error())
			if err != nil {
				log.Error().Err(err).Msgf("[EvmClient] [RecoverEvents] failed to extract recover range from error message: %s", err.Error())
				return err
			}
			log.Info().Str("Chain", c.EvmConfig.ID).Uint64("Adjusted RecoverRange", recoverRange).
				Msgf("[EvmClient] [RecoverEvents] recover range extracted from error message: %d", recoverRange)
			setQueryRange(&query, fromBlock, recoverRange, currentBlockNumber)
			continue
		}
		if len(logs) > 0 {
			log.Info().Str("Chain", c.EvmConfig.ID).Msgf("[EvmClient] [RecoverEvents] found %d logs, [%d, %d]/%d, time: %s",
				len(logs), fromBlock, query.ToBlock, currentBlockNumber, time.Since(start))
			logsChan <- logs
			logCounter += len(logs)
		}
		//Set fromBlock to the next block number for next iteration
		fromBlock = query.ToBlock.Uint64() + 1
	}
	log.Info().
		Str("Chain", c.EvmConfig.ID).
		Uint64("CurrentBlockNumber", currentBlockNumber).
		Int("TotalLogs", logCounter).
		Msg("[EvmClient] [FinishRecover] recovered all events")
	return nil
}

func setQueryRange(query *ethereum.FilterQuery, fromBlock uint64, recoverRange uint64, currentBlockNumber uint64) {
	toBlock := fromBlock + recoverRange - 1
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
	return 0, fmt.Errorf("no match found")
}

func (c *EvmClient) fetchBlocks(ctx context.Context, blockHeightsChan <-chan map[uint64]uint8) {
	for {
		select {
		case <-ctx.Done():
			return
		case blockHeights := <-blockHeightsChan:
			for blockNumber, _ := range blockHeights {
				go func(blockNumber uint64) {
					block, err := c.Client.BlockByNumber(context.Background(), big.NewInt(int64(blockNumber)))
					if err != nil {
						log.Error().Err(err).Msgf("[EvmClient] [startFetchBlock] failed to fetch block %d", blockNumber)
					} else if block == nil {
						log.Error().Msgf("[EvmClient] [startFetchBlock] block %d not found", blockNumber)
					} else {
						log.Info().Uint64("BlockNumber", block.NumberU64()).
							Str("BlockHash", hex.EncodeToString(block.Hash().Bytes())).
							Uint64("BlockTime", block.Time()).
							Msgf("[EvmClient] [startFetchBlock] found block")

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
							log.Error().Err(err).Msgf("[EvmClient] [startFetchBlock] failed to save block header %d", blockNumber)
						}
					}
				}(blockNumber)
			}
		}
	}
}
