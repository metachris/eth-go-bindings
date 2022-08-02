// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc165

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

// Erc165MetaData contains all meta data concerning the Erc165 contract.
var Erc165MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506101b7806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c806301ffc9a714610030575b600080fd5b61004a600480360381019061004591906100df565b610060565b6040516100579190610117565b60405180910390f35b60007f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b6000813590506100d98161016a565b92915050565b6000602082840312156100f157600080fd5b60006100ff848285016100ca565b91505092915050565b61011181610132565b82525050565b600060208201905061012c6000830184610108565b92915050565b60008115159050919050565b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b6101738161013e565b811461017e57600080fd5b5056fea26469706673582212206301fac24ced774e086182432be8e21829d3db0195a3239cda932117c86325c564736f6c63430008040033",
}

// Erc165ABI is the input ABI used to generate the binding from.
// Deprecated: Use Erc165MetaData.ABI instead.
var Erc165ABI = Erc165MetaData.ABI

// Erc165Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Erc165MetaData.Bin instead.
var Erc165Bin = Erc165MetaData.Bin

// DeployErc165 deploys a new Ethereum contract, binding an instance of Erc165 to it.
func DeployErc165(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Erc165, error) {
	parsed, err := Erc165MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Erc165Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Erc165{Erc165Caller: Erc165Caller{contract: contract}, Erc165Transactor: Erc165Transactor{contract: contract}, Erc165Filterer: Erc165Filterer{contract: contract}}, nil
}

// Erc165 is an auto generated Go binding around an Ethereum contract.
type Erc165 struct {
	Erc165Caller     // Read-only binding to the contract
	Erc165Transactor // Write-only binding to the contract
	Erc165Filterer   // Log filterer for contract events
}

// Erc165Caller is an auto generated read-only Go binding around an Ethereum contract.
type Erc165Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc165Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Erc165Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc165Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Erc165Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc165Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Erc165Session struct {
	Contract     *Erc165           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Erc165CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Erc165CallerSession struct {
	Contract *Erc165Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// Erc165TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Erc165TransactorSession struct {
	Contract     *Erc165Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Erc165Raw is an auto generated low-level Go binding around an Ethereum contract.
type Erc165Raw struct {
	Contract *Erc165 // Generic contract binding to access the raw methods on
}

// Erc165CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Erc165CallerRaw struct {
	Contract *Erc165Caller // Generic read-only contract binding to access the raw methods on
}

// Erc165TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Erc165TransactorRaw struct {
	Contract *Erc165Transactor // Generic write-only contract binding to access the raw methods on
}

// NewErc165 creates a new instance of Erc165, bound to a specific deployed contract.
func NewErc165(address common.Address, backend bind.ContractBackend) (*Erc165, error) {
	contract, err := bindErc165(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Erc165{Erc165Caller: Erc165Caller{contract: contract}, Erc165Transactor: Erc165Transactor{contract: contract}, Erc165Filterer: Erc165Filterer{contract: contract}}, nil
}

// NewErc165Caller creates a new read-only instance of Erc165, bound to a specific deployed contract.
func NewErc165Caller(address common.Address, caller bind.ContractCaller) (*Erc165Caller, error) {
	contract, err := bindErc165(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Erc165Caller{contract: contract}, nil
}

// NewErc165Transactor creates a new write-only instance of Erc165, bound to a specific deployed contract.
func NewErc165Transactor(address common.Address, transactor bind.ContractTransactor) (*Erc165Transactor, error) {
	contract, err := bindErc165(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Erc165Transactor{contract: contract}, nil
}

// NewErc165Filterer creates a new log filterer instance of Erc165, bound to a specific deployed contract.
func NewErc165Filterer(address common.Address, filterer bind.ContractFilterer) (*Erc165Filterer, error) {
	contract, err := bindErc165(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Erc165Filterer{contract: contract}, nil
}

// bindErc165 binds a generic wrapper to an already deployed contract.
func bindErc165(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Erc165ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc165 *Erc165Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc165.Contract.Erc165Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc165 *Erc165Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc165.Contract.Erc165Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc165 *Erc165Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc165.Contract.Erc165Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc165 *Erc165CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc165.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc165 *Erc165TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc165.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc165 *Erc165TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc165.Contract.contract.Transact(opts, method, params...)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Erc165 *Erc165Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Erc165.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Erc165 *Erc165Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Erc165.Contract.SupportsInterface(&_Erc165.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Erc165 *Erc165CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Erc165.Contract.SupportsInterface(&_Erc165.CallOpts, interfaceId)
}
