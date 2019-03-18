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

// PoolABI is the input ABI used to generate the binding from.
const PoolABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_node\",\"type\":\"address\"}],\"name\":\"isMasterNode\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_node\",\"type\":\"address\"}],\"name\":\"isApprovedNode\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_node\",\"type\":\"address\"}],\"name\":\"addApprovedNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"poolDomain\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getData\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getSeedNode\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_url\",\"type\":\"string\"}],\"name\":\"setPoolDomain\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_data\",\"type\":\"string\"}],\"name\":\"setData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"masterNodes\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_node\",\"type\":\"address\"}],\"name\":\"addMasterNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_node\",\"type\":\"string\"}],\"name\":\"setSeedNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"data\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"approvedNodes\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getPoolDomain\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_node\",\"type\":\"address\"}],\"name\":\"removeMasterNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"seedNode\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_node\",\"type\":\"address\"}],\"name\":\"removeApprovedNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// PoolBin is the compiled bytecode used for deploying new contracts.
const PoolBin = `0x608060405234801561001057600080fd5b50604051602080610da08339810180604052602081101561003057600080fd5b505160028054600160a060020a031916600160a060020a03909216919091179055610d40806100606000396000f3fe608060405234801561001057600080fd5b5060043610610149576000357c0100000000000000000000000000000000000000000000000000000000900480636bc3af27116100ca578063914964771161008e57806391496477146104fb578063a6f9dae114610503578063dd9cf97114610529578063e0a18ee71461054f578063ed82836e1461055757610149565b80636bc3af27146103fb57806373d4a13a146104a1578063806ba4d2146104a9578063893d20e8146104cf5780638da5cb5b146104f357610149565b80633c714980116101115780633c7149801461025b5780633d2d3a671461026357806347064d6a1461030957806350ca875f146103af5780636b8a98a9146103d557610149565b80630458a3351461014e578063053f6ef614610188578063324045de146101ae5780633b5e0ebd146101d65780633bc5de3014610253575b600080fd5b6101746004803603602081101561016457600080fd5b5035600160a060020a031661057d565b604080519115158252519081900360200190f35b6101746004803603602081101561019e57600080fd5b5035600160a060020a031661059b565b6101d4600480360360208110156101c457600080fd5b5035600160a060020a03166105b9565b005b6101de610630565b6040805160208082528351818301528351919283929083019185019080838360005b83811015610218578181015183820152602001610200565b50505050905090810190601f1680156102455780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101de6106be565b6101de610755565b6101d46004803603602081101561027957600080fd5b81019060208101813564010000000081111561029457600080fd5b8201836020820111156102a657600080fd5b803590602001918460018302840111640100000000831117156102c857600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506107b6945050505050565b6101d46004803603602081101561031f57600080fd5b81019060208101813564010000000081111561033a57600080fd5b82018360208201111561034c57600080fd5b8035906020019184600183028401116401000000008311171561036e57600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061081d945050505050565b610174600480360360208110156103c557600080fd5b5035600160a060020a0316610880565b6101d4600480360360208110156103eb57600080fd5b5035600160a060020a0316610895565b6101d46004803603602081101561041157600080fd5b81019060208101813564010000000081111561042c57600080fd5b82018360208201111561043e57600080fd5b8035906020019184600183028401116401000000008311171561046057600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610909945050505050565b6101de61096c565b610174600480360360208110156104bf57600080fd5b5035600160a060020a03166109c7565b6104d76109dc565b60408051600160a060020a039092168252519081900360200190f35b6104d76109eb565b6101de6109fa565b6101d46004803603602081101561051957600080fd5b5035600160a060020a0316610a5b565b6101d46004803603602081101561053f57600080fd5b5035600160a060020a0316610ada565b6101de610b6f565b6101d46004803603602081101561056d57600080fd5b5035600160a060020a0316610bca565b600160a060020a031660009081526020819052604090205460ff1690565b600160a060020a031660009081526001602052604090205460ff1690565b600254600160a060020a03163314610609576040805160e560020a62461bcd0281526020600482015260196024820152600080516020610cf5833981519152604482015290519081900360640190fd5b600160a060020a03166000908152600160208190526040909120805460ff19169091179055565b6005805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156106b65780601f1061068b576101008083540402835291602001916106b6565b820191906000526020600020905b81548152906001019060200180831161069957829003601f168201915b505050505081565b60048054604080516020601f600260001961010060018816150201909516949094049384018190048102820181019092528281526060939092909183018282801561074a5780601f1061071f5761010080835404028352916020019161074a565b820191906000526020600020905b81548152906001019060200180831161072d57829003601f168201915b505050505090505b90565b60038054604080516020601f600260001961010060018816150201909516949094049384018190048102820181019092528281526060939092909183018282801561074a5780601f1061071f5761010080835404028352916020019161074a565b600254600160a060020a03163314610806576040805160e560020a62461bcd0281526020600482015260196024820152600080516020610cf5833981519152604482015290519081900360640190fd5b8051610819906005906020840190610c5c565b5050565b600254600160a060020a0316331461086d576040805160e560020a62461bcd0281526020600482015260196024820152600080516020610cf5833981519152604482015290519081900360640190fd5b8051610819906004906020840190610c5c565b60006020819052908152604090205460ff1681565b600254600160a060020a031633146108e5576040805160e560020a62461bcd0281526020600482015260196024820152600080516020610cf5833981519152604482015290519081900360640190fd5b600160a060020a03166000908152602081905260409020805460ff19166001179055565b600254600160a060020a03163314610959576040805160e560020a62461bcd0281526020600482015260196024820152600080516020610cf5833981519152604482015290519081900360640190fd5b8051610819906003906020840190610c5c565b6004805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156106b65780601f1061068b576101008083540402835291602001916106b6565b60016020526000908152604090205460ff1681565b600254600160a060020a031690565b600254600160a060020a031681565b60058054604080516020601f600260001961010060018816150201909516949094049384018190048102820181019092528281526060939092909183018282801561074a5780601f1061071f5761010080835404028352916020019161074a565b600254600160a060020a03163314610aab576040805160e560020a62461bcd0281526020600482015260196024820152600080516020610cf5833981519152604482015290519081900360640190fd5b6002805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600254600160a060020a03163314610b2a576040805160e560020a62461bcd0281526020600482015260196024820152600080516020610cf5833981519152604482015290519081900360640190fd5b600160a060020a03811660009081526020819052604090205460ff1615610b6c57600160a060020a0381166000908152602081905260409020805460ff191690555b50565b6003805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156106b65780601f1061068b576101008083540402835291602001916106b6565b600254600160a060020a03163314610c1a576040805160e560020a62461bcd0281526020600482015260196024820152600080516020610cf5833981519152604482015290519081900360640190fd5b600160a060020a03811660009081526001602052604090205460ff1615610b6c57600160a060020a03166000908152600160205260409020805460ff19169055565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610c9d57805160ff1916838001178555610cca565b82800160010185558215610cca579182015b82811115610cca578251825591602001919060010190610caf565b50610cd6929150610cda565b5090565b61075291905b80821115610cd65760008155600101610ce056fe4d757374206265206f776e6572206f6620636f6e747261637400000000000000a165627a7a723058206eec045e9fb252fed8d555d5d4ddbb028efc5f112606682baf58248b8e93e7440029`

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
const PoolFactoryBin = `0x608060405234801561001057600080fd5b50611219806100206000396000f3fe608060405234801561001057600080fd5b506004361061009a576000357c0100000000000000000000000000000000000000000000000000000000900480637cf25517116100785780637cf255171461017a5780639049f9d2146101a0578063d4cced9a146101c6578063d88ff1f4146101ec5761009a565b806312708d721461009f57806341d1de97146100e7578063556f77db14610104575b600080fd5b6100cb600480360360408110156100b557600080fd5b50600160a060020a0381351690602001356101f4565b60408051600160a060020a039092168252519081900360200190f35b6100cb600480360360208110156100fd57600080fd5b503561022b565b61012a6004803603602081101561011a57600080fd5b5035600160a060020a0316610253565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561016657818101518382015260200161014e565b505050509050019250505060405180910390f35b6100cb6004803603602081101561019057600080fd5b5035600160a060020a03166102c7565b6100cb600480360360208110156101b657600080fd5b5035600160a060020a03166102e5565b6100cb600480360360208110156101dc57600080fd5b5035600160a060020a03166103c0565b61012a6103db565b60006020528160005260406000208181548110151561020f57fe5b600091825260209091200154600160a060020a03169150829050565b600280548290811061023957fe5b600091825260209091200154600160a060020a0316905081565b600160a060020a038116600090815260208181526040918290208054835181840281018401909452808452606093928301828280156102bb57602002820191906000526020600020905b8154600160a060020a0316815260019091019060200180831161029d575b50505050509050919050565b600160a060020a039081166000908152600160205260409020541690565b600080826102f161043d565b600160a060020a03909116815260405190819003602001906000f08015801561031e573d6000803e3d6000fd5b50600160a060020a0380851660008181526020818152604080832080546001818101835591855283852001805496881673ffffffffffffffffffffffffffffffffffffffff19978816811790915580855292819052908320805486169094179093556002805493840181559091527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace9091018054909216179055915050919050565b600160205260009081526040902054600160a060020a031681565b6060600280548060200260200160405190810160405280929190818152602001828054801561043357602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311610415575b5050505050905090565b604051610da08061044e8339019056fe608060405234801561001057600080fd5b50604051602080610da08339810180604052602081101561003057600080fd5b505160028054600160a060020a031916600160a060020a03909216919091179055610d40806100606000396000f3fe608060405234801561001057600080fd5b5060043610610149576000357c0100000000000000000000000000000000000000000000000000000000900480636bc3af27116100ca578063914964771161008e57806391496477146104fb578063a6f9dae114610503578063dd9cf97114610529578063e0a18ee71461054f578063ed82836e1461055757610149565b80636bc3af27146103fb57806373d4a13a146104a1578063806ba4d2146104a9578063893d20e8146104cf5780638da5cb5b146104f357610149565b80633c714980116101115780633c7149801461025b5780633d2d3a671461026357806347064d6a1461030957806350ca875f146103af5780636b8a98a9146103d557610149565b80630458a3351461014e578063053f6ef614610188578063324045de146101ae5780633b5e0ebd146101d65780633bc5de3014610253575b600080fd5b6101746004803603602081101561016457600080fd5b5035600160a060020a031661057d565b604080519115158252519081900360200190f35b6101746004803603602081101561019e57600080fd5b5035600160a060020a031661059b565b6101d4600480360360208110156101c457600080fd5b5035600160a060020a03166105b9565b005b6101de610630565b6040805160208082528351818301528351919283929083019185019080838360005b83811015610218578181015183820152602001610200565b50505050905090810190601f1680156102455780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101de6106be565b6101de610755565b6101d46004803603602081101561027957600080fd5b81019060208101813564010000000081111561029457600080fd5b8201836020820111156102a657600080fd5b803590602001918460018302840111640100000000831117156102c857600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506107b6945050505050565b6101d46004803603602081101561031f57600080fd5b81019060208101813564010000000081111561033a57600080fd5b82018360208201111561034c57600080fd5b8035906020019184600183028401116401000000008311171561036e57600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061081d945050505050565b610174600480360360208110156103c557600080fd5b5035600160a060020a0316610880565b6101d4600480360360208110156103eb57600080fd5b5035600160a060020a0316610895565b6101d46004803603602081101561041157600080fd5b81019060208101813564010000000081111561042c57600080fd5b82018360208201111561043e57600080fd5b8035906020019184600183028401116401000000008311171561046057600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610909945050505050565b6101de61096c565b610174600480360360208110156104bf57600080fd5b5035600160a060020a03166109c7565b6104d76109dc565b60408051600160a060020a039092168252519081900360200190f35b6104d76109eb565b6101de6109fa565b6101d46004803603602081101561051957600080fd5b5035600160a060020a0316610a5b565b6101d46004803603602081101561053f57600080fd5b5035600160a060020a0316610ada565b6101de610b6f565b6101d46004803603602081101561056d57600080fd5b5035600160a060020a0316610bca565b600160a060020a031660009081526020819052604090205460ff1690565b600160a060020a031660009081526001602052604090205460ff1690565b600254600160a060020a03163314610609576040805160e560020a62461bcd0281526020600482015260196024820152600080516020610cf5833981519152604482015290519081900360640190fd5b600160a060020a03166000908152600160208190526040909120805460ff19169091179055565b6005805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156106b65780601f1061068b576101008083540402835291602001916106b6565b820191906000526020600020905b81548152906001019060200180831161069957829003601f168201915b505050505081565b60048054604080516020601f600260001961010060018816150201909516949094049384018190048102820181019092528281526060939092909183018282801561074a5780601f1061071f5761010080835404028352916020019161074a565b820191906000526020600020905b81548152906001019060200180831161072d57829003601f168201915b505050505090505b90565b60038054604080516020601f600260001961010060018816150201909516949094049384018190048102820181019092528281526060939092909183018282801561074a5780601f1061071f5761010080835404028352916020019161074a565b600254600160a060020a03163314610806576040805160e560020a62461bcd0281526020600482015260196024820152600080516020610cf5833981519152604482015290519081900360640190fd5b8051610819906005906020840190610c5c565b5050565b600254600160a060020a0316331461086d576040805160e560020a62461bcd0281526020600482015260196024820152600080516020610cf5833981519152604482015290519081900360640190fd5b8051610819906004906020840190610c5c565b60006020819052908152604090205460ff1681565b600254600160a060020a031633146108e5576040805160e560020a62461bcd0281526020600482015260196024820152600080516020610cf5833981519152604482015290519081900360640190fd5b600160a060020a03166000908152602081905260409020805460ff19166001179055565b600254600160a060020a03163314610959576040805160e560020a62461bcd0281526020600482015260196024820152600080516020610cf5833981519152604482015290519081900360640190fd5b8051610819906003906020840190610c5c565b6004805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156106b65780601f1061068b576101008083540402835291602001916106b6565b60016020526000908152604090205460ff1681565b600254600160a060020a031690565b600254600160a060020a031681565b60058054604080516020601f600260001961010060018816150201909516949094049384018190048102820181019092528281526060939092909183018282801561074a5780601f1061071f5761010080835404028352916020019161074a565b600254600160a060020a03163314610aab576040805160e560020a62461bcd0281526020600482015260196024820152600080516020610cf5833981519152604482015290519081900360640190fd5b6002805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600254600160a060020a03163314610b2a576040805160e560020a62461bcd0281526020600482015260196024820152600080516020610cf5833981519152604482015290519081900360640190fd5b600160a060020a03811660009081526020819052604090205460ff1615610b6c57600160a060020a0381166000908152602081905260409020805460ff191690555b50565b6003805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156106b65780601f1061068b576101008083540402835291602001916106b6565b600254600160a060020a03163314610c1a576040805160e560020a62461bcd0281526020600482015260196024820152600080516020610cf5833981519152604482015290519081900360640190fd5b600160a060020a03811660009081526001602052604090205460ff1615610b6c57600160a060020a03166000908152600160205260409020805460ff19169055565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610c9d57805160ff1916838001178555610cca565b82800160010185558215610cca579182015b82811115610cca578251825591602001919060010190610caf565b50610cd6929150610cda565b5090565b61075291905b80821115610cd65760008155600101610ce056fe4d757374206265206f776e6572206f6620636f6e747261637400000000000000a165627a7a723058206eec045e9fb252fed8d555d5d4ddbb028efc5f112606682baf58248b8e93e7440029a165627a7a72305820268935df9938b5a438019d6611a57198866ad88f3f7a8270c348f57861ea45140029`

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
