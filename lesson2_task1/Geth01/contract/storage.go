// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// StroageMetaData contains all meta data concerning the Stroage contract.
var StroageMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"MyEvent\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"number\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"retrieve\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"store\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b506101a88061001c5f395ff3fe608060405234801561000f575f5ffd5b506004361061003f575f3560e01c80632e64cec1146100435780636057361d146100615780638381f58a1461007d575b5f5ffd5b61004b61009b565b6040516100589190610100565b60405180910390f35b61007b60048036038101906100769190610147565b6100a3565b005b6100856100e3565b6040516100929190610100565b60405180910390f35b5f5f54905090565b805f819055507f6c2b4666ba8da5a95717621d879a77de725f3d816709b9cbe9f059b8f875e284816040516100d89190610100565b60405180910390a150565b5f5481565b5f819050919050565b6100fa816100e8565b82525050565b5f6020820190506101135f8301846100f1565b92915050565b5f5ffd5b610126816100e8565b8114610130575f5ffd5b50565b5f813590506101418161011d565b92915050565b5f6020828403121561015c5761015b610119565b5b5f61016984828501610133565b9150509291505056fea264697066735822122047dc8e7bdc7813a410fb92c4ea7b39c46322eac56c966cab260956a56732742864736f6c634300081c0033",
}

// StroageABI is the input ABI used to generate the binding from.
// Deprecated: Use StroageMetaData.ABI instead.
var StroageABI = StroageMetaData.ABI

// StroageBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StroageMetaData.Bin instead.
var StroageBin = StroageMetaData.Bin

// DeployStroage deploys a new Ethereum contract, binding an instance of Stroage to it.
func DeployStroage(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Stroage, error) {
	parsed, err := StroageMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StroageBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Stroage{StroageCaller: StroageCaller{contract: contract}, StroageTransactor: StroageTransactor{contract: contract}, StroageFilterer: StroageFilterer{contract: contract}}, nil
}

// Stroage is an auto generated Go binding around an Ethereum contract.
type Stroage struct {
	StroageCaller     // Read-only binding to the contract
	StroageTransactor // Write-only binding to the contract
	StroageFilterer   // Log filterer for contract events
}

// StroageCaller is an auto generated read-only Go binding around an Ethereum contract.
type StroageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StroageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StroageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StroageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StroageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StroageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StroageSession struct {
	Contract     *Stroage          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StroageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StroageCallerSession struct {
	Contract *StroageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// StroageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StroageTransactorSession struct {
	Contract     *StroageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StroageRaw is an auto generated low-level Go binding around an Ethereum contract.
type StroageRaw struct {
	Contract *Stroage // Generic contract binding to access the raw methods on
}

// StroageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StroageCallerRaw struct {
	Contract *StroageCaller // Generic read-only contract binding to access the raw methods on
}

// StroageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StroageTransactorRaw struct {
	Contract *StroageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStroage creates a new instance of Stroage, bound to a specific deployed contract.
func NewStroage(address common.Address, backend bind.ContractBackend) (*Stroage, error) {
	contract, err := bindStroage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Stroage{StroageCaller: StroageCaller{contract: contract}, StroageTransactor: StroageTransactor{contract: contract}, StroageFilterer: StroageFilterer{contract: contract}}, nil
}

// NewStroageCaller creates a new read-only instance of Stroage, bound to a specific deployed contract.
func NewStroageCaller(address common.Address, caller bind.ContractCaller) (*StroageCaller, error) {
	contract, err := bindStroage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StroageCaller{contract: contract}, nil
}

// NewStroageTransactor creates a new write-only instance of Stroage, bound to a specific deployed contract.
func NewStroageTransactor(address common.Address, transactor bind.ContractTransactor) (*StroageTransactor, error) {
	contract, err := bindStroage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StroageTransactor{contract: contract}, nil
}

// NewStroageFilterer creates a new log filterer instance of Stroage, bound to a specific deployed contract.
func NewStroageFilterer(address common.Address, filterer bind.ContractFilterer) (*StroageFilterer, error) {
	contract, err := bindStroage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StroageFilterer{contract: contract}, nil
}

// bindStroage binds a generic wrapper to an already deployed contract.
func bindStroage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StroageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Stroage *StroageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Stroage.Contract.StroageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Stroage *StroageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stroage.Contract.StroageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Stroage *StroageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Stroage.Contract.StroageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Stroage *StroageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Stroage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Stroage *StroageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stroage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Stroage *StroageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Stroage.Contract.contract.Transact(opts, method, params...)
}

