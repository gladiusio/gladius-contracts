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
const MarketBin = `0x60806040526005809080546001816001161561010002031660029004610026929190610090565b5034801561003357600080fd5b50604051604080610ed18339810180604052604081101561005357600080fd5b508051602090910151600280546001600160a01b039384166001600160a01b03199182161790915560038054939092169216919091179055610132565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106100c95780548555610105565b8280016001018555821561010557600052602060002091601f016020900482015b828111156101055782548255916001019190600101906100ea565b50610111929150610115565b5090565b61012f91905b80821115610111576000815560010161011b565b90565b610d90806101416000396000f3fe608060405234801561001057600080fd5b506004361061010b5760003560e01c80638da5cb5b116100a2578063b42cb6f511610071578063b42cb6f514610387578063d00e1e721461038f578063d88ff1f4146103c9578063db20ee33146103d1578063e9e24234146103fd5761010b565b80638da5cb5b1461032b5780639a06b11314610333578063a6f9dae11461033b578063aebce2d4146103615761010b565b806372591be5116100de57806372591be5146102cf57806373d4a13a146102f55780637cf25517146102fd578063893d20e8146103235761010b565b806317e421c8146101105780633bc5de301461018657806347064d6a146102035780636bd969b3146102ab575b600080fd5b6101366004803603602081101561012657600080fd5b50356001600160a01b031661041a565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561017257818101518382015260200161015a565b505050509050019250505060405180910390f35b61018e61048e565b6040805160208082528351818301528351919283929083019185019080838360005b838110156101c85781810151838201526020016101b0565b50505050905090810190601f1680156101f55780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6102a96004803603602081101561021957600080fd5b81019060208101813564010000000081111561023457600080fd5b82018360208201111561024657600080fd5b8035906020019184600183028401116401000000008311171561026857600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610525945050505050565b005b6102b36105ee565b604080516001600160a01b039092168252519081900360200190f35b6102a9600480360360208110156102e557600080fd5b50356001600160a01b03166105fd565b61018e610697565b6102b36004803603602081101561031357600080fd5b50356001600160a01b0316610725565b6102b36107ab565b6102b36107ba565b6102b36107c9565b6102a96004803603602081101561035157600080fd5b50356001600160a01b0316610849565b6101366004803603602081101561037757600080fd5b50356001600160a01b03166108e3565b6101366109c8565b6103b5600480360360208110156103a557600080fd5b50356001600160a01b0316610a29565b604080519115158252519081900360200190f35b610136610b98565b6102b3600480360360408110156103e757600080fd5b506001600160a01b038135169060200135610c70565b6102b36004803603602081101561041357600080fd5b5035610ca5565b6001600160a01b0381166000908152602081815260409182902080548351818402810184019094528084526060939283018282801561048257602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610464575b50505050509050919050565b60048054604080516020601f600260001961010060018816150201909516949094049384018190048102820181019092528281526060939092909183018282801561051a5780601f106104ef5761010080835404028352916020019161051a565b820191906000526020600020905b8154815290600101906020018083116104fd57829003601f168201915b505050505090505b90565b6002546005906001600160a01b031633146105d657604051600160e51b62461bcd0281526020600482019081528254600260001961010060018416150201909116046024830181905290918291604490910190849080156105c75780601f1061059c576101008083540402835291602001916105c7565b820191906000526020600020905b8154815290600101906020018083116105aa57829003601f168201915b50509250505060405180910390fd5b5080516105ea906004906020840190610ccc565b5050565b6003546001600160a01b031690565b6002546005906001600160a01b0316331461067457604051600160e51b62461bcd0281526020600482019081528254600260001961010060018416150201909116046024830181905290918291604490910190849080156105c75780601f1061059c576101008083540402835291602001916105c7565b50600380546001600160a01b0319166001600160a01b0392909216919091179055565b6004805460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152929183018282801561071d5780601f106106f25761010080835404028352916020019161071d565b820191906000526020600020905b81548152906001019060200180831161070057829003601f168201915b505050505081565b60035460408051600160e01b637cf255170281526001600160a01b03848116600483015291516000939290921691637cf2551791602480820192602092909190829003018186803b15801561077957600080fd5b505afa15801561078d573d6000803e3d6000fd5b505050506040513d60208110156107a357600080fd5b505192915050565b6002546001600160a01b031690565b6002546001600160a01b031681565b60035460408051600160e11b634824fce902815233600482015290516000926001600160a01b031691639049f9d291602480830192602092919082900301818787803b15801561081857600080fd5b505af115801561082c573d6000803e3d6000fd5b505050506040513d602081101561084257600080fd5b5051905090565b6002546005906001600160a01b031633146108c057604051600160e51b62461bcd0281526020600482019081528254600260001961010060018416150201909116046024830181905290918291604490910190849080156105c75780601f1061059c576101008083540402835291602001916105c7565b50600280546001600160a01b0319166001600160a01b0392909216919091179055565b60035460408051600160e01b63556f77db0281526001600160a01b0384811660048301529151606093929092169163556f77db91602480820192600092909190829003018186803b15801561093757600080fd5b505afa15801561094b573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052602081101561097457600080fd5b81019080805164010000000081111561098c57600080fd5b8201602081018481111561099f57600080fd5b81518560208202830111640100000000821117156109bc57600080fd5b50909695505050505050565b6060600180548060200260200160405190810160405280929190818152602001828054801561051a57602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610a02575050505050905090565b6002546000906005906001600160a01b03163314610aa357604051600160e51b62461bcd0281526020600482019081528254600260001961010060018416150201909116046024830181905290918291604490910190849080156105c75780601f1061059c576101008083540402835291602001916105c7565b506000829050600080826001600160a01b031663893d20e86040518163ffffffff1660e01b815260040160206040518083038186803b158015610ae557600080fd5b505afa158015610af9573d6000803e3d6000fd5b505050506040513d6020811015610b0f57600080fd5b50516001600160a01b0390811682526020828101939093526040909101600090812080546001818101835591835293822090930180549483166001600160a01b031995861617905582548084018455908390527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf601805491861691909316179091559050919050565b60035460408051600160e21b633623fc7d02815290516060926001600160a01b03169163d88ff1f4916004808301926000929190829003018186803b158015610be057600080fd5b505afa158015610bf4573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526020811015610c1d57600080fd5b810190808051640100000000811115610c3557600080fd5b82016020810184811115610c4857600080fd5b8151856020820283011164010000000082111715610c6557600080fd5b509094505050505090565b60006020528160005260406000208181548110610c8957fe5b6000918252602090912001546001600160a01b03169150829050565b60018181548110610cb257fe5b6000918252602090912001546001600160a01b0316905081565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610d0d57805160ff1916838001178555610d3a565b82800160010185558215610d3a579182015b82811115610d3a578251825591602001919060010190610d1f565b50610d46929150610d4a565b5090565b61052291905b80821115610d465760008155600101610d5056fea165627a7a72305820ed146cdfcc8f9a051808d1ee3a2b97bea312159a5c90de29d0523420c1799ef10029`

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
const PoolABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_node\",\"type\":\"address\"}],\"name\":\"isMasterNode\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_node\",\"type\":\"address\"}],\"name\":\"isApprovedNode\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAllApprovedNodes\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAllInfrastructureNodes\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_node\",\"type\":\"address\"}],\"name\":\"addApprovedNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"certificateBundle\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"poolDomain\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getData\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getSeedNode\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_url\",\"type\":\"string\"}],\"name\":\"setPoolDomain\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_node\",\"type\":\"address\"}],\"name\":\"removeInfrastructureNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_data\",\"type\":\"string\"}],\"name\":\"setData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"masterNodes\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cdnDomain\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_node\",\"type\":\"address\"}],\"name\":\"addMasterNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_node\",\"type\":\"string\"}],\"name\":\"setSeedNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_url\",\"type\":\"string\"}],\"name\":\"setCDNDomain\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"data\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"infrastructureNodes\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"approvedNodes\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCertificateBundle\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"infrastructureNodesList\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getPoolDomain\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_node\",\"type\":\"address\"}],\"name\":\"isInfrastructureNode\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_url\",\"type\":\"string\"}],\"name\":\"setCertificateBundle\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"masterNodesList\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAllMasterNodes\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_node\",\"type\":\"address\"}],\"name\":\"removeMasterNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"seedNode\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCDNDomain\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"approvedNodesList\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_node\",\"type\":\"address\"}],\"name\":\"removeApprovedNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_node\",\"type\":\"address\"}],\"name\":\"addInfrastructureNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// PoolBin is the compiled bytecode used for deploying new contracts.
const PoolBin = `0x608060405234801561001057600080fd5b506040516020806116b08339810180604052602081101561003057600080fd5b5051600680546001600160a01b0319166001600160a01b03909216919091179055611650806100606000396000f3fe608060405234801561001057600080fd5b506004361061021c5760003560e01c806378a3b06411610125578063beb09f2b116100ad578063e0a18ee71161007c578063e0a18ee714610890578063e2a5462f14610898578063ed7754cd146108a0578063ed82836e146108bd578063f7e9495a146108e35761021c565b8063beb09f2b146107a1578063c943230114610845578063cf94702314610862578063dd9cf9711461086a5761021c565b8063893d20e8116100f4578063893d20e81461073d5780638da5cb5b14610745578063914964771461074d5780639240a18214610755578063a6f9dae11461077b5761021c565b806378a3b064146106b0578063806ba4d2146106d6578063836ac1db146106fc57806383d96809146107045761021c565b80633d2d3a67116101a8578063653f2c7811610177578063653f2c78146105325780636b8a98a91461053a5780636bc3af271461056057806372ae65d51461060457806373d4a13a146106a85761021c565b80633d2d3a671461039e5780633f2706d71461044257806347064d6a1461046857806350ca875f1461050c5761021c565b8063324045de116101ef578063324045de146102e157806338562934146103095780633b5e0ebd146103865780633bc5de301461038e5780633c714980146103965761021c565b80630458a33514610221578063053f6ef61461025b57806306ddccbd1461028157806319dc52a0146102d9575b600080fd5b6102476004803603602081101561023757600080fd5b50356001600160a01b0316610909565b604080519115158252519081900360200190f35b6102476004803603602081101561027157600080fd5b50356001600160a01b0316610927565b610289610945565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156102c55781810151838201526020016102ad565b505050509050019250505060405180910390f35b6102896109a8565b610307600480360360208110156102f757600080fd5b50356001600160a01b0316610a08565b005b610311610abe565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561034b578181015183820152602001610333565b50505050905090810190601f1680156103785780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b610311610b4c565b610311610ba7565b610311610c34565b610307600480360360208110156103b457600080fd5b810190602081018135600160201b8111156103ce57600080fd5b8201836020820111156103e057600080fd5b803590602001918460018302840111600160201b8311171561040157600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610c95945050505050565b6103076004803603602081101561045857600080fd5b50356001600160a01b0316610cfc565b6103076004803603602081101561047e57600080fd5b810190602081018135600160201b81111561049857600080fd5b8201836020820111156104aa57600080fd5b803590602001918460018302840111600160201b831117156104cb57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610d91945050505050565b6102476004803603602081101561052257600080fd5b50356001600160a01b0316610df4565b610311610e09565b6103076004803603602081101561055057600080fd5b50356001600160a01b0316610e64565b6103076004803603602081101561057657600080fd5b810190602081018135600160201b81111561059057600080fd5b8201836020820111156105a257600080fd5b803590602001918460018302840111600160201b831117156105c357600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610f17945050505050565b6103076004803603602081101561061a57600080fd5b810190602081018135600160201b81111561063457600080fd5b82018360208201111561064657600080fd5b803590602001918460018302840111600160201b8311171561066757600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610f7a945050505050565b610311610fdd565b610247600480360360208110156106c657600080fd5b50356001600160a01b0316611038565b610247600480360360208110156106ec57600080fd5b50356001600160a01b031661104d565b610311611062565b6107216004803603602081101561071a57600080fd5b50356110c3565b604080516001600160a01b039092168252519081900360200190f35b6107216110ea565b6107216110f9565b610311611108565b6102476004803603602081101561076b57600080fd5b50356001600160a01b0316611169565b6103076004803603602081101561079157600080fd5b50356001600160a01b0316611187565b610307600480360360208110156107b757600080fd5b810190602081018135600160201b8111156107d157600080fd5b8201836020820111156107e357600080fd5b803590602001918460018302840111600160201b8311171561080457600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506111f9945050505050565b6107216004803603602081101561085b57600080fd5b503561125c565b610289611269565b6103076004803603602081101561088057600080fd5b50356001600160a01b03166112c9565b61031161135b565b6103116113b6565b610721600480360360208110156108b657600080fd5b5035611417565b610307600480360360208110156108d357600080fd5b50356001600160a01b0316611424565b610307600480360360208110156108f957600080fd5b50356001600160a01b03166114b6565b6001600160a01b031660009081526020819052604090205460ff1690565b6001600160a01b031660009081526002602052604090205460ff1690565b6060600380548060200260200160405190810160405280929190818152602001828054801561099d57602002820191906000526020600020905b81546001600160a01b0316815260019091019060200180831161097f575b505050505090505b90565b6060600580548060200260200160405190810160405280929190818152602001828054801561099d576020028201919060005260206000209081546001600160a01b0316815260019091019060200180831161097f575050505050905090565b6006546001600160a01b03163314610a585760408051600160e51b62461bcd0281526020600482015260196024820152600080516020611605833981519152604482015290519081900360640190fd5b6001600160a01b03166000818152600260205260408120805460ff191660019081179091556003805491820181559091527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b0180546001600160a01b0319169091179055565b600b805460408051602060026001851615610100026000190190941693909304601f81018490048402820184019092528181529291830182828015610b445780601f10610b1957610100808354040283529160200191610b44565b820191906000526020600020905b815481529060010190602001808311610b2757829003601f168201915b505050505081565b6009805460408051602060026001851615610100026000190190941693909304601f81018490048402820184019092528181529291830182828015610b445780601f10610b1957610100808354040283529160200191610b44565b60088054604080516020601f600260001961010060018816150201909516949094049384018190048102820181019092528281526060939092909183018282801561099d5780601f10610c085761010080835404028352916020019161099d565b820191906000526020600020905b815481529060010190602001808311610c1657509395945050505050565b60078054604080516020601f600260001961010060018816150201909516949094049384018190048102820181019092528281526060939092909183018282801561099d5780601f10610c085761010080835404028352916020019161099d565b6006546001600160a01b03163314610ce55760408051600160e51b62461bcd0281526020600482015260196024820152600080516020611605833981519152604482015290519081900360640190fd5b8051610cf890600990602084019061156c565b5050565b6006546001600160a01b03163314610d4c5760408051600160e51b62461bcd0281526020600482015260196024820152600080516020611605833981519152604482015290519081900360640190fd5b6001600160a01b03811660009081526004602052604090205460ff1615610d8e576001600160a01b0381166000908152600460205260409020805460ff191690555b50565b6006546001600160a01b03163314610de15760408051600160e51b62461bcd0281526020600482015260196024820152600080516020611605833981519152604482015290519081900360640190fd5b8051610cf890600890602084019061156c565b60006020819052908152604090205460ff1681565b600a805460408051602060026001851615610100026000190190941693909304601f81018490048402820184019092528181529291830182828015610b445780601f10610b1957610100808354040283529160200191610b44565b6006546001600160a01b03163314610eb45760408051600160e51b62461bcd0281526020600482015260196024820152600080516020611605833981519152604482015290519081900360640190fd5b6001600160a01b03166000818152602081905260408120805460ff191660019081179091558054808201825591527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf60180546001600160a01b0319169091179055565b6006546001600160a01b03163314610f675760408051600160e51b62461bcd0281526020600482015260196024820152600080516020611605833981519152604482015290519081900360640190fd5b8051610cf890600790602084019061156c565b6006546001600160a01b03163314610fca5760408051600160e51b62461bcd0281526020600482015260196024820152600080516020611605833981519152604482015290519081900360640190fd5b8051610cf890600a90602084019061156c565b6008805460408051602060026001851615610100026000190190941693909304601f81018490048402820184019092528181529291830182828015610b445780601f10610b1957610100808354040283529160200191610b44565b60046020526000908152604090205460ff1681565b60026020526000908152604090205460ff1681565b600b8054604080516020601f600260001961010060018816150201909516949094049384018190048102820181019092528281526060939092909183018282801561099d5780601f10610c085761010080835404028352916020019161099d565b600581815481106110d057fe5b6000918252602090912001546001600160a01b0316905081565b6006546001600160a01b031690565b6006546001600160a01b031681565b60098054604080516020601f600260001961010060018816150201909516949094049384018190048102820181019092528281526060939092909183018282801561099d5780601f10610c085761010080835404028352916020019161099d565b6001600160a01b031660009081526004602052604090205460ff1690565b6006546001600160a01b031633146111d75760408051600160e51b62461bcd0281526020600482015260196024820152600080516020611605833981519152604482015290519081900360640190fd5b600680546001600160a01b0319166001600160a01b0392909216919091179055565b6006546001600160a01b031633146112495760408051600160e51b62461bcd0281526020600482015260196024820152600080516020611605833981519152604482015290519081900360640190fd5b8051610cf890600b90602084019061156c565b600181815481106110d057fe5b6060600180548060200260200160405190810160405280929190818152602001828054801561099d576020028201919060005260206000209081546001600160a01b0316815260019091019060200180831161097f575050505050905090565b6006546001600160a01b031633146113195760408051600160e51b62461bcd0281526020600482015260196024820152600080516020611605833981519152604482015290519081900360640190fd5b6001600160a01b03811660009081526020819052604090205460ff1615610d8e576001600160a01b03166000908152602081905260409020805460ff19169055565b6007805460408051602060026001851615610100026000190190941693909304601f81018490048402820184019092528181529291830182828015610b445780601f10610b1957610100808354040283529160200191610b44565b600a8054604080516020601f600260001961010060018816150201909516949094049384018190048102820181019092528281526060939092909183018282801561099d5780601f10610c085761010080835404028352916020019161099d565b600381815481106110d057fe5b6006546001600160a01b031633146114745760408051600160e51b62461bcd0281526020600482015260196024820152600080516020611605833981519152604482015290519081900360640190fd5b6001600160a01b03811660009081526002602052604090205460ff1615610d8e576001600160a01b03166000908152600260205260409020805460ff19169055565b6006546001600160a01b031633146115065760408051600160e51b62461bcd0281526020600482015260196024820152600080516020611605833981519152604482015290519081900360640190fd5b6001600160a01b03166000818152600460205260408120805460ff191660019081179091556005805491820181559091527f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db00180546001600160a01b0319169091179055565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106115ad57805160ff19168380011785556115da565b828001600101855582156115da579182015b828111156115da5782518255916020019190600101906115bf565b506115e69291506115ea565b5090565b6109a591905b808211156115e657600081556001016115f056fe4d757374206265206f776e6572206f6620636f6e747261637400000000000000a165627a7a723058204458deeb964872d760a39485321701772e2e3c4400cec77d9bcf737ae90804e70029`

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

