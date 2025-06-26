package evm

import (
	"context"
	"fmt"
	"math/big"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rs/zerolog/log"
	"github.com/scalarorg/data-models/scalarnet"
)

var (
	ALL_EVENTS = []string{
		EVENT_EVM_CONTRACT_CALL,
		EVENT_EVM_CONTRACT_CALL_WITH_TOKEN,
		EVENT_EVM_TOKEN_SENT,
		EVENT_EVM_CONTRACT_CALL_APPROVED,
		EVENT_EVM_COMMAND_EXECUTED,
		EVENT_EVM_TOKEN_DEPLOYED,
		EVENT_EVM_SWITCHED_PHASE,
		EVENT_EVM_REDEEM_TOKEN,
	}
)

// Go routine for process missing logs
func (c *EvmClient) ProcessMissingLogs() {
	gatewayAbi, err := getScalarGatewayAbi()
	if err != nil {
		log.Error().Err(err).Msg("[EvmClient] [ProcessMissingLogs] failed to get scalar gateway abi")
		return
	}
	mapEvents := map[string]*abi.Event{}
	for _, event := range gatewayAbi.Events {
		mapEvents[event.ID.String()] = &event
	}
	for {
		logs := c.MissingLogs.GetLogs(10)
		if len(logs) == 0 {
			if c.MissingLogs.IsRecovered() {
				log.Info().Str("Chain", c.EvmConfig.ID).Msg("[EvmClient] [ProcessMissingLogs] no logs to process, recovered flag is true, exit")
				break
			} else {
				log.Info().Str("Chain", c.EvmConfig.ID).Msg("[EvmClient] [ProcessMissingLogs] no logs to process, recover is in progress, sleep 1 second then continue")
				time.Sleep(time.Second)
				continue
			}
		}
		log.Info().Str("Chain", c.EvmConfig.ID).Int("Number of logs", len(logs)).Msg("[EvmClient] [ProcessMissingLogs] processing logs")
		for _, txLog := range logs {
			topic := txLog.Topics[0].String()
			event, ok := mapEvents[topic]
			if !ok {
				log.Error().Str("topic", topic).Any("txLog", txLog).Msg("[EvmClient] [ProcessMissingLogs] event not found")
				continue
			}
			log.Debug().
				Str("chainId", c.EvmConfig.GetId()).
				Str("eventName", event.Name).
				Str("txHash", txLog.TxHash.String()).
				Msg("[EvmClient] [ProcessMissingLogs] start processing missing event")

			err := c.handleEventLog(event, txLog)
			if err != nil {
				log.Error().Err(err).Msg("[EvmClient] [ProcessMissingLogs] failed to handle event log")
			}

		}
	}
	log.Info().Str("Chain", c.EvmConfig.ID).Msg("[EvmClient] [ProcessMissingLogs] finished processing all missing evm events")
}

