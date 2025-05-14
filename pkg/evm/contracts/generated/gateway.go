// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ScalarGatewaySession is an auto generated low-level Go binding around an user-defined struct.
type ScalarGatewaySession struct {
	Sequence uint64
	Phase    uint8
}

// IScalarGatewayMetaData contains all meta data concerning the IScalarGateway contract.
var IScalarGatewayMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"authModule_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenDeployer_\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allTokensFrozen\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"approveContractCall\",\"inputs\":[{\"name\":\"params\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"commandId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"approveContractCallWithMint\",\"inputs\":[{\"name\":\"params\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"commandId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"authModule\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"burnToken\",\"inputs\":[{\"name\":\"params\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"callContract\",\"inputs\":[{\"name\":\"destinationChain\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"destinationContractAddress\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"callContractWithToken\",\"inputs\":[{\"name\":\"destinationChain\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"destinationContractAddress\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"contractId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"deployToken\",\"inputs\":[{\"name\":\"params\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"input\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAddress\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBool\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBytes\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getInt\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSession\",\"inputs\":[{\"name\":\"_custodianGroupId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structScalarGateway.Session\",\"components\":[{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"phase\",\"type\":\"uint8\",\"internalType\":\"enumScalarGateway.Phase\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getString\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUint\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"governance\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"implementation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isCommandExecuted\",\"inputs\":[{\"name\":\"commandId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isContractCallAndMintApproved\",\"inputs\":[{\"name\":\"commandId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChain\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"sourceAddress\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"contractAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"payloadHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isContractCallApproved\",\"inputs\":[{\"name\":\"commandId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChain\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"sourceAddress\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"contractAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"payloadHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mintLimiter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mintToken\",\"inputs\":[{\"name\":\"params\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"redeemToken\",\"inputs\":[{\"name\":\"params\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerCustodianGroup\",\"inputs\":[{\"name\":\"params\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"sendToken\",\"inputs\":[{\"name\":\"destinationChain\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"destinationAddress\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"sessions\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"phase\",\"type\":\"uint8\",\"internalType\":\"enumScalarGateway.Phase\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setTokenMintLimits\",\"inputs\":[{\"name\":\"symbols\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"limits\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setup\",\"inputs\":[{\"name\":\"params\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"switchPhase\",\"inputs\":[{\"name\":\"params\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"tokenAddresses\",\"inputs\":[{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokenCustodianGroupIds\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokenDeployer\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokenFrozen\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"tokenMintAmount\",\"inputs\":[{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokenMintLimit\",\"inputs\":[{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferGovernance\",\"inputs\":[{\"name\":\"newGovernance\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferMintLimiter\",\"inputs\":[{\"name\":\"newMintLimiter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOperatorship\",\"inputs\":[{\"name\":\"newOperatorsData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgrade\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"newImplementationCodeHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"setupParams\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"validateContractCall\",\"inputs\":[{\"name\":\"commandId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChain\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"sourceAddress\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"payloadHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"valid\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"validateContractCallAndMint\",\"inputs\":[{\"name\":\"commandId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChain\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"sourceAddress\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"payloadHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"valid\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ContractCall\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"destinationChain\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"destinationContractAddress\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"payloadHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"payload\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ContractCallApproved\",\"inputs\":[{\"name\":\"commandId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"sourceChain\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"sourceAddress\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"contractAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"payloadHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"sourceTxHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"sourceEventIndex\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ContractCallApprovedWithMint\",\"inputs\":[{\"name\":\"commandId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"sourceChain\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"sourceAddress\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"contractAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"payloadHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"symbol\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"sourceTxHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"sourceEventIndex\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ContractCallExecuted\",\"inputs\":[{\"name\":\"commandId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ContractCallWithToken\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"destinationChain\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"destinationContractAddress\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"payloadHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"payload\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"symbol\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Executed\",\"inputs\":[{\"name\":\"commandId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GovernanceTransferred\",\"inputs\":[{\"name\":\"previousGovernance\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newGovernance\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MintLimiterTransferred\",\"inputs\":[{\"name\":\"previousGovernance\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newGovernance\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorshipTransferred\",\"inputs\":[{\"name\":\"newOperatorsData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RedeemToken\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sequence\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"custodianGroupUID\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"destinationChain\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"destinationContractAddress\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"payloadHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"payload\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"symbol\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RegisterCustodianGroup\",\"inputs\":[{\"name\":\"custodianGroupId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"sequence\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"phase\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumScalarGateway.Phase\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SwitchPhase\",\"inputs\":[{\"name\":\"custodianGroupId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"sequence\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"from\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumScalarGateway.Phase\"},{\"name\":\"to\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumScalarGateway.Phase\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenDeployed\",\"inputs\":[{\"name\":\"symbol\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"tokenAddresses\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenMintLimitUpdated\",\"inputs\":[{\"name\":\"symbol\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"limit\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenSent\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"destinationChain\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"destinationAddress\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"BurnFailed\",\"inputs\":[{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"ExceedMintLimit\",\"inputs\":[{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"InvalidAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAuthModule\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidChainId\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCodeHash\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCommands\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCustodianGroupId\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidGovernance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidImplementation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidMintLimiter\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPhase\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSession\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSetMintLimitsParams\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTokenDeployer\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MintFailed\",\"inputs\":[{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"NotGovernance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializedSession\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotMintLimiter\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotProxy\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PhaseAlreadyExists\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PhaseNotChanged\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SetupFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TokenAlreadyExists\",\"inputs\":[{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"TokenContractDoesNotExist\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenDeployFailed\",\"inputs\":[{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"TokenDoesNotExist\",\"inputs\":[{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"TokenTransferFailed\",\"inputs\":[]}]",
}

// IScalarGatewayABI is the input ABI used to generate the binding from.
// Deprecated: Use IScalarGatewayMetaData.ABI instead.
var IScalarGatewayABI = IScalarGatewayMetaData.ABI

// IScalarGateway is an auto generated Go binding around an Ethereum contract.
type IScalarGateway struct {
	IScalarGatewayCaller     // Read-only binding to the contract
	IScalarGatewayTransactor // Write-only binding to the contract
	IScalarGatewayFilterer   // Log filterer for contract events
}

