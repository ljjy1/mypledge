// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindcode

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

// MockTestERC20MetaData contains all meta data concerning the MockTestERC20 contract.
var MockTestERC20MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"initialSupply\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523461039557610a178038038061001981610399565b9283398101906060818303126103955780516001600160401b03811161039557826100459183016103be565b60208201519092906001600160401b038111610395576040916100699184016103be565b91015182516001600160401b0381116102a657600354600181811c9116801561038b575b602082101461028857601f8111610328575b506020601f82116001146102c557819293945f926102ba575b50508160011b915f199060031b1c1916176003555b81516001600160401b0381116102a657600454600181811c9116801561029c575b602082101461028857601f8111610225575b50602092601f82116001146101c457928192935f926101b9575b50508160011b915f199060031b1c1916176004555b33156101a65760025481810180911161019257600255335f525f60205260405f208181540190556040519081525f7fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60203393a360405161060790816104108239f35b634e487b7160e01b5f52601160045260245ffd5b63ec442f0560e01b5f525f60045260245ffd5b015190505f8061011a565b601f1982169360045f52805f20915f5b86811061020d57508360019596106101f5575b505050811b0160045561012f565b01515f1960f88460031b161c191690555f80806101e7565b919260206001819286850151815501940192016101d4565b60045f527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b601f830160051c8101916020841061027e575b601f0160051c01905b8181106102735750610100565b5f8155600101610266565b909150819061025d565b634e487b7160e01b5f52602260045260245ffd5b90607f16906100ee565b634e487b7160e01b5f52604160045260245ffd5b015190505f806100b8565b601f1982169060035f52805f20915f5b818110610310575095836001959697106102f8575b505050811b016003556100cd565b01515f1960f88460031b161c191690555f80806102ea565b9192602060018192868b0151815501940192016102d5565b60035f527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b601f830160051c81019160208410610381575b601f0160051c01905b818110610376575061009f565b5f8155600101610369565b9091508190610360565b90607f169061008d565b5f80fd5b6040519190601f01601f191682016001600160401b038111838210176102a657604052565b81601f82011215610395578051906001600160401b0382116102a6576103ed601f8301601f1916602001610399565b928284526020838301011161039557815f9260208093018386015e830101529056fe6080806040526004361015610012575f80fd5b5f3560e01c90816306fdde03146103ef57508063095ea7b31461036d57806318160ddd1461035057806323b872dd14610271578063313ce5671461025657806370a082311461021f57806395d89b4114610104578063a9059cbb146100d35763dd62ed3e1461007f575f80fd5b346100cf5760403660031901126100cf576100986104e8565b6100a06104fe565b6001600160a01b039182165f908152600160209081526040808320949093168252928352819020549051908152f35b5f80fd5b346100cf5760403660031901126100cf576100f96100ef6104e8565b6024359033610514565b602060405160018152f35b346100cf575f3660031901126100cf576040515f6004548060011c90600181168015610215575b602083108114610201578285529081156101e55750600114610190575b50819003601f01601f191681019067ffffffffffffffff82118183101761017c57610178829182604052826104be565b0390f35b634e487b7160e01b5f52604160045260245ffd5b905060045f527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b5f905b8282106101cf57506020915082010182610148565b60018160209254838588010152019101906101ba565b90506020925060ff191682840152151560051b82010182610148565b634e487b7160e01b5f52602260045260245ffd5b91607f169161012b565b346100cf5760203660031901126100cf576001600160a01b036102406104e8565b165f525f602052602060405f2054604051908152f35b346100cf575f3660031901126100cf57602060405160128152f35b346100cf5760603660031901126100cf5761028a6104e8565b6102926104fe565b6001600160a01b0382165f818152600160209081526040808320338452909152902054909260443592915f1981106102d0575b506100f99350610514565b83811061033557841561032257331561030f576100f9945f52600160205260405f2060018060a01b0333165f526020528360405f2091039055846102c5565b634a1406b160e11b5f525f60045260245ffd5b63e602df0560e01b5f525f60045260245ffd5b8390637dc7a0d960e11b5f523360045260245260445260645ffd5b346100cf575f3660031901126100cf576020600254604051908152f35b346100cf5760403660031901126100cf576103866104e8565b602435903315610322576001600160a01b031690811561030f57335f52600160205260405f20825f526020528060405f20556040519081527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560203392a3602060405160018152f35b346100cf575f3660031901126100cf575f6003548060011c906001811680156104b4575b602083108114610201578285529081156101e5575060011461045f5750819003601f01601f191681019067ffffffffffffffff82118183101761017c57610178829182604052826104be565b905060035f527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b5f905b82821061049e57506020915082010182610148565b6001816020925483858801015201910190610489565b91607f1691610413565b602060409281835280519182918282860152018484015e5f828201840152601f01601f1916010190565b600435906001600160a01b03821682036100cf57565b602435906001600160a01b03821682036100cf57565b6001600160a01b03169081156105be576001600160a01b03169182156105ab57815f525f60205260405f205481811061059257817fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92602092855f525f84520360405f2055845f525f825260405f20818154019055604051908152a3565b8263391434e360e21b5f5260045260245260445260645ffd5b63ec442f0560e01b5f525f60045260245ffd5b634b637e8f60e11b5f525f60045260245ffdfea264697066735822122016b0077c0571bfe30f817ebc97445779fceefc6bbf3e655e15ed6841ccaf8edf64736f6c634300081c0033",
}

// MockTestERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use MockTestERC20MetaData.ABI instead.
var MockTestERC20ABI = MockTestERC20MetaData.ABI

// MockTestERC20Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MockTestERC20MetaData.Bin instead.
var MockTestERC20Bin = MockTestERC20MetaData.Bin

// DeployMockTestERC20 deploys a new Ethereum contract, binding an instance of MockTestERC20 to it.
func DeployMockTestERC20(auth *bind.TransactOpts, backend bind.ContractBackend, name string, symbol string, initialSupply *big.Int) (common.Address, *types.Transaction, *MockTestERC20, error) {
	parsed, err := MockTestERC20MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MockTestERC20Bin), backend, name, symbol, initialSupply)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MockTestERC20{MockTestERC20Caller: MockTestERC20Caller{contract: contract}, MockTestERC20Transactor: MockTestERC20Transactor{contract: contract}, MockTestERC20Filterer: MockTestERC20Filterer{contract: contract}}, nil
}

// MockTestERC20 is an auto generated Go binding around an Ethereum contract.
type MockTestERC20 struct {
	MockTestERC20Caller     // Read-only binding to the contract
	MockTestERC20Transactor // Write-only binding to the contract
	MockTestERC20Filterer   // Log filterer for contract events
}

// MockTestERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type MockTestERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockTestERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type MockTestERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockTestERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MockTestERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockTestERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MockTestERC20Session struct {
	Contract     *MockTestERC20    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MockTestERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MockTestERC20CallerSession struct {
	Contract *MockTestERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// MockTestERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MockTestERC20TransactorSession struct {
	Contract     *MockTestERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// MockTestERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type MockTestERC20Raw struct {
	Contract *MockTestERC20 // Generic contract binding to access the raw methods on
}

// MockTestERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MockTestERC20CallerRaw struct {
	Contract *MockTestERC20Caller // Generic read-only contract binding to access the raw methods on
}

// MockTestERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MockTestERC20TransactorRaw struct {
	Contract *MockTestERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewMockTestERC20 creates a new instance of MockTestERC20, bound to a specific deployed contract.
func NewMockTestERC20(address common.Address, backend bind.ContractBackend) (*MockTestERC20, error) {
	contract, err := bindMockTestERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MockTestERC20{MockTestERC20Caller: MockTestERC20Caller{contract: contract}, MockTestERC20Transactor: MockTestERC20Transactor{contract: contract}, MockTestERC20Filterer: MockTestERC20Filterer{contract: contract}}, nil
}

// NewMockTestERC20Caller creates a new read-only instance of MockTestERC20, bound to a specific deployed contract.
func NewMockTestERC20Caller(address common.Address, caller bind.ContractCaller) (*MockTestERC20Caller, error) {
	contract, err := bindMockTestERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MockTestERC20Caller{contract: contract}, nil
}

// NewMockTestERC20Transactor creates a new write-only instance of MockTestERC20, bound to a specific deployed contract.
func NewMockTestERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*MockTestERC20Transactor, error) {
	contract, err := bindMockTestERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MockTestERC20Transactor{contract: contract}, nil
}

// NewMockTestERC20Filterer creates a new log filterer instance of MockTestERC20, bound to a specific deployed contract.
func NewMockTestERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*MockTestERC20Filterer, error) {
	contract, err := bindMockTestERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MockTestERC20Filterer{contract: contract}, nil
}

