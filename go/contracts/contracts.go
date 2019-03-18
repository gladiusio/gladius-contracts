// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// MarketABI is the input ABI used to generate the binding from.
const MarketABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_ownerAddress\",\"type\":\"address\"}],\"name\":\"getOwnerMarketPools\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getData\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_data\",\"type\":\"string\"}],\"name\":\"setData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getPoolFactory\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_factory\",\"type\":\"address\"}],\"name\":\"changePoolFactory\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"data\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"name\":\"getPoolOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"createPool\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_ownerAddress\",\"type\":\"address\"}],\"name\":\"getOwnerAllPools\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getMarketPools\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_poolAddress\",\"type\":\"address\"}],\"name\":\"addToMarketplace\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAllPools\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"marketPoolOwners\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"marketPoolsList\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_poolFactory\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// MarketBin is the compiled bytecode used for deploying new contracts.
const MarketBin = `0x60806040526005809080546001816001161561010002031660029004610026929190610090565b5034801561003357600080fd5b506040516040806110108339810180604052604081101561005357600080fd5b50805160209091015160028054600160a060020a03938416600160a060020a03199182161790915560038054939092169216919091179055610132565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106100c95780548555610105565b8280016001018555821561010557600052602060002091601f016020900482015b828111156101055782548255916001019190600101906100ea565b50610111929150610115565b5090565b61012f91905b80821115610111576000815560010161011b565b90565b610ecf806101416000396000f3fe608060405234801561001057600080fd5b5060043610610128576000357c0100000000000000000000000000000000000000000000000000000000900480638da5cb5b116100bf578063b42cb6f51161008e578063b42cb6f5146103a4578063d00e1e72146103ac578063d88ff1f4146103e6578063db20ee33146103ee578063e9e242341461041a57610128565b80638da5cb5b146103485780639a06b11314610350578063a6f9dae114610358578063aebce2d41461037e57610128565b806372591be5116100fb57806372591be5146102ec57806373d4a13a146103125780637cf255171461031a578063893d20e81461034057610128565b806317e421c81461012d5780633bc5de30146101a357806347064d6a146102205780636bd969b3146102c8575b600080fd5b6101536004803603602081101561014357600080fd5b5035600160a060020a0316610437565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561018f578181015183820152602001610177565b505050509050019250505060405180910390f35b6101ab6104ab565b6040805160208082528351818301528351919283929083019185019080838360005b838110156101e55781810151838201526020016101cd565b50505050905090810190601f1680156102125780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6102c66004803603602081101561023657600080fd5b81019060208101813564010000000081111561025157600080fd5b82018360208201111561026357600080fd5b8035906020019184600183028401116401000000008311171561028557600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610542945050505050565b005b6102d061060b565b60408051600160a060020a039092168252519081900360200190f35b6102c66004803603602081101561030257600080fd5b5035600160a060020a031661061a565b6101ab6106c1565b6102d06004803603602081101561033057600080fd5b5035600160a060020a031661074f565b6102d06107eb565b6102d06107fa565b6102d0610809565b6102c66004803603602081101561036e57600080fd5b5035600160a060020a03166108a5565b6101536004803603602081101561039457600080fd5b5035600160a060020a031661094c565b610153610a47565b6103d2600480360360208110156103c257600080fd5b5035600160a060020a0316610aa8565b604080519115158252519081900360200190f35b610153610cbe565b6102d06004803603604081101561040457600080fd5b50600160a060020a038135169060200135610dac565b6102d06004803603602081101561043057600080fd5b5035610de3565b600160a060020a0381166000908152602081815260409182902080548351818402810184019094528084526060939283018282801561049f57602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311610481575b50505050509050919050565b60048054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156105375780601f1061050c57610100808354040283529160200191610537565b820191906000526020600020905b81548152906001019060200180831161051a57829003601f168201915b505050505090505b90565b600254600590600160a060020a031633146105f35760405160e560020a62461bcd0281526020600482019081528254600260001961010060018416150201909116046024830181905290918291604490910190849080156105e45780601f106105b9576101008083540402835291602001916105e4565b820191906000526020600020905b8154815290600101906020018083116105c757829003601f168201915b50509250505060405180910390fd5b508051610607906004906020840190610e0b565b5050565b600354600160a060020a031690565b600254600590600160a060020a031633146106915760405160e560020a62461bcd0281526020600482019081528254600260001961010060018416150201909116046024830181905290918291604490910190849080156105e45780601f106105b9576101008083540402835291602001916105e4565b506003805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b6004805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156107475780601f1061071c57610100808354040283529160200191610747565b820191906000526020600020905b81548152906001019060200180831161072a57829003601f168201915b505050505081565b600354604080517f7cf25517000000000000000000000000000000000000000000000000000000008152600160a060020a03848116600483015291516000939290921691637cf2551791602480820192602092909190829003018186803b1580156107b957600080fd5b505afa1580156107cd573d6000803e3d6000fd5b505050506040513d60208110156107e357600080fd5b505192915050565b600254600160a060020a031690565b600254600160a060020a031681565b600354604080517f9049f9d200000000000000000000000000000000000000000000000000000000815233600482015290516000928392600160a060020a0390911691639049f9d29160248082019260209290919082900301818787803b15801561087357600080fd5b505af1158015610887573d6000803e3d6000fd5b505050506040513d602081101561089d57600080fd5b505191505090565b600254600590600160a060020a0316331461091c5760405160e560020a62461bcd0281526020600482019081528254600260001961010060018416150201909116046024830181905290918291604490910190849080156105e45780601f106105b9576101008083540402835291602001916105e4565b506002805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600354604080517f556f77db000000000000000000000000000000000000000000000000000000008152600160a060020a0384811660048301529151606093929092169163556f77db91602480820192600092909190829003018186803b1580156109b657600080fd5b505afa1580156109ca573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260208110156109f357600080fd5b810190808051640100000000811115610a0b57600080fd5b82016020810184811115610a1e57600080fd5b8151856020820283011164010000000082111715610a3b57600080fd5b50909695505050505050565b6060600180548060200260200160405190810160405280929190818152602001828054801561053757602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311610a81575050505050905090565b600254600090600590600160a060020a03163314610b225760405160e560020a62461bcd0281526020600482019081528254600260001961010060018416150201909116046024830181905290918291604490910190849080156105e45780601f106105b9576101008083540402835291602001916105e4565b50600082905060008082600160a060020a031663893d20e86040518163ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040160206040518083038186803b158015610b8057600080fd5b505afa158015610b94573d6000803e3d6000fd5b505050506040513d6020811015610baa57600080fd5b5051600160a060020a0390811682526020808301939093526040918201600090812080546001808201835591835291859020909101805473ffffffffffffffffffffffffffffffffffffffff1916928616928317905582517f893d20e800000000000000000000000000000000000000000000000000000000815292519093919263893d20e8926004808301939192829003018186803b158015610c4d57600080fd5b505afa158015610c61573d6000803e3d6000fd5b505050506040513d6020811015610c7757600080fd5b50518154600180820184556000938452602090932001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a039092169190911790559392505050565b600354604080517fd88ff1f40000000000000000000000000000000000000000000000000000000081529051606092600160a060020a03169163d88ff1f4916004808301926000929190829003018186803b158015610d1c57600080fd5b505afa158015610d30573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526020811015610d5957600080fd5b810190808051640100000000811115610d7157600080fd5b82016020810184811115610d8457600080fd5b8151856020820283011164010000000082111715610da157600080fd5b509094505050505090565b600060205281600052604060002081815481101515610dc757fe5b600091825260209091200154600160a060020a03169150829050565b6001805482908110610df157fe5b600091825260209091200154600160a060020a0316905081565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610e4c57805160ff1916838001178555610e79565b82800160010185558215610e79579182015b82811115610e79578251825591602001919060010190610e5e565b50610e85929150610e89565b5090565b61053f91905b80821115610e855760008155600101610e8f56fea165627a7a723058208fe72988acdfec8eff19cb3732650314f16e30d3e39546a0a91bbcd3060afd440029`

// DeployMarket deploys a new Ethereum contract, binding an instance of Market to it.
func DeployMarket(auth *bind.TransactOpts, backend bind.ContractBackend, _owner common.Address, _poolFactory common.Address) (common.Address, *types.Transaction, *Market, error) {
	parsed, err := abi.JSON(strings.NewReader(MarketABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MarketBin), backend, _owner, _poolFactory)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Market{MarketCaller: MarketCaller{contract: contract}, MarketTransactor: MarketTransactor{contract: contract}, MarketFilterer: MarketFilterer{contract: contract}}, nil
}

// Market is an auto generated Go binding around an Ethereum contract.
type Market struct {
	MarketCaller     // Read-only binding to the contract
	MarketTransactor // Write-only binding to the contract
	MarketFilterer   // Log filterer for contract events
}

