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

// DebtTokenMetaData contains all meta data concerning the DebtToken contract.
var DebtTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidParameter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAuthorizedMinter\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DebtBurned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DebtMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"MinterAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"minters\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"setMinter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523461038257610ae78038038061001981610386565b9283398101906060818303126103825780516001600160401b03811161038257826100459183016103ab565b60208201519092906001600160401b038111610382576040916100699184016103ab565b9101516001600160a01b0381169081900361038257801561036f575f80546001600160a01b03198116831782556001600160a01b0316907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09080a381516001600160401b03811161028257600154600181811c91168015610365575b602082101461026457601f8111610302575b50602092601f82116001146102a157928192935f92610296575b50508160011b915f199060031b1c1916176001555b80516001600160401b03811161028257600254600181811c91168015610278575b602082101461026457601f8111610201575b50602091601f82116001146101a1579181925f92610196575b50508160011b915f199060031b1c1916176002555b6040516106ea90816103fd8239f35b015190505f80610172565b601f1982169260025f52805f20915f5b8581106101e9575083600195106101d1575b505050811b01600255610187565b01515f1960f88460031b161c191690555f80806101c3565b919260206001819286850151815501940192016101b1565b60025f527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace601f830160051c8101916020841061025a575b601f0160051c01905b81811061024f5750610159565b5f8155600101610242565b9091508190610239565b634e487b7160e01b5f52602260045260245ffd5b90607f1690610147565b634e487b7160e01b5f52604160045260245ffd5b015190505f80610111565b601f1982169360015f52805f20915f5b8681106102ea57508360019596106102d2575b505050811b01600155610126565b01515f1960f88460031b161c191690555f80806102c4565b919260206001819286850151815501940192016102b1565b60015f527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf6601f830160051c8101916020841061035b575b601f0160051c01905b81811061035057506100f7565b5f8155600101610343565b909150819061033a565b90607f16906100e5565b631e4fbdf760e01b5f525f60045260245ffd5b5f80fd5b6040519190601f01601f191682016001600160401b0381118382101761028257604052565b81601f82011215610382578051906001600160401b038211610282576103da601f8301601f1916602001610386565b928284526020838301011161038257815f9260208093018386015e830101529056fe60806040526004361015610011575f80fd5b5f3560e01c806306fdde031461054b57806318160ddd1461052e578063313ce5671461051357806340c10f191461047457806370a082311461043c578063715018a6146103e55780638da5cb5b146103be57806395d89b41146102a05780639dc29fac146101e0578063cf456ae714610163578063f2fde38b146100de5763f46eccc41461009d575f80fd5b346100da5760203660031901126100da576001600160a01b036100be61064a565b165f526003602052602060ff60405f2054166040519015158152f35b5f80fd5b346100da5760203660031901126100da576100f761064a565b6100ff61068e565b6001600160a01b03168015610150575f80546001600160a01b03198116831782556001600160a01b0316907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09080a3005b631e4fbdf760e01b5f525f60045260245ffd5b346100da5760403660031901126100da5761017c61064a565b602435908115158092036100da5760207f763efcc94f241a365ee1267a4046c4e650be372dd27a6948d4a23e224a26ebe3916101b661068e565b60018060a01b031692835f526003825260405f2060ff1981541660ff8316179055604051908152a2005b346100da5760403660031901126100da576101f961064a565b60243590335f52600360205260ff60405f205416156102915760018060a01b031690815f5260046020528060405f2054106102825760207f97a507cdb275af68c600db8c00bdcae6ed23981cc9f27fa4f2b86a5bdc076eb891835f526004825260405f20610268828254610681565b905561027681600554610681565b600555604051908152a2005b631e9acf1760e31b5f5260045ffd5b63019a5d1d60e71b5f5260045ffd5b346100da575f3660031901126100da576040515f6002548060011c906001811680156103b4575b6020831081146103a057828552908115610384575060011461032e575b50819003601f01601f191681019067ffffffffffffffff82118183101761031a57604082905281906103169082610620565b0390f35b634e487b7160e01b5f52604160045260245ffd5b60025f9081529091507f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace5b82821061036e575060209150820101826102e4565b6001816020925483858801015201910190610359565b90506020925060ff191682840152151560051b820101826102e4565b634e487b7160e01b5f52602260045260245ffd5b91607f16916102c7565b346100da575f3660031901126100da575f546040516001600160a01b039091168152602090f35b346100da575f3660031901126100da576103fd61068e565b5f80546001600160a01b0319811682556001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a3005b346100da5760203660031901126100da576001600160a01b0361045d61064a565b165f526004602052602060405f2054604051908152f35b346100da5760403660031901126100da5761048d61064a565b60243590335f52600360205260ff60405f20541615610291576001600160a01b03169081156105045780156105045760207f9b147abdd144aea38aa2d6db5c7851352d6de64a7e633d19f87b03a79febf1c391835f526004825260405f206104f6828254610660565b905561027681600554610660565b630309cb8760e51b5f5260045ffd5b346100da575f3660031901126100da57602060405160128152f35b346100da575f3660031901126100da576020600554604051908152f35b346100da575f3660031901126100da576040515f6001548060011c90600181168015610616575b6020831081146103a05782855290811561038457506001146105c05750819003601f01601f191681019067ffffffffffffffff82118183101761031a57604082905281906103169082610620565b60015f9081529091507fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf65b828210610600575060209150820101826102e4565b60018160209254838588010152019101906105eb565b91607f1691610572565b602060409281835280519182918282860152018484015e5f828201840152601f01601f1916010190565b600435906001600160a01b03821682036100da57565b9190820180921161066d57565b634e487b7160e01b5f52601160045260245ffd5b9190820391821161066d57565b5f546001600160a01b031633036106a157565b63118cdaa760e01b5f523360045260245ffdfea2646970667358221220921ad9a5166695792b6fa74e2c0fd5e2e061d7f309caecb40eb431fe54d6bae464736f6c634300081c0033",
}

// DebtTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use DebtTokenMetaData.ABI instead.
var DebtTokenABI = DebtTokenMetaData.ABI

// DebtTokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DebtTokenMetaData.Bin instead.
var DebtTokenBin = DebtTokenMetaData.Bin

// DeployDebtToken deploys a new Ethereum contract, binding an instance of DebtToken to it.
func DeployDebtToken(auth *bind.TransactOpts, backend bind.ContractBackend, _name string, _symbol string, _owner common.Address) (common.Address, *types.Transaction, *DebtToken, error) {
	parsed, err := DebtTokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DebtTokenBin), backend, _name, _symbol, _owner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DebtToken{DebtTokenCaller: DebtTokenCaller{contract: contract}, DebtTokenTransactor: DebtTokenTransactor{contract: contract}, DebtTokenFilterer: DebtTokenFilterer{contract: contract}}, nil
}

// DebtToken is an auto generated Go binding around an Ethereum contract.
type DebtToken struct {
	DebtTokenCaller     // Read-only binding to the contract
	DebtTokenTransactor // Write-only binding to the contract
	DebtTokenFilterer   // Log filterer for contract events
}

// DebtTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type DebtTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DebtTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DebtTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DebtTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DebtTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DebtTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DebtTokenSession struct {
	Contract     *DebtToken        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DebtTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DebtTokenCallerSession struct {
	Contract *DebtTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// DebtTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DebtTokenTransactorSession struct {
	Contract     *DebtTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// DebtTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type DebtTokenRaw struct {
	Contract *DebtToken // Generic contract binding to access the raw methods on
}

// DebtTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DebtTokenCallerRaw struct {
	Contract *DebtTokenCaller // Generic read-only contract binding to access the raw methods on
}

// DebtTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DebtTokenTransactorRaw struct {
	Contract *DebtTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDebtToken creates a new instance of DebtToken, bound to a specific deployed contract.
func NewDebtToken(address common.Address, backend bind.ContractBackend) (*DebtToken, error) {
	contract, err := bindDebtToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DebtToken{DebtTokenCaller: DebtTokenCaller{contract: contract}, DebtTokenTransactor: DebtTokenTransactor{contract: contract}, DebtTokenFilterer: DebtTokenFilterer{contract: contract}}, nil
}

// NewDebtTokenCaller creates a new read-only instance of DebtToken, bound to a specific deployed contract.
func NewDebtTokenCaller(address common.Address, caller bind.ContractCaller) (*DebtTokenCaller, error) {
	contract, err := bindDebtToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DebtTokenCaller{contract: contract}, nil
}

// NewDebtTokenTransactor creates a new write-only instance of DebtToken, bound to a specific deployed contract.
func NewDebtTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*DebtTokenTransactor, error) {
	contract, err := bindDebtToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DebtTokenTransactor{contract: contract}, nil
}

// NewDebtTokenFilterer creates a new log filterer instance of DebtToken, bound to a specific deployed contract.
func NewDebtTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*DebtTokenFilterer, error) {
	contract, err := bindDebtToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DebtTokenFilterer{contract: contract}, nil
}