// bindMockTestERC20 binds a generic wrapper to an already deployed contract.
func bindMockTestERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MockTestERC20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockTestERC20 *MockTestERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockTestERC20.Contract.MockTestERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockTestERC20 *MockTestERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockTestERC20.Contract.MockTestERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockTestERC20 *MockTestERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockTestERC20.Contract.MockTestERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockTestERC20 *MockTestERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockTestERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockTestERC20 *MockTestERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockTestERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockTestERC20 *MockTestERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockTestERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_MockTestERC20 *MockTestERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MockTestERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_MockTestERC20 *MockTestERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _MockTestERC20.Contract.Allowance(&_MockTestERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_MockTestERC20 *MockTestERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _MockTestERC20.Contract.Allowance(&_MockTestERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_MockTestERC20 *MockTestERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MockTestERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_MockTestERC20 *MockTestERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _MockTestERC20.Contract.BalanceOf(&_MockTestERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_MockTestERC20 *MockTestERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _MockTestERC20.Contract.BalanceOf(&_MockTestERC20.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MockTestERC20 *MockTestERC20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _MockTestERC20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MockTestERC20 *MockTestERC20Session) Decimals() (uint8, error) {
	return _MockTestERC20.Contract.Decimals(&_MockTestERC20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MockTestERC20 *MockTestERC20CallerSession) Decimals() (uint8, error) {
	return _MockTestERC20.Contract.Decimals(&_MockTestERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MockTestERC20 *MockTestERC20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MockTestERC20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MockTestERC20 *MockTestERC20Session) Name() (string, error) {
	return _MockTestERC20.Contract.Name(&_MockTestERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MockTestERC20 *MockTestERC20CallerSession) Name() (string, error) {
	return _MockTestERC20.Contract.Name(&_MockTestERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MockTestERC20 *MockTestERC20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MockTestERC20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MockTestERC20 *MockTestERC20Session) Symbol() (string, error) {
	return _MockTestERC20.Contract.Symbol(&_MockTestERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MockTestERC20 *MockTestERC20CallerSession) Symbol() (string, error) {
	return _MockTestERC20.Contract.Symbol(&_MockTestERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MockTestERC20 *MockTestERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MockTestERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MockTestERC20 *MockTestERC20Session) TotalSupply() (*big.Int, error) {
	return _MockTestERC20.Contract.TotalSupply(&_MockTestERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MockTestERC20 *MockTestERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _MockTestERC20.Contract.TotalSupply(&_MockTestERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_MockTestERC20 *MockTestERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _MockTestERC20.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_MockTestERC20 *MockTestERC20Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _MockTestERC20.Contract.Approve(&_MockTestERC20.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_MockTestERC20 *MockTestERC20TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _MockTestERC20.Contract.Approve(&_MockTestERC20.TransactOpts, spender, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_MockTestERC20 *MockTestERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MockTestERC20.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_MockTestERC20 *MockTestERC20Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MockTestERC20.Contract.Transfer(&_MockTestERC20.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_MockTestERC20 *MockTestERC20TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MockTestERC20.Contract.Transfer(&_MockTestERC20.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_MockTestERC20 *MockTestERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MockTestERC20.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_MockTestERC20 *MockTestERC20Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MockTestERC20.Contract.TransferFrom(&_MockTestERC20.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_MockTestERC20 *MockTestERC20TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MockTestERC20.Contract.TransferFrom(&_MockTestERC20.TransactOpts, from, to, value)
}

// MockTestERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the MockTestERC20 contract.
type MockTestERC20ApprovalIterator struct {
	Event *MockTestERC20Approval // Event containing the contract specifics and raw log

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
func (it *MockTestERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockTestERC20Approval)
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
		it.Event = new(MockTestERC20Approval)
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
func (it *MockTestERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockTestERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockTestERC20Approval represents a Approval event raised by the MockTestERC20 contract.
type MockTestERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_MockTestERC20 *MockTestERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*MockTestERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MockTestERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &MockTestERC20ApprovalIterator{contract: _MockTestERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_MockTestERC20 *MockTestERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MockTestERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MockTestERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockTestERC20Approval)
				if err := _MockTestERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_MockTestERC20 *MockTestERC20Filterer) ParseApproval(log types.Log) (*MockTestERC20Approval, error) {
	event := new(MockTestERC20Approval)
	if err := _MockTestERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MockTestERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the MockTestERC20 contract.
type MockTestERC20TransferIterator struct {
	Event *MockTestERC20Transfer // Event containing the contract specifics and raw log

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
func (it *MockTestERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockTestERC20Transfer)
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
		it.Event = new(MockTestERC20Transfer)
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
func (it *MockTestERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockTestERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockTestERC20Transfer represents a Transfer event raised by the MockTestERC20 contract.
type MockTestERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MockTestERC20 *MockTestERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MockTestERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MockTestERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MockTestERC20TransferIterator{contract: _MockTestERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MockTestERC20 *MockTestERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MockTestERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MockTestERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockTestERC20Transfer)
				if err := _MockTestERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MockTestERC20 *MockTestERC20Filterer) ParseTransfer(log types.Log) (*MockTestERC20Transfer, error) {
	event := new(MockTestERC20Transfer)
	if err := _MockTestERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