// MarketCaller is an auto generated read-only Go binding around an Ethereum contract.
type MarketCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MarketTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MarketFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MarketSession struct {
	Contract     *Market           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MarketCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MarketCallerSession struct {
	Contract *MarketCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MarketTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MarketTransactorSession struct {
	Contract     *MarketTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MarketRaw is an auto generated low-level Go binding around an Ethereum contract.
type MarketRaw struct {
	Contract *Market // Generic contract binding to access the raw methods on
}

// MarketCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MarketCallerRaw struct {
	Contract *MarketCaller // Generic read-only contract binding to access the raw methods on
}

// MarketTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MarketTransactorRaw struct {
	Contract *MarketTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMarket creates a new instance of Market, bound to a specific deployed contract.
func NewMarket(address common.Address, backend bind.ContractBackend) (*Market, error) {
	contract, err := bindMarket(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Market{MarketCaller: MarketCaller{contract: contract}, MarketTransactor: MarketTransactor{contract: contract}, MarketFilterer: MarketFilterer{contract: contract}}, nil
}

// NewMarketCaller creates a new read-only instance of Market, bound to a specific deployed contract.
func NewMarketCaller(address common.Address, caller bind.ContractCaller) (*MarketCaller, error) {
	contract, err := bindMarket(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MarketCaller{contract: contract}, nil
}

// NewMarketTransactor creates a new write-only instance of Market, bound to a specific deployed contract.
func NewMarketTransactor(address common.Address, transactor bind.ContractTransactor) (*MarketTransactor, error) {
	contract, err := bindMarket(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MarketTransactor{contract: contract}, nil
}

// NewMarketFilterer creates a new log filterer instance of Market, bound to a specific deployed contract.
func NewMarketFilterer(address common.Address, filterer bind.ContractFilterer) (*MarketFilterer, error) {
	contract, err := bindMarket(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MarketFilterer{contract: contract}, nil
}

// bindMarket binds a generic wrapper to an already deployed contract.
func bindMarket(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MarketABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Market *MarketRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Market.Contract.MarketCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Market *MarketRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Market.Contract.MarketTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Market *MarketRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Market.Contract.MarketTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Market *MarketCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Market.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Market *MarketTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Market.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Market *MarketTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Market.Contract.contract.Transact(opts, method, params...)
}

// Data is a free data retrieval call binding the contract method 0x73d4a13a.
//
// Solidity: function data() constant returns(string)
func (_Market *MarketCaller) Data(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Market.contract.Call(opts, out, "data")
	return *ret0, err
}

// Data is a free data retrieval call binding the contract method 0x73d4a13a.
//
// Solidity: function data() constant returns(string)
func (_Market *MarketSession) Data() (string, error) {
	return _Market.Contract.Data(&_Market.CallOpts)
}

// Data is a free data retrieval call binding the contract method 0x73d4a13a.
//
// Solidity: function data() constant returns(string)
func (_Market *MarketCallerSession) Data() (string, error) {
	return _Market.Contract.Data(&_Market.CallOpts)
}

// GetAllPools is a free data retrieval call binding the contract method 0xd88ff1f4.
//
// Solidity: function getAllPools() constant returns(address[])
func (_Market *MarketCaller) GetAllPools(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Market.contract.Call(opts, out, "getAllPools")
	return *ret0, err
}

// GetAllPools is a free data retrieval call binding the contract method 0xd88ff1f4.
//
// Solidity: function getAllPools() constant returns(address[])
func (_Market *MarketSession) GetAllPools() ([]common.Address, error) {
	return _Market.Contract.GetAllPools(&_Market.CallOpts)
}

// GetAllPools is a free data retrieval call binding the contract method 0xd88ff1f4.
//
// Solidity: function getAllPools() constant returns(address[])
func (_Market *MarketCallerSession) GetAllPools() ([]common.Address, error) {
	return _Market.Contract.GetAllPools(&_Market.CallOpts)
}

// GetData is a free data retrieval call binding the contract method 0x3bc5de30.
//
// Solidity: function getData() constant returns(string)
func (_Market *MarketCaller) GetData(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Market.contract.Call(opts, out, "getData")
	return *ret0, err
}

// GetData is a free data retrieval call binding the contract method 0x3bc5de30.
//
// Solidity: function getData() constant returns(string)
func (_Market *MarketSession) GetData() (string, error) {
	return _Market.Contract.GetData(&_Market.CallOpts)
}

// GetData is a free data retrieval call binding the contract method 0x3bc5de30.
//
// Solidity: function getData() constant returns(string)
func (_Market *MarketCallerSession) GetData() (string, error) {
	return _Market.Contract.GetData(&_Market.CallOpts)
}

// GetMarketPools is a free data retrieval call binding the contract method 0xb42cb6f5.
//
// Solidity: function getMarketPools() constant returns(address[])
func (_Market *MarketCaller) GetMarketPools(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Market.contract.Call(opts, out, "getMarketPools")
	return *ret0, err
}

// GetMarketPools is a free data retrieval call binding the contract method 0xb42cb6f5.
//
// Solidity: function getMarketPools() constant returns(address[])
func (_Market *MarketSession) GetMarketPools() ([]common.Address, error) {
	return _Market.Contract.GetMarketPools(&_Market.CallOpts)
}

// GetMarketPools is a free data retrieval call binding the contract method 0xb42cb6f5.
//
// Solidity: function getMarketPools() constant returns(address[])
func (_Market *MarketCallerSession) GetMarketPools() ([]common.Address, error) {
	return _Market.Contract.GetMarketPools(&_Market.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_Market *MarketCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Market.contract.Call(opts, out, "getOwner")
	return *ret0, err
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_Market *MarketSession) GetOwner() (common.Address, error) {
	return _Market.Contract.GetOwner(&_Market.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_Market *MarketCallerSession) GetOwner() (common.Address, error) {
	return _Market.Contract.GetOwner(&_Market.CallOpts)
}

// GetOwnerAllPools is a free data retrieval call binding the contract method 0xaebce2d4.
//
// Solidity: function getOwnerAllPools(address _ownerAddress) constant returns(address[])
func (_Market *MarketCaller) GetOwnerAllPools(opts *bind.CallOpts, _ownerAddress common.Address) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Market.contract.Call(opts, out, "getOwnerAllPools", _ownerAddress)
	return *ret0, err
}

// GetOwnerAllPools is a free data retrieval call binding the contract method 0xaebce2d4.
//
// Solidity: function getOwnerAllPools(address _ownerAddress) constant returns(address[])
func (_Market *MarketSession) GetOwnerAllPools(_ownerAddress common.Address) ([]common.Address, error) {
	return _Market.Contract.GetOwnerAllPools(&_Market.CallOpts, _ownerAddress)
}

// GetOwnerAllPools is a free data retrieval call binding the contract method 0xaebce2d4.
//
// Solidity: function getOwnerAllPools(address _ownerAddress) constant returns(address[])
func (_Market *MarketCallerSession) GetOwnerAllPools(_ownerAddress common.Address) ([]common.Address, error) {
	return _Market.Contract.GetOwnerAllPools(&_Market.CallOpts, _ownerAddress)
}

// GetOwnerMarketPools is a free data retrieval call binding the contract method 0x17e421c8.
//
// Solidity: function getOwnerMarketPools(address _ownerAddress) constant returns(address[])
func (_Market *MarketCaller) GetOwnerMarketPools(opts *bind.CallOpts, _ownerAddress common.Address) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Market.contract.Call(opts, out, "getOwnerMarketPools", _ownerAddress)
	return *ret0, err
}

// GetOwnerMarketPools is a free data retrieval call binding the contract method 0x17e421c8.
//
// Solidity: function getOwnerMarketPools(address _ownerAddress) constant returns(address[])
func (_Market *MarketSession) GetOwnerMarketPools(_ownerAddress common.Address) ([]common.Address, error) {
	return _Market.Contract.GetOwnerMarketPools(&_Market.CallOpts, _ownerAddress)
}

// GetOwnerMarketPools is a free data retrieval call binding the contract method 0x17e421c8.
//
// Solidity: function getOwnerMarketPools(address _ownerAddress) constant returns(address[])
func (_Market *MarketCallerSession) GetOwnerMarketPools(_ownerAddress common.Address) ([]common.Address, error) {
	return _Market.Contract.GetOwnerMarketPools(&_Market.CallOpts, _ownerAddress)
}

// GetPoolFactory is a free data retrieval call binding the contract method 0x6bd969b3.
//
// Solidity: function getPoolFactory() constant returns(address)
func (_Market *MarketCaller) GetPoolFactory(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Market.contract.Call(opts, out, "getPoolFactory")
	return *ret0, err
}

// GetPoolFactory is a free data retrieval call binding the contract method 0x6bd969b3.
//
// Solidity: function getPoolFactory() constant returns(address)
func (_Market *MarketSession) GetPoolFactory() (common.Address, error) {
	return _Market.Contract.GetPoolFactory(&_Market.CallOpts)
}

// GetPoolFactory is a free data retrieval call binding the contract method 0x6bd969b3.
//
// Solidity: function getPoolFactory() constant returns(address)
func (_Market *MarketCallerSession) GetPoolFactory() (common.Address, error) {
	return _Market.Contract.GetPoolFactory(&_Market.CallOpts)
}

// GetPoolOwner is a free data retrieval call binding the contract method 0x7cf25517.
//
// Solidity: function getPoolOwner(address _pool) constant returns(address)
func (_Market *MarketCaller) GetPoolOwner(opts *bind.CallOpts, _pool common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Market.contract.Call(opts, out, "getPoolOwner", _pool)
	return *ret0, err
}

// GetPoolOwner is a free data retrieval call binding the contract method 0x7cf25517.
//
// Solidity: function getPoolOwner(address _pool) constant returns(address)
func (_Market *MarketSession) GetPoolOwner(_pool common.Address) (common.Address, error) {
	return _Market.Contract.GetPoolOwner(&_Market.CallOpts, _pool)
}

// GetPoolOwner is a free data retrieval call binding the contract method 0x7cf25517.
//
// Solidity: function getPoolOwner(address _pool) constant returns(address)
func (_Market *MarketCallerSession) GetPoolOwner(_pool common.Address) (common.Address, error) {
	return _Market.Contract.GetPoolOwner(&_Market.CallOpts, _pool)
}

// MarketPoolOwners is a free data retrieval call binding the contract method 0xdb20ee33.
//
// Solidity: function marketPoolOwners(address , uint256 ) constant returns(address)
func (_Market *MarketCaller) MarketPoolOwners(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Market.contract.Call(opts, out, "marketPoolOwners", arg0, arg1)
	return *ret0, err
}

// MarketPoolOwners is a free data retrieval call binding the contract method 0xdb20ee33.
//
// Solidity: function marketPoolOwners(address , uint256 ) constant returns(address)
func (_Market *MarketSession) MarketPoolOwners(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _Market.Contract.MarketPoolOwners(&_Market.CallOpts, arg0, arg1)
}

// MarketPoolOwners is a free data retrieval call binding the contract method 0xdb20ee33.
//
// Solidity: function marketPoolOwners(address , uint256 ) constant returns(address)
func (_Market *MarketCallerSession) MarketPoolOwners(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _Market.Contract.MarketPoolOwners(&_Market.CallOpts, arg0, arg1)
}

// MarketPoolsList is a free data retrieval call binding the contract method 0xe9e24234.
//
// Solidity: function marketPoolsList(uint256 ) constant returns(address)
func (_Market *MarketCaller) MarketPoolsList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Market.contract.Call(opts, out, "marketPoolsList", arg0)
	return *ret0, err
}

// MarketPoolsList is a free data retrieval call binding the contract method 0xe9e24234.
//
// Solidity: function marketPoolsList(uint256 ) constant returns(address)
func (_Market *MarketSession) MarketPoolsList(arg0 *big.Int) (common.Address, error) {
	return _Market.Contract.MarketPoolsList(&_Market.CallOpts, arg0)
}

// MarketPoolsList is a free data retrieval call binding the contract method 0xe9e24234.
//
// Solidity: function marketPoolsList(uint256 ) constant returns(address)
func (_Market *MarketCallerSession) MarketPoolsList(arg0 *big.Int) (common.Address, error) {
	return _Market.Contract.MarketPoolsList(&_Market.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Market *MarketCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Market.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Market *MarketSession) Owner() (common.Address, error) {
	return _Market.Contract.Owner(&_Market.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Market *MarketCallerSession) Owner() (common.Address, error) {
	return _Market.Contract.Owner(&_Market.CallOpts)
}

// AddToMarketplace is a paid mutator transaction binding the contract method 0xd00e1e72.
//
// Solidity: function addToMarketplace(address _poolAddress) returns(bool)
func (_Market *MarketTransactor) AddToMarketplace(opts *bind.TransactOpts, _poolAddress common.Address) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "addToMarketplace", _poolAddress)
}

// AddToMarketplace is a paid mutator transaction binding the contract method 0xd00e1e72.
//
// Solidity: function addToMarketplace(address _poolAddress) returns(bool)
func (_Market *MarketSession) AddToMarketplace(_poolAddress common.Address) (*types.Transaction, error) {
	return _Market.Contract.AddToMarketplace(&_Market.TransactOpts, _poolAddress)
}

// AddToMarketplace is a paid mutator transaction binding the contract method 0xd00e1e72.
//
// Solidity: function addToMarketplace(address _poolAddress) returns(bool)
func (_Market *MarketTransactorSession) AddToMarketplace(_poolAddress common.Address) (*types.Transaction, error) {
	return _Market.Contract.AddToMarketplace(&_Market.TransactOpts, _poolAddress)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address _owner) returns()
func (_Market *MarketTransactor) ChangeOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "changeOwner", _owner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address _owner) returns()
func (_Market *MarketSession) ChangeOwner(_owner common.Address) (*types.Transaction, error) {
	return _Market.Contract.ChangeOwner(&_Market.TransactOpts, _owner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address _owner) returns()
func (_Market *MarketTransactorSession) ChangeOwner(_owner common.Address) (*types.Transaction, error) {
	return _Market.Contract.ChangeOwner(&_Market.TransactOpts, _owner)
}

// ChangePoolFactory is a paid mutator transaction binding the contract method 0x72591be5.
//
// Solidity: function changePoolFactory(address _factory) returns()
func (_Market *MarketTransactor) ChangePoolFactory(opts *bind.TransactOpts, _factory common.Address) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "changePoolFactory", _factory)
}

// ChangePoolFactory is a paid mutator transaction binding the contract method 0x72591be5.
//
// Solidity: function changePoolFactory(address _factory) returns()
func (_Market *MarketSession) ChangePoolFactory(_factory common.Address) (*types.Transaction, error) {
	return _Market.Contract.ChangePoolFactory(&_Market.TransactOpts, _factory)
}

// ChangePoolFactory is a paid mutator transaction binding the contract method 0x72591be5.
//
// Solidity: function changePoolFactory(address _factory) returns()
func (_Market *MarketTransactorSession) ChangePoolFactory(_factory common.Address) (*types.Transaction, error) {
	return _Market.Contract.ChangePoolFactory(&_Market.TransactOpts, _factory)
}

// CreatePool is a paid mutator transaction binding the contract method 0x9a06b113.
//
// Solidity: function createPool() returns(address)
func (_Market *MarketTransactor) CreatePool(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "createPool")
}

// CreatePool is a paid mutator transaction binding the contract method 0x9a06b113.
//
// Solidity: function createPool() returns(address)
func (_Market *MarketSession) CreatePool() (*types.Transaction, error) {
	return _Market.Contract.CreatePool(&_Market.TransactOpts)
}

// CreatePool is a paid mutator transaction binding the contract method 0x9a06b113.
//
// Solidity: function createPool() returns(address)
func (_Market *MarketTransactorSession) CreatePool() (*types.Transaction, error) {
	return _Market.Contract.CreatePool(&_Market.TransactOpts)
}

// SetData is a paid mutator transaction binding the contract method 0x47064d6a.
//
// Solidity: function setData(string _data) returns()
func (_Market *MarketTransactor) SetData(opts *bind.TransactOpts, _data string) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "setData", _data)
}

