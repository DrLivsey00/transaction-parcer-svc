// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package parser

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

// TokenFilterMetaData contains all meta data concerning the TokenFilter contract.
var TokenFilterMetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenOwner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenOwner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"tokenOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]",
}

// TokenFilterABI is the input ABI used to generate the binding from.
// Deprecated: Use TokenFilterMetaData.ABI instead.
var TokenFilterABI = TokenFilterMetaData.ABI

// TokenFilter is an auto generated Go binding around an Ethereum contract.
type TokenFilter struct {
	TokenFilterCaller     // Read-only binding to the contract
	TokenFilterTransactor // Write-only binding to the contract
	TokenFilterFilterer   // Log filterer for contract events
}

// TokenFilterCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenFilterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenFilterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenFilterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenFilterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenFilterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenFilterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenFilterSession struct {
	Contract     *TokenFilter      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenFilterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenFilterCallerSession struct {
	Contract *TokenFilterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// TokenFilterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenFilterTransactorSession struct {
	Contract     *TokenFilterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// TokenFilterRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenFilterRaw struct {
	Contract *TokenFilter // Generic contract binding to access the raw methods on
}

// TokenFilterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenFilterCallerRaw struct {
	Contract *TokenFilterCaller // Generic read-only contract binding to access the raw methods on
}

// TokenFilterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenFilterTransactorRaw struct {
	Contract *TokenFilterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenFilter creates a new instance of TokenFilter, bound to a specific deployed contract.
func NewTokenFilter(address common.Address, backend bind.ContractBackend) (*TokenFilter, error) {
	contract, err := bindTokenFilter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenFilter{TokenFilterCaller: TokenFilterCaller{contract: contract}, TokenFilterTransactor: TokenFilterTransactor{contract: contract}, TokenFilterFilterer: TokenFilterFilterer{contract: contract}}, nil
}

// NewTokenFilterCaller creates a new read-only instance of TokenFilter, bound to a specific deployed contract.
func NewTokenFilterCaller(address common.Address, caller bind.ContractCaller) (*TokenFilterCaller, error) {
	contract, err := bindTokenFilter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenFilterCaller{contract: contract}, nil
}

// NewTokenFilterTransactor creates a new write-only instance of TokenFilter, bound to a specific deployed contract.
func NewTokenFilterTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenFilterTransactor, error) {
	contract, err := bindTokenFilter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenFilterTransactor{contract: contract}, nil
}

// NewTokenFilterFilterer creates a new log filterer instance of TokenFilter, bound to a specific deployed contract.
func NewTokenFilterFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenFilterFilterer, error) {
	contract, err := bindTokenFilter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenFilterFilterer{contract: contract}, nil
}

