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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"v1\",\"type\":\"uint256[]\"}],\"name\":\"getMedian\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"e8787f41": "getMedian(uint256[])",
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

// GetMedian is a free data retrieval call binding the contract method 0xe8787f41.
//
// Solidity: function getMedian(uint256[] v1) view returns(uint256)
func (_Median *MedianCaller) GetMedian(opts *bind.CallOpts, v1 []*big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Median.contract.Call(opts, &out, "getMedian", v1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMedian is a free data retrieval call binding the contract method 0xe8787f41.
//
// Solidity: function getMedian(uint256[] v1) view returns(uint256)
func (_Median *MedianSession) GetMedian(v1 []*big.Int) (*big.Int, error) {
	return _Median.Contract.GetMedian(&_Median.CallOpts, v1)
}

// GetMedian is a free data retrieval call binding the contract method 0xe8787f41.
//
// Solidity: function getMedian(uint256[] v1) view returns(uint256)
func (_Median *MedianCallerSession) GetMedian(v1 []*big.Int) (*big.Int, error) {
	return _Median.Contract.GetMedian(&_Median.CallOpts, v1)
}

// TestMetaData contains all meta data concerning the Test contract.
var TestMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"med\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"vals\",\"type\":\"uint256[]\"}],\"name\":\"testMe\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"cd53c747": "med()",
		"517cd0b3": "testMe(uint256[])",
	},
	Bin: "0x6080604052600180546001600160a01b03191673030000000000000000000000000000000000000117905534801561003657600080fd5b50610249806100466000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c8063517cd0b31461003b578063cd53c74714610050575b600080fd5b61004e6100493660046100f8565b61006b565b005b61005960005481565b60405190815260200160405180910390f35b60015460405163e8787f4160e01b81526001600160a01b039091169063e8787f419061009b9084906004016101b6565b602060405180830381865afa1580156100b8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906100dc91906101fa565b60005550565b634e487b7160e01b600052604160045260246000fd5b6000602080838503121561010b57600080fd5b823567ffffffffffffffff8082111561012357600080fd5b818501915085601f83011261013757600080fd5b813581811115610149576101496100e2565b8060051b604051601f19603f8301168101818110858211171561016e5761016e6100e2565b60405291825284820192508381018501918883111561018c57600080fd5b938501935b828510156101aa57843584529385019392850192610191565b98975050505050505050565b6020808252825182820181905260009190848201906040850190845b818110156101ee578351835292840192918401916001016101d2565b50909695505050505050565b60006020828403121561020c57600080fd5b505191905056fea264697066735822122051d8afcf870d2655cba0124e02b39aae0aaa5f86cc9018138e315d3b7bfcbf6364736f6c634300080f0033",
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

// Med is a free data retrieval call binding the contract method 0xcd53c747.
//
// Solidity: function med() view returns(uint256)
func (_Test *TestCaller) Med(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Test.contract.Call(opts, &out, "med")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Med is a free data retrieval call binding the contract method 0xcd53c747.
//
// Solidity: function med() view returns(uint256)
func (_Test *TestSession) Med() (*big.Int, error) {
	return _Test.Contract.Med(&_Test.CallOpts)
}

// Med is a free data retrieval call binding the contract method 0xcd53c747.
//
// Solidity: function med() view returns(uint256)
func (_Test *TestCallerSession) Med() (*big.Int, error) {
	return _Test.Contract.Med(&_Test.CallOpts)
}

// TestMe is a paid mutator transaction binding the contract method 0x517cd0b3.
//
// Solidity: function testMe(uint256[] vals) returns()
func (_Test *TestTransactor) TestMe(opts *bind.TransactOpts, vals []*big.Int) (*types.Transaction, error) {
	return _Test.contract.Transact(opts, "testMe", vals)
}

// TestMe is a paid mutator transaction binding the contract method 0x517cd0b3.
//
// Solidity: function testMe(uint256[] vals) returns()
func (_Test *TestSession) TestMe(vals []*big.Int) (*types.Transaction, error) {
	return _Test.Contract.TestMe(&_Test.TransactOpts, vals)
}

// TestMe is a paid mutator transaction binding the contract method 0x517cd0b3.
//
// Solidity: function testMe(uint256[] vals) returns()
func (_Test *TestTransactorSession) TestMe(vals []*big.Int) (*types.Transaction, error) {
	return _Test.Contract.TestMe(&_Test.TransactOpts, vals)
}