// SetData is a paid mutator transaction binding the contract method 0x47064d6a.
//
// Solidity: function setData(string _data) returns()
func (_Market *MarketSession) SetData(_data string) (*types.Transaction, error) {
	return _Market.Contract.SetData(&_Market.TransactOpts, _data)
}

// SetData is a paid mutator transaction binding the contract method 0x47064d6a.
//
// Solidity: function setData(string _data) returns()
func (_Market *MarketTransactorSession) SetData(_data string) (*types.Transaction, error) {
	return _Market.Contract.SetData(&_Market.TransactOpts, _data)
}

// PoolABI is the input ABI used to generate the binding from.
const PoolABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_node\",\"type\":\"address\"}],\"name\":\"isMasterNode\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_node\",\"type\":\"address\"}],\"name\":\"isApprovedNode\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_node\",\"type\":\"address\"}],\"name\":\"addApprovedNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"poolDomain\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getData\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getSeedNode\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_url\",\"type\":\"string\"}],\"name\":\"setPoolDomain\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_data\",\"type\":\"string\"}],\"name\":\"setData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"masterNodes\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cdnDomain\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_node\",\"type\":\"address\"}],\"name\":\"addMasterNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_node\",\"type\":\"string\"}],\"name\":\"setSeedNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_url\",\"type\":\"string\"}],\"name\":\"setCDNDomain\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"data\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"approvedNodes\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getPoolDomain\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_node\",\"type\":\"address\"}],\"name\":\"removeMasterNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"seedNode\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCDNDomain\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_node\",\"type\":\"address\"}],\"name\":\"removeApprovedNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// PoolBin is the compiled bytecode used for deploying new contracts.
const PoolBin = `0x608060405234801561001057600080fd5b50604051602080610f968339810180604052602081101561003057600080fd5b505160028054600160a060020a031916600160a060020a03909216919091179055610f36806100606000396000f3fe608060405234801561001057600080fd5b506004361061016a576000357c0100000000000000000000000000000000000000000000000000000000900480636bc3af27116100e0578063914964771161009957806391496477146105ca578063a6f9dae1146105d2578063dd9cf971146105f8578063e0a18ee71461061e578063e2a5462f14610626578063ed82836e1461062e5761016a565b80636bc3af271461042457806372ae65d5146104ca57806373d4a13a14610570578063806ba4d214610578578063893d20e81461059e5780638da5cb5b146105c25761016a565b80633c714980116101325780633c7149801461027c5780633d2d3a671461028457806347064d6a1461032a57806350ca875f146103d0578063653f2c78146103f65780636b8a98a9146103fe5761016a565b80630458a3351461016f578063053f6ef6146101a9578063324045de146101cf5780633b5e0ebd146101f75780633bc5de3014610274575b600080fd5b6101956004803603602081101561018557600080fd5b5035600160a060020a0316610654565b604080519115158252519081900360200190f35b610195600480360360208110156101bf57600080fd5b5035600160a060020a0316610672565b6101f5600480360360208110156101e557600080fd5b5035600160a060020a0316610690565b005b6101ff610707565b6040805160208082528351818301528351919283929083019185019080838360005b83811015610239578181015183820152602001610221565b50505050905090810190601f1680156102665780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101ff610795565b6101ff61082c565b6101f56004803603602081101561029a57600080fd5b8101906020810181356401000000008111156102b557600080fd5b8201836020820111156102c757600080fd5b803590602001918460018302840111640100000000831117156102e957600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061088d945050505050565b6101f56004803603602081101561034057600080fd5b81019060208101813564010000000081111561035b57600080fd5b82018360208201111561036d57600080fd5b8035906020019184600183028401116401000000008311171561038f57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506108f4945050505050565b610195600480360360208110156103e657600080fd5b5035600160a060020a0316610957565b6101ff61096c565b6101f56004803603602081101561041457600080fd5b5035600160a060020a03166109c7565b6101f56004803603602081101561043a57600080fd5b81019060208101813564010000000081111561045557600080fd5b82018360208201111561046757600080fd5b8035906020019184600183028401116401000000008311171561048957600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610a3b945050505050565b6101f5600480360360208110156104e057600080fd5b8101906020810181356401000000008111156104fb57600080fd5b82018360208201111561050d57600080fd5b8035906020019184600183028401116401000000008311171561052f57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610a9e945050505050565b6101ff610b01565b6101956004803603602081101561058e57600080fd5b5035600160a060020a0316610b5c565b6105a6610b71565b60408051600160a060020a039092168252519081900360200190f35b6105a6610b80565b6101ff610b8f565b6101f5600480360360208110156105e857600080fd5b5035600160a060020a0316610bf0565b6101f56004803603602081101561060e57600080fd5b5035600160a060020a0316610c6f565b6101ff610d04565b6101ff610d5f565b6101f56004803603602081101561064457600080fd5b5035600160a060020a0316610dc0565b600160a060020a031660009081526020819052604090205460ff1690565b600160a060020a031660009081526001602052604090205460ff1690565b600254600160a060020a031633146106e0576040805160e560020a62461bcd0281526020600482015260196024820152600080516020610eeb833981519152604482015290519081900360640190fd5b600160a060020a03166000908152600160208190526040909120805460ff19169091179055565b6005805460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152929183018282801561078d5780601f106107625761010080835404028352916020019161078d565b820191906000526020600020905b81548152906001019060200180831161077057829003601f168201915b505050505081565b60048054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156108215780601f106107f657610100808354040283529160200191610821565b820191906000526020600020905b81548152906001019060200180831161080457829003601f168201915b505050505090505b90565b60038054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156108215780601f106107f657610100808354040283529160200191610821565b600254600160a060020a031633146108dd576040805160e560020a62461bcd0281526020600482015260196024820152600080516020610eeb833981519152604482015290519081900360640190fd5b80516108f0906005906020840190610e52565b5050565b600254600160a060020a03163314610944576040805160e560020a62461bcd0281526020600482015260196024820152600080516020610eeb833981519152604482015290519081900360640190fd5b80516108f0906004906020840190610e52565b60006020819052908152604090205460ff1681565b6006805460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152929183018282801561078d5780601f106107625761010080835404028352916020019161078d565b600254600160a060020a03163314610a17576040805160e560020a62461bcd0281526020600482015260196024820152600080516020610eeb833981519152604482015290519081900360640190fd5b600160a060020a03166000908152602081905260409020805460ff19166001179055565b600254600160a060020a03163314610a8b576040805160e560020a62461bcd0281526020600482015260196024820152600080516020610eeb833981519152604482015290519081900360640190fd5b80516108f0906003906020840190610e52565b600254600160a060020a03163314610aee576040805160e560020a62461bcd0281526020600482015260196024820152600080516020610eeb833981519152604482015290519081900360640190fd5b80516108f0906006906020840190610e52565b6004805460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152929183018282801561078d5780601f106107625761010080835404028352916020019161078d565b60016020526000908152604090205460ff1681565b600254600160a060020a031690565b600254600160a060020a031681565b60058054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156108215780601f106107f657610100808354040283529160200191610821565b600254600160a060020a03163314610c40576040805160e560020a62461bcd0281526020600482015260196024820152600080516020610eeb833981519152604482015290519081900360640190fd5b6002805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600254600160a060020a03163314610cbf576040805160e560020a62461bcd0281526020600482015260196024820152600080516020610eeb833981519152604482015290519081900360640190fd5b600160a060020a03811660009081526020819052604090205460ff1615610d0157600160a060020a0381166000908152602081905260409020805460ff191690555b50565b6003805460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152929183018282801561078d5780601f106107625761010080835404028352916020019161078d565b60068054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156108215780601f106107f657610100808354040283529160200191610821565b600254600160a060020a03163314610e10576040805160e560020a62461bcd0281526020600482015260196024820152600080516020610eeb833981519152604482015290519081900360640190fd5b600160a060020a03811660009081526001602052604090205460ff1615610d0157600160a060020a03166000908152600160205260409020805460ff19169055565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610e9357805160ff1916838001178555610ec0565b82800160010185558215610ec0579182015b82811115610ec0578251825591602001919060010190610ea5565b50610ecc929150610ed0565b5090565b61082991905b80821115610ecc5760008155600101610ed656fe4d757374206265206f776e6572206f6620636f6e747261637400000000000000a165627a7a723058204722d05781582105ead9cdbeb45aafed394ed9e050945c888d793b6f8537f28f0029`

