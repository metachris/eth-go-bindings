// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc777

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
)

// Erc777MetaData contains all meta data concerning the Erc777 contract.
var Erc777MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenHolder\",\"type\":\"address\"}],\"name\":\"AuthorizedOperator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"operatorData\",\"type\":\"bytes\"}],\"name\":\"Burned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"operatorData\",\"type\":\"bytes\"}],\"name\":\"Minted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenHolder\",\"type\":\"address\"}],\"name\":\"RevokedOperator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"operatorData\",\"type\":\"bytes\"}],\"name\":\"Sent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"authorizeOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenHolder\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultOperators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"granularity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenHolder\",\"type\":\"address\"}],\"name\":\"isOperatorFor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"operatorData\",\"type\":\"bytes\"}],\"name\":\"operatorBurn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"operatorData\",\"type\":\"bytes\"}],\"name\":\"operatorSend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"revokeOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"send\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// Erc777ABI is the input ABI used to generate the binding from.
// Deprecated: Use Erc777MetaData.ABI instead.
var Erc777ABI = Erc777MetaData.ABI

// Erc777 is an auto generated Go binding around an Ethereum contract.
type Erc777 struct {
	Erc777Caller     // Read-only binding to the contract
	Erc777Transactor // Write-only binding to the contract
	Erc777Filterer   // Log filterer for contract events
}

// Erc777Caller is an auto generated read-only Go binding around an Ethereum contract.
type Erc777Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc777Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Erc777Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc777Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Erc777Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc777Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Erc777Session struct {
	Contract     *Erc777           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Erc777CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Erc777CallerSession struct {
	Contract *Erc777Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// Erc777TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Erc777TransactorSession struct {
	Contract     *Erc777Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Erc777Raw is an auto generated low-level Go binding around an Ethereum contract.
type Erc777Raw struct {
	Contract *Erc777 // Generic contract binding to access the raw methods on
}

// Erc777CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Erc777CallerRaw struct {
	Contract *Erc777Caller // Generic read-only contract binding to access the raw methods on
}

// Erc777TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Erc777TransactorRaw struct {
	Contract *Erc777Transactor // Generic write-only contract binding to access the raw methods on
}