// bindDebtToken binds a generic wrapper to an already deployed contract.
func bindDebtToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DebtTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DebtToken *DebtTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DebtToken.Contract.DebtTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DebtToken *DebtTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DebtToken.Contract.DebtTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DebtToken *DebtTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DebtToken.Contract.DebtTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DebtToken *DebtTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DebtToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DebtToken *DebtTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DebtToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DebtToken *DebtTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DebtToken.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_DebtToken *DebtTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DebtToken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_DebtToken *DebtTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _DebtToken.Contract.BalanceOf(&_DebtToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_DebtToken *DebtTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _DebtToken.Contract.BalanceOf(&_DebtToken.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_DebtToken *DebtTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _DebtToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_DebtToken *DebtTokenSession) Decimals() (uint8, error) {
	return _DebtToken.Contract.Decimals(&_DebtToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_DebtToken *DebtTokenCallerSession) Decimals() (uint8, error) {
	return _DebtToken.Contract.Decimals(&_DebtToken.CallOpts)
}

// Minters is a free data retrieval call binding the contract method 0xf46eccc4.
//
// Solidity: function minters(address ) view returns(bool)
func (_DebtToken *DebtTokenCaller) Minters(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _DebtToken.contract.Call(opts, &out, "minters", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Minters is a free data retrieval call binding the contract method 0xf46eccc4.
//
// Solidity: function minters(address ) view returns(bool)
func (_DebtToken *DebtTokenSession) Minters(arg0 common.Address) (bool, error) {
	return _DebtToken.Contract.Minters(&_DebtToken.CallOpts, arg0)
}

// Minters is a free data retrieval call binding the contract method 0xf46eccc4.
//
// Solidity: function minters(address ) view returns(bool)
func (_DebtToken *DebtTokenCallerSession) Minters(arg0 common.Address) (bool, error) {
	return _DebtToken.Contract.Minters(&_DebtToken.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DebtToken *DebtTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DebtToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DebtToken *DebtTokenSession) Name() (string, error) {
	return _DebtToken.Contract.Name(&_DebtToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DebtToken *DebtTokenCallerSession) Name() (string, error) {
	return _DebtToken.Contract.Name(&_DebtToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DebtToken *DebtTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DebtToken.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DebtToken *DebtTokenSession) Owner() (common.Address, error) {
	return _DebtToken.Contract.Owner(&_DebtToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DebtToken *DebtTokenCallerSession) Owner() (common.Address, error) {
	return _DebtToken.Contract.Owner(&_DebtToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_DebtToken *DebtTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DebtToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_DebtToken *DebtTokenSession) Symbol() (string, error) {
	return _DebtToken.Contract.Symbol(&_DebtToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_DebtToken *DebtTokenCallerSession) Symbol() (string, error) {
	return _DebtToken.Contract.Symbol(&_DebtToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_DebtToken *DebtTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DebtToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_DebtToken *DebtTokenSession) TotalSupply() (*big.Int, error) {
	return _DebtToken.Contract.TotalSupply(&_DebtToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_DebtToken *DebtTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _DebtToken.Contract.TotalSupply(&_DebtToken.CallOpts)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 amount) returns()
func (_DebtToken *DebtTokenTransactor) Burn(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DebtToken.contract.Transact(opts, "burn", account, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 amount) returns()
func (_DebtToken *DebtTokenSession) Burn(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DebtToken.Contract.Burn(&_DebtToken.TransactOpts, account, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 amount) returns()
func (_DebtToken *DebtTokenTransactorSession) Burn(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DebtToken.Contract.Burn(&_DebtToken.TransactOpts, account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_DebtToken *DebtTokenTransactor) Mint(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DebtToken.contract.Transact(opts, "mint", account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_DebtToken *DebtTokenSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DebtToken.Contract.Mint(&_DebtToken.TransactOpts, account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_DebtToken *DebtTokenTransactorSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DebtToken.Contract.Mint(&_DebtToken.TransactOpts, account, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DebtToken *DebtTokenTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DebtToken.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DebtToken *DebtTokenSession) RenounceOwnership() (*types.Transaction, error) {
	return _DebtToken.Contract.RenounceOwnership(&_DebtToken.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DebtToken *DebtTokenTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _DebtToken.Contract.RenounceOwnership(&_DebtToken.TransactOpts)
}

// SetMinter is a paid mutator transaction binding the contract method 0xcf456ae7.
//
// Solidity: function setMinter(address minter, bool status) returns()
func (_DebtToken *DebtTokenTransactor) SetMinter(opts *bind.TransactOpts, minter common.Address, status bool) (*types.Transaction, error) {
	return _DebtToken.contract.Transact(opts, "setMinter", minter, status)
}

// SetMinter is a paid mutator transaction binding the contract method 0xcf456ae7.
//
// Solidity: function setMinter(address minter, bool status) returns()
func (_DebtToken *DebtTokenSession) SetMinter(minter common.Address, status bool) (*types.Transaction, error) {
	return _DebtToken.Contract.SetMinter(&_DebtToken.TransactOpts, minter, status)
}

// SetMinter is a paid mutator transaction binding the contract method 0xcf456ae7.
//
// Solidity: function setMinter(address minter, bool status) returns()
func (_DebtToken *DebtTokenTransactorSession) SetMinter(minter common.Address, status bool) (*types.Transaction, error) {
	return _DebtToken.Contract.SetMinter(&_DebtToken.TransactOpts, minter, status)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DebtToken *DebtTokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _DebtToken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DebtToken *DebtTokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DebtToken.Contract.TransferOwnership(&_DebtToken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DebtToken *DebtTokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DebtToken.Contract.TransferOwnership(&_DebtToken.TransactOpts, newOwner)
}

// DebtTokenDebtBurnedIterator is returned from FilterDebtBurned and is used to iterate over the raw logs and unpacked data for DebtBurned events raised by the DebtToken contract.
type DebtTokenDebtBurnedIterator struct {
	Event *DebtTokenDebtBurned // Event containing the contract specifics and raw log

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
func (it *DebtTokenDebtBurnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DebtTokenDebtBurned)
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
		it.Event = new(DebtTokenDebtBurned)
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
func (it *DebtTokenDebtBurnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DebtTokenDebtBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DebtTokenDebtBurned represents a DebtBurned event raised by the DebtToken contract.
type DebtTokenDebtBurned struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDebtBurned is a free log retrieval operation binding the contract event 0x97a507cdb275af68c600db8c00bdcae6ed23981cc9f27fa4f2b86a5bdc076eb8.
//
// Solidity: event DebtBurned(address indexed account, uint256 amount)
func (_DebtToken *DebtTokenFilterer) FilterDebtBurned(opts *bind.FilterOpts, account []common.Address) (*DebtTokenDebtBurnedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _DebtToken.contract.FilterLogs(opts, "DebtBurned", accountRule)
	if err != nil {
		return nil, err
	}
	return &DebtTokenDebtBurnedIterator{contract: _DebtToken.contract, event: "DebtBurned", logs: logs, sub: sub}, nil
}

// WatchDebtBurned is a free log subscription operation binding the contract event 0x97a507cdb275af68c600db8c00bdcae6ed23981cc9f27fa4f2b86a5bdc076eb8.
//
// Solidity: event DebtBurned(address indexed account, uint256 amount)
func (_DebtToken *DebtTokenFilterer) WatchDebtBurned(opts *bind.WatchOpts, sink chan<- *DebtTokenDebtBurned, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _DebtToken.contract.WatchLogs(opts, "DebtBurned", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DebtTokenDebtBurned)
				if err := _DebtToken.contract.UnpackLog(event, "DebtBurned", log); err != nil {
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

// ParseDebtBurned is a log parse operation binding the contract event 0x97a507cdb275af68c600db8c00bdcae6ed23981cc9f27fa4f2b86a5bdc076eb8.
//
// Solidity: event DebtBurned(address indexed account, uint256 amount)
func (_DebtToken *DebtTokenFilterer) ParseDebtBurned(log types.Log) (*DebtTokenDebtBurned, error) {
	event := new(DebtTokenDebtBurned)
	if err := _DebtToken.contract.UnpackLog(event, "DebtBurned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DebtTokenDebtMintedIterator is returned from FilterDebtMinted and is used to iterate over the raw logs and unpacked data for DebtMinted events raised by the DebtToken contract.
type DebtTokenDebtMintedIterator struct {
	Event *DebtTokenDebtMinted // Event containing the contract specifics and raw log

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
func (it *DebtTokenDebtMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DebtTokenDebtMinted)
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
		it.Event = new(DebtTokenDebtMinted)
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
func (it *DebtTokenDebtMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DebtTokenDebtMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DebtTokenDebtMinted represents a DebtMinted event raised by the DebtToken contract.
type DebtTokenDebtMinted struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDebtMinted is a free log retrieval operation binding the contract event 0x9b147abdd144aea38aa2d6db5c7851352d6de64a7e633d19f87b03a79febf1c3.
//
// Solidity: event DebtMinted(address indexed account, uint256 amount)
func (_DebtToken *DebtTokenFilterer) FilterDebtMinted(opts *bind.FilterOpts, account []common.Address) (*DebtTokenDebtMintedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _DebtToken.contract.FilterLogs(opts, "DebtMinted", accountRule)
	if err != nil {
		return nil, err
	}
	return &DebtTokenDebtMintedIterator{contract: _DebtToken.contract, event: "DebtMinted", logs: logs, sub: sub}, nil
}

// WatchDebtMinted is a free log subscription operation binding the contract event 0x9b147abdd144aea38aa2d6db5c7851352d6de64a7e633d19f87b03a79febf1c3.
//
// Solidity: event DebtMinted(address indexed account, uint256 amount)
func (_DebtToken *DebtTokenFilterer) WatchDebtMinted(opts *bind.WatchOpts, sink chan<- *DebtTokenDebtMinted, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _DebtToken.contract.WatchLogs(opts, "DebtMinted", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DebtTokenDebtMinted)
				if err := _DebtToken.contract.UnpackLog(event, "DebtMinted", log); err != nil {
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

// ParseDebtMinted is a log parse operation binding the contract event 0x9b147abdd144aea38aa2d6db5c7851352d6de64a7e633d19f87b03a79febf1c3.
//
// Solidity: event DebtMinted(address indexed account, uint256 amount)
func (_DebtToken *DebtTokenFilterer) ParseDebtMinted(log types.Log) (*DebtTokenDebtMinted, error) {
	event := new(DebtTokenDebtMinted)
	if err := _DebtToken.contract.UnpackLog(event, "DebtMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DebtTokenMinterAddedIterator is returned from FilterMinterAdded and is used to iterate over the raw logs and unpacked data for MinterAdded events raised by the DebtToken contract.
type DebtTokenMinterAddedIterator struct {
	Event *DebtTokenMinterAdded // Event containing the contract specifics and raw log

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
func (it *DebtTokenMinterAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DebtTokenMinterAdded)
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
		it.Event = new(DebtTokenMinterAdded)
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
func (it *DebtTokenMinterAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DebtTokenMinterAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DebtTokenMinterAdded represents a MinterAdded event raised by the DebtToken contract.
type DebtTokenMinterAdded struct {
	Minter common.Address
	Status bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMinterAdded is a free log retrieval operation binding the contract event 0x763efcc94f241a365ee1267a4046c4e650be372dd27a6948d4a23e224a26ebe3.
//
// Solidity: event MinterAdded(address indexed minter, bool status)
func (_DebtToken *DebtTokenFilterer) FilterMinterAdded(opts *bind.FilterOpts, minter []common.Address) (*DebtTokenMinterAddedIterator, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}

	logs, sub, err := _DebtToken.contract.FilterLogs(opts, "MinterAdded", minterRule)
	if err != nil {
		return nil, err
	}
	return &DebtTokenMinterAddedIterator{contract: _DebtToken.contract, event: "MinterAdded", logs: logs, sub: sub}, nil
}

// WatchMinterAdded is a free log subscription operation binding the contract event 0x763efcc94f241a365ee1267a4046c4e650be372dd27a6948d4a23e224a26ebe3.
//
// Solidity: event MinterAdded(address indexed minter, bool status)
func (_DebtToken *DebtTokenFilterer) WatchMinterAdded(opts *bind.WatchOpts, sink chan<- *DebtTokenMinterAdded, minter []common.Address) (event.Subscription, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}

	logs, sub, err := _DebtToken.contract.WatchLogs(opts, "MinterAdded", minterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DebtTokenMinterAdded)
				if err := _DebtToken.contract.UnpackLog(event, "MinterAdded", log); err != nil {
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

// ParseMinterAdded is a log parse operation binding the contract event 0x763efcc94f241a365ee1267a4046c4e650be372dd27a6948d4a23e224a26ebe3.
//
// Solidity: event MinterAdded(address indexed minter, bool status)
func (_DebtToken *DebtTokenFilterer) ParseMinterAdded(log types.Log) (*DebtTokenMinterAdded, error) {
	event := new(DebtTokenMinterAdded)
	if err := _DebtToken.contract.UnpackLog(event, "MinterAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DebtTokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the DebtToken contract.
type DebtTokenOwnershipTransferredIterator struct {
	Event *DebtTokenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DebtTokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DebtTokenOwnershipTransferred)
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
		it.Event = new(DebtTokenOwnershipTransferred)
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
func (it *DebtTokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DebtTokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DebtTokenOwnershipTransferred represents a OwnershipTransferred event raised by the DebtToken contract.
type DebtTokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DebtToken *DebtTokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DebtTokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DebtToken.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DebtTokenOwnershipTransferredIterator{contract: _DebtToken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DebtToken *DebtTokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DebtTokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DebtToken.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DebtTokenOwnershipTransferred)
				if err := _DebtToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_DebtToken *DebtTokenFilterer) ParseOwnershipTransferred(log types.Log) (*DebtTokenOwnershipTransferred, error) {
	event := new(DebtTokenOwnershipTransferred)
	if err := _DebtToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
