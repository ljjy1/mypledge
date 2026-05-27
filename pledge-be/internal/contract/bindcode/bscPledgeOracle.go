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

// BscPledgeOracleMetaData contains all meta data concerning the BscPledgeOracle contract.
var BscPledgeOracleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"asset\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"tokenDecimals\",\"type\":\"uint8\"}],\"name\":\"AggregatorSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"asset\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"ManualPriceSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"assets\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"}],\"name\":\"ManualPricesSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldDivisor\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newDivisor\",\"type\":\"uint256\"}],\"name\":\"PriceDivisorUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getAssetsAggregator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"assets\",\"type\":\"uint256[]\"}],\"name\":\"getPrices\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"underlying\",\"type\":\"uint256\"}],\"name\":\"getUnderlyingAggregator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"underlying\",\"type\":\"uint256\"}],\"name\":\"getUnderlyingPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceDivisor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"tokenDecimals\",\"type\":\"uint8\"}],\"name\":\"setAssetsAggregator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"setPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newDivisor\",\"type\":\"uint256\"}],\"name\":\"setPriceDivisor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"assets\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"}],\"name\":\"setPrices\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"underlying\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"tokenDecimals\",\"type\":\"uint8\"}],\"name\":\"setUnderlyingAggregator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"underlying\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"setUnderlyingPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60803460bd57601f610da238819003918201601f19168301916001600160401b0383118484101760c15780849260209460405283398101031260bd57516001600160a01b0381169081900360bd57801560aa575f80546001600160a01b031981168317825560405192916001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09080a36001600455610ccc90816100d68239f35b631e4fbdf760e01b5f525f60045260245ffd5b5f80fd5b634e487b7160e01b5f52604160045260245ffdfe60806040526004361015610011575f80fd5b5f3560e01c8062e4768b1461068e57806309cb3a4e146105665780630b905a2e1461053157806341976e0914610505578063715018a6146104ae57806375e443aa1461046457806383532667146104085780638da5cb5b146103e1578063b2861e7b14610348578063b889a989146102f0578063be7b4adc146102ab578063d05eaae014610189578063d690a8c31461016c578063da663257146101465763f2fde38b146100bd575f80fd5b34610142576020366003190112610142576100d66106f9565b6100de610858565b6001600160a01b0316801561012f575f80546001600160a01b03198116831782556001600160a01b0316907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09080a3005b631e4fbdf760e01b5f525f60045260245ffd5b5f80fd5b34610142576020366003190112610142576020610164600435610810565b604051908152f35b34610142575f366003190112610142576020600454604051908152f35b346101425760403660031901126101425760043567ffffffffffffffff8111610142576101ba903690600401610783565b9060243567ffffffffffffffff8111610142576101db903690600401610783565b6101e6939193610858565b808203610274575f5b82811061024757507fff1a86ea42b3a30c49373419293ba99ec1535aaef99ebd6f4be0a307a54ad1d693610242916102346040519586956040875260408701916107ec565b9184830360208601526107ec565b0390a1005b8061025560019284886107dc565b356102618286886107dc565b355f52600360205260405f2055016101ef565b60405162461bcd60e51b815260206004820152600f60248201526e098cadccee8d040dad2e6dac2e8c6d608b1b6044820152606490fd5b34610142576060366003190112610142576102ee6102c76106f9565b6102cf61070f565b6102d7610773565b916102e0610858565b6001600160a01b031661087e565b005b34610142576020366003190112610142576001600160a01b036103116106f9565b165f9081526001602090815260408083205460028352928190205481516001600160a01b03909416845260ff169183019190915290f35b3461014257602036600319011261014257600435610364610858565b80156103a35760407fcd572cdad82b8d7c7cc3eccd904c997bf961bb9e917c48a2c1214b200fb3e36591600454908060045582519182526020820152a1005b60405162461bcd60e51b815260206004820152601660248201527544697669736f722063616e6e6f74206265207a65726f60501b6044820152606490fd5b34610142575f366003190112610142575f546040516001600160a01b039091168152602090f35b34610142576040366003190112610142577fb4e7056b10bc2d2d4fe5b2ff374b24edf94410f4620866b12b3bcf1765d91894604060043560243561044a610858565b815f52600360205280835f205582519182526020820152a1005b34610142576020366003190112610142576004355f9081526001602090815260408083205460028352928190205481516001600160a01b03909416845260ff169183019190915290f35b34610142575f366003190112610142576104c6610858565b5f80546001600160a01b0319811682556001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a3005b346101425760203660031901126101425760206101646001600160a01b0361052b6106f9565b16610810565b34610142576060366003190112610142576102ee61054d61070f565b610555610773565b9061055e610858565b60043561087e565b346101425760203660031901126101425760043567ffffffffffffffff81116101425736602382011215610142578060040135906105a38261075b565b916105b16040519384610725565b8083526024602084019160051b8301019136831161014257602401905b82821061067e578351846105fa6105e48361075b565b926105f26040519485610725565b80845261075b565b602083019190601f19013683375f5b81518110156106395780610628610622600193856107b4565b51610810565b61063282876107b4565b5201610609565b505090604051918291602083019060208452518091526040830191905f5b818110610665575050500390f35b8251845285945060209384019390920191600101610657565b81358152602091820191016105ce565b34610142576040366003190112610142577fb4e7056b10bc2d2d4fe5b2ff374b24edf94410f4620866b12b3bcf1765d9189460406106ca6106f9565b602435906106d6610858565b60018060a01b031690815f52600360205280835f205582519182526020820152a1005b600435906001600160a01b038216820361014257565b602435906001600160a01b038216820361014257565b90601f8019910116810190811067ffffffffffffffff82111761074757604052565b634e487b7160e01b5f52604160045260245ffd5b67ffffffffffffffff81116107475760051b60200190565b6044359060ff8216820361014257565b9181601f840112156101425782359167ffffffffffffffff8311610142576020808501948460051b01011161014257565b80518210156107c85760209160051b010190565b634e487b7160e01b5f52603260045260245ffd5b91908110156107c85760051b0190565b81835290916001600160fb1b0383116101425760209260051b809284830137010190565b5f818152600160205260409020546001600160a01b03168061083c57505f52600360205260405f205490565b610855915f52600260205260ff60405f20541690610a01565b90565b5f546001600160a01b0316330361086b57565b63118cdaa760e01b5f523360045260245ffd5b916001600160a01b0390911690811561095d5760ff169182151580610952575b15610914577f5fa5791e5634d03047f423c78a1370531e6d289f6caf006a94fa039a3b5dce9392606092825f52600160205260405f20816bffffffffffffffffffffffff60a01b825416179055825f52600260205260405f208260ff1982541617905560405192835260208301526040820152a1565b60405162461bcd60e51b8152602060048201526016602482015275496e76616c696420746f6b656e20646563696d616c7360501b6044820152606490fd5b50601e83111561089e565b60405162461bcd60e51b815260206004820152601260248201527124b73b30b634b21030b3b3b932b3b0ba37b960711b6044820152606490fd5b519069ffffffffffffffffffff8216820361014257565b604d81116109bc57600a0a90565b634e487b7160e01b5f52601160045260245ffd5b81156109da570490565b634e487b7160e01b5f52601260045260245ffd5b818102929181159184041417156109bc57565b604051633fabe5a360e21b81526001600160a01b03919091169060a081600481855afa908115610b91575f915f925f925f92610c40575b505f841315610c075769ffffffffffffffffffff809116911610610bd45715610b9c5760206004926040519384809263313ce56760e01b82525afa918215610b91575f92610b53575b506004549160ff16816012821015610b155750601203601281116109bc5760ff91610aae610ab4926109ae565b906109ee565b925b166012811015610ae057601203601281116109bc5761085592610aae610adb926109ae565b6109d0565b6012811115610b0b5760111981019081116109bc5761085592610b05610adb926109ae565b906109d0565b50610855916109d0565b9060128195939511610b2c575b505060ff90610ab6565b9193509060111981019081116109bc5760ff91610b05610b4b926109ae565b92905f610b22565b9091506020813d602011610b89575b81610b6f60209383610725565b81010312610142575160ff8116810361014257905f610a81565b3d9150610b62565b6040513d5f823e3d90fd5b60405162461bcd60e51b815260206004820152601060248201526f125b98dbdb5c1b195d19481c9bdd5b9960821b6044820152606490fd5b60405162461bcd60e51b815260206004820152600b60248201526a14dd185b19481c9bdd5b9960aa1b6044820152606490fd5b60405162461bcd60e51b815260206004820152601160248201527004f7261636c65207072696365203c3d203607c1b6044820152606490fd5b935050505060a0813d60a011610c8e575b81610c5e60a09383610725565b8101031261014257610c6f81610997565b6020820151610c85608060608501519401610997565b9092915f610a38565b3d9150610c5156fea2646970667358221220c979d02c6692527b14061ef7619e7610a6571c2577dd03ac59bc283b018b5e8b64736f6c634300081c0033",
}

