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

// FitMetaData contains all meta data concerning the Fit contract.
var FitMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"v1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"v2\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"v3\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"v4\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"v5\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"v6\",\"type\":\"uint256\"}],\"name\":\"getFit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"aba95e3b": "getFit(uint256,uint256,uint256,uint256,uint256,uint256)",
	},
}

// FitABI is the input ABI used to generate the binding from.
// Deprecated: Use FitMetaData.ABI instead.
var FitABI = FitMetaData.ABI

// Deprecated: Use FitMetaData.Sigs instead.
// FitFuncSigs maps the 4-byte function signature to its string representation.
var FitFuncSigs = FitMetaData.Sigs

// Fit is an auto generated Go binding around an Ethereum contract.
type Fit struct {
	FitCaller     // Read-only binding to the contract
	FitTransactor // Write-only binding to the contract
	FitFilterer   // Log filterer for contract events
}

// FitCaller is an auto generated read-only Go binding around an Ethereum contract.
type FitCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FitTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FitTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FitFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FitFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FitSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FitSession struct {
	Contract     *Fit              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FitCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FitCallerSession struct {
	Contract *FitCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// FitTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FitTransactorSession struct {
	Contract     *FitTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FitRaw is an auto generated low-level Go binding around an Ethereum contract.
type FitRaw struct {
	Contract *Fit // Generic contract binding to access the raw methods on
}

// FitCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FitCallerRaw struct {
	Contract *FitCaller // Generic read-only contract binding to access the raw methods on
}

// FitTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FitTransactorRaw struct {
	Contract *FitTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFit creates a new instance of Fit, bound to a specific deployed contract.
func NewFit(address common.Address, backend bind.ContractBackend) (*Fit, error) {
	contract, err := bindFit(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Fit{FitCaller: FitCaller{contract: contract}, FitTransactor: FitTransactor{contract: contract}, FitFilterer: FitFilterer{contract: contract}}, nil
}

// NewFitCaller creates a new read-only instance of Fit, bound to a specific deployed contract.
func NewFitCaller(address common.Address, caller bind.ContractCaller) (*FitCaller, error) {
	contract, err := bindFit(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FitCaller{contract: contract}, nil
}

// NewFitTransactor creates a new write-only instance of Fit, bound to a specific deployed contract.
func NewFitTransactor(address common.Address, transactor bind.ContractTransactor) (*FitTransactor, error) {
	contract, err := bindFit(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FitTransactor{contract: contract}, nil
}

// NewFitFilterer creates a new log filterer instance of Fit, bound to a specific deployed contract.
func NewFitFilterer(address common.Address, filterer bind.ContractFilterer) (*FitFilterer, error) {
	contract, err := bindFit(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FitFilterer{contract: contract}, nil
}

// bindFit binds a generic wrapper to an already deployed contract.
func bindFit(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FitABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Fit *FitRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Fit.Contract.FitCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Fit *FitRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fit.Contract.FitTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Fit *FitRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Fit.Contract.FitTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Fit *FitCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Fit.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Fit *FitTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fit.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Fit *FitTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Fit.Contract.contract.Transact(opts, method, params...)
}

// GetFit is a free data retrieval call binding the contract method 0xaba95e3b.
//
// Solidity: function getFit(uint256 v1, uint256 v2, uint256 v3, uint256 v4, uint256 v5, uint256 v6) view returns(uint256)
func (_Fit *FitCaller) GetFit(opts *bind.CallOpts, v1 *big.Int, v2 *big.Int, v3 *big.Int, v4 *big.Int, v5 *big.Int, v6 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Fit.contract.Call(opts, &out, "getFit", v1, v2, v3, v4, v5, v6)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFit is a free data retrieval call binding the contract method 0xaba95e3b.
//
// Solidity: function getFit(uint256 v1, uint256 v2, uint256 v3, uint256 v4, uint256 v5, uint256 v6) view returns(uint256)
func (_Fit *FitSession) GetFit(v1 *big.Int, v2 *big.Int, v3 *big.Int, v4 *big.Int, v5 *big.Int, v6 *big.Int) (*big.Int, error) {
	return _Fit.Contract.GetFit(&_Fit.CallOpts, v1, v2, v3, v4, v5, v6)
}

// GetFit is a free data retrieval call binding the contract method 0xaba95e3b.
//
// Solidity: function getFit(uint256 v1, uint256 v2, uint256 v3, uint256 v4, uint256 v5, uint256 v6) view returns(uint256)
func (_Fit *FitCallerSession) GetFit(v1 *big.Int, v2 *big.Int, v3 *big.Int, v4 *big.Int, v5 *big.Int, v6 *big.Int) (*big.Int, error) {
	return _Fit.Contract.GetFit(&_Fit.CallOpts, v1, v2, v3, v4, v5, v6)
}

// TestMetaData contains all meta data concerning the Test contract.
var TestMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"res\",\"type\":\"bytes32\"}],\"name\":\"Debug\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"last\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"v1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"v2\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"v3\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"v4\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"v5\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"v6\",\"type\":\"uint256\"}],\"name\":\"testMe\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"47799da8": "last()",
		"5a7c4229": "testMe(uint256,uint256,uint256,uint256,uint256,uint256)",
	},
	Bin: "0x6080604052600180546001600160a01b03191673030000000000000000000000000000000000000117905534801561003657600080fd5b50610198806100466000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c806347799da81461003b5780635a7c422914610056575b600080fd5b61004460005481565b60405190815260200160405180910390f35b610069610064366004610106565b61006b565b005b60015460405163aba95e3b60e01b8152600481018890526024810187905260448101869052606481018590526084810184905260a481018390526001600160a01b039091169063aba95e3b9060c401602060405180830381865afa1580156100d7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906100fb9190610149565b600055505050505050565b60008060008060008060c0878903121561011f57600080fd5b505084359660208601359650604086013595606081013595506080810135945060a0013592509050565b60006020828403121561015b57600080fd5b505191905056fea2646970667358221220d166f74017985dbd4896ee37afa506d33ad1ef41421ae310f0555b0c7e4a047864736f6c634300080f0033",
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
// Solidity: function last() view returns(uint256)
func (_Test *TestCaller) Last(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Test.contract.Call(opts, &out, "last")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Last is a free data retrieval call binding the contract method 0x47799da8.
//
// Solidity: function last() view returns(uint256)
func (_Test *TestSession) Last() (*big.Int, error) {
	return _Test.Contract.Last(&_Test.CallOpts)
}

// Last is a free data retrieval call binding the contract method 0x47799da8.
//
// Solidity: function last() view returns(uint256)
func (_Test *TestCallerSession) Last() (*big.Int, error) {
	return _Test.Contract.Last(&_Test.CallOpts)
}

// TestMe is a paid mutator transaction binding the contract method 0x5a7c4229.
//
// Solidity: function testMe(uint256 v1, uint256 v2, uint256 v3, uint256 v4, uint256 v5, uint256 v6) returns()
func (_Test *TestTransactor) TestMe(opts *bind.TransactOpts, v1 *big.Int, v2 *big.Int, v3 *big.Int, v4 *big.Int, v5 *big.Int, v6 *big.Int) (*types.Transaction, error) {
	return _Test.contract.Transact(opts, "testMe", v1, v2, v3, v4, v5, v6)
}

// TestMe is a paid mutator transaction binding the contract method 0x5a7c4229.
//
// Solidity: function testMe(uint256 v1, uint256 v2, uint256 v3, uint256 v4, uint256 v5, uint256 v6) returns()
func (_Test *TestSession) TestMe(v1 *big.Int, v2 *big.Int, v3 *big.Int, v4 *big.Int, v5 *big.Int, v6 *big.Int) (*types.Transaction, error) {
	return _Test.Contract.TestMe(&_Test.TransactOpts, v1, v2, v3, v4, v5, v6)
}

// TestMe is a paid mutator transaction binding the contract method 0x5a7c4229.
//
// Solidity: function testMe(uint256 v1, uint256 v2, uint256 v3, uint256 v4, uint256 v5, uint256 v6) returns()
func (_Test *TestTransactorSession) TestMe(v1 *big.Int, v2 *big.Int, v3 *big.Int, v4 *big.Int, v5 *big.Int, v6 *big.Int) (*types.Transaction, error) {
	return _Test.Contract.TestMe(&_Test.TransactOpts, v1, v2, v3, v4, v5, v6)
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

