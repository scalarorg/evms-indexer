package evm

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"
)

// Go routine for process missing logs
// func (c *EvmClient) ProcessMissingLogs() {
// 	gatewayAbi, err := getScalarGatewayAbi()
// 	if err != nil {
// 		log.Error().Err(err).Msg("[EvmClient] [ProcessMissingLogs] failed to get scalar gateway abi")
// 		return
// 	}
// 	mapEvents := map[string]*abi.Event{}
// 	for _, event := range gatewayAbi.Events {
// 		mapEvents[event.ID.String()] = &event
// 	}
// 	for {
// 		logs := c.MissingLogs.GetLogs(100)
// 		if len(logs) == 0 {
// 			if c.MissingLogs.IsRecovered() {
// 				log.Info().Str("Chain", c.EvmConfig.ID).Msg("[EvmClient] [ProcessMissingLogs] no logs to process, recovered flag is true, exit")
// 				break
// 			} else {
// 				log.Info().Str("Chain", c.EvmConfig.ID).Msg("[EvmClient] [ProcessMissingLogs] no logs to process, recover is in progress, sleep 1 second then continue")
// 				time.Sleep(time.Second)
// 				continue
// 			}
// 		}
// 		log.Info().Str("Chain", c.EvmConfig.ID).Int("Number of logs", len(logs)).Msg("[EvmClient] [ProcessMissingLogs] processing logs")
// 		for _, txLog := range logs {
// 			topic := txLog.Topics[0].String()
// 			event, ok := mapEvents[topic]
// 			if !ok {
// 				log.Error().Str("topic", topic).Any("txLog", txLog).Msg("[EvmClient] [ProcessMissingLogs] event not found")
// 				continue
// 			}
// 			log.Debug().
// 				Str("chainId", c.EvmConfig.GetId()).
// 				Str("eventName", event.Name).
// 				Str("txHash", txLog.TxHash.String()).
// 				Msg("[EvmClient] [ProcessMissingLogs] start processing missing event")

// 			err := c.handleEventLog(event, txLog)
// 			if err != nil {
// 				log.Error().Err(err).Msg("[EvmClient] [ProcessMissingLogs] failed to handle event log")
// 			}

// 		}
// 	}
// 	log.Info().Str("Chain", c.EvmConfig.ID).Msg("[EvmClient] [ProcessMissingLogs] finished processing all missing evm events")
// }

func (c *EvmClient) RecoverEvents(ctx context.Context, eventNames []string, currentBlockNumber uint64) error {
	// Use the new simplified batch recovery method
	return c.RecoverAllEventsBatch(ctx)
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

// For speed up recover process, we split the recover range into multiple chunks and recover them in parallel
// func (c *EvmClient) RecoverEventsByRange(ctx context.Context, topics []common.Hash,
// 	fromBlock uint64, toBlock uint64, recoverRange uint64) error {

// }