// BscPledgeOracleABI is the input ABI used to generate the binding from.
// Deprecated: Use BscPledgeOracleMetaData.ABI instead.
var BscPledgeOracleABI = BscPledgeOracleMetaData.ABI

// BscPledgeOracleBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BscPledgeOracleMetaData.Bin instead.
var BscPledgeOracleBin = BscPledgeOracleMetaData.Bin

// DeployBscPledgeOracle deploys a new Ethereum contract, binding an instance of BscPledgeOracle to it.
func DeployBscPledgeOracle(auth *bind.TransactOpts, backend bind.ContractBackend, _owner common.Address) (common.Address, *types.Transaction, *BscPledgeOracle, error) {
	parsed, err := BscPledgeOracleMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BscPledgeOracleBin), backend, _owner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BscPledgeOracle{BscPledgeOracleCaller: BscPledgeOracleCaller{contract: contract}, BscPledgeOracleTransactor: BscPledgeOracleTransactor{contract: contract}, BscPledgeOracleFilterer: BscPledgeOracleFilterer{contract: contract}}, nil
}

// BscPledgeOracle is an auto generated Go binding around an Ethereum contract.
type BscPledgeOracle struct {
	BscPledgeOracleCaller     // Read-only binding to the contract
	BscPledgeOracleTransactor // Write-only binding to the contract
	BscPledgeOracleFilterer   // Log filterer for contract events
}

// BscPledgeOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type BscPledgeOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BscPledgeOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BscPledgeOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BscPledgeOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BscPledgeOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BscPledgeOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BscPledgeOracleSession struct {
	Contract     *BscPledgeOracle  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BscPledgeOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BscPledgeOracleCallerSession struct {
	Contract *BscPledgeOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// BscPledgeOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BscPledgeOracleTransactorSession struct {
	Contract     *BscPledgeOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// BscPledgeOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type BscPledgeOracleRaw struct {
	Contract *BscPledgeOracle // Generic contract binding to access the raw methods on
}

// BscPledgeOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BscPledgeOracleCallerRaw struct {
	Contract *BscPledgeOracleCaller // Generic read-only contract binding to access the raw methods on
}

// BscPledgeOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BscPledgeOracleTransactorRaw struct {
	Contract *BscPledgeOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBscPledgeOracle creates a new instance of BscPledgeOracle, bound to a specific deployed contract.
func NewBscPledgeOracle(address common.Address, backend bind.ContractBackend) (*BscPledgeOracle, error) {
	contract, err := bindBscPledgeOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BscPledgeOracle{BscPledgeOracleCaller: BscPledgeOracleCaller{contract: contract}, BscPledgeOracleTransactor: BscPledgeOracleTransactor{contract: contract}, BscPledgeOracleFilterer: BscPledgeOracleFilterer{contract: contract}}, nil
}

// NewBscPledgeOracleCaller creates a new read-only instance of BscPledgeOracle, bound to a specific deployed contract.
func NewBscPledgeOracleCaller(address common.Address, caller bind.ContractCaller) (*BscPledgeOracleCaller, error) {
	contract, err := bindBscPledgeOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BscPledgeOracleCaller{contract: contract}, nil
}

// NewBscPledgeOracleTransactor creates a new write-only instance of BscPledgeOracle, bound to a specific deployed contract.
func NewBscPledgeOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*BscPledgeOracleTransactor, error) {
	contract, err := bindBscPledgeOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BscPledgeOracleTransactor{contract: contract}, nil
}

// NewBscPledgeOracleFilterer creates a new log filterer instance of BscPledgeOracle, bound to a specific deployed contract.
func NewBscPledgeOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*BscPledgeOracleFilterer, error) {
	contract, err := bindBscPledgeOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BscPledgeOracleFilterer{contract: contract}, nil
}

// bindBscPledgeOracle binds a generic wrapper to an already deployed contract.
func bindBscPledgeOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BscPledgeOracleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BscPledgeOracle *BscPledgeOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BscPledgeOracle.Contract.BscPledgeOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BscPledgeOracle *BscPledgeOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BscPledgeOracle.Contract.BscPledgeOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BscPledgeOracle *BscPledgeOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BscPledgeOracle.Contract.BscPledgeOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BscPledgeOracle *BscPledgeOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BscPledgeOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BscPledgeOracle *BscPledgeOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BscPledgeOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BscPledgeOracle *BscPledgeOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BscPledgeOracle.Contract.contract.Transact(opts, method, params...)
}