// DeployPool deploys a new Ethereum contract, binding an instance of Pool to it.
func DeployPool(auth *bind.TransactOpts, backend bind.ContractBackend, _owner common.Address) (common.Address, *types.Transaction, *Pool, error) {
	parsed, err := abi.JSON(strings.NewReader(PoolABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PoolBin), backend, _owner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Pool{PoolCaller: PoolCaller{contract: contract}, PoolTransactor: PoolTransactor{contract: contract}, PoolFilterer: PoolFilterer{contract: contract}}, nil
}

// Pool is an auto generated Go binding around an Ethereum contract.
type Pool struct {
	PoolCaller     // Read-only binding to the contract
	PoolTransactor // Write-only binding to the contract
	PoolFilterer   // Log filterer for contract events
}

// PoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type PoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PoolSession struct {
	Contract     *Pool             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PoolCallerSession struct {
	Contract *PoolCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PoolTransactorSession struct {
	Contract     *PoolTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type PoolRaw struct {
	Contract *Pool // Generic contract binding to access the raw methods on
}

// PoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PoolCallerRaw struct {
	Contract *PoolCaller // Generic read-only contract binding to access the raw methods on
}

// PoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PoolTransactorRaw struct {
	Contract *PoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPool creates a new instance of Pool, bound to a specific deployed contract.
func NewPool(address common.Address, backend bind.ContractBackend) (*Pool, error) {
	contract, err := bindPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pool{PoolCaller: PoolCaller{contract: contract}, PoolTransactor: PoolTransactor{contract: contract}, PoolFilterer: PoolFilterer{contract: contract}}, nil
}

// NewPoolCaller creates a new read-only instance of Pool, bound to a specific deployed contract.
func NewPoolCaller(address common.Address, caller bind.ContractCaller) (*PoolCaller, error) {
	contract, err := bindPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PoolCaller{contract: contract}, nil
}

// NewPoolTransactor creates a new write-only instance of Pool, bound to a specific deployed contract.
func NewPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*PoolTransactor, error) {
	contract, err := bindPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PoolTransactor{contract: contract}, nil
}

// NewPoolFilterer creates a new log filterer instance of Pool, bound to a specific deployed contract.
func NewPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*PoolFilterer, error) {
	contract, err := bindPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PoolFilterer{contract: contract}, nil
}

// bindPool binds a generic wrapper to an already deployed contract.
func bindPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pool *PoolRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Pool.Contract.PoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pool *PoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pool.Contract.PoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pool *PoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pool.Contract.PoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pool *PoolCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Pool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pool *PoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pool *PoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pool.Contract.contract.Transact(opts, method, params...)
}

// ApprovedNodes is a free data retrieval call binding the contract method 0x806ba4d2.
//
// Solidity: function approvedNodes(address ) constant returns(bool)
func (_Pool *PoolCaller) ApprovedNodes(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Pool.contract.Call(opts, out, "approvedNodes", arg0)
	return *ret0, err
}

// ApprovedNodes is a free data retrieval call binding the contract method 0x806ba4d2.
//
// Solidity: function approvedNodes(address ) constant returns(bool)
func (_Pool *PoolSession) ApprovedNodes(arg0 common.Address) (bool, error) {
	return _Pool.Contract.ApprovedNodes(&_Pool.CallOpts, arg0)
}

// ApprovedNodes is a free data retrieval call binding the contract method 0x806ba4d2.
//
// Solidity: function approvedNodes(address ) constant returns(bool)
func (_Pool *PoolCallerSession) ApprovedNodes(arg0 common.Address) (bool, error) {
	return _Pool.Contract.ApprovedNodes(&_Pool.CallOpts, arg0)
}

// CdnDomain is a free data retrieval call binding the contract method 0x653f2c78.
//
// Solidity: function cdnDomain() constant returns(string)
func (_Pool *PoolCaller) CdnDomain(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Pool.contract.Call(opts, out, "cdnDomain")
	return *ret0, err
}

// CdnDomain is a free data retrieval call binding the contract method 0x653f2c78.
//
// Solidity: function cdnDomain() constant returns(string)
func (_Pool *PoolSession) CdnDomain() (string, error) {
	return _Pool.Contract.CdnDomain(&_Pool.CallOpts)
}

// CdnDomain is a free data retrieval call binding the contract method 0x653f2c78.
//
// Solidity: function cdnDomain() constant returns(string)
func (_Pool *PoolCallerSession) CdnDomain() (string, error) {
	return _Pool.Contract.CdnDomain(&_Pool.CallOpts)
}

