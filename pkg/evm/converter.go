package evm

import (
	"encoding/hex"
	"fmt"
	"strings"

	chains "github.com/scalarorg/data-models/chains"
	"github.com/scalarorg/data-models/scalarnet"
	contracts "github.com/scalarorg/evms-indexer/pkg/evm/contracts/generated"
	"github.com/scalarorg/evms-indexer/pkg/types"
	"github.com/scalarorg/evms-indexer/pkg/utils"
)

func (c *EvmClient) ContractCallEvent2Model(event *contracts.IScalarGatewayContractCall) (chains.ContractCall, error) {
	//id := strings.ToLower(fmt.Sprintf("%s-%d", event.Raw.TxHash.String(), event.Raw.Index))
	//Calculate eventId by Txhash-logIndex among logs in txreceipt (AxelarEvmModule)
	//https://github.com/scalarorg/scalar-core/blob/main/vald/evm/gateway_tx_confirmation.go#L73
	//Dec 30, use logIndex directly to avoid redundant request. This must aggrees with the scalar-core vald module
	// receipt, err := c.Client.TransactionReceipt(context.Background(), event.Raw.TxHash)
	// if err != nil {
	// 	return models.RelayData{}, fmt.Errorf("failed to get transaction receipt: %w", err)
	// }
	// var id string
	// for ind, log := range receipt.Logs {
	// 	if log.Index == event.Raw.Index {
	// 		id = fmt.Sprintf("%s-%d", event.Raw.TxHash.String(), ind)
	// 		break
	// 	}
	// }
	eventId := fmt.Sprintf("%s-%d", utils.NormalizeHash(event.Raw.TxHash.String()), event.Raw.Index)
	senderAddress := event.Sender.String()

	chainInfoBytes := types.ChainInfoBytes{}
	err := chainInfoBytes.FromString(event.DestinationChain)
	if err != nil {
		return chains.ContractCall{}, fmt.Errorf("failed to convert destination chain: %w", err)
	}
	contractCall := chains.ContractCall{
		EventID:     eventId,
		TxHash:      event.Raw.TxHash.String(),
		BlockNumber: event.Raw.BlockNumber,
		LogIndex:    event.Raw.Index,
		SourceChain: c.EvmConfig.GetId(),
		//3 follows field are used for query to get back payload, so need to convert to lower case
		DestinationChain: event.DestinationChain,
		SourceAddress:    utils.NormalizeAddress(senderAddress, chainInfoBytes.ChainType()),
		PayloadHash:      utils.NormalizeHash(hex.EncodeToString(event.PayloadHash[:])),
		Payload:          event.Payload,
	}
	return contractCall, nil
}

