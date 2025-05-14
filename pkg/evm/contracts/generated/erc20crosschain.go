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

// IScalarERC20CrossChainMetaData contains all meta data concerning the IScalarERC20CrossChain contract.
var IScalarERC20CrossChainMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"gateway_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"gasReceiver_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"decimals_\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"contractId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decreaseAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"subtractedValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"commandId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChain\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"sourceAddress\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"executeWithToken\",\"inputs\":[{\"name\":\"commandId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChain\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"sourceAddress\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenSymbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"faucet\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"gasService\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIAxelarGasService\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"gateway\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIAxelarGateway\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"implementation\",\"inputs\":[],\"outputs\":[{\"name\":\"implementation_\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"increaseAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"addedValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"owner_\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingOwner\",\"inputs\":[],\"outputs\":[{\"name\":\"owner_\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proposeOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setup\",\"inputs\":[{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferRemote\",\"inputs\":[{\"name\":\"destinationChain\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"destinationContractAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"encodedMetadata\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"upgrade\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"newImplementationCodeHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"params\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Executed\",\"inputs\":[{\"name\":\"sourceChain\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"sourceAddress\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"sourceTx\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FalseSender\",\"inputs\":[{\"name\":\"sourceChain\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"sourceAddress\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferStarted\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TransferRemote\",\"inputs\":[{\"name\":\"destinationChain\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"destinationContractAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAccount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAddressString\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCodeHash\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidImplementation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOwnerAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotApprovedByGateway\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotProxy\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SetupFailed\",\"inputs\":[]}]",
}

// IScalarERC20CrossChainABI is the input ABI used to generate the binding from.
// Deprecated: Use IScalarERC20CrossChainMetaData.ABI instead.
var IScalarERC20CrossChainABI = IScalarERC20CrossChainMetaData.ABI

// IScalarERC20CrossChain is an auto generated Go binding around an Ethereum contract.
type IScalarERC20CrossChain struct {
	IScalarERC20CrossChainCaller     // Read-only binding to the contract
	IScalarERC20CrossChainTransactor // Write-only binding to the contract
	IScalarERC20CrossChainFilterer   // Log filterer for contract events
}

// IScalarERC20CrossChainCaller is an auto generated read-only Go binding around an Ethereum contract.
type IScalarERC20CrossChainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IScalarERC20CrossChainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IScalarERC20CrossChainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IScalarERC20CrossChainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IScalarERC20CrossChainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IScalarERC20CrossChainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IScalarERC20CrossChainSession struct {
	Contract     *IScalarERC20CrossChain // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IScalarERC20CrossChainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IScalarERC20CrossChainCallerSession struct {
	Contract *IScalarERC20CrossChainCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// IScalarERC20CrossChainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IScalarERC20CrossChainTransactorSession struct {
	Contract     *IScalarERC20CrossChainTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// IScalarERC20CrossChainRaw is an auto generated low-level Go binding around an Ethereum contract.
type IScalarERC20CrossChainRaw struct {
	Contract *IScalarERC20CrossChain // Generic contract binding to access the raw methods on
}

// IScalarERC20CrossChainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IScalarERC20CrossChainCallerRaw struct {
	Contract *IScalarERC20CrossChainCaller // Generic read-only contract binding to access the raw methods on
}

// IScalarERC20CrossChainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IScalarERC20CrossChainTransactorRaw struct {
	Contract *IScalarERC20CrossChainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIScalarERC20CrossChain creates a new instance of IScalarERC20CrossChain, bound to a specific deployed contract.
func NewIScalarERC20CrossChain(address common.Address, backend bind.ContractBackend) (*IScalarERC20CrossChain, error) {
	contract, err := bindIScalarERC20CrossChain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IScalarERC20CrossChain{IScalarERC20CrossChainCaller: IScalarERC20CrossChainCaller{contract: contract}, IScalarERC20CrossChainTransactor: IScalarERC20CrossChainTransactor{contract: contract}, IScalarERC20CrossChainFilterer: IScalarERC20CrossChainFilterer{contract: contract}}, nil
}

// NewIScalarERC20CrossChainCaller creates a new read-only instance of IScalarERC20CrossChain, bound to a specific deployed contract.
func NewIScalarERC20CrossChainCaller(address common.Address, caller bind.ContractCaller) (*IScalarERC20CrossChainCaller, error) {
	contract, err := bindIScalarERC20CrossChain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IScalarERC20CrossChainCaller{contract: contract}, nil
}

// NewIScalarERC20CrossChainTransactor creates a new write-only instance of IScalarERC20CrossChain, bound to a specific deployed contract.
func NewIScalarERC20CrossChainTransactor(address common.Address, transactor bind.ContractTransactor) (*IScalarERC20CrossChainTransactor, error) {
	contract, err := bindIScalarERC20CrossChain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IScalarERC20CrossChainTransactor{contract: contract}, nil
}

// NewIScalarERC20CrossChainFilterer creates a new log filterer instance of IScalarERC20CrossChain, bound to a specific deployed contract.
func NewIScalarERC20CrossChainFilterer(address common.Address, filterer bind.ContractFilterer) (*IScalarERC20CrossChainFilterer, error) {
	contract, err := bindIScalarERC20CrossChain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IScalarERC20CrossChainFilterer{contract: contract}, nil
}

