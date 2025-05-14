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

// IScalarExecutableMetaData contains all meta data concerning the IScalarExecutable contract.
var IScalarExecutableMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"NotApprovedByGateway\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"commandId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"sourceChain\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sourceAddress\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"commandId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"sourceChain\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sourceAddress\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"tokenSymbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"executeWithToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gateway\",\"outputs\":[{\"internalType\":\"contractIAxelarGateway\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IScalarExecutableABI is the input ABI used to generate the binding from.
// Deprecated: Use IScalarExecutableMetaData.ABI instead.
var IScalarExecutableABI = IScalarExecutableMetaData.ABI

// IScalarExecutable is an auto generated Go binding around an Ethereum contract.
type IScalarExecutable struct {
	IScalarExecutableCaller     // Read-only binding to the contract
	IScalarExecutableTransactor // Write-only binding to the contract
	IScalarExecutableFilterer   // Log filterer for contract events
}

// IScalarExecutableCaller is an auto generated read-only Go binding around an Ethereum contract.
type IScalarExecutableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IScalarExecutableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IScalarExecutableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IScalarExecutableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IScalarExecutableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IScalarExecutableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IScalarExecutableSession struct {
	Contract     *IScalarExecutable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IScalarExecutableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IScalarExecutableCallerSession struct {
	Contract *IScalarExecutableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// IScalarExecutableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IScalarExecutableTransactorSession struct {
	Contract     *IScalarExecutableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// IScalarExecutableRaw is an auto generated low-level Go binding around an Ethereum contract.
type IScalarExecutableRaw struct {
	Contract *IScalarExecutable // Generic contract binding to access the raw methods on
}

// IScalarExecutableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IScalarExecutableCallerRaw struct {
	Contract *IScalarExecutableCaller // Generic read-only contract binding to access the raw methods on
}

// IScalarExecutableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IScalarExecutableTransactorRaw struct {
	Contract *IScalarExecutableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIScalarExecutable creates a new instance of IScalarExecutable, bound to a specific deployed contract.
func NewIScalarExecutable(address common.Address, backend bind.ContractBackend) (*IScalarExecutable, error) {
	contract, err := bindIScalarExecutable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IScalarExecutable{IScalarExecutableCaller: IScalarExecutableCaller{contract: contract}, IScalarExecutableTransactor: IScalarExecutableTransactor{contract: contract}, IScalarExecutableFilterer: IScalarExecutableFilterer{contract: contract}}, nil
}

// NewIScalarExecutableCaller creates a new read-only instance of IScalarExecutable, bound to a specific deployed contract.
func NewIScalarExecutableCaller(address common.Address, caller bind.ContractCaller) (*IScalarExecutableCaller, error) {
	contract, err := bindIScalarExecutable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IScalarExecutableCaller{contract: contract}, nil
}

// NewIScalarExecutableTransactor creates a new write-only instance of IScalarExecutable, bound to a specific deployed contract.
func NewIScalarExecutableTransactor(address common.Address, transactor bind.ContractTransactor) (*IScalarExecutableTransactor, error) {
	contract, err := bindIScalarExecutable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IScalarExecutableTransactor{contract: contract}, nil
}

// NewIScalarExecutableFilterer creates a new log filterer instance of IScalarExecutable, bound to a specific deployed contract.
func NewIScalarExecutableFilterer(address common.Address, filterer bind.ContractFilterer) (*IScalarExecutableFilterer, error) {
	contract, err := bindIScalarExecutable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IScalarExecutableFilterer{contract: contract}, nil
}

// bindIScalarExecutable binds a generic wrapper to an already deployed contract.
func bindIScalarExecutable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IScalarExecutableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IScalarExecutable *IScalarExecutableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IScalarExecutable.Contract.IScalarExecutableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IScalarExecutable *IScalarExecutableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IScalarExecutable.Contract.IScalarExecutableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IScalarExecutable *IScalarExecutableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IScalarExecutable.Contract.IScalarExecutableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IScalarExecutable *IScalarExecutableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IScalarExecutable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IScalarExecutable *IScalarExecutableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IScalarExecutable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IScalarExecutable *IScalarExecutableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IScalarExecutable.Contract.contract.Transact(opts, method, params...)
}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_IScalarExecutable *IScalarExecutableCaller) Gateway(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IScalarExecutable.contract.Call(opts, &out, "gateway")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_IScalarExecutable *IScalarExecutableSession) Gateway() (common.Address, error) {
	return _IScalarExecutable.Contract.Gateway(&_IScalarExecutable.CallOpts)
}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_IScalarExecutable *IScalarExecutableCallerSession) Gateway() (common.Address, error) {
	return _IScalarExecutable.Contract.Gateway(&_IScalarExecutable.CallOpts)
}