// Data is a free data retrieval call binding the contract method 0x73d4a13a.
//
// Solidity: function data() constant returns(string)
func (_Pool *PoolCaller) Data(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Pool.contract.Call(opts, out, "data")
	return *ret0, err
}

// Data is a free data retrieval call binding the contract method 0x73d4a13a.
//
// Solidity: function data() constant returns(string)
func (_Pool *PoolSession) Data() (string, error) {
	return _Pool.Contract.Data(&_Pool.CallOpts)
}

// Data is a free data retrieval call binding the contract method 0x73d4a13a.
//
// Solidity: function data() constant returns(string)
func (_Pool *PoolCallerSession) Data() (string, error) {
	return _Pool.Contract.Data(&_Pool.CallOpts)
}

// GetCDNDomain is a free data retrieval call binding the contract method 0xe2a5462f.
//
// Solidity: function getCDNDomain() constant returns(string)
func (_Pool *PoolCaller) GetCDNDomain(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Pool.contract.Call(opts, out, "getCDNDomain")
	return *ret0, err
}

// GetCDNDomain is a free data retrieval call binding the contract method 0xe2a5462f.
//
// Solidity: function getCDNDomain() constant returns(string)
func (_Pool *PoolSession) GetCDNDomain() (string, error) {
	return _Pool.Contract.GetCDNDomain(&_Pool.CallOpts)
}

// GetCDNDomain is a free data retrieval call binding the contract method 0xe2a5462f.
//
// Solidity: function getCDNDomain() constant returns(string)
func (_Pool *PoolCallerSession) GetCDNDomain() (string, error) {
	return _Pool.Contract.GetCDNDomain(&_Pool.CallOpts)
}

// GetData is a free data retrieval call binding the contract method 0x3bc5de30.
//
// Solidity: function getData() constant returns(string)
func (_Pool *PoolCaller) GetData(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Pool.contract.Call(opts, out, "getData")
	return *ret0, err
}

// GetData is a free data retrieval call binding the contract method 0x3bc5de30.
//
// Solidity: function getData() constant returns(string)
func (_Pool *PoolSession) GetData() (string, error) {
	return _Pool.Contract.GetData(&_Pool.CallOpts)
}

// GetData is a free data retrieval call binding the contract method 0x3bc5de30.
//
// Solidity: function getData() constant returns(string)
func (_Pool *PoolCallerSession) GetData() (string, error) {
	return _Pool.Contract.GetData(&_Pool.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_Pool *PoolCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Pool.contract.Call(opts, out, "getOwner")
	return *ret0, err
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_Pool *PoolSession) GetOwner() (common.Address, error) {
	return _Pool.Contract.GetOwner(&_Pool.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_Pool *PoolCallerSession) GetOwner() (common.Address, error) {
	return _Pool.Contract.GetOwner(&_Pool.CallOpts)
}

// GetPoolDomain is a free data retrieval call binding the contract method 0x91496477.
//
// Solidity: function getPoolDomain() constant returns(string)
func (_Pool *PoolCaller) GetPoolDomain(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Pool.contract.Call(opts, out, "getPoolDomain")
	return *ret0, err
}

// GetPoolDomain is a free data retrieval call binding the contract method 0x91496477.
//
// Solidity: function getPoolDomain() constant returns(string)
func (_Pool *PoolSession) GetPoolDomain() (string, error) {
	return _Pool.Contract.GetPoolDomain(&_Pool.CallOpts)
}

// GetPoolDomain is a free data retrieval call binding the contract method 0x91496477.
//
// Solidity: function getPoolDomain() constant returns(string)
func (_Pool *PoolCallerSession) GetPoolDomain() (string, error) {
	return _Pool.Contract.GetPoolDomain(&_Pool.CallOpts)
}

// GetSeedNode is a free data retrieval call binding the contract method 0x3c714980.
//
// Solidity: function getSeedNode() constant returns(string)
func (_Pool *PoolCaller) GetSeedNode(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Pool.contract.Call(opts, out, "getSeedNode")
	return *ret0, err
}

// GetSeedNode is a free data retrieval call binding the contract method 0x3c714980.
//
// Solidity: function getSeedNode() constant returns(string)
func (_Pool *PoolSession) GetSeedNode() (string, error) {
	return _Pool.Contract.GetSeedNode(&_Pool.CallOpts)
}

// GetSeedNode is a free data retrieval call binding the contract method 0x3c714980.
//
// Solidity: function getSeedNode() constant returns(string)
func (_Pool *PoolCallerSession) GetSeedNode() (string, error) {
	return _Pool.Contract.GetSeedNode(&_Pool.CallOpts)
}

// IsApprovedNode is a free data retrieval call binding the contract method 0x053f6ef6.
//
// Solidity: function isApprovedNode(address _node) constant returns(bool)
func (_Pool *PoolCaller) IsApprovedNode(opts *bind.CallOpts, _node common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Pool.contract.Call(opts, out, "isApprovedNode", _node)
	return *ret0, err
}

// IsApprovedNode is a free data retrieval call binding the contract method 0x053f6ef6.
//
// Solidity: function isApprovedNode(address _node) constant returns(bool)
func (_Pool *PoolSession) IsApprovedNode(_node common.Address) (bool, error) {
	return _Pool.Contract.IsApprovedNode(&_Pool.CallOpts, _node)
}

// IsApprovedNode is a free data retrieval call binding the contract method 0x053f6ef6.
//
// Solidity: function isApprovedNode(address _node) constant returns(bool)
func (_Pool *PoolCallerSession) IsApprovedNode(_node common.Address) (bool, error) {
	return _Pool.Contract.IsApprovedNode(&_Pool.CallOpts, _node)
}

// IsMasterNode is a free data retrieval call binding the contract method 0x0458a335.
//
// Solidity: function isMasterNode(address _node) constant returns(bool)
func (_Pool *PoolCaller) IsMasterNode(opts *bind.CallOpts, _node common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Pool.contract.Call(opts, out, "isMasterNode", _node)
	return *ret0, err
}

// IsMasterNode is a free data retrieval call binding the contract method 0x0458a335.
//
// Solidity: function isMasterNode(address _node) constant returns(bool)
func (_Pool *PoolSession) IsMasterNode(_node common.Address) (bool, error) {
	return _Pool.Contract.IsMasterNode(&_Pool.CallOpts, _node)
}

// IsMasterNode is a free data retrieval call binding the contract method 0x0458a335.
//
// Solidity: function isMasterNode(address _node) constant returns(bool)
func (_Pool *PoolCallerSession) IsMasterNode(_node common.Address) (bool, error) {
	return _Pool.Contract.IsMasterNode(&_Pool.CallOpts, _node)
}

// MasterNodes is a free data retrieval call binding the contract method 0x50ca875f.
//
// Solidity: function masterNodes(address ) constant returns(bool)
func (_Pool *PoolCaller) MasterNodes(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Pool.contract.Call(opts, out, "masterNodes", arg0)
	return *ret0, err
}

// MasterNodes is a free data retrieval call binding the contract method 0x50ca875f.
//
// Solidity: function masterNodes(address ) constant returns(bool)
func (_Pool *PoolSession) MasterNodes(arg0 common.Address) (bool, error) {
	return _Pool.Contract.MasterNodes(&_Pool.CallOpts, arg0)
}

// MasterNodes is a free data retrieval call binding the contract method 0x50ca875f.
//
// Solidity: function masterNodes(address ) constant returns(bool)
func (_Pool *PoolCallerSession) MasterNodes(arg0 common.Address) (bool, error) {
	return _Pool.Contract.MasterNodes(&_Pool.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Pool *PoolCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Pool.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Pool *PoolSession) Owner() (common.Address, error) {
	return _Pool.Contract.Owner(&_Pool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Pool *PoolCallerSession) Owner() (common.Address, error) {
	return _Pool.Contract.Owner(&_Pool.CallOpts)
}

// PoolDomain is a free data retrieval call binding the contract method 0x3b5e0ebd.
//
// Solidity: function poolDomain() constant returns(string)
func (_Pool *PoolCaller) PoolDomain(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Pool.contract.Call(opts, out, "poolDomain")
	return *ret0, err
}

// PoolDomain is a free data retrieval call binding the contract method 0x3b5e0ebd.
//
// Solidity: function poolDomain() constant returns(string)
func (_Pool *PoolSession) PoolDomain() (string, error) {
	return _Pool.Contract.PoolDomain(&_Pool.CallOpts)
}

// PoolDomain is a free data retrieval call binding the contract method 0x3b5e0ebd.
//
// Solidity: function poolDomain() constant returns(string)
func (_Pool *PoolCallerSession) PoolDomain() (string, error) {
	return _Pool.Contract.PoolDomain(&_Pool.CallOpts)
}

// SeedNode is a free data retrieval call binding the contract method 0xe0a18ee7.
//
// Solidity: function seedNode() constant returns(string)
func (_Pool *PoolCaller) SeedNode(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Pool.contract.Call(opts, out, "seedNode")
	return *ret0, err
}

// SeedNode is a free data retrieval call binding the contract method 0xe0a18ee7.
//
// Solidity: function seedNode() constant returns(string)
func (_Pool *PoolSession) SeedNode() (string, error) {
	return _Pool.Contract.SeedNode(&_Pool.CallOpts)
}

// SeedNode is a free data retrieval call binding the contract method 0xe0a18ee7.
//
// Solidity: function seedNode() constant returns(string)
func (_Pool *PoolCallerSession) SeedNode() (string, error) {
	return _Pool.Contract.SeedNode(&_Pool.CallOpts)
}

// AddApprovedNode is a paid mutator transaction binding the contract method 0x324045de.
//
// Solidity: function addApprovedNode(address _node) returns()
func (_Pool *PoolTransactor) AddApprovedNode(opts *bind.TransactOpts, _node common.Address) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "addApprovedNode", _node)
}

// AddApprovedNode is a paid mutator transaction binding the contract method 0x324045de.
//
// Solidity: function addApprovedNode(address _node) returns()
func (_Pool *PoolSession) AddApprovedNode(_node common.Address) (*types.Transaction, error) {
	return _Pool.Contract.AddApprovedNode(&_Pool.TransactOpts, _node)
}

// AddApprovedNode is a paid mutator transaction binding the contract method 0x324045de.
//
// Solidity: function addApprovedNode(address _node) returns()
func (_Pool *PoolTransactorSession) AddApprovedNode(_node common.Address) (*types.Transaction, error) {
	return _Pool.Contract.AddApprovedNode(&_Pool.TransactOpts, _node)
}

// AddMasterNode is a paid mutator transaction binding the contract method 0x6b8a98a9.
//
// Solidity: function addMasterNode(address _node) returns()
func (_Pool *PoolTransactor) AddMasterNode(opts *bind.TransactOpts, _node common.Address) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "addMasterNode", _node)
}