// Number is a free data retrieval call binding the contract method 0x8381f58a.
//
// Solidity: function number() view returns(uint256)
func (_Stroage *StroageCaller) Number(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stroage.contract.Call(opts, &out, "number")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Number is a free data retrieval call binding the contract method 0x8381f58a.
//
// Solidity: function number() view returns(uint256)
func (_Stroage *StroageSession) Number() (*big.Int, error) {
	return _Stroage.Contract.Number(&_Stroage.CallOpts)
}

// Number is a free data retrieval call binding the contract method 0x8381f58a.
//
// Solidity: function number() view returns(uint256)
func (_Stroage *StroageCallerSession) Number() (*big.Int, error) {
	return _Stroage.Contract.Number(&_Stroage.CallOpts)
}

// Retrieve is a free data retrieval call binding the contract method 0x2e64cec1.
//
// Solidity: function retrieve() view returns(uint256)
func (_Stroage *StroageCaller) Retrieve(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stroage.contract.Call(opts, &out, "retrieve")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Retrieve is a free data retrieval call binding the contract method 0x2e64cec1.
//
// Solidity: function retrieve() view returns(uint256)
func (_Stroage *StroageSession) Retrieve() (*big.Int, error) {
	return _Stroage.Contract.Retrieve(&_Stroage.CallOpts)
}

// Retrieve is a free data retrieval call binding the contract method 0x2e64cec1.
//
// Solidity: function retrieve() view returns(uint256)
func (_Stroage *StroageCallerSession) Retrieve() (*big.Int, error) {
	return _Stroage.Contract.Retrieve(&_Stroage.CallOpts)
}

// Store is a paid mutator transaction binding the contract method 0x6057361d.
//
// Solidity: function store(uint256 num) returns()
func (_Stroage *StroageTransactor) Store(opts *bind.TransactOpts, num *big.Int) (*types.Transaction, error) {
	return _Stroage.contract.Transact(opts, "store", num)
}

// Store is a paid mutator transaction binding the contract method 0x6057361d.
//
// Solidity: function store(uint256 num) returns()
func (_Stroage *StroageSession) Store(num *big.Int) (*types.Transaction, error) {
	return _Stroage.Contract.Store(&_Stroage.TransactOpts, num)
}

// Store is a paid mutator transaction binding the contract method 0x6057361d.
//
// Solidity: function store(uint256 num) returns()
func (_Stroage *StroageTransactorSession) Store(num *big.Int) (*types.Transaction, error) {
	return _Stroage.Contract.Store(&_Stroage.TransactOpts, num)
}

// StroageMyEventIterator is returned from FilterMyEvent and is used to iterate over the raw logs and unpacked data for MyEvent events raised by the Stroage contract.
type StroageMyEventIterator struct {
	Event *StroageMyEvent // Event containing the contract specifics and raw log

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
func (it *StroageMyEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StroageMyEvent)
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
		it.Event = new(StroageMyEvent)
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
func (it *StroageMyEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StroageMyEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StroageMyEvent represents a MyEvent event raised by the Stroage contract.
type StroageMyEvent struct {
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMyEvent is a free log retrieval operation binding the contract event 0x6c2b4666ba8da5a95717621d879a77de725f3d816709b9cbe9f059b8f875e284.
//
// Solidity: event MyEvent(uint256 value)
func (_Stroage *StroageFilterer) FilterMyEvent(opts *bind.FilterOpts) (*StroageMyEventIterator, error) {

	logs, sub, err := _Stroage.contract.FilterLogs(opts, "MyEvent")
	if err != nil {
		return nil, err
	}
	return &StroageMyEventIterator{contract: _Stroage.contract, event: "MyEvent", logs: logs, sub: sub}, nil
}

// WatchMyEvent is a free log subscription operation binding the contract event 0x6c2b4666ba8da5a95717621d879a77de725f3d816709b9cbe9f059b8f875e284.
//
// Solidity: event MyEvent(uint256 value)
func (_Stroage *StroageFilterer) WatchMyEvent(opts *bind.WatchOpts, sink chan<- *StroageMyEvent) (event.Subscription, error) {

	logs, sub, err := _Stroage.contract.WatchLogs(opts, "MyEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StroageMyEvent)
				if err := _Stroage.contract.UnpackLog(event, "MyEvent", log); err != nil {
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

// ParseMyEvent is a log parse operation binding the contract event 0x6c2b4666ba8da5a95717621d879a77de725f3d816709b9cbe9f059b8f875e284.
//
// Solidity: event MyEvent(uint256 value)
func (_Stroage *StroageFilterer) ParseMyEvent(log types.Log) (*StroageMyEvent, error) {
	event := new(StroageMyEvent)
	if err := _Stroage.contract.UnpackLog(event, "MyEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
