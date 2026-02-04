// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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

// TestErrorMetaData contains all meta data concerning the TestError contract.
var TestErrorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"NoAuthorization\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"IsOwner\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_Owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TestErrorABI is the input ABI used to generate the binding from.
// Deprecated: Use TestErrorMetaData.ABI instead.
var TestErrorABI = TestErrorMetaData.ABI

// TestError is an auto generated Go binding around an Ethereum contract.
type TestError struct {
	TestErrorCaller     // Read-only binding to the contract
	TestErrorTransactor // Write-only binding to the contract
	TestErrorFilterer   // Log filterer for contract events
}

// TestErrorCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestErrorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestErrorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestErrorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestErrorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestErrorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestErrorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestErrorSession struct {
	Contract     *TestError        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestErrorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestErrorCallerSession struct {
	Contract *TestErrorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// TestErrorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestErrorTransactorSession struct {
	Contract     *TestErrorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// TestErrorRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestErrorRaw struct {
	Contract *TestError // Generic contract binding to access the raw methods on
}

// TestErrorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestErrorCallerRaw struct {
	Contract *TestErrorCaller // Generic read-only contract binding to access the raw methods on
}

// TestErrorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestErrorTransactorRaw struct {
	Contract *TestErrorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestError creates a new instance of TestError, bound to a specific deployed contract.
func NewTestError(address common.Address, backend bind.ContractBackend) (*TestError, error) {
	contract, err := bindTestError(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestError{TestErrorCaller: TestErrorCaller{contract: contract}, TestErrorTransactor: TestErrorTransactor{contract: contract}, TestErrorFilterer: TestErrorFilterer{contract: contract}}, nil
}

// NewTestErrorCaller creates a new read-only instance of TestError, bound to a specific deployed contract.
func NewTestErrorCaller(address common.Address, caller bind.ContractCaller) (*TestErrorCaller, error) {
	contract, err := bindTestError(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestErrorCaller{contract: contract}, nil
}

// NewTestErrorTransactor creates a new write-only instance of TestError, bound to a specific deployed contract.
func NewTestErrorTransactor(address common.Address, transactor bind.ContractTransactor) (*TestErrorTransactor, error) {
	contract, err := bindTestError(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestErrorTransactor{contract: contract}, nil
}

// NewTestErrorFilterer creates a new log filterer instance of TestError, bound to a specific deployed contract.
func NewTestErrorFilterer(address common.Address, filterer bind.ContractFilterer) (*TestErrorFilterer, error) {
	contract, err := bindTestError(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestErrorFilterer{contract: contract}, nil
}

// bindTestError binds a generic wrapper to an already deployed contract.
func bindTestError(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TestErrorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestError *TestErrorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestError.Contract.TestErrorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestError *TestErrorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestError.Contract.TestErrorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestError *TestErrorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestError.Contract.TestErrorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestError *TestErrorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestError.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestError *TestErrorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestError.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestError *TestErrorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestError.Contract.contract.Transact(opts, method, params...)
}

// IsOwner is a free data retrieval call binding the contract method 0xdd0860a8.
//
// Solidity: function IsOwner(address _user) view returns()
func (_TestError *TestErrorCaller) IsOwner(opts *bind.CallOpts, _user common.Address) error {
	var out []interface{}
	err := _TestError.contract.Call(opts, &out, "IsOwner", _user)

	if err != nil {
		return err
	}

	return err

}

// IsOwner is a free data retrieval call binding the contract method 0xdd0860a8.
//
// Solidity: function IsOwner(address _user) view returns()
func (_TestError *TestErrorSession) IsOwner(_user common.Address) error {
	return _TestError.Contract.IsOwner(&_TestError.CallOpts, _user)
}

// IsOwner is a free data retrieval call binding the contract method 0xdd0860a8.
//
// Solidity: function IsOwner(address _user) view returns()
func (_TestError *TestErrorCallerSession) IsOwner(_user common.Address) error {
	return _TestError.Contract.IsOwner(&_TestError.CallOpts, _user)
}

// ContractOwner is a free data retrieval call binding the contract method 0xdcc23455.
//
// Solidity: function _Owner() view returns(address)
func (_TestError *TestErrorCaller) ContractOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TestError.contract.Call(opts, &out, "_Owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ContractOwner is a free data retrieval call binding the contract method 0xdcc23455.
//
// Solidity: function _Owner() view returns(address)
func (_TestError *TestErrorSession) ContractOwner() (common.Address, error) {
	return _TestError.Contract.ContractOwner(&_TestError.CallOpts)
}

// ContractOwner is a free data retrieval call binding the contract method 0xdcc23455.
//
// Solidity: function _Owner() view returns(address)
func (_TestError *TestErrorCallerSession) ContractOwner() (common.Address, error) {
	return _TestError.Contract.ContractOwner(&_TestError.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TestError *TestErrorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TestError.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TestError *TestErrorSession) Owner() (common.Address, error) {
	return _TestError.Contract.Owner(&_TestError.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TestError *TestErrorCallerSession) Owner() (common.Address, error) {
	return _TestError.Contract.Owner(&_TestError.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TestError *TestErrorTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestError.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TestError *TestErrorSession) RenounceOwnership() (*types.Transaction, error) {
	return _TestError.Contract.RenounceOwnership(&_TestError.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TestError *TestErrorTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _TestError.Contract.RenounceOwnership(&_TestError.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TestError *TestErrorTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TestError.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TestError *TestErrorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TestError.Contract.TransferOwnership(&_TestError.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TestError *TestErrorTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TestError.Contract.TransferOwnership(&_TestError.TransactOpts, newOwner)
}

// TestErrorOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TestError contract.
type TestErrorOwnershipTransferredIterator struct {
	Event *TestErrorOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TestErrorOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestErrorOwnershipTransferred)
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
		it.Event = new(TestErrorOwnershipTransferred)
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
func (it *TestErrorOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestErrorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestErrorOwnershipTransferred represents a OwnershipTransferred event raised by the TestError contract.
type TestErrorOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TestError *TestErrorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TestErrorOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TestError.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TestErrorOwnershipTransferredIterator{contract: _TestError.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TestError *TestErrorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TestErrorOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TestError.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestErrorOwnershipTransferred)
				if err := _TestError.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TestError *TestErrorFilterer) ParseOwnershipTransferred(log types.Log) (*TestErrorOwnershipTransferred, error) {
	event := new(TestErrorOwnershipTransferred)
	if err := _TestError.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