// bindTokenFilter binds a generic wrapper to an already deployed contract.
func bindTokenFilter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TokenFilterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenFilter *TokenFilterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenFilter.Contract.TokenFilterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenFilter *TokenFilterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenFilter.Contract.TokenFilterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenFilter *TokenFilterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenFilter.Contract.TokenFilterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenFilter *TokenFilterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenFilter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenFilter *TokenFilterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenFilter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenFilter *TokenFilterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenFilter.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address tokenOwner, address spender) view returns(uint256 remaining)
func (_TokenFilter *TokenFilterCaller) Allowance(opts *bind.CallOpts, tokenOwner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TokenFilter.contract.Call(opts, &out, "allowance", tokenOwner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address tokenOwner, address spender) view returns(uint256 remaining)
func (_TokenFilter *TokenFilterSession) Allowance(tokenOwner common.Address, spender common.Address) (*big.Int, error) {
	return _TokenFilter.Contract.Allowance(&_TokenFilter.CallOpts, tokenOwner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address tokenOwner, address spender) view returns(uint256 remaining)
func (_TokenFilter *TokenFilterCallerSession) Allowance(tokenOwner common.Address, spender common.Address) (*big.Int, error) {
	return _TokenFilter.Contract.Allowance(&_TokenFilter.CallOpts, tokenOwner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address tokenOwner) view returns(uint256 balance)
func (_TokenFilter *TokenFilterCaller) BalanceOf(opts *bind.CallOpts, tokenOwner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TokenFilter.contract.Call(opts, &out, "balanceOf", tokenOwner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address tokenOwner) view returns(uint256 balance)
func (_TokenFilter *TokenFilterSession) BalanceOf(tokenOwner common.Address) (*big.Int, error) {
	return _TokenFilter.Contract.BalanceOf(&_TokenFilter.CallOpts, tokenOwner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address tokenOwner) view returns(uint256 balance)
func (_TokenFilter *TokenFilterCallerSession) BalanceOf(tokenOwner common.Address) (*big.Int, error) {
	return _TokenFilter.Contract.BalanceOf(&_TokenFilter.CallOpts, tokenOwner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_TokenFilter *TokenFilterCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _TokenFilter.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_TokenFilter *TokenFilterSession) Decimals() (uint8, error) {
	return _TokenFilter.Contract.Decimals(&_TokenFilter.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_TokenFilter *TokenFilterCallerSession) Decimals() (uint8, error) {
	return _TokenFilter.Contract.Decimals(&_TokenFilter.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TokenFilter *TokenFilterCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TokenFilter.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TokenFilter *TokenFilterSession) Name() (string, error) {
	return _TokenFilter.Contract.Name(&_TokenFilter.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TokenFilter *TokenFilterCallerSession) Name() (string, error) {
	return _TokenFilter.Contract.Name(&_TokenFilter.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TokenFilter *TokenFilterCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TokenFilter.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TokenFilter *TokenFilterSession) Symbol() (string, error) {
	return _TokenFilter.Contract.Symbol(&_TokenFilter.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TokenFilter *TokenFilterCallerSession) Symbol() (string, error) {
	return _TokenFilter.Contract.Symbol(&_TokenFilter.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TokenFilter *TokenFilterCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TokenFilter.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TokenFilter *TokenFilterSession) TotalSupply() (*big.Int, error) {
	return _TokenFilter.Contract.TotalSupply(&_TokenFilter.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TokenFilter *TokenFilterCallerSession) TotalSupply() (*big.Int, error) {
	return _TokenFilter.Contract.TotalSupply(&_TokenFilter.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 tokens) returns(bool success)
func (_TokenFilter *TokenFilterTransactor) Approve(opts *bind.TransactOpts, spender common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _TokenFilter.contract.Transact(opts, "approve", spender, tokens)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 tokens) returns(bool success)
func (_TokenFilter *TokenFilterSession) Approve(spender common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _TokenFilter.Contract.Approve(&_TokenFilter.TransactOpts, spender, tokens)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 tokens) returns(bool success)
func (_TokenFilter *TokenFilterTransactorSession) Approve(spender common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _TokenFilter.Contract.Approve(&_TokenFilter.TransactOpts, spender, tokens)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 tokens) returns(bool success)
func (_TokenFilter *TokenFilterTransactor) Transfer(opts *bind.TransactOpts, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _TokenFilter.contract.Transact(opts, "transfer", to, tokens)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 tokens) returns(bool success)
func (_TokenFilter *TokenFilterSession) Transfer(to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _TokenFilter.Contract.Transfer(&_TokenFilter.TransactOpts, to, tokens)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 tokens) returns(bool success)
func (_TokenFilter *TokenFilterTransactorSession) Transfer(to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _TokenFilter.Contract.Transfer(&_TokenFilter.TransactOpts, to, tokens)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokens) returns(bool success)
func (_TokenFilter *TokenFilterTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _TokenFilter.contract.Transact(opts, "transferFrom", from, to, tokens)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokens) returns(bool success)
func (_TokenFilter *TokenFilterSession) TransferFrom(from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _TokenFilter.Contract.TransferFrom(&_TokenFilter.TransactOpts, from, to, tokens)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokens) returns(bool success)
func (_TokenFilter *TokenFilterTransactorSession) TransferFrom(from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _TokenFilter.Contract.TransferFrom(&_TokenFilter.TransactOpts, from, to, tokens)
}

// TokenFilterApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the TokenFilter contract.
type TokenFilterApprovalIterator struct {
	Event *TokenFilterApproval // Event containing the contract specifics and raw log

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
func (it *TokenFilterApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenFilterApproval)
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
		it.Event = new(TokenFilterApproval)
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
func (it *TokenFilterApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenFilterApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenFilterApproval represents a Approval event raised by the TokenFilter contract.
type TokenFilterApproval struct {
	TokenOwner common.Address
	Spender    common.Address
	Tokens     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed tokenOwner, address indexed spender, uint256 tokens)
func (_TokenFilter *TokenFilterFilterer) FilterApproval(opts *bind.FilterOpts, tokenOwner []common.Address, spender []common.Address) (*TokenFilterApprovalIterator, error) {

	var tokenOwnerRule []interface{}
	for _, tokenOwnerItem := range tokenOwner {
		tokenOwnerRule = append(tokenOwnerRule, tokenOwnerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _TokenFilter.contract.FilterLogs(opts, "Approval", tokenOwnerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &TokenFilterApprovalIterator{contract: _TokenFilter.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed tokenOwner, address indexed spender, uint256 tokens)
func (_TokenFilter *TokenFilterFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *TokenFilterApproval, tokenOwner []common.Address, spender []common.Address) (event.Subscription, error) {

	var tokenOwnerRule []interface{}
	for _, tokenOwnerItem := range tokenOwner {
		tokenOwnerRule = append(tokenOwnerRule, tokenOwnerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _TokenFilter.contract.WatchLogs(opts, "Approval", tokenOwnerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenFilterApproval)
				if err := _TokenFilter.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed tokenOwner, address indexed spender, uint256 tokens)
func (_TokenFilter *TokenFilterFilterer) ParseApproval(log types.Log) (*TokenFilterApproval, error) {
	event := new(TokenFilterApproval)
	if err := _TokenFilter.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenFilterTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the TokenFilter contract.
type TokenFilterTransferIterator struct {
	Event *TokenFilterTransfer // Event containing the contract specifics and raw log

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
func (it *TokenFilterTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenFilterTransfer)
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
		it.Event = new(TokenFilterTransfer)
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
func (it *TokenFilterTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenFilterTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenFilterTransfer represents a Transfer event raised by the TokenFilter contract.
type TokenFilterTransfer struct {
	From   common.Address
	To     common.Address
	Tokens *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 tokens)
func (_TokenFilter *TokenFilterFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*TokenFilterTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TokenFilter.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TokenFilterTransferIterator{contract: _TokenFilter.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 tokens)
func (_TokenFilter *TokenFilterFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *TokenFilterTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TokenFilter.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenFilterTransfer)
				if err := _TokenFilter.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 tokens)
func (_TokenFilter *TokenFilterFilterer) ParseTransfer(log types.Log) (*TokenFilterTransfer, error) {
	event := new(TokenFilterTransfer)
	if err := _TokenFilter.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