// Execute is a paid mutator transaction binding the contract method 0x49160658.
//
// Solidity: function execute(bytes32 commandId, string sourceChain, string sourceAddress, bytes payload) returns()
func (_IScalarExecutable *IScalarExecutableTransactor) Execute(opts *bind.TransactOpts, commandId [32]byte, sourceChain string, sourceAddress string, payload []byte) (*types.Transaction, error) {
	return _IScalarExecutable.contract.Transact(opts, "execute", commandId, sourceChain, sourceAddress, payload)
}

// Execute is a paid mutator transaction binding the contract method 0x49160658.
//
// Solidity: function execute(bytes32 commandId, string sourceChain, string sourceAddress, bytes payload) returns()
func (_IScalarExecutable *IScalarExecutableSession) Execute(commandId [32]byte, sourceChain string, sourceAddress string, payload []byte) (*types.Transaction, error) {
	return _IScalarExecutable.Contract.Execute(&_IScalarExecutable.TransactOpts, commandId, sourceChain, sourceAddress, payload)
}

// Execute is a paid mutator transaction binding the contract method 0x49160658.
//
// Solidity: function execute(bytes32 commandId, string sourceChain, string sourceAddress, bytes payload) returns()
func (_IScalarExecutable *IScalarExecutableTransactorSession) Execute(commandId [32]byte, sourceChain string, sourceAddress string, payload []byte) (*types.Transaction, error) {
	return _IScalarExecutable.Contract.Execute(&_IScalarExecutable.TransactOpts, commandId, sourceChain, sourceAddress, payload)
}

// ExecuteWithToken is a paid mutator transaction binding the contract method 0x1a98b2e0.
//
// Solidity: function executeWithToken(bytes32 commandId, string sourceChain, string sourceAddress, bytes payload, string tokenSymbol, uint256 amount) returns()
func (_IScalarExecutable *IScalarExecutableTransactor) ExecuteWithToken(opts *bind.TransactOpts, commandId [32]byte, sourceChain string, sourceAddress string, payload []byte, tokenSymbol string, amount *big.Int) (*types.Transaction, error) {
	return _IScalarExecutable.contract.Transact(opts, "executeWithToken", commandId, sourceChain, sourceAddress, payload, tokenSymbol, amount)
}

// ExecuteWithToken is a paid mutator transaction binding the contract method 0x1a98b2e0.
//
// Solidity: function executeWithToken(bytes32 commandId, string sourceChain, string sourceAddress, bytes payload, string tokenSymbol, uint256 amount) returns()
func (_IScalarExecutable *IScalarExecutableSession) ExecuteWithToken(commandId [32]byte, sourceChain string, sourceAddress string, payload []byte, tokenSymbol string, amount *big.Int) (*types.Transaction, error) {
	return _IScalarExecutable.Contract.ExecuteWithToken(&_IScalarExecutable.TransactOpts, commandId, sourceChain, sourceAddress, payload, tokenSymbol, amount)
}

// ExecuteWithToken is a paid mutator transaction binding the contract method 0x1a98b2e0.
//
// Solidity: function executeWithToken(bytes32 commandId, string sourceChain, string sourceAddress, bytes payload, string tokenSymbol, uint256 amount) returns()
func (_IScalarExecutable *IScalarExecutableTransactorSession) ExecuteWithToken(commandId [32]byte, sourceChain string, sourceAddress string, payload []byte, tokenSymbol string, amount *big.Int) (*types.Transaction, error) {
	return _IScalarExecutable.Contract.ExecuteWithToken(&_IScalarExecutable.TransactOpts, commandId, sourceChain, sourceAddress, payload, tokenSymbol, amount)
}