// NewErc777 creates a new instance of Erc777, bound to a specific deployed contract.
func NewErc777(address common.Address, backend bind.ContractBackend) (*Erc777, error) {
	contract, err := bindErc777(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Erc777{Erc777Caller: Erc777Caller{contract: contract}, Erc777Transactor: Erc777Transactor{contract: contract}, Erc777Filterer: Erc777Filterer{contract: contract}}, nil
}

// NewErc777Caller creates a new read-only instance of Erc777, bound to a specific deployed contract.
func NewErc777Caller(address common.Address, caller bind.ContractCaller) (*Erc777Caller, error) {
	contract, err := bindErc777(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Erc777Caller{contract: contract}, nil
}

// NewErc777Transactor creates a new write-only instance of Erc777, bound to a specific deployed contract.
func NewErc777Transactor(address common.Address, transactor bind.ContractTransactor) (*Erc777Transactor, error) {
	contract, err := bindErc777(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Erc777Transactor{contract: contract}, nil
}

// NewErc777Filterer creates a new log filterer instance of Erc777, bound to a specific deployed contract.
func NewErc777Filterer(address common.Address, filterer bind.ContractFilterer) (*Erc777Filterer, error) {
	contract, err := bindErc777(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Erc777Filterer{contract: contract}, nil
}

// bindErc777 binds a generic wrapper to an already deployed contract.
func bindErc777(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Erc777ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc777 *Erc777Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc777.Contract.Erc777Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc777 *Erc777Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc777.Contract.Erc777Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc777 *Erc777Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc777.Contract.Erc777Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc777 *Erc777CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc777.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc777 *Erc777TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc777.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc777 *Erc777TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc777.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address holder, address spender) view returns(uint256)
func (_Erc777 *Erc777Caller) Allowance(opts *bind.CallOpts, holder common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Erc777.contract.Call(opts, &out, "allowance", holder, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address holder, address spender) view returns(uint256)
func (_Erc777 *Erc777Session) Allowance(holder common.Address, spender common.Address) (*big.Int, error) {
	return _Erc777.Contract.Allowance(&_Erc777.CallOpts, holder, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address holder, address spender) view returns(uint256)
func (_Erc777 *Erc777CallerSession) Allowance(holder common.Address, spender common.Address) (*big.Int, error) {
	return _Erc777.Contract.Allowance(&_Erc777.CallOpts, holder, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address tokenHolder) view returns(uint256)
func (_Erc777 *Erc777Caller) BalanceOf(opts *bind.CallOpts, tokenHolder common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Erc777.contract.Call(opts, &out, "balanceOf", tokenHolder)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address tokenHolder) view returns(uint256)
func (_Erc777 *Erc777Session) BalanceOf(tokenHolder common.Address) (*big.Int, error) {
	return _Erc777.Contract.BalanceOf(&_Erc777.CallOpts, tokenHolder)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address tokenHolder) view returns(uint256)
func (_Erc777 *Erc777CallerSession) BalanceOf(tokenHolder common.Address) (*big.Int, error) {
	return _Erc777.Contract.BalanceOf(&_Erc777.CallOpts, tokenHolder)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_Erc777 *Erc777Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Erc777.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_Erc777 *Erc777Session) Decimals() (uint8, error) {
	return _Erc777.Contract.Decimals(&_Erc777.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_Erc777 *Erc777CallerSession) Decimals() (uint8, error) {
	return _Erc777.Contract.Decimals(&_Erc777.CallOpts)
}

// DefaultOperators is a free data retrieval call binding the contract method 0x06e48538.
//
// Solidity: function defaultOperators() view returns(address[])
func (_Erc777 *Erc777Caller) DefaultOperators(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Erc777.contract.Call(opts, &out, "defaultOperators")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// DefaultOperators is a free data retrieval call binding the contract method 0x06e48538.
//
// Solidity: function defaultOperators() view returns(address[])
func (_Erc777 *Erc777Session) DefaultOperators() ([]common.Address, error) {
	return _Erc777.Contract.DefaultOperators(&_Erc777.CallOpts)
}

// DefaultOperators is a free data retrieval call binding the contract method 0x06e48538.
//
// Solidity: function defaultOperators() view returns(address[])
func (_Erc777 *Erc777CallerSession) DefaultOperators() ([]common.Address, error) {
	return _Erc777.Contract.DefaultOperators(&_Erc777.CallOpts)
}

// Granularity is a free data retrieval call binding the contract method 0x556f0dc7.
//
// Solidity: function granularity() view returns(uint256)
func (_Erc777 *Erc777Caller) Granularity(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc777.contract.Call(opts, &out, "granularity")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Granularity is a free data retrieval call binding the contract method 0x556f0dc7.
//
// Solidity: function granularity() view returns(uint256)
func (_Erc777 *Erc777Session) Granularity() (*big.Int, error) {
	return _Erc777.Contract.Granularity(&_Erc777.CallOpts)
}

// Granularity is a free data retrieval call binding the contract method 0x556f0dc7.
//
// Solidity: function granularity() view returns(uint256)
func (_Erc777 *Erc777CallerSession) Granularity() (*big.Int, error) {
	return _Erc777.Contract.Granularity(&_Erc777.CallOpts)
}

// IsOperatorFor is a free data retrieval call binding the contract method 0xd95b6371.
//
// Solidity: function isOperatorFor(address operator, address tokenHolder) view returns(bool)
func (_Erc777 *Erc777Caller) IsOperatorFor(opts *bind.CallOpts, operator common.Address, tokenHolder common.Address) (bool, error) {
	var out []interface{}
	err := _Erc777.contract.Call(opts, &out, "isOperatorFor", operator, tokenHolder)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperatorFor is a free data retrieval call binding the contract method 0xd95b6371.
//
// Solidity: function isOperatorFor(address operator, address tokenHolder) view returns(bool)
func (_Erc777 *Erc777Session) IsOperatorFor(operator common.Address, tokenHolder common.Address) (bool, error) {
	return _Erc777.Contract.IsOperatorFor(&_Erc777.CallOpts, operator, tokenHolder)
}

// IsOperatorFor is a free data retrieval call binding the contract method 0xd95b6371.
//
// Solidity: function isOperatorFor(address operator, address tokenHolder) view returns(bool)
func (_Erc777 *Erc777CallerSession) IsOperatorFor(operator common.Address, tokenHolder common.Address) (bool, error) {
	return _Erc777.Contract.IsOperatorFor(&_Erc777.CallOpts, operator, tokenHolder)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Erc777 *Erc777Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Erc777.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Erc777 *Erc777Session) Name() (string, error) {
	return _Erc777.Contract.Name(&_Erc777.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Erc777 *Erc777CallerSession) Name() (string, error) {
	return _Erc777.Contract.Name(&_Erc777.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Erc777 *Erc777Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Erc777.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Erc777 *Erc777Session) Symbol() (string, error) {
	return _Erc777.Contract.Symbol(&_Erc777.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Erc777 *Erc777CallerSession) Symbol() (string, error) {
	return _Erc777.Contract.Symbol(&_Erc777.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Erc777 *Erc777Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc777.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Erc777 *Erc777Session) TotalSupply() (*big.Int, error) {
	return _Erc777.Contract.TotalSupply(&_Erc777.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Erc777 *Erc777CallerSession) TotalSupply() (*big.Int, error) {
	return _Erc777.Contract.TotalSupply(&_Erc777.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Erc777 *Erc777Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Erc777.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Erc777 *Erc777Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Erc777.Contract.Approve(&_Erc777.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Erc777 *Erc777TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Erc777.Contract.Approve(&_Erc777.TransactOpts, spender, value)
}

// AuthorizeOperator is a paid mutator transaction binding the contract method 0x959b8c3f.
//
// Solidity: function authorizeOperator(address operator) returns()
func (_Erc777 *Erc777Transactor) AuthorizeOperator(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _Erc777.contract.Transact(opts, "authorizeOperator", operator)
}

// AuthorizeOperator is a paid mutator transaction binding the contract method 0x959b8c3f.
//
// Solidity: function authorizeOperator(address operator) returns()
func (_Erc777 *Erc777Session) AuthorizeOperator(operator common.Address) (*types.Transaction, error) {
	return _Erc777.Contract.AuthorizeOperator(&_Erc777.TransactOpts, operator)
}

// AuthorizeOperator is a paid mutator transaction binding the contract method 0x959b8c3f.
//
// Solidity: function authorizeOperator(address operator) returns()
func (_Erc777 *Erc777TransactorSession) AuthorizeOperator(operator common.Address) (*types.Transaction, error) {
	return _Erc777.Contract.AuthorizeOperator(&_Erc777.TransactOpts, operator)
}

// Burn is a paid mutator transaction binding the contract method 0xfe9d9303.
//
// Solidity: function burn(uint256 amount, bytes data) returns()
func (_Erc777 *Erc777Transactor) Burn(opts *bind.TransactOpts, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Erc777.contract.Transact(opts, "burn", amount, data)
}

// Burn is a paid mutator transaction binding the contract method 0xfe9d9303.
//
// Solidity: function burn(uint256 amount, bytes data) returns()
func (_Erc777 *Erc777Session) Burn(amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Erc777.Contract.Burn(&_Erc777.TransactOpts, amount, data)
}

// Burn is a paid mutator transaction binding the contract method 0xfe9d9303.
//
// Solidity: function burn(uint256 amount, bytes data) returns()
func (_Erc777 *Erc777TransactorSession) Burn(amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Erc777.Contract.Burn(&_Erc777.TransactOpts, amount, data)
}

// OperatorBurn is a paid mutator transaction binding the contract method 0xfc673c4f.
//
// Solidity: function operatorBurn(address account, uint256 amount, bytes data, bytes operatorData) returns()
func (_Erc777 *Erc777Transactor) OperatorBurn(opts *bind.TransactOpts, account common.Address, amount *big.Int, data []byte, operatorData []byte) (*types.Transaction, error) {
	return _Erc777.contract.Transact(opts, "operatorBurn", account, amount, data, operatorData)
}

// OperatorBurn is a paid mutator transaction binding the contract method 0xfc673c4f.
//
// Solidity: function operatorBurn(address account, uint256 amount, bytes data, bytes operatorData) returns()
func (_Erc777 *Erc777Session) OperatorBurn(account common.Address, amount *big.Int, data []byte, operatorData []byte) (*types.Transaction, error) {
	return _Erc777.Contract.OperatorBurn(&_Erc777.TransactOpts, account, amount, data, operatorData)
}

// OperatorBurn is a paid mutator transaction binding the contract method 0xfc673c4f.
//
// Solidity: function operatorBurn(address account, uint256 amount, bytes data, bytes operatorData) returns()
func (_Erc777 *Erc777TransactorSession) OperatorBurn(account common.Address, amount *big.Int, data []byte, operatorData []byte) (*types.Transaction, error) {
	return _Erc777.Contract.OperatorBurn(&_Erc777.TransactOpts, account, amount, data, operatorData)
}

// OperatorSend is a paid mutator transaction binding the contract method 0x62ad1b83.
//
// Solidity: function operatorSend(address sender, address recipient, uint256 amount, bytes data, bytes operatorData) returns()
func (_Erc777 *Erc777Transactor) OperatorSend(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int, data []byte, operatorData []byte) (*types.Transaction, error) {
	return _Erc777.contract.Transact(opts, "operatorSend", sender, recipient, amount, data, operatorData)
}

// OperatorSend is a paid mutator transaction binding the contract method 0x62ad1b83.
//
// Solidity: function operatorSend(address sender, address recipient, uint256 amount, bytes data, bytes operatorData) returns()
func (_Erc777 *Erc777Session) OperatorSend(sender common.Address, recipient common.Address, amount *big.Int, data []byte, operatorData []byte) (*types.Transaction, error) {
	return _Erc777.Contract.OperatorSend(&_Erc777.TransactOpts, sender, recipient, amount, data, operatorData)
}

// OperatorSend is a paid mutator transaction binding the contract method 0x62ad1b83.
//
// Solidity: function operatorSend(address sender, address recipient, uint256 amount, bytes data, bytes operatorData) returns()
func (_Erc777 *Erc777TransactorSession) OperatorSend(sender common.Address, recipient common.Address, amount *big.Int, data []byte, operatorData []byte) (*types.Transaction, error) {
	return _Erc777.Contract.OperatorSend(&_Erc777.TransactOpts, sender, recipient, amount, data, operatorData)
}

// RevokeOperator is a paid mutator transaction binding the contract method 0xfad8b32a.
//
// Solidity: function revokeOperator(address operator) returns()
func (_Erc777 *Erc777Transactor) RevokeOperator(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _Erc777.contract.Transact(opts, "revokeOperator", operator)
}

// RevokeOperator is a paid mutator transaction binding the contract method 0xfad8b32a.
//
// Solidity: function revokeOperator(address operator) returns()
func (_Erc777 *Erc777Session) RevokeOperator(operator common.Address) (*types.Transaction, error) {
	return _Erc777.Contract.RevokeOperator(&_Erc777.TransactOpts, operator)
}

// RevokeOperator is a paid mutator transaction binding the contract method 0xfad8b32a.
//
// Solidity: function revokeOperator(address operator) returns()
func (_Erc777 *Erc777TransactorSession) RevokeOperator(operator common.Address) (*types.Transaction, error) {
	return _Erc777.Contract.RevokeOperator(&_Erc777.TransactOpts, operator)
}

// Send is a paid mutator transaction binding the contract method 0x9bd9bbc6.
//
// Solidity: function send(address recipient, uint256 amount, bytes data) returns()
func (_Erc777 *Erc777Transactor) Send(opts *bind.TransactOpts, recipient common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Erc777.contract.Transact(opts, "send", recipient, amount, data)
}

// Send is a paid mutator transaction binding the contract method 0x9bd9bbc6.
//
// Solidity: function send(address recipient, uint256 amount, bytes data) returns()
func (_Erc777 *Erc777Session) Send(recipient common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Erc777.Contract.Send(&_Erc777.TransactOpts, recipient, amount, data)
}

// Send is a paid mutator transaction binding the contract method 0x9bd9bbc6.
//
// Solidity: function send(address recipient, uint256 amount, bytes data) returns()
func (_Erc777 *Erc777TransactorSession) Send(recipient common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Erc777.Contract.Send(&_Erc777.TransactOpts, recipient, amount, data)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Erc777 *Erc777Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc777.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Erc777 *Erc777Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc777.Contract.Transfer(&_Erc777.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Erc777 *Erc777TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc777.Contract.Transfer(&_Erc777.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address holder, address recipient, uint256 amount) returns(bool)
func (_Erc777 *Erc777Transactor) TransferFrom(opts *bind.TransactOpts, holder common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc777.contract.Transact(opts, "transferFrom", holder, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address holder, address recipient, uint256 amount) returns(bool)
func (_Erc777 *Erc777Session) TransferFrom(holder common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc777.Contract.TransferFrom(&_Erc777.TransactOpts, holder, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address holder, address recipient, uint256 amount) returns(bool)
func (_Erc777 *Erc777TransactorSession) TransferFrom(holder common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc777.Contract.TransferFrom(&_Erc777.TransactOpts, holder, recipient, amount)
}

// Erc777ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Erc777 contract.
type Erc777ApprovalIterator struct {
	Event *Erc777Approval // Event containing the contract specifics and raw log

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
func (it *Erc777ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc777Approval)
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
		it.Event = new(Erc777Approval)
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
func (it *Erc777ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc777ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc777Approval represents a Approval event raised by the Erc777 contract.
type Erc777Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Erc777 *Erc777Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*Erc777ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Erc777.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &Erc777ApprovalIterator{contract: _Erc777.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Erc777 *Erc777Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *Erc777Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Erc777.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc777Approval)
				if err := _Erc777.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Erc777 *Erc777Filterer) ParseApproval(log types.Log) (*Erc777Approval, error) {
	event := new(Erc777Approval)
	if err := _Erc777.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc777AuthorizedOperatorIterator is returned from FilterAuthorizedOperator and is used to iterate over the raw logs and unpacked data for AuthorizedOperator events raised by the Erc777 contract.
type Erc777AuthorizedOperatorIterator struct {
	Event *Erc777AuthorizedOperator // Event containing the contract specifics and raw log

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
func (it *Erc777AuthorizedOperatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc777AuthorizedOperator)
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
		it.Event = new(Erc777AuthorizedOperator)
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
func (it *Erc777AuthorizedOperatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc777AuthorizedOperatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc777AuthorizedOperator represents a AuthorizedOperator event raised by the Erc777 contract.
type Erc777AuthorizedOperator struct {
	Operator    common.Address
	TokenHolder common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAuthorizedOperator is a free log retrieval operation binding the contract event 0xf4caeb2d6ca8932a215a353d0703c326ec2d81fc68170f320eb2ab49e9df61f9.
//
// Solidity: event AuthorizedOperator(address indexed operator, address indexed tokenHolder)
func (_Erc777 *Erc777Filterer) FilterAuthorizedOperator(opts *bind.FilterOpts, operator []common.Address, tokenHolder []common.Address) (*Erc777AuthorizedOperatorIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var tokenHolderRule []interface{}
	for _, tokenHolderItem := range tokenHolder {
		tokenHolderRule = append(tokenHolderRule, tokenHolderItem)
	}

	logs, sub, err := _Erc777.contract.FilterLogs(opts, "AuthorizedOperator", operatorRule, tokenHolderRule)
	if err != nil {
		return nil, err
	}
	return &Erc777AuthorizedOperatorIterator{contract: _Erc777.contract, event: "AuthorizedOperator", logs: logs, sub: sub}, nil
}

// WatchAuthorizedOperator is a free log subscription operation binding the contract event 0xf4caeb2d6ca8932a215a353d0703c326ec2d81fc68170f320eb2ab49e9df61f9.
//
// Solidity: event AuthorizedOperator(address indexed operator, address indexed tokenHolder)
func (_Erc777 *Erc777Filterer) WatchAuthorizedOperator(opts *bind.WatchOpts, sink chan<- *Erc777AuthorizedOperator, operator []common.Address, tokenHolder []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var tokenHolderRule []interface{}
	for _, tokenHolderItem := range tokenHolder {
		tokenHolderRule = append(tokenHolderRule, tokenHolderItem)
	}

	logs, sub, err := _Erc777.contract.WatchLogs(opts, "AuthorizedOperator", operatorRule, tokenHolderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc777AuthorizedOperator)
				if err := _Erc777.contract.UnpackLog(event, "AuthorizedOperator", log); err != nil {
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

// ParseAuthorizedOperator is a log parse operation binding the contract event 0xf4caeb2d6ca8932a215a353d0703c326ec2d81fc68170f320eb2ab49e9df61f9.
//
// Solidity: event AuthorizedOperator(address indexed operator, address indexed tokenHolder)
func (_Erc777 *Erc777Filterer) ParseAuthorizedOperator(log types.Log) (*Erc777AuthorizedOperator, error) {
	event := new(Erc777AuthorizedOperator)
	if err := _Erc777.contract.UnpackLog(event, "AuthorizedOperator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc777BurnedIterator is returned from FilterBurned and is used to iterate over the raw logs and unpacked data for Burned events raised by the Erc777 contract.
type Erc777BurnedIterator struct {
	Event *Erc777Burned // Event containing the contract specifics and raw log

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
func (it *Erc777BurnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc777Burned)
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
		it.Event = new(Erc777Burned)
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
func (it *Erc777BurnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc777BurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc777Burned represents a Burned event raised by the Erc777 contract.
type Erc777Burned struct {
	Operator     common.Address
	From         common.Address
	Amount       *big.Int
	Data         []byte
	OperatorData []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterBurned is a free log retrieval operation binding the contract event 0xa78a9be3a7b862d26933ad85fb11d80ef66b8f972d7cbba06621d583943a4098.
//
// Solidity: event Burned(address indexed operator, address indexed from, uint256 amount, bytes data, bytes operatorData)
func (_Erc777 *Erc777Filterer) FilterBurned(opts *bind.FilterOpts, operator []common.Address, from []common.Address) (*Erc777BurnedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Erc777.contract.FilterLogs(opts, "Burned", operatorRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &Erc777BurnedIterator{contract: _Erc777.contract, event: "Burned", logs: logs, sub: sub}, nil
}

// WatchBurned is a free log subscription operation binding the contract event 0xa78a9be3a7b862d26933ad85fb11d80ef66b8f972d7cbba06621d583943a4098.
//
// Solidity: event Burned(address indexed operator, address indexed from, uint256 amount, bytes data, bytes operatorData)
func (_Erc777 *Erc777Filterer) WatchBurned(opts *bind.WatchOpts, sink chan<- *Erc777Burned, operator []common.Address, from []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Erc777.contract.WatchLogs(opts, "Burned", operatorRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc777Burned)
				if err := _Erc777.contract.UnpackLog(event, "Burned", log); err != nil {
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

// ParseBurned is a log parse operation binding the contract event 0xa78a9be3a7b862d26933ad85fb11d80ef66b8f972d7cbba06621d583943a4098.
//
// Solidity: event Burned(address indexed operator, address indexed from, uint256 amount, bytes data, bytes operatorData)
func (_Erc777 *Erc777Filterer) ParseBurned(log types.Log) (*Erc777Burned, error) {
	event := new(Erc777Burned)
	if err := _Erc777.contract.UnpackLog(event, "Burned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc777MintedIterator is returned from FilterMinted and is used to iterate over the raw logs and unpacked data for Minted events raised by the Erc777 contract.
type Erc777MintedIterator struct {
	Event *Erc777Minted // Event containing the contract specifics and raw log

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
func (it *Erc777MintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc777Minted)
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
		it.Event = new(Erc777Minted)
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
func (it *Erc777MintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc777MintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc777Minted represents a Minted event raised by the Erc777 contract.
type Erc777Minted struct {
	Operator     common.Address
	To           common.Address
	Amount       *big.Int
	Data         []byte
	OperatorData []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMinted is a free log retrieval operation binding the contract event 0x2fe5be0146f74c5bce36c0b80911af6c7d86ff27e89d5cfa61fc681327954e5d.
//
// Solidity: event Minted(address indexed operator, address indexed to, uint256 amount, bytes data, bytes operatorData)
func (_Erc777 *Erc777Filterer) FilterMinted(opts *bind.FilterOpts, operator []common.Address, to []common.Address) (*Erc777MintedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Erc777.contract.FilterLogs(opts, "Minted", operatorRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Erc777MintedIterator{contract: _Erc777.contract, event: "Minted", logs: logs, sub: sub}, nil
}

// WatchMinted is a free log subscription operation binding the contract event 0x2fe5be0146f74c5bce36c0b80911af6c7d86ff27e89d5cfa61fc681327954e5d.
//
// Solidity: event Minted(address indexed operator, address indexed to, uint256 amount, bytes data, bytes operatorData)
func (_Erc777 *Erc777Filterer) WatchMinted(opts *bind.WatchOpts, sink chan<- *Erc777Minted, operator []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Erc777.contract.WatchLogs(opts, "Minted", operatorRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc777Minted)
				if err := _Erc777.contract.UnpackLog(event, "Minted", log); err != nil {
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

// ParseMinted is a log parse operation binding the contract event 0x2fe5be0146f74c5bce36c0b80911af6c7d86ff27e89d5cfa61fc681327954e5d.
//
// Solidity: event Minted(address indexed operator, address indexed to, uint256 amount, bytes data, bytes operatorData)
func (_Erc777 *Erc777Filterer) ParseMinted(log types.Log) (*Erc777Minted, error) {
	event := new(Erc777Minted)
	if err := _Erc777.contract.UnpackLog(event, "Minted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc777RevokedOperatorIterator is returned from FilterRevokedOperator and is used to iterate over the raw logs and unpacked data for RevokedOperator events raised by the Erc777 contract.
type Erc777RevokedOperatorIterator struct {
	Event *Erc777RevokedOperator // Event containing the contract specifics and raw log

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
func (it *Erc777RevokedOperatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc777RevokedOperator)
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
		it.Event = new(Erc777RevokedOperator)
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
func (it *Erc777RevokedOperatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc777RevokedOperatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc777RevokedOperator represents a RevokedOperator event raised by the Erc777 contract.
type Erc777RevokedOperator struct {
	Operator    common.Address
	TokenHolder common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRevokedOperator is a free log retrieval operation binding the contract event 0x50546e66e5f44d728365dc3908c63bc5cfeeab470722c1677e3073a6ac294aa1.
//
// Solidity: event RevokedOperator(address indexed operator, address indexed tokenHolder)
func (_Erc777 *Erc777Filterer) FilterRevokedOperator(opts *bind.FilterOpts, operator []common.Address, tokenHolder []common.Address) (*Erc777RevokedOperatorIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var tokenHolderRule []interface{}
	for _, tokenHolderItem := range tokenHolder {
		tokenHolderRule = append(tokenHolderRule, tokenHolderItem)
	}

	logs, sub, err := _Erc777.contract.FilterLogs(opts, "RevokedOperator", operatorRule, tokenHolderRule)
	if err != nil {
		return nil, err
	}
	return &Erc777RevokedOperatorIterator{contract: _Erc777.contract, event: "RevokedOperator", logs: logs, sub: sub}, nil
}

// WatchRevokedOperator is a free log subscription operation binding the contract event 0x50546e66e5f44d728365dc3908c63bc5cfeeab470722c1677e3073a6ac294aa1.
//
// Solidity: event RevokedOperator(address indexed operator, address indexed tokenHolder)
func (_Erc777 *Erc777Filterer) WatchRevokedOperator(opts *bind.WatchOpts, sink chan<- *Erc777RevokedOperator, operator []common.Address, tokenHolder []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var tokenHolderRule []interface{}
	for _, tokenHolderItem := range tokenHolder {
		tokenHolderRule = append(tokenHolderRule, tokenHolderItem)
	}

	logs, sub, err := _Erc777.contract.WatchLogs(opts, "RevokedOperator", operatorRule, tokenHolderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc777RevokedOperator)
				if err := _Erc777.contract.UnpackLog(event, "RevokedOperator", log); err != nil {
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

// ParseRevokedOperator is a log parse operation binding the contract event 0x50546e66e5f44d728365dc3908c63bc5cfeeab470722c1677e3073a6ac294aa1.
//
// Solidity: event RevokedOperator(address indexed operator, address indexed tokenHolder)
func (_Erc777 *Erc777Filterer) ParseRevokedOperator(log types.Log) (*Erc777RevokedOperator, error) {
	event := new(Erc777RevokedOperator)
	if err := _Erc777.contract.UnpackLog(event, "RevokedOperator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc777SentIterator is returned from FilterSent and is used to iterate over the raw logs and unpacked data for Sent events raised by the Erc777 contract.
type Erc777SentIterator struct {
	Event *Erc777Sent // Event containing the contract specifics and raw log

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
func (it *Erc777SentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc777Sent)
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
		it.Event = new(Erc777Sent)
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
func (it *Erc777SentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc777SentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc777Sent represents a Sent event raised by the Erc777 contract.
type Erc777Sent struct {
	Operator     common.Address
	From         common.Address
	To           common.Address
	Amount       *big.Int
	Data         []byte
	OperatorData []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSent is a free log retrieval operation binding the contract event 0x06b541ddaa720db2b10a4d0cdac39b8d360425fc073085fac19bc82614677987.
//
// Solidity: event Sent(address indexed operator, address indexed from, address indexed to, uint256 amount, bytes data, bytes operatorData)
func (_Erc777 *Erc777Filterer) FilterSent(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*Erc777SentIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Erc777.contract.FilterLogs(opts, "Sent", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Erc777SentIterator{contract: _Erc777.contract, event: "Sent", logs: logs, sub: sub}, nil
}

// WatchSent is a free log subscription operation binding the contract event 0x06b541ddaa720db2b10a4d0cdac39b8d360425fc073085fac19bc82614677987.
//
// Solidity: event Sent(address indexed operator, address indexed from, address indexed to, uint256 amount, bytes data, bytes operatorData)
func (_Erc777 *Erc777Filterer) WatchSent(opts *bind.WatchOpts, sink chan<- *Erc777Sent, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Erc777.contract.WatchLogs(opts, "Sent", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc777Sent)
				if err := _Erc777.contract.UnpackLog(event, "Sent", log); err != nil {
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

// ParseSent is a log parse operation binding the contract event 0x06b541ddaa720db2b10a4d0cdac39b8d360425fc073085fac19bc82614677987.
//
// Solidity: event Sent(address indexed operator, address indexed from, address indexed to, uint256 amount, bytes data, bytes operatorData)
func (_Erc777 *Erc777Filterer) ParseSent(log types.Log) (*Erc777Sent, error) {
	event := new(Erc777Sent)
	if err := _Erc777.contract.UnpackLog(event, "Sent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc777TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Erc777 contract.
type Erc777TransferIterator struct {
	Event *Erc777Transfer // Event containing the contract specifics and raw log

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
func (it *Erc777TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc777Transfer)
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
		it.Event = new(Erc777Transfer)
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
func (it *Erc777TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc777TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc777Transfer represents a Transfer event raised by the Erc777 contract.
type Erc777Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Erc777 *Erc777Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*Erc777TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Erc777.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Erc777TransferIterator{contract: _Erc777.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Erc777 *Erc777Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *Erc777Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Erc777.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc777Transfer)
				if err := _Erc777.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Erc777 *Erc777Filterer) ParseTransfer(log types.Log) (*Erc777Transfer, error) {
	event := new(Erc777Transfer)
	if err := _Erc777.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