// GetAssetsAggregator is a free data retrieval call binding the contract method 0xb889a989.
//
// Solidity: function getAssetsAggregator(address asset) view returns(address, uint256)
func (_BscPledgeOracle *BscPledgeOracleCaller) GetAssetsAggregator(opts *bind.CallOpts, asset common.Address) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _BscPledgeOracle.contract.Call(opts, &out, "getAssetsAggregator", asset)

	if err != nil {
		return *new(common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetAssetsAggregator is a free data retrieval call binding the contract method 0xb889a989.
//
// Solidity: function getAssetsAggregator(address asset) view returns(address, uint256)
func (_BscPledgeOracle *BscPledgeOracleSession) GetAssetsAggregator(asset common.Address) (common.Address, *big.Int, error) {
	return _BscPledgeOracle.Contract.GetAssetsAggregator(&_BscPledgeOracle.CallOpts, asset)
}

// GetAssetsAggregator is a free data retrieval call binding the contract method 0xb889a989.
//
// Solidity: function getAssetsAggregator(address asset) view returns(address, uint256)
func (_BscPledgeOracle *BscPledgeOracleCallerSession) GetAssetsAggregator(asset common.Address) (common.Address, *big.Int, error) {
	return _BscPledgeOracle.Contract.GetAssetsAggregator(&_BscPledgeOracle.CallOpts, asset)
}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address asset) view returns(uint256)
func (_BscPledgeOracle *BscPledgeOracleCaller) GetPrice(opts *bind.CallOpts, asset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BscPledgeOracle.contract.Call(opts, &out, "getPrice", asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address asset) view returns(uint256)
func (_BscPledgeOracle *BscPledgeOracleSession) GetPrice(asset common.Address) (*big.Int, error) {
	return _BscPledgeOracle.Contract.GetPrice(&_BscPledgeOracle.CallOpts, asset)
}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address asset) view returns(uint256)
func (_BscPledgeOracle *BscPledgeOracleCallerSession) GetPrice(asset common.Address) (*big.Int, error) {
	return _BscPledgeOracle.Contract.GetPrice(&_BscPledgeOracle.CallOpts, asset)
}

// GetPrices is a free data retrieval call binding the contract method 0x09cb3a4e.
//
// Solidity: function getPrices(uint256[] assets) view returns(uint256[])
func (_BscPledgeOracle *BscPledgeOracleCaller) GetPrices(opts *bind.CallOpts, assets []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _BscPledgeOracle.contract.Call(opts, &out, "getPrices", assets)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetPrices is a free data retrieval call binding the contract method 0x09cb3a4e.
//
// Solidity: function getPrices(uint256[] assets) view returns(uint256[])
func (_BscPledgeOracle *BscPledgeOracleSession) GetPrices(assets []*big.Int) ([]*big.Int, error) {
	return _BscPledgeOracle.Contract.GetPrices(&_BscPledgeOracle.CallOpts, assets)
}

// GetPrices is a free data retrieval call binding the contract method 0x09cb3a4e.
//
// Solidity: function getPrices(uint256[] assets) view returns(uint256[])
func (_BscPledgeOracle *BscPledgeOracleCallerSession) GetPrices(assets []*big.Int) ([]*big.Int, error) {
	return _BscPledgeOracle.Contract.GetPrices(&_BscPledgeOracle.CallOpts, assets)
}

// GetUnderlyingAggregator is a free data retrieval call binding the contract method 0x75e443aa.
//
// Solidity: function getUnderlyingAggregator(uint256 underlying) view returns(address, uint256)
func (_BscPledgeOracle *BscPledgeOracleCaller) GetUnderlyingAggregator(opts *bind.CallOpts, underlying *big.Int) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _BscPledgeOracle.contract.Call(opts, &out, "getUnderlyingAggregator", underlying)

	if err != nil {
		return *new(common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetUnderlyingAggregator is a free data retrieval call binding the contract method 0x75e443aa.
//
// Solidity: function getUnderlyingAggregator(uint256 underlying) view returns(address, uint256)
func (_BscPledgeOracle *BscPledgeOracleSession) GetUnderlyingAggregator(underlying *big.Int) (common.Address, *big.Int, error) {
	return _BscPledgeOracle.Contract.GetUnderlyingAggregator(&_BscPledgeOracle.CallOpts, underlying)
}

// GetUnderlyingAggregator is a free data retrieval call binding the contract method 0x75e443aa.
//
// Solidity: function getUnderlyingAggregator(uint256 underlying) view returns(address, uint256)
func (_BscPledgeOracle *BscPledgeOracleCallerSession) GetUnderlyingAggregator(underlying *big.Int) (common.Address, *big.Int, error) {
	return _BscPledgeOracle.Contract.GetUnderlyingAggregator(&_BscPledgeOracle.CallOpts, underlying)
}

// GetUnderlyingPrice is a free data retrieval call binding the contract method 0xda663257.
//
// Solidity: function getUnderlyingPrice(uint256 underlying) view returns(uint256)
func (_BscPledgeOracle *BscPledgeOracleCaller) GetUnderlyingPrice(opts *bind.CallOpts, underlying *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BscPledgeOracle.contract.Call(opts, &out, "getUnderlyingPrice", underlying)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUnderlyingPrice is a free data retrieval call binding the contract method 0xda663257.
//
// Solidity: function getUnderlyingPrice(uint256 underlying) view returns(uint256)
func (_BscPledgeOracle *BscPledgeOracleSession) GetUnderlyingPrice(underlying *big.Int) (*big.Int, error) {
	return _BscPledgeOracle.Contract.GetUnderlyingPrice(&_BscPledgeOracle.CallOpts, underlying)
}

// GetUnderlyingPrice is a free data retrieval call binding the contract method 0xda663257.
//
// Solidity: function getUnderlyingPrice(uint256 underlying) view returns(uint256)
func (_BscPledgeOracle *BscPledgeOracleCallerSession) GetUnderlyingPrice(underlying *big.Int) (*big.Int, error) {
	return _BscPledgeOracle.Contract.GetUnderlyingPrice(&_BscPledgeOracle.CallOpts, underlying)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BscPledgeOracle *BscPledgeOracleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BscPledgeOracle.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BscPledgeOracle *BscPledgeOracleSession) Owner() (common.Address, error) {
	return _BscPledgeOracle.Contract.Owner(&_BscPledgeOracle.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BscPledgeOracle *BscPledgeOracleCallerSession) Owner() (common.Address, error) {
	return _BscPledgeOracle.Contract.Owner(&_BscPledgeOracle.CallOpts)
}

// PriceDivisor is a free data retrieval call binding the contract method 0xd690a8c3.
//
// Solidity: function priceDivisor() view returns(uint256)
func (_BscPledgeOracle *BscPledgeOracleCaller) PriceDivisor(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BscPledgeOracle.contract.Call(opts, &out, "priceDivisor")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PriceDivisor is a free data retrieval call binding the contract method 0xd690a8c3.
//
// Solidity: function priceDivisor() view returns(uint256)
func (_BscPledgeOracle *BscPledgeOracleSession) PriceDivisor() (*big.Int, error) {
	return _BscPledgeOracle.Contract.PriceDivisor(&_BscPledgeOracle.CallOpts)
}

// PriceDivisor is a free data retrieval call binding the contract method 0xd690a8c3.
//
// Solidity: function priceDivisor() view returns(uint256)
func (_BscPledgeOracle *BscPledgeOracleCallerSession) PriceDivisor() (*big.Int, error) {
	return _BscPledgeOracle.Contract.PriceDivisor(&_BscPledgeOracle.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BscPledgeOracle *BscPledgeOracleTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BscPledgeOracle.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BscPledgeOracle *BscPledgeOracleSession) RenounceOwnership() (*types.Transaction, error) {
	return _BscPledgeOracle.Contract.RenounceOwnership(&_BscPledgeOracle.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BscPledgeOracle *BscPledgeOracleTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BscPledgeOracle.Contract.RenounceOwnership(&_BscPledgeOracle.TransactOpts)
}

// SetAssetsAggregator is a paid mutator transaction binding the contract method 0xbe7b4adc.
//
// Solidity: function setAssetsAggregator(address asset, address aggregator, uint8 tokenDecimals) returns()
func (_BscPledgeOracle *BscPledgeOracleTransactor) SetAssetsAggregator(opts *bind.TransactOpts, asset common.Address, aggregator common.Address, tokenDecimals uint8) (*types.Transaction, error) {
	return _BscPledgeOracle.contract.Transact(opts, "setAssetsAggregator", asset, aggregator, tokenDecimals)
}

// SetAssetsAggregator is a paid mutator transaction binding the contract method 0xbe7b4adc.
//
// Solidity: function setAssetsAggregator(address asset, address aggregator, uint8 tokenDecimals) returns()
func (_BscPledgeOracle *BscPledgeOracleSession) SetAssetsAggregator(asset common.Address, aggregator common.Address, tokenDecimals uint8) (*types.Transaction, error) {
	return _BscPledgeOracle.Contract.SetAssetsAggregator(&_BscPledgeOracle.TransactOpts, asset, aggregator, tokenDecimals)
}

// SetAssetsAggregator is a paid mutator transaction binding the contract method 0xbe7b4adc.
//
// Solidity: function setAssetsAggregator(address asset, address aggregator, uint8 tokenDecimals) returns()
func (_BscPledgeOracle *BscPledgeOracleTransactorSession) SetAssetsAggregator(asset common.Address, aggregator common.Address, tokenDecimals uint8) (*types.Transaction, error) {
	return _BscPledgeOracle.Contract.SetAssetsAggregator(&_BscPledgeOracle.TransactOpts, asset, aggregator, tokenDecimals)
}

// SetPrice is a paid mutator transaction binding the contract method 0x00e4768b.
//
// Solidity: function setPrice(address asset, uint256 price) returns()
func (_BscPledgeOracle *BscPledgeOracleTransactor) SetPrice(opts *bind.TransactOpts, asset common.Address, price *big.Int) (*types.Transaction, error) {
	return _BscPledgeOracle.contract.Transact(opts, "setPrice", asset, price)
}

// SetPrice is a paid mutator transaction binding the contract method 0x00e4768b.
//
// Solidity: function setPrice(address asset, uint256 price) returns()
func (_BscPledgeOracle *BscPledgeOracleSession) SetPrice(asset common.Address, price *big.Int) (*types.Transaction, error) {
	return _BscPledgeOracle.Contract.SetPrice(&_BscPledgeOracle.TransactOpts, asset, price)
}

// SetPrice is a paid mutator transaction binding the contract method 0x00e4768b.
//
// Solidity: function setPrice(address asset, uint256 price) returns()
func (_BscPledgeOracle *BscPledgeOracleTransactorSession) SetPrice(asset common.Address, price *big.Int) (*types.Transaction, error) {
	return _BscPledgeOracle.Contract.SetPrice(&_BscPledgeOracle.TransactOpts, asset, price)
}

// SetPriceDivisor is a paid mutator transaction binding the contract method 0xb2861e7b.
//
// Solidity: function setPriceDivisor(uint256 newDivisor) returns()
func (_BscPledgeOracle *BscPledgeOracleTransactor) SetPriceDivisor(opts *bind.TransactOpts, newDivisor *big.Int) (*types.Transaction, error) {
	return _BscPledgeOracle.contract.Transact(opts, "setPriceDivisor", newDivisor)
}

// SetPriceDivisor is a paid mutator transaction binding the contract method 0xb2861e7b.
//
// Solidity: function setPriceDivisor(uint256 newDivisor) returns()
func (_BscPledgeOracle *BscPledgeOracleSession) SetPriceDivisor(newDivisor *big.Int) (*types.Transaction, error) {
	return _BscPledgeOracle.Contract.SetPriceDivisor(&_BscPledgeOracle.TransactOpts, newDivisor)
}

// SetPriceDivisor is a paid mutator transaction binding the contract method 0xb2861e7b.
//
// Solidity: function setPriceDivisor(uint256 newDivisor) returns()
func (_BscPledgeOracle *BscPledgeOracleTransactorSession) SetPriceDivisor(newDivisor *big.Int) (*types.Transaction, error) {
	return _BscPledgeOracle.Contract.SetPriceDivisor(&_BscPledgeOracle.TransactOpts, newDivisor)
}

// SetPrices is a paid mutator transaction binding the contract method 0xd05eaae0.
//
// Solidity: function setPrices(uint256[] assets, uint256[] prices) returns()
func (_BscPledgeOracle *BscPledgeOracleTransactor) SetPrices(opts *bind.TransactOpts, assets []*big.Int, prices []*big.Int) (*types.Transaction, error) {
	return _BscPledgeOracle.contract.Transact(opts, "setPrices", assets, prices)
}

// SetPrices is a paid mutator transaction binding the contract method 0xd05eaae0.
//
// Solidity: function setPrices(uint256[] assets, uint256[] prices) returns()
func (_BscPledgeOracle *BscPledgeOracleSession) SetPrices(assets []*big.Int, prices []*big.Int) (*types.Transaction, error) {
	return _BscPledgeOracle.Contract.SetPrices(&_BscPledgeOracle.TransactOpts, assets, prices)
}

// SetPrices is a paid mutator transaction binding the contract method 0xd05eaae0.
//
// Solidity: function setPrices(uint256[] assets, uint256[] prices) returns()
func (_BscPledgeOracle *BscPledgeOracleTransactorSession) SetPrices(assets []*big.Int, prices []*big.Int) (*types.Transaction, error) {
	return _BscPledgeOracle.Contract.SetPrices(&_BscPledgeOracle.TransactOpts, assets, prices)
}

// SetUnderlyingAggregator is a paid mutator transaction binding the contract method 0x0b905a2e.
//
// Solidity: function setUnderlyingAggregator(uint256 underlying, address aggregator, uint8 tokenDecimals) returns()
func (_BscPledgeOracle *BscPledgeOracleTransactor) SetUnderlyingAggregator(opts *bind.TransactOpts, underlying *big.Int, aggregator common.Address, tokenDecimals uint8) (*types.Transaction, error) {
	return _BscPledgeOracle.contract.Transact(opts, "setUnderlyingAggregator", underlying, aggregator, tokenDecimals)
}

// SetUnderlyingAggregator is a paid mutator transaction binding the contract method 0x0b905a2e.
//
// Solidity: function setUnderlyingAggregator(uint256 underlying, address aggregator, uint8 tokenDecimals) returns()
func (_BscPledgeOracle *BscPledgeOracleSession) SetUnderlyingAggregator(underlying *big.Int, aggregator common.Address, tokenDecimals uint8) (*types.Transaction, error) {
	return _BscPledgeOracle.Contract.SetUnderlyingAggregator(&_BscPledgeOracle.TransactOpts, underlying, aggregator, tokenDecimals)
}

// SetUnderlyingAggregator is a paid mutator transaction binding the contract method 0x0b905a2e.
//
// Solidity: function setUnderlyingAggregator(uint256 underlying, address aggregator, uint8 tokenDecimals) returns()
func (_BscPledgeOracle *BscPledgeOracleTransactorSession) SetUnderlyingAggregator(underlying *big.Int, aggregator common.Address, tokenDecimals uint8) (*types.Transaction, error) {
	return _BscPledgeOracle.Contract.SetUnderlyingAggregator(&_BscPledgeOracle.TransactOpts, underlying, aggregator, tokenDecimals)
}

// SetUnderlyingPrice is a paid mutator transaction binding the contract method 0x83532667.
//
// Solidity: function setUnderlyingPrice(uint256 underlying, uint256 price) returns()
func (_BscPledgeOracle *BscPledgeOracleTransactor) SetUnderlyingPrice(opts *bind.TransactOpts, underlying *big.Int, price *big.Int) (*types.Transaction, error) {
	return _BscPledgeOracle.contract.Transact(opts, "setUnderlyingPrice", underlying, price)
}

// SetUnderlyingPrice is a paid mutator transaction binding the contract method 0x83532667.
//
// Solidity: function setUnderlyingPrice(uint256 underlying, uint256 price) returns()
func (_BscPledgeOracle *BscPledgeOracleSession) SetUnderlyingPrice(underlying *big.Int, price *big.Int) (*types.Transaction, error) {
	return _BscPledgeOracle.Contract.SetUnderlyingPrice(&_BscPledgeOracle.TransactOpts, underlying, price)
}

// SetUnderlyingPrice is a paid mutator transaction binding the contract method 0x83532667.
//
// Solidity: function setUnderlyingPrice(uint256 underlying, uint256 price) returns()
func (_BscPledgeOracle *BscPledgeOracleTransactorSession) SetUnderlyingPrice(underlying *big.Int, price *big.Int) (*types.Transaction, error) {
	return _BscPledgeOracle.Contract.SetUnderlyingPrice(&_BscPledgeOracle.TransactOpts, underlying, price)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BscPledgeOracle *BscPledgeOracleTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BscPledgeOracle.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BscPledgeOracle *BscPledgeOracleSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BscPledgeOracle.Contract.TransferOwnership(&_BscPledgeOracle.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BscPledgeOracle *BscPledgeOracleTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BscPledgeOracle.Contract.TransferOwnership(&_BscPledgeOracle.TransactOpts, newOwner)
}

// BscPledgeOracleAggregatorSetIterator is returned from FilterAggregatorSet and is used to iterate over the raw logs and unpacked data for AggregatorSet events raised by the BscPledgeOracle contract.
type BscPledgeOracleAggregatorSetIterator struct {
	Event *BscPledgeOracleAggregatorSet // Event containing the contract specifics and raw log

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
func (it *BscPledgeOracleAggregatorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BscPledgeOracleAggregatorSet)
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
		it.Event = new(BscPledgeOracleAggregatorSet)
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
func (it *BscPledgeOracleAggregatorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BscPledgeOracleAggregatorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BscPledgeOracleAggregatorSet represents a AggregatorSet event raised by the BscPledgeOracle contract.
type BscPledgeOracleAggregatorSet struct {
	Asset         *big.Int
	Aggregator    common.Address
	TokenDecimals uint8
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAggregatorSet is a free log retrieval operation binding the contract event 0x5fa5791e5634d03047f423c78a1370531e6d289f6caf006a94fa039a3b5dce93.
//
// Solidity: event AggregatorSet(uint256 asset, address aggregator, uint8 tokenDecimals)
func (_BscPledgeOracle *BscPledgeOracleFilterer) FilterAggregatorSet(opts *bind.FilterOpts) (*BscPledgeOracleAggregatorSetIterator, error) {

	logs, sub, err := _BscPledgeOracle.contract.FilterLogs(opts, "AggregatorSet")
	if err != nil {
		return nil, err
	}
	return &BscPledgeOracleAggregatorSetIterator{contract: _BscPledgeOracle.contract, event: "AggregatorSet", logs: logs, sub: sub}, nil
}

// WatchAggregatorSet is a free log subscription operation binding the contract event 0x5fa5791e5634d03047f423c78a1370531e6d289f6caf006a94fa039a3b5dce93.
//
// Solidity: event AggregatorSet(uint256 asset, address aggregator, uint8 tokenDecimals)
func (_BscPledgeOracle *BscPledgeOracleFilterer) WatchAggregatorSet(opts *bind.WatchOpts, sink chan<- *BscPledgeOracleAggregatorSet) (event.Subscription, error) {

	logs, sub, err := _BscPledgeOracle.contract.WatchLogs(opts, "AggregatorSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BscPledgeOracleAggregatorSet)
				if err := _BscPledgeOracle.contract.UnpackLog(event, "AggregatorSet", log); err != nil {
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

// ParseAggregatorSet is a log parse operation binding the contract event 0x5fa5791e5634d03047f423c78a1370531e6d289f6caf006a94fa039a3b5dce93.
//
// Solidity: event AggregatorSet(uint256 asset, address aggregator, uint8 tokenDecimals)
func (_BscPledgeOracle *BscPledgeOracleFilterer) ParseAggregatorSet(log types.Log) (*BscPledgeOracleAggregatorSet, error) {
	event := new(BscPledgeOracleAggregatorSet)
	if err := _BscPledgeOracle.contract.UnpackLog(event, "AggregatorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BscPledgeOracleManualPriceSetIterator is returned from FilterManualPriceSet and is used to iterate over the raw logs and unpacked data for ManualPriceSet events raised by the BscPledgeOracle contract.
type BscPledgeOracleManualPriceSetIterator struct {
	Event *BscPledgeOracleManualPriceSet // Event containing the contract specifics and raw log

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
func (it *BscPledgeOracleManualPriceSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BscPledgeOracleManualPriceSet)
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
		it.Event = new(BscPledgeOracleManualPriceSet)
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
func (it *BscPledgeOracleManualPriceSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BscPledgeOracleManualPriceSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BscPledgeOracleManualPriceSet represents a ManualPriceSet event raised by the BscPledgeOracle contract.
type BscPledgeOracleManualPriceSet struct {
	Asset *big.Int
	Price *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterManualPriceSet is a free log retrieval operation binding the contract event 0xb4e7056b10bc2d2d4fe5b2ff374b24edf94410f4620866b12b3bcf1765d91894.
//
// Solidity: event ManualPriceSet(uint256 asset, uint256 price)
func (_BscPledgeOracle *BscPledgeOracleFilterer) FilterManualPriceSet(opts *bind.FilterOpts) (*BscPledgeOracleManualPriceSetIterator, error) {

	logs, sub, err := _BscPledgeOracle.contract.FilterLogs(opts, "ManualPriceSet")
	if err != nil {
		return nil, err
	}
	return &BscPledgeOracleManualPriceSetIterator{contract: _BscPledgeOracle.contract, event: "ManualPriceSet", logs: logs, sub: sub}, nil
}

// WatchManualPriceSet is a free log subscription operation binding the contract event 0xb4e7056b10bc2d2d4fe5b2ff374b24edf94410f4620866b12b3bcf1765d91894.
//
// Solidity: event ManualPriceSet(uint256 asset, uint256 price)
func (_BscPledgeOracle *BscPledgeOracleFilterer) WatchManualPriceSet(opts *bind.WatchOpts, sink chan<- *BscPledgeOracleManualPriceSet) (event.Subscription, error) {

	logs, sub, err := _BscPledgeOracle.contract.WatchLogs(opts, "ManualPriceSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BscPledgeOracleManualPriceSet)
				if err := _BscPledgeOracle.contract.UnpackLog(event, "ManualPriceSet", log); err != nil {
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

// ParseManualPriceSet is a log parse operation binding the contract event 0xb4e7056b10bc2d2d4fe5b2ff374b24edf94410f4620866b12b3bcf1765d91894.
//
// Solidity: event ManualPriceSet(uint256 asset, uint256 price)
func (_BscPledgeOracle *BscPledgeOracleFilterer) ParseManualPriceSet(log types.Log) (*BscPledgeOracleManualPriceSet, error) {
	event := new(BscPledgeOracleManualPriceSet)
	if err := _BscPledgeOracle.contract.UnpackLog(event, "ManualPriceSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BscPledgeOracleManualPricesSetIterator is returned from FilterManualPricesSet and is used to iterate over the raw logs and unpacked data for ManualPricesSet events raised by the BscPledgeOracle contract.
type BscPledgeOracleManualPricesSetIterator struct {
	Event *BscPledgeOracleManualPricesSet // Event containing the contract specifics and raw log

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
func (it *BscPledgeOracleManualPricesSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BscPledgeOracleManualPricesSet)
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
		it.Event = new(BscPledgeOracleManualPricesSet)
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
func (it *BscPledgeOracleManualPricesSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BscPledgeOracleManualPricesSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BscPledgeOracleManualPricesSet represents a ManualPricesSet event raised by the BscPledgeOracle contract.
type BscPledgeOracleManualPricesSet struct {
	Assets []*big.Int
	Prices []*big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterManualPricesSet is a free log retrieval operation binding the contract event 0xff1a86ea42b3a30c49373419293ba99ec1535aaef99ebd6f4be0a307a54ad1d6.
//
// Solidity: event ManualPricesSet(uint256[] assets, uint256[] prices)
func (_BscPledgeOracle *BscPledgeOracleFilterer) FilterManualPricesSet(opts *bind.FilterOpts) (*BscPledgeOracleManualPricesSetIterator, error) {

	logs, sub, err := _BscPledgeOracle.contract.FilterLogs(opts, "ManualPricesSet")
	if err != nil {
		return nil, err
	}
	return &BscPledgeOracleManualPricesSetIterator{contract: _BscPledgeOracle.contract, event: "ManualPricesSet", logs: logs, sub: sub}, nil
}

// WatchManualPricesSet is a free log subscription operation binding the contract event 0xff1a86ea42b3a30c49373419293ba99ec1535aaef99ebd6f4be0a307a54ad1d6.
//
// Solidity: event ManualPricesSet(uint256[] assets, uint256[] prices)
func (_BscPledgeOracle *BscPledgeOracleFilterer) WatchManualPricesSet(opts *bind.WatchOpts, sink chan<- *BscPledgeOracleManualPricesSet) (event.Subscription, error) {

	logs, sub, err := _BscPledgeOracle.contract.WatchLogs(opts, "ManualPricesSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BscPledgeOracleManualPricesSet)
				if err := _BscPledgeOracle.contract.UnpackLog(event, "ManualPricesSet", log); err != nil {
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

// ParseManualPricesSet is a log parse operation binding the contract event 0xff1a86ea42b3a30c49373419293ba99ec1535aaef99ebd6f4be0a307a54ad1d6.
//
// Solidity: event ManualPricesSet(uint256[] assets, uint256[] prices)
func (_BscPledgeOracle *BscPledgeOracleFilterer) ParseManualPricesSet(log types.Log) (*BscPledgeOracleManualPricesSet, error) {
	event := new(BscPledgeOracleManualPricesSet)
	if err := _BscPledgeOracle.contract.UnpackLog(event, "ManualPricesSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BscPledgeOracleOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BscPledgeOracle contract.
type BscPledgeOracleOwnershipTransferredIterator struct {
	Event *BscPledgeOracleOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BscPledgeOracleOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BscPledgeOracleOwnershipTransferred)
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
		it.Event = new(BscPledgeOracleOwnershipTransferred)
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
func (it *BscPledgeOracleOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BscPledgeOracleOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BscPledgeOracleOwnershipTransferred represents a OwnershipTransferred event raised by the BscPledgeOracle contract.
type BscPledgeOracleOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BscPledgeOracle *BscPledgeOracleFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BscPledgeOracleOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BscPledgeOracle.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BscPledgeOracleOwnershipTransferredIterator{contract: _BscPledgeOracle.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BscPledgeOracle *BscPledgeOracleFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BscPledgeOracleOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BscPledgeOracle.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BscPledgeOracleOwnershipTransferred)
				if err := _BscPledgeOracle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_BscPledgeOracle *BscPledgeOracleFilterer) ParseOwnershipTransferred(log types.Log) (*BscPledgeOracleOwnershipTransferred, error) {
	event := new(BscPledgeOracleOwnershipTransferred)
	if err := _BscPledgeOracle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BscPledgeOraclePriceDivisorUpdatedIterator is returned from FilterPriceDivisorUpdated and is used to iterate over the raw logs and unpacked data for PriceDivisorUpdated events raised by the BscPledgeOracle contract.
type BscPledgeOraclePriceDivisorUpdatedIterator struct {
	Event *BscPledgeOraclePriceDivisorUpdated // Event containing the contract specifics and raw log

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
func (it *BscPledgeOraclePriceDivisorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BscPledgeOraclePriceDivisorUpdated)
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
		it.Event = new(BscPledgeOraclePriceDivisorUpdated)
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
func (it *BscPledgeOraclePriceDivisorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BscPledgeOraclePriceDivisorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BscPledgeOraclePriceDivisorUpdated represents a PriceDivisorUpdated event raised by the BscPledgeOracle contract.
type BscPledgeOraclePriceDivisorUpdated struct {
	OldDivisor *big.Int
	NewDivisor *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPriceDivisorUpdated is a free log retrieval operation binding the contract event 0xcd572cdad82b8d7c7cc3eccd904c997bf961bb9e917c48a2c1214b200fb3e365.
//
// Solidity: event PriceDivisorUpdated(uint256 oldDivisor, uint256 newDivisor)
func (_BscPledgeOracle *BscPledgeOracleFilterer) FilterPriceDivisorUpdated(opts *bind.FilterOpts) (*BscPledgeOraclePriceDivisorUpdatedIterator, error) {

	logs, sub, err := _BscPledgeOracle.contract.FilterLogs(opts, "PriceDivisorUpdated")
	if err != nil {
		return nil, err
	}
	return &BscPledgeOraclePriceDivisorUpdatedIterator{contract: _BscPledgeOracle.contract, event: "PriceDivisorUpdated", logs: logs, sub: sub}, nil
}

// WatchPriceDivisorUpdated is a free log subscription operation binding the contract event 0xcd572cdad82b8d7c7cc3eccd904c997bf961bb9e917c48a2c1214b200fb3e365.
//
// Solidity: event PriceDivisorUpdated(uint256 oldDivisor, uint256 newDivisor)
func (_BscPledgeOracle *BscPledgeOracleFilterer) WatchPriceDivisorUpdated(opts *bind.WatchOpts, sink chan<- *BscPledgeOraclePriceDivisorUpdated) (event.Subscription, error) {

	logs, sub, err := _BscPledgeOracle.contract.WatchLogs(opts, "PriceDivisorUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BscPledgeOraclePriceDivisorUpdated)
				if err := _BscPledgeOracle.contract.UnpackLog(event, "PriceDivisorUpdated", log); err != nil {
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

// ParsePriceDivisorUpdated is a log parse operation binding the contract event 0xcd572cdad82b8d7c7cc3eccd904c997bf961bb9e917c48a2c1214b200fb3e365.
//
// Solidity: event PriceDivisorUpdated(uint256 oldDivisor, uint256 newDivisor)
func (_BscPledgeOracle *BscPledgeOracleFilterer) ParsePriceDivisorUpdated(log types.Log) (*BscPledgeOraclePriceDivisorUpdated, error) {
	event := new(BscPledgeOraclePriceDivisorUpdated)
	if err := _BscPledgeOracle.contract.UnpackLog(event, "PriceDivisorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
