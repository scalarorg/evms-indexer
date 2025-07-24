package parser

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/scalarorg/data-models/chains"
	contracts "github.com/scalarorg/evms-indexer/pkg/evm/contracts/generated"
	"github.com/scalarorg/evms-indexer/pkg/types"
	"github.com/scalarorg/evms-indexer/pkg/utils"
)

func ContractCallEvent2Model(sourceChain string, event *contracts.IScalarGatewayContractCall) *chains.ContractCall {
	eventId := fmt.Sprintf("%s-%d", utils.NormalizeHash(event.Raw.TxHash.String()), event.Raw.Index)
	senderAddress := event.Sender.String()

	contractCall := chains.ContractCall{
		EventID:     eventId,
		TxHash:      event.Raw.TxHash.String(),
		BlockNumber: event.Raw.BlockNumber,
		LogIndex:    event.Raw.Index,
		SourceChain: sourceChain,
		//3 follows field are used for query to get back payload, so need to convert to lower case
		DestinationChain: event.DestinationChain,
		SourceAddress:    strings.ToLower(senderAddress),
		PayloadHash:      utils.NormalizeHash(hex.EncodeToString(event.PayloadHash[:])),
		Payload:          event.Payload,
	}
	return &contractCall
}

func ContractCallApprovedEvent2Model(sourceChain string, event *contracts.IScalarGatewayContractCallApproved) *chains.ContractCallApproved {
	txHash := event.Raw.TxHash.String()
	eventId := strings.ToLower(fmt.Sprintf("%s-%d-%d", txHash, event.SourceEventIndex, event.Raw.Index))
	sourceEventIndex := uint64(0)
	if event.SourceEventIndex != nil && event.SourceEventIndex.IsUint64() {
		sourceEventIndex = event.SourceEventIndex.Uint64()
	}
	record := chains.ContractCallApproved{
		EventID:          eventId,
		SourceChain:      event.SourceChain,
		DestinationChain: sourceChain,
		TxHash:           strings.ToLower(txHash),
		CommandID:        hex.EncodeToString(event.CommandId[:]),
		Sender:           strings.ToLower(event.SourceAddress),
		ContractAddress:  strings.ToLower(event.ContractAddress.String()),
		PayloadHash:      strings.ToLower(hex.EncodeToString(event.PayloadHash[:])),
		SourceTxHash:     strings.ToLower(hex.EncodeToString(event.SourceTxHash[:])),
		SourceEventIndex: sourceEventIndex,
	}
	return &record
}

func ContractCallWithTokenEvent2Model(sourceChain string, event *contracts.IScalarGatewayContractCallWithToken) *chains.ContractCallWithToken {
	eventId := fmt.Sprintf("%s-%d", utils.NormalizeHash(event.Raw.TxHash.String()), event.Raw.Index)
	senderAddress := event.Sender.String()

	chainInfoBytes := types.ChainInfoBytes{}
	err := chainInfoBytes.FromString(event.DestinationChain)
	if err != nil {
		log.Error().Err(err).Msg("failed to get chain info")
		return nil
	}
	chainType := chainInfoBytes.ChainType()
	if chainType != types.ChainTypeBitcoin {
		log.Error().Any("chainType", chainType).Msg("chainType is not bitcoin")
		return nil
	}
	//Extract user destination address
	payload := ContractCallWithTokenPayload{}
	err = payload.Parse(event.Payload)
	if err != nil {
		log.Error().Err(err).Msg("failed to parse payload")
		return nil
	}
	destinationAddress, err := payload.GetDestinationAddress(chainInfoBytes.ChainID())
	if err != nil {
		log.Error().Err(err).Msg("failed to get destination address")
		return nil
	}

	callContract := chains.ContractCall{
		EventID:     eventId,
		TxHash:      utils.NormalizeHash(event.Raw.TxHash.String()),
		BlockNumber: event.Raw.BlockNumber,
		LogIndex:    event.Raw.Index,
		SourceChain: sourceChain,
		//3 follows field are used for query to get back payload, so need to convert to lower case
		DestinationChain:   event.DestinationChain,
		DestinationAddress: utils.NormalizeAddress(destinationAddress, chainType),
		SourceAddress:      utils.NormalizeAddress(senderAddress, chainType),
		PayloadHash:        utils.NormalizeHash(hex.EncodeToString(event.PayloadHash[:])),
		Payload:            event.Payload,
	}
	contractCallWithToken := chains.ContractCallWithToken{
		ContractCall:         callContract,
		TokenContractAddress: utils.NormalizeAddress(event.DestinationContractAddress, chainType),
		Symbol:               event.Symbol,
		Amount:               event.Amount.Uint64(),
	}
	return &contractCallWithToken
}

