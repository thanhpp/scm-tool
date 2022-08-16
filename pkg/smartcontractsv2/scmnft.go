// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package smartcontractsv2

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

// Smartcontractsv2MetaData contains all meta data concerning the Smartcontractsv2 contract.
var Smartcontractsv2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"multiSafeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"}],\"name\":\"selfMintNFT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// Smartcontractsv2ABI is the input ABI used to generate the binding from.
// Deprecated: Use Smartcontractsv2MetaData.ABI instead.
var Smartcontractsv2ABI = Smartcontractsv2MetaData.ABI

// Smartcontractsv2 is an auto generated Go binding around an Ethereum contract.
type Smartcontractsv2 struct {
	Smartcontractsv2Caller     // Read-only binding to the contract
	Smartcontractsv2Transactor // Write-only binding to the contract
	Smartcontractsv2Filterer   // Log filterer for contract events
}

// Smartcontractsv2Caller is an auto generated read-only Go binding around an Ethereum contract.
type Smartcontractsv2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Smartcontractsv2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Smartcontractsv2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Smartcontractsv2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Smartcontractsv2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Smartcontractsv2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Smartcontractsv2Session struct {
	Contract     *Smartcontractsv2 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Smartcontractsv2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Smartcontractsv2CallerSession struct {
	Contract *Smartcontractsv2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// Smartcontractsv2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Smartcontractsv2TransactorSession struct {
	Contract     *Smartcontractsv2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// Smartcontractsv2Raw is an auto generated low-level Go binding around an Ethereum contract.
type Smartcontractsv2Raw struct {
	Contract *Smartcontractsv2 // Generic contract binding to access the raw methods on
}

// Smartcontractsv2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Smartcontractsv2CallerRaw struct {
	Contract *Smartcontractsv2Caller // Generic read-only contract binding to access the raw methods on
}

// Smartcontractsv2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Smartcontractsv2TransactorRaw struct {
	Contract *Smartcontractsv2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewSmartcontractsv2 creates a new instance of Smartcontractsv2, bound to a specific deployed contract.
func NewSmartcontractsv2(address common.Address, backend bind.ContractBackend) (*Smartcontractsv2, error) {
	contract, err := bindSmartcontractsv2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Smartcontractsv2{Smartcontractsv2Caller: Smartcontractsv2Caller{contract: contract}, Smartcontractsv2Transactor: Smartcontractsv2Transactor{contract: contract}, Smartcontractsv2Filterer: Smartcontractsv2Filterer{contract: contract}}, nil
}

// NewSmartcontractsv2Caller creates a new read-only instance of Smartcontractsv2, bound to a specific deployed contract.
func NewSmartcontractsv2Caller(address common.Address, caller bind.ContractCaller) (*Smartcontractsv2Caller, error) {
	contract, err := bindSmartcontractsv2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Smartcontractsv2Caller{contract: contract}, nil
}

// NewSmartcontractsv2Transactor creates a new write-only instance of Smartcontractsv2, bound to a specific deployed contract.
func NewSmartcontractsv2Transactor(address common.Address, transactor bind.ContractTransactor) (*Smartcontractsv2Transactor, error) {
	contract, err := bindSmartcontractsv2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Smartcontractsv2Transactor{contract: contract}, nil
}

// NewSmartcontractsv2Filterer creates a new log filterer instance of Smartcontractsv2, bound to a specific deployed contract.
func NewSmartcontractsv2Filterer(address common.Address, filterer bind.ContractFilterer) (*Smartcontractsv2Filterer, error) {
	contract, err := bindSmartcontractsv2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Smartcontractsv2Filterer{contract: contract}, nil
}

// bindSmartcontractsv2 binds a generic wrapper to an already deployed contract.
func bindSmartcontractsv2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Smartcontractsv2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Smartcontractsv2 *Smartcontractsv2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Smartcontractsv2.Contract.Smartcontractsv2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Smartcontractsv2 *Smartcontractsv2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Smartcontractsv2.Contract.Smartcontractsv2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Smartcontractsv2 *Smartcontractsv2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Smartcontractsv2.Contract.Smartcontractsv2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Smartcontractsv2 *Smartcontractsv2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Smartcontractsv2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Smartcontractsv2 *Smartcontractsv2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Smartcontractsv2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Smartcontractsv2 *Smartcontractsv2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Smartcontractsv2.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Smartcontractsv2 *Smartcontractsv2Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Smartcontractsv2.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Smartcontractsv2 *Smartcontractsv2Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Smartcontractsv2.Contract.BalanceOf(&_Smartcontractsv2.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Smartcontractsv2 *Smartcontractsv2CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Smartcontractsv2.Contract.BalanceOf(&_Smartcontractsv2.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Smartcontractsv2 *Smartcontractsv2Caller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Smartcontractsv2.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Smartcontractsv2 *Smartcontractsv2Session) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Smartcontractsv2.Contract.GetApproved(&_Smartcontractsv2.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Smartcontractsv2 *Smartcontractsv2CallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Smartcontractsv2.Contract.GetApproved(&_Smartcontractsv2.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Smartcontractsv2 *Smartcontractsv2Caller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Smartcontractsv2.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Smartcontractsv2 *Smartcontractsv2Session) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Smartcontractsv2.Contract.IsApprovedForAll(&_Smartcontractsv2.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Smartcontractsv2 *Smartcontractsv2CallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Smartcontractsv2.Contract.IsApprovedForAll(&_Smartcontractsv2.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Smartcontractsv2 *Smartcontractsv2Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Smartcontractsv2.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Smartcontractsv2 *Smartcontractsv2Session) Name() (string, error) {
	return _Smartcontractsv2.Contract.Name(&_Smartcontractsv2.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Smartcontractsv2 *Smartcontractsv2CallerSession) Name() (string, error) {
	return _Smartcontractsv2.Contract.Name(&_Smartcontractsv2.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Smartcontractsv2 *Smartcontractsv2Caller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Smartcontractsv2.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Smartcontractsv2 *Smartcontractsv2Session) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Smartcontractsv2.Contract.OwnerOf(&_Smartcontractsv2.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Smartcontractsv2 *Smartcontractsv2CallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Smartcontractsv2.Contract.OwnerOf(&_Smartcontractsv2.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Smartcontractsv2 *Smartcontractsv2Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Smartcontractsv2.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Smartcontractsv2 *Smartcontractsv2Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Smartcontractsv2.Contract.SupportsInterface(&_Smartcontractsv2.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Smartcontractsv2 *Smartcontractsv2CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Smartcontractsv2.Contract.SupportsInterface(&_Smartcontractsv2.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Smartcontractsv2 *Smartcontractsv2Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Smartcontractsv2.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Smartcontractsv2 *Smartcontractsv2Session) Symbol() (string, error) {
	return _Smartcontractsv2.Contract.Symbol(&_Smartcontractsv2.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Smartcontractsv2 *Smartcontractsv2CallerSession) Symbol() (string, error) {
	return _Smartcontractsv2.Contract.Symbol(&_Smartcontractsv2.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Smartcontractsv2 *Smartcontractsv2Caller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Smartcontractsv2.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Smartcontractsv2 *Smartcontractsv2Session) TokenURI(tokenId *big.Int) (string, error) {
	return _Smartcontractsv2.Contract.TokenURI(&_Smartcontractsv2.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Smartcontractsv2 *Smartcontractsv2CallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Smartcontractsv2.Contract.TokenURI(&_Smartcontractsv2.CallOpts, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Smartcontractsv2 *Smartcontractsv2Transactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Smartcontractsv2.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Smartcontractsv2 *Smartcontractsv2Session) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Smartcontractsv2.Contract.Approve(&_Smartcontractsv2.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Smartcontractsv2 *Smartcontractsv2TransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Smartcontractsv2.Contract.Approve(&_Smartcontractsv2.TransactOpts, to, tokenId)
}

// MultiSafeTransferFrom is a paid mutator transaction binding the contract method 0x7dd4fc65.
//
// Solidity: function multiSafeTransferFrom(address from, address to, uint256[] tokenIds) returns()
func (_Smartcontractsv2 *Smartcontractsv2Transactor) MultiSafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenIds []*big.Int) (*types.Transaction, error) {
	return _Smartcontractsv2.contract.Transact(opts, "multiSafeTransferFrom", from, to, tokenIds)
}

// MultiSafeTransferFrom is a paid mutator transaction binding the contract method 0x7dd4fc65.
//
// Solidity: function multiSafeTransferFrom(address from, address to, uint256[] tokenIds) returns()
func (_Smartcontractsv2 *Smartcontractsv2Session) MultiSafeTransferFrom(from common.Address, to common.Address, tokenIds []*big.Int) (*types.Transaction, error) {
	return _Smartcontractsv2.Contract.MultiSafeTransferFrom(&_Smartcontractsv2.TransactOpts, from, to, tokenIds)
}

// MultiSafeTransferFrom is a paid mutator transaction binding the contract method 0x7dd4fc65.
//
// Solidity: function multiSafeTransferFrom(address from, address to, uint256[] tokenIds) returns()
func (_Smartcontractsv2 *Smartcontractsv2TransactorSession) MultiSafeTransferFrom(from common.Address, to common.Address, tokenIds []*big.Int) (*types.Transaction, error) {
	return _Smartcontractsv2.Contract.MultiSafeTransferFrom(&_Smartcontractsv2.TransactOpts, from, to, tokenIds)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Smartcontractsv2 *Smartcontractsv2Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Smartcontractsv2.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Smartcontractsv2 *Smartcontractsv2Session) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Smartcontractsv2.Contract.SafeTransferFrom(&_Smartcontractsv2.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Smartcontractsv2 *Smartcontractsv2TransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Smartcontractsv2.Contract.SafeTransferFrom(&_Smartcontractsv2.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Smartcontractsv2 *Smartcontractsv2Transactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Smartcontractsv2.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Smartcontractsv2 *Smartcontractsv2Session) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Smartcontractsv2.Contract.SafeTransferFrom0(&_Smartcontractsv2.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Smartcontractsv2 *Smartcontractsv2TransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Smartcontractsv2.Contract.SafeTransferFrom0(&_Smartcontractsv2.TransactOpts, from, to, tokenId, data)
}

// SelfMintNFT is a paid mutator transaction binding the contract method 0x862a7edd.
//
// Solidity: function selfMintNFT(string tokenURI) returns(uint256)
func (_Smartcontractsv2 *Smartcontractsv2Transactor) SelfMintNFT(opts *bind.TransactOpts, tokenURI string) (*types.Transaction, error) {
	return _Smartcontractsv2.contract.Transact(opts, "selfMintNFT", tokenURI)
}

// SelfMintNFT is a paid mutator transaction binding the contract method 0x862a7edd.
//
// Solidity: function selfMintNFT(string tokenURI) returns(uint256)
func (_Smartcontractsv2 *Smartcontractsv2Session) SelfMintNFT(tokenURI string) (*types.Transaction, error) {
	return _Smartcontractsv2.Contract.SelfMintNFT(&_Smartcontractsv2.TransactOpts, tokenURI)
}

// SelfMintNFT is a paid mutator transaction binding the contract method 0x862a7edd.
//
// Solidity: function selfMintNFT(string tokenURI) returns(uint256)
func (_Smartcontractsv2 *Smartcontractsv2TransactorSession) SelfMintNFT(tokenURI string) (*types.Transaction, error) {
	return _Smartcontractsv2.Contract.SelfMintNFT(&_Smartcontractsv2.TransactOpts, tokenURI)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Smartcontractsv2 *Smartcontractsv2Transactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Smartcontractsv2.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Smartcontractsv2 *Smartcontractsv2Session) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Smartcontractsv2.Contract.SetApprovalForAll(&_Smartcontractsv2.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Smartcontractsv2 *Smartcontractsv2TransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Smartcontractsv2.Contract.SetApprovalForAll(&_Smartcontractsv2.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Smartcontractsv2 *Smartcontractsv2Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Smartcontractsv2.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Smartcontractsv2 *Smartcontractsv2Session) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Smartcontractsv2.Contract.TransferFrom(&_Smartcontractsv2.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Smartcontractsv2 *Smartcontractsv2TransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Smartcontractsv2.Contract.TransferFrom(&_Smartcontractsv2.TransactOpts, from, to, tokenId)
}

// Smartcontractsv2ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Smartcontractsv2 contract.
type Smartcontractsv2ApprovalIterator struct {
	Event *Smartcontractsv2Approval // Event containing the contract specifics and raw log

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
func (it *Smartcontractsv2ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Smartcontractsv2Approval)
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
		it.Event = new(Smartcontractsv2Approval)
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
func (it *Smartcontractsv2ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Smartcontractsv2ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Smartcontractsv2Approval represents a Approval event raised by the Smartcontractsv2 contract.
type Smartcontractsv2Approval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Smartcontractsv2 *Smartcontractsv2Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*Smartcontractsv2ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Smartcontractsv2.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &Smartcontractsv2ApprovalIterator{contract: _Smartcontractsv2.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Smartcontractsv2 *Smartcontractsv2Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *Smartcontractsv2Approval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Smartcontractsv2.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Smartcontractsv2Approval)
				if err := _Smartcontractsv2.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Smartcontractsv2 *Smartcontractsv2Filterer) ParseApproval(log types.Log) (*Smartcontractsv2Approval, error) {
	event := new(Smartcontractsv2Approval)
	if err := _Smartcontractsv2.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Smartcontractsv2ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Smartcontractsv2 contract.
type Smartcontractsv2ApprovalForAllIterator struct {
	Event *Smartcontractsv2ApprovalForAll // Event containing the contract specifics and raw log

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
func (it *Smartcontractsv2ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Smartcontractsv2ApprovalForAll)
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
		it.Event = new(Smartcontractsv2ApprovalForAll)
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
func (it *Smartcontractsv2ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Smartcontractsv2ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Smartcontractsv2ApprovalForAll represents a ApprovalForAll event raised by the Smartcontractsv2 contract.
type Smartcontractsv2ApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Smartcontractsv2 *Smartcontractsv2Filterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*Smartcontractsv2ApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Smartcontractsv2.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &Smartcontractsv2ApprovalForAllIterator{contract: _Smartcontractsv2.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Smartcontractsv2 *Smartcontractsv2Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *Smartcontractsv2ApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Smartcontractsv2.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Smartcontractsv2ApprovalForAll)
				if err := _Smartcontractsv2.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Smartcontractsv2 *Smartcontractsv2Filterer) ParseApprovalForAll(log types.Log) (*Smartcontractsv2ApprovalForAll, error) {
	event := new(Smartcontractsv2ApprovalForAll)
	if err := _Smartcontractsv2.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Smartcontractsv2TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Smartcontractsv2 contract.
type Smartcontractsv2TransferIterator struct {
	Event *Smartcontractsv2Transfer // Event containing the contract specifics and raw log

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
func (it *Smartcontractsv2TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Smartcontractsv2Transfer)
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
		it.Event = new(Smartcontractsv2Transfer)
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
func (it *Smartcontractsv2TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Smartcontractsv2TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Smartcontractsv2Transfer represents a Transfer event raised by the Smartcontractsv2 contract.
type Smartcontractsv2Transfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Smartcontractsv2 *Smartcontractsv2Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*Smartcontractsv2TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Smartcontractsv2.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &Smartcontractsv2TransferIterator{contract: _Smartcontractsv2.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Smartcontractsv2 *Smartcontractsv2Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *Smartcontractsv2Transfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Smartcontractsv2.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Smartcontractsv2Transfer)
				if err := _Smartcontractsv2.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Smartcontractsv2 *Smartcontractsv2Filterer) ParseTransfer(log types.Log) (*Smartcontractsv2Transfer, error) {
	event := new(Smartcontractsv2Transfer)
	if err := _Smartcontractsv2.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
