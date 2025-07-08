package evm

import (
	"context"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rs/zerolog/log"
)

func (c *EvmClient) SubscribeAllEvents(ctx context.Context, topics []common.Hash, logsChan chan<- types.Log) error {
	log.Info().Hex("GatewayAddress", c.GatewayAddress.Bytes()).
		Int("topics length", len(topics)).
		Msgf("[EvmClient] [SubscribeAllEvents] subscribing to events")
	query := ethereum.FilterQuery{
		Addresses: []common.Address{c.GatewayAddress},
		Topics: [][]common.Hash{
			topics,
		},
	}

	sub, err := c.Client.SubscribeFilterLogs(context.Background(), query, logsChan)
	if err != nil {
		log.Error().Err(err).Msgf("[EvmClient] [SubscribeAllEvents] failed to subscribe to events")
		return err
	}
	c.subscriptions = sub
	return nil
}