// IScalarGatewayCaller is an auto generated read-only Go binding around an Ethereum contract.
type IScalarGatewayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IScalarGatewayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IScalarGatewayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IScalarGatewayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IScalarGatewayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IScalarGatewaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IScalarGatewaySession struct {
	Contract     *IScalarGateway   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IScalarGatewayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IScalarGatewayCallerSession struct {
	Contract *IScalarGatewayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IScalarGatewayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IScalarGatewayTransactorSession struct {
	Contract     *IScalarGatewayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IScalarGatewayRaw is an auto generated low-level Go binding around an Ethereum contract.
type IScalarGatewayRaw struct {
	Contract *IScalarGateway // Generic contract binding to access the raw methods on
}

// IScalarGatewayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IScalarGatewayCallerRaw struct {
	Contract *IScalarGatewayCaller // Generic read-only contract binding to access the raw methods on
}

// IScalarGatewayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IScalarGatewayTransactorRaw struct {
	Contract *IScalarGatewayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIScalarGateway creates a new instance of IScalarGateway, bound to a specific deployed contract.
func NewIScalarGateway(address common.Address, backend bind.ContractBackend) (*IScalarGateway, error) {
	contract, err := bindIScalarGateway(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IScalarGateway{IScalarGatewayCaller: IScalarGatewayCaller{contract: contract}, IScalarGatewayTransactor: IScalarGatewayTransactor{contract: contract}, IScalarGatewayFilterer: IScalarGatewayFilterer{contract: contract}}, nil
}

// NewIScalarGatewayCaller creates a new read-only instance of IScalarGateway, bound to a specific deployed contract.
func NewIScalarGatewayCaller(address common.Address, caller bind.ContractCaller) (*IScalarGatewayCaller, error) {
	contract, err := bindIScalarGateway(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IScalarGatewayCaller{contract: contract}, nil
}

// NewIScalarGatewayTransactor creates a new write-only instance of IScalarGateway, bound to a specific deployed contract.
func NewIScalarGatewayTransactor(address common.Address, transactor bind.ContractTransactor) (*IScalarGatewayTransactor, error) {
	contract, err := bindIScalarGateway(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IScalarGatewayTransactor{contract: contract}, nil
}

// NewIScalarGatewayFilterer creates a new log filterer instance of IScalarGateway, bound to a specific deployed contract.
func NewIScalarGatewayFilterer(address common.Address, filterer bind.ContractFilterer) (*IScalarGatewayFilterer, error) {
	contract, err := bindIScalarGateway(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IScalarGatewayFilterer{contract: contract}, nil
}

// bindIScalarGateway binds a generic wrapper to an already deployed contract.
func bindIScalarGateway(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IScalarGatewayMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IScalarGateway *IScalarGatewayRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IScalarGateway.Contract.IScalarGatewayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IScalarGateway *IScalarGatewayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IScalarGateway.Contract.IScalarGatewayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IScalarGateway *IScalarGatewayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IScalarGateway.Contract.IScalarGatewayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IScalarGateway *IScalarGatewayCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IScalarGateway.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IScalarGateway *IScalarGatewayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IScalarGateway.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IScalarGateway *IScalarGatewayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IScalarGateway.Contract.contract.Transact(opts, method, params...)
}

// AllTokensFrozen is a free data retrieval call binding the contract method 0xaa1e1f0a.
//
// Solidity: function allTokensFrozen() pure returns(bool)
func (_IScalarGateway *IScalarGatewayCaller) AllTokensFrozen(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _IScalarGateway.contract.Call(opts, &out, "allTokensFrozen")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllTokensFrozen is a free data retrieval call binding the contract method 0xaa1e1f0a.
//
// Solidity: function allTokensFrozen() pure returns(bool)
func (_IScalarGateway *IScalarGatewaySession) AllTokensFrozen() (bool, error) {
	return _IScalarGateway.Contract.AllTokensFrozen(&_IScalarGateway.CallOpts)
}

// AllTokensFrozen is a free data retrieval call binding the contract method 0xaa1e1f0a.
//
// Solidity: function allTokensFrozen() pure returns(bool)
func (_IScalarGateway *IScalarGatewayCallerSession) AllTokensFrozen() (bool, error) {
	return _IScalarGateway.Contract.AllTokensFrozen(&_IScalarGateway.CallOpts)
}

// AuthModule is a free data retrieval call binding the contract method 0x64940c56.
//
// Solidity: function authModule() view returns(address)
func (_IScalarGateway *IScalarGatewayCaller) AuthModule(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IScalarGateway.contract.Call(opts, &out, "authModule")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AuthModule is a free data retrieval call binding the contract method 0x64940c56.
//
// Solidity: function authModule() view returns(address)
func (_IScalarGateway *IScalarGatewaySession) AuthModule() (common.Address, error) {
	return _IScalarGateway.Contract.AuthModule(&_IScalarGateway.CallOpts)
}

// AuthModule is a free data retrieval call binding the contract method 0x64940c56.
//
// Solidity: function authModule() view returns(address)
func (_IScalarGateway *IScalarGatewayCallerSession) AuthModule() (common.Address, error) {
	return _IScalarGateway.Contract.AuthModule(&_IScalarGateway.CallOpts)
}

// ContractId is a free data retrieval call binding the contract method 0x8291286c.
//
// Solidity: function contractId() pure returns(bytes32)
func (_IScalarGateway *IScalarGatewayCaller) ContractId(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IScalarGateway.contract.Call(opts, &out, "contractId")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ContractId is a free data retrieval call binding the contract method 0x8291286c.
//
// Solidity: function contractId() pure returns(bytes32)
func (_IScalarGateway *IScalarGatewaySession) ContractId() ([32]byte, error) {
	return _IScalarGateway.Contract.ContractId(&_IScalarGateway.CallOpts)
}

// ContractId is a free data retrieval call binding the contract method 0x8291286c.
//
// Solidity: function contractId() pure returns(bytes32)
func (_IScalarGateway *IScalarGatewayCallerSession) ContractId() ([32]byte, error) {
	return _IScalarGateway.Contract.ContractId(&_IScalarGateway.CallOpts)
}

// GetAddress is a free data retrieval call binding the contract method 0x21f8a721.
//
// Solidity: function getAddress(bytes32 key) view returns(address)
func (_IScalarGateway *IScalarGatewayCaller) GetAddress(opts *bind.CallOpts, key [32]byte) (common.Address, error) {
	var out []interface{}
	err := _IScalarGateway.contract.Call(opts, &out, "getAddress", key)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddress is a free data retrieval call binding the contract method 0x21f8a721.
//
// Solidity: function getAddress(bytes32 key) view returns(address)
func (_IScalarGateway *IScalarGatewaySession) GetAddress(key [32]byte) (common.Address, error) {
	return _IScalarGateway.Contract.GetAddress(&_IScalarGateway.CallOpts, key)
}

// GetAddress is a free data retrieval call binding the contract method 0x21f8a721.
//
// Solidity: function getAddress(bytes32 key) view returns(address)
func (_IScalarGateway *IScalarGatewayCallerSession) GetAddress(key [32]byte) (common.Address, error) {
	return _IScalarGateway.Contract.GetAddress(&_IScalarGateway.CallOpts, key)
}

// GetBool is a free data retrieval call binding the contract method 0x7ae1cfca.
//
// Solidity: function getBool(bytes32 key) view returns(bool)
func (_IScalarGateway *IScalarGatewayCaller) GetBool(opts *bind.CallOpts, key [32]byte) (bool, error) {
	var out []interface{}
	err := _IScalarGateway.contract.Call(opts, &out, "getBool", key)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetBool is a free data retrieval call binding the contract method 0x7ae1cfca.
//
// Solidity: function getBool(bytes32 key) view returns(bool)
func (_IScalarGateway *IScalarGatewaySession) GetBool(key [32]byte) (bool, error) {
	return _IScalarGateway.Contract.GetBool(&_IScalarGateway.CallOpts, key)
}

// GetBool is a free data retrieval call binding the contract method 0x7ae1cfca.
//
// Solidity: function getBool(bytes32 key) view returns(bool)
func (_IScalarGateway *IScalarGatewayCallerSession) GetBool(key [32]byte) (bool, error) {
	return _IScalarGateway.Contract.GetBool(&_IScalarGateway.CallOpts, key)
}

// GetBytes is a free data retrieval call binding the contract method 0xc031a180.
//
// Solidity: function getBytes(bytes32 key) view returns(bytes)
func (_IScalarGateway *IScalarGatewayCaller) GetBytes(opts *bind.CallOpts, key [32]byte) ([]byte, error) {
	var out []interface{}
	err := _IScalarGateway.contract.Call(opts, &out, "getBytes", key)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetBytes is a free data retrieval call binding the contract method 0xc031a180.
//
// Solidity: function getBytes(bytes32 key) view returns(bytes)
func (_IScalarGateway *IScalarGatewaySession) GetBytes(key [32]byte) ([]byte, error) {
	return _IScalarGateway.Contract.GetBytes(&_IScalarGateway.CallOpts, key)
}

// GetBytes is a free data retrieval call binding the contract method 0xc031a180.
//
// Solidity: function getBytes(bytes32 key) view returns(bytes)
func (_IScalarGateway *IScalarGatewayCallerSession) GetBytes(key [32]byte) ([]byte, error) {
	return _IScalarGateway.Contract.GetBytes(&_IScalarGateway.CallOpts, key)
}

// GetInt is a free data retrieval call binding the contract method 0xdc97d962.
//
// Solidity: function getInt(bytes32 key) view returns(int256)
func (_IScalarGateway *IScalarGatewayCaller) GetInt(opts *bind.CallOpts, key [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _IScalarGateway.contract.Call(opts, &out, "getInt", key)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetInt is a free data retrieval call binding the contract method 0xdc97d962.
//
// Solidity: function getInt(bytes32 key) view returns(int256)
func (_IScalarGateway *IScalarGatewaySession) GetInt(key [32]byte) (*big.Int, error) {
	return _IScalarGateway.Contract.GetInt(&_IScalarGateway.CallOpts, key)
}

// GetInt is a free data retrieval call binding the contract method 0xdc97d962.
//
// Solidity: function getInt(bytes32 key) view returns(int256)
func (_IScalarGateway *IScalarGatewayCallerSession) GetInt(key [32]byte) (*big.Int, error) {
	return _IScalarGateway.Contract.GetInt(&_IScalarGateway.CallOpts, key)
}

// GetSession is a free data retrieval call binding the contract method 0x39b240bd.
//
// Solidity: function getSession(bytes32 _custodianGroupId) view returns((uint64,uint8))
func (_IScalarGateway *IScalarGatewayCaller) GetSession(opts *bind.CallOpts, _custodianGroupId [32]byte) (ScalarGatewaySession, error) {
	var out []interface{}
	err := _IScalarGateway.contract.Call(opts, &out, "getSession", _custodianGroupId)

	if err != nil {
		return *new(ScalarGatewaySession), err
	}

	out0 := *abi.ConvertType(out[0], new(ScalarGatewaySession)).(*ScalarGatewaySession)

	return out0, err

}

// GetSession is a free data retrieval call binding the contract method 0x39b240bd.
//
// Solidity: function getSession(bytes32 _custodianGroupId) view returns((uint64,uint8))
func (_IScalarGateway *IScalarGatewaySession) GetSession(_custodianGroupId [32]byte) (ScalarGatewaySession, error) {
	return _IScalarGateway.Contract.GetSession(&_IScalarGateway.CallOpts, _custodianGroupId)
}

// GetSession is a free data retrieval call binding the contract method 0x39b240bd.
//
// Solidity: function getSession(bytes32 _custodianGroupId) view returns((uint64,uint8))
func (_IScalarGateway *IScalarGatewayCallerSession) GetSession(_custodianGroupId [32]byte) (ScalarGatewaySession, error) {
	return _IScalarGateway.Contract.GetSession(&_IScalarGateway.CallOpts, _custodianGroupId)
}

// GetString is a free data retrieval call binding the contract method 0x986e791a.
//
// Solidity: function getString(bytes32 key) view returns(string)
func (_IScalarGateway *IScalarGatewayCaller) GetString(opts *bind.CallOpts, key [32]byte) (string, error) {
	var out []interface{}
	err := _IScalarGateway.contract.Call(opts, &out, "getString", key)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetString is a free data retrieval call binding the contract method 0x986e791a.
//
// Solidity: function getString(bytes32 key) view returns(string)
func (_IScalarGateway *IScalarGatewaySession) GetString(key [32]byte) (string, error) {
	return _IScalarGateway.Contract.GetString(&_IScalarGateway.CallOpts, key)
}

// GetString is a free data retrieval call binding the contract method 0x986e791a.
//
// Solidity: function getString(bytes32 key) view returns(string)
func (_IScalarGateway *IScalarGatewayCallerSession) GetString(key [32]byte) (string, error) {
	return _IScalarGateway.Contract.GetString(&_IScalarGateway.CallOpts, key)
}

// GetUint is a free data retrieval call binding the contract method 0xbd02d0f5.
//
// Solidity: function getUint(bytes32 key) view returns(uint256)
func (_IScalarGateway *IScalarGatewayCaller) GetUint(opts *bind.CallOpts, key [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _IScalarGateway.contract.Call(opts, &out, "getUint", key)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUint is a free data retrieval call binding the contract method 0xbd02d0f5.
//
// Solidity: function getUint(bytes32 key) view returns(uint256)
func (_IScalarGateway *IScalarGatewaySession) GetUint(key [32]byte) (*big.Int, error) {
	return _IScalarGateway.Contract.GetUint(&_IScalarGateway.CallOpts, key)
}

// GetUint is a free data retrieval call binding the contract method 0xbd02d0f5.
//
// Solidity: function getUint(bytes32 key) view returns(uint256)
func (_IScalarGateway *IScalarGatewayCallerSession) GetUint(key [32]byte) (*big.Int, error) {
	return _IScalarGateway.Contract.GetUint(&_IScalarGateway.CallOpts, key)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_IScalarGateway *IScalarGatewayCaller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IScalarGateway.contract.Call(opts, &out, "governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_IScalarGateway *IScalarGatewaySession) Governance() (common.Address, error) {
	return _IScalarGateway.Contract.Governance(&_IScalarGateway.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_IScalarGateway *IScalarGatewayCallerSession) Governance() (common.Address, error) {
	return _IScalarGateway.Contract.Governance(&_IScalarGateway.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_IScalarGateway *IScalarGatewayCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IScalarGateway.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_IScalarGateway *IScalarGatewaySession) Implementation() (common.Address, error) {
	return _IScalarGateway.Contract.Implementation(&_IScalarGateway.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_IScalarGateway *IScalarGatewayCallerSession) Implementation() (common.Address, error) {
	return _IScalarGateway.Contract.Implementation(&_IScalarGateway.CallOpts)
}

// IsCommandExecuted is a free data retrieval call binding the contract method 0xd26ff210.
//
// Solidity: function isCommandExecuted(bytes32 commandId) view returns(bool)
func (_IScalarGateway *IScalarGatewayCaller) IsCommandExecuted(opts *bind.CallOpts, commandId [32]byte) (bool, error) {
	var out []interface{}
	err := _IScalarGateway.contract.Call(opts, &out, "isCommandExecuted", commandId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCommandExecuted is a free data retrieval call binding the contract method 0xd26ff210.
//
// Solidity: function isCommandExecuted(bytes32 commandId) view returns(bool)
func (_IScalarGateway *IScalarGatewaySession) IsCommandExecuted(commandId [32]byte) (bool, error) {
	return _IScalarGateway.Contract.IsCommandExecuted(&_IScalarGateway.CallOpts, commandId)
}

// IsCommandExecuted is a free data retrieval call binding the contract method 0xd26ff210.
//
// Solidity: function isCommandExecuted(bytes32 commandId) view returns(bool)
func (_IScalarGateway *IScalarGatewayCallerSession) IsCommandExecuted(commandId [32]byte) (bool, error) {
	return _IScalarGateway.Contract.IsCommandExecuted(&_IScalarGateway.CallOpts, commandId)
}

// IsContractCallAndMintApproved is a free data retrieval call binding the contract method 0xbc00c216.
//
// Solidity: function isContractCallAndMintApproved(bytes32 commandId, string sourceChain, string sourceAddress, address contractAddress, bytes32 payloadHash, string symbol, uint256 amount) view returns(bool)
func (_IScalarGateway *IScalarGatewayCaller) IsContractCallAndMintApproved(opts *bind.CallOpts, commandId [32]byte, sourceChain string, sourceAddress string, contractAddress common.Address, payloadHash [32]byte, symbol string, amount *big.Int) (bool, error) {
	var out []interface{}
	err := _IScalarGateway.contract.Call(opts, &out, "isContractCallAndMintApproved", commandId, sourceChain, sourceAddress, contractAddress, payloadHash, symbol, amount)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsContractCallAndMintApproved is a free data retrieval call binding the contract method 0xbc00c216.
//
// Solidity: function isContractCallAndMintApproved(bytes32 commandId, string sourceChain, string sourceAddress, address contractAddress, bytes32 payloadHash, string symbol, uint256 amount) view returns(bool)
func (_IScalarGateway *IScalarGatewaySession) IsContractCallAndMintApproved(commandId [32]byte, sourceChain string, sourceAddress string, contractAddress common.Address, payloadHash [32]byte, symbol string, amount *big.Int) (bool, error) {
	return _IScalarGateway.Contract.IsContractCallAndMintApproved(&_IScalarGateway.CallOpts, commandId, sourceChain, sourceAddress, contractAddress, payloadHash, symbol, amount)
}

// IsContractCallAndMintApproved is a free data retrieval call binding the contract method 0xbc00c216.
//
// Solidity: function isContractCallAndMintApproved(bytes32 commandId, string sourceChain, string sourceAddress, address contractAddress, bytes32 payloadHash, string symbol, uint256 amount) view returns(bool)
func (_IScalarGateway *IScalarGatewayCallerSession) IsContractCallAndMintApproved(commandId [32]byte, sourceChain string, sourceAddress string, contractAddress common.Address, payloadHash [32]byte, symbol string, amount *big.Int) (bool, error) {
	return _IScalarGateway.Contract.IsContractCallAndMintApproved(&_IScalarGateway.CallOpts, commandId, sourceChain, sourceAddress, contractAddress, payloadHash, symbol, amount)
}

// IsContractCallApproved is a free data retrieval call binding the contract method 0xf6a5f9f5.
//
// Solidity: function isContractCallApproved(bytes32 commandId, string sourceChain, string sourceAddress, address contractAddress, bytes32 payloadHash) view returns(bool)
func (_IScalarGateway *IScalarGatewayCaller) IsContractCallApproved(opts *bind.CallOpts, commandId [32]byte, sourceChain string, sourceAddress string, contractAddress common.Address, payloadHash [32]byte) (bool, error) {
	var out []interface{}
	err := _IScalarGateway.contract.Call(opts, &out, "isContractCallApproved", commandId, sourceChain, sourceAddress, contractAddress, payloadHash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsContractCallApproved is a free data retrieval call binding the contract method 0xf6a5f9f5.
//
// Solidity: function isContractCallApproved(bytes32 commandId, string sourceChain, string sourceAddress, address contractAddress, bytes32 payloadHash) view returns(bool)
func (_IScalarGateway *IScalarGatewaySession) IsContractCallApproved(commandId [32]byte, sourceChain string, sourceAddress string, contractAddress common.Address, payloadHash [32]byte) (bool, error) {
	return _IScalarGateway.Contract.IsContractCallApproved(&_IScalarGateway.CallOpts, commandId, sourceChain, sourceAddress, contractAddress, payloadHash)
}

// IsContractCallApproved is a free data retrieval call binding the contract method 0xf6a5f9f5.
//
// Solidity: function isContractCallApproved(bytes32 commandId, string sourceChain, string sourceAddress, address contractAddress, bytes32 payloadHash) view returns(bool)
func (_IScalarGateway *IScalarGatewayCallerSession) IsContractCallApproved(commandId [32]byte, sourceChain string, sourceAddress string, contractAddress common.Address, payloadHash [32]byte) (bool, error) {
	return _IScalarGateway.Contract.IsContractCallApproved(&_IScalarGateway.CallOpts, commandId, sourceChain, sourceAddress, contractAddress, payloadHash)
}

// MintLimiter is a free data retrieval call binding the contract method 0xc82fe87a.
//
// Solidity: function mintLimiter() view returns(address)
func (_IScalarGateway *IScalarGatewayCaller) MintLimiter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IScalarGateway.contract.Call(opts, &out, "mintLimiter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MintLimiter is a free data retrieval call binding the contract method 0xc82fe87a.
//
// Solidity: function mintLimiter() view returns(address)
func (_IScalarGateway *IScalarGatewaySession) MintLimiter() (common.Address, error) {
	return _IScalarGateway.Contract.MintLimiter(&_IScalarGateway.CallOpts)
}

// MintLimiter is a free data retrieval call binding the contract method 0xc82fe87a.
//
// Solidity: function mintLimiter() view returns(address)
func (_IScalarGateway *IScalarGatewayCallerSession) MintLimiter() (common.Address, error) {
	return _IScalarGateway.Contract.MintLimiter(&_IScalarGateway.CallOpts)
}

// Sessions is a free data retrieval call binding the contract method 0x7dbd2832.
//
// Solidity: function sessions(bytes32 ) view returns(uint64 sequence, uint8 phase)
func (_IScalarGateway *IScalarGatewayCaller) Sessions(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Sequence uint64
	Phase    uint8
}, error) {
	var out []interface{}
	err := _IScalarGateway.contract.Call(opts, &out, "sessions", arg0)

	outstruct := new(struct {
		Sequence uint64
		Phase    uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Sequence = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.Phase = *abi.ConvertType(out[1], new(uint8)).(*uint8)

	return *outstruct, err

}

// Sessions is a free data retrieval call binding the contract method 0x7dbd2832.
//
// Solidity: function sessions(bytes32 ) view returns(uint64 sequence, uint8 phase)
func (_IScalarGateway *IScalarGatewaySession) Sessions(arg0 [32]byte) (struct {
	Sequence uint64
	Phase    uint8
}, error) {
	return _IScalarGateway.Contract.Sessions(&_IScalarGateway.CallOpts, arg0)
}

// Sessions is a free data retrieval call binding the contract method 0x7dbd2832.
//
// Solidity: function sessions(bytes32 ) view returns(uint64 sequence, uint8 phase)
func (_IScalarGateway *IScalarGatewayCallerSession) Sessions(arg0 [32]byte) (struct {
	Sequence uint64
	Phase    uint8
}, error) {
	return _IScalarGateway.Contract.Sessions(&_IScalarGateway.CallOpts, arg0)
}

// TokenAddresses is a free data retrieval call binding the contract method 0x935b13f6.
//
// Solidity: function tokenAddresses(string symbol) view returns(address)
func (_IScalarGateway *IScalarGatewayCaller) TokenAddresses(opts *bind.CallOpts, symbol string) (common.Address, error) {
	var out []interface{}
	err := _IScalarGateway.contract.Call(opts, &out, "tokenAddresses", symbol)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenAddresses is a free data retrieval call binding the contract method 0x935b13f6.
//
// Solidity: function tokenAddresses(string symbol) view returns(address)
func (_IScalarGateway *IScalarGatewaySession) TokenAddresses(symbol string) (common.Address, error) {
	return _IScalarGateway.Contract.TokenAddresses(&_IScalarGateway.CallOpts, symbol)
}

// TokenAddresses is a free data retrieval call binding the contract method 0x935b13f6.
//
// Solidity: function tokenAddresses(string symbol) view returns(address)
func (_IScalarGateway *IScalarGatewayCallerSession) TokenAddresses(symbol string) (common.Address, error) {
	return _IScalarGateway.Contract.TokenAddresses(&_IScalarGateway.CallOpts, symbol)
}

// TokenCustodianGroupIds is a free data retrieval call binding the contract method 0xa8ce83ab.
//
// Solidity: function tokenCustodianGroupIds(string ) view returns(bytes32)
func (_IScalarGateway *IScalarGatewayCaller) TokenCustodianGroupIds(opts *bind.CallOpts, arg0 string) ([32]byte, error) {
	var out []interface{}
	err := _IScalarGateway.contract.Call(opts, &out, "tokenCustodianGroupIds", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TokenCustodianGroupIds is a free data retrieval call binding the contract method 0xa8ce83ab.
//
// Solidity: function tokenCustodianGroupIds(string ) view returns(bytes32)
func (_IScalarGateway *IScalarGatewaySession) TokenCustodianGroupIds(arg0 string) ([32]byte, error) {
	return _IScalarGateway.Contract.TokenCustodianGroupIds(&_IScalarGateway.CallOpts, arg0)
}

// TokenCustodianGroupIds is a free data retrieval call binding the contract method 0xa8ce83ab.
//
// Solidity: function tokenCustodianGroupIds(string ) view returns(bytes32)
func (_IScalarGateway *IScalarGatewayCallerSession) TokenCustodianGroupIds(arg0 string) ([32]byte, error) {
	return _IScalarGateway.Contract.TokenCustodianGroupIds(&_IScalarGateway.CallOpts, arg0)
}

// TokenDeployer is a free data retrieval call binding the contract method 0x2a2dae0a.
//
// Solidity: function tokenDeployer() view returns(address)
func (_IScalarGateway *IScalarGatewayCaller) TokenDeployer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IScalarGateway.contract.Call(opts, &out, "tokenDeployer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenDeployer is a free data retrieval call binding the contract method 0x2a2dae0a.
//
// Solidity: function tokenDeployer() view returns(address)
func (_IScalarGateway *IScalarGatewaySession) TokenDeployer() (common.Address, error) {
	return _IScalarGateway.Contract.TokenDeployer(&_IScalarGateway.CallOpts)
}

// TokenDeployer is a free data retrieval call binding the contract method 0x2a2dae0a.
//
// Solidity: function tokenDeployer() view returns(address)
func (_IScalarGateway *IScalarGatewayCallerSession) TokenDeployer() (common.Address, error) {
	return _IScalarGateway.Contract.TokenDeployer(&_IScalarGateway.CallOpts)
}

// TokenFrozen is a free data retrieval call binding the contract method 0x7b1b769e.
//
// Solidity: function tokenFrozen(string ) pure returns(bool)
func (_IScalarGateway *IScalarGatewayCaller) TokenFrozen(opts *bind.CallOpts, arg0 string) (bool, error) {
	var out []interface{}
	err := _IScalarGateway.contract.Call(opts, &out, "tokenFrozen", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TokenFrozen is a free data retrieval call binding the contract method 0x7b1b769e.
//
// Solidity: function tokenFrozen(string ) pure returns(bool)
func (_IScalarGateway *IScalarGatewaySession) TokenFrozen(arg0 string) (bool, error) {
	return _IScalarGateway.Contract.TokenFrozen(&_IScalarGateway.CallOpts, arg0)
}

// TokenFrozen is a free data retrieval call binding the contract method 0x7b1b769e.
//
// Solidity: function tokenFrozen(string ) pure returns(bool)
func (_IScalarGateway *IScalarGatewayCallerSession) TokenFrozen(arg0 string) (bool, error) {
	return _IScalarGateway.Contract.TokenFrozen(&_IScalarGateway.CallOpts, arg0)
}

// TokenMintAmount is a free data retrieval call binding the contract method 0xcec7b359.
//
// Solidity: function tokenMintAmount(string symbol) view returns(uint256)
func (_IScalarGateway *IScalarGatewayCaller) TokenMintAmount(opts *bind.CallOpts, symbol string) (*big.Int, error) {
	var out []interface{}
	err := _IScalarGateway.contract.Call(opts, &out, "tokenMintAmount", symbol)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenMintAmount is a free data retrieval call binding the contract method 0xcec7b359.
//
// Solidity: function tokenMintAmount(string symbol) view returns(uint256)
func (_IScalarGateway *IScalarGatewaySession) TokenMintAmount(symbol string) (*big.Int, error) {
	return _IScalarGateway.Contract.TokenMintAmount(&_IScalarGateway.CallOpts, symbol)
}

// TokenMintAmount is a free data retrieval call binding the contract method 0xcec7b359.
//
// Solidity: function tokenMintAmount(string symbol) view returns(uint256)
func (_IScalarGateway *IScalarGatewayCallerSession) TokenMintAmount(symbol string) (*big.Int, error) {
	return _IScalarGateway.Contract.TokenMintAmount(&_IScalarGateway.CallOpts, symbol)
}

// TokenMintLimit is a free data retrieval call binding the contract method 0x269eb65e.
//
// Solidity: function tokenMintLimit(string symbol) view returns(uint256)
func (_IScalarGateway *IScalarGatewayCaller) TokenMintLimit(opts *bind.CallOpts, symbol string) (*big.Int, error) {
	var out []interface{}
	err := _IScalarGateway.contract.Call(opts, &out, "tokenMintLimit", symbol)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenMintLimit is a free data retrieval call binding the contract method 0x269eb65e.
//
// Solidity: function tokenMintLimit(string symbol) view returns(uint256)
func (_IScalarGateway *IScalarGatewaySession) TokenMintLimit(symbol string) (*big.Int, error) {
	return _IScalarGateway.Contract.TokenMintLimit(&_IScalarGateway.CallOpts, symbol)
}

// TokenMintLimit is a free data retrieval call binding the contract method 0x269eb65e.
//
// Solidity: function tokenMintLimit(string symbol) view returns(uint256)
func (_IScalarGateway *IScalarGatewayCallerSession) TokenMintLimit(symbol string) (*big.Int, error) {
	return _IScalarGateway.Contract.TokenMintLimit(&_IScalarGateway.CallOpts, symbol)
}

// ApproveContractCall is a paid mutator transaction binding the contract method 0xfbe0a31b.
//
// Solidity: function approveContractCall(bytes params, bytes32 commandId) returns()
func (_IScalarGateway *IScalarGatewayTransactor) ApproveContractCall(opts *bind.TransactOpts, params []byte, commandId [32]byte) (*types.Transaction, error) {
	return _IScalarGateway.contract.Transact(opts, "approveContractCall", params, commandId)
}

// ApproveContractCall is a paid mutator transaction binding the contract method 0xfbe0a31b.
//
// Solidity: function approveContractCall(bytes params, bytes32 commandId) returns()
func (_IScalarGateway *IScalarGatewaySession) ApproveContractCall(params []byte, commandId [32]byte) (*types.Transaction, error) {
	return _IScalarGateway.Contract.ApproveContractCall(&_IScalarGateway.TransactOpts, params, commandId)
}

// ApproveContractCall is a paid mutator transaction binding the contract method 0xfbe0a31b.
//
// Solidity: function approveContractCall(bytes params, bytes32 commandId) returns()
func (_IScalarGateway *IScalarGatewayTransactorSession) ApproveContractCall(params []byte, commandId [32]byte) (*types.Transaction, error) {
	return _IScalarGateway.Contract.ApproveContractCall(&_IScalarGateway.TransactOpts, params, commandId)
}

// ApproveContractCallWithMint is a paid mutator transaction binding the contract method 0x585a9fd4.
//
// Solidity: function approveContractCallWithMint(bytes params, bytes32 commandId) returns()
func (_IScalarGateway *IScalarGatewayTransactor) ApproveContractCallWithMint(opts *bind.TransactOpts, params []byte, commandId [32]byte) (*types.Transaction, error) {
	return _IScalarGateway.contract.Transact(opts, "approveContractCallWithMint", params, commandId)
}

// ApproveContractCallWithMint is a paid mutator transaction binding the contract method 0x585a9fd4.
//
// Solidity: function approveContractCallWithMint(bytes params, bytes32 commandId) returns()
func (_IScalarGateway *IScalarGatewaySession) ApproveContractCallWithMint(params []byte, commandId [32]byte) (*types.Transaction, error) {
	return _IScalarGateway.Contract.ApproveContractCallWithMint(&_IScalarGateway.TransactOpts, params, commandId)
}

// ApproveContractCallWithMint is a paid mutator transaction binding the contract method 0x585a9fd4.
//
// Solidity: function approveContractCallWithMint(bytes params, bytes32 commandId) returns()
func (_IScalarGateway *IScalarGatewayTransactorSession) ApproveContractCallWithMint(params []byte, commandId [32]byte) (*types.Transaction, error) {
	return _IScalarGateway.Contract.ApproveContractCallWithMint(&_IScalarGateway.TransactOpts, params, commandId)
}

// BurnToken is a paid mutator transaction binding the contract method 0x4656ae2e.
//
// Solidity: function burnToken(bytes params, bytes32 ) returns()
func (_IScalarGateway *IScalarGatewayTransactor) BurnToken(opts *bind.TransactOpts, params []byte, arg1 [32]byte) (*types.Transaction, error) {
	return _IScalarGateway.contract.Transact(opts, "burnToken", params, arg1)
}

// BurnToken is a paid mutator transaction binding the contract method 0x4656ae2e.
//
// Solidity: function burnToken(bytes params, bytes32 ) returns()
func (_IScalarGateway *IScalarGatewaySession) BurnToken(params []byte, arg1 [32]byte) (*types.Transaction, error) {
	return _IScalarGateway.Contract.BurnToken(&_IScalarGateway.TransactOpts, params, arg1)
}

// BurnToken is a paid mutator transaction binding the contract method 0x4656ae2e.
//
// Solidity: function burnToken(bytes params, bytes32 ) returns()
func (_IScalarGateway *IScalarGatewayTransactorSession) BurnToken(params []byte, arg1 [32]byte) (*types.Transaction, error) {
	return _IScalarGateway.Contract.BurnToken(&_IScalarGateway.TransactOpts, params, arg1)
}

// CallContract is a paid mutator transaction binding the contract method 0x1c92115f.
//
// Solidity: function callContract(string destinationChain, string destinationContractAddress, bytes payload) returns()
func (_IScalarGateway *IScalarGatewayTransactor) CallContract(opts *bind.TransactOpts, destinationChain string, destinationContractAddress string, payload []byte) (*types.Transaction, error) {
	return _IScalarGateway.contract.Transact(opts, "callContract", destinationChain, destinationContractAddress, payload)
}

// CallContract is a paid mutator transaction binding the contract method 0x1c92115f.
//
// Solidity: function callContract(string destinationChain, string destinationContractAddress, bytes payload) returns()
func (_IScalarGateway *IScalarGatewaySession) CallContract(destinationChain string, destinationContractAddress string, payload []byte) (*types.Transaction, error) {
	return _IScalarGateway.Contract.CallContract(&_IScalarGateway.TransactOpts, destinationChain, destinationContractAddress, payload)
}

// CallContract is a paid mutator transaction binding the contract method 0x1c92115f.
//
// Solidity: function callContract(string destinationChain, string destinationContractAddress, bytes payload) returns()
func (_IScalarGateway *IScalarGatewayTransactorSession) CallContract(destinationChain string, destinationContractAddress string, payload []byte) (*types.Transaction, error) {
	return _IScalarGateway.Contract.CallContract(&_IScalarGateway.TransactOpts, destinationChain, destinationContractAddress, payload)
}

// CallContractWithToken is a paid mutator transaction binding the contract method 0xb5417084.
//
// Solidity: function callContractWithToken(string destinationChain, string destinationContractAddress, bytes payload, string symbol, uint256 amount) returns()
func (_IScalarGateway *IScalarGatewayTransactor) CallContractWithToken(opts *bind.TransactOpts, destinationChain string, destinationContractAddress string, payload []byte, symbol string, amount *big.Int) (*types.Transaction, error) {
	return _IScalarGateway.contract.Transact(opts, "callContractWithToken", destinationChain, destinationContractAddress, payload, symbol, amount)
}

// CallContractWithToken is a paid mutator transaction binding the contract method 0xb5417084.
//
// Solidity: function callContractWithToken(string destinationChain, string destinationContractAddress, bytes payload, string symbol, uint256 amount) returns()
func (_IScalarGateway *IScalarGatewaySession) CallContractWithToken(destinationChain string, destinationContractAddress string, payload []byte, symbol string, amount *big.Int) (*types.Transaction, error) {
	return _IScalarGateway.Contract.CallContractWithToken(&_IScalarGateway.TransactOpts, destinationChain, destinationContractAddress, payload, symbol, amount)
}

// CallContractWithToken is a paid mutator transaction binding the contract method 0xb5417084.
//
// Solidity: function callContractWithToken(string destinationChain, string destinationContractAddress, bytes payload, string symbol, uint256 amount) returns()
func (_IScalarGateway *IScalarGatewayTransactorSession) CallContractWithToken(destinationChain string, destinationContractAddress string, payload []byte, symbol string, amount *big.Int) (*types.Transaction, error) {
	return _IScalarGateway.Contract.CallContractWithToken(&_IScalarGateway.TransactOpts, destinationChain, destinationContractAddress, payload, symbol, amount)
}

// DeployToken is a paid mutator transaction binding the contract method 0x886a625d.
//
// Solidity: function deployToken(bytes params, bytes32 ) returns()
func (_IScalarGateway *IScalarGatewayTransactor) DeployToken(opts *bind.TransactOpts, params []byte, arg1 [32]byte) (*types.Transaction, error) {
	return _IScalarGateway.contract.Transact(opts, "deployToken", params, arg1)
}

// DeployToken is a paid mutator transaction binding the contract method 0x886a625d.
//
// Solidity: function deployToken(bytes params, bytes32 ) returns()
func (_IScalarGateway *IScalarGatewaySession) DeployToken(params []byte, arg1 [32]byte) (*types.Transaction, error) {
	return _IScalarGateway.Contract.DeployToken(&_IScalarGateway.TransactOpts, params, arg1)
}

// DeployToken is a paid mutator transaction binding the contract method 0x886a625d.
//
// Solidity: function deployToken(bytes params, bytes32 ) returns()
func (_IScalarGateway *IScalarGatewayTransactorSession) DeployToken(params []byte, arg1 [32]byte) (*types.Transaction, error) {
	return _IScalarGateway.Contract.DeployToken(&_IScalarGateway.TransactOpts, params, arg1)
}

// Execute is a paid mutator transaction binding the contract method 0x09c5eabe.
//
// Solidity: function execute(bytes input) returns()
func (_IScalarGateway *IScalarGatewayTransactor) Execute(opts *bind.TransactOpts, input []byte) (*types.Transaction, error) {
	return _IScalarGateway.contract.Transact(opts, "execute", input)
}

// Execute is a paid mutator transaction binding the contract method 0x09c5eabe.
//
// Solidity: function execute(bytes input) returns()
func (_IScalarGateway *IScalarGatewaySession) Execute(input []byte) (*types.Transaction, error) {
	return _IScalarGateway.Contract.Execute(&_IScalarGateway.TransactOpts, input)
}

// Execute is a paid mutator transaction binding the contract method 0x09c5eabe.
//
// Solidity: function execute(bytes input) returns()
func (_IScalarGateway *IScalarGatewayTransactorSession) Execute(input []byte) (*types.Transaction, error) {
	return _IScalarGateway.Contract.Execute(&_IScalarGateway.TransactOpts, input)
}

// MintToken is a paid mutator transaction binding the contract method 0x146e2d78.
//
// Solidity: function mintToken(bytes params, bytes32 ) returns()
func (_IScalarGateway *IScalarGatewayTransactor) MintToken(opts *bind.TransactOpts, params []byte, arg1 [32]byte) (*types.Transaction, error) {
	return _IScalarGateway.contract.Transact(opts, "mintToken", params, arg1)
}

// MintToken is a paid mutator transaction binding the contract method 0x146e2d78.
//
// Solidity: function mintToken(bytes params, bytes32 ) returns()
func (_IScalarGateway *IScalarGatewaySession) MintToken(params []byte, arg1 [32]byte) (*types.Transaction, error) {
	return _IScalarGateway.Contract.MintToken(&_IScalarGateway.TransactOpts, params, arg1)
}

// MintToken is a paid mutator transaction binding the contract method 0x146e2d78.
//
// Solidity: function mintToken(bytes params, bytes32 ) returns()
func (_IScalarGateway *IScalarGatewayTransactorSession) MintToken(params []byte, arg1 [32]byte) (*types.Transaction, error) {
	return _IScalarGateway.Contract.MintToken(&_IScalarGateway.TransactOpts, params, arg1)
}

// RedeemToken is a paid mutator transaction binding the contract method 0x8d48b75e.
//
// Solidity: function redeemToken(bytes params, address account) returns()
func (_IScalarGateway *IScalarGatewayTransactor) RedeemToken(opts *bind.TransactOpts, params []byte, account common.Address) (*types.Transaction, error) {
	return _IScalarGateway.contract.Transact(opts, "redeemToken", params, account)
}

// RedeemToken is a paid mutator transaction binding the contract method 0x8d48b75e.
//
// Solidity: function redeemToken(bytes params, address account) returns()
func (_IScalarGateway *IScalarGatewaySession) RedeemToken(params []byte, account common.Address) (*types.Transaction, error) {
	return _IScalarGateway.Contract.RedeemToken(&_IScalarGateway.TransactOpts, params, account)
}

// RedeemToken is a paid mutator transaction binding the contract method 0x8d48b75e.
//
// Solidity: function redeemToken(bytes params, address account) returns()
func (_IScalarGateway *IScalarGatewayTransactorSession) RedeemToken(params []byte, account common.Address) (*types.Transaction, error) {
	return _IScalarGateway.Contract.RedeemToken(&_IScalarGateway.TransactOpts, params, account)
}

// RegisterCustodianGroup is a paid mutator transaction binding the contract method 0xec3bfb35.
//
// Solidity: function registerCustodianGroup(bytes params, bytes32 ) returns()
func (_IScalarGateway *IScalarGatewayTransactor) RegisterCustodianGroup(opts *bind.TransactOpts, params []byte, arg1 [32]byte) (*types.Transaction, error) {
	return _IScalarGateway.contract.Transact(opts, "registerCustodianGroup", params, arg1)
}

// RegisterCustodianGroup is a paid mutator transaction binding the contract method 0xec3bfb35.
//
// Solidity: function registerCustodianGroup(bytes params, bytes32 ) returns()
func (_IScalarGateway *IScalarGatewaySession) RegisterCustodianGroup(params []byte, arg1 [32]byte) (*types.Transaction, error) {
	return _IScalarGateway.Contract.RegisterCustodianGroup(&_IScalarGateway.TransactOpts, params, arg1)
}

// RegisterCustodianGroup is a paid mutator transaction binding the contract method 0xec3bfb35.
//
// Solidity: function registerCustodianGroup(bytes params, bytes32 ) returns()
func (_IScalarGateway *IScalarGatewayTransactorSession) RegisterCustodianGroup(params []byte, arg1 [32]byte) (*types.Transaction, error) {
	return _IScalarGateway.Contract.RegisterCustodianGroup(&_IScalarGateway.TransactOpts, params, arg1)
}

// SendToken is a paid mutator transaction binding the contract method 0x26ef699d.
//
// Solidity: function sendToken(string destinationChain, string destinationAddress, string symbol, uint256 amount) returns()
func (_IScalarGateway *IScalarGatewayTransactor) SendToken(opts *bind.TransactOpts, destinationChain string, destinationAddress string, symbol string, amount *big.Int) (*types.Transaction, error) {
	return _IScalarGateway.contract.Transact(opts, "sendToken", destinationChain, destinationAddress, symbol, amount)
}

// SendToken is a paid mutator transaction binding the contract method 0x26ef699d.
//
// Solidity: function sendToken(string destinationChain, string destinationAddress, string symbol, uint256 amount) returns()
func (_IScalarGateway *IScalarGatewaySession) SendToken(destinationChain string, destinationAddress string, symbol string, amount *big.Int) (*types.Transaction, error) {
	return _IScalarGateway.Contract.SendToken(&_IScalarGateway.TransactOpts, destinationChain, destinationAddress, symbol, amount)
}

// SendToken is a paid mutator transaction binding the contract method 0x26ef699d.
//
// Solidity: function sendToken(string destinationChain, string destinationAddress, string symbol, uint256 amount) returns()
func (_IScalarGateway *IScalarGatewayTransactorSession) SendToken(destinationChain string, destinationAddress string, symbol string, amount *big.Int) (*types.Transaction, error) {
	return _IScalarGateway.Contract.SendToken(&_IScalarGateway.TransactOpts, destinationChain, destinationAddress, symbol, amount)
}

// SetTokenMintLimits is a paid mutator transaction binding the contract method 0x67ace8eb.
//
// Solidity: function setTokenMintLimits(string[] symbols, uint256[] limits) returns()
func (_IScalarGateway *IScalarGatewayTransactor) SetTokenMintLimits(opts *bind.TransactOpts, symbols []string, limits []*big.Int) (*types.Transaction, error) {
	return _IScalarGateway.contract.Transact(opts, "setTokenMintLimits", symbols, limits)
}

// SetTokenMintLimits is a paid mutator transaction binding the contract method 0x67ace8eb.
//
// Solidity: function setTokenMintLimits(string[] symbols, uint256[] limits) returns()
func (_IScalarGateway *IScalarGatewaySession) SetTokenMintLimits(symbols []string, limits []*big.Int) (*types.Transaction, error) {
	return _IScalarGateway.Contract.SetTokenMintLimits(&_IScalarGateway.TransactOpts, symbols, limits)
}

// SetTokenMintLimits is a paid mutator transaction binding the contract method 0x67ace8eb.
//
// Solidity: function setTokenMintLimits(string[] symbols, uint256[] limits) returns()
func (_IScalarGateway *IScalarGatewayTransactorSession) SetTokenMintLimits(symbols []string, limits []*big.Int) (*types.Transaction, error) {
	return _IScalarGateway.Contract.SetTokenMintLimits(&_IScalarGateway.TransactOpts, symbols, limits)
}

// Setup is a paid mutator transaction binding the contract method 0x9ded06df.
//
// Solidity: function setup(bytes params) returns()
func (_IScalarGateway *IScalarGatewayTransactor) Setup(opts *bind.TransactOpts, params []byte) (*types.Transaction, error) {
	return _IScalarGateway.contract.Transact(opts, "setup", params)
}

// Setup is a paid mutator transaction binding the contract method 0x9ded06df.
//
// Solidity: function setup(bytes params) returns()
func (_IScalarGateway *IScalarGatewaySession) Setup(params []byte) (*types.Transaction, error) {
	return _IScalarGateway.Contract.Setup(&_IScalarGateway.TransactOpts, params)
}

// Setup is a paid mutator transaction binding the contract method 0x9ded06df.
//
// Solidity: function setup(bytes params) returns()
func (_IScalarGateway *IScalarGatewayTransactorSession) Setup(params []byte) (*types.Transaction, error) {
	return _IScalarGateway.Contract.Setup(&_IScalarGateway.TransactOpts, params)
}

// SwitchPhase is a paid mutator transaction binding the contract method 0x13616aee.
//
// Solidity: function switchPhase(bytes params, bytes32 ) returns()
func (_IScalarGateway *IScalarGatewayTransactor) SwitchPhase(opts *bind.TransactOpts, params []byte, arg1 [32]byte) (*types.Transaction, error) {
	return _IScalarGateway.contract.Transact(opts, "switchPhase", params, arg1)
}

// SwitchPhase is a paid mutator transaction binding the contract method 0x13616aee.
//
// Solidity: function switchPhase(bytes params, bytes32 ) returns()
func (_IScalarGateway *IScalarGatewaySession) SwitchPhase(params []byte, arg1 [32]byte) (*types.Transaction, error) {
	return _IScalarGateway.Contract.SwitchPhase(&_IScalarGateway.TransactOpts, params, arg1)
}

// SwitchPhase is a paid mutator transaction binding the contract method 0x13616aee.
//
// Solidity: function switchPhase(bytes params, bytes32 ) returns()
func (_IScalarGateway *IScalarGatewayTransactorSession) SwitchPhase(params []byte, arg1 [32]byte) (*types.Transaction, error) {
	return _IScalarGateway.Contract.SwitchPhase(&_IScalarGateway.TransactOpts, params, arg1)
}

// TransferGovernance is a paid mutator transaction binding the contract method 0xd38bfff4.
//
// Solidity: function transferGovernance(address newGovernance) returns()
func (_IScalarGateway *IScalarGatewayTransactor) TransferGovernance(opts *bind.TransactOpts, newGovernance common.Address) (*types.Transaction, error) {
	return _IScalarGateway.contract.Transact(opts, "transferGovernance", newGovernance)
}

// TransferGovernance is a paid mutator transaction binding the contract method 0xd38bfff4.
//
// Solidity: function transferGovernance(address newGovernance) returns()
func (_IScalarGateway *IScalarGatewaySession) TransferGovernance(newGovernance common.Address) (*types.Transaction, error) {
	return _IScalarGateway.Contract.TransferGovernance(&_IScalarGateway.TransactOpts, newGovernance)
}

// TransferGovernance is a paid mutator transaction binding the contract method 0xd38bfff4.
//
// Solidity: function transferGovernance(address newGovernance) returns()
func (_IScalarGateway *IScalarGatewayTransactorSession) TransferGovernance(newGovernance common.Address) (*types.Transaction, error) {
	return _IScalarGateway.Contract.TransferGovernance(&_IScalarGateway.TransactOpts, newGovernance)
}

// TransferMintLimiter is a paid mutator transaction binding the contract method 0x41d8f26b.
//
// Solidity: function transferMintLimiter(address newMintLimiter) returns()
func (_IScalarGateway *IScalarGatewayTransactor) TransferMintLimiter(opts *bind.TransactOpts, newMintLimiter common.Address) (*types.Transaction, error) {
	return _IScalarGateway.contract.Transact(opts, "transferMintLimiter", newMintLimiter)
}

// TransferMintLimiter is a paid mutator transaction binding the contract method 0x41d8f26b.
//
// Solidity: function transferMintLimiter(address newMintLimiter) returns()
func (_IScalarGateway *IScalarGatewaySession) TransferMintLimiter(newMintLimiter common.Address) (*types.Transaction, error) {
	return _IScalarGateway.Contract.TransferMintLimiter(&_IScalarGateway.TransactOpts, newMintLimiter)
}

// TransferMintLimiter is a paid mutator transaction binding the contract method 0x41d8f26b.
//
// Solidity: function transferMintLimiter(address newMintLimiter) returns()
func (_IScalarGateway *IScalarGatewayTransactorSession) TransferMintLimiter(newMintLimiter common.Address) (*types.Transaction, error) {
	return _IScalarGateway.Contract.TransferMintLimiter(&_IScalarGateway.TransactOpts, newMintLimiter)
}

// TransferOperatorship is a paid mutator transaction binding the contract method 0x97b87ba6.
//
// Solidity: function transferOperatorship(bytes newOperatorsData, bytes32 ) returns()
func (_IScalarGateway *IScalarGatewayTransactor) TransferOperatorship(opts *bind.TransactOpts, newOperatorsData []byte, arg1 [32]byte) (*types.Transaction, error) {
	return _IScalarGateway.contract.Transact(opts, "transferOperatorship", newOperatorsData, arg1)
}

// TransferOperatorship is a paid mutator transaction binding the contract method 0x97b87ba6.
//
// Solidity: function transferOperatorship(bytes newOperatorsData, bytes32 ) returns()
func (_IScalarGateway *IScalarGatewaySession) TransferOperatorship(newOperatorsData []byte, arg1 [32]byte) (*types.Transaction, error) {
	return _IScalarGateway.Contract.TransferOperatorship(&_IScalarGateway.TransactOpts, newOperatorsData, arg1)
}

// TransferOperatorship is a paid mutator transaction binding the contract method 0x97b87ba6.
//
// Solidity: function transferOperatorship(bytes newOperatorsData, bytes32 ) returns()
func (_IScalarGateway *IScalarGatewayTransactorSession) TransferOperatorship(newOperatorsData []byte, arg1 [32]byte) (*types.Transaction, error) {
	return _IScalarGateway.Contract.TransferOperatorship(&_IScalarGateway.TransactOpts, newOperatorsData, arg1)
}

// Upgrade is a paid mutator transaction binding the contract method 0xa3499c73.
//
// Solidity: function upgrade(address newImplementation, bytes32 newImplementationCodeHash, bytes setupParams) returns()
func (_IScalarGateway *IScalarGatewayTransactor) Upgrade(opts *bind.TransactOpts, newImplementation common.Address, newImplementationCodeHash [32]byte, setupParams []byte) (*types.Transaction, error) {
	return _IScalarGateway.contract.Transact(opts, "upgrade", newImplementation, newImplementationCodeHash, setupParams)
}

// Upgrade is a paid mutator transaction binding the contract method 0xa3499c73.
//
// Solidity: function upgrade(address newImplementation, bytes32 newImplementationCodeHash, bytes setupParams) returns()
func (_IScalarGateway *IScalarGatewaySession) Upgrade(newImplementation common.Address, newImplementationCodeHash [32]byte, setupParams []byte) (*types.Transaction, error) {
	return _IScalarGateway.Contract.Upgrade(&_IScalarGateway.TransactOpts, newImplementation, newImplementationCodeHash, setupParams)
}

// Upgrade is a paid mutator transaction binding the contract method 0xa3499c73.
//
// Solidity: function upgrade(address newImplementation, bytes32 newImplementationCodeHash, bytes setupParams) returns()
func (_IScalarGateway *IScalarGatewayTransactorSession) Upgrade(newImplementation common.Address, newImplementationCodeHash [32]byte, setupParams []byte) (*types.Transaction, error) {
	return _IScalarGateway.Contract.Upgrade(&_IScalarGateway.TransactOpts, newImplementation, newImplementationCodeHash, setupParams)
}

// ValidateContractCall is a paid mutator transaction binding the contract method 0x5f6970c3.
//
// Solidity: function validateContractCall(bytes32 commandId, string sourceChain, string sourceAddress, bytes32 payloadHash) returns(bool valid)
func (_IScalarGateway *IScalarGatewayTransactor) ValidateContractCall(opts *bind.TransactOpts, commandId [32]byte, sourceChain string, sourceAddress string, payloadHash [32]byte) (*types.Transaction, error) {
	return _IScalarGateway.contract.Transact(opts, "validateContractCall", commandId, sourceChain, sourceAddress, payloadHash)
}

// ValidateContractCall is a paid mutator transaction binding the contract method 0x5f6970c3.
//
// Solidity: function validateContractCall(bytes32 commandId, string sourceChain, string sourceAddress, bytes32 payloadHash) returns(bool valid)
func (_IScalarGateway *IScalarGatewaySession) ValidateContractCall(commandId [32]byte, sourceChain string, sourceAddress string, payloadHash [32]byte) (*types.Transaction, error) {
	return _IScalarGateway.Contract.ValidateContractCall(&_IScalarGateway.TransactOpts, commandId, sourceChain, sourceAddress, payloadHash)
}

// ValidateContractCall is a paid mutator transaction binding the contract method 0x5f6970c3.
//
// Solidity: function validateContractCall(bytes32 commandId, string sourceChain, string sourceAddress, bytes32 payloadHash) returns(bool valid)
func (_IScalarGateway *IScalarGatewayTransactorSession) ValidateContractCall(commandId [32]byte, sourceChain string, sourceAddress string, payloadHash [32]byte) (*types.Transaction, error) {
	return _IScalarGateway.Contract.ValidateContractCall(&_IScalarGateway.TransactOpts, commandId, sourceChain, sourceAddress, payloadHash)
}

// ValidateContractCallAndMint is a paid mutator transaction binding the contract method 0x1876eed9.
//
// Solidity: function validateContractCallAndMint(bytes32 commandId, string sourceChain, string sourceAddress, bytes32 payloadHash, string symbol, uint256 amount) returns(bool valid)
func (_IScalarGateway *IScalarGatewayTransactor) ValidateContractCallAndMint(opts *bind.TransactOpts, commandId [32]byte, sourceChain string, sourceAddress string, payloadHash [32]byte, symbol string, amount *big.Int) (*types.Transaction, error) {
	return _IScalarGateway.contract.Transact(opts, "validateContractCallAndMint", commandId, sourceChain, sourceAddress, payloadHash, symbol, amount)
}

// ValidateContractCallAndMint is a paid mutator transaction binding the contract method 0x1876eed9.
//
// Solidity: function validateContractCallAndMint(bytes32 commandId, string sourceChain, string sourceAddress, bytes32 payloadHash, string symbol, uint256 amount) returns(bool valid)
func (_IScalarGateway *IScalarGatewaySession) ValidateContractCallAndMint(commandId [32]byte, sourceChain string, sourceAddress string, payloadHash [32]byte, symbol string, amount *big.Int) (*types.Transaction, error) {
	return _IScalarGateway.Contract.ValidateContractCallAndMint(&_IScalarGateway.TransactOpts, commandId, sourceChain, sourceAddress, payloadHash, symbol, amount)
}

// ValidateContractCallAndMint is a paid mutator transaction binding the contract method 0x1876eed9.
//
// Solidity: function validateContractCallAndMint(bytes32 commandId, string sourceChain, string sourceAddress, bytes32 payloadHash, string symbol, uint256 amount) returns(bool valid)
func (_IScalarGateway *IScalarGatewayTransactorSession) ValidateContractCallAndMint(commandId [32]byte, sourceChain string, sourceAddress string, payloadHash [32]byte, symbol string, amount *big.Int) (*types.Transaction, error) {
	return _IScalarGateway.Contract.ValidateContractCallAndMint(&_IScalarGateway.TransactOpts, commandId, sourceChain, sourceAddress, payloadHash, symbol, amount)
}

// IScalarGatewayContractCallIterator is returned from FilterContractCall and is used to iterate over the raw logs and unpacked data for ContractCall events raised by the IScalarGateway contract.
type IScalarGatewayContractCallIterator struct {
	Event *IScalarGatewayContractCall // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IScalarGatewayContractCallIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IScalarGatewayContractCall)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IScalarGatewayContractCall)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IScalarGatewayContractCallIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IScalarGatewayContractCallIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IScalarGatewayContractCall represents a ContractCall event raised by the IScalarGateway contract.
type IScalarGatewayContractCall struct {
	Sender                     common.Address
	DestinationChain           string
	DestinationContractAddress string
	PayloadHash                [32]byte
	Payload                    []byte
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterContractCall is a free log retrieval operation binding the contract event 0x30ae6cc78c27e651745bf2ad08a11de83910ac1e347a52f7ac898c0fbef94dae.
//
// Solidity: event ContractCall(address indexed sender, string destinationChain, string destinationContractAddress, bytes32 indexed payloadHash, bytes payload)
func (_IScalarGateway *IScalarGatewayFilterer) FilterContractCall(opts *bind.FilterOpts, sender []common.Address, payloadHash [][32]byte) (*IScalarGatewayContractCallIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var payloadHashRule []interface{}
	for _, payloadHashItem := range payloadHash {
		payloadHashRule = append(payloadHashRule, payloadHashItem)
	}

	logs, sub, err := _IScalarGateway.contract.FilterLogs(opts, "ContractCall", senderRule, payloadHashRule)
	if err != nil {
		return nil, err
	}
	return &IScalarGatewayContractCallIterator{contract: _IScalarGateway.contract, event: "ContractCall", logs: logs, sub: sub}, nil
}

// WatchContractCall is a free log subscription operation binding the contract event 0x30ae6cc78c27e651745bf2ad08a11de83910ac1e347a52f7ac898c0fbef94dae.
//
// Solidity: event ContractCall(address indexed sender, string destinationChain, string destinationContractAddress, bytes32 indexed payloadHash, bytes payload)
func (_IScalarGateway *IScalarGatewayFilterer) WatchContractCall(opts *bind.WatchOpts, sink chan<- *IScalarGatewayContractCall, sender []common.Address, payloadHash [][32]byte) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var payloadHashRule []interface{}
	for _, payloadHashItem := range payloadHash {
		payloadHashRule = append(payloadHashRule, payloadHashItem)
	}

	logs, sub, err := _IScalarGateway.contract.WatchLogs(opts, "ContractCall", senderRule, payloadHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IScalarGatewayContractCall)
				if err := _IScalarGateway.contract.UnpackLog(event, "ContractCall", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseContractCall is a log parse operation binding the contract event 0x30ae6cc78c27e651745bf2ad08a11de83910ac1e347a52f7ac898c0fbef94dae.
//
// Solidity: event ContractCall(address indexed sender, string destinationChain, string destinationContractAddress, bytes32 indexed payloadHash, bytes payload)
func (_IScalarGateway *IScalarGatewayFilterer) ParseContractCall(log types.Log) (*IScalarGatewayContractCall, error) {
	event := new(IScalarGatewayContractCall)
	if err := _IScalarGateway.contract.UnpackLog(event, "ContractCall", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IScalarGatewayContractCallApprovedIterator is returned from FilterContractCallApproved and is used to iterate over the raw logs and unpacked data for ContractCallApproved events raised by the IScalarGateway contract.
type IScalarGatewayContractCallApprovedIterator struct {
	Event *IScalarGatewayContractCallApproved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IScalarGatewayContractCallApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IScalarGatewayContractCallApproved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IScalarGatewayContractCallApproved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IScalarGatewayContractCallApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IScalarGatewayContractCallApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IScalarGatewayContractCallApproved represents a ContractCallApproved event raised by the IScalarGateway contract.
type IScalarGatewayContractCallApproved struct {
	CommandId        [32]byte
	SourceChain      string
	SourceAddress    string
	ContractAddress  common.Address
	PayloadHash      [32]byte
	SourceTxHash     [32]byte
	SourceEventIndex *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterContractCallApproved is a free log retrieval operation binding the contract event 0x44e4f8f6bd682c5a3aeba93601ab07cb4d1f21b2aab1ae4880d9577919309aa4.
//
// Solidity: event ContractCallApproved(bytes32 indexed commandId, string sourceChain, string sourceAddress, address indexed contractAddress, bytes32 indexed payloadHash, bytes32 sourceTxHash, uint256 sourceEventIndex)
func (_IScalarGateway *IScalarGatewayFilterer) FilterContractCallApproved(opts *bind.FilterOpts, commandId [][32]byte, contractAddress []common.Address, payloadHash [][32]byte) (*IScalarGatewayContractCallApprovedIterator, error) {

	var commandIdRule []interface{}
	for _, commandIdItem := range commandId {
		commandIdRule = append(commandIdRule, commandIdItem)
	}

	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}
	var payloadHashRule []interface{}
	for _, payloadHashItem := range payloadHash {
		payloadHashRule = append(payloadHashRule, payloadHashItem)
	}

	logs, sub, err := _IScalarGateway.contract.FilterLogs(opts, "ContractCallApproved", commandIdRule, contractAddressRule, payloadHashRule)
	if err != nil {
		return nil, err
	}
	return &IScalarGatewayContractCallApprovedIterator{contract: _IScalarGateway.contract, event: "ContractCallApproved", logs: logs, sub: sub}, nil
}

// WatchContractCallApproved is a free log subscription operation binding the contract event 0x44e4f8f6bd682c5a3aeba93601ab07cb4d1f21b2aab1ae4880d9577919309aa4.
//
// Solidity: event ContractCallApproved(bytes32 indexed commandId, string sourceChain, string sourceAddress, address indexed contractAddress, bytes32 indexed payloadHash, bytes32 sourceTxHash, uint256 sourceEventIndex)
func (_IScalarGateway *IScalarGatewayFilterer) WatchContractCallApproved(opts *bind.WatchOpts, sink chan<- *IScalarGatewayContractCallApproved, commandId [][32]byte, contractAddress []common.Address, payloadHash [][32]byte) (event.Subscription, error) {

	var commandIdRule []interface{}
	for _, commandIdItem := range commandId {
		commandIdRule = append(commandIdRule, commandIdItem)
	}

	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}
	var payloadHashRule []interface{}
	for _, payloadHashItem := range payloadHash {
		payloadHashRule = append(payloadHashRule, payloadHashItem)
	}

	logs, sub, err := _IScalarGateway.contract.WatchLogs(opts, "ContractCallApproved", commandIdRule, contractAddressRule, payloadHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IScalarGatewayContractCallApproved)
				if err := _IScalarGateway.contract.UnpackLog(event, "ContractCallApproved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseContractCallApproved is a log parse operation binding the contract event 0x44e4f8f6bd682c5a3aeba93601ab07cb4d1f21b2aab1ae4880d9577919309aa4.
//
// Solidity: event ContractCallApproved(bytes32 indexed commandId, string sourceChain, string sourceAddress, address indexed contractAddress, bytes32 indexed payloadHash, bytes32 sourceTxHash, uint256 sourceEventIndex)
func (_IScalarGateway *IScalarGatewayFilterer) ParseContractCallApproved(log types.Log) (*IScalarGatewayContractCallApproved, error) {
	event := new(IScalarGatewayContractCallApproved)
	if err := _IScalarGateway.contract.UnpackLog(event, "ContractCallApproved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IScalarGatewayContractCallApprovedWithMintIterator is returned from FilterContractCallApprovedWithMint and is used to iterate over the raw logs and unpacked data for ContractCallApprovedWithMint events raised by the IScalarGateway contract.
type IScalarGatewayContractCallApprovedWithMintIterator struct {
	Event *IScalarGatewayContractCallApprovedWithMint // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IScalarGatewayContractCallApprovedWithMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IScalarGatewayContractCallApprovedWithMint)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IScalarGatewayContractCallApprovedWithMint)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IScalarGatewayContractCallApprovedWithMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IScalarGatewayContractCallApprovedWithMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IScalarGatewayContractCallApprovedWithMint represents a ContractCallApprovedWithMint event raised by the IScalarGateway contract.
type IScalarGatewayContractCallApprovedWithMint struct {
	CommandId        [32]byte
	SourceChain      string
	SourceAddress    string
	ContractAddress  common.Address
	PayloadHash      [32]byte
	Symbol           string
	Amount           *big.Int
	SourceTxHash     [32]byte
	SourceEventIndex *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterContractCallApprovedWithMint is a free log retrieval operation binding the contract event 0x9991faa1f435675159ffae64b66d7ecfdb55c29755869a18db8497b4392347e0.
//
// Solidity: event ContractCallApprovedWithMint(bytes32 indexed commandId, string sourceChain, string sourceAddress, address indexed contractAddress, bytes32 indexed payloadHash, string symbol, uint256 amount, bytes32 sourceTxHash, uint256 sourceEventIndex)
func (_IScalarGateway *IScalarGatewayFilterer) FilterContractCallApprovedWithMint(opts *bind.FilterOpts, commandId [][32]byte, contractAddress []common.Address, payloadHash [][32]byte) (*IScalarGatewayContractCallApprovedWithMintIterator, error) {

	var commandIdRule []interface{}
	for _, commandIdItem := range commandId {
		commandIdRule = append(commandIdRule, commandIdItem)
	}

	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}
	var payloadHashRule []interface{}
	for _, payloadHashItem := range payloadHash {
		payloadHashRule = append(payloadHashRule, payloadHashItem)
	}

	logs, sub, err := _IScalarGateway.contract.FilterLogs(opts, "ContractCallApprovedWithMint", commandIdRule, contractAddressRule, payloadHashRule)
	if err != nil {
		return nil, err
	}
	return &IScalarGatewayContractCallApprovedWithMintIterator{contract: _IScalarGateway.contract, event: "ContractCallApprovedWithMint", logs: logs, sub: sub}, nil
}

// WatchContractCallApprovedWithMint is a free log subscription operation binding the contract event 0x9991faa1f435675159ffae64b66d7ecfdb55c29755869a18db8497b4392347e0.
//
// Solidity: event ContractCallApprovedWithMint(bytes32 indexed commandId, string sourceChain, string sourceAddress, address indexed contractAddress, bytes32 indexed payloadHash, string symbol, uint256 amount, bytes32 sourceTxHash, uint256 sourceEventIndex)
func (_IScalarGateway *IScalarGatewayFilterer) WatchContractCallApprovedWithMint(opts *bind.WatchOpts, sink chan<- *IScalarGatewayContractCallApprovedWithMint, commandId [][32]byte, contractAddress []common.Address, payloadHash [][32]byte) (event.Subscription, error) {

	var commandIdRule []interface{}
	for _, commandIdItem := range commandId {
		commandIdRule = append(commandIdRule, commandIdItem)
	}

	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}
	var payloadHashRule []interface{}
	for _, payloadHashItem := range payloadHash {
		payloadHashRule = append(payloadHashRule, payloadHashItem)
	}

	logs, sub, err := _IScalarGateway.contract.WatchLogs(opts, "ContractCallApprovedWithMint", commandIdRule, contractAddressRule, payloadHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IScalarGatewayContractCallApprovedWithMint)
				if err := _IScalarGateway.contract.UnpackLog(event, "ContractCallApprovedWithMint", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseContractCallApprovedWithMint is a log parse operation binding the contract event 0x9991faa1f435675159ffae64b66d7ecfdb55c29755869a18db8497b4392347e0.
//
// Solidity: event ContractCallApprovedWithMint(bytes32 indexed commandId, string sourceChain, string sourceAddress, address indexed contractAddress, bytes32 indexed payloadHash, string symbol, uint256 amount, bytes32 sourceTxHash, uint256 sourceEventIndex)
func (_IScalarGateway *IScalarGatewayFilterer) ParseContractCallApprovedWithMint(log types.Log) (*IScalarGatewayContractCallApprovedWithMint, error) {
	event := new(IScalarGatewayContractCallApprovedWithMint)
	if err := _IScalarGateway.contract.UnpackLog(event, "ContractCallApprovedWithMint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IScalarGatewayContractCallExecutedIterator is returned from FilterContractCallExecuted and is used to iterate over the raw logs and unpacked data for ContractCallExecuted events raised by the IScalarGateway contract.
type IScalarGatewayContractCallExecutedIterator struct {
	Event *IScalarGatewayContractCallExecuted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IScalarGatewayContractCallExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IScalarGatewayContractCallExecuted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IScalarGatewayContractCallExecuted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IScalarGatewayContractCallExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IScalarGatewayContractCallExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IScalarGatewayContractCallExecuted represents a ContractCallExecuted event raised by the IScalarGateway contract.
type IScalarGatewayContractCallExecuted struct {
	CommandId [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterContractCallExecuted is a free log retrieval operation binding the contract event 0x91057b069763121972ce22b18b2f319b1520dd4c72f1f94a6395e81ceaf63f41.
//
// Solidity: event ContractCallExecuted(bytes32 indexed commandId)
func (_IScalarGateway *IScalarGatewayFilterer) FilterContractCallExecuted(opts *bind.FilterOpts, commandId [][32]byte) (*IScalarGatewayContractCallExecutedIterator, error) {

	var commandIdRule []interface{}
	for _, commandIdItem := range commandId {
		commandIdRule = append(commandIdRule, commandIdItem)
	}

	logs, sub, err := _IScalarGateway.contract.FilterLogs(opts, "ContractCallExecuted", commandIdRule)
	if err != nil {
		return nil, err
	}
	return &IScalarGatewayContractCallExecutedIterator{contract: _IScalarGateway.contract, event: "ContractCallExecuted", logs: logs, sub: sub}, nil
}

// WatchContractCallExecuted is a free log subscription operation binding the contract event 0x91057b069763121972ce22b18b2f319b1520dd4c72f1f94a6395e81ceaf63f41.
//
// Solidity: event ContractCallExecuted(bytes32 indexed commandId)
func (_IScalarGateway *IScalarGatewayFilterer) WatchContractCallExecuted(opts *bind.WatchOpts, sink chan<- *IScalarGatewayContractCallExecuted, commandId [][32]byte) (event.Subscription, error) {

	var commandIdRule []interface{}
	for _, commandIdItem := range commandId {
		commandIdRule = append(commandIdRule, commandIdItem)
	}

	logs, sub, err := _IScalarGateway.contract.WatchLogs(opts, "ContractCallExecuted", commandIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IScalarGatewayContractCallExecuted)
				if err := _IScalarGateway.contract.UnpackLog(event, "ContractCallExecuted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseContractCallExecuted is a log parse operation binding the contract event 0x91057b069763121972ce22b18b2f319b1520dd4c72f1f94a6395e81ceaf63f41.
//
// Solidity: event ContractCallExecuted(bytes32 indexed commandId)
func (_IScalarGateway *IScalarGatewayFilterer) ParseContractCallExecuted(log types.Log) (*IScalarGatewayContractCallExecuted, error) {
	event := new(IScalarGatewayContractCallExecuted)
	if err := _IScalarGateway.contract.UnpackLog(event, "ContractCallExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IScalarGatewayContractCallWithTokenIterator is returned from FilterContractCallWithToken and is used to iterate over the raw logs and unpacked data for ContractCallWithToken events raised by the IScalarGateway contract.
type IScalarGatewayContractCallWithTokenIterator struct {
	Event *IScalarGatewayContractCallWithToken // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IScalarGatewayContractCallWithTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IScalarGatewayContractCallWithToken)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IScalarGatewayContractCallWithToken)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IScalarGatewayContractCallWithTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IScalarGatewayContractCallWithTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IScalarGatewayContractCallWithToken represents a ContractCallWithToken event raised by the IScalarGateway contract.
type IScalarGatewayContractCallWithToken struct {
	Sender                     common.Address
	DestinationChain           string
	DestinationContractAddress string
	PayloadHash                [32]byte
	Payload                    []byte
	Symbol                     string
	Amount                     *big.Int
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterContractCallWithToken is a free log retrieval operation binding the contract event 0x7e50569d26be643bda7757722291ec66b1be66d8283474ae3fab5a98f878a7a2.
//
// Solidity: event ContractCallWithToken(address indexed sender, string destinationChain, string destinationContractAddress, bytes32 indexed payloadHash, bytes payload, string symbol, uint256 amount)
func (_IScalarGateway *IScalarGatewayFilterer) FilterContractCallWithToken(opts *bind.FilterOpts, sender []common.Address, payloadHash [][32]byte) (*IScalarGatewayContractCallWithTokenIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var payloadHashRule []interface{}
	for _, payloadHashItem := range payloadHash {
		payloadHashRule = append(payloadHashRule, payloadHashItem)
	}

	logs, sub, err := _IScalarGateway.contract.FilterLogs(opts, "ContractCallWithToken", senderRule, payloadHashRule)
	if err != nil {
		return nil, err
	}
	return &IScalarGatewayContractCallWithTokenIterator{contract: _IScalarGateway.contract, event: "ContractCallWithToken", logs: logs, sub: sub}, nil
}

// WatchContractCallWithToken is a free log subscription operation binding the contract event 0x7e50569d26be643bda7757722291ec66b1be66d8283474ae3fab5a98f878a7a2.
//
// Solidity: event ContractCallWithToken(address indexed sender, string destinationChain, string destinationContractAddress, bytes32 indexed payloadHash, bytes payload, string symbol, uint256 amount)
func (_IScalarGateway *IScalarGatewayFilterer) WatchContractCallWithToken(opts *bind.WatchOpts, sink chan<- *IScalarGatewayContractCallWithToken, sender []common.Address, payloadHash [][32]byte) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var payloadHashRule []interface{}
	for _, payloadHashItem := range payloadHash {
		payloadHashRule = append(payloadHashRule, payloadHashItem)
	}

	logs, sub, err := _IScalarGateway.contract.WatchLogs(opts, "ContractCallWithToken", senderRule, payloadHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IScalarGatewayContractCallWithToken)
				if err := _IScalarGateway.contract.UnpackLog(event, "ContractCallWithToken", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseContractCallWithToken is a log parse operation binding the contract event 0x7e50569d26be643bda7757722291ec66b1be66d8283474ae3fab5a98f878a7a2.
//
// Solidity: event ContractCallWithToken(address indexed sender, string destinationChain, string destinationContractAddress, bytes32 indexed payloadHash, bytes payload, string symbol, uint256 amount)
func (_IScalarGateway *IScalarGatewayFilterer) ParseContractCallWithToken(log types.Log) (*IScalarGatewayContractCallWithToken, error) {
	event := new(IScalarGatewayContractCallWithToken)
	if err := _IScalarGateway.contract.UnpackLog(event, "ContractCallWithToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IScalarGatewayExecutedIterator is returned from FilterExecuted and is used to iterate over the raw logs and unpacked data for Executed events raised by the IScalarGateway contract.
type IScalarGatewayExecutedIterator struct {
	Event *IScalarGatewayExecuted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IScalarGatewayExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IScalarGatewayExecuted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IScalarGatewayExecuted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IScalarGatewayExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IScalarGatewayExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IScalarGatewayExecuted represents a Executed event raised by the IScalarGateway contract.
type IScalarGatewayExecuted struct {
	CommandId [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterExecuted is a free log retrieval operation binding the contract event 0xa74c8847d513feba22a0f0cb38d53081abf97562cdb293926ba243689e7c41ca.
//
// Solidity: event Executed(bytes32 indexed commandId)
func (_IScalarGateway *IScalarGatewayFilterer) FilterExecuted(opts *bind.FilterOpts, commandId [][32]byte) (*IScalarGatewayExecutedIterator, error) {

	var commandIdRule []interface{}
	for _, commandIdItem := range commandId {
		commandIdRule = append(commandIdRule, commandIdItem)
	}

	logs, sub, err := _IScalarGateway.contract.FilterLogs(opts, "Executed", commandIdRule)
	if err != nil {
		return nil, err
	}
	return &IScalarGatewayExecutedIterator{contract: _IScalarGateway.contract, event: "Executed", logs: logs, sub: sub}, nil
}

// WatchExecuted is a free log subscription operation binding the contract event 0xa74c8847d513feba22a0f0cb38d53081abf97562cdb293926ba243689e7c41ca.
//
// Solidity: event Executed(bytes32 indexed commandId)
func (_IScalarGateway *IScalarGatewayFilterer) WatchExecuted(opts *bind.WatchOpts, sink chan<- *IScalarGatewayExecuted, commandId [][32]byte) (event.Subscription, error) {

	var commandIdRule []interface{}
	for _, commandIdItem := range commandId {
		commandIdRule = append(commandIdRule, commandIdItem)
	}

	logs, sub, err := _IScalarGateway.contract.WatchLogs(opts, "Executed", commandIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IScalarGatewayExecuted)
				if err := _IScalarGateway.contract.UnpackLog(event, "Executed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseExecuted is a log parse operation binding the contract event 0xa74c8847d513feba22a0f0cb38d53081abf97562cdb293926ba243689e7c41ca.
//
// Solidity: event Executed(bytes32 indexed commandId)
func (_IScalarGateway *IScalarGatewayFilterer) ParseExecuted(log types.Log) (*IScalarGatewayExecuted, error) {
	event := new(IScalarGatewayExecuted)
	if err := _IScalarGateway.contract.UnpackLog(event, "Executed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IScalarGatewayGovernanceTransferredIterator is returned from FilterGovernanceTransferred and is used to iterate over the raw logs and unpacked data for GovernanceTransferred events raised by the IScalarGateway contract.
type IScalarGatewayGovernanceTransferredIterator struct {
	Event *IScalarGatewayGovernanceTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IScalarGatewayGovernanceTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IScalarGatewayGovernanceTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IScalarGatewayGovernanceTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IScalarGatewayGovernanceTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IScalarGatewayGovernanceTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IScalarGatewayGovernanceTransferred represents a GovernanceTransferred event raised by the IScalarGateway contract.
type IScalarGatewayGovernanceTransferred struct {
	PreviousGovernance common.Address
	NewGovernance      common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterGovernanceTransferred is a free log retrieval operation binding the contract event 0x5f56bee8cffbe9a78652a74a60705edede02af10b0bbb888ca44b79a0d42ce80.
//
// Solidity: event GovernanceTransferred(address indexed previousGovernance, address indexed newGovernance)
func (_IScalarGateway *IScalarGatewayFilterer) FilterGovernanceTransferred(opts *bind.FilterOpts, previousGovernance []common.Address, newGovernance []common.Address) (*IScalarGatewayGovernanceTransferredIterator, error) {

	var previousGovernanceRule []interface{}
	for _, previousGovernanceItem := range previousGovernance {
		previousGovernanceRule = append(previousGovernanceRule, previousGovernanceItem)
	}
	var newGovernanceRule []interface{}
	for _, newGovernanceItem := range newGovernance {
		newGovernanceRule = append(newGovernanceRule, newGovernanceItem)
	}

	logs, sub, err := _IScalarGateway.contract.FilterLogs(opts, "GovernanceTransferred", previousGovernanceRule, newGovernanceRule)
	if err != nil {
		return nil, err
	}
	return &IScalarGatewayGovernanceTransferredIterator{contract: _IScalarGateway.contract, event: "GovernanceTransferred", logs: logs, sub: sub}, nil
}

// WatchGovernanceTransferred is a free log subscription operation binding the contract event 0x5f56bee8cffbe9a78652a74a60705edede02af10b0bbb888ca44b79a0d42ce80.
//
// Solidity: event GovernanceTransferred(address indexed previousGovernance, address indexed newGovernance)
func (_IScalarGateway *IScalarGatewayFilterer) WatchGovernanceTransferred(opts *bind.WatchOpts, sink chan<- *IScalarGatewayGovernanceTransferred, previousGovernance []common.Address, newGovernance []common.Address) (event.Subscription, error) {

	var previousGovernanceRule []interface{}
	for _, previousGovernanceItem := range previousGovernance {
		previousGovernanceRule = append(previousGovernanceRule, previousGovernanceItem)
	}
	var newGovernanceRule []interface{}
	for _, newGovernanceItem := range newGovernance {
		newGovernanceRule = append(newGovernanceRule, newGovernanceItem)
	}

	logs, sub, err := _IScalarGateway.contract.WatchLogs(opts, "GovernanceTransferred", previousGovernanceRule, newGovernanceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IScalarGatewayGovernanceTransferred)
				if err := _IScalarGateway.contract.UnpackLog(event, "GovernanceTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseGovernanceTransferred is a log parse operation binding the contract event 0x5f56bee8cffbe9a78652a74a60705edede02af10b0bbb888ca44b79a0d42ce80.
//
// Solidity: event GovernanceTransferred(address indexed previousGovernance, address indexed newGovernance)
func (_IScalarGateway *IScalarGatewayFilterer) ParseGovernanceTransferred(log types.Log) (*IScalarGatewayGovernanceTransferred, error) {
	event := new(IScalarGatewayGovernanceTransferred)
	if err := _IScalarGateway.contract.UnpackLog(event, "GovernanceTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IScalarGatewayMintLimiterTransferredIterator is returned from FilterMintLimiterTransferred and is used to iterate over the raw logs and unpacked data for MintLimiterTransferred events raised by the IScalarGateway contract.
type IScalarGatewayMintLimiterTransferredIterator struct {
	Event *IScalarGatewayMintLimiterTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IScalarGatewayMintLimiterTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IScalarGatewayMintLimiterTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IScalarGatewayMintLimiterTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IScalarGatewayMintLimiterTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IScalarGatewayMintLimiterTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IScalarGatewayMintLimiterTransferred represents a MintLimiterTransferred event raised by the IScalarGateway contract.
type IScalarGatewayMintLimiterTransferred struct {
	PreviousGovernance common.Address
	NewGovernance      common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterMintLimiterTransferred is a free log retrieval operation binding the contract event 0xa9303c860c5de3c0c866c354d281785c89778ac5ca2dffdf12841c45cd4e1e6e.
//
// Solidity: event MintLimiterTransferred(address indexed previousGovernance, address indexed newGovernance)
func (_IScalarGateway *IScalarGatewayFilterer) FilterMintLimiterTransferred(opts *bind.FilterOpts, previousGovernance []common.Address, newGovernance []common.Address) (*IScalarGatewayMintLimiterTransferredIterator, error) {

	var previousGovernanceRule []interface{}
	for _, previousGovernanceItem := range previousGovernance {
		previousGovernanceRule = append(previousGovernanceRule, previousGovernanceItem)
	}
	var newGovernanceRule []interface{}
	for _, newGovernanceItem := range newGovernance {
		newGovernanceRule = append(newGovernanceRule, newGovernanceItem)
	}

	logs, sub, err := _IScalarGateway.contract.FilterLogs(opts, "MintLimiterTransferred", previousGovernanceRule, newGovernanceRule)
	if err != nil {
		return nil, err
	}
	return &IScalarGatewayMintLimiterTransferredIterator{contract: _IScalarGateway.contract, event: "MintLimiterTransferred", logs: logs, sub: sub}, nil
}

// WatchMintLimiterTransferred is a free log subscription operation binding the contract event 0xa9303c860c5de3c0c866c354d281785c89778ac5ca2dffdf12841c45cd4e1e6e.
//
// Solidity: event MintLimiterTransferred(address indexed previousGovernance, address indexed newGovernance)
func (_IScalarGateway *IScalarGatewayFilterer) WatchMintLimiterTransferred(opts *bind.WatchOpts, sink chan<- *IScalarGatewayMintLimiterTransferred, previousGovernance []common.Address, newGovernance []common.Address) (event.Subscription, error) {

	var previousGovernanceRule []interface{}
	for _, previousGovernanceItem := range previousGovernance {
		previousGovernanceRule = append(previousGovernanceRule, previousGovernanceItem)
	}
	var newGovernanceRule []interface{}
	for _, newGovernanceItem := range newGovernance {
		newGovernanceRule = append(newGovernanceRule, newGovernanceItem)
	}

	logs, sub, err := _IScalarGateway.contract.WatchLogs(opts, "MintLimiterTransferred", previousGovernanceRule, newGovernanceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IScalarGatewayMintLimiterTransferred)
				if err := _IScalarGateway.contract.UnpackLog(event, "MintLimiterTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMintLimiterTransferred is a log parse operation binding the contract event 0xa9303c860c5de3c0c866c354d281785c89778ac5ca2dffdf12841c45cd4e1e6e.
//
// Solidity: event MintLimiterTransferred(address indexed previousGovernance, address indexed newGovernance)
func (_IScalarGateway *IScalarGatewayFilterer) ParseMintLimiterTransferred(log types.Log) (*IScalarGatewayMintLimiterTransferred, error) {
	event := new(IScalarGatewayMintLimiterTransferred)
	if err := _IScalarGateway.contract.UnpackLog(event, "MintLimiterTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IScalarGatewayOperatorshipTransferredIterator is returned from FilterOperatorshipTransferred and is used to iterate over the raw logs and unpacked data for OperatorshipTransferred events raised by the IScalarGateway contract.
type IScalarGatewayOperatorshipTransferredIterator struct {
	Event *IScalarGatewayOperatorshipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IScalarGatewayOperatorshipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IScalarGatewayOperatorshipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IScalarGatewayOperatorshipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IScalarGatewayOperatorshipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IScalarGatewayOperatorshipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IScalarGatewayOperatorshipTransferred represents a OperatorshipTransferred event raised by the IScalarGateway contract.
type IScalarGatewayOperatorshipTransferred struct {
	NewOperatorsData []byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterOperatorshipTransferred is a free log retrieval operation binding the contract event 0x192e759e55f359cd9832b5c0c6e38e4b6df5c5ca33f3bd5c90738e865a521872.
//
// Solidity: event OperatorshipTransferred(bytes newOperatorsData)
func (_IScalarGateway *IScalarGatewayFilterer) FilterOperatorshipTransferred(opts *bind.FilterOpts) (*IScalarGatewayOperatorshipTransferredIterator, error) {

	logs, sub, err := _IScalarGateway.contract.FilterLogs(opts, "OperatorshipTransferred")
	if err != nil {
		return nil, err
	}
	return &IScalarGatewayOperatorshipTransferredIterator{contract: _IScalarGateway.contract, event: "OperatorshipTransferred", logs: logs, sub: sub}, nil
}

// WatchOperatorshipTransferred is a free log subscription operation binding the contract event 0x192e759e55f359cd9832b5c0c6e38e4b6df5c5ca33f3bd5c90738e865a521872.
//
// Solidity: event OperatorshipTransferred(bytes newOperatorsData)
func (_IScalarGateway *IScalarGatewayFilterer) WatchOperatorshipTransferred(opts *bind.WatchOpts, sink chan<- *IScalarGatewayOperatorshipTransferred) (event.Subscription, error) {

	logs, sub, err := _IScalarGateway.contract.WatchLogs(opts, "OperatorshipTransferred")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IScalarGatewayOperatorshipTransferred)
				if err := _IScalarGateway.contract.UnpackLog(event, "OperatorshipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOperatorshipTransferred is a log parse operation binding the contract event 0x192e759e55f359cd9832b5c0c6e38e4b6df5c5ca33f3bd5c90738e865a521872.
//
// Solidity: event OperatorshipTransferred(bytes newOperatorsData)
func (_IScalarGateway *IScalarGatewayFilterer) ParseOperatorshipTransferred(log types.Log) (*IScalarGatewayOperatorshipTransferred, error) {
	event := new(IScalarGatewayOperatorshipTransferred)
	if err := _IScalarGateway.contract.UnpackLog(event, "OperatorshipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IScalarGatewayRedeemTokenIterator is returned from FilterRedeemToken and is used to iterate over the raw logs and unpacked data for RedeemToken events raised by the IScalarGateway contract.
type IScalarGatewayRedeemTokenIterator struct {
	Event *IScalarGatewayRedeemToken // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IScalarGatewayRedeemTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IScalarGatewayRedeemToken)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IScalarGatewayRedeemToken)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IScalarGatewayRedeemTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IScalarGatewayRedeemTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IScalarGatewayRedeemToken represents a RedeemToken event raised by the IScalarGateway contract.
type IScalarGatewayRedeemToken struct {
	Sender                     common.Address
	Sequence                   uint64
	CustodianGroupUID          [32]byte
	DestinationChain           string
	DestinationContractAddress string
	PayloadHash                [32]byte
	Payload                    []byte
	Symbol                     string
	Amount                     *big.Int
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterRedeemToken is a free log retrieval operation binding the contract event 0x985ccdc26bd7af05b8e15ab7607dc07f483a1f6a6ba9aa8daabee4e0cba3de11.
//
// Solidity: event RedeemToken(address indexed sender, uint64 indexed sequence, bytes32 custodianGroupUID, string destinationChain, string destinationContractAddress, bytes32 indexed payloadHash, bytes payload, string symbol, uint256 amount)
func (_IScalarGateway *IScalarGatewayFilterer) FilterRedeemToken(opts *bind.FilterOpts, sender []common.Address, sequence []uint64, payloadHash [][32]byte) (*IScalarGatewayRedeemTokenIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var sequenceRule []interface{}
	for _, sequenceItem := range sequence {
		sequenceRule = append(sequenceRule, sequenceItem)
	}

	var payloadHashRule []interface{}
	for _, payloadHashItem := range payloadHash {
		payloadHashRule = append(payloadHashRule, payloadHashItem)
	}

	logs, sub, err := _IScalarGateway.contract.FilterLogs(opts, "RedeemToken", senderRule, sequenceRule, payloadHashRule)
	if err != nil {
		return nil, err
	}
	return &IScalarGatewayRedeemTokenIterator{contract: _IScalarGateway.contract, event: "RedeemToken", logs: logs, sub: sub}, nil
}

// WatchRedeemToken is a free log subscription operation binding the contract event 0x985ccdc26bd7af05b8e15ab7607dc07f483a1f6a6ba9aa8daabee4e0cba3de11.
//
// Solidity: event RedeemToken(address indexed sender, uint64 indexed sequence, bytes32 custodianGroupUID, string destinationChain, string destinationContractAddress, bytes32 indexed payloadHash, bytes payload, string symbol, uint256 amount)
func (_IScalarGateway *IScalarGatewayFilterer) WatchRedeemToken(opts *bind.WatchOpts, sink chan<- *IScalarGatewayRedeemToken, sender []common.Address, sequence []uint64, payloadHash [][32]byte) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var sequenceRule []interface{}
	for _, sequenceItem := range sequence {
		sequenceRule = append(sequenceRule, sequenceItem)
	}

	var payloadHashRule []interface{}
	for _, payloadHashItem := range payloadHash {
		payloadHashRule = append(payloadHashRule, payloadHashItem)
	}

	logs, sub, err := _IScalarGateway.contract.WatchLogs(opts, "RedeemToken", senderRule, sequenceRule, payloadHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IScalarGatewayRedeemToken)
				if err := _IScalarGateway.contract.UnpackLog(event, "RedeemToken", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRedeemToken is a log parse operation binding the contract event 0x985ccdc26bd7af05b8e15ab7607dc07f483a1f6a6ba9aa8daabee4e0cba3de11.
//
// Solidity: event RedeemToken(address indexed sender, uint64 indexed sequence, bytes32 custodianGroupUID, string destinationChain, string destinationContractAddress, bytes32 indexed payloadHash, bytes payload, string symbol, uint256 amount)
func (_IScalarGateway *IScalarGatewayFilterer) ParseRedeemToken(log types.Log) (*IScalarGatewayRedeemToken, error) {
	event := new(IScalarGatewayRedeemToken)
	if err := _IScalarGateway.contract.UnpackLog(event, "RedeemToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IScalarGatewayRegisterCustodianGroupIterator is returned from FilterRegisterCustodianGroup and is used to iterate over the raw logs and unpacked data for RegisterCustodianGroup events raised by the IScalarGateway contract.
type IScalarGatewayRegisterCustodianGroupIterator struct {
	Event *IScalarGatewayRegisterCustodianGroup // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IScalarGatewayRegisterCustodianGroupIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IScalarGatewayRegisterCustodianGroup)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IScalarGatewayRegisterCustodianGroup)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IScalarGatewayRegisterCustodianGroupIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IScalarGatewayRegisterCustodianGroupIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IScalarGatewayRegisterCustodianGroup represents a RegisterCustodianGroup event raised by the IScalarGateway contract.
type IScalarGatewayRegisterCustodianGroup struct {
	CustodianGroupId [32]byte
	Sequence         uint64
	Phase            uint8
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterRegisterCustodianGroup is a free log retrieval operation binding the contract event 0xe7ed286e40bd8ca4ee93649688af9b82b893998465bc78ce2a3ca042f6537fbc.
//
// Solidity: event RegisterCustodianGroup(bytes32 indexed custodianGroupId, uint64 sequence, uint8 phase)
func (_IScalarGateway *IScalarGatewayFilterer) FilterRegisterCustodianGroup(opts *bind.FilterOpts, custodianGroupId [][32]byte) (*IScalarGatewayRegisterCustodianGroupIterator, error) {

	var custodianGroupIdRule []interface{}
	for _, custodianGroupIdItem := range custodianGroupId {
		custodianGroupIdRule = append(custodianGroupIdRule, custodianGroupIdItem)
	}

	logs, sub, err := _IScalarGateway.contract.FilterLogs(opts, "RegisterCustodianGroup", custodianGroupIdRule)
	if err != nil {
		return nil, err
	}
	return &IScalarGatewayRegisterCustodianGroupIterator{contract: _IScalarGateway.contract, event: "RegisterCustodianGroup", logs: logs, sub: sub}, nil
}

// WatchRegisterCustodianGroup is a free log subscription operation binding the contract event 0xe7ed286e40bd8ca4ee93649688af9b82b893998465bc78ce2a3ca042f6537fbc.
//
// Solidity: event RegisterCustodianGroup(bytes32 indexed custodianGroupId, uint64 sequence, uint8 phase)
func (_IScalarGateway *IScalarGatewayFilterer) WatchRegisterCustodianGroup(opts *bind.WatchOpts, sink chan<- *IScalarGatewayRegisterCustodianGroup, custodianGroupId [][32]byte) (event.Subscription, error) {

	var custodianGroupIdRule []interface{}
	for _, custodianGroupIdItem := range custodianGroupId {
		custodianGroupIdRule = append(custodianGroupIdRule, custodianGroupIdItem)
	}

	logs, sub, err := _IScalarGateway.contract.WatchLogs(opts, "RegisterCustodianGroup", custodianGroupIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IScalarGatewayRegisterCustodianGroup)
				if err := _IScalarGateway.contract.UnpackLog(event, "RegisterCustodianGroup", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRegisterCustodianGroup is a log parse operation binding the contract event 0xe7ed286e40bd8ca4ee93649688af9b82b893998465bc78ce2a3ca042f6537fbc.
//
// Solidity: event RegisterCustodianGroup(bytes32 indexed custodianGroupId, uint64 sequence, uint8 phase)
func (_IScalarGateway *IScalarGatewayFilterer) ParseRegisterCustodianGroup(log types.Log) (*IScalarGatewayRegisterCustodianGroup, error) {
	event := new(IScalarGatewayRegisterCustodianGroup)
	if err := _IScalarGateway.contract.UnpackLog(event, "RegisterCustodianGroup", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IScalarGatewaySwitchPhaseIterator is returned from FilterSwitchPhase and is used to iterate over the raw logs and unpacked data for SwitchPhase events raised by the IScalarGateway contract.
type IScalarGatewaySwitchPhaseIterator struct {
	Event *IScalarGatewaySwitchPhase // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IScalarGatewaySwitchPhaseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IScalarGatewaySwitchPhase)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IScalarGatewaySwitchPhase)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IScalarGatewaySwitchPhaseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IScalarGatewaySwitchPhaseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IScalarGatewaySwitchPhase represents a SwitchPhase event raised by the IScalarGateway contract.
type IScalarGatewaySwitchPhase struct {
	CustodianGroupId [32]byte
	Sequence         uint64
	From             uint8
	To               uint8
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSwitchPhase is a free log retrieval operation binding the contract event 0x688700c66a6bfdcb89a49d388166fec61901058bc742019b144019a784fd4278.
//
// Solidity: event SwitchPhase(bytes32 indexed custodianGroupId, uint64 indexed sequence, uint8 from, uint8 to)
func (_IScalarGateway *IScalarGatewayFilterer) FilterSwitchPhase(opts *bind.FilterOpts, custodianGroupId [][32]byte, sequence []uint64) (*IScalarGatewaySwitchPhaseIterator, error) {

	var custodianGroupIdRule []interface{}
	for _, custodianGroupIdItem := range custodianGroupId {
		custodianGroupIdRule = append(custodianGroupIdRule, custodianGroupIdItem)
	}
	var sequenceRule []interface{}
	for _, sequenceItem := range sequence {
		sequenceRule = append(sequenceRule, sequenceItem)
	}

	logs, sub, err := _IScalarGateway.contract.FilterLogs(opts, "SwitchPhase", custodianGroupIdRule, sequenceRule)
	if err != nil {
		return nil, err
	}
	return &IScalarGatewaySwitchPhaseIterator{contract: _IScalarGateway.contract, event: "SwitchPhase", logs: logs, sub: sub}, nil
}

// WatchSwitchPhase is a free log subscription operation binding the contract event 0x688700c66a6bfdcb89a49d388166fec61901058bc742019b144019a784fd4278.
//
// Solidity: event SwitchPhase(bytes32 indexed custodianGroupId, uint64 indexed sequence, uint8 from, uint8 to)
func (_IScalarGateway *IScalarGatewayFilterer) WatchSwitchPhase(opts *bind.WatchOpts, sink chan<- *IScalarGatewaySwitchPhase, custodianGroupId [][32]byte, sequence []uint64) (event.Subscription, error) {

	var custodianGroupIdRule []interface{}
	for _, custodianGroupIdItem := range custodianGroupId {
		custodianGroupIdRule = append(custodianGroupIdRule, custodianGroupIdItem)
	}
	var sequenceRule []interface{}
	for _, sequenceItem := range sequence {
		sequenceRule = append(sequenceRule, sequenceItem)
	}

	logs, sub, err := _IScalarGateway.contract.WatchLogs(opts, "SwitchPhase", custodianGroupIdRule, sequenceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IScalarGatewaySwitchPhase)
				if err := _IScalarGateway.contract.UnpackLog(event, "SwitchPhase", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSwitchPhase is a log parse operation binding the contract event 0x688700c66a6bfdcb89a49d388166fec61901058bc742019b144019a784fd4278.
//
// Solidity: event SwitchPhase(bytes32 indexed custodianGroupId, uint64 indexed sequence, uint8 from, uint8 to)
func (_IScalarGateway *IScalarGatewayFilterer) ParseSwitchPhase(log types.Log) (*IScalarGatewaySwitchPhase, error) {
	event := new(IScalarGatewaySwitchPhase)
	if err := _IScalarGateway.contract.UnpackLog(event, "SwitchPhase", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IScalarGatewayTokenDeployedIterator is returned from FilterTokenDeployed and is used to iterate over the raw logs and unpacked data for TokenDeployed events raised by the IScalarGateway contract.
type IScalarGatewayTokenDeployedIterator struct {
	Event *IScalarGatewayTokenDeployed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IScalarGatewayTokenDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IScalarGatewayTokenDeployed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IScalarGatewayTokenDeployed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IScalarGatewayTokenDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IScalarGatewayTokenDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IScalarGatewayTokenDeployed represents a TokenDeployed event raised by the IScalarGateway contract.
type IScalarGatewayTokenDeployed struct {
	Symbol         string
	TokenAddresses common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterTokenDeployed is a free log retrieval operation binding the contract event 0xbf90b5a1ec9763e8bf4b9245cef0c28db92bab309fc2c5177f17814f38246938.
//
// Solidity: event TokenDeployed(string symbol, address tokenAddresses)
func (_IScalarGateway *IScalarGatewayFilterer) FilterTokenDeployed(opts *bind.FilterOpts) (*IScalarGatewayTokenDeployedIterator, error) {

	logs, sub, err := _IScalarGateway.contract.FilterLogs(opts, "TokenDeployed")
	if err != nil {
		return nil, err
	}
	return &IScalarGatewayTokenDeployedIterator{contract: _IScalarGateway.contract, event: "TokenDeployed", logs: logs, sub: sub}, nil
}

// WatchTokenDeployed is a free log subscription operation binding the contract event 0xbf90b5a1ec9763e8bf4b9245cef0c28db92bab309fc2c5177f17814f38246938.
//
// Solidity: event TokenDeployed(string symbol, address tokenAddresses)
func (_IScalarGateway *IScalarGatewayFilterer) WatchTokenDeployed(opts *bind.WatchOpts, sink chan<- *IScalarGatewayTokenDeployed) (event.Subscription, error) {

	logs, sub, err := _IScalarGateway.contract.WatchLogs(opts, "TokenDeployed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IScalarGatewayTokenDeployed)
				if err := _IScalarGateway.contract.UnpackLog(event, "TokenDeployed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokenDeployed is a log parse operation binding the contract event 0xbf90b5a1ec9763e8bf4b9245cef0c28db92bab309fc2c5177f17814f38246938.
//
// Solidity: event TokenDeployed(string symbol, address tokenAddresses)
func (_IScalarGateway *IScalarGatewayFilterer) ParseTokenDeployed(log types.Log) (*IScalarGatewayTokenDeployed, error) {
	event := new(IScalarGatewayTokenDeployed)
	if err := _IScalarGateway.contract.UnpackLog(event, "TokenDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IScalarGatewayTokenMintLimitUpdatedIterator is returned from FilterTokenMintLimitUpdated and is used to iterate over the raw logs and unpacked data for TokenMintLimitUpdated events raised by the IScalarGateway contract.
type IScalarGatewayTokenMintLimitUpdatedIterator struct {
	Event *IScalarGatewayTokenMintLimitUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IScalarGatewayTokenMintLimitUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IScalarGatewayTokenMintLimitUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IScalarGatewayTokenMintLimitUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IScalarGatewayTokenMintLimitUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IScalarGatewayTokenMintLimitUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IScalarGatewayTokenMintLimitUpdated represents a TokenMintLimitUpdated event raised by the IScalarGateway contract.
type IScalarGatewayTokenMintLimitUpdated struct {
	Symbol string
	Limit  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokenMintLimitUpdated is a free log retrieval operation binding the contract event 0xd99446c1d76385bb5519ccfb5274abcfd5896dfc22405e40010fde217f018a18.
//
// Solidity: event TokenMintLimitUpdated(string symbol, uint256 limit)
func (_IScalarGateway *IScalarGatewayFilterer) FilterTokenMintLimitUpdated(opts *bind.FilterOpts) (*IScalarGatewayTokenMintLimitUpdatedIterator, error) {

	logs, sub, err := _IScalarGateway.contract.FilterLogs(opts, "TokenMintLimitUpdated")
	if err != nil {
		return nil, err
	}
	return &IScalarGatewayTokenMintLimitUpdatedIterator{contract: _IScalarGateway.contract, event: "TokenMintLimitUpdated", logs: logs, sub: sub}, nil
}

// WatchTokenMintLimitUpdated is a free log subscription operation binding the contract event 0xd99446c1d76385bb5519ccfb5274abcfd5896dfc22405e40010fde217f018a18.
//
// Solidity: event TokenMintLimitUpdated(string symbol, uint256 limit)
func (_IScalarGateway *IScalarGatewayFilterer) WatchTokenMintLimitUpdated(opts *bind.WatchOpts, sink chan<- *IScalarGatewayTokenMintLimitUpdated) (event.Subscription, error) {

	logs, sub, err := _IScalarGateway.contract.WatchLogs(opts, "TokenMintLimitUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IScalarGatewayTokenMintLimitUpdated)
				if err := _IScalarGateway.contract.UnpackLog(event, "TokenMintLimitUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokenMintLimitUpdated is a log parse operation binding the contract event 0xd99446c1d76385bb5519ccfb5274abcfd5896dfc22405e40010fde217f018a18.
//
// Solidity: event TokenMintLimitUpdated(string symbol, uint256 limit)
func (_IScalarGateway *IScalarGatewayFilterer) ParseTokenMintLimitUpdated(log types.Log) (*IScalarGatewayTokenMintLimitUpdated, error) {
	event := new(IScalarGatewayTokenMintLimitUpdated)
	if err := _IScalarGateway.contract.UnpackLog(event, "TokenMintLimitUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IScalarGatewayTokenSentIterator is returned from FilterTokenSent and is used to iterate over the raw logs and unpacked data for TokenSent events raised by the IScalarGateway contract.
type IScalarGatewayTokenSentIterator struct {
	Event *IScalarGatewayTokenSent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IScalarGatewayTokenSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IScalarGatewayTokenSent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IScalarGatewayTokenSent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IScalarGatewayTokenSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IScalarGatewayTokenSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IScalarGatewayTokenSent represents a TokenSent event raised by the IScalarGateway contract.
type IScalarGatewayTokenSent struct {
	Sender             common.Address
	DestinationChain   string
	DestinationAddress string
	Symbol             string
	Amount             *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterTokenSent is a free log retrieval operation binding the contract event 0x651d93f66c4329630e8d0f62488eff599e3be484da587335e8dc0fcf46062726.
//
// Solidity: event TokenSent(address indexed sender, string destinationChain, string destinationAddress, string symbol, uint256 amount)
func (_IScalarGateway *IScalarGatewayFilterer) FilterTokenSent(opts *bind.FilterOpts, sender []common.Address) (*IScalarGatewayTokenSentIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IScalarGateway.contract.FilterLogs(opts, "TokenSent", senderRule)
	if err != nil {
		return nil, err
	}
	return &IScalarGatewayTokenSentIterator{contract: _IScalarGateway.contract, event: "TokenSent", logs: logs, sub: sub}, nil
}

// WatchTokenSent is a free log subscription operation binding the contract event 0x651d93f66c4329630e8d0f62488eff599e3be484da587335e8dc0fcf46062726.
//
// Solidity: event TokenSent(address indexed sender, string destinationChain, string destinationAddress, string symbol, uint256 amount)
func (_IScalarGateway *IScalarGatewayFilterer) WatchTokenSent(opts *bind.WatchOpts, sink chan<- *IScalarGatewayTokenSent, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IScalarGateway.contract.WatchLogs(opts, "TokenSent", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IScalarGatewayTokenSent)
				if err := _IScalarGateway.contract.UnpackLog(event, "TokenSent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokenSent is a log parse operation binding the contract event 0x651d93f66c4329630e8d0f62488eff599e3be484da587335e8dc0fcf46062726.
//
// Solidity: event TokenSent(address indexed sender, string destinationChain, string destinationAddress, string symbol, uint256 amount)
func (_IScalarGateway *IScalarGatewayFilterer) ParseTokenSent(log types.Log) (*IScalarGatewayTokenSent, error) {
	event := new(IScalarGatewayTokenSent)
	if err := _IScalarGateway.contract.UnpackLog(event, "TokenSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IScalarGatewayUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the IScalarGateway contract.
type IScalarGatewayUpgradedIterator struct {
	Event *IScalarGatewayUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IScalarGatewayUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IScalarGatewayUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IScalarGatewayUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IScalarGatewayUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IScalarGatewayUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IScalarGatewayUpgraded represents a Upgraded event raised by the IScalarGateway contract.
type IScalarGatewayUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_IScalarGateway *IScalarGatewayFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*IScalarGatewayUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _IScalarGateway.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &IScalarGatewayUpgradedIterator{contract: _IScalarGateway.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_IScalarGateway *IScalarGatewayFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *IScalarGatewayUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _IScalarGateway.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IScalarGatewayUpgraded)
				if err := _IScalarGateway.contract.UnpackLog(event, "Upgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_IScalarGateway *IScalarGatewayFilterer) ParseUpgraded(log types.Log) (*IScalarGatewayUpgraded, error) {
	event := new(IScalarGatewayUpgraded)
	if err := _IScalarGateway.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