// ApprovedNodesList is a free data retrieval call binding the contract method 0xed7754cd.
//
// Solidity: function approvedNodesList(uint256 ) constant returns(address)
func (_Pool *PoolCaller) ApprovedNodesList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Pool.contract.Call(opts, out, "approvedNodesList", arg0)
	return *ret0, err
}

// ApprovedNodesList is a free data retrieval call binding the contract method 0xed7754cd.
//
// Solidity: function approvedNodesList(uint256 ) constant returns(address)
func (_Pool *PoolSession) ApprovedNodesList(arg0 *big.Int) (common.Address, error) {
	return _Pool.Contract.ApprovedNodesList(&_Pool.CallOpts, arg0)
}

// ApprovedNodesList is a free data retrieval call binding the contract method 0xed7754cd.
//
// Solidity: function approvedNodesList(uint256 ) constant returns(address)
func (_Pool *PoolCallerSession) ApprovedNodesList(arg0 *big.Int) (common.Address, error) {
	return _Pool.Contract.ApprovedNodesList(&_Pool.CallOpts, arg0)
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

// CertificateBundle is a free data retrieval call binding the contract method 0x38562934.
//
// Solidity: function certificateBundle() constant returns(string)
func (_Pool *PoolCaller) CertificateBundle(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Pool.contract.Call(opts, out, "certificateBundle")
	return *ret0, err
}

// CertificateBundle is a free data retrieval call binding the contract method 0x38562934.
//
// Solidity: function certificateBundle() constant returns(string)
func (_Pool *PoolSession) CertificateBundle() (string, error) {
	return _Pool.Contract.CertificateBundle(&_Pool.CallOpts)
}

// CertificateBundle is a free data retrieval call binding the contract method 0x38562934.
//
// Solidity: function certificateBundle() constant returns(string)
func (_Pool *PoolCallerSession) CertificateBundle() (string, error) {
	return _Pool.Contract.CertificateBundle(&_Pool.CallOpts)
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

// GetAllApprovedNodes is a free data retrieval call binding the contract method 0x06ddccbd.
//
// Solidity: function getAllApprovedNodes() constant returns(address[])
func (_Pool *PoolCaller) GetAllApprovedNodes(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Pool.contract.Call(opts, out, "getAllApprovedNodes")
	return *ret0, err
}

// GetAllApprovedNodes is a free data retrieval call binding the contract method 0x06ddccbd.
//
// Solidity: function getAllApprovedNodes() constant returns(address[])
func (_Pool *PoolSession) GetAllApprovedNodes() ([]common.Address, error) {
	return _Pool.Contract.GetAllApprovedNodes(&_Pool.CallOpts)
}

// GetAllApprovedNodes is a free data retrieval call binding the contract method 0x06ddccbd.
//
// Solidity: function getAllApprovedNodes() constant returns(address[])
func (_Pool *PoolCallerSession) GetAllApprovedNodes() ([]common.Address, error) {
	return _Pool.Contract.GetAllApprovedNodes(&_Pool.CallOpts)
}

// GetAllInfrastructureNodes is a free data retrieval call binding the contract method 0x19dc52a0.
//
// Solidity: function getAllInfrastructureNodes() constant returns(address[])
func (_Pool *PoolCaller) GetAllInfrastructureNodes(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Pool.contract.Call(opts, out, "getAllInfrastructureNodes")
	return *ret0, err
}

// GetAllInfrastructureNodes is a free data retrieval call binding the contract method 0x19dc52a0.
//
// Solidity: function getAllInfrastructureNodes() constant returns(address[])
func (_Pool *PoolSession) GetAllInfrastructureNodes() ([]common.Address, error) {
	return _Pool.Contract.GetAllInfrastructureNodes(&_Pool.CallOpts)
}

// GetAllInfrastructureNodes is a free data retrieval call binding the contract method 0x19dc52a0.
//
// Solidity: function getAllInfrastructureNodes() constant returns(address[])
func (_Pool *PoolCallerSession) GetAllInfrastructureNodes() ([]common.Address, error) {
	return _Pool.Contract.GetAllInfrastructureNodes(&_Pool.CallOpts)
}

// GetAllMasterNodes is a free data retrieval call binding the contract method 0xcf947023.
//
// Solidity: function getAllMasterNodes() constant returns(address[])
func (_Pool *PoolCaller) GetAllMasterNodes(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Pool.contract.Call(opts, out, "getAllMasterNodes")
	return *ret0, err
}

// GetAllMasterNodes is a free data retrieval call binding the contract method 0xcf947023.
//
// Solidity: function getAllMasterNodes() constant returns(address[])
func (_Pool *PoolSession) GetAllMasterNodes() ([]common.Address, error) {
	return _Pool.Contract.GetAllMasterNodes(&_Pool.CallOpts)
}

// GetAllMasterNodes is a free data retrieval call binding the contract method 0xcf947023.
//
// Solidity: function getAllMasterNodes() constant returns(address[])
func (_Pool *PoolCallerSession) GetAllMasterNodes() ([]common.Address, error) {
	return _Pool.Contract.GetAllMasterNodes(&_Pool.CallOpts)
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

// GetCertificateBundle is a free data retrieval call binding the contract method 0x836ac1db.
//
// Solidity: function getCertificateBundle() constant returns(string)
func (_Pool *PoolCaller) GetCertificateBundle(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Pool.contract.Call(opts, out, "getCertificateBundle")
	return *ret0, err
}

// GetCertificateBundle is a free data retrieval call binding the contract method 0x836ac1db.
//
// Solidity: function getCertificateBundle() constant returns(string)
func (_Pool *PoolSession) GetCertificateBundle() (string, error) {
	return _Pool.Contract.GetCertificateBundle(&_Pool.CallOpts)
}

// GetCertificateBundle is a free data retrieval call binding the contract method 0x836ac1db.
//
// Solidity: function getCertificateBundle() constant returns(string)
func (_Pool *PoolCallerSession) GetCertificateBundle() (string, error) {
	return _Pool.Contract.GetCertificateBundle(&_Pool.CallOpts)
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

// InfrastructureNodes is a free data retrieval call binding the contract method 0x78a3b064.
//
// Solidity: function infrastructureNodes(address ) constant returns(bool)
func (_Pool *PoolCaller) InfrastructureNodes(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Pool.contract.Call(opts, out, "infrastructureNodes", arg0)
	return *ret0, err
}

// InfrastructureNodes is a free data retrieval call binding the contract method 0x78a3b064.
//
// Solidity: function infrastructureNodes(address ) constant returns(bool)
func (_Pool *PoolSession) InfrastructureNodes(arg0 common.Address) (bool, error) {
	return _Pool.Contract.InfrastructureNodes(&_Pool.CallOpts, arg0)
}

// InfrastructureNodes is a free data retrieval call binding the contract method 0x78a3b064.
//
// Solidity: function infrastructureNodes(address ) constant returns(bool)
func (_Pool *PoolCallerSession) InfrastructureNodes(arg0 common.Address) (bool, error) {
	return _Pool.Contract.InfrastructureNodes(&_Pool.CallOpts, arg0)
}

// InfrastructureNodesList is a free data retrieval call binding the contract method 0x83d96809.
//
// Solidity: function infrastructureNodesList(uint256 ) constant returns(address)
func (_Pool *PoolCaller) InfrastructureNodesList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Pool.contract.Call(opts, out, "infrastructureNodesList", arg0)
	return *ret0, err
}

// InfrastructureNodesList is a free data retrieval call binding the contract method 0x83d96809.
//
// Solidity: function infrastructureNodesList(uint256 ) constant returns(address)
func (_Pool *PoolSession) InfrastructureNodesList(arg0 *big.Int) (common.Address, error) {
	return _Pool.Contract.InfrastructureNodesList(&_Pool.CallOpts, arg0)
}

// InfrastructureNodesList is a free data retrieval call binding the contract method 0x83d96809.
//
// Solidity: function infrastructureNodesList(uint256 ) constant returns(address)
func (_Pool *PoolCallerSession) InfrastructureNodesList(arg0 *big.Int) (common.Address, error) {
	return _Pool.Contract.InfrastructureNodesList(&_Pool.CallOpts, arg0)
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

// IsInfrastructureNode is a free data retrieval call binding the contract method 0x9240a182.
//
// Solidity: function isInfrastructureNode(address _node) constant returns(bool)
func (_Pool *PoolCaller) IsInfrastructureNode(opts *bind.CallOpts, _node common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Pool.contract.Call(opts, out, "isInfrastructureNode", _node)
	return *ret0, err
}

// IsInfrastructureNode is a free data retrieval call binding the contract method 0x9240a182.
//
// Solidity: function isInfrastructureNode(address _node) constant returns(bool)
func (_Pool *PoolSession) IsInfrastructureNode(_node common.Address) (bool, error) {
	return _Pool.Contract.IsInfrastructureNode(&_Pool.CallOpts, _node)
}

// IsInfrastructureNode is a free data retrieval call binding the contract method 0x9240a182.
//
// Solidity: function isInfrastructureNode(address _node) constant returns(bool)
func (_Pool *PoolCallerSession) IsInfrastructureNode(_node common.Address) (bool, error) {
	return _Pool.Contract.IsInfrastructureNode(&_Pool.CallOpts, _node)
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

// MasterNodesList is a free data retrieval call binding the contract method 0xc9432301.
//
// Solidity: function masterNodesList(uint256 ) constant returns(address)
func (_Pool *PoolCaller) MasterNodesList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Pool.contract.Call(opts, out, "masterNodesList", arg0)
	return *ret0, err
}

// MasterNodesList is a free data retrieval call binding the contract method 0xc9432301.
//
// Solidity: function masterNodesList(uint256 ) constant returns(address)
func (_Pool *PoolSession) MasterNodesList(arg0 *big.Int) (common.Address, error) {
	return _Pool.Contract.MasterNodesList(&_Pool.CallOpts, arg0)
}

// MasterNodesList is a free data retrieval call binding the contract method 0xc9432301.
//
// Solidity: function masterNodesList(uint256 ) constant returns(address)
func (_Pool *PoolCallerSession) MasterNodesList(arg0 *big.Int) (common.Address, error) {
	return _Pool.Contract.MasterNodesList(&_Pool.CallOpts, arg0)
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

// AddInfrastructureNode is a paid mutator transaction binding the contract method 0xf7e9495a.
//
// Solidity: function addInfrastructureNode(address _node) returns()
func (_Pool *PoolTransactor) AddInfrastructureNode(opts *bind.TransactOpts, _node common.Address) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "addInfrastructureNode", _node)
}

// AddInfrastructureNode is a paid mutator transaction binding the contract method 0xf7e9495a.
//
// Solidity: function addInfrastructureNode(address _node) returns()
func (_Pool *PoolSession) AddInfrastructureNode(_node common.Address) (*types.Transaction, error) {
	return _Pool.Contract.AddInfrastructureNode(&_Pool.TransactOpts, _node)
}

// AddInfrastructureNode is a paid mutator transaction binding the contract method 0xf7e9495a.
//
// Solidity: function addInfrastructureNode(address _node) returns()
func (_Pool *PoolTransactorSession) AddInfrastructureNode(_node common.Address) (*types.Transaction, error) {
	return _Pool.Contract.AddInfrastructureNode(&_Pool.TransactOpts, _node)
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

// RemoveInfrastructureNode is a paid mutator transaction binding the contract method 0x3f2706d7.
//
// Solidity: function removeInfrastructureNode(address _node) returns()
func (_Pool *PoolTransactor) RemoveInfrastructureNode(opts *bind.TransactOpts, _node common.Address) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "removeInfrastructureNode", _node)
}

// RemoveInfrastructureNode is a paid mutator transaction binding the contract method 0x3f2706d7.
//
// Solidity: function removeInfrastructureNode(address _node) returns()
func (_Pool *PoolSession) RemoveInfrastructureNode(_node common.Address) (*types.Transaction, error) {
	return _Pool.Contract.RemoveInfrastructureNode(&_Pool.TransactOpts, _node)
}

// RemoveInfrastructureNode is a paid mutator transaction binding the contract method 0x3f2706d7.
//
// Solidity: function removeInfrastructureNode(address _node) returns()
func (_Pool *PoolTransactorSession) RemoveInfrastructureNode(_node common.Address) (*types.Transaction, error) {
	return _Pool.Contract.RemoveInfrastructureNode(&_Pool.TransactOpts, _node)
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

// SetCertificateBundle is a paid mutator transaction binding the contract method 0xbeb09f2b.
//
// Solidity: function setCertificateBundle(string _url) returns()
func (_Pool *PoolTransactor) SetCertificateBundle(opts *bind.TransactOpts, _url string) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "setCertificateBundle", _url)
}

// SetCertificateBundle is a paid mutator transaction binding the contract method 0xbeb09f2b.
//
// Solidity: function setCertificateBundle(string _url) returns()
func (_Pool *PoolSession) SetCertificateBundle(_url string) (*types.Transaction, error) {
	return _Pool.Contract.SetCertificateBundle(&_Pool.TransactOpts, _url)
}

// SetCertificateBundle is a paid mutator transaction binding the contract method 0xbeb09f2b.
//
// Solidity: function setCertificateBundle(string _url) returns()
func (_Pool *PoolTransactorSession) SetCertificateBundle(_url string) (*types.Transaction, error) {
	return _Pool.Contract.SetCertificateBundle(&_Pool.TransactOpts, _url)
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
const PoolFactoryBin = `0x608060405234801561001057600080fd5b50611afd806100206000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c80637cf255171161005b5780637cf255171461015d5780639049f9d214610183578063d4cced9a146101a9578063d88ff1f4146101cf5761007d565b806312708d721461008257806341d1de97146100ca578063556f77db146100e7575b600080fd5b6100ae6004803603604081101561009857600080fd5b506001600160a01b0381351690602001356101d7565b604080516001600160a01b039092168252519081900360200190f35b6100ae600480360360208110156100e057600080fd5b503561020c565b61010d600480360360208110156100fd57600080fd5b50356001600160a01b0316610233565b60408051602080825283518183015283519192839290830191858101910280838360005b83811015610149578181015183820152602001610131565b505050509050019250505060405180910390f35b6100ae6004803603602081101561017357600080fd5b50356001600160a01b03166102a7565b6100ae6004803603602081101561019957600080fd5b50356001600160a01b03166102c5565b6100ae600480360360208110156101bf57600080fd5b50356001600160a01b0316610397565b61010d6103b2565b600060205281600052604060002081815481106101f057fe5b6000918252602090912001546001600160a01b03169150829050565b6002818154811061021957fe5b6000918252602090912001546001600160a01b0316905081565b6001600160a01b0381166000908152602081815260409182902080548351818402810184019094528084526060939283018282801561029b57602002820191906000526020600020905b81546001600160a01b0316815260019091019060200180831161027d575b50505050509050919050565b6001600160a01b039081166000908152600160205260409020541690565b600080826040516102d590610414565b6001600160a01b03909116815260405190819003602001906000f080158015610302573d6000803e3d6000fd5b506001600160a01b038085166000818152602081815260408083208054600181810183559185528385200180549688166001600160a01b0319978816811790915580855292819052908320805486169094179093556002805493840181559091527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace9091018054909216179055915050919050565b6001602052600090815260409020546001600160a01b031681565b6060600280548060200260200160405190810160405280929190818152602001828054801561040a57602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116103ec575b5050505050905090565b6116b0806104228339019056fe608060405234801561001057600080fd5b506040516020806116b08339810180604052602081101561003057600080fd5b5051600680546001600160a01b0319166001600160a01b03909216919091179055611650806100606000396000f3fe608060405234801561001057600080fd5b506004361061021c5760003560e01c806378a3b06411610125578063beb09f2b116100ad578063e0a18ee71161007c578063e0a18ee714610890578063e2a5462f14610898578063ed7754cd146108a0578063ed82836e146108bd578063f7e9495a146108e35761021c565b8063beb09f2b146107a1578063c943230114610845578063cf94702314610862578063dd9cf9711461086a5761021c565b8063893d20e8116100f4578063893d20e81461073d5780638da5cb5b14610745578063914964771461074d5780639240a18214610755578063a6f9dae11461077b5761021c565b806378a3b064146106b0578063806ba4d2146106d6578063836ac1db146106fc57806383d96809146107045761021c565b80633d2d3a67116101a8578063653f2c7811610177578063653f2c78146105325780636b8a98a91461053a5780636bc3af271461056057806372ae65d51461060457806373d4a13a146106a85761021c565b80633d2d3a671461039e5780633f2706d71461044257806347064d6a1461046857806350ca875f1461050c5761021c565b8063324045de116101ef578063324045de146102e157806338562934146103095780633b5e0ebd146103865780633bc5de301461038e5780633c714980146103965761021c565b80630458a33514610221578063053f6ef61461025b57806306ddccbd1461028157806319dc52a0146102d9575b600080fd5b6102476004803603602081101561023757600080fd5b50356001600160a01b0316610909565b604080519115158252519081900360200190f35b6102476004803603602081101561027157600080fd5b50356001600160a01b0316610927565b610289610945565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156102c55781810151838201526020016102ad565b505050509050019250505060405180910390f35b6102896109a8565b610307600480360360208110156102f757600080fd5b50356001600160a01b0316610a08565b005b610311610abe565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561034b578181015183820152602001610333565b50505050905090810190601f1680156103785780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b610311610b4c565b610311610ba7565b610311610c34565b610307600480360360208110156103b457600080fd5b810190602081018135600160201b8111156103ce57600080fd5b8201836020820111156103e057600080fd5b803590602001918460018302840111600160201b8311171561040157600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610c95945050505050565b6103076004803603602081101561045857600080fd5b50356001600160a01b0316610cfc565b6103076004803603602081101561047e57600080fd5b810190602081018135600160201b81111561049857600080fd5b8201836020820111156104aa57600080fd5b803590602001918460018302840111600160201b831117156104cb57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610d91945050505050565b6102476004803603602081101561052257600080fd5b50356001600160a01b0316610df4565b610311610e09565b6103076004803603602081101561055057600080fd5b50356001600160a01b0316610e64565b6103076004803603602081101561057657600080fd5b810190602081018135600160201b81111561059057600080fd5b8201836020820111156105a257600080fd5b803590602001918460018302840111600160201b831117156105c357600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610f17945050505050565b6103076004803603602081101561061a57600080fd5b810190602081018135600160201b81111561063457600080fd5b82018360208201111561064657600080fd5b803590602001918460018302840111600160201b8311171561066757600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610f7a945050505050565b610311610fdd565b610247600480360360208110156106c657600080fd5b50356001600160a01b0316611038565b610247600480360360208110156106ec57600080fd5b50356001600160a01b031661104d565b610311611062565b6107216004803603602081101561071a57600080fd5b50356110c3565b604080516001600160a01b039092168252519081900360200190f35b6107216110ea565b6107216110f9565b610311611108565b6102476004803603602081101561076b57600080fd5b50356001600160a01b0316611169565b6103076004803603602081101561079157600080fd5b50356001600160a01b0316611187565b610307600480360360208110156107b757600080fd5b810190602081018135600160201b8111156107d157600080fd5b8201836020820111156107e357600080fd5b803590602001918460018302840111600160201b8311171561080457600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506111f9945050505050565b6107216004803603602081101561085b57600080fd5b503561125c565b610289611269565b6103076004803603602081101561088057600080fd5b50356001600160a01b03166112c9565b61031161135b565b6103116113b6565b610721600480360360208110156108b657600080fd5b5035611417565b610307600480360360208110156108d357600080fd5b50356001600160a01b0316611424565b610307600480360360208110156108f957600080fd5b50356001600160a01b03166114b6565b6001600160a01b031660009081526020819052604090205460ff1690565b6001600160a01b031660009081526002602052604090205460ff1690565b6060600380548060200260200160405190810160405280929190818152602001828054801561099d57602002820191906000526020600020905b81546001600160a01b0316815260019091019060200180831161097f575b505050505090505b90565b6060600580548060200260200160405190810160405280929190818152602001828054801561099d576020028201919060005260206000209081546001600160a01b0316815260019091019060200180831161097f575050505050905090565b6006546001600160a01b03163314610a585760408051600160e51b62461bcd0281526020600482015260196024820152600080516020611605833981519152604482015290519081900360640190fd5b6001600160a01b03166000818152600260205260408120805460ff191660019081179091556003805491820181559091527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b0180546001600160a01b0319169091179055565b600b805460408051602060026001851615610100026000190190941693909304601f81018490048402820184019092528181529291830182828015610b445780601f10610b1957610100808354040283529160200191610b44565b820191906000526020600020905b815481529060010190602001808311610b2757829003601f168201915b505050505081565b6009805460408051602060026001851615610100026000190190941693909304601f81018490048402820184019092528181529291830182828015610b445780601f10610b1957610100808354040283529160200191610b44565b60088054604080516020601f600260001961010060018816150201909516949094049384018190048102820181019092528281526060939092909183018282801561099d5780601f10610c085761010080835404028352916020019161099d565b820191906000526020600020905b815481529060010190602001808311610c1657509395945050505050565b60078054604080516020601f600260001961010060018816150201909516949094049384018190048102820181019092528281526060939092909183018282801561099d5780601f10610c085761010080835404028352916020019161099d565b6006546001600160a01b03163314610ce55760408051600160e51b62461bcd0281526020600482015260196024820152600080516020611605833981519152604482015290519081900360640190fd5b8051610cf890600990602084019061156c565b5050565b6006546001600160a01b03163314610d4c5760408051600160e51b62461bcd0281526020600482015260196024820152600080516020611605833981519152604482015290519081900360640190fd5b6001600160a01b03811660009081526004602052604090205460ff1615610d8e576001600160a01b0381166000908152600460205260409020805460ff191690555b50565b6006546001600160a01b03163314610de15760408051600160e51b62461bcd0281526020600482015260196024820152600080516020611605833981519152604482015290519081900360640190fd5b8051610cf890600890602084019061156c565b60006020819052908152604090205460ff1681565b600a805460408051602060026001851615610100026000190190941693909304601f81018490048402820184019092528181529291830182828015610b445780601f10610b1957610100808354040283529160200191610b44565b6006546001600160a01b03163314610eb45760408051600160e51b62461bcd0281526020600482015260196024820152600080516020611605833981519152604482015290519081900360640190fd5b6001600160a01b03166000818152602081905260408120805460ff191660019081179091558054808201825591527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf60180546001600160a01b0319169091179055565b6006546001600160a01b03163314610f675760408051600160e51b62461bcd0281526020600482015260196024820152600080516020611605833981519152604482015290519081900360640190fd5b8051610cf890600790602084019061156c565b6006546001600160a01b03163314610fca5760408051600160e51b62461bcd0281526020600482015260196024820152600080516020611605833981519152604482015290519081900360640190fd5b8051610cf890600a90602084019061156c565b6008805460408051602060026001851615610100026000190190941693909304601f81018490048402820184019092528181529291830182828015610b445780601f10610b1957610100808354040283529160200191610b44565b60046020526000908152604090205460ff1681565b60026020526000908152604090205460ff1681565b600b8054604080516020601f600260001961010060018816150201909516949094049384018190048102820181019092528281526060939092909183018282801561099d5780601f10610c085761010080835404028352916020019161099d565b600581815481106110d057fe5b6000918252602090912001546001600160a01b0316905081565b6006546001600160a01b031690565b6006546001600160a01b031681565b60098054604080516020601f600260001961010060018816150201909516949094049384018190048102820181019092528281526060939092909183018282801561099d5780601f10610c085761010080835404028352916020019161099d565b6001600160a01b031660009081526004602052604090205460ff1690565b6006546001600160a01b031633146111d75760408051600160e51b62461bcd0281526020600482015260196024820152600080516020611605833981519152604482015290519081900360640190fd5b600680546001600160a01b0319166001600160a01b0392909216919091179055565b6006546001600160a01b031633146112495760408051600160e51b62461bcd0281526020600482015260196024820152600080516020611605833981519152604482015290519081900360640190fd5b8051610cf890600b90602084019061156c565b600181815481106110d057fe5b6060600180548060200260200160405190810160405280929190818152602001828054801561099d576020028201919060005260206000209081546001600160a01b0316815260019091019060200180831161097f575050505050905090565b6006546001600160a01b031633146113195760408051600160e51b62461bcd0281526020600482015260196024820152600080516020611605833981519152604482015290519081900360640190fd5b6001600160a01b03811660009081526020819052604090205460ff1615610d8e576001600160a01b03166000908152602081905260409020805460ff19169055565b6007805460408051602060026001851615610100026000190190941693909304601f81018490048402820184019092528181529291830182828015610b445780601f10610b1957610100808354040283529160200191610b44565b600a8054604080516020601f600260001961010060018816150201909516949094049384018190048102820181019092528281526060939092909183018282801561099d5780601f10610c085761010080835404028352916020019161099d565b600381815481106110d057fe5b6006546001600160a01b031633146114745760408051600160e51b62461bcd0281526020600482015260196024820152600080516020611605833981519152604482015290519081900360640190fd5b6001600160a01b03811660009081526002602052604090205460ff1615610d8e576001600160a01b03166000908152600260205260409020805460ff19169055565b6006546001600160a01b031633146115065760408051600160e51b62461bcd0281526020600482015260196024820152600080516020611605833981519152604482015290519081900360640190fd5b6001600160a01b03166000818152600460205260408120805460ff191660019081179091556005805491820181559091527f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db00180546001600160a01b0319169091179055565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106115ad57805160ff19168380011785556115da565b828001600101855582156115da579182015b828111156115da5782518255916020019190600101906115bf565b506115e69291506115ea565b5090565b6109a591905b808211156115e657600081556001016115f056fe4d757374206265206f776e6572206f6620636f6e747261637400000000000000a165627a7a723058204458deeb964872d760a39485321701772e2e3c4400cec77d9bcf737ae90804e70029a165627a7a72305820c311c6cf57b91a4d38cc0ef134906bece0bec1fe131ef54074504b36e8b4c46a0029`

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
