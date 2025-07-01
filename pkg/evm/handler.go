package evm

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/scalarorg/data-models/chains"
	contracts "github.com/scalarorg/evms-indexer/pkg/evm/contracts/generated"
)

// func (ec *EvmClient) handleContractCall(event *contracts.IScalarGatewayContractCall) error {
// 	//0. Preprocess the event
// 	ec.preprocessContractCall(event)
// 	//1. Convert into a RelayData instance then store to the db
// 	contractCall, err := ec.ContractCallEvent2Model(event)
// 	if err != nil {
// 		return fmt.Errorf("failed to convert ContractCallEvent to RelayData: %w", err)
// 	}
// 	//2. update last checkpoint
// 	lastCheckpoint, err := ec.dbAdapter.GetLastEventCheckPoint(ec.EvmConfig.GetId(), EVENT_EVM_CONTRACT_CALL)
// 	if err != nil {
// 		log.Debug().Str("chainId", ec.EvmConfig.GetId()).
// 			Str("eventName", EVENT_EVM_CONTRACT_CALL).
// 			Msg("[EvmClient] [handleContractCall] Get event from begining")
// 	}
// 	if event.Raw.BlockNumber > lastCheckpoint.BlockNumber ||
// 		(event.Raw.BlockNumber == lastCheckpoint.BlockNumber && event.Raw.TxIndex > lastCheckpoint.LogIndex) {
// 		lastCheckpoint.BlockNumber = event.Raw.BlockNumber
// 		lastCheckpoint.TxHash = event.Raw.TxHash.String()
// 		lastCheckpoint.LogIndex = event.Raw.Index
// 		lastCheckpoint.EventKey = fmt.Sprintf("%s-%d-%d", event.Raw.TxHash.String(), event.Raw.BlockNumber, event.Raw.Index)
// 	}
// 	//3. store relay data to the db, update last checkpoint
// 	err = ec.dbAdapter.CreateContractCall(contractCall, lastCheckpoint)
// 	if err != nil {
// 		return fmt.Errorf("failed to create evm contract call: %w", err)
// 	}
// 	//2. Send to the bus
// 	confirmTxs := events.ConfirmTxsRequest{
// 		ChainName: ec.EvmConfig.GetId(),
// 		TxHashs:   map[string]string{contractCall.TxHash: contractCall.DestinationChain},
// 	}
// 	if ec.eventBus != nil {
// 		ec.eventBus.BroadcastEvent(&events.EventEnvelope{
// 			EventType:        EVENT_EVM_CONTRACT_CALL,
// 			DestinationChain: events.SCALAR_NETWORK_NAME,
// 			Data:             confirmTxs,
// 		})
// 	} else {
// 		log.Warn().Msg("[EvmClient] [handleContractCall] event bus is undefined")
// 	}
// 	return nil
// }
// func (ec *EvmClient) preprocessContractCall(event *contracts.IScalarGatewayContractCall) error {
// 	log.Info().
// 		Str("sender", event.Sender.Hex()).
// 		Str("destinationChain", event.DestinationChain).
// 		Str("destinationContractAddress", event.DestinationContractAddress).
// 		Str("payloadHash", hex.EncodeToString(event.PayloadHash[:])).
// 		Str("txHash", event.Raw.TxHash.String()).
// 		Uint("logIndex", event.Raw.Index).
// 		Uint("txIndex", event.Raw.TxIndex).
// 		Str("logData", hex.EncodeToString(event.Raw.Data)).
// 		Msg("[EvmClient] [preprocessContractCall] Start handle Contract call")
// 	//Todo: validate the event
// 	return nil
// }