// AddMasterNode is a paid mutator transaction binding the contract method 0x6b8a98a9.
//
// Solidity: function addMasterNode(address _node) returns()
func (_Pool *PoolSession) AddMasterNode(_node common.Address) (*types.Transaction, error) {
	return _Pool.Contract.AddMasterNode(&_Pool.TransactOpts, _node)
}

// AddMasterNode is a paid mutator transaction binding the contract method 0x6b8a98a9.
//
// Solidity: function addMasterNode(address _node) returns()
func (_Pool *PoolTransactorSession) AddMasterNode(_node common.Address) (*types.Transaction, error) {
	return _Pool.Contract.AddMasterNode(&_Pool.TransactOpts, _node)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address _owner) returns()
func (_Pool *PoolTransactor) ChangeOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "changeOwner", _owner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address _owner) returns()
func (_Pool *PoolSession) ChangeOwner(_owner common.Address) (*types.Transaction, error) {
	return _Pool.Contract.ChangeOwner(&_Pool.TransactOpts, _owner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address _owner) returns()
func (_Pool *PoolTransactorSession) ChangeOwner(_owner common.Address) (*types.Transaction, error) {
	return _Pool.Contract.ChangeOwner(&_Pool.TransactOpts, _owner)
}

// RemoveApprovedNode is a paid mutator transaction binding the contract method 0xed82836e.
//
// Solidity: function removeApprovedNode(address _node) returns()
func (_Pool *PoolTransactor) RemoveApprovedNode(opts *bind.TransactOpts, _node common.Address) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "removeApprovedNode", _node)
}

// RemoveApprovedNode is a paid mutator transaction binding the contract method 0xed82836e.
//
// Solidity: function removeApprovedNode(address _node) returns()
func (_Pool *PoolSession) RemoveApprovedNode(_node common.Address) (*types.Transaction, error) {
	return _Pool.Contract.RemoveApprovedNode(&_Pool.TransactOpts, _node)
}

// RemoveApprovedNode is a paid mutator transaction binding the contract method 0xed82836e.
//
// Solidity: function removeApprovedNode(address _node) returns()
func (_Pool *PoolTransactorSession) RemoveApprovedNode(_node common.Address) (*types.Transaction, error) {
	return _Pool.Contract.RemoveApprovedNode(&_Pool.TransactOpts, _node)
}

// RemoveMasterNode is a paid mutator transaction binding the contract method 0xdd9cf971.
//
// Solidity: function removeMasterNode(address _node) returns()
func (_Pool *PoolTransactor) RemoveMasterNode(opts *bind.TransactOpts, _node common.Address) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "removeMasterNode", _node)
}

// RemoveMasterNode is a paid mutator transaction binding the contract method 0xdd9cf971.
//
// Solidity: function removeMasterNode(address _node) returns()
func (_Pool *PoolSession) RemoveMasterNode(_node common.Address) (*types.Transaction, error) {
	return _Pool.Contract.RemoveMasterNode(&_Pool.TransactOpts, _node)
}

// RemoveMasterNode is a paid mutator transaction binding the contract method 0xdd9cf971.
//
// Solidity: function removeMasterNode(address _node) returns()
func (_Pool *PoolTransactorSession) RemoveMasterNode(_node common.Address) (*types.Transaction, error) {
	return _Pool.Contract.RemoveMasterNode(&_Pool.TransactOpts, _node)
}

// SetCDNDomain is a paid mutator transaction binding the contract method 0x72ae65d5.
//
// Solidity: function setCDNDomain(string _url) returns()
func (_Pool *PoolTransactor) SetCDNDomain(opts *bind.TransactOpts, _url string) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "setCDNDomain", _url)
}

// SetCDNDomain is a paid mutator transaction binding the contract method 0x72ae65d5.
//
// Solidity: function setCDNDomain(string _url) returns()
func (_Pool *PoolSession) SetCDNDomain(_url string) (*types.Transaction, error) {
	return _Pool.Contract.SetCDNDomain(&_Pool.TransactOpts, _url)
}

// SetCDNDomain is a paid mutator transaction binding the contract method 0x72ae65d5.
//
// Solidity: function setCDNDomain(string _url) returns()
func (_Pool *PoolTransactorSession) SetCDNDomain(_url string) (*types.Transaction, error) {
	return _Pool.Contract.SetCDNDomain(&_Pool.TransactOpts, _url)
}

// SetData is a paid mutator transaction binding the contract method 0x47064d6a.
//
// Solidity: function setData(string _data) returns()
func (_Pool *PoolTransactor) SetData(opts *bind.TransactOpts, _data string) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "setData", _data)
}

// SetData is a paid mutator transaction binding the contract method 0x47064d6a.
//
// Solidity: function setData(string _data) returns()
func (_Pool *PoolSession) SetData(_data string) (*types.Transaction, error) {
	return _Pool.Contract.SetData(&_Pool.TransactOpts, _data)
}

// SetData is a paid mutator transaction binding the contract method 0x47064d6a.
//
// Solidity: function setData(string _data) returns()
func (_Pool *PoolTransactorSession) SetData(_data string) (*types.Transaction, error) {
	return _Pool.Contract.SetData(&_Pool.TransactOpts, _data)
}

// SetPoolDomain is a paid mutator transaction binding the contract method 0x3d2d3a67.
//
// Solidity: function setPoolDomain(string _url) returns()
func (_Pool *PoolTransactor) SetPoolDomain(opts *bind.TransactOpts, _url string) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "setPoolDomain", _url)
}

// SetPoolDomain is a paid mutator transaction binding the contract method 0x3d2d3a67.
//
// Solidity: function setPoolDomain(string _url) returns()
func (_Pool *PoolSession) SetPoolDomain(_url string) (*types.Transaction, error) {
	return _Pool.Contract.SetPoolDomain(&_Pool.TransactOpts, _url)
}

// SetPoolDomain is a paid mutator transaction binding the contract method 0x3d2d3a67.
//
// Solidity: function setPoolDomain(string _url) returns()
func (_Pool *PoolTransactorSession) SetPoolDomain(_url string) (*types.Transaction, error) {
	return _Pool.Contract.SetPoolDomain(&_Pool.TransactOpts, _url)
}

// SetSeedNode is a paid mutator transaction binding the contract method 0x6bc3af27.
//
// Solidity: function setSeedNode(string _node) returns()
func (_Pool *PoolTransactor) SetSeedNode(opts *bind.TransactOpts, _node string) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "setSeedNode", _node)
}

// SetSeedNode is a paid mutator transaction binding the contract method 0x6bc3af27.
//
// Solidity: function setSeedNode(string _node) returns()
func (_Pool *PoolSession) SetSeedNode(_node string) (*types.Transaction, error) {
	return _Pool.Contract.SetSeedNode(&_Pool.TransactOpts, _node)
}

// SetSeedNode is a paid mutator transaction binding the contract method 0x6bc3af27.
//
// Solidity: function setSeedNode(string _node) returns()
func (_Pool *PoolTransactorSession) SetSeedNode(_node string) (*types.Transaction, error) {
	return _Pool.Contract.SetSeedNode(&_Pool.TransactOpts, _node)
}