func (c *EvmClient) ContractCallWithToken2Model(event *contracts.IScalarGatewayContractCallWithToken) (*chains.ContractCallWithToken, error) {
	eventId := fmt.Sprintf("%s-%d", utils.NormalizeHash(event.Raw.TxHash.String()), event.Raw.Index)
	senderAddress := event.Sender.String()

	chainInfoBytes := types.ChainInfoBytes{}
	err := chainInfoBytes.FromString(event.DestinationChain)
	if err != nil {
		return nil, fmt.Errorf("failed to convert destination chain: %w", err)
	}
	chainType := chainInfoBytes.ChainType()
	if chainType != types.ChainTypeBitcoin {
		return nil, fmt.Errorf("unsupported chain type: %s", chainType)
	}
	//Extract user destination address
	payload := ContractCallWithTokenPayload{}
	err = payload.Parse(event.Payload)
	if err != nil {
		return nil, fmt.Errorf("failed to parse payload: %w", err)
	}
	destinationAddress, err := payload.GetDestinationAddress(chainInfoBytes.ChainID())
	if err != nil {
		return nil, fmt.Errorf("failed to get destination address: %w", err)
	}

	callContract := chains.ContractCall{
		EventID:     eventId,
		TxHash:      utils.NormalizeHash(event.Raw.TxHash.String()),
		BlockNumber: event.Raw.BlockNumber,
		LogIndex:    event.Raw.Index,
		SourceChain: c.EvmConfig.GetId(),
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
	return &contractCallWithToken, nil
}

// Use contractCallToken to store the redeem token event
func (c *EvmClient) RedeemTokenEvent2Model(event *contracts.IScalarGatewayRedeemToken) (*chains.ContractCallWithToken, error) {
	eventId := fmt.Sprintf("%s-%d", utils.NormalizeHash(event.Raw.TxHash.String()), event.Raw.Index)
	senderAddress := event.Sender.String()

	chainInfoBytes := types.ChainInfoBytes{}
	err := chainInfoBytes.FromString(event.DestinationChain)
	if err != nil {
		return nil, fmt.Errorf("failed to convert destination chain: %w", err)
	}

	payload := ContractCallWithTokenPayload{}
	err = payload.Parse(event.Payload)
	if err != nil {
		return nil, fmt.Errorf("failed to parse payload: %w", err)
	}
	destinationAddress, err := payload.GetDestinationAddress(chainInfoBytes.ChainID())
	if err != nil {
		return nil, fmt.Errorf("failed to get destination address: %w", err)
	}

	callContract := chains.ContractCall{
		EventID:     eventId,
		TxHash:      utils.NormalizeHash(event.Raw.TxHash.String()),
		BlockNumber: event.Raw.BlockNumber,
		LogIndex:    event.Raw.Index,
		SourceChain: c.EvmConfig.GetId(),
		//3 follows field are used for query to get back payload, so need to convert to lower case
		DestinationChain:   event.DestinationChain,
		DestinationAddress: utils.NormalizeAddress(destinationAddress, chainInfoBytes.ChainType()),
		SourceAddress:      utils.NormalizeAddress(senderAddress, chainInfoBytes.ChainType()),
		PayloadHash:        utils.NormalizeHash(hex.EncodeToString(event.PayloadHash[:])),
		Payload:            event.Payload,
	}
	contractCallWithToken := chains.ContractCallWithToken{
		ContractCall:         callContract,
		TokenContractAddress: utils.NormalizeAddress(event.DestinationContractAddress, chainInfoBytes.ChainType()),
		Symbol:               event.Symbol,
		Amount:               event.Amount.Uint64(),
		CustodianGroupUid:    hex.EncodeToString(event.CustodianGroupUID[:]),
		SessionSequence:      event.Sequence,
	}
	return &contractCallWithToken, nil
}

func (c *EvmClient) TokenSentEvent2Model(event *contracts.IScalarGatewayTokenSent) (chains.TokenSent, error) {
	normalizedTxHash := utils.NormalizeHash(event.Raw.TxHash.String())
	eventId := fmt.Sprintf("%s-%d", normalizedTxHash, event.Raw.Index)
	senderAddress := event.Sender.String()
	tokenSent := chains.TokenSent{
		EventID:     eventId,
		SourceChain: c.EvmConfig.GetId(),
		TxHash:      normalizedTxHash,
		BlockNumber: event.Raw.BlockNumber,
		LogIndex:    event.Raw.Index,
		//3 follows field are used for query to get back payload, so need to convert to lower case
		SourceAddress:        strings.ToLower(senderAddress),
		DestinationChain:     event.DestinationChain,
		DestinationAddress:   strings.ToLower(event.DestinationAddress),
		Symbol:               event.Symbol,
		TokenContractAddress: c.GetTokenContractAddressFromSymbol(event.Symbol),
		Amount:               event.Amount.Uint64(),
		Status:               chains.TokenSentStatusPending,
	}
	return tokenSent, nil
}

// Todo: Implement this function
func (c *EvmClient) GetTokenContractAddressFromSymbol(symbol string) string {
	address, ok := c.TokenAddresses[symbol]
	if ok {
		return address
	}
	//Try to get from the contract
	tokenAddress, err := c.Gateway.TokenAddresses(nil, symbol)
	if err != nil {
		return ""
	}
	c.TokenAddresses[symbol] = tokenAddress.String()
	return tokenAddress.String()
}
func (c *EvmClient) ContractCallApprovedEvent2Model(event *contracts.IScalarGatewayContractCallApproved) (scalarnet.ContractCallApproved, error) {
	txHash := event.Raw.TxHash.String()
	eventId := strings.ToLower(fmt.Sprintf("%s-%d-%d", txHash, event.SourceEventIndex, event.Raw.Index))
	sourceEventIndex := uint64(0)
	if event.SourceEventIndex != nil && event.SourceEventIndex.IsUint64() {
		sourceEventIndex = event.SourceEventIndex.Uint64()
	}
	record := scalarnet.ContractCallApproved{
		EventID:          eventId,
		SourceChain:      event.SourceChain,
		DestinationChain: c.EvmConfig.GetId(),
		TxHash:           strings.ToLower(txHash),
		CommandID:        hex.EncodeToString(event.CommandId[:]),
		Sender:           strings.ToLower(event.SourceAddress),
		ContractAddress:  strings.ToLower(event.ContractAddress.String()),
		PayloadHash:      strings.ToLower(hex.EncodeToString(event.PayloadHash[:])),
		SourceTxHash:     strings.ToLower(hex.EncodeToString(event.SourceTxHash[:])),
		SourceEventIndex: sourceEventIndex,
	}
	return record, nil
}

func (c *EvmClient) CommandExecutedEvent2Model(event *contracts.IScalarGatewayExecuted) chains.CommandExecuted {
	cmdExecuted := chains.CommandExecuted{
		SourceChain: c.EvmConfig.GetId(),
		Address:     event.Raw.Address.String(),
		TxHash:      strings.ToLower(event.Raw.TxHash.String()),
		BlockNumber: uint64(event.Raw.BlockNumber),
		LogIndex:    uint(event.Raw.Index),
		CommandID:   hex.EncodeToString(event.CommandId[:]),
	}
	return cmdExecuted
}

func (c *EvmClient) TokenDeployedEvent2Model(event *contracts.IScalarGatewayTokenDeployed) chains.TokenDeployed {
	tokenDeployed := chains.TokenDeployed{
		Chain:        c.EvmConfig.GetId(),
		BlockNumber:  uint64(event.Raw.BlockNumber),
		TxHash:       event.Raw.TxHash.String(),
		Symbol:       event.Symbol,
		TokenAddress: event.TokenAddresses.String(),
	}
	return tokenDeployed
}

func (c *EvmClient) SwitchPhaseEvent2Model(event *contracts.IScalarGatewaySwitchPhase) chains.SwitchedPhase {
	switchPhase := chains.SwitchedPhase{
		Chain:             c.EvmConfig.GetId(),
		BlockNumber:       uint64(event.Raw.BlockNumber),
		TxHash:            event.Raw.TxHash.String(),
		CustodianGroupUid: hex.EncodeToString(event.CustodianGroupId[:]),
		SessionSequence:   event.Sequence,
		From:              event.From,
		To:                event.To,
	}
	return switchPhase
}