func (ec *EvmClient) HandleContractCallWithToken(event *contracts.IScalarGatewayContractCallWithToken) error {
	//0. Preprocess the event
	ec.preprocessContractCallWithToken(event)
	//Get block header
	err := ec.fetchBlockHeader(event.Raw.BlockNumber)
	if err != nil {
		log.Error().Err(err).Msgf("[EvmClient] [HandleContractCallWithToken] failed to fetch block header %d", event.Raw.BlockNumber)
	}
	//1. Convert into a RelayData instance then store to the db
	contractCallWithToken, err := ec.ContractCallWithToken2Model(event)
	if err != nil {
		return fmt.Errorf("failed to convert ContractCallEvent to ContractCallWithToken: %w", err)
	}
	//2. update last checkpoint
	lastCheckpoint, err := ec.dbAdapter.GetLastEventCheckPoint(ec.EvmConfig.GetId(), EVENT_EVM_CONTRACT_CALL_WITH_TOKEN, ec.EvmConfig.StartBlock)
	if err != nil {
		log.Debug().Str("chainId", ec.EvmConfig.GetId()).
			Str("eventName", EVENT_EVM_CONTRACT_CALL_WITH_TOKEN).
			Msg("[EvmClient] [handleContractCallWithToken] Get event from begining")
	}
	if event.Raw.BlockNumber > lastCheckpoint.BlockNumber ||
		(event.Raw.BlockNumber == lastCheckpoint.BlockNumber && event.Raw.TxIndex > lastCheckpoint.LogIndex) {
		lastCheckpoint.BlockNumber = event.Raw.BlockNumber
		lastCheckpoint.TxHash = event.Raw.TxHash.String()
		lastCheckpoint.LogIndex = event.Raw.Index
		lastCheckpoint.EventKey = fmt.Sprintf("%s-%d-%d", event.Raw.TxHash.String(), event.Raw.BlockNumber, event.Raw.Index)
	}
	//3. store relay data to the db, update last checkpoint
	err = ec.dbAdapter.CreateContractCallWithToken(contractCallWithToken, lastCheckpoint)
	if err != nil {
		return fmt.Errorf("failed to create evm contract call: %w", err)
	}
	return nil
}

func (ec *EvmClient) HandleRedeemToken(event *contracts.IScalarGatewayRedeemToken) error {
	//0. Preprocess the event
	log.Info().Str("Chain", ec.EvmConfig.ID).Any("event", event).Msg("[EvmClient] [HandleRedeemToken] Start processing evm redeem token")
	err := ec.fetchBlockHeader(event.Raw.BlockNumber)
	if err != nil {
		log.Error().Err(err).Msgf("[EvmClient] [HandleRedeemToken] failed to fetch block header %d", event.Raw.BlockNumber)
	}
	//1. Convert into a RelayData instance then store to the db
	redeemToken, err := ec.RedeemTokenEvent2Model(event)
	if err != nil {
		return fmt.Errorf("failed to convert ContractCallEvent to ContractCallWithToken: %w", err)
	}
	//2. update last checkpoint
	lastCheckpoint, err := ec.dbAdapter.GetLastEventCheckPoint(ec.EvmConfig.GetId(), EVENT_EVM_REDEEM_TOKEN, ec.EvmConfig.StartBlock)
	if err != nil {
		log.Debug().Str("chainId", ec.EvmConfig.GetId()).
			Str("eventName", EVENT_EVM_CONTRACT_CALL_WITH_TOKEN).
			Msg("[EvmClient] [handleContractCallWithToken] Get event from begining")
	}
	if event.Raw.BlockNumber > lastCheckpoint.BlockNumber ||
		(event.Raw.BlockNumber == lastCheckpoint.BlockNumber && event.Raw.TxIndex > lastCheckpoint.LogIndex) {
		lastCheckpoint.BlockNumber = event.Raw.BlockNumber
		lastCheckpoint.TxHash = event.Raw.TxHash.String()
		lastCheckpoint.LogIndex = event.Raw.Index
		lastCheckpoint.EventKey = fmt.Sprintf("%s-%d-%d", event.Raw.TxHash.String(), event.Raw.BlockNumber, event.Raw.Index)
	}
	//3. store relay data to the db, update last checkpoint
	err = ec.dbAdapter.CreateContractCallWithToken(redeemToken, lastCheckpoint)
	if err != nil {
		return fmt.Errorf("failed to create evm contract call: %w", err)
	}
	return nil
}