// bindIScalarERC20CrossChain binds a generic wrapper to an already deployed contract.
func bindIScalarERC20CrossChain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IScalarERC20CrossChainMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IScalarERC20CrossChain *IScalarERC20CrossChainRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IScalarERC20CrossChain.Contract.IScalarERC20CrossChainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IScalarERC20CrossChain *IScalarERC20CrossChainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IScalarERC20CrossChain.Contract.IScalarERC20CrossChainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IScalarERC20CrossChain *IScalarERC20CrossChainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IScalarERC20CrossChain.Contract.IScalarERC20CrossChainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IScalarERC20CrossChain *IScalarERC20CrossChainCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IScalarERC20CrossChain.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IScalarERC20CrossChain *IScalarERC20CrossChainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IScalarERC20CrossChain.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IScalarERC20CrossChain *IScalarERC20CrossChainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IScalarERC20CrossChain.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_IScalarERC20CrossChain *IScalarERC20CrossChainCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IScalarERC20CrossChain.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_IScalarERC20CrossChain *IScalarERC20CrossChainSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _IScalarERC20CrossChain.Contract.Allowance(&_IScalarERC20CrossChain.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_IScalarERC20CrossChain *IScalarERC20CrossChainCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _IScalarERC20CrossChain.Contract.Allowance(&_IScalarERC20CrossChain.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_IScalarERC20CrossChain *IScalarERC20CrossChainCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IScalarERC20CrossChain.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_IScalarERC20CrossChain *IScalarERC20CrossChainSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _IScalarERC20CrossChain.Contract.BalanceOf(&_IScalarERC20CrossChain.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_IScalarERC20CrossChain *IScalarERC20CrossChainCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _IScalarERC20CrossChain.Contract.BalanceOf(&_IScalarERC20CrossChain.CallOpts, arg0)
}

// ContractId is a free data retrieval call binding the contract method 0x8291286c.
//
// Solidity: function contractId() pure returns(bytes32)
func (_IScalarERC20CrossChain *IScalarERC20CrossChainCaller) ContractId(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IScalarERC20CrossChain.contract.Call(opts, &out, "contractId")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ContractId is a free data retrieval call binding the contract method 0x8291286c.
//
// Solidity: function contractId() pure returns(bytes32)
func (_IScalarERC20CrossChain *IScalarERC20CrossChainSession) ContractId() ([32]byte, error) {
	return _IScalarERC20CrossChain.Contract.ContractId(&_IScalarERC20CrossChain.CallOpts)
}

// ContractId is a free data retrieval call binding the contract method 0x8291286c.
//
// Solidity: function contractId() pure returns(bytes32)
func (_IScalarERC20CrossChain *IScalarERC20CrossChainCallerSession) ContractId() ([32]byte, error) {
	return _IScalarERC20CrossChain.Contract.ContractId(&_IScalarERC20CrossChain.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IScalarERC20CrossChain *IScalarERC20CrossChainCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _IScalarERC20CrossChain.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IScalarERC20CrossChain *IScalarERC20CrossChainSession) Decimals() (uint8, error) {
	return _IScalarERC20CrossChain.Contract.Decimals(&_IScalarERC20CrossChain.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IScalarERC20CrossChain *IScalarERC20CrossChainCallerSession) Decimals() (uint8, error) {
	return _IScalarERC20CrossChain.Contract.Decimals(&_IScalarERC20CrossChain.CallOpts)
}

// GasService is a free data retrieval call binding the contract method 0x6a22d8cc.
//
// Solidity: function gasService() view returns(address)
func (_IScalarERC20CrossChain *IScalarERC20CrossChainCaller) GasService(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IScalarERC20CrossChain.contract.Call(opts, &out, "gasService")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GasService is a free data retrieval call binding the contract method 0x6a22d8cc.
//
// Solidity: function gasService() view returns(address)
func (_IScalarERC20CrossChain *IScalarERC20CrossChainSession) GasService() (common.Address, error) {
	return _IScalarERC20CrossChain.Contract.GasService(&_IScalarERC20CrossChain.CallOpts)
}

// GasService is a free data retrieval call binding the contract method 0x6a22d8cc.
//
// Solidity: function gasService() view returns(address)
func (_IScalarERC20CrossChain *IScalarERC20CrossChainCallerSession) GasService() (common.Address, error) {
	return _IScalarERC20CrossChain.Contract.GasService(&_IScalarERC20CrossChain.CallOpts)
}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_IScalarERC20CrossChain *IScalarERC20CrossChainCaller) Gateway(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IScalarERC20CrossChain.contract.Call(opts, &out, "gateway")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_IScalarERC20CrossChain *IScalarERC20CrossChainSession) Gateway() (common.Address, error) {
	return _IScalarERC20CrossChain.Contract.Gateway(&_IScalarERC20CrossChain.CallOpts)
}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_IScalarERC20CrossChain *IScalarERC20CrossChainCallerSession) Gateway() (common.Address, error) {
	return _IScalarERC20CrossChain.Contract.Gateway(&_IScalarERC20CrossChain.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address implementation_)
func (_IScalarERC20CrossChain *IScalarERC20CrossChainCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IScalarERC20CrossChain.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address implementation_)
func (_IScalarERC20CrossChain *IScalarERC20CrossChainSession) Implementation() (common.Address, error) {
	return _IScalarERC20CrossChain.Contract.Implementation(&_IScalarERC20CrossChain.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address implementation_)
func (_IScalarERC20CrossChain *IScalarERC20CrossChainCallerSession) Implementation() (common.Address, error) {
	return _IScalarERC20CrossChain.Contract.Implementation(&_IScalarERC20CrossChain.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IScalarERC20CrossChain *IScalarERC20CrossChainCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IScalarERC20CrossChain.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IScalarERC20CrossChain *IScalarERC20CrossChainSession) Name() (string, error) {
	return _IScalarERC20CrossChain.Contract.Name(&_IScalarERC20CrossChain.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IScalarERC20CrossChain *IScalarERC20CrossChainCallerSession) Name() (string, error) {
	return _IScalarERC20CrossChain.Contract.Name(&_IScalarERC20CrossChain.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address owner_)
func (_IScalarERC20CrossChain *IScalarERC20CrossChainCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IScalarERC20CrossChain.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address owner_)
func (_IScalarERC20CrossChain *IScalarERC20CrossChainSession) Owner() (common.Address, error) {
	return _IScalarERC20CrossChain.Contract.Owner(&_IScalarERC20CrossChain.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address owner_)
func (_IScalarERC20CrossChain *IScalarERC20CrossChainCallerSession) Owner() (common.Address, error) {
	return _IScalarERC20CrossChain.Contract.Owner(&_IScalarERC20CrossChain.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address owner_)
func (_IScalarERC20CrossChain *IScalarERC20CrossChainCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IScalarERC20CrossChain.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address owner_)
func (_IScalarERC20CrossChain *IScalarERC20CrossChainSession) PendingOwner() (common.Address, error) {
	return _IScalarERC20CrossChain.Contract.PendingOwner(&_IScalarERC20CrossChain.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address owner_)
func (_IScalarERC20CrossChain *IScalarERC20CrossChainCallerSession) PendingOwner() (common.Address, error) {
	return _IScalarERC20CrossChain.Contract.PendingOwner(&_IScalarERC20CrossChain.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IScalarERC20CrossChain *IScalarERC20CrossChainCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IScalarERC20CrossChain.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IScalarERC20CrossChain *IScalarERC20CrossChainSession) Symbol() (string, error) {
	return _IScalarERC20CrossChain.Contract.Symbol(&_IScalarERC20CrossChain.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IScalarERC20CrossChain *IScalarERC20CrossChainCallerSession) Symbol() (string, error) {
	return _IScalarERC20CrossChain.Contract.Symbol(&_IScalarERC20CrossChain.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IScalarERC20CrossChain *IScalarERC20CrossChainCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IScalarERC20CrossChain.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IScalarERC20CrossChain *IScalarERC20CrossChainSession) TotalSupply() (*big.Int, error) {
	return _IScalarERC20CrossChain.Contract.TotalSupply(&_IScalarERC20CrossChain.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IScalarERC20CrossChain *IScalarERC20CrossChainCallerSession) TotalSupply() (*big.Int, error) {
	return _IScalarERC20CrossChain.Contract.TotalSupply(&_IScalarERC20CrossChain.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_IScalarERC20CrossChain *IScalarERC20CrossChainTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IScalarERC20CrossChain.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_IScalarERC20CrossChain *IScalarERC20CrossChainSession) AcceptOwnership() (*types.Transaction, error) {
	return _IScalarERC20CrossChain.Contract.AcceptOwnership(&_IScalarERC20CrossChain.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_IScalarERC20CrossChain *IScalarERC20CrossChainTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _IScalarERC20CrossChain.Contract.AcceptOwnership(&_IScalarERC20CrossChain.TransactOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IScalarERC20CrossChain *IScalarERC20CrossChainTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IScalarERC20CrossChain.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IScalarERC20CrossChain *IScalarERC20CrossChainSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IScalarERC20CrossChain.Contract.Approve(&_IScalarERC20CrossChain.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IScalarERC20CrossChain *IScalarERC20CrossChainTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IScalarERC20CrossChain.Contract.Approve(&_IScalarERC20CrossChain.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_IScalarERC20CrossChain *IScalarERC20CrossChainTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _IScalarERC20CrossChain.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_IScalarERC20CrossChain *IScalarERC20CrossChainSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _IScalarERC20CrossChain.Contract.DecreaseAllowance(&_IScalarERC20CrossChain.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_IScalarERC20CrossChain *IScalarERC20CrossChainTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _IScalarERC20CrossChain.Contract.DecreaseAllowance(&_IScalarERC20CrossChain.TransactOpts, spender, subtractedValue)
}

// Execute is a paid mutator transaction binding the contract method 0x49160658.
//
// Solidity: function execute(bytes32 commandId, string sourceChain, string sourceAddress, bytes payload) returns()
func (_IScalarERC20CrossChain *IScalarERC20CrossChainTransactor) Execute(opts *bind.TransactOpts, commandId [32]byte, sourceChain string, sourceAddress string, payload []byte) (*types.Transaction, error) {
	return _IScalarERC20CrossChain.contract.Transact(opts, "execute", commandId, sourceChain, sourceAddress, payload)
}

// Execute is a paid mutator transaction binding the contract method 0x49160658.
//
// Solidity: function execute(bytes32 commandId, string sourceChain, string sourceAddress, bytes payload) returns()
func (_IScalarERC20CrossChain *IScalarERC20CrossChainSession) Execute(commandId [32]byte, sourceChain string, sourceAddress string, payload []byte) (*types.Transaction, error) {
	return _IScalarERC20CrossChain.Contract.Execute(&_IScalarERC20CrossChain.TransactOpts, commandId, sourceChain, sourceAddress, payload)
}

// Execute is a paid mutator transaction binding the contract method 0x49160658.
//
// Solidity: function execute(bytes32 commandId, string sourceChain, string sourceAddress, bytes payload) returns()
func (_IScalarERC20CrossChain *IScalarERC20CrossChainTransactorSession) Execute(commandId [32]byte, sourceChain string, sourceAddress string, payload []byte) (*types.Transaction, error) {
	return _IScalarERC20CrossChain.Contract.Execute(&_IScalarERC20CrossChain.TransactOpts, commandId, sourceChain, sourceAddress, payload)
}

// ExecuteWithToken is a paid mutator transaction binding the contract method 0x1a98b2e0.
//
// Solidity: function executeWithToken(bytes32 commandId, string sourceChain, string sourceAddress, bytes payload, string tokenSymbol, uint256 amount) returns()
func (_IScalarERC20CrossChain *IScalarERC20CrossChainTransactor) ExecuteWithToken(opts *bind.TransactOpts, commandId [32]byte, sourceChain string, sourceAddress string, payload []byte, tokenSymbol string, amount *big.Int) (*types.Transaction, error) {
	return _IScalarERC20CrossChain.contract.Transact(opts, "executeWithToken", commandId, sourceChain, sourceAddress, payload, tokenSymbol, amount)
}

// ExecuteWithToken is a paid mutator transaction binding the contract method 0x1a98b2e0.
//
// Solidity: function executeWithToken(bytes32 commandId, string sourceChain, string sourceAddress, bytes payload, string tokenSymbol, uint256 amount) returns()
func (_IScalarERC20CrossChain *IScalarERC20CrossChainSession) ExecuteWithToken(commandId [32]byte, sourceChain string, sourceAddress string, payload []byte, tokenSymbol string, amount *big.Int) (*types.Transaction, error) {
	return _IScalarERC20CrossChain.Contract.ExecuteWithToken(&_IScalarERC20CrossChain.TransactOpts, commandId, sourceChain, sourceAddress, payload, tokenSymbol, amount)
}

// ExecuteWithToken is a paid mutator transaction binding the contract method 0x1a98b2e0.
//
// Solidity: function executeWithToken(bytes32 commandId, string sourceChain, string sourceAddress, bytes payload, string tokenSymbol, uint256 amount) returns()
func (_IScalarERC20CrossChain *IScalarERC20CrossChainTransactorSession) ExecuteWithToken(commandId [32]byte, sourceChain string, sourceAddress string, payload []byte, tokenSymbol string, amount *big.Int) (*types.Transaction, error) {
	return _IScalarERC20CrossChain.Contract.ExecuteWithToken(&_IScalarERC20CrossChain.TransactOpts, commandId, sourceChain, sourceAddress, payload, tokenSymbol, amount)
}

// Faucet is a paid mutator transaction binding the contract method 0x57915897.
//
// Solidity: function faucet(uint256 amount) returns()
func (_IScalarERC20CrossChain *IScalarERC20CrossChainTransactor) Faucet(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _IScalarERC20CrossChain.contract.Transact(opts, "faucet", amount)
}

// Faucet is a paid mutator transaction binding the contract method 0x57915897.
//
// Solidity: function faucet(uint256 amount) returns()
func (_IScalarERC20CrossChain *IScalarERC20CrossChainSession) Faucet(amount *big.Int) (*types.Transaction, error) {
	return _IScalarERC20CrossChain.Contract.Faucet(&_IScalarERC20CrossChain.TransactOpts, amount)
}

// Faucet is a paid mutator transaction binding the contract method 0x57915897.
//
// Solidity: function faucet(uint256 amount) returns()
func (_IScalarERC20CrossChain *IScalarERC20CrossChainTransactorSession) Faucet(amount *big.Int) (*types.Transaction, error) {
	return _IScalarERC20CrossChain.Contract.Faucet(&_IScalarERC20CrossChain.TransactOpts, amount)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_IScalarERC20CrossChain *IScalarERC20CrossChainTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _IScalarERC20CrossChain.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_IScalarERC20CrossChain *IScalarERC20CrossChainSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _IScalarERC20CrossChain.Contract.IncreaseAllowance(&_IScalarERC20CrossChain.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_IScalarERC20CrossChain *IScalarERC20CrossChainTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _IScalarERC20CrossChain.Contract.IncreaseAllowance(&_IScalarERC20CrossChain.TransactOpts, spender, addedValue)
}

// ProposeOwnership is a paid mutator transaction binding the contract method 0x710bf322.
//
// Solidity: function proposeOwnership(address newOwner) returns()
func (_IScalarERC20CrossChain *IScalarERC20CrossChainTransactor) ProposeOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _IScalarERC20CrossChain.contract.Transact(opts, "proposeOwnership", newOwner)
}

// ProposeOwnership is a paid mutator transaction binding the contract method 0x710bf322.
//
// Solidity: function proposeOwnership(address newOwner) returns()
func (_IScalarERC20CrossChain *IScalarERC20CrossChainSession) ProposeOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _IScalarERC20CrossChain.Contract.ProposeOwnership(&_IScalarERC20CrossChain.TransactOpts, newOwner)
}

// ProposeOwnership is a paid mutator transaction binding the contract method 0x710bf322.
//
// Solidity: function proposeOwnership(address newOwner) returns()
func (_IScalarERC20CrossChain *IScalarERC20CrossChainTransactorSession) ProposeOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _IScalarERC20CrossChain.Contract.ProposeOwnership(&_IScalarERC20CrossChain.TransactOpts, newOwner)
}

// Setup is a paid mutator transaction binding the contract method 0x9ded06df.
//
// Solidity: function setup(bytes data) returns()
func (_IScalarERC20CrossChain *IScalarERC20CrossChainTransactor) Setup(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _IScalarERC20CrossChain.contract.Transact(opts, "setup", data)
}

// Setup is a paid mutator transaction binding the contract method 0x9ded06df.
//
// Solidity: function setup(bytes data) returns()
func (_IScalarERC20CrossChain *IScalarERC20CrossChainSession) Setup(data []byte) (*types.Transaction, error) {
	return _IScalarERC20CrossChain.Contract.Setup(&_IScalarERC20CrossChain.TransactOpts, data)
}

// Setup is a paid mutator transaction binding the contract method 0x9ded06df.
//
// Solidity: function setup(bytes data) returns()
func (_IScalarERC20CrossChain *IScalarERC20CrossChainTransactorSession) Setup(data []byte) (*types.Transaction, error) {
	return _IScalarERC20CrossChain.Contract.Setup(&_IScalarERC20CrossChain.TransactOpts, data)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IScalarERC20CrossChain *IScalarERC20CrossChainTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IScalarERC20CrossChain.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IScalarERC20CrossChain *IScalarERC20CrossChainSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IScalarERC20CrossChain.Contract.Transfer(&_IScalarERC20CrossChain.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IScalarERC20CrossChain *IScalarERC20CrossChainTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IScalarERC20CrossChain.Contract.Transfer(&_IScalarERC20CrossChain.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IScalarERC20CrossChain *IScalarERC20CrossChainTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IScalarERC20CrossChain.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IScalarERC20CrossChain *IScalarERC20CrossChainSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IScalarERC20CrossChain.Contract.TransferFrom(&_IScalarERC20CrossChain.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IScalarERC20CrossChain *IScalarERC20CrossChainTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IScalarERC20CrossChain.Contract.TransferFrom(&_IScalarERC20CrossChain.TransactOpts, sender, recipient, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IScalarERC20CrossChain *IScalarERC20CrossChainTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _IScalarERC20CrossChain.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IScalarERC20CrossChain *IScalarERC20CrossChainSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _IScalarERC20CrossChain.Contract.TransferOwnership(&_IScalarERC20CrossChain.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IScalarERC20CrossChain *IScalarERC20CrossChainTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _IScalarERC20CrossChain.Contract.TransferOwnership(&_IScalarERC20CrossChain.TransactOpts, newOwner)
}

// TransferRemote is a paid mutator transaction binding the contract method 0xc16f0259.
//
// Solidity: function transferRemote(string destinationChain, address destinationContractAddress, uint256 amount, bytes encodedMetadata) payable returns()
func (_IScalarERC20CrossChain *IScalarERC20CrossChainTransactor) TransferRemote(opts *bind.TransactOpts, destinationChain string, destinationContractAddress common.Address, amount *big.Int, encodedMetadata []byte) (*types.Transaction, error) {
	return _IScalarERC20CrossChain.contract.Transact(opts, "transferRemote", destinationChain, destinationContractAddress, amount, encodedMetadata)
}

// TransferRemote is a paid mutator transaction binding the contract method 0xc16f0259.
//
// Solidity: function transferRemote(string destinationChain, address destinationContractAddress, uint256 amount, bytes encodedMetadata) payable returns()
func (_IScalarERC20CrossChain *IScalarERC20CrossChainSession) TransferRemote(destinationChain string, destinationContractAddress common.Address, amount *big.Int, encodedMetadata []byte) (*types.Transaction, error) {
	return _IScalarERC20CrossChain.Contract.TransferRemote(&_IScalarERC20CrossChain.TransactOpts, destinationChain, destinationContractAddress, amount, encodedMetadata)
}

// TransferRemote is a paid mutator transaction binding the contract method 0xc16f0259.
//
// Solidity: function transferRemote(string destinationChain, address destinationContractAddress, uint256 amount, bytes encodedMetadata) payable returns()
func (_IScalarERC20CrossChain *IScalarERC20CrossChainTransactorSession) TransferRemote(destinationChain string, destinationContractAddress common.Address, amount *big.Int, encodedMetadata []byte) (*types.Transaction, error) {
	return _IScalarERC20CrossChain.Contract.TransferRemote(&_IScalarERC20CrossChain.TransactOpts, destinationChain, destinationContractAddress, amount, encodedMetadata)
}

// Upgrade is a paid mutator transaction binding the contract method 0xa3499c73.
//
// Solidity: function upgrade(address newImplementation, bytes32 newImplementationCodeHash, bytes params) returns()
func (_IScalarERC20CrossChain *IScalarERC20CrossChainTransactor) Upgrade(opts *bind.TransactOpts, newImplementation common.Address, newImplementationCodeHash [32]byte, params []byte) (*types.Transaction, error) {
	return _IScalarERC20CrossChain.contract.Transact(opts, "upgrade", newImplementation, newImplementationCodeHash, params)
}

// Upgrade is a paid mutator transaction binding the contract method 0xa3499c73.
//
// Solidity: function upgrade(address newImplementation, bytes32 newImplementationCodeHash, bytes params) returns()
func (_IScalarERC20CrossChain *IScalarERC20CrossChainSession) Upgrade(newImplementation common.Address, newImplementationCodeHash [32]byte, params []byte) (*types.Transaction, error) {
	return _IScalarERC20CrossChain.Contract.Upgrade(&_IScalarERC20CrossChain.TransactOpts, newImplementation, newImplementationCodeHash, params)
}

// Upgrade is a paid mutator transaction binding the contract method 0xa3499c73.
//
// Solidity: function upgrade(address newImplementation, bytes32 newImplementationCodeHash, bytes params) returns()
func (_IScalarERC20CrossChain *IScalarERC20CrossChainTransactorSession) Upgrade(newImplementation common.Address, newImplementationCodeHash [32]byte, params []byte) (*types.Transaction, error) {
	return _IScalarERC20CrossChain.Contract.Upgrade(&_IScalarERC20CrossChain.TransactOpts, newImplementation, newImplementationCodeHash, params)
}

// IScalarERC20CrossChainApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IScalarERC20CrossChain contract.
type IScalarERC20CrossChainApprovalIterator struct {
	Event *IScalarERC20CrossChainApproval // Event containing the contract specifics and raw log

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
func (it *IScalarERC20CrossChainApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IScalarERC20CrossChainApproval)
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
		it.Event = new(IScalarERC20CrossChainApproval)
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
func (it *IScalarERC20CrossChainApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IScalarERC20CrossChainApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IScalarERC20CrossChainApproval represents a Approval event raised by the IScalarERC20CrossChain contract.
type IScalarERC20CrossChainApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IScalarERC20CrossChain *IScalarERC20CrossChainFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IScalarERC20CrossChainApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IScalarERC20CrossChain.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IScalarERC20CrossChainApprovalIterator{contract: _IScalarERC20CrossChain.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IScalarERC20CrossChain *IScalarERC20CrossChainFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IScalarERC20CrossChainApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IScalarERC20CrossChain.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IScalarERC20CrossChainApproval)
				if err := _IScalarERC20CrossChain.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IScalarERC20CrossChain *IScalarERC20CrossChainFilterer) ParseApproval(log types.Log) (*IScalarERC20CrossChainApproval, error) {
	event := new(IScalarERC20CrossChainApproval)
	if err := _IScalarERC20CrossChain.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IScalarERC20CrossChainExecutedIterator is returned from FilterExecuted and is used to iterate over the raw logs and unpacked data for Executed events raised by the IScalarERC20CrossChain contract.
type IScalarERC20CrossChainExecutedIterator struct {
	Event *IScalarERC20CrossChainExecuted // Event containing the contract specifics and raw log

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
func (it *IScalarERC20CrossChainExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IScalarERC20CrossChainExecuted)
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
		it.Event = new(IScalarERC20CrossChainExecuted)
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
func (it *IScalarERC20CrossChainExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IScalarERC20CrossChainExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IScalarERC20CrossChainExecuted represents a Executed event raised by the IScalarERC20CrossChain contract.
type IScalarERC20CrossChainExecuted struct {
	SourceChain   string
	SourceAddress string
	Amount        *big.Int
	SourceTx      [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterExecuted is a free log retrieval operation binding the contract event 0x2be7b106d990d3d787c7682fa446d06333d7d01706d1b4ec26f8a06b638c50ee.
//
// Solidity: event Executed(string sourceChain, string sourceAddress, uint256 amount, bytes32 sourceTx)
func (_IScalarERC20CrossChain *IScalarERC20CrossChainFilterer) FilterExecuted(opts *bind.FilterOpts) (*IScalarERC20CrossChainExecutedIterator, error) {

	logs, sub, err := _IScalarERC20CrossChain.contract.FilterLogs(opts, "Executed")
	if err != nil {
		return nil, err
	}
	return &IScalarERC20CrossChainExecutedIterator{contract: _IScalarERC20CrossChain.contract, event: "Executed", logs: logs, sub: sub}, nil
}

// WatchExecuted is a free log subscription operation binding the contract event 0x2be7b106d990d3d787c7682fa446d06333d7d01706d1b4ec26f8a06b638c50ee.
//
// Solidity: event Executed(string sourceChain, string sourceAddress, uint256 amount, bytes32 sourceTx)
func (_IScalarERC20CrossChain *IScalarERC20CrossChainFilterer) WatchExecuted(opts *bind.WatchOpts, sink chan<- *IScalarERC20CrossChainExecuted) (event.Subscription, error) {

	logs, sub, err := _IScalarERC20CrossChain.contract.WatchLogs(opts, "Executed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IScalarERC20CrossChainExecuted)
				if err := _IScalarERC20CrossChain.contract.UnpackLog(event, "Executed", log); err != nil {
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

// ParseExecuted is a log parse operation binding the contract event 0x2be7b106d990d3d787c7682fa446d06333d7d01706d1b4ec26f8a06b638c50ee.
//
// Solidity: event Executed(string sourceChain, string sourceAddress, uint256 amount, bytes32 sourceTx)
func (_IScalarERC20CrossChain *IScalarERC20CrossChainFilterer) ParseExecuted(log types.Log) (*IScalarERC20CrossChainExecuted, error) {
	event := new(IScalarERC20CrossChainExecuted)
	if err := _IScalarERC20CrossChain.contract.UnpackLog(event, "Executed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IScalarERC20CrossChainFalseSenderIterator is returned from FilterFalseSender and is used to iterate over the raw logs and unpacked data for FalseSender events raised by the IScalarERC20CrossChain contract.
type IScalarERC20CrossChainFalseSenderIterator struct {
	Event *IScalarERC20CrossChainFalseSender // Event containing the contract specifics and raw log

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
func (it *IScalarERC20CrossChainFalseSenderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IScalarERC20CrossChainFalseSender)
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
		it.Event = new(IScalarERC20CrossChainFalseSender)
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
func (it *IScalarERC20CrossChainFalseSenderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IScalarERC20CrossChainFalseSenderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IScalarERC20CrossChainFalseSender represents a FalseSender event raised by the IScalarERC20CrossChain contract.
type IScalarERC20CrossChainFalseSender struct {
	SourceChain   string
	SourceAddress string
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterFalseSender is a free log retrieval operation binding the contract event 0x03daa0914739c74a32f06993c062b0ec6a2c0b8a39b5a0b8bff2a12ef55b9dc5.
//
// Solidity: event FalseSender(string sourceChain, string sourceAddress)
func (_IScalarERC20CrossChain *IScalarERC20CrossChainFilterer) FilterFalseSender(opts *bind.FilterOpts) (*IScalarERC20CrossChainFalseSenderIterator, error) {

	logs, sub, err := _IScalarERC20CrossChain.contract.FilterLogs(opts, "FalseSender")
	if err != nil {
		return nil, err
	}
	return &IScalarERC20CrossChainFalseSenderIterator{contract: _IScalarERC20CrossChain.contract, event: "FalseSender", logs: logs, sub: sub}, nil
}

// WatchFalseSender is a free log subscription operation binding the contract event 0x03daa0914739c74a32f06993c062b0ec6a2c0b8a39b5a0b8bff2a12ef55b9dc5.
//
// Solidity: event FalseSender(string sourceChain, string sourceAddress)
func (_IScalarERC20CrossChain *IScalarERC20CrossChainFilterer) WatchFalseSender(opts *bind.WatchOpts, sink chan<- *IScalarERC20CrossChainFalseSender) (event.Subscription, error) {

	logs, sub, err := _IScalarERC20CrossChain.contract.WatchLogs(opts, "FalseSender")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IScalarERC20CrossChainFalseSender)
				if err := _IScalarERC20CrossChain.contract.UnpackLog(event, "FalseSender", log); err != nil {
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

// ParseFalseSender is a log parse operation binding the contract event 0x03daa0914739c74a32f06993c062b0ec6a2c0b8a39b5a0b8bff2a12ef55b9dc5.
//
// Solidity: event FalseSender(string sourceChain, string sourceAddress)
func (_IScalarERC20CrossChain *IScalarERC20CrossChainFilterer) ParseFalseSender(log types.Log) (*IScalarERC20CrossChainFalseSender, error) {
	event := new(IScalarERC20CrossChainFalseSender)
	if err := _IScalarERC20CrossChain.contract.UnpackLog(event, "FalseSender", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IScalarERC20CrossChainOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the IScalarERC20CrossChain contract.
type IScalarERC20CrossChainOwnershipTransferStartedIterator struct {
	Event *IScalarERC20CrossChainOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *IScalarERC20CrossChainOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IScalarERC20CrossChainOwnershipTransferStarted)
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
		it.Event = new(IScalarERC20CrossChainOwnershipTransferStarted)
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
func (it *IScalarERC20CrossChainOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IScalarERC20CrossChainOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IScalarERC20CrossChainOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the IScalarERC20CrossChain contract.
type IScalarERC20CrossChainOwnershipTransferStarted struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0xd9be0e8e07417e00f2521db636cb53e316fd288f5051f16d2aa2bf0c3938a876.
//
// Solidity: event OwnershipTransferStarted(address indexed newOwner)
func (_IScalarERC20CrossChain *IScalarERC20CrossChainFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, newOwner []common.Address) (*IScalarERC20CrossChainOwnershipTransferStartedIterator, error) {

	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _IScalarERC20CrossChain.contract.FilterLogs(opts, "OwnershipTransferStarted", newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &IScalarERC20CrossChainOwnershipTransferStartedIterator{contract: _IScalarERC20CrossChain.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0xd9be0e8e07417e00f2521db636cb53e316fd288f5051f16d2aa2bf0c3938a876.
//
// Solidity: event OwnershipTransferStarted(address indexed newOwner)
func (_IScalarERC20CrossChain *IScalarERC20CrossChainFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *IScalarERC20CrossChainOwnershipTransferStarted, newOwner []common.Address) (event.Subscription, error) {

	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _IScalarERC20CrossChain.contract.WatchLogs(opts, "OwnershipTransferStarted", newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IScalarERC20CrossChainOwnershipTransferStarted)
				if err := _IScalarERC20CrossChain.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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

// ParseOwnershipTransferStarted is a log parse operation binding the contract event 0xd9be0e8e07417e00f2521db636cb53e316fd288f5051f16d2aa2bf0c3938a876.
//
// Solidity: event OwnershipTransferStarted(address indexed newOwner)
func (_IScalarERC20CrossChain *IScalarERC20CrossChainFilterer) ParseOwnershipTransferStarted(log types.Log) (*IScalarERC20CrossChainOwnershipTransferStarted, error) {
	event := new(IScalarERC20CrossChainOwnershipTransferStarted)
	if err := _IScalarERC20CrossChain.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IScalarERC20CrossChainOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the IScalarERC20CrossChain contract.
type IScalarERC20CrossChainOwnershipTransferredIterator struct {
	Event *IScalarERC20CrossChainOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *IScalarERC20CrossChainOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IScalarERC20CrossChainOwnershipTransferred)
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
		it.Event = new(IScalarERC20CrossChainOwnershipTransferred)
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
func (it *IScalarERC20CrossChainOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IScalarERC20CrossChainOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IScalarERC20CrossChainOwnershipTransferred represents a OwnershipTransferred event raised by the IScalarERC20CrossChain contract.
type IScalarERC20CrossChainOwnershipTransferred struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x04dba622d284ed0014ee4b9a6a68386be1a4c08a4913ae272de89199cc686163.
//
// Solidity: event OwnershipTransferred(address indexed newOwner)
func (_IScalarERC20CrossChain *IScalarERC20CrossChainFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, newOwner []common.Address) (*IScalarERC20CrossChainOwnershipTransferredIterator, error) {

	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _IScalarERC20CrossChain.contract.FilterLogs(opts, "OwnershipTransferred", newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &IScalarERC20CrossChainOwnershipTransferredIterator{contract: _IScalarERC20CrossChain.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x04dba622d284ed0014ee4b9a6a68386be1a4c08a4913ae272de89199cc686163.
//
// Solidity: event OwnershipTransferred(address indexed newOwner)
func (_IScalarERC20CrossChain *IScalarERC20CrossChainFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *IScalarERC20CrossChainOwnershipTransferred, newOwner []common.Address) (event.Subscription, error) {

	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _IScalarERC20CrossChain.contract.WatchLogs(opts, "OwnershipTransferred", newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IScalarERC20CrossChainOwnershipTransferred)
				if err := _IScalarERC20CrossChain.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x04dba622d284ed0014ee4b9a6a68386be1a4c08a4913ae272de89199cc686163.
//
// Solidity: event OwnershipTransferred(address indexed newOwner)
func (_IScalarERC20CrossChain *IScalarERC20CrossChainFilterer) ParseOwnershipTransferred(log types.Log) (*IScalarERC20CrossChainOwnershipTransferred, error) {
	event := new(IScalarERC20CrossChainOwnershipTransferred)
	if err := _IScalarERC20CrossChain.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IScalarERC20CrossChainTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IScalarERC20CrossChain contract.
type IScalarERC20CrossChainTransferIterator struct {
	Event *IScalarERC20CrossChainTransfer // Event containing the contract specifics and raw log

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
func (it *IScalarERC20CrossChainTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IScalarERC20CrossChainTransfer)
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
		it.Event = new(IScalarERC20CrossChainTransfer)
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
func (it *IScalarERC20CrossChainTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IScalarERC20CrossChainTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IScalarERC20CrossChainTransfer represents a Transfer event raised by the IScalarERC20CrossChain contract.
type IScalarERC20CrossChainTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IScalarERC20CrossChain *IScalarERC20CrossChainFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IScalarERC20CrossChainTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IScalarERC20CrossChain.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IScalarERC20CrossChainTransferIterator{contract: _IScalarERC20CrossChain.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IScalarERC20CrossChain *IScalarERC20CrossChainFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IScalarERC20CrossChainTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IScalarERC20CrossChain.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IScalarERC20CrossChainTransfer)
				if err := _IScalarERC20CrossChain.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IScalarERC20CrossChain *IScalarERC20CrossChainFilterer) ParseTransfer(log types.Log) (*IScalarERC20CrossChainTransfer, error) {
	event := new(IScalarERC20CrossChainTransfer)
	if err := _IScalarERC20CrossChain.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IScalarERC20CrossChainTransferRemoteIterator is returned from FilterTransferRemote and is used to iterate over the raw logs and unpacked data for TransferRemote events raised by the IScalarERC20CrossChain contract.
type IScalarERC20CrossChainTransferRemoteIterator struct {
	Event *IScalarERC20CrossChainTransferRemote // Event containing the contract specifics and raw log

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
func (it *IScalarERC20CrossChainTransferRemoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IScalarERC20CrossChainTransferRemote)
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
		it.Event = new(IScalarERC20CrossChainTransferRemote)
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
func (it *IScalarERC20CrossChainTransferRemoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IScalarERC20CrossChainTransferRemoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IScalarERC20CrossChainTransferRemote represents a TransferRemote event raised by the IScalarERC20CrossChain contract.
type IScalarERC20CrossChainTransferRemote struct {
	DestinationChain           string
	DestinationContractAddress common.Address
	Sender                     common.Address
	Amount                     *big.Int
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterTransferRemote is a free log retrieval operation binding the contract event 0xc4147435f095e6892364936c3909f49a00fc612968e039d109a8284011253036.
//
// Solidity: event TransferRemote(string destinationChain, address destinationContractAddress, address sender, uint256 amount)
func (_IScalarERC20CrossChain *IScalarERC20CrossChainFilterer) FilterTransferRemote(opts *bind.FilterOpts) (*IScalarERC20CrossChainTransferRemoteIterator, error) {

	logs, sub, err := _IScalarERC20CrossChain.contract.FilterLogs(opts, "TransferRemote")
	if err != nil {
		return nil, err
	}
	return &IScalarERC20CrossChainTransferRemoteIterator{contract: _IScalarERC20CrossChain.contract, event: "TransferRemote", logs: logs, sub: sub}, nil
}

// WatchTransferRemote is a free log subscription operation binding the contract event 0xc4147435f095e6892364936c3909f49a00fc612968e039d109a8284011253036.
//
// Solidity: event TransferRemote(string destinationChain, address destinationContractAddress, address sender, uint256 amount)
func (_IScalarERC20CrossChain *IScalarERC20CrossChainFilterer) WatchTransferRemote(opts *bind.WatchOpts, sink chan<- *IScalarERC20CrossChainTransferRemote) (event.Subscription, error) {

	logs, sub, err := _IScalarERC20CrossChain.contract.WatchLogs(opts, "TransferRemote")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IScalarERC20CrossChainTransferRemote)
				if err := _IScalarERC20CrossChain.contract.UnpackLog(event, "TransferRemote", log); err != nil {
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

// ParseTransferRemote is a log parse operation binding the contract event 0xc4147435f095e6892364936c3909f49a00fc612968e039d109a8284011253036.
//
// Solidity: event TransferRemote(string destinationChain, address destinationContractAddress, address sender, uint256 amount)
func (_IScalarERC20CrossChain *IScalarERC20CrossChainFilterer) ParseTransferRemote(log types.Log) (*IScalarERC20CrossChainTransferRemote, error) {
	event := new(IScalarERC20CrossChainTransferRemote)
	if err := _IScalarERC20CrossChain.contract.UnpackLog(event, "TransferRemote", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IScalarERC20CrossChainUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the IScalarERC20CrossChain contract.
type IScalarERC20CrossChainUpgradedIterator struct {
	Event *IScalarERC20CrossChainUpgraded // Event containing the contract specifics and raw log

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
func (it *IScalarERC20CrossChainUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IScalarERC20CrossChainUpgraded)
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
		it.Event = new(IScalarERC20CrossChainUpgraded)
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
func (it *IScalarERC20CrossChainUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IScalarERC20CrossChainUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IScalarERC20CrossChainUpgraded represents a Upgraded event raised by the IScalarERC20CrossChain contract.
type IScalarERC20CrossChainUpgraded struct {
	NewImplementation common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed newImplementation)
func (_IScalarERC20CrossChain *IScalarERC20CrossChainFilterer) FilterUpgraded(opts *bind.FilterOpts, newImplementation []common.Address) (*IScalarERC20CrossChainUpgradedIterator, error) {

	var newImplementationRule []interface{}
	for _, newImplementationItem := range newImplementation {
		newImplementationRule = append(newImplementationRule, newImplementationItem)
	}

	logs, sub, err := _IScalarERC20CrossChain.contract.FilterLogs(opts, "Upgraded", newImplementationRule)
	if err != nil {
		return nil, err
	}
	return &IScalarERC20CrossChainUpgradedIterator{contract: _IScalarERC20CrossChain.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed newImplementation)
func (_IScalarERC20CrossChain *IScalarERC20CrossChainFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *IScalarERC20CrossChainUpgraded, newImplementation []common.Address) (event.Subscription, error) {

	var newImplementationRule []interface{}
	for _, newImplementationItem := range newImplementation {
		newImplementationRule = append(newImplementationRule, newImplementationItem)
	}

	logs, sub, err := _IScalarERC20CrossChain.contract.WatchLogs(opts, "Upgraded", newImplementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IScalarERC20CrossChainUpgraded)
				if err := _IScalarERC20CrossChain.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
// Solidity: event Upgraded(address indexed newImplementation)
func (_IScalarERC20CrossChain *IScalarERC20CrossChainFilterer) ParseUpgraded(log types.Log) (*IScalarERC20CrossChainUpgraded, error) {
	event := new(IScalarERC20CrossChainUpgraded)
	if err := _IScalarERC20CrossChain.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