// PoolFactoryABI is the input ABI used to generate the binding from.
const PoolFactoryABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ownerToPools\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allPools\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"getOwnedPool\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"name\":\"getPoolOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"createPool\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"poolToOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAllPools\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// PoolFactoryBin is the compiled bytecode used for deploying new contracts.
const PoolFactoryBin = `0x608060405234801561001057600080fd5b5061140f806100206000396000f3fe608060405234801561001057600080fd5b506004361061009a576000357c0100000000000000000000000000000000000000000000000000000000900480637cf25517116100785780637cf255171461017a5780639049f9d2146101a0578063d4cced9a146101c6578063d88ff1f4146101ec5761009a565b806312708d721461009f57806341d1de97146100e7578063556f77db14610104575b600080fd5b6100cb600480360360408110156100b557600080fd5b50600160a060020a0381351690602001356101f4565b60408051600160a060020a039092168252519081900360200190f35b6100cb600480360360208110156100fd57600080fd5b503561022b565b61012a6004803603602081101561011a57600080fd5b5035600160a060020a0316610253565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561016657818101518382015260200161014e565b505050509050019250505060405180910390f35b6100cb6004803603602081101561019057600080fd5b5035600160a060020a03166102c7565b6100cb600480360360208110156101b657600080fd5b5035600160a060020a03166102e5565b6100cb600480360360208110156101dc57600080fd5b5035600160a060020a03166103c0565b61012a6103db565b60006020528160005260406000208181548110151561020f57fe5b600091825260209091200154600160a060020a03169150829050565b600280548290811061023957fe5b600091825260209091200154600160a060020a0316905081565b600160a060020a038116600090815260208181526040918290208054835181840281018401909452808452606093928301828280156102bb57602002820191906000526020600020905b8154600160a060020a0316815260019091019060200180831161029d575b50505050509050919050565b600160a060020a039081166000908152600160205260409020541690565b600080826102f161043d565b600160a060020a03909116815260405190819003602001906000f08015801561031e573d6000803e3d6000fd5b50600160a060020a0380851660008181526020818152604080832080546001818101835591855283852001805496881673ffffffffffffffffffffffffffffffffffffffff19978816811790915580855292819052908320805486169094179093556002805493840181559091527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace9091018054909216179055915050919050565b600160205260009081526040902054600160a060020a031681565b6060600280548060200260200160405190810160405280929190818152602001828054801561043357602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311610415575b5050505050905090565b604051610f968061044e8339019056fe608060405234801561001057600080fd5b50604051602080610f968339810180604052602081101561003057600080fd5b505160028054600160a060020a031916600160a060020a03909216919091179055610f36806100606000396000f3fe608060405234801561001057600080fd5b506004361061016a576000357c0100000000000000000000000000000000000000000000000000000000900480636bc3af27116100e0578063914964771161009957806391496477146105ca578063a6f9dae1146105d2578063dd9cf971146105f8578063e0a18ee71461061e578063e2a5462f14610626578063ed82836e1461062e5761016a565b80636bc3af271461042457806372ae65d5146104ca57806373d4a13a14610570578063806ba4d214610578578063893d20e81461059e5780638da5cb5b146105c25761016a565b80633c714980116101325780633c7149801461027c5780633d2d3a671461028457806347064d6a1461032a57806350ca875f146103d0578063653f2c78146103f65780636b8a98a9146103fe5761016a565b80630458a3351461016f578063053f6ef6146101a9578063324045de146101cf5780633b5e0ebd146101f75780633bc5de3014610274575b600080fd5b6101956004803603602081101561018557600080fd5b5035600160a060020a0316610654565b604080519115158252519081900360200190f35b610195600480360360208110156101bf57600080fd5b5035600160a060020a0316610672565b6101f5600480360360208110156101e557600080fd5b5035600160a060020a0316610690565b005b6101ff610707565b6040805160208082528351818301528351919283929083019185019080838360005b83811015610239578181015183820152602001610221565b50505050905090810190601f1680156102665780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101ff610795565b6101ff61082c565b6101f56004803603602081101561029a57600080fd5b8101906020810181356401000000008111156102b557600080fd5b8201836020820111156102c757600080fd5b803590602001918460018302840111640100000000831117156102e957600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061088d945050505050565b6101f56004803603602081101561034057600080fd5b81019060208101813564010000000081111561035b57600080fd5b82018360208201111561036d57600080fd5b8035906020019184600183028401116401000000008311171561038f57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506108f4945050505050565b610195600480360360208110156103e657600080fd5b5035600160a060020a0316610957565b6101ff61096c565b6101f56004803603602081101561041457600080fd5b5035600160a060020a03166109c7565b6101f56004803603602081101561043a57600080fd5b81019060208101813564010000000081111561045557600080fd5b82018360208201111561046757600080fd5b8035906020019184600183028401116401000000008311171561048957600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610a3b945050505050565b6101f5600480360360208110156104e057600080fd5b8101906020810181356401000000008111156104fb57600080fd5b82018360208201111561050d57600080fd5b8035906020019184600183028401116401000000008311171561052f57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610a9e945050505050565b6101ff610b01565b6101956004803603602081101561058e57600080fd5b5035600160a060020a0316610b5c565b6105a6610b71565b60408051600160a060020a039092168252519081900360200190f35b6105a6610b80565b6101ff610b8f565b6101f5600480360360208110156105e857600080fd5b5035600160a060020a0316610bf0565b6101f56004803603602081101561060e57600080fd5b5035600160a060020a0316610c6f565b6101ff610d04565b6101ff610d5f565b6101f56004803603602081101561064457600080fd5b5035600160a060020a0316610dc0565b600160a060020a031660009081526020819052604090205460ff1690565b600160a060020a031660009081526001602052604090205460ff1690565b600254600160a060020a031633146106e0576040805160e560020a62461bcd0281526020600482015260196024820152600080516020610eeb833981519152604482015290519081900360640190fd5b600160a060020a03166000908152600160208190526040909120805460ff19169091179055565b6005805460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152929183018282801561078d5780601f106107625761010080835404028352916020019161078d565b820191906000526020600020905b81548152906001019060200180831161077057829003601f168201915b505050505081565b60048054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156108215780601f106107f657610100808354040283529160200191610821565b820191906000526020600020905b81548152906001019060200180831161080457829003601f168201915b505050505090505b90565b60038054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156108215780601f106107f657610100808354040283529160200191610821565b600254600160a060020a031633146108dd576040805160e560020a62461bcd0281526020600482015260196024820152600080516020610eeb833981519152604482015290519081900360640190fd5b80516108f0906005906020840190610e52565b5050565b600254600160a060020a03163314610944576040805160e560020a62461bcd0281526020600482015260196024820152600080516020610eeb833981519152604482015290519081900360640190fd5b80516108f0906004906020840190610e52565b60006020819052908152604090205460ff1681565b6006805460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152929183018282801561078d5780601f106107625761010080835404028352916020019161078d565b600254600160a060020a03163314610a17576040805160e560020a62461bcd0281526020600482015260196024820152600080516020610eeb833981519152604482015290519081900360640190fd5b600160a060020a03166000908152602081905260409020805460ff19166001179055565b600254600160a060020a03163314610a8b576040805160e560020a62461bcd0281526020600482015260196024820152600080516020610eeb833981519152604482015290519081900360640190fd5b80516108f0906003906020840190610e52565b600254600160a060020a03163314610aee576040805160e560020a62461bcd0281526020600482015260196024820152600080516020610eeb833981519152604482015290519081900360640190fd5b80516108f0906006906020840190610e52565b6004805460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152929183018282801561078d5780601f106107625761010080835404028352916020019161078d565b60016020526000908152604090205460ff1681565b600254600160a060020a031690565b600254600160a060020a031681565b60058054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156108215780601f106107f657610100808354040283529160200191610821565b600254600160a060020a03163314610c40576040805160e560020a62461bcd0281526020600482015260196024820152600080516020610eeb833981519152604482015290519081900360640190fd5b6002805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600254600160a060020a03163314610cbf576040805160e560020a62461bcd0281526020600482015260196024820152600080516020610eeb833981519152604482015290519081900360640190fd5b600160a060020a03811660009081526020819052604090205460ff1615610d0157600160a060020a0381166000908152602081905260409020805460ff191690555b50565b6003805460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152929183018282801561078d5780601f106107625761010080835404028352916020019161078d565b60068054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156108215780601f106107f657610100808354040283529160200191610821565b600254600160a060020a03163314610e10576040805160e560020a62461bcd0281526020600482015260196024820152600080516020610eeb833981519152604482015290519081900360640190fd5b600160a060020a03811660009081526001602052604090205460ff1615610d0157600160a060020a03166000908152600160205260409020805460ff19169055565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610e9357805160ff1916838001178555610ec0565b82800160010185558215610ec0579182015b82811115610ec0578251825591602001919060010190610ea5565b50610ecc929150610ed0565b5090565b61082991905b80821115610ecc5760008155600101610ed656fe4d757374206265206f776e6572206f6620636f6e747261637400000000000000a165627a7a723058204722d05781582105ead9cdbeb45aafed394ed9e050945c888d793b6f8537f28f0029a165627a7a72305820a422d5e5ff44ed496b4553e569b1d779511c7cd7563f7b01f98f5c1ee7a6803c0029`