func (ec *EvmClient) preprocessContractCallWithToken(event *contracts.IScalarGatewayContractCallWithToken) error {
	log.Info().
		Str("sender", event.Sender.Hex()).
		Str("destinationChain", event.DestinationChain).
		Str("destinationContractAddress", event.DestinationContractAddress).
		Str("payloadHash", hex.EncodeToString(event.PayloadHash[:])).
		Str("Symbol", event.Symbol).
		Uint64("Amount", event.Amount.Uint64()).
		Str("txHash", event.Raw.TxHash.String()).
		Uint("logIndex", event.Raw.Index).
		Uint("txIndex", event.Raw.TxIndex).
		Str("logData", hex.EncodeToString(event.Raw.Data)).
		Msg("[EvmClient] [preprocessContractCallWithToken] Start handle Contract call with token")
	//Todo: validate the event
	return nil
}

func (ec *EvmClient) HandleTokenSent(event *contracts.IScalarGatewayTokenSent) error {
	//0. Preprocess the event
	ec.preprocessTokenSent(event)
	err := ec.fetchBlockHeader(event.Raw.BlockNumber)
	if err != nil {
		log.Error().Err(err).Msgf("[EvmClient] [HandleTokenSent] failed to fetch block header %d", event.Raw.BlockNumber)
	}
	//1. Convert into a RelayData instance then store to the db
	tokenSent, err := ec.TokenSentEvent2Model(event)
	if err != nil {
		log.Error().Err(err).Msg("[EvmClient] [HandleTokenSent] failed to convert TokenSentEvent to model data")
		return err
	}
	//For evm, the token sent is verified immediately by the scalarnet
	tokenSent.Status = chains.TokenSentStatusVerifying
	//2. update last checkpoint
	lastCheckpoint, err := ec.dbAdapter.GetLastEventCheckPoint(ec.EvmConfig.GetId(), EVENT_EVM_TOKEN_SENT, ec.EvmConfig.StartBlock)
	if err != nil {
		log.Debug().Str("chainId", ec.EvmConfig.GetId()).
			Str("eventName", EVENT_EVM_TOKEN_SENT).
			Msg("[EvmClient] [handleTokenSent] Get event from begining")
	}
	if event.Raw.BlockNumber > lastCheckpoint.BlockNumber ||
		(event.Raw.BlockNumber == lastCheckpoint.BlockNumber && event.Raw.TxIndex > lastCheckpoint.LogIndex) {
		lastCheckpoint.BlockNumber = event.Raw.BlockNumber
		lastCheckpoint.TxHash = event.Raw.TxHash.String()
		lastCheckpoint.LogIndex = event.Raw.Index
		lastCheckpoint.EventKey = fmt.Sprintf("%s-%d-%d", event.Raw.TxHash.String(), event.Raw.BlockNumber, event.Raw.Index)
	}
	//3. store relay data to the db, update last checkpoint
	err = ec.dbAdapter.SaveTokenSent(tokenSent, lastCheckpoint)
	if err != nil {
		return fmt.Errorf("failed to create evm token send: %w", err)
	}
	return nil
}

func (ec *EvmClient) preprocessTokenSent(event *contracts.IScalarGatewayTokenSent) error {
	log.Info().
		Str("sender", event.Sender.Hex()).
		Str("destinationChain", event.DestinationChain).
		Str("destinationAddress", event.DestinationAddress).
		Str("txHash", event.Raw.TxHash.String()).
		Str("symbol", event.Symbol).
		Uint64("amount", event.Amount.Uint64()).
		Uint("logIndex", event.Raw.Index).
		Uint("txIndex", event.Raw.TxIndex).
		Str("logData", hex.EncodeToString(event.Raw.Data)).
		Msg("[EvmClient] [preprocessTokenSent] Start handle TokenSent")
	//Todo: validate the event
	return nil
}

