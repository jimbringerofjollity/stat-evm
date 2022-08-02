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
)

// MedianMetaData contains all meta data concerning the Median contract.
var MedianMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"getMedian\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"1ff7dabc": "getMedian()",
	},
}

// MedianABI is the input ABI used to generate the binding from.
// Deprecated: Use MedianMetaData.ABI instead.
var MedianABI = MedianMetaData.ABI

// Deprecated: Use MedianMetaData.Sigs instead.
// MedianFuncSigs maps the 4-byte function signature to its string representation.
var MedianFuncSigs = MedianMetaData.Sigs

// Median is an auto generated Go binding around an Ethereum contract.
type Median struct {
	MedianCaller     // Read-only binding to the contract
	MedianTransactor // Write-only binding to the contract
	MedianFilterer   // Log filterer for contract events
}

// MedianCaller is an auto generated read-only Go binding around an Ethereum contract.
type MedianCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MedianTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MedianTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MedianFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MedianFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MedianSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MedianSession struct {
	Contract     *Median           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MedianCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MedianCallerSession struct {
	Contract *MedianCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MedianTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MedianTransactorSession struct {
	Contract     *MedianTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MedianRaw is an auto generated low-level Go binding around an Ethereum contract.
type MedianRaw struct {
	Contract *Median // Generic contract binding to access the raw methods on
}

// MedianCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MedianCallerRaw struct {
	Contract *MedianCaller // Generic read-only contract binding to access the raw methods on
}

// MedianTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MedianTransactorRaw struct {
	Contract *MedianTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMedian creates a new instance of Median, bound to a specific deployed contract.
func NewMedian(address common.Address, backend bind.ContractBackend) (*Median, error) {
	contract, err := bindMedian(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Median{MedianCaller: MedianCaller{contract: contract}, MedianTransactor: MedianTransactor{contract: contract}, MedianFilterer: MedianFilterer{contract: contract}}, nil
}

// NewMedianCaller creates a new read-only instance of Median, bound to a specific deployed contract.
func NewMedianCaller(address common.Address, caller bind.ContractCaller) (*MedianCaller, error) {
	contract, err := bindMedian(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MedianCaller{contract: contract}, nil
}

// NewMedianTransactor creates a new write-only instance of Median, bound to a specific deployed contract.
func NewMedianTransactor(address common.Address, transactor bind.ContractTransactor) (*MedianTransactor, error) {
	contract, err := bindMedian(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MedianTransactor{contract: contract}, nil
}

// NewMedianFilterer creates a new log filterer instance of Median, bound to a specific deployed contract.
func NewMedianFilterer(address common.Address, filterer bind.ContractFilterer) (*MedianFilterer, error) {
	contract, err := bindMedian(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MedianFilterer{contract: contract}, nil
}

// bindMedian binds a generic wrapper to an already deployed contract.
func bindMedian(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MedianABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Median *MedianRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Median.Contract.MedianCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Median *MedianRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Median.Contract.MedianTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Median *MedianRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Median.Contract.MedianTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Median *MedianCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Median.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Median *MedianTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Median.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Median *MedianTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Median.Contract.contract.Transact(opts, method, params...)
}

// GetMedian is a free data retrieval call binding the contract method 0x1ff7dabc.
//
// Solidity: function getMedian() view returns(bytes32)
func (_Median *MedianCaller) GetMedian(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Median.contract.Call(opts, &out, "getMedian")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetMedian is a free data retrieval call binding the contract method 0x1ff7dabc.
//
// Solidity: function getMedian() view returns(bytes32)
func (_Median *MedianSession) GetMedian() ([32]byte, error) {
	return _Median.Contract.GetMedian(&_Median.CallOpts)
}

// GetMedian is a free data retrieval call binding the contract method 0x1ff7dabc.
//
// Solidity: function getMedian() view returns(bytes32)
func (_Median *MedianCallerSession) GetMedian() ([32]byte, error) {
	return _Median.Contract.GetMedian(&_Median.CallOpts)
}

// TestMetaData contains all meta data concerning the Test contract.
var TestMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"res\",\"type\":\"bytes32\"}],\"name\":\"Debug\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"last\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"testMe\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"47799da8": "last()",
		"52415840": "testMe()",
	},
	Bin: "0x6080604052600180546001600160a01b03191673030000000000000000000000000000000000000117905534801561003657600080fd5b5061011e806100466000396000f3fe6080604052348015600f57600080fd5b506004361060325760003560e01c806347799da814603757806352415840146051575b600080fd5b603f60005481565b60405190815260200160405180910390f35b60576059565b005b600160009054906101000a90046001600160a01b03166001600160a01b0316631ff7dabc6040518163ffffffff1660e01b8152600401602060405180830381865afa15801560ab573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019060cd919060d0565b50565b60006020828403121560e157600080fd5b505191905056fea26469706673582212203a65a79ad7c5fc57b73236451f48cad842ff7488fde9ee29b729060ea7c9cd7c64736f6c634300080f0033",
}

// TestABI is the input ABI used to generate the binding from.
// Deprecated: Use TestMetaData.ABI instead.
var TestABI = TestMetaData.ABI

// Deprecated: Use TestMetaData.Sigs instead.
// TestFuncSigs maps the 4-byte function signature to its string representation.
var TestFuncSigs = TestMetaData.Sigs

// TestBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TestMetaData.Bin instead.
var TestBin = TestMetaData.Bin

// DeployTest deploys a new Ethereum contract, binding an instance of Test to it.
func DeployTest(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Test, error) {
	parsed, err := TestMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TestBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Test{TestCaller: TestCaller{contract: contract}, TestTransactor: TestTransactor{contract: contract}, TestFilterer: TestFilterer{contract: contract}}, nil
}

// Test is an auto generated Go binding around an Ethereum contract.
type Test struct {
	TestCaller     // Read-only binding to the contract
	TestTransactor // Write-only binding to the contract
	TestFilterer   // Log filterer for contract events
}

// TestCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestSession struct {
	Contract     *Test             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestCallerSession struct {
	Contract *TestCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestTransactorSession struct {
	Contract     *TestTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestRaw struct {
	Contract *Test // Generic contract binding to access the raw methods on
}

// TestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestCallerRaw struct {
	Contract *TestCaller // Generic read-only contract binding to access the raw methods on
}

// TestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestTransactorRaw struct {
	Contract *TestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTest creates a new instance of Test, bound to a specific deployed contract.
func NewTest(address common.Address, backend bind.ContractBackend) (*Test, error) {
	contract, err := bindTest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Test{TestCaller: TestCaller{contract: contract}, TestTransactor: TestTransactor{contract: contract}, TestFilterer: TestFilterer{contract: contract}}, nil
}

// NewTestCaller creates a new read-only instance of Test, bound to a specific deployed contract.
func NewTestCaller(address common.Address, caller bind.ContractCaller) (*TestCaller, error) {
	contract, err := bindTest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestCaller{contract: contract}, nil
}

// NewTestTransactor creates a new write-only instance of Test, bound to a specific deployed contract.
func NewTestTransactor(address common.Address, transactor bind.ContractTransactor) (*TestTransactor, error) {
	contract, err := bindTest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestTransactor{contract: contract}, nil
}

// NewTestFilterer creates a new log filterer instance of Test, bound to a specific deployed contract.
func NewTestFilterer(address common.Address, filterer bind.ContractFilterer) (*TestFilterer, error) {
	contract, err := bindTest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestFilterer{contract: contract}, nil
}

// bindTest binds a generic wrapper to an already deployed contract.
func bindTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TestABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Test *TestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Test.Contract.TestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Test *TestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Test.Contract.TestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Test *TestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Test.Contract.TestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Test *TestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Test.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Test *TestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Test.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Test *TestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Test.Contract.contract.Transact(opts, method, params...)
}

// Last is a free data retrieval call binding the contract method 0x47799da8.
//
// Solidity: function last() view returns(bytes32)
func (_Test *TestCaller) Last(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Test.contract.Call(opts, &out, "last")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Last is a free data retrieval call binding the contract method 0x47799da8.
//
// Solidity: function last() view returns(bytes32)
func (_Test *TestSession) Last() ([32]byte, error) {
	return _Test.Contract.Last(&_Test.CallOpts)
}

// Last is a free data retrieval call binding the contract method 0x47799da8.
//
// Solidity: function last() view returns(bytes32)
func (_Test *TestCallerSession) Last() ([32]byte, error) {
	return _Test.Contract.Last(&_Test.CallOpts)
}

// TestMe is a paid mutator transaction binding the contract method 0x52415840.
//
// Solidity: function testMe() returns()
func (_Test *TestTransactor) TestMe(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Test.contract.Transact(opts, "testMe")
}

// TestMe is a paid mutator transaction binding the contract method 0x52415840.
//
// Solidity: function testMe() returns()
func (_Test *TestSession) TestMe() (*types.Transaction, error) {
	return _Test.Contract.TestMe(&_Test.TransactOpts)
}

// TestMe is a paid mutator transaction binding the contract method 0x52415840.
//
// Solidity: function testMe() returns()
func (_Test *TestTransactorSession) TestMe() (*types.Transaction, error) {
	return _Test.Contract.TestMe(&_Test.TransactOpts)
}

// TestDebugIterator is returned from FilterDebug and is used to iterate over the raw logs and unpacked data for Debug events raised by the Test contract.
type TestDebugIterator struct {
	Event *TestDebug // Event containing the contract specifics and raw log

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
func (it *TestDebugIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestDebug)
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
		it.Event = new(TestDebug)
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
func (it *TestDebugIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestDebugIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestDebug represents a Debug event raised by the Test contract.
type TestDebug struct {
	Message string
	Res     [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDebug is a free log retrieval operation binding the contract event 0x9b727ad7a6d995527744fb3b7ccd68fd2d4a769c01209e041e1014abd3518113.
//
// Solidity: event Debug(string message, bytes32 res)
func (_Test *TestFilterer) FilterDebug(opts *bind.FilterOpts) (*TestDebugIterator, error) {

	logs, sub, err := _Test.contract.FilterLogs(opts, "Debug")
	if err != nil {
		return nil, err
	}
	return &TestDebugIterator{contract: _Test.contract, event: "Debug", logs: logs, sub: sub}, nil
}

// WatchDebug is a free log subscription operation binding the contract event 0x9b727ad7a6d995527744fb3b7ccd68fd2d4a769c01209e041e1014abd3518113.
//
// Solidity: event Debug(string message, bytes32 res)
func (_Test *TestFilterer) WatchDebug(opts *bind.WatchOpts, sink chan<- *TestDebug) (event.Subscription, error) {

	logs, sub, err := _Test.contract.WatchLogs(opts, "Debug")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestDebug)
				if err := _Test.contract.UnpackLog(event, "Debug", log); err != nil {
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

// ParseDebug is a log parse operation binding the contract event 0x9b727ad7a6d995527744fb3b7ccd68fd2d4a769c01209e041e1014abd3518113.
//
// Solidity: event Debug(string message, bytes32 res)
func (_Test *TestFilterer) ParseDebug(log types.Log) (*TestDebug, error) {
	event := new(TestDebug)
	if err := _Test.contract.UnpackLog(event, "Debug", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