// DeployPoolFactory deploys a new Ethereum contract, binding an instance of PoolFactory to it.
func DeployPoolFactory(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PoolFactory, error) {
	parsed, err := abi.JSON(strings.NewReader(PoolFactoryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PoolFactoryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PoolFactory{PoolFactoryCaller: PoolFactoryCaller{contract: contract}, PoolFactoryTransactor: PoolFactoryTransactor{contract: contract}, PoolFactoryFilterer: PoolFactoryFilterer{contract: contract}}, nil
}

// PoolFactory is an auto generated Go binding around an Ethereum contract.
type PoolFactory struct {
	PoolFactoryCaller     // Read-only binding to the contract
	PoolFactoryTransactor // Write-only binding to the contract
	PoolFactoryFilterer   // Log filterer for contract events
}

// PoolFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type PoolFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PoolFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PoolFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PoolFactorySession struct {
	Contract     *PoolFactory      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PoolFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PoolFactoryCallerSession struct {
	Contract *PoolFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// PoolFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PoolFactoryTransactorSession struct {
	Contract     *PoolFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// PoolFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type PoolFactoryRaw struct {
	Contract *PoolFactory // Generic contract binding to access the raw methods on
}

// PoolFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PoolFactoryCallerRaw struct {
	Contract *PoolFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// PoolFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PoolFactoryTransactorRaw struct {
	Contract *PoolFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPoolFactory creates a new instance of PoolFactory, bound to a specific deployed contract.
func NewPoolFactory(address common.Address, backend bind.ContractBackend) (*PoolFactory, error) {
	contract, err := bindPoolFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PoolFactory{PoolFactoryCaller: PoolFactoryCaller{contract: contract}, PoolFactoryTransactor: PoolFactoryTransactor{contract: contract}, PoolFactoryFilterer: PoolFactoryFilterer{contract: contract}}, nil
}

// NewPoolFactoryCaller creates a new read-only instance of PoolFactory, bound to a specific deployed contract.
func NewPoolFactoryCaller(address common.Address, caller bind.ContractCaller) (*PoolFactoryCaller, error) {
	contract, err := bindPoolFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PoolFactoryCaller{contract: contract}, nil
}

// NewPoolFactoryTransactor creates a new write-only instance of PoolFactory, bound to a specific deployed contract.
func NewPoolFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*PoolFactoryTransactor, error) {
	contract, err := bindPoolFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PoolFactoryTransactor{contract: contract}, nil
}

// NewPoolFactoryFilterer creates a new log filterer instance of PoolFactory, bound to a specific deployed contract.
func NewPoolFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*PoolFactoryFilterer, error) {
	contract, err := bindPoolFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PoolFactoryFilterer{contract: contract}, nil
}

// bindPoolFactory binds a generic wrapper to an already deployed contract.
func bindPoolFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PoolFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PoolFactory *PoolFactoryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PoolFactory.Contract.PoolFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PoolFactory *PoolFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoolFactory.Contract.PoolFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PoolFactory *PoolFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PoolFactory.Contract.PoolFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PoolFactory *PoolFactoryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PoolFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PoolFactory *PoolFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoolFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PoolFactory *PoolFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PoolFactory.Contract.contract.Transact(opts, method, params...)
}

// AllPools is a free data retrieval call binding the contract method 0x41d1de97.
//
// Solidity: function allPools(uint256 ) constant returns(address)
func (_PoolFactory *PoolFactoryCaller) AllPools(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PoolFactory.contract.Call(opts, out, "allPools", arg0)
	return *ret0, err
}

// AllPools is a free data retrieval call binding the contract method 0x41d1de97.
//
// Solidity: function allPools(uint256 ) constant returns(address)
func (_PoolFactory *PoolFactorySession) AllPools(arg0 *big.Int) (common.Address, error) {
	return _PoolFactory.Contract.AllPools(&_PoolFactory.CallOpts, arg0)
}

// AllPools is a free data retrieval call binding the contract method 0x41d1de97.
//
// Solidity: function allPools(uint256 ) constant returns(address)
func (_PoolFactory *PoolFactoryCallerSession) AllPools(arg0 *big.Int) (common.Address, error) {
	return _PoolFactory.Contract.AllPools(&_PoolFactory.CallOpts, arg0)
}

// GetAllPools is a free data retrieval call binding the contract method 0xd88ff1f4.
//
// Solidity: function getAllPools() constant returns(address[])
func (_PoolFactory *PoolFactoryCaller) GetAllPools(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _PoolFactory.contract.Call(opts, out, "getAllPools")
	return *ret0, err
}

// GetAllPools is a free data retrieval call binding the contract method 0xd88ff1f4.
//
// Solidity: function getAllPools() constant returns(address[])
func (_PoolFactory *PoolFactorySession) GetAllPools() ([]common.Address, error) {
	return _PoolFactory.Contract.GetAllPools(&_PoolFactory.CallOpts)
}

// GetAllPools is a free data retrieval call binding the contract method 0xd88ff1f4.
//
// Solidity: function getAllPools() constant returns(address[])
func (_PoolFactory *PoolFactoryCallerSession) GetAllPools() ([]common.Address, error) {
	return _PoolFactory.Contract.GetAllPools(&_PoolFactory.CallOpts)
}

// GetOwnedPool is a free data retrieval call binding the contract method 0x556f77db.
//
// Solidity: function getOwnedPool(address _owner) constant returns(address[])
func (_PoolFactory *PoolFactoryCaller) GetOwnedPool(opts *bind.CallOpts, _owner common.Address) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _PoolFactory.contract.Call(opts, out, "getOwnedPool", _owner)
	return *ret0, err
}

// GetOwnedPool is a free data retrieval call binding the contract method 0x556f77db.
//
// Solidity: function getOwnedPool(address _owner) constant returns(address[])
func (_PoolFactory *PoolFactorySession) GetOwnedPool(_owner common.Address) ([]common.Address, error) {
	return _PoolFactory.Contract.GetOwnedPool(&_PoolFactory.CallOpts, _owner)
}

// GetOwnedPool is a free data retrieval call binding the contract method 0x556f77db.
//
// Solidity: function getOwnedPool(address _owner) constant returns(address[])
func (_PoolFactory *PoolFactoryCallerSession) GetOwnedPool(_owner common.Address) ([]common.Address, error) {
	return _PoolFactory.Contract.GetOwnedPool(&_PoolFactory.CallOpts, _owner)
}

// GetPoolOwner is a free data retrieval call binding the contract method 0x7cf25517.
//
// Solidity: function getPoolOwner(address _pool) constant returns(address)
func (_PoolFactory *PoolFactoryCaller) GetPoolOwner(opts *bind.CallOpts, _pool common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PoolFactory.contract.Call(opts, out, "getPoolOwner", _pool)
	return *ret0, err
}

// GetPoolOwner is a free data retrieval call binding the contract method 0x7cf25517.
//
// Solidity: function getPoolOwner(address _pool) constant returns(address)
func (_PoolFactory *PoolFactorySession) GetPoolOwner(_pool common.Address) (common.Address, error) {
	return _PoolFactory.Contract.GetPoolOwner(&_PoolFactory.CallOpts, _pool)
}

// GetPoolOwner is a free data retrieval call binding the contract method 0x7cf25517.
//
// Solidity: function getPoolOwner(address _pool) constant returns(address)
func (_PoolFactory *PoolFactoryCallerSession) GetPoolOwner(_pool common.Address) (common.Address, error) {
	return _PoolFactory.Contract.GetPoolOwner(&_PoolFactory.CallOpts, _pool)
}

// OwnerToPools is a free data retrieval call binding the contract method 0x12708d72.
//
// Solidity: function ownerToPools(address , uint256 ) constant returns(address)
func (_PoolFactory *PoolFactoryCaller) OwnerToPools(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PoolFactory.contract.Call(opts, out, "ownerToPools", arg0, arg1)
	return *ret0, err
}

// OwnerToPools is a free data retrieval call binding the contract method 0x12708d72.
//
// Solidity: function ownerToPools(address , uint256 ) constant returns(address)
func (_PoolFactory *PoolFactorySession) OwnerToPools(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _PoolFactory.Contract.OwnerToPools(&_PoolFactory.CallOpts, arg0, arg1)
}

// OwnerToPools is a free data retrieval call binding the contract method 0x12708d72.
//
// Solidity: function ownerToPools(address , uint256 ) constant returns(address)
func (_PoolFactory *PoolFactoryCallerSession) OwnerToPools(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _PoolFactory.Contract.OwnerToPools(&_PoolFactory.CallOpts, arg0, arg1)
}

// PoolToOwner is a free data retrieval call binding the contract method 0xd4cced9a.
//
// Solidity: function poolToOwner(address ) constant returns(address)
func (_PoolFactory *PoolFactoryCaller) PoolToOwner(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PoolFactory.contract.Call(opts, out, "poolToOwner", arg0)
	return *ret0, err
}

// PoolToOwner is a free data retrieval call binding the contract method 0xd4cced9a.
//
// Solidity: function poolToOwner(address ) constant returns(address)
func (_PoolFactory *PoolFactorySession) PoolToOwner(arg0 common.Address) (common.Address, error) {
	return _PoolFactory.Contract.PoolToOwner(&_PoolFactory.CallOpts, arg0)
}

// PoolToOwner is a free data retrieval call binding the contract method 0xd4cced9a.
//
// Solidity: function poolToOwner(address ) constant returns(address)
func (_PoolFactory *PoolFactoryCallerSession) PoolToOwner(arg0 common.Address) (common.Address, error) {
	return _PoolFactory.Contract.PoolToOwner(&_PoolFactory.CallOpts, arg0)
}

// CreatePool is a paid mutator transaction binding the contract method 0x9049f9d2.
//
// Solidity: function createPool(address _owner) returns(address)
func (_PoolFactory *PoolFactoryTransactor) CreatePool(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _PoolFactory.contract.Transact(opts, "createPool", _owner)
}

// CreatePool is a paid mutator transaction binding the contract method 0x9049f9d2.
//
// Solidity: function createPool(address _owner) returns(address)
func (_PoolFactory *PoolFactorySession) CreatePool(_owner common.Address) (*types.Transaction, error) {
	return _PoolFactory.Contract.CreatePool(&_PoolFactory.TransactOpts, _owner)
}

// CreatePool is a paid mutator transaction binding the contract method 0x9049f9d2.
//
// Solidity: function createPool(address _owner) returns(address)
func (_PoolFactory *PoolFactoryTransactorSession) CreatePool(_owner common.Address) (*types.Transaction, error) {
	return _PoolFactory.Contract.CreatePool(&_PoolFactory.TransactOpts, _owner)
}
