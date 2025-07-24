package evm

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rs/zerolog/log"
	evmAbi "github.com/scalarorg/evms-indexer/pkg/evm/abi"
	contracts "github.com/scalarorg/evms-indexer/pkg/evm/contracts/generated"
	"github.com/scalarorg/evms-indexer/pkg/evm/parser"
)

type ValidEvmEvent interface {
	*contracts.IScalarGatewayContractCallApproved |
		*contracts.IScalarGatewayContractCall |
		*contracts.IScalarGatewayContractCallWithToken |
		*contracts.IScalarGatewayRedeemToken |
		*contracts.IScalarGatewayExecuted |
		*contracts.IScalarGatewayTokenSent |
		*contracts.IScalarGatewaySwitchPhase |
		*contracts.IScalarGatewayTokenDeployed
}

// Return contracts.{EVENT} and gorm.Model
func parseEventLog(sourceChain string, event *abi.Event, txLog types.Log) (any, any, error) {
	switch event.Name {
	case evmAbi.EVENT_EVM_TOKEN_SENT:
		tokenSent := &contracts.IScalarGatewayTokenSent{
			Raw: txLog,
		}
		err := parser.ParseEventData(&txLog, event, tokenSent)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to parse event %s: %w", event.Name, err)
		}
		model := parser.TokenSentEvent2Model(sourceChain, tokenSent)
		return tokenSent, model, nil
	case evmAbi.EVENT_EVM_CONTRACT_CALL_WITH_TOKEN:
		contractCallWithToken := &contracts.IScalarGatewayContractCallWithToken{
			Raw: txLog,
		}
		err := parser.ParseEventData(&txLog, event, contractCallWithToken)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to parse event %s: %w", event.Name, err)
		}
		model := parser.ContractCallWithTokenEvent2Model(sourceChain, contractCallWithToken)
		log.Info().Any("model", model).Any("contractCallWithToken", contractCallWithToken).Msg("ContractCallWithTokenEvent2Model")
		return contractCallWithToken, model, nil
	case evmAbi.EVENT_EVM_CONTRACT_CALL:
		contractCall := &contracts.IScalarGatewayContractCall{
			Raw: txLog,
		}
		err := parser.ParseEventData(&txLog, event, contractCall)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to parse event %s: %w", event.Name, err)
		}
		model := parser.ContractCallEvent2Model(sourceChain, contractCall)
		return contractCall, model, nil
	case evmAbi.EVENT_EVM_CONTRACT_CALL_APPROVED:
		contractCallApproved := &contracts.IScalarGatewayContractCallApproved{
			Raw: txLog,
		}
		err := parser.ParseEventData(&txLog, event, contractCallApproved)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to parse event %s: %w", event.Name, err)
		}
		model := parser.ContractCallApprovedEvent2Model(sourceChain, contractCallApproved)
		return contractCallApproved, model, nil
	case evmAbi.EVENT_EVM_REDEEM_TOKEN:
		redeemToken := &contracts.IScalarGatewayRedeemToken{
			Raw: txLog,
		}
		err := parser.ParseEventData(&txLog, event, redeemToken)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to parse event %s: %w", event.Name, err)
		}
		model := parser.RedeemTokenEvent2Model(sourceChain, redeemToken)
		return redeemToken, model, nil
	case evmAbi.EVENT_EVM_COMMAND_EXECUTED:
		executed := &contracts.IScalarGatewayExecuted{
			Raw: txLog,
		}
		err := parser.ParseEventData(&txLog, event, executed)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to parse event %s: %w", event.Name, err)
		}
		model := parser.CommandExecutedEvent2Model(sourceChain, executed)
		return executed, model, nil
	case evmAbi.EVENT_EVM_TOKEN_DEPLOYED:
		tokenDeployed := &contracts.IScalarGatewayTokenDeployed{
			Raw: txLog,
		}
		err := parser.ParseEventData(&txLog, event, tokenDeployed)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to parse event %s: %w", event.Name, err)
		}
		model := parser.TokenDeployedEvent2Model(sourceChain, tokenDeployed)
		return tokenDeployed, model, nil
	case evmAbi.EVENT_EVM_SWITCHED_PHASE:
		switchedPhase := &contracts.IScalarGatewaySwitchPhase{
			Raw: txLog,
		}
		err := parser.ParseEventData(&txLog, event, switchedPhase)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to parse event %s: %w", event.Name, err)
		}
		model := parser.SwitchPhaseEvent2Model(sourceChain, switchedPhase)
		return switchedPhase, model, nil
	default:
		return nil, nil, fmt.Errorf("invalid event type for %s: %T", event.Name, txLog)
	}
}