func (c *EvmClient) RecoverEvents(ctx context.Context, eventNames []string, currentBlockNumber uint64) error {
	if c.dbAdapter == nil {
		log.Warn().Msgf("[EvmClient] [RecoverEvents] dbAdapter is nil, skip recovering events")
		return nil
	}
	topics := []common.Hash{}
	mapEvents := map[string]abi.Event{}
	for _, eventName := range eventNames {
		event, ok := scalarGatewayAbi.Events[eventName]
		if ok {
			log.Info().Str("Chain", c.EvmConfig.ID).
				Str("EventName", event.Name).
				Str("EventID", event.ID.String()).
				Msgf("[EvmClient] [RecoverEvents] adding event to topics")
			topics = append(topics, event.ID)
			mapEvents[event.ID.String()] = event
		}
	}
	lastCheckpoint, err := c.dbAdapter.GetLastCheckPoint(c.EvmConfig.GetId())
	if err != nil {
		log.Warn().Err(err).Msgf("[EvmClient] [RecoverEvents] failed to get last checkpoint use default value")
	}
	if lastCheckpoint.BlockNumber == 0 {
		lastCheckpoint.BlockNumber = c.EvmConfig.StartBlock
	}

	log.Info().Str("Chain", c.EvmConfig.ID).
		Str("GatewayAddress", c.GatewayAddress.String()).
		Str("EventNames", strings.Join(eventNames, ",")).
		Any("LastCheckpoint", lastCheckpoint).Msg("[EvmClient] [RecoverEvents] start recovering events")
	recoverRange := uint64(100000)
	if c.EvmConfig.RecoverRange > 0 && c.EvmConfig.RecoverRange < 100000 {
		recoverRange = c.EvmConfig.RecoverRange
	}
	fromBlock := lastCheckpoint.BlockNumber
	toBlock := fromBlock + recoverRange - 1
	if toBlock > currentBlockNumber {
		toBlock = currentBlockNumber
	}
	query := ethereum.FilterQuery{
		Addresses: []common.Address{c.GatewayAddress},
		Topics:    [][]common.Hash{topics},
	}
	logCounter := 0
	for fromBlock < currentBlockNumber {
		query.FromBlock = big.NewInt(int64(fromBlock))
		query.ToBlock = big.NewInt(int64(toBlock))
		start := time.Now()
		logs, err := c.Client.FilterLogs(context.Background(), query)
		elapsed := time.Since(start)
		if err != nil {
			log.Error().Err(err).
				Str("Chain", c.EvmConfig.ID).
				Uint64("FromBlock", query.FromBlock.Uint64()).
				Uint64("ToBlock", toBlock).
				Str("Endpoint", c.EvmConfig.RPCUrl).
				Msg("[EvmClient] [RecoverEvents] error when getting logs")
			recoverRangeInt, err := extractRecoverRange(err.Error())
			if err != nil {
				log.Error().Err(err).Msgf("[EvmClient] [RecoverEvents] failed to extract recover range from error message: %s", err.Error())
				return err
			}
			recoverRange = uint64(recoverRangeInt)
			log.Info().Str("Chain", c.EvmConfig.ID).Int("AdjustedRecoverRange", recoverRangeInt).
				Msgf("[EvmClient] [RecoverEvents] recover range extracted from error message: %d", recoverRangeInt)
			toBlock = fromBlock + recoverRange - 1
			continue
		} else {
			log.Info().Str("Chain", c.EvmConfig.ID).
				Uint64("FromBlock", query.FromBlock.Uint64()).
				Uint64("ToBlock", toBlock).
				Int("Logs found", len(logs)).
				Str("Elapsed", elapsed.String()).
				Msgf("[EvmClient] [RecoverAllRedeemSessions] progress %d/%d", currentBlockNumber-query.FromBlock.Uint64()+1, currentBlockNumber-c.EvmConfig.StartBlock+1)
		}

		if len(logs) > 0 {
			log.Info().Str("Chain", c.EvmConfig.ID).Msgf("[EvmClient] [RecoverEvents] found %d logs, fromBlock: %d, toBlock: %d", len(logs), fromBlock, query.ToBlock)
			c.MissingLogs.AppendLogs(logs)
			logCounter += len(logs)
			c.UpdateLastCheckPoint(mapEvents, logs, query.ToBlock.Uint64())
		} else {
			log.Info().Str("Chain", c.EvmConfig.ID).Msgf("[EvmClient] [RecoverEvents] no logs found, fromBlock: %d, toBlock: %d", fromBlock, query.ToBlock)
		}
		//Set fromBlock to the next block number for next iteration
		fromBlock = toBlock + 1
		toBlock = fromBlock + recoverRange - 1
		if toBlock > currentBlockNumber {
			toBlock = currentBlockNumber
		}
	}
	log.Info().
		Str("Chain", c.EvmConfig.ID).
		Uint64("CurrentBlockNumber", currentBlockNumber).
		Int("TotalLogs", logCounter).
		Msg("[EvmClient] [FinishRecover] recovered all events")
	return nil
}
func extractRecoverRange(errMsg string) (int, error) {
	re := regexp.MustCompile(`You can make eth_getLogs requests with up to a (\d+) block range`)
	match := re.FindStringSubmatch(errMsg)

	if len(match) > 1 {
		value, err := strconv.Atoi(match[1])
		if err != nil {
			return 0, fmt.Errorf("failed to parse number: %w", err)
		}
		return value, nil
	}
	return 0, fmt.Errorf("no match found")
}

// For speed up recover process, we split the recover range into multiple chunks and recover them in parallel
// func (c *EvmClient) RecoverEventsByRange(ctx context.Context, topics []common.Hash,
// 	fromBlock uint64, toBlock uint64, recoverRange uint64) error {

// }
func (c *EvmClient) UpdateLastCheckPoint(events map[string]abi.Event, logs []types.Log, lastBlock uint64) {
	eventCheckPoints := map[string]scalarnet.EventCheckPoint{}
	for _, txLog := range logs {
		topic := txLog.Topics[0].String()
		event, ok := events[topic]
		if !ok {
			log.Error().Str("topic", topic).Any("txLog", txLog).Msg("[EvmClient] [UpdateLastCheckPoint] event not found")
			continue
		}
		checkpoint, ok := eventCheckPoints[event.Name]
		if !ok {
			checkpoint = scalarnet.EventCheckPoint{
				ChainName:   c.EvmConfig.ID,
				EventName:   event.Name,
				BlockNumber: lastBlock,
				LogIndex:    txLog.Index,
				TxHash:      txLog.TxHash.String(),
			}
			eventCheckPoints[event.Name] = checkpoint
		} else {
			checkpoint.BlockNumber = lastBlock
			checkpoint.LogIndex = txLog.Index
			checkpoint.TxHash = txLog.TxHash.String()
		}
	}
	c.dbAdapter.UpdateLastEventCheckPoints(eventCheckPoints)
}