func (ec *EvmClient) HandleContractCallApproved(event *contracts.IScalarGatewayContractCallApproved) error {
	//0. Preprocess the event
	err := ec.preprocessContractCallApproved(event)
	if err != nil {
		return fmt.Errorf("failed to preprocess contract call approved: %w", err)
	}
	err = ec.fetchBlockHeader(event.Raw.BlockNumber)
	if err != nil {
		log.Error().Err(err).Msgf("[EvmClient] [HandleRedeemToken] failed to fetch block header %d", event.Raw.BlockNumber)
	}
	//1. Convert into a RelayData instance then store to the db
	contractCallApproved, err := ec.ContractCallApprovedEvent2Model(event)
	if err != nil {
		return fmt.Errorf("failed to convert ContractCallApprovedEvent to RelayData: %w", err)
	}
	err = ec.dbAdapter.SaveSingleValue(&contractCallApproved)
	if err != nil {
		return fmt.Errorf("failed to create contract call approved: %w", err)
	}
	// Find relayData from the db by combination (contractAddress, sourceAddress, payloadHash)
	// This contract call (initiated by the user call to the source chain) is approved by EVM network
	// So anyone can execute it on the EVM by broadcast the corresponding payload to protocol's smart contract on the destination chain
	destContractAddress := strings.TrimLeft(event.ContractAddress.Hex(), "0x")
	sourceAddress := strings.TrimLeft(event.SourceAddress, "0x")
	payloadHash := strings.TrimLeft(hex.EncodeToString(event.PayloadHash[:]), "0x")
	relayDatas, err := ec.dbAdapter.FindContractCallByParams(sourceAddress, destContractAddress, payloadHash)
	if err != nil {
		log.Error().Err(err).Msg("[EvmClient] [handleContractCallApproved] find relay data")
		return err
	}
	log.Debug().Str("contractAddress", event.ContractAddress.String()).
		Str("sourceAddress", event.SourceAddress).
		Str("payloadHash", hex.EncodeToString(event.PayloadHash[:])).
		Any("relayDatas count", len(relayDatas)).
		Msg("[EvmClient] [handleContractCallApproved] query relaydata by ContractCall")
	return nil
}

func (ec *EvmClient) preprocessContractCallApproved(event *contracts.IScalarGatewayContractCallApproved) error {
	log.Info().Any("event", event).Msgf("[EvmClient] [handleContractCallApproved]")
	//Todo: validate the event
	return nil
}

func (ec *EvmClient) HandleCommandExecuted(event *contracts.IScalarGatewayExecuted) error {
	//0. Preprocess the event
	//ec.preprocessCommandExecuted(event)
	err := ec.fetchBlockHeader(event.Raw.BlockNumber)
	if err != nil {
		log.Error().Err(err).Msgf("[EvmClient] [HandleCommandExecuted] failed to fetch block header %d", event.Raw.BlockNumber)
	}
	//1. Convert into a RelayData instance then store to the db
	cmdExecuted := ec.CommandExecutedEvent2Model(event)
	if ec.dbAdapter != nil {
		err = ec.dbAdapter.SaveCommandExecuted(&cmdExecuted)
		if err != nil {
			log.Error().Err(err).Msg("[EvmClient] [HandleCommandExecuted] failed to save evm executed to the db")
			return fmt.Errorf("failed to create evm executed: %w", err)
		}
	}
	return nil
}

func (ec *EvmClient) preprocessCommandExecuted(event *contracts.IScalarGatewayExecuted) error {
	log.Info().Any("event", event).Msg("[EvmClient] [ExecutedHandler] Start processing evm command executed")
	//Todo: validate the event
	return nil
}

func (ec *EvmClient) HandleTokenDeployed(event *contracts.IScalarGatewayTokenDeployed) error {
	//0. Preprocess the event
	log.Info().Any("event", event).Msg("[EvmClient] [HandleTokenDeployed] Start processing evm token deployed")
	//1. Convert into a RelayData instance then store to the db
	tokenDeployed := ec.TokenDeployedEvent2Model(event)
	if ec.dbAdapter != nil {
		err := ec.dbAdapter.SaveTokenDeployed(&tokenDeployed)
		if err != nil {
			return fmt.Errorf("failed to create evm token deployed: %w", err)
		}
	}
	return nil
}

func (ec *EvmClient) HandleSwitchPhase(event *contracts.IScalarGatewaySwitchPhase) error {
	//0. Preprocess the event
	log.Info().Str("Chain", ec.EvmConfig.ID).Any("event", event).Msg("[EvmClient] [HandleSwitchPhase] Start processing evm switch phase")
	//1. Convert into a RelayData instance then store to the db
	switchPhase := ec.SwitchPhaseEvent2Model(event)
	if ec.dbAdapter != nil {
		err := ec.dbAdapter.SaveSingleValue(&switchPhase)
		if err != nil {
			return fmt.Errorf("failed to create evm switch phase: %w", err)
		}
	}
	return nil
}