func CommandExecutedEvent2Model(sourceChain string, event *contracts.IScalarGatewayExecuted) *chains.CommandExecuted {
	return &chains.CommandExecuted{
		SourceChain: sourceChain,
		Address:     event.Raw.Address.String(),
		TxHash:      strings.ToLower(event.Raw.TxHash.String()),
		BlockNumber: uint64(event.Raw.BlockNumber),
		LogIndex:    uint(event.Raw.Index),
		CommandID:   hex.EncodeToString(event.CommandId[:]),
	}
}

func TokenDeployedEvent2Model(sourceChain string, event *contracts.IScalarGatewayTokenDeployed) *chains.TokenDeployed {
	return &chains.TokenDeployed{
		SourceChain:  sourceChain,
		BlockNumber:  uint64(event.Raw.BlockNumber),
		TxHash:       event.Raw.TxHash.String(),
		Symbol:       event.Symbol,
		TokenAddress: event.TokenAddresses.String(),
	}
}

func SwitchPhaseEvent2Model(sourceChain string, event *contracts.IScalarGatewaySwitchPhase) *chains.SwitchedPhase {
	return &chains.SwitchedPhase{
		SourceChain:       sourceChain,
		BlockNumber:       uint64(event.Raw.BlockNumber),
		TxHash:            event.Raw.TxHash.String(),
		CustodianGroupUid: hex.EncodeToString(event.CustodianGroupId[:]),
		SessionSequence:   event.Sequence,
		From:              event.From,
		To:                event.To,
	}
}

func TokenSentEvent2Model(sourceChain string, event *contracts.IScalarGatewayTokenSent) *chains.TokenSent {
	normalizedTxHash := utils.NormalizeHash(event.Raw.TxHash.String())
	eventId := fmt.Sprintf("%s-%d", normalizedTxHash, event.Raw.Index)
	senderAddress := event.Sender.String()
	return &chains.TokenSent{
		EventID:     eventId,
		SourceChain: sourceChain,
		TxHash:      normalizedTxHash,
		BlockNumber: event.Raw.BlockNumber,
		LogIndex:    event.Raw.Index,
		//3 follows field are used for query to get back payload, so need to convert to lower case
		SourceAddress:      strings.ToLower(senderAddress),
		DestinationChain:   event.DestinationChain,
		DestinationAddress: strings.ToLower(event.DestinationAddress),
		Symbol:             event.Symbol,
		//TokenContractAddress: c.GetTokenContractAddressFromSymbol(event.Symbol),
		Amount: event.Amount.Uint64(),
		Status: chains.TokenSentStatusPending,
	}
}

func RedeemTokenEvent2Model(sourceChain string, event *contracts.IScalarGatewayRedeemToken) *chains.EvmRedeemTx {
	eventId := fmt.Sprintf("%s-%d", utils.NormalizeHash(event.Raw.TxHash.String()), event.Raw.Index)
	senderAddress := event.Sender.String()

	chainInfoBytes := types.ChainInfoBytes{}
	err := chainInfoBytes.FromString(event.DestinationChain)
	if err != nil {
		return nil
	}

	payload := ContractCallWithTokenPayload{}
	err = payload.Parse(event.Payload)
	if err != nil {
		return nil
	}
	destinationAddress, err := payload.GetDestinationAddress(chainInfoBytes.ChainID())
	if err != nil {
		return nil
	}

	return &chains.EvmRedeemTx{
		EventID:              eventId,
		TxHash:               event.Raw.TxHash.String(),
		BlockNumber:          uint64(event.Raw.BlockNumber),
		LogIndex:             uint(event.Raw.Index),
		SourceChain:          sourceChain,
		SourceAddress:        senderAddress,
		DestinationChain:     event.DestinationChain,
		DestinationAddress:   destinationAddress,
		Payload:              event.Payload,
		PayloadHash:          hex.EncodeToString(event.PayloadHash[:]),
		TokenContractAddress: event.DestinationContractAddress,
		Symbol:               event.Symbol,
		Amount:               event.Amount.Uint64(),
		CustodianGroupUid:    hex.EncodeToString(event.CustodianGroupUID[:]),
		SessionSequence:      event.Sequence,
	}
}
