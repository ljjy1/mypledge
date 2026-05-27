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

// PledgePoolCreatePoolParams is an auto generated low-level Go binding around an user-defined struct.
type PledgePoolCreatePoolParams struct {
	SettleTime             *big.Int
	EndTime                *big.Int
	InterestRate           *big.Int
	MaxSupply              *big.Int
	MortgageRate           *big.Int
	LendToken              common.Address
	BorrowToken            common.Address
	LendDebtToken          common.Address
	BorrowDebtToken        common.Address
	AutoLiquidateThreshold *big.Int
}

// PledgePoolMetaData contains all meta data concerning the PledgePool contract.
var PledgePoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_swapRouter\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_feeAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Borrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"borrowDebtTokenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lendTokenAmount\",\"type\":\"uint256\"}],\"name\":\"ClaimBorrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"lender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ClaimLendDebtToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"}],\"name\":\"CreatePledgePool\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"burnAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"redeemAmount\",\"type\":\"uint256\"}],\"name\":\"DestroyBorrowDebtToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"lender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"burnAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"redeemAmount\",\"type\":\"uint256\"}],\"name\":\"DestroyLendDebtToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EmergencyWithdrawBorrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"lender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EmergencyWithdrawLend\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"finishAmountLend\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"finishAmountBorrow\",\"type\":\"uint256\"}],\"name\":\"FinishPool\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"lender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Lend\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liquidationAmountLend\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liquidationAmountBorrow\",\"type\":\"uint256\"}],\"name\":\"LiquidatePool\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"refunder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RefundBorrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"refunder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RefundLend\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newLendFee\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newBorrowFee\",\"type\":\"uint256\"}],\"name\":\"SetFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldFeeAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newFeeAddress\",\"type\":\"address\"}],\"name\":\"SetFeeAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"oldPaused\",\"type\":\"bool\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"newPaused\",\"type\":\"bool\"}],\"name\":\"SetGlobalPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"oldMinAmount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newMinAmount\",\"type\":\"uint256\"}],\"name\":\"SetMinAmount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOracle\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOracle\",\"type\":\"address\"}],\"name\":\"SetOracle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldSwapAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newSwapAddress\",\"type\":\"address\"}],\"name\":\"SetSwapRouterAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"settleAmountLend\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"settleAmountBorrow\",\"type\":\"uint256\"}],\"name\":\"SettlePool\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_borrowTokenAmount\",\"type\":\"uint256\"}],\"name\":\"borrow\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"borrowFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"borrowInfoMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"returnBorrowAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"hasNoRefund\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"hasNoClaim\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"checkCanFinish\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"checkCanLiquidate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"checkCanSettle\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"claimBorrow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"claimLendDebtToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"settleTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"interestRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mortgageRate\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"lendToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrowToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"lendDebtToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrowDebtToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"autoLiquidateThreshold\",\"type\":\"uint256\"}],\"internalType\":\"structPledgePool.CreatePoolParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"createPledgePool\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_borrowDebtTokenAmount\",\"type\":\"uint256\"}],\"name\":\"destroyBorrowDebtToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_lendDebtTokenAmount\",\"type\":\"uint256\"}],\"name\":\"destroyLendDebtToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"emergencyWithdrawBorrow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"emergencyWithdrawLend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeAddress\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"finishPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"getPoolState\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_lendAmount\",\"type\":\"uint256\"}],\"name\":\"lend\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lendFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"lendInfoMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"lendAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"refundLendAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"hasNoRefund\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"hasNoClaim\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"liquidatePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pledgePoolInfoList\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"settleTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"interestRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lendSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mortgageRate\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"lendToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrowToken\",\"type\":\"address\"},{\"internalType\":\"enumPledgePool.PoolState\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"lendDebtToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrowDebtToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"autoLiquidateThreshold\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"poolDataInfoList\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"settleAmountLend\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"settleAmountBorrow\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"finishAmountLend\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"finishAmountBorrow\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationAmountLend\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationAmountBorrow\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"refundBorrow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"refundLend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_lendFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_borrowFee\",\"type\":\"uint256\"}],\"name\":\"setFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_feeAddress\",\"type\":\"address\"}],\"name\":\"setFeeAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setGlobalPaused\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minAmount\",\"type\":\"uint256\"}],\"name\":\"setMinAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"}],\"name\":\"setOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_swapRouter\",\"type\":\"address\"}],\"name\":\"setSwapRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"settlePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swapRouter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60803461027957601f61498738819003918201601f19168301916001600160401b0383118484101761027d578084926080946040528339810103126102795761004781610291565b61005360208301610291565b60408301516001600160a01b03811693908490036102795760606100779101610291565b60017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00556001600160a01b0316908115610266575f80546001600160a01b031981168417825560405193916001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09080a368056bc75e2d63100000600155600254916001600160a01b0384161561022457506001600160a01b03169182156101d4578315610184576001600160a81b031990911660089190911b610100600160a81b031617600255600380546001600160a01b0319908116929092179055600480549091169190911790555f60058190556006556040516146e190816102a68239f35b60405162461bcd60e51b815260206004820152602260248201527f4665654164647265737320616464726573732063616e6e6f7420626520656d70604482015261747960f01b6064820152608490fd5b60405162461bcd60e51b815260206004820152602260248201527f53776170526f7574657220616464726573732063616e6e6f7420626520656d70604482015261747960f01b6064820152608490fd5b62461bcd60e51b815260206004820152601e60248201527f4f7261636c6520616464726573732063616e6e6f7420626520656d70747900006044820152606490fd5b631e4fbdf760e01b5f525f60045260245ffd5b5f80fd5b634e487b7160e01b5f52604160045260245ffd5b51906001600160a01b03821682036102795756fe60806040526004361015610011575f80fd5b5f5f3560e01c80630beccc401461306a5780630ecbcdab14612e7757806310b72d6814612a0657806317f1854c1461291f5780633ab4a44514612632578063412736571461256e57806341275358146125455780634945c25b146121215780634aea0aec146121035780634def20da14611ee45780634ec2d87514611ebc57806350e0a6b814611e9d57806351dbf7f914611e2057806352f7c98814611ccb57806361a552dc14611ca85780636f68d40e14611c89578063715018a614611c2f5780637adbf97314611b5d5780637cad0a3414611af15780637dc0d1d014611ac45780637e32e47b146118e35780638705fcd41461181b578063897b0637146117865780638da5cb5b1461175f5780639110df851461147a5780639b2cb5d81461145c578063a62ff164146111f1578063b1597517146111a0578063c31c9c0714611177578063c4bd6fe614611121578063db45037d14610cf0578063dcfddacf14610c6f578063dd182bb514610a54578063e0f86bcc146105a3578063e626648a14610585578063eec8d5061461026d578063f2fde38b146101e75763ff920635146101bc575f80fd5b346101e45760203660031901126101e45760206101da600435613e77565b6040519015158152f35b80fd5b50346101e45760203660031901126101e4576102016134ba565b610209614592565b6001600160a01b031680156102595781546001600160a01b03198116821783556001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08380a380f35b631e4fbdf760e01b82526004829052602482fd5b50346101e45760203660031901126101e45760043561028a613ff1565b61029960ff60025416156134ec565b6102af6102a58261348a565b505442101561383a565b60ff60086102bc8361348a565b50015460a01c1660058110156105715715158061054a575b6102dd9061387f565b6102e68161348a565b50906102f1816134d0565b5033845260096020526040842082855260205260408420928354918215906103198215613946565b6004830154905461033461032d828461380e565b1515613e38565b600287019485549360ff851661051557670de0b6b3a76400008202918204670de0b6b3a764000014171561050157916103836103889261037d83670de0b6b3a7640000966135e2565b9261380e565b6135bb565b049485156104b0576007936001809360ff1916179055016103aa8582546137b5565b90550180546001600160a01b031680610422575083808080863382f115610417575b60018060a01b03905416906040519283527f366d54abb3194a1830bd9174cd4ada6d761f18ad3975cdb0895e36f5d2fae03e60203394a460015f51602061468c5f395f51905f525580f35b6040513d85823e3d90fd5b6040516370a0823160e01b8152306004820152602081602481855afa80156104a5578591879161046a575b509161045e8261046594101561364c565b3390614029565b6103cc565b9150506020813d60201161049d575b8161048660209383613585565b810103126104995751849061045e61044d565b5f80fd5b3d9150610479565b6040513d88823e3d90fd5b60405162461bcd60e51b8152602060048201526024808201527f526566756e6420616d6f756e74206d75737420626520677265617465722074686044820152630616e20360e41b6064820152608490fd5b634e487b7160e01b89526011600452602489fd5b60405162461bcd60e51b815260206004820152600d60248201526c12185cc81b9bc81c99599d5b99609a1b6044820152606490fd5b5060ff60086105588361348a565b50015460a01c16600581101561057157600414156102d4565b634e487b7160e01b83526021600452602483fd5b50346101e457806003193601126101e4576020600654604051908152f35b50346101e45760203660031901126101e4576004356105c0614592565b6105c98161348a565b50600881019081549160ff8360a01c166005811015610a405760016105ee91146137c2565b60018201549161060083421015613698565b61063c61061961060f876134d0565b509483549061380e565b660b342eb7c380006106358654926103836002870154856135bb565b04906137b5565b936305f5e10061064e600554876135bb565b049487600761065d88846137b5565b940180549094936001600160a01b03918216911681810361073e57505050600285015560018401547fe197bf82539ba04efe8d438ac6dc6e530ed505bb8442c1c746badc834f159777956040959493929091818110610733576106d7916106c39161380e565b60065485546001600160a01b03169061454d565b600385015580610710575b50505b805460ff60a01b1916600160a11b17905560028101546003919091015482519182526020820152a280f35b905460045461072c92916001600160a01b039182169116614505565b5f806106e2565b50506106d7886106c3565b600354604051633fc8cef360e01b8152959695929190602090849060049082906001600160a01b03165afa928315610a35578493610a04575b50806109fe57825b826109f857835b6001600160a01b039182169116036108ad57505085546001600160a01b031615610848575b5050907fe197bf82539ba04efe8d438ac6dc6e530ed505bb8442c1c746badc834f15977795859493926002604097015580610825575b5050600183015481811061081a57610810916107fc9161380e565b60065483546001600160a01b03169061454d565b60038301556106e5565b5050610810866107fc565b905460045461084192916001600160a01b039182169116614505565b5f806107e1565b6001600160a01b0316803b156108a957818591600460405180948193630d0e30db60e41b83525af1801561089e57156107ab578161088b91979695949397613585565b61089a5790919293875f6107ab565b8780fd5b6040513d84823e3d90fd5b5080fd5b809499506108c29350819598979692506141e2565b916001850191825484116109a0576108db9184916142a5565b86811061095b577fe197bf82539ba04efe8d438ac6dc6e530ed505bb8442c1c746badc834f159777966040968183111561094f5761091d82610934939461380e565b90546004546001600160a01b039081169116614505565b60028501555b5481811061081a57610810916107fc9161380e565b5050600285015561093a565b60405162461bcd60e51b815260206004820152601d60248201527f66696e697368506f6f6c3a20536c69707061676520746f6f20686967680000006044820152606490fd5b60405162461bcd60e51b815260206004820152602a60248201527f66696e697368506f6f6c3a20696e73756666696369656e7420626f72726f772060448201526918dbdb1b185d195c985b60b21b6064820152608490fd5b82610786565b8061077f565b610a2791935060203d602011610a2e575b610a1f8183613585565b81019061381b565b915f610777565b503d610a15565b6040513d86823e3d90fd5b634e487b7160e01b86526021600452602486fd5b50346101e45760203660031901126101e457600435610a71613ff1565b610a8060ff60025416156134ec565b60ff6008610a8d8361348a565b50015460a01c166005811015610571576004610aa99114613d61565b610ab28161348a565b503383526009602052604083208284526020526040832090600481015415610c39578154908115610c035760076002840191610af260ff84541615613dad565b01805490926001600160a01b039091169081610b82575050848080808654818115610b79575b3390f115610a35575b600160ff1982541617905560018060a01b039054169054916040519283527f1925c2d32a3486e44382eebcaea2888a167e91fde83a533b9ca611b0fc7931d260203394a460015f51602061468c5f395f51905f525580f35b506108fc610b18565b6040516370a0823160e01b815230600482015291602083602481845afa928315610bf8578893610bc2575b5061045e82610bbd94101561364c565b610b21565b92506020833d602011610bf0575b81610bdd60209383613585565b810103126104995791519161045e610bad565b3d9150610bd0565b6040513d8a823e3d90fd5b60405162461bcd60e51b815260206004820152600e60248201526d139bc81b195b9908185b5bdd5b9d60921b6044820152606490fd5b60405162461bcd60e51b815260206004820152600e60248201526d4e6f206c656e6420737570706c7960901b6044820152606490fd5b50346101e45760403660031901126101e4576040906001600160a01b03610c946134ba565b1681526009602052818120602435825260205220805490610cec6002600183015492015460405193849360ff808460081c16931691859260609295949195608085019685526020850152151560408401521515910152565b0390f35b50346101e457610cff36613474565b90610d08613ff1565b610d1760ff60025416156134ec565b60ff6008610d248361348a565b50015460a01c16600581101561110d5760021480156110e7575b610d4790613528565b610d508161348a565b5090610d5b816134d0565b50338552600960205260408520828652602052610d7d60408620541515613946565b83156110945760098301805490919086906001600160a01b0316803b156108a957604051632770a7eb60e21b8152336004820152602481018890529082908290604490829084905af1801561089e5761107b575b5050670de0b6b3a76400008502858104670de0b6b3a764000003611067578154610dfa916135e2565b60ff600886015460a01c16600581101561105357670de0b6b3a7640000929190600203610f4d576002610e3e92610e376001890154421015613698565b01546135bb565b6007909401549304926001600160a01b031680610eca57508480848015610ec0575b8280929181923390f115610a35575b60018060a01b039054169160405193845260208401527ff2ddf7a85d68bf37293e4fd7e4f19e0309db0714a4ad5251906341cc9ebb90ba60403394a45b60015f51602061468c5f395f51905f525580f35b6108fc9150610e60565b6040516370a0823160e01b8152306004820152602081602481855afa8015610f425785918891610f0b575b509161045e82610f0694101561364c565b610e6f565b9150506020813d602011610f3a575b81610f2760209383613585565b810103126104995751849061045e610ef5565b3d9150610f1a565b6040513d89823e3d90fd5b6004610f5f92610e3788544211613600565b6007909401549304926001600160a01b031680610fdb57508480848015610fd1575b8280929181923390f115610a35575b60018060a01b039054169160405193845260208401527ff2ddf7a85d68bf37293e4fd7e4f19e0309db0714a4ad5251906341cc9ebb90ba60403394a4610eac565b6108fc9150610f81565b6040516370a0823160e01b8152306004820152602081602481855afa8015610f42578591889161101c575b509161045e8261101794101561364c565b610f90565b9150506020813d60201161104b575b8161103860209383613585565b810103126104995751849061045e611006565b3d915061102b565b634e487b7160e01b88526021600452602488fd5b634e487b7160e01b87526011600452602487fd5b8161108591613585565b61109057855f610dd1565b8580fd5b60405162461bcd60e51b815260206004820152602560248201527f44657374726f7920616d6f756e74206d75737420626520677265617465722074604482015264068616e20360dc1b6064820152608490fd5b5060ff60086110f58361348a565b50015460a01c16600581101561110d57600314610d3e565b634e487b7160e01b84526021600452602484fd5b50346101e457806003193601126101e45761113a614592565b60025460ff808216158080157f989810a81181eea4877dd45b6d6eea562c0ad665780a7c2738e0ea75042876938680a3169060ff19161760025580f35b50346101e457806003193601126101e4576003546040516001600160a01b039091168152602090f35b50346101e45760203660031901126101e45760ff60086111c160043561348a565b50015460a01c169060058210156111dd57602082604051908152f35b634e487b7160e01b81526021600452602490fd5b50346101e45760203660031901126101e45760043561120e613ff1565b61121d60ff60025416156134ec565b6112296102a58261348a565b60ff60086112368361348a565b50015460a01c16600581101561057157151580611435575b6112579061387f565b6112608161348a565b509061126b816134d0565b50338452600a60205260408420828552602052604084209081549081159161129383156138da565b60016005870154920154906112a9828411613e38565b60028501936112bc60ff86541615613dad565b670de0b6b3a76400008202918204670de0b6b3a76400001417156114215782610383670de0b6b3a76400009361037d6008966112f7956135e2565b0494019160018060a01b0383541680155f146113a257508580868015611398575b8280929181923390f11561138d576001915b8260ff19825416179055016113408482546137b5565b905560018060a01b03905416906040519283527f445b80a962209ad38e06ec22380c7f6377277223b1ee8d229c2bcb60d80e3d5460203394a460015f51602061468c5f395f51905f525580f35b6040513d87823e3d90fd5b6108fc9150611318565b6040516370a0823160e01b81523060048201529290602084602481845afa8015610bf857879489916113e8575b50936113e39161045e82600197101561364c565b61132a565b9450506020843d602011611419575b8161140460209383613585565b810103126104995792518693906113e36113cf565b3d91506113f7565b634e487b7160e01b88526011600452602488fd5b5060ff60086114438361348a565b50015460a01c166005811015610571576004141561124e565b50346101e457806003193601126101e4576020600154604051908152f35b5061148436613474565b9061148d613ff1565b61149c60ff60025416156134ec565b6114b16114a88261348a565b505442106136db565b60ff60086114be8361348a565b50015460a01c16600581101561110d576114d8901561371d565b816114e28261348a565b50923385526009602052604085208386526020526040852090600785019260018060a01b0384541680155f1461160e57505050611520341515613769565b60015434106115be5760029061154960038601546115423460048901546137b5565b1115613dec565b600434955b6115598784546137b5565b8355016115678682546137b5565b90550161ffff19815416905560018060a01b03905416906040519283527f77c494147cc26c9ccb1f3f1926fb5bffa128633d935fdf7147dff0da6f740db460203394a460015f51602061468c5f395f51905f525580f35b60405162461bcd60e51b815260206004820152602260248201527f455448206d7573742062652067726561746572207468616e206d696e416d6f756044820152611b9d60f21b6064820152608490fd5b81939296911561171a5760015484106116c95761163760038301546115428660048601546137b5565b6040516370a0823160e01b8152336004820152602081602481855afa9081156116be578991611686575b50846004939261167860029761168194101561364c565b309033906140aa565b61154e565b90506020929192813d6020116116b6575b816116a460209383613585565b81010312610499575190919084611661565b3d9150611697565b6040513d8b823e3d90fd5b60405162461bcd60e51b8152602060048201526024808201527f4552433230206d7573742062652067726561746572207468616e206d696e416d6044820152631bdd5b9d60e21b6064820152608490fd5b60405162461bcd60e51b815260206004820152601c60248201527f4552433230206d7573742062652067726561746572207468616e2030000000006044820152606490fd5b50346101e457806003193601126101e457546040516001600160a01b039091168152602090f35b50346101e45760203660031901126101e4576004356117a3614592565b80156117d757806001547ffa6189b739625142c695478e9d0095a1cb9e6fad92ad8a727e0055a5cc85b06b8480a360015580f35b606460405162461bcd60e51b815260206004820152602060248201527f6d696e416d6f756e74206d7573742062652067726561746572207468616e20306044820152fd5b50346101e45760203660031901126101e4576004356001600160a01b038116908190036108a95761184a614592565b801561189357600454816001600160a01b0382167fd44190acf9d04bdb5d3a1aafff7e6dee8b40b93dfb8c5d3f0eea4b9f4539c3f78580a36001600160a01b0319161760045580f35b60405162461bcd60e51b815260206004820152602260248201527f4665654164647265737320616464726573732063616e6e6f7420626520656d70604482015261747960f01b6064820152608490fd5b50346101e45760203660031901126101e457600435611900613ff1565b61190f60ff60025416156134ec565b60ff600861191c8361348a565b50015460a01c1660058110156105715760046119389114613d61565b6119418161348a565b50338352600a602052604083208284526020526040832090600581015415611a8c578154906119718215156138da565b6008600284019161198660ff84541615613dad565b01805490926001600160a01b039091169081611a16575050848080808654818115611a0d575b3390f115610a35575b600160ff1982541617905560018060a01b039054169054916040519283527fded74dfac71815a7a7980ee897eb0a579034909191f76db76a03e2f17b92c27b60203394a460015f51602061468c5f395f51905f525580f35b506108fc6119ac565b6040516370a0823160e01b815230600482015291602083602481845afa928315610bf8578893611a56575b5061045e82611a5194101561364c565b6119b5565b92506020833d602011611a84575b81611a7160209383613585565b810103126104995791519161045e611a41565b3d9150611a64565b60405162461bcd60e51b815260206004820152601060248201526f4e6f20626f72726f7720737570706c7960801b6044820152606490fd5b50346101e457806003193601126101e45760025460405160089190911c6001600160a01b03168152602090f35b50346101e45760203660031901126101e457600435906008548210156101e45760c0611b1c836134d0565b5080549060018101549060028101546003820154906005600484015493015493604051958652602086015260408501526060840152608083015260a0820152f35b50346101e45760203660031901126101e457611b776134ba565b611b7f614592565b6001600160a01b038116908115611bea5760025491600883901c6001600160a01b03167fb7261e9c33aa7c56209c3bf60b424a8f9551ce28876c0ab3d0c487695e9434878580a3610100600160a81b031990911660089190911b610100600160a81b03161760025580f35b60405162461bcd60e51b815260206004820152601e60248201527f4f7261636c6520616464726573732063616e6e6f7420626520656d70747900006044820152606490fd5b50346101e457806003193601126101e457611c48614592565b80546001600160a01b03198116825581906001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a380f35b50346101e45760203660031901126101e45760206101da600435613d12565b50346101e457806003193601126101e457602060ff600254166040519015158152f35b50346101e457611cda36613474565b90611ce3614592565b80151580611e12575b15611da75781151580611d99575b15611d2e5781817f032dc6a2d839eb179729a55633fdf1c41a1fc4739394154117005db2b354b9b58580a360055560065580f35b60405162461bcd60e51b815260206004820152603e60248201527f626f72726f77466565206d7573742062652067726561746572207468616e203060448201527f20616e64206c657373207468616e206f7220657175616c20746f2031653800006064820152608490fd5b506305f5e100821115611cfa565b60405162461bcd60e51b815260206004820152603c60248201527f6c656e64466565206d7573742062652067726561746572207468616e2030206160448201527f6e64206c657373207468616e206f7220657175616c20746f20316538000000006064820152608490fd5b506305f5e100811115611cec565b50346101e45760403660031901126101e4576040906001600160a01b03611e456134ba565b168152600a602052818120602435825260205220805490610cec6002600183015492015460405193849360ff808460081c16931691859260609295949195608085019685526020850152151560408401521515910152565b50346101e45760203660031901126101e45760206101da600435613cd5565b50346101e45760203660031901126101e457611ed6614592565b611ee160043561397d565b80f35b50346101e45760203660031901126101e457600435611f01613ff1565b611f1060ff60025416156134ec565b611f1c6102a58261348a565b60ff6008611f298361348a565b50015460a01c166005811015610571571515806120dc575b611f4a9061387f565b611f538161348a565b5090611f5e816134d0565b50338452600960205260408420828552602052604084209081546002811593611f878515613946565b019260ff845460081c166120a857670de0b6b3a76400008202918204670de0b6b3a764000014171561209457611fd8600992611fd1670de0b6b3a7640000936004890154906135e2565b90546135bb565b049301908460018060a01b03835416803b156108a9576040516340c10f1960e01b8152336004820152602481018790529082908290604490829084905af1801561089e5761207b575b505061010061ff001982541617905560018060a01b03905416906040519283527f535481ce60a76f235eeee0e6601988d1031b55eef7e85a043cd7a861f94fb59f60203394a460015f51602061468c5f395f51905f525580f35b8161208591613585565b61209057845f612021565b8480fd5b634e487b7160e01b86526011600452602486fd5b60405162461bcd60e51b815260206004820152600c60248201526b486173206e6f20636c61696d60a01b6044820152606490fd5b5060ff60086120ea8361348a565b50015460a01c1660058110156105715760041415611f41565b50346101e457806003193601126101e4576020600554604051908152f35b50346101e4576101403660031901126101e45761213c614592565b6004359081156124f6576024359182156124b1578083111561245c576001600160a01b03612168613919565b1615612409576001600160a01b0361217e61392f565b16156123b45760843580156123635760443580156123125760075494600160401b8610156122fe57600186016007556121b68661348a565b50506121c18661348a565b50938455600184015560028301556064356003830155600682015560a4356001600160a01b03811681036122fa576007820180546001600160a01b0319166001600160a01b0392831617905560c43590811681036122fa576008820180546001600160a81b0319166001600160a01b03909216919091179055612242613919565b6009820180546001600160a01b0319166001600160a01b0390921691909117905561226b61392f565b600a820180546001600160a01b0319166001600160a01b0390921691909117905561012435600b9091015560085491600160401b8310156122e6576122b8836001602095016008556134d0565b5050807f67d5edbe9035b9e077288ae99fef26dd09577336cade83fbf1270ad01af587f96040519380a28152f35b634e487b7160e01b82526041600452602482fd5b8280fd5b634e487b7160e01b85526041600452602485fd5b60405162461bcd60e51b815260206004820152602360248201527f496e74657265737452617465206d75737420626520677265617465722074686160448201526206e20360ec1b6064820152608490fd5b60405162461bcd60e51b815260206004820152602360248201527f4d6f72746761676552617465206d75737420626520677265617465722074686160448201526206e20360ec1b6064820152608490fd5b60405162461bcd60e51b815260206004820152602760248201527f426f72726f7744656274546f6b656e20616464726573732063616e6e6f7420626044820152666520656d70747960c81b6064820152608490fd5b60405162461bcd60e51b815260206004820152602560248201527f4c656e6444656274546f6b656e20616464726573732063616e6e6f7420626520604482015264656d70747960d81b6064820152608490fd5b60405162461bcd60e51b815260206004820152602760248201527f456e6454696d65206d7573742062652067726561746572207468616e20536574604482015266746c6554696d6560c81b6064820152608490fd5b60405162461bcd60e51b815260206004820152601e60248201527f456e6454696d65206d7573742062652067726561746572207468616e203000006044820152606490fd5b60405162461bcd60e51b815260206004820152602160248201527f536574746c6554696d65206d7573742062652067726561746572207468616e206044820152600360fc1b6064820152608490fd5b50346101e457806003193601126101e4576004546040516001600160a01b039091168152602090f35b50346101e45760203660031901126101e4576125886134ba565b612590614592565b6001600160a01b031680156125e257600354816001600160a01b0382167f4558149b3c5427365f76d4ff19bef30aba41f17e5e601d4661330d8d2b6876278580a36001600160a01b0319161760035580f35b60405162461bcd60e51b815260206004820152602260248201527f53776170526f7574657220616464726573732063616e6e6f7420626520656d70604482015261747960f01b6064820152608490fd5b50346101e45760203660031901126101e45760043561264f613ff1565b61265e60ff60025416156134ec565b61266a6102a58261348a565b60ff60086126778361348a565b50015460a01c166005811015610571571515806128f8575b6126989061387f565b6126a18161348a565b506126ab826134d0565b5091338452600a602052604084208185526020526040842092835460028115956126d587156138da565b019060ff825460081c166128c1576305f5e1006126f884546006880154906135bb565b0495670de0b6b3a76400008202918204670de0b6b3a76400001417156110675761273a612733670de0b6b3a7640000926005880154906135e2565b80976135bb565b0494600a8501928760018060a01b03855416803b156108a9576040516340c10f1960e01b8152336004820152602481018a90529082908290604490829084905af1801561089e576128ac575b5050670de0b6b3a76400009161279c91546135bb565b6007909501549404936001600160a01b0316806128345750858085801561282a575b8280929181923390f11561138d575b61010061ff001982541617905560018060a01b039054169160405193845260208401527f1226b632b60c533dd17d143c6f2616eb2415a03ce2ce3168a8311c8787e2555660403394a460015f51602061468c5f395f51905f525580f35b6108fc91506127be565b6040516370a0823160e01b8152306004820152602081602481855afa8015610bf85786918991612875575b509161045e8261287094101561364c565b6127cd565b9150506020813d6020116128a4575b8161289160209383613585565b810103126104995751859061045e61285f565b3d9150612884565b816128b691613585565b61089a57875f612786565b60405162461bcd60e51b815260206004820152600f60248201526e105b1c9958591e4818db185a5b5959608a1b6044820152606490fd5b5060ff60086129068361348a565b50015460a01c166005811015610571576004141561268f565b50346101e45760203660031901126101e4576004356007548110156108a9576129479061348a565b508054906001810154906002810154600382015460048301546005840154600685015460018060a01b036007870154169160088701549360ff8560a01c169560018060a01b0360098a01541697600b60018060a01b03600a8c0154169a01549a6040519c8d5260208d015260408c015260608b015260808a015260a089015260c088015260e087015260018060a01b03166101008601526005811015610a40576101a09550610120850152610140840152610160830152610180820152f35b50346101e45760203660031901126101e457600435612a23613ff1565b612a3260ff60025416156134ec565b612a3b8161348a565b50600881019081549160ff8360a01c166005811015610a40576001612a6091146137c2565b60018201549182421015612e2757612a7d61061961060f876134d0565b936305f5e100612a8f600554876135bb565b0494876007612a9e88846137b5565b940180549094936001600160a01b039182169116818103612b7c57505050600485015560018401547f40a58cc34fb5788d695e1e808a735bf11c933c79a9c7af0726d365a2d10d29d3956040959493929091818110612b7157612b04916106c39161380e565b600585015580612b4e575b50505b805460ff60a01b1916600360a01b17905560048101546005919091015482519182526020820152a260015f51602061468c5f395f51905f525580f35b9054600454612b6a92916001600160a01b039182169116614505565b5f80612b0f565b5050612b04886106c3565b600354604051633fc8cef360e01b8152959695929190602090849060049082906001600160a01b03165afa928315610a35578493612e06575b5080612e0057825b82612dfa57835b6001600160a01b03918216911603612cc457505085546001600160a01b031615612c72575b5050907f40a58cc34fb5788d695e1e808a735bf11c933c79a9c7af0726d365a2d10d29d395859493926004604097015580612c4f575b50506001830154818110612c4457612c3a916107fc9161380e565b6005830155612b12565b5050612c3a866107fc565b9054600454612c6b92916001600160a01b039182169116614505565b5f80612c1f565b6001600160a01b0316803b156108a957818591600460405180948193630d0e30db60e41b83525af1801561089e5715612be95781612cb591979695949397613585565b61089a5790919293875f612be9565b80949950612cd99350819598979692506141e2565b91600185019182548411612d9f57612cf29184916142a5565b868110612d5b577f40a58cc34fb5788d695e1e808a735bf11c933c79a9c7af0726d365a2d10d29d39660409681831115612d4f5761091d82612d34939461380e565b60048501555b54818110612c4457612c3a916107fc9161380e565b50506004850155612d3a565b606460405162461bcd60e51b815260206004820152602060248201527f6c6971756964617465506f6f6c3a20536c69707061676520746f6f20686967686044820152fd5b60405162461bcd60e51b815260206004820152602d60248201527f6c6971756964617465506f6f6c3a20696e73756666696369656e7420626f727260448201526c1bddc818dbdb1b185d195c985b609a1b6064820152608490fd5b82612bc4565b80612bbd565b612e2091935060203d602011610a2e57610a1f8183613585565b915f612bb5565b60405162461bcd60e51b815260206004820152602260248201527f506f6f6c20616c726561647920656e6465642c207573652066696e697368506f6044820152611bdb60f21b6064820152608490fd5b50612e8136613474565b90612e8a613ff1565b612e9960ff60025416156134ec565b612ea56114a88261348a565b60ff6008612eb28361348a565b50015460a01c16600581101561110d57612ecc901561371d565b81612ed68261348a565b5092338552600a602052604085208386526020526040852090600885019260018060a01b0384541680155f14612f8c57505050600290612f17341515613769565b600534955b612f278784546137b5565b835501612f358682546137b5565b90550161ffff19815416905560018060a01b03905416906040519283527ff9a33434428db5f0416c03e38307599ad0b9b9965d6c070eb08e87cc1f0ca50e60203394a460015f51602061468c5f395f51905f525580f35b819392969115613014576040516370a0823160e01b8152336004820152602081602481855afa9081156116be578991612fdc575b508460059392611678600297612fd794101561364c565b612f1c565b90506020929192813d60201161300c575b81612ffa60209383613585565b81010312610499575190919084612fc0565b3d9150612fed565b60405162461bcd60e51b815260206004820152602860248201527f626f72726f77546f6b656e416d6f756e74206d75737420626520677265617465604482015267072207468616e20360c41b6064820152608490fd5b50346104995761307936613474565b90613082613ff1565b61309160ff60025416156134ec565b60ff600861309e8361348a565b50015460a01c16600581101561346057600214801561343a575b6130c190613528565b6130ca8161348a565b50906130d5816134d0565b5083156133e057600a8301546001600160a01b0316803b1561049957604051632770a7eb60e21b815233600482015260248101869052905f908290604490829084905af180156133d5576133c0575b506305f5e10061313a82546006860154906135bb565b04670de0b6b3a76400008502858104670de0b6b3a7640000036110675790613161916135e2565b906008840193845460ff8160a01c1660058110156133ac576002036132a257509160036131a192610e376001670de0b6b3a7640000960154421015613698565b835491900492906001600160a01b03168061322a57508480848015613220575b8280929181923390f115610a35575b60018060a01b039054169160405193845260208401527f6523113a4bcd4da7dc9c771b759cc100a53e2925903ef0789e7be97aa234160a60403394a460015f51602061468c5f395f51905f525580f35b6108fc91506131c1565b6040516370a0823160e01b8152306004820152602081602481855afa8015610f42578591889161326b575b509161045e8261326694101561364c565b6131d0565b9150506020813d60201161329a575b8161328760209383613585565b810103126104995751849061045e613255565b3d915061327a565b94926005670de0b6b3a764000093610e376132bf94544211613600565b04926001600160a01b0316806133345750848084801561332a575b8280929181923390f115610a35575b60018060a01b039054169160405193845260208401527f6523113a4bcd4da7dc9c771b759cc100a53e2925903ef0789e7be97aa234160a60403394a4610eac565b6108fc91506132da565b6040516370a0823160e01b8152306004820152602081602481855afa8015610f425785918891613375575b509161045e8261337094101561364c565b6132e9565b9150506020813d6020116133a4575b8161339160209383613585565b810103126104995751849061045e61335f565b3d9150613384565b634e487b7160e01b89526021600452602489fd5b6133cd9195505f90613585565b5f935f613124565b6040513d5f823e3d90fd5b60405162461bcd60e51b815260206004820152602c60248201527f626f72726f7744656274546f6b656e416d6f756e74206d75737420626520677260448201526b06561746572207468616e20360a41b6064820152608490fd5b5060ff60086134488361348a565b50015460a01c166005811015613460576003146130b8565b634e487b7160e01b5f52602160045260245ffd5b6040906003190112610499576004359060243590565b6007548110156134a65760075f52600c60205f20910201905f90565b634e487b7160e01b5f52603260045260245ffd5b600435906001600160a01b038216820361049957565b6008548110156134a65760085f52600660205f20910201905f90565b156134f357565b60405162461bcd60e51b815260206004820152600d60248201526c11db1bd8985b0814185d5cd959609a1b6044820152606490fd5b1561352f57565b60405162461bcd60e51b815260206004820152602860248201527f506f6f6c207374617465206d7573742062652046494e495348206f72204c49516044820152672aa4a220aa24a7a760c11b6064820152608490fd5b90601f8019910116810190811067ffffffffffffffff8211176135a757604052565b634e487b7160e01b5f52604160045260245ffd5b818102929181159184041417156135ce57565b634e487b7160e01b5f52601160045260245ffd5b81156135ec570490565b634e487b7160e01b5f52601260045260245ffd5b1561360757565b60405162461bcd60e51b815260206004820152601760248201527f4e6f74207265616368656420736574746c652074696d650000000000000000006044820152606490fd5b1561365357565b60405162461bcd60e51b815260206004820152601b60248201527f45524332302062616c616e6365206973206e6f7420656e6f75676800000000006044820152606490fd5b1561369f57565b60405162461bcd60e51b81526020600482015260146024820152734e6f74207265616368656420656e642074696d6560601b6044820152606490fd5b156136e257565b60405162461bcd60e51b81526020600482015260136024820152724c657373207468616e20746869732074696d6560681b6044820152606490fd5b1561372457565b60405162461bcd60e51b815260206004820152601860248201527f506f6f6c207374617465206d757374206265204d4154434800000000000000006044820152606490fd5b1561377057565b60405162461bcd60e51b815260206004820152601a60248201527f455448206d7573742062652067726561746572207468616e20300000000000006044820152606490fd5b919082018092116135ce57565b156137c957565b60405162461bcd60e51b815260206004820152601c60248201527f506f6f6c207374617465206d75737420626520455845435554494f4e000000006044820152606490fd5b919082039182116135ce57565b9081602091031261049957516001600160a01b03811681036104995790565b1561384157565b60405162461bcd60e51b815260206004820152601660248201527547726561746572207468616e20746869732074696d6560501b6044820152606490fd5b1561388657565b60405162461bcd60e51b815260206004820152602660248201527f506f6f6c207374617465206d757374206e6f74206265204d41544348206f7220604482015265554e444f4e4560d01b6064820152608490fd5b156138e157565b60405162461bcd60e51b815260206004820152601060248201526f139bc8189bdc9c9bddc8185b5bdd5b9d60821b6044820152606490fd5b60e4356001600160a01b03811681036104995790565b610104356001600160a01b03811681036104995790565b1561394d57565b60405162461bcd60e51b8152602060048201526008602482015267139bdd081b195b9960c21b6044820152606490fd5b6139868161348a565b506008810160ff815460a01c166005811015613460576139a6901561371d565b6139b38254421015613600565b600482018054158015613cc9575b613cb357600783015482546001600160a01b03908116929116828103613aab575090506005830154926305f5e1008402938085046305f5e10014901517156135ce576305f5e10091613a1b6006613a2e93015480966135e2565b905490808211613aa35750935b846135bb565b04905b8115613a8d5791817fecb83b282735569068a276ac3b9f868bdb0048d952568d94f8a79a83eb001bbe936040936001613a69886134d0565b508581550155805460ff60a01b1916600160a01b17905582519182526020820152a2565b805460ff60a01b1916600160a21b179055505050565b905093613a28565b6002546040516341976e0960e01b8152600481019490945260081c6001600160a01b031690602084602481855afa9384156133d5575f94613c7e575b506020906024604051809481936341976e0960e01b835260048301525afa9081156133d5575f91613c4c575b508215613bf8578015613ba757613b2e8360058701546135bb565b946305f5e1008602958087046305f5e10014901517156135ce57613b7b92613b68600661038393015497613b6289866135bb565b906135e2565b905490808211613b9f5750955b866135bb565b90806305f5e10002906305f5e1008204036135ce57613b99916135e2565b90613a31565b905095613b75565b60405162461bcd60e51b8152602060048201526024808201527f736574746c65506f6f6c3a206c656e6420746f6b656e207072696365206973206044820152637a65726f60e01b6064820152608490fd5b60405162461bcd60e51b815260206004820152602660248201527f736574746c65506f6f6c3a20626f72726f7720746f6b656e207072696365206960448201526573207a65726f60d01b6064820152608490fd5b90506020813d602011613c76575b81613c6760209383613585565b8101031261049957515f613b13565b3d9150613c5a565b9093506020813d602011613cab575b81613c9a60209383613585565b810103126104995751926020613ae7565b3d9150613c8d565b50805460ff60a01b1916600160a21b1790555050565b506005830154156139c1565b613cde9061348a565b5060ff600882015460a01c16600581101561346057600103613d0d57600101544210613d0957600190565b5f90565b505f90565b613d1b9061348a565b5060ff600882015460a01c16600581101561346057613d0d5780544210613d0d57600481015415908115613d54575b50613d0957600190565b600591500154155f613d4a565b15613d6857565b60405162461bcd60e51b815260206004820152601960248201527f506f6f6c207374617465206d75737420626520554e444f4e45000000000000006044820152606490fd5b15613db457565b60405162461bcd60e51b815260206004820152601060248201526f105b1c9958591e481c99599d5b99195960821b6044820152606490fd5b15613df357565b60405162461bcd60e51b815260206004820152601960248201527f4578636565647320746865206d6178696d756d206c696d6974000000000000006044820152606490fd5b15613e3f57565b60405162461bcd60e51b815260206004820152601060248201526f139bc81c99599d5b9908185b5bdd5b9d60821b6044820152606490fd5b613e808161348a565b5090600882015460ff8160a01c16600581101561346057600103613fea576001830154421015613fea5760025460078401546040516341976e0960e01b81526001600160a01b0391821660048201529260089290921c1690602083602481855afa9283156133d5575f93613fb5575b506040516341976e0960e01b81526001600160a01b03909116600482015290602090829060249082905afa9081156133d5575f91613f81575b50613f5291613f47613f4c926001613f3f876134d0565b5001546135bb565b6135e2565b916134d0565b5054906305f5e1008102908082046305f5e10014901517156135ce57600b91613f7a916135e2565b9101541190565b90506020813d602011613fad575b81613f9c60209383613585565b810103126104995751613f52613f28565b3d9150613f8f565b9092506020813d602011613fe2575b81613fd160209383613585565b810103126104995751916020613eef565b3d9150613fc4565b5050505f90565b60025f51602061468c5f395f51905f52541461401a5760025f51602061468c5f395f51905f5255565b633ee5aeb560e01b5f5260045ffd5b916040519163a9059cbb60e01b5f5260018060a01b031660045260245260205f60448180865af19060015f5114821615614089575b604052156140695750565b635274afe760e01b5f9081526001600160a01b0391909116600452602490fd5b9060018115166140a157823b15153d1516169061405e565b503d5f823e3d90fd5b6040516323b872dd60e01b5f9081526001600160a01b039384166004529290931660245260449390935260209060648180865af19060015f51148216156140fb575b6040525f606052156140695750565b9060018115166140a157823b15153d151616906140ec565b6020818303126104995780519067ffffffffffffffff821161049957019080601f830112156104995781519167ffffffffffffffff83116135a7578260051b9060208201936141656040519586613585565b845260208085019282010192831161049957602001905b8282106141895750505090565b815181526020918201910161417c565b90602080835192838152019201905f5b8181106141b65750505090565b82516001600160a01b03168452602093840193909201916001016141a9565b8051156134a65760200190565b916141f15f9261422b946145b8565b600354604080516307c0329d60e21b815260048101949094526024840152919384926001600160a01b031691839182916044830190614199565b03915afa80156133d557614246915f9161424a575b506141d5565b5190565b61426691503d805f833e61425e8183613585565b810190614113565b5f614240565b90608092614291919695949683525f602084015260a0604084015260a0830190614199565b6001600160a01b0390951660608201520152565b6003546001600160a01b03169290916142be81846145b8565b926001600160a01b03168061436957505061012c4201928342116135ce575f9261430c92604051809681958294637ff36ab560e01b8452886004850152608060248501526084840190614199565b90306044840152606483015203925af19081156133d5575f9161434f575b505b80515f198101919082116135ce5780518210156134a65760209160051b01015190565b61436391503d805f833e61425e8183613585565b5f61432a565b5f9492945060405163095ea7b360e01b5f52836004528560245260205f60448180865af19060015f51148216156144f6575b6040521561444a575b506001600160a01b03166144175761012c4201908142116135ce576143e5935f8094604051968795869485936318cbafe560e01b855230916004860161426c565b03925af19081156133d5575f916143fd575b5061432c565b61441191503d805f833e61425e8183613585565b5f6143f7565b61012c4201908142116135ce576143e5935f8094604051968795869485936338ed173960e01b855230916004860161426c565b60405163095ea7b360e01b5f52836004525f60245260205f60448180865af19060015f51148216156144de575b604052156144b45760405163095ea7b360e01b5f52836004528560245260205f60448180865af19060015f51148216156144c6575b6040526143a4575b635274afe760e01b5f5260045260245ffd5b9060018115166140a157823b15153d151616906144ac565b9060018115166140a157823b15153d15161690614477565b90823b15153d1516169061439b565b6001600160a01b0316919082614542575f8093508092819282908215614538575b6001600160a01b031690f1156133d557565b6108fc9150614526565b61454b92614029565b565b916305f5e10061456061457094836135bb565b04918280614573575b505061380e565b90565b60045461458b926001600160a01b0390911690614505565b5f82614569565b5f546001600160a01b031633036145a557565b63118cdaa760e01b5f523360045260245ffd5b600354604051633fc8cef360e01b81529193929190602090829060049082906001600160a01b03165afa9081156133d5575f9161466c575b5060405191614600606084613585565b60028352604036602085013791938492906001600160a01b0381166146675750815b61462b846141d5565b6001600160a01b039182169052811661465f5750905b8051600110156134a6576001600160a01b0390911660409190910152565b905090614641565b614622565b614685915060203d602011610a2e57610a1f8183613585565b5f6145f056fe9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00a26469706673582212205d794c34f9ab6e8120953cd74793cafb104e58bc8700e1b5e95cf77a209548a664736f6c634300081c0033",
}

// PledgePoolABI is the input ABI used to generate the binding from.
// Deprecated: Use PledgePoolMetaData.ABI instead.
var PledgePoolABI = PledgePoolMetaData.ABI

// PledgePoolBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PledgePoolMetaData.Bin instead.
var PledgePoolBin = PledgePoolMetaData.Bin

// DeployPledgePool deploys a new Ethereum contract, binding an instance of PledgePool to it.
func DeployPledgePool(auth *bind.TransactOpts, backend bind.ContractBackend, _oracle common.Address, _swapRouter common.Address, _feeAddress common.Address, _owner common.Address) (common.Address, *types.Transaction, *PledgePool, error) {
	parsed, err := PledgePoolMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PledgePoolBin), backend, _oracle, _swapRouter, _feeAddress, _owner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PledgePool{PledgePoolCaller: PledgePoolCaller{contract: contract}, PledgePoolTransactor: PledgePoolTransactor{contract: contract}, PledgePoolFilterer: PledgePoolFilterer{contract: contract}}, nil
}

// PledgePool is an auto generated Go binding around an Ethereum contract.
type PledgePool struct {
	PledgePoolCaller     // Read-only binding to the contract
	PledgePoolTransactor // Write-only binding to the contract
	PledgePoolFilterer   // Log filterer for contract events
}

// PledgePoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type PledgePoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PledgePoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PledgePoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PledgePoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PledgePoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PledgePoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PledgePoolSession struct {
	Contract     *PledgePool       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PledgePoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PledgePoolCallerSession struct {
	Contract *PledgePoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// PledgePoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PledgePoolTransactorSession struct {
	Contract     *PledgePoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// PledgePoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type PledgePoolRaw struct {
	Contract *PledgePool // Generic contract binding to access the raw methods on
}

// PledgePoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PledgePoolCallerRaw struct {
	Contract *PledgePoolCaller // Generic read-only contract binding to access the raw methods on
}

// PledgePoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PledgePoolTransactorRaw struct {
	Contract *PledgePoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPledgePool creates a new instance of PledgePool, bound to a specific deployed contract.
func NewPledgePool(address common.Address, backend bind.ContractBackend) (*PledgePool, error) {
	contract, err := bindPledgePool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PledgePool{PledgePoolCaller: PledgePoolCaller{contract: contract}, PledgePoolTransactor: PledgePoolTransactor{contract: contract}, PledgePoolFilterer: PledgePoolFilterer{contract: contract}}, nil
}

// NewPledgePoolCaller creates a new read-only instance of PledgePool, bound to a specific deployed contract.
func NewPledgePoolCaller(address common.Address, caller bind.ContractCaller) (*PledgePoolCaller, error) {
	contract, err := bindPledgePool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PledgePoolCaller{contract: contract}, nil
}

// NewPledgePoolTransactor creates a new write-only instance of PledgePool, bound to a specific deployed contract.
func NewPledgePoolTransactor(address common.Address, transactor bind.ContractTransactor) (*PledgePoolTransactor, error) {
	contract, err := bindPledgePool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PledgePoolTransactor{contract: contract}, nil
}

// NewPledgePoolFilterer creates a new log filterer instance of PledgePool, bound to a specific deployed contract.
func NewPledgePoolFilterer(address common.Address, filterer bind.ContractFilterer) (*PledgePoolFilterer, error) {
	contract, err := bindPledgePool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PledgePoolFilterer{contract: contract}, nil
}

// bindPledgePool binds a generic wrapper to an already deployed contract.
func bindPledgePool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PledgePoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PledgePool *PledgePoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PledgePool.Contract.PledgePoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PledgePool *PledgePoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PledgePool.Contract.PledgePoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PledgePool *PledgePoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PledgePool.Contract.PledgePoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PledgePool *PledgePoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PledgePool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PledgePool *PledgePoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PledgePool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PledgePool *PledgePoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PledgePool.Contract.contract.Transact(opts, method, params...)
}

// BorrowFee is a free data retrieval call binding the contract method 0xe626648a.
//
// Solidity: function borrowFee() view returns(uint256)
func (_PledgePool *PledgePoolCaller) BorrowFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PledgePool.contract.Call(opts, &out, "borrowFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowFee is a free data retrieval call binding the contract method 0xe626648a.
//
// Solidity: function borrowFee() view returns(uint256)
func (_PledgePool *PledgePoolSession) BorrowFee() (*big.Int, error) {
	return _PledgePool.Contract.BorrowFee(&_PledgePool.CallOpts)
}

// BorrowFee is a free data retrieval call binding the contract method 0xe626648a.
//
// Solidity: function borrowFee() view returns(uint256)
func (_PledgePool *PledgePoolCallerSession) BorrowFee() (*big.Int, error) {
	return _PledgePool.Contract.BorrowFee(&_PledgePool.CallOpts)
}

// BorrowInfoMap is a free data retrieval call binding the contract method 0x51dbf7f9.
//
// Solidity: function borrowInfoMap(address , uint256 ) view returns(uint256 borrowAmount, uint256 returnBorrowAmount, bool hasNoRefund, bool hasNoClaim)
func (_PledgePool *PledgePoolCaller) BorrowInfoMap(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	BorrowAmount       *big.Int
	ReturnBorrowAmount *big.Int
	HasNoRefund        bool
	HasNoClaim         bool
}, error) {
	var out []interface{}
	err := _PledgePool.contract.Call(opts, &out, "borrowInfoMap", arg0, arg1)

	outstruct := new(struct {
		BorrowAmount       *big.Int
		ReturnBorrowAmount *big.Int
		HasNoRefund        bool
		HasNoClaim         bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BorrowAmount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ReturnBorrowAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.HasNoRefund = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.HasNoClaim = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// BorrowInfoMap is a free data retrieval call binding the contract method 0x51dbf7f9.
//
// Solidity: function borrowInfoMap(address , uint256 ) view returns(uint256 borrowAmount, uint256 returnBorrowAmount, bool hasNoRefund, bool hasNoClaim)
func (_PledgePool *PledgePoolSession) BorrowInfoMap(arg0 common.Address, arg1 *big.Int) (struct {
	BorrowAmount       *big.Int
	ReturnBorrowAmount *big.Int
	HasNoRefund        bool
	HasNoClaim         bool
}, error) {
	return _PledgePool.Contract.BorrowInfoMap(&_PledgePool.CallOpts, arg0, arg1)
}

// BorrowInfoMap is a free data retrieval call binding the contract method 0x51dbf7f9.
//
// Solidity: function borrowInfoMap(address , uint256 ) view returns(uint256 borrowAmount, uint256 returnBorrowAmount, bool hasNoRefund, bool hasNoClaim)
func (_PledgePool *PledgePoolCallerSession) BorrowInfoMap(arg0 common.Address, arg1 *big.Int) (struct {
	BorrowAmount       *big.Int
	ReturnBorrowAmount *big.Int
	HasNoRefund        bool
	HasNoClaim         bool
}, error) {
	return _PledgePool.Contract.BorrowInfoMap(&_PledgePool.CallOpts, arg0, arg1)
}

// CheckCanFinish is a free data retrieval call binding the contract method 0x50e0a6b8.
//
// Solidity: function checkCanFinish(uint256 _pid) view returns(bool)
func (_PledgePool *PledgePoolCaller) CheckCanFinish(opts *bind.CallOpts, _pid *big.Int) (bool, error) {
	var out []interface{}
	err := _PledgePool.contract.Call(opts, &out, "checkCanFinish", _pid)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckCanFinish is a free data retrieval call binding the contract method 0x50e0a6b8.
//
// Solidity: function checkCanFinish(uint256 _pid) view returns(bool)
func (_PledgePool *PledgePoolSession) CheckCanFinish(_pid *big.Int) (bool, error) {
	return _PledgePool.Contract.CheckCanFinish(&_PledgePool.CallOpts, _pid)
}

// CheckCanFinish is a free data retrieval call binding the contract method 0x50e0a6b8.
//
// Solidity: function checkCanFinish(uint256 _pid) view returns(bool)
func (_PledgePool *PledgePoolCallerSession) CheckCanFinish(_pid *big.Int) (bool, error) {
	return _PledgePool.Contract.CheckCanFinish(&_PledgePool.CallOpts, _pid)
}

// CheckCanLiquidate is a free data retrieval call binding the contract method 0xff920635.
//
// Solidity: function checkCanLiquidate(uint256 _pid) view returns(bool)
func (_PledgePool *PledgePoolCaller) CheckCanLiquidate(opts *bind.CallOpts, _pid *big.Int) (bool, error) {
	var out []interface{}
	err := _PledgePool.contract.Call(opts, &out, "checkCanLiquidate", _pid)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckCanLiquidate is a free data retrieval call binding the contract method 0xff920635.
//
// Solidity: function checkCanLiquidate(uint256 _pid) view returns(bool)
func (_PledgePool *PledgePoolSession) CheckCanLiquidate(_pid *big.Int) (bool, error) {
	return _PledgePool.Contract.CheckCanLiquidate(&_PledgePool.CallOpts, _pid)
}

// CheckCanLiquidate is a free data retrieval call binding the contract method 0xff920635.
//
// Solidity: function checkCanLiquidate(uint256 _pid) view returns(bool)
func (_PledgePool *PledgePoolCallerSession) CheckCanLiquidate(_pid *big.Int) (bool, error) {
	return _PledgePool.Contract.CheckCanLiquidate(&_PledgePool.CallOpts, _pid)
}

// CheckCanSettle is a free data retrieval call binding the contract method 0x6f68d40e.
//
// Solidity: function checkCanSettle(uint256 _pid) view returns(bool)
func (_PledgePool *PledgePoolCaller) CheckCanSettle(opts *bind.CallOpts, _pid *big.Int) (bool, error) {
	var out []interface{}
	err := _PledgePool.contract.Call(opts, &out, "checkCanSettle", _pid)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckCanSettle is a free data retrieval call binding the contract method 0x6f68d40e.
//
// Solidity: function checkCanSettle(uint256 _pid) view returns(bool)
func (_PledgePool *PledgePoolSession) CheckCanSettle(_pid *big.Int) (bool, error) {
	return _PledgePool.Contract.CheckCanSettle(&_PledgePool.CallOpts, _pid)
}

// CheckCanSettle is a free data retrieval call binding the contract method 0x6f68d40e.
//
// Solidity: function checkCanSettle(uint256 _pid) view returns(bool)
func (_PledgePool *PledgePoolCallerSession) CheckCanSettle(_pid *big.Int) (bool, error) {
	return _PledgePool.Contract.CheckCanSettle(&_PledgePool.CallOpts, _pid)
}

// FeeAddress is a free data retrieval call binding the contract method 0x41275358.
//
// Solidity: function feeAddress() view returns(address)
func (_PledgePool *PledgePoolCaller) FeeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PledgePool.contract.Call(opts, &out, "feeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeAddress is a free data retrieval call binding the contract method 0x41275358.
//
// Solidity: function feeAddress() view returns(address)
func (_PledgePool *PledgePoolSession) FeeAddress() (common.Address, error) {
	return _PledgePool.Contract.FeeAddress(&_PledgePool.CallOpts)
}

// FeeAddress is a free data retrieval call binding the contract method 0x41275358.
//
// Solidity: function feeAddress() view returns(address)
func (_PledgePool *PledgePoolCallerSession) FeeAddress() (common.Address, error) {
	return _PledgePool.Contract.FeeAddress(&_PledgePool.CallOpts)
}

// GetPoolState is a free data retrieval call binding the contract method 0xb1597517.
//
// Solidity: function getPoolState(uint256 _pid) view returns(uint256)
func (_PledgePool *PledgePoolCaller) GetPoolState(opts *bind.CallOpts, _pid *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PledgePool.contract.Call(opts, &out, "getPoolState", _pid)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPoolState is a free data retrieval call binding the contract method 0xb1597517.
//
// Solidity: function getPoolState(uint256 _pid) view returns(uint256)
func (_PledgePool *PledgePoolSession) GetPoolState(_pid *big.Int) (*big.Int, error) {
	return _PledgePool.Contract.GetPoolState(&_PledgePool.CallOpts, _pid)
}

// GetPoolState is a free data retrieval call binding the contract method 0xb1597517.
//
// Solidity: function getPoolState(uint256 _pid) view returns(uint256)
func (_PledgePool *PledgePoolCallerSession) GetPoolState(_pid *big.Int) (*big.Int, error) {
	return _PledgePool.Contract.GetPoolState(&_PledgePool.CallOpts, _pid)
}

// GlobalPaused is a free data retrieval call binding the contract method 0x61a552dc.
//
// Solidity: function globalPaused() view returns(bool)
func (_PledgePool *PledgePoolCaller) GlobalPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PledgePool.contract.Call(opts, &out, "globalPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GlobalPaused is a free data retrieval call binding the contract method 0x61a552dc.
//
// Solidity: function globalPaused() view returns(bool)
func (_PledgePool *PledgePoolSession) GlobalPaused() (bool, error) {
	return _PledgePool.Contract.GlobalPaused(&_PledgePool.CallOpts)
}

// GlobalPaused is a free data retrieval call binding the contract method 0x61a552dc.
//
// Solidity: function globalPaused() view returns(bool)
func (_PledgePool *PledgePoolCallerSession) GlobalPaused() (bool, error) {
	return _PledgePool.Contract.GlobalPaused(&_PledgePool.CallOpts)
}

// LendFee is a free data retrieval call binding the contract method 0x4aea0aec.
//
// Solidity: function lendFee() view returns(uint256)
func (_PledgePool *PledgePoolCaller) LendFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PledgePool.contract.Call(opts, &out, "lendFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LendFee is a free data retrieval call binding the contract method 0x4aea0aec.
//
// Solidity: function lendFee() view returns(uint256)
func (_PledgePool *PledgePoolSession) LendFee() (*big.Int, error) {
	return _PledgePool.Contract.LendFee(&_PledgePool.CallOpts)
}

// LendFee is a free data retrieval call binding the contract method 0x4aea0aec.
//
// Solidity: function lendFee() view returns(uint256)
func (_PledgePool *PledgePoolCallerSession) LendFee() (*big.Int, error) {
	return _PledgePool.Contract.LendFee(&_PledgePool.CallOpts)
}

// LendInfoMap is a free data retrieval call binding the contract method 0xdcfddacf.
//
// Solidity: function lendInfoMap(address , uint256 ) view returns(uint256 lendAmount, uint256 refundLendAmount, bool hasNoRefund, bool hasNoClaim)
func (_PledgePool *PledgePoolCaller) LendInfoMap(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	LendAmount       *big.Int
	RefundLendAmount *big.Int
	HasNoRefund      bool
	HasNoClaim       bool
}, error) {
	var out []interface{}
	err := _PledgePool.contract.Call(opts, &out, "lendInfoMap", arg0, arg1)

	outstruct := new(struct {
		LendAmount       *big.Int
		RefundLendAmount *big.Int
		HasNoRefund      bool
		HasNoClaim       bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LendAmount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.RefundLendAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.HasNoRefund = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.HasNoClaim = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// LendInfoMap is a free data retrieval call binding the contract method 0xdcfddacf.
//
// Solidity: function lendInfoMap(address , uint256 ) view returns(uint256 lendAmount, uint256 refundLendAmount, bool hasNoRefund, bool hasNoClaim)
func (_PledgePool *PledgePoolSession) LendInfoMap(arg0 common.Address, arg1 *big.Int) (struct {
	LendAmount       *big.Int
	RefundLendAmount *big.Int
	HasNoRefund      bool
	HasNoClaim       bool
}, error) {
	return _PledgePool.Contract.LendInfoMap(&_PledgePool.CallOpts, arg0, arg1)
}

// LendInfoMap is a free data retrieval call binding the contract method 0xdcfddacf.
//
// Solidity: function lendInfoMap(address , uint256 ) view returns(uint256 lendAmount, uint256 refundLendAmount, bool hasNoRefund, bool hasNoClaim)
func (_PledgePool *PledgePoolCallerSession) LendInfoMap(arg0 common.Address, arg1 *big.Int) (struct {
	LendAmount       *big.Int
	RefundLendAmount *big.Int
	HasNoRefund      bool
	HasNoClaim       bool
}, error) {
	return _PledgePool.Contract.LendInfoMap(&_PledgePool.CallOpts, arg0, arg1)
}

// MinAmount is a free data retrieval call binding the contract method 0x9b2cb5d8.
//
// Solidity: function minAmount() view returns(uint256)
func (_PledgePool *PledgePoolCaller) MinAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PledgePool.contract.Call(opts, &out, "minAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinAmount is a free data retrieval call binding the contract method 0x9b2cb5d8.
//
// Solidity: function minAmount() view returns(uint256)
func (_PledgePool *PledgePoolSession) MinAmount() (*big.Int, error) {
	return _PledgePool.Contract.MinAmount(&_PledgePool.CallOpts)
}

// MinAmount is a free data retrieval call binding the contract method 0x9b2cb5d8.
//
// Solidity: function minAmount() view returns(uint256)
func (_PledgePool *PledgePoolCallerSession) MinAmount() (*big.Int, error) {
	return _PledgePool.Contract.MinAmount(&_PledgePool.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_PledgePool *PledgePoolCaller) Oracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PledgePool.contract.Call(opts, &out, "oracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_PledgePool *PledgePoolSession) Oracle() (common.Address, error) {
	return _PledgePool.Contract.Oracle(&_PledgePool.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_PledgePool *PledgePoolCallerSession) Oracle() (common.Address, error) {
	return _PledgePool.Contract.Oracle(&_PledgePool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PledgePool *PledgePoolCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PledgePool.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PledgePool *PledgePoolSession) Owner() (common.Address, error) {
	return _PledgePool.Contract.Owner(&_PledgePool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PledgePool *PledgePoolCallerSession) Owner() (common.Address, error) {
	return _PledgePool.Contract.Owner(&_PledgePool.CallOpts)
}

// PledgePoolInfoList is a free data retrieval call binding the contract method 0x17f1854c.
//
// Solidity: function pledgePoolInfoList(uint256 ) view returns(uint256 settleTime, uint256 endTime, uint256 interestRate, uint256 maxSupply, uint256 lendSupply, uint256 borrowSupply, uint256 mortgageRate, address lendToken, address borrowToken, uint8 state, address lendDebtToken, address borrowDebtToken, uint256 autoLiquidateThreshold)
func (_PledgePool *PledgePoolCaller) PledgePoolInfoList(opts *bind.CallOpts, arg0 *big.Int) (struct {
	SettleTime             *big.Int
	EndTime                *big.Int
	InterestRate           *big.Int
	MaxSupply              *big.Int
	LendSupply             *big.Int
	BorrowSupply           *big.Int
	MortgageRate           *big.Int
	LendToken              common.Address
	BorrowToken            common.Address
	State                  uint8
	LendDebtToken          common.Address
	BorrowDebtToken        common.Address
	AutoLiquidateThreshold *big.Int
}, error) {
	var out []interface{}
	err := _PledgePool.contract.Call(opts, &out, "pledgePoolInfoList", arg0)

	outstruct := new(struct {
		SettleTime             *big.Int
		EndTime                *big.Int
		InterestRate           *big.Int
		MaxSupply              *big.Int
		LendSupply             *big.Int
		BorrowSupply           *big.Int
		MortgageRate           *big.Int
		LendToken              common.Address
		BorrowToken            common.Address
		State                  uint8
		LendDebtToken          common.Address
		BorrowDebtToken        common.Address
		AutoLiquidateThreshold *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SettleTime = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.EndTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.InterestRate = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.MaxSupply = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.LendSupply = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.BorrowSupply = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.MortgageRate = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.LendToken = *abi.ConvertType(out[7], new(common.Address)).(*common.Address)
	outstruct.BorrowToken = *abi.ConvertType(out[8], new(common.Address)).(*common.Address)
	outstruct.State = *abi.ConvertType(out[9], new(uint8)).(*uint8)
	outstruct.LendDebtToken = *abi.ConvertType(out[10], new(common.Address)).(*common.Address)
	outstruct.BorrowDebtToken = *abi.ConvertType(out[11], new(common.Address)).(*common.Address)
	outstruct.AutoLiquidateThreshold = *abi.ConvertType(out[12], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PledgePoolInfoList is a free data retrieval call binding the contract method 0x17f1854c.
//
// Solidity: function pledgePoolInfoList(uint256 ) view returns(uint256 settleTime, uint256 endTime, uint256 interestRate, uint256 maxSupply, uint256 lendSupply, uint256 borrowSupply, uint256 mortgageRate, address lendToken, address borrowToken, uint8 state, address lendDebtToken, address borrowDebtToken, uint256 autoLiquidateThreshold)
func (_PledgePool *PledgePoolSession) PledgePoolInfoList(arg0 *big.Int) (struct {
	SettleTime             *big.Int
	EndTime                *big.Int
	InterestRate           *big.Int
	MaxSupply              *big.Int
	LendSupply             *big.Int
	BorrowSupply           *big.Int
	MortgageRate           *big.Int
	LendToken              common.Address
	BorrowToken            common.Address
	State                  uint8
	LendDebtToken          common.Address
	BorrowDebtToken        common.Address
	AutoLiquidateThreshold *big.Int
}, error) {
	return _PledgePool.Contract.PledgePoolInfoList(&_PledgePool.CallOpts, arg0)
}

// PledgePoolInfoList is a free data retrieval call binding the contract method 0x17f1854c.
//
// Solidity: function pledgePoolInfoList(uint256 ) view returns(uint256 settleTime, uint256 endTime, uint256 interestRate, uint256 maxSupply, uint256 lendSupply, uint256 borrowSupply, uint256 mortgageRate, address lendToken, address borrowToken, uint8 state, address lendDebtToken, address borrowDebtToken, uint256 autoLiquidateThreshold)
func (_PledgePool *PledgePoolCallerSession) PledgePoolInfoList(arg0 *big.Int) (struct {
	SettleTime             *big.Int
	EndTime                *big.Int
	InterestRate           *big.Int
	MaxSupply              *big.Int
	LendSupply             *big.Int
	BorrowSupply           *big.Int
	MortgageRate           *big.Int
	LendToken              common.Address
	BorrowToken            common.Address
	State                  uint8
	LendDebtToken          common.Address
	BorrowDebtToken        common.Address
	AutoLiquidateThreshold *big.Int
}, error) {
	return _PledgePool.Contract.PledgePoolInfoList(&_PledgePool.CallOpts, arg0)
}

// PoolDataInfoList is a free data retrieval call binding the contract method 0x7cad0a34.
//
// Solidity: function poolDataInfoList(uint256 ) view returns(uint256 settleAmountLend, uint256 settleAmountBorrow, uint256 finishAmountLend, uint256 finishAmountBorrow, uint256 liquidationAmountLend, uint256 liquidationAmountBorrow)
func (_PledgePool *PledgePoolCaller) PoolDataInfoList(opts *bind.CallOpts, arg0 *big.Int) (struct {
	SettleAmountLend        *big.Int
	SettleAmountBorrow      *big.Int
	FinishAmountLend        *big.Int
	FinishAmountBorrow      *big.Int
	LiquidationAmountLend   *big.Int
	LiquidationAmountBorrow *big.Int
}, error) {
	var out []interface{}
	err := _PledgePool.contract.Call(opts, &out, "poolDataInfoList", arg0)

	outstruct := new(struct {
		SettleAmountLend        *big.Int
		SettleAmountBorrow      *big.Int
		FinishAmountLend        *big.Int
		FinishAmountBorrow      *big.Int
		LiquidationAmountLend   *big.Int
		LiquidationAmountBorrow *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SettleAmountLend = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.SettleAmountBorrow = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.FinishAmountLend = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.FinishAmountBorrow = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.LiquidationAmountLend = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.LiquidationAmountBorrow = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PoolDataInfoList is a free data retrieval call binding the contract method 0x7cad0a34.
//
// Solidity: function poolDataInfoList(uint256 ) view returns(uint256 settleAmountLend, uint256 settleAmountBorrow, uint256 finishAmountLend, uint256 finishAmountBorrow, uint256 liquidationAmountLend, uint256 liquidationAmountBorrow)
func (_PledgePool *PledgePoolSession) PoolDataInfoList(arg0 *big.Int) (struct {
	SettleAmountLend        *big.Int
	SettleAmountBorrow      *big.Int
	FinishAmountLend        *big.Int
	FinishAmountBorrow      *big.Int
	LiquidationAmountLend   *big.Int
	LiquidationAmountBorrow *big.Int
}, error) {
	return _PledgePool.Contract.PoolDataInfoList(&_PledgePool.CallOpts, arg0)
}

// PoolDataInfoList is a free data retrieval call binding the contract method 0x7cad0a34.
//
// Solidity: function poolDataInfoList(uint256 ) view returns(uint256 settleAmountLend, uint256 settleAmountBorrow, uint256 finishAmountLend, uint256 finishAmountBorrow, uint256 liquidationAmountLend, uint256 liquidationAmountBorrow)
func (_PledgePool *PledgePoolCallerSession) PoolDataInfoList(arg0 *big.Int) (struct {
	SettleAmountLend        *big.Int
	SettleAmountBorrow      *big.Int
	FinishAmountLend        *big.Int
	FinishAmountBorrow      *big.Int
	LiquidationAmountLend   *big.Int
	LiquidationAmountBorrow *big.Int
}, error) {
	return _PledgePool.Contract.PoolDataInfoList(&_PledgePool.CallOpts, arg0)
}

// SwapRouter is a free data retrieval call binding the contract method 0xc31c9c07.
//
// Solidity: function swapRouter() view returns(address)
func (_PledgePool *PledgePoolCaller) SwapRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PledgePool.contract.Call(opts, &out, "swapRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SwapRouter is a free data retrieval call binding the contract method 0xc31c9c07.
//
// Solidity: function swapRouter() view returns(address)
func (_PledgePool *PledgePoolSession) SwapRouter() (common.Address, error) {
	return _PledgePool.Contract.SwapRouter(&_PledgePool.CallOpts)
}

// SwapRouter is a free data retrieval call binding the contract method 0xc31c9c07.
//
// Solidity: function swapRouter() view returns(address)
func (_PledgePool *PledgePoolCallerSession) SwapRouter() (common.Address, error) {
	return _PledgePool.Contract.SwapRouter(&_PledgePool.CallOpts)
}

// Borrow is a paid mutator transaction binding the contract method 0x0ecbcdab.
//
// Solidity: function borrow(uint256 _pid, uint256 _borrowTokenAmount) payable returns()
func (_PledgePool *PledgePoolTransactor) Borrow(opts *bind.TransactOpts, _pid *big.Int, _borrowTokenAmount *big.Int) (*types.Transaction, error) {
	return _PledgePool.contract.Transact(opts, "borrow", _pid, _borrowTokenAmount)
}

// Borrow is a paid mutator transaction binding the contract method 0x0ecbcdab.
//
// Solidity: function borrow(uint256 _pid, uint256 _borrowTokenAmount) payable returns()
func (_PledgePool *PledgePoolSession) Borrow(_pid *big.Int, _borrowTokenAmount *big.Int) (*types.Transaction, error) {
	return _PledgePool.Contract.Borrow(&_PledgePool.TransactOpts, _pid, _borrowTokenAmount)
}

// Borrow is a paid mutator transaction binding the contract method 0x0ecbcdab.
//
// Solidity: function borrow(uint256 _pid, uint256 _borrowTokenAmount) payable returns()
func (_PledgePool *PledgePoolTransactorSession) Borrow(_pid *big.Int, _borrowTokenAmount *big.Int) (*types.Transaction, error) {
	return _PledgePool.Contract.Borrow(&_PledgePool.TransactOpts, _pid, _borrowTokenAmount)
}

// ClaimBorrow is a paid mutator transaction binding the contract method 0x3ab4a445.
//
// Solidity: function claimBorrow(uint256 _pid) returns()
func (_PledgePool *PledgePoolTransactor) ClaimBorrow(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _PledgePool.contract.Transact(opts, "claimBorrow", _pid)
}

// ClaimBorrow is a paid mutator transaction binding the contract method 0x3ab4a445.
//
// Solidity: function claimBorrow(uint256 _pid) returns()
func (_PledgePool *PledgePoolSession) ClaimBorrow(_pid *big.Int) (*types.Transaction, error) {
	return _PledgePool.Contract.ClaimBorrow(&_PledgePool.TransactOpts, _pid)
}

// ClaimBorrow is a paid mutator transaction binding the contract method 0x3ab4a445.
//
// Solidity: function claimBorrow(uint256 _pid) returns()
func (_PledgePool *PledgePoolTransactorSession) ClaimBorrow(_pid *big.Int) (*types.Transaction, error) {
	return _PledgePool.Contract.ClaimBorrow(&_PledgePool.TransactOpts, _pid)
}

// ClaimLendDebtToken is a paid mutator transaction binding the contract method 0x4def20da.
//
// Solidity: function claimLendDebtToken(uint256 _pid) returns()
func (_PledgePool *PledgePoolTransactor) ClaimLendDebtToken(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _PledgePool.contract.Transact(opts, "claimLendDebtToken", _pid)
}

// ClaimLendDebtToken is a paid mutator transaction binding the contract method 0x4def20da.
//
// Solidity: function claimLendDebtToken(uint256 _pid) returns()
func (_PledgePool *PledgePoolSession) ClaimLendDebtToken(_pid *big.Int) (*types.Transaction, error) {
	return _PledgePool.Contract.ClaimLendDebtToken(&_PledgePool.TransactOpts, _pid)
}

// ClaimLendDebtToken is a paid mutator transaction binding the contract method 0x4def20da.
//
// Solidity: function claimLendDebtToken(uint256 _pid) returns()
func (_PledgePool *PledgePoolTransactorSession) ClaimLendDebtToken(_pid *big.Int) (*types.Transaction, error) {
	return _PledgePool.Contract.ClaimLendDebtToken(&_PledgePool.TransactOpts, _pid)
}

// CreatePledgePool is a paid mutator transaction binding the contract method 0x4945c25b.
//
// Solidity: function createPledgePool((uint256,uint256,uint256,uint256,uint256,address,address,address,address,uint256) params) returns(uint256)
func (_PledgePool *PledgePoolTransactor) CreatePledgePool(opts *bind.TransactOpts, params PledgePoolCreatePoolParams) (*types.Transaction, error) {
	return _PledgePool.contract.Transact(opts, "createPledgePool", params)
}

// CreatePledgePool is a paid mutator transaction binding the contract method 0x4945c25b.
//
// Solidity: function createPledgePool((uint256,uint256,uint256,uint256,uint256,address,address,address,address,uint256) params) returns(uint256)
func (_PledgePool *PledgePoolSession) CreatePledgePool(params PledgePoolCreatePoolParams) (*types.Transaction, error) {
	return _PledgePool.Contract.CreatePledgePool(&_PledgePool.TransactOpts, params)
}

// CreatePledgePool is a paid mutator transaction binding the contract method 0x4945c25b.
//
// Solidity: function createPledgePool((uint256,uint256,uint256,uint256,uint256,address,address,address,address,uint256) params) returns(uint256)
func (_PledgePool *PledgePoolTransactorSession) CreatePledgePool(params PledgePoolCreatePoolParams) (*types.Transaction, error) {
	return _PledgePool.Contract.CreatePledgePool(&_PledgePool.TransactOpts, params)
}

// DestroyBorrowDebtToken is a paid mutator transaction binding the contract method 0x0beccc40.
//
// Solidity: function destroyBorrowDebtToken(uint256 _pid, uint256 _borrowDebtTokenAmount) returns()
func (_PledgePool *PledgePoolTransactor) DestroyBorrowDebtToken(opts *bind.TransactOpts, _pid *big.Int, _borrowDebtTokenAmount *big.Int) (*types.Transaction, error) {
	return _PledgePool.contract.Transact(opts, "destroyBorrowDebtToken", _pid, _borrowDebtTokenAmount)
}

// DestroyBorrowDebtToken is a paid mutator transaction binding the contract method 0x0beccc40.
//
// Solidity: function destroyBorrowDebtToken(uint256 _pid, uint256 _borrowDebtTokenAmount) returns()
func (_PledgePool *PledgePoolSession) DestroyBorrowDebtToken(_pid *big.Int, _borrowDebtTokenAmount *big.Int) (*types.Transaction, error) {
	return _PledgePool.Contract.DestroyBorrowDebtToken(&_PledgePool.TransactOpts, _pid, _borrowDebtTokenAmount)
}

// DestroyBorrowDebtToken is a paid mutator transaction binding the contract method 0x0beccc40.
//
// Solidity: function destroyBorrowDebtToken(uint256 _pid, uint256 _borrowDebtTokenAmount) returns()
func (_PledgePool *PledgePoolTransactorSession) DestroyBorrowDebtToken(_pid *big.Int, _borrowDebtTokenAmount *big.Int) (*types.Transaction, error) {
	return _PledgePool.Contract.DestroyBorrowDebtToken(&_PledgePool.TransactOpts, _pid, _borrowDebtTokenAmount)
}

// DestroyLendDebtToken is a paid mutator transaction binding the contract method 0xdb45037d.
//
// Solidity: function destroyLendDebtToken(uint256 _pid, uint256 _lendDebtTokenAmount) returns()
func (_PledgePool *PledgePoolTransactor) DestroyLendDebtToken(opts *bind.TransactOpts, _pid *big.Int, _lendDebtTokenAmount *big.Int) (*types.Transaction, error) {
	return _PledgePool.contract.Transact(opts, "destroyLendDebtToken", _pid, _lendDebtTokenAmount)
}

// DestroyLendDebtToken is a paid mutator transaction binding the contract method 0xdb45037d.
//
// Solidity: function destroyLendDebtToken(uint256 _pid, uint256 _lendDebtTokenAmount) returns()
func (_PledgePool *PledgePoolSession) DestroyLendDebtToken(_pid *big.Int, _lendDebtTokenAmount *big.Int) (*types.Transaction, error) {
	return _PledgePool.Contract.DestroyLendDebtToken(&_PledgePool.TransactOpts, _pid, _lendDebtTokenAmount)
}

// DestroyLendDebtToken is a paid mutator transaction binding the contract method 0xdb45037d.
//
// Solidity: function destroyLendDebtToken(uint256 _pid, uint256 _lendDebtTokenAmount) returns()
func (_PledgePool *PledgePoolTransactorSession) DestroyLendDebtToken(_pid *big.Int, _lendDebtTokenAmount *big.Int) (*types.Transaction, error) {
	return _PledgePool.Contract.DestroyLendDebtToken(&_PledgePool.TransactOpts, _pid, _lendDebtTokenAmount)
}

// EmergencyWithdrawBorrow is a paid mutator transaction binding the contract method 0x7e32e47b.
//
// Solidity: function emergencyWithdrawBorrow(uint256 _pid) returns()
func (_PledgePool *PledgePoolTransactor) EmergencyWithdrawBorrow(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _PledgePool.contract.Transact(opts, "emergencyWithdrawBorrow", _pid)
}

// EmergencyWithdrawBorrow is a paid mutator transaction binding the contract method 0x7e32e47b.
//
// Solidity: function emergencyWithdrawBorrow(uint256 _pid) returns()
func (_PledgePool *PledgePoolSession) EmergencyWithdrawBorrow(_pid *big.Int) (*types.Transaction, error) {
	return _PledgePool.Contract.EmergencyWithdrawBorrow(&_PledgePool.TransactOpts, _pid)
}

// EmergencyWithdrawBorrow is a paid mutator transaction binding the contract method 0x7e32e47b.
//
// Solidity: function emergencyWithdrawBorrow(uint256 _pid) returns()
func (_PledgePool *PledgePoolTransactorSession) EmergencyWithdrawBorrow(_pid *big.Int) (*types.Transaction, error) {
	return _PledgePool.Contract.EmergencyWithdrawBorrow(&_PledgePool.TransactOpts, _pid)
}

// EmergencyWithdrawLend is a paid mutator transaction binding the contract method 0xdd182bb5.
//
// Solidity: function emergencyWithdrawLend(uint256 _pid) returns()
func (_PledgePool *PledgePoolTransactor) EmergencyWithdrawLend(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _PledgePool.contract.Transact(opts, "emergencyWithdrawLend", _pid)
}

// EmergencyWithdrawLend is a paid mutator transaction binding the contract method 0xdd182bb5.
//
// Solidity: function emergencyWithdrawLend(uint256 _pid) returns()
func (_PledgePool *PledgePoolSession) EmergencyWithdrawLend(_pid *big.Int) (*types.Transaction, error) {
	return _PledgePool.Contract.EmergencyWithdrawLend(&_PledgePool.TransactOpts, _pid)
}

// EmergencyWithdrawLend is a paid mutator transaction binding the contract method 0xdd182bb5.
//
// Solidity: function emergencyWithdrawLend(uint256 _pid) returns()
func (_PledgePool *PledgePoolTransactorSession) EmergencyWithdrawLend(_pid *big.Int) (*types.Transaction, error) {
	return _PledgePool.Contract.EmergencyWithdrawLend(&_PledgePool.TransactOpts, _pid)
}

// FinishPool is a paid mutator transaction binding the contract method 0xe0f86bcc.
//
// Solidity: function finishPool(uint256 _pid) returns()
func (_PledgePool *PledgePoolTransactor) FinishPool(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _PledgePool.contract.Transact(opts, "finishPool", _pid)
}

// FinishPool is a paid mutator transaction binding the contract method 0xe0f86bcc.
//
// Solidity: function finishPool(uint256 _pid) returns()
func (_PledgePool *PledgePoolSession) FinishPool(_pid *big.Int) (*types.Transaction, error) {
	return _PledgePool.Contract.FinishPool(&_PledgePool.TransactOpts, _pid)
}

// FinishPool is a paid mutator transaction binding the contract method 0xe0f86bcc.
//
// Solidity: function finishPool(uint256 _pid) returns()
func (_PledgePool *PledgePoolTransactorSession) FinishPool(_pid *big.Int) (*types.Transaction, error) {
	return _PledgePool.Contract.FinishPool(&_PledgePool.TransactOpts, _pid)
}

// Lend is a paid mutator transaction binding the contract method 0x9110df85.
//
// Solidity: function lend(uint256 _pid, uint256 _lendAmount) payable returns()
func (_PledgePool *PledgePoolTransactor) Lend(opts *bind.TransactOpts, _pid *big.Int, _lendAmount *big.Int) (*types.Transaction, error) {
	return _PledgePool.contract.Transact(opts, "lend", _pid, _lendAmount)
}

// Lend is a paid mutator transaction binding the contract method 0x9110df85.
//
// Solidity: function lend(uint256 _pid, uint256 _lendAmount) payable returns()
func (_PledgePool *PledgePoolSession) Lend(_pid *big.Int, _lendAmount *big.Int) (*types.Transaction, error) {
	return _PledgePool.Contract.Lend(&_PledgePool.TransactOpts, _pid, _lendAmount)
}

// Lend is a paid mutator transaction binding the contract method 0x9110df85.
//
// Solidity: function lend(uint256 _pid, uint256 _lendAmount) payable returns()
func (_PledgePool *PledgePoolTransactorSession) Lend(_pid *big.Int, _lendAmount *big.Int) (*types.Transaction, error) {
	return _PledgePool.Contract.Lend(&_PledgePool.TransactOpts, _pid, _lendAmount)
}

// LiquidatePool is a paid mutator transaction binding the contract method 0x10b72d68.
//
// Solidity: function liquidatePool(uint256 _pid) returns()
func (_PledgePool *PledgePoolTransactor) LiquidatePool(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _PledgePool.contract.Transact(opts, "liquidatePool", _pid)
}

// LiquidatePool is a paid mutator transaction binding the contract method 0x10b72d68.
//
// Solidity: function liquidatePool(uint256 _pid) returns()
func (_PledgePool *PledgePoolSession) LiquidatePool(_pid *big.Int) (*types.Transaction, error) {
	return _PledgePool.Contract.LiquidatePool(&_PledgePool.TransactOpts, _pid)
}

// LiquidatePool is a paid mutator transaction binding the contract method 0x10b72d68.
//
// Solidity: function liquidatePool(uint256 _pid) returns()
func (_PledgePool *PledgePoolTransactorSession) LiquidatePool(_pid *big.Int) (*types.Transaction, error) {
	return _PledgePool.Contract.LiquidatePool(&_PledgePool.TransactOpts, _pid)
}

// RefundBorrow is a paid mutator transaction binding the contract method 0xa62ff164.
//
// Solidity: function refundBorrow(uint256 _pid) returns()
func (_PledgePool *PledgePoolTransactor) RefundBorrow(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _PledgePool.contract.Transact(opts, "refundBorrow", _pid)
}

// RefundBorrow is a paid mutator transaction binding the contract method 0xa62ff164.
//
// Solidity: function refundBorrow(uint256 _pid) returns()
func (_PledgePool *PledgePoolSession) RefundBorrow(_pid *big.Int) (*types.Transaction, error) {
	return _PledgePool.Contract.RefundBorrow(&_PledgePool.TransactOpts, _pid)
}

// RefundBorrow is a paid mutator transaction binding the contract method 0xa62ff164.
//
// Solidity: function refundBorrow(uint256 _pid) returns()
func (_PledgePool *PledgePoolTransactorSession) RefundBorrow(_pid *big.Int) (*types.Transaction, error) {
	return _PledgePool.Contract.RefundBorrow(&_PledgePool.TransactOpts, _pid)
}

// RefundLend is a paid mutator transaction binding the contract method 0xeec8d506.
//
// Solidity: function refundLend(uint256 _pid) returns()
func (_PledgePool *PledgePoolTransactor) RefundLend(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _PledgePool.contract.Transact(opts, "refundLend", _pid)
}

// RefundLend is a paid mutator transaction binding the contract method 0xeec8d506.
//
// Solidity: function refundLend(uint256 _pid) returns()
func (_PledgePool *PledgePoolSession) RefundLend(_pid *big.Int) (*types.Transaction, error) {
	return _PledgePool.Contract.RefundLend(&_PledgePool.TransactOpts, _pid)
}

// RefundLend is a paid mutator transaction binding the contract method 0xeec8d506.
//
// Solidity: function refundLend(uint256 _pid) returns()
func (_PledgePool *PledgePoolTransactorSession) RefundLend(_pid *big.Int) (*types.Transaction, error) {
	return _PledgePool.Contract.RefundLend(&_PledgePool.TransactOpts, _pid)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PledgePool *PledgePoolTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PledgePool.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PledgePool *PledgePoolSession) RenounceOwnership() (*types.Transaction, error) {
	return _PledgePool.Contract.RenounceOwnership(&_PledgePool.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PledgePool *PledgePoolTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PledgePool.Contract.RenounceOwnership(&_PledgePool.TransactOpts)
}

// SetFee is a paid mutator transaction binding the contract method 0x52f7c988.
//
// Solidity: function setFee(uint256 _lendFee, uint256 _borrowFee) returns()
func (_PledgePool *PledgePoolTransactor) SetFee(opts *bind.TransactOpts, _lendFee *big.Int, _borrowFee *big.Int) (*types.Transaction, error) {
	return _PledgePool.contract.Transact(opts, "setFee", _lendFee, _borrowFee)
}

// SetFee is a paid mutator transaction binding the contract method 0x52f7c988.
//
// Solidity: function setFee(uint256 _lendFee, uint256 _borrowFee) returns()
func (_PledgePool *PledgePoolSession) SetFee(_lendFee *big.Int, _borrowFee *big.Int) (*types.Transaction, error) {
	return _PledgePool.Contract.SetFee(&_PledgePool.TransactOpts, _lendFee, _borrowFee)
}

// SetFee is a paid mutator transaction binding the contract method 0x52f7c988.
//
// Solidity: function setFee(uint256 _lendFee, uint256 _borrowFee) returns()
func (_PledgePool *PledgePoolTransactorSession) SetFee(_lendFee *big.Int, _borrowFee *big.Int) (*types.Transaction, error) {
	return _PledgePool.Contract.SetFee(&_PledgePool.TransactOpts, _lendFee, _borrowFee)
}

// SetFeeAddress is a paid mutator transaction binding the contract method 0x8705fcd4.
//
// Solidity: function setFeeAddress(address _feeAddress) returns()
func (_PledgePool *PledgePoolTransactor) SetFeeAddress(opts *bind.TransactOpts, _feeAddress common.Address) (*types.Transaction, error) {
	return _PledgePool.contract.Transact(opts, "setFeeAddress", _feeAddress)
}

// SetFeeAddress is a paid mutator transaction binding the contract method 0x8705fcd4.
//
// Solidity: function setFeeAddress(address _feeAddress) returns()
func (_PledgePool *PledgePoolSession) SetFeeAddress(_feeAddress common.Address) (*types.Transaction, error) {
	return _PledgePool.Contract.SetFeeAddress(&_PledgePool.TransactOpts, _feeAddress)
}

// SetFeeAddress is a paid mutator transaction binding the contract method 0x8705fcd4.
//
// Solidity: function setFeeAddress(address _feeAddress) returns()
func (_PledgePool *PledgePoolTransactorSession) SetFeeAddress(_feeAddress common.Address) (*types.Transaction, error) {
	return _PledgePool.Contract.SetFeeAddress(&_PledgePool.TransactOpts, _feeAddress)
}

// SetGlobalPaused is a paid mutator transaction binding the contract method 0xc4bd6fe6.
//
// Solidity: function setGlobalPaused() returns()
func (_PledgePool *PledgePoolTransactor) SetGlobalPaused(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PledgePool.contract.Transact(opts, "setGlobalPaused")
}

// SetGlobalPaused is a paid mutator transaction binding the contract method 0xc4bd6fe6.
//
// Solidity: function setGlobalPaused() returns()
func (_PledgePool *PledgePoolSession) SetGlobalPaused() (*types.Transaction, error) {
	return _PledgePool.Contract.SetGlobalPaused(&_PledgePool.TransactOpts)
}

// SetGlobalPaused is a paid mutator transaction binding the contract method 0xc4bd6fe6.
//
// Solidity: function setGlobalPaused() returns()
func (_PledgePool *PledgePoolTransactorSession) SetGlobalPaused() (*types.Transaction, error) {
	return _PledgePool.Contract.SetGlobalPaused(&_PledgePool.TransactOpts)
}

// SetMinAmount is a paid mutator transaction binding the contract method 0x897b0637.
//
// Solidity: function setMinAmount(uint256 _minAmount) returns()
func (_PledgePool *PledgePoolTransactor) SetMinAmount(opts *bind.TransactOpts, _minAmount *big.Int) (*types.Transaction, error) {
	return _PledgePool.contract.Transact(opts, "setMinAmount", _minAmount)
}

// SetMinAmount is a paid mutator transaction binding the contract method 0x897b0637.
//
// Solidity: function setMinAmount(uint256 _minAmount) returns()
func (_PledgePool *PledgePoolSession) SetMinAmount(_minAmount *big.Int) (*types.Transaction, error) {
	return _PledgePool.Contract.SetMinAmount(&_PledgePool.TransactOpts, _minAmount)
}

// SetMinAmount is a paid mutator transaction binding the contract method 0x897b0637.
//
// Solidity: function setMinAmount(uint256 _minAmount) returns()
func (_PledgePool *PledgePoolTransactorSession) SetMinAmount(_minAmount *big.Int) (*types.Transaction, error) {
	return _PledgePool.Contract.SetMinAmount(&_PledgePool.TransactOpts, _minAmount)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address _oracle) returns()
func (_PledgePool *PledgePoolTransactor) SetOracle(opts *bind.TransactOpts, _oracle common.Address) (*types.Transaction, error) {
	return _PledgePool.contract.Transact(opts, "setOracle", _oracle)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address _oracle) returns()
func (_PledgePool *PledgePoolSession) SetOracle(_oracle common.Address) (*types.Transaction, error) {
	return _PledgePool.Contract.SetOracle(&_PledgePool.TransactOpts, _oracle)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address _oracle) returns()
func (_PledgePool *PledgePoolTransactorSession) SetOracle(_oracle common.Address) (*types.Transaction, error) {
	return _PledgePool.Contract.SetOracle(&_PledgePool.TransactOpts, _oracle)
}

// SetSwapRouter is a paid mutator transaction binding the contract method 0x41273657.
//
// Solidity: function setSwapRouter(address _swapRouter) returns()
func (_PledgePool *PledgePoolTransactor) SetSwapRouter(opts *bind.TransactOpts, _swapRouter common.Address) (*types.Transaction, error) {
	return _PledgePool.contract.Transact(opts, "setSwapRouter", _swapRouter)
}

// SetSwapRouter is a paid mutator transaction binding the contract method 0x41273657.
//
// Solidity: function setSwapRouter(address _swapRouter) returns()
func (_PledgePool *PledgePoolSession) SetSwapRouter(_swapRouter common.Address) (*types.Transaction, error) {
	return _PledgePool.Contract.SetSwapRouter(&_PledgePool.TransactOpts, _swapRouter)
}

// SetSwapRouter is a paid mutator transaction binding the contract method 0x41273657.
//
// Solidity: function setSwapRouter(address _swapRouter) returns()
func (_PledgePool *PledgePoolTransactorSession) SetSwapRouter(_swapRouter common.Address) (*types.Transaction, error) {
	return _PledgePool.Contract.SetSwapRouter(&_PledgePool.TransactOpts, _swapRouter)
}

// SettlePool is a paid mutator transaction binding the contract method 0x4ec2d875.
//
// Solidity: function settlePool(uint256 _pid) returns()
func (_PledgePool *PledgePoolTransactor) SettlePool(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _PledgePool.contract.Transact(opts, "settlePool", _pid)
}

// SettlePool is a paid mutator transaction binding the contract method 0x4ec2d875.
//
// Solidity: function settlePool(uint256 _pid) returns()
func (_PledgePool *PledgePoolSession) SettlePool(_pid *big.Int) (*types.Transaction, error) {
	return _PledgePool.Contract.SettlePool(&_PledgePool.TransactOpts, _pid)
}

// SettlePool is a paid mutator transaction binding the contract method 0x4ec2d875.
//
// Solidity: function settlePool(uint256 _pid) returns()
func (_PledgePool *PledgePoolTransactorSession) SettlePool(_pid *big.Int) (*types.Transaction, error) {
	return _PledgePool.Contract.SettlePool(&_PledgePool.TransactOpts, _pid)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PledgePool *PledgePoolTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PledgePool.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PledgePool *PledgePoolSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PledgePool.Contract.TransferOwnership(&_PledgePool.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PledgePool *PledgePoolTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PledgePool.Contract.TransferOwnership(&_PledgePool.TransactOpts, newOwner)
}

// PledgePoolBorrowIterator is returned from FilterBorrow and is used to iterate over the raw logs and unpacked data for Borrow events raised by the PledgePool contract.
type PledgePoolBorrowIterator struct {
	Event *PledgePoolBorrow // Event containing the contract specifics and raw log

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
func (it *PledgePoolBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PledgePoolBorrow)
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
		it.Event = new(PledgePoolBorrow)
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
func (it *PledgePoolBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PledgePoolBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PledgePoolBorrow represents a Borrow event raised by the PledgePool contract.
type PledgePoolBorrow struct {
	Pid      *big.Int
	Token    common.Address
	Borrower common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBorrow is a free log retrieval operation binding the contract event 0xf9a33434428db5f0416c03e38307599ad0b9b9965d6c070eb08e87cc1f0ca50e.
//
// Solidity: event Borrow(uint256 indexed pid, address indexed token, address indexed borrower, uint256 amount)
func (_PledgePool *PledgePoolFilterer) FilterBorrow(opts *bind.FilterOpts, pid []*big.Int, token []common.Address, borrower []common.Address) (*PledgePoolBorrowIterator, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _PledgePool.contract.FilterLogs(opts, "Borrow", pidRule, tokenRule, borrowerRule)
	if err != nil {
		return nil, err
	}
	return &PledgePoolBorrowIterator{contract: _PledgePool.contract, event: "Borrow", logs: logs, sub: sub}, nil
}

// WatchBorrow is a free log subscription operation binding the contract event 0xf9a33434428db5f0416c03e38307599ad0b9b9965d6c070eb08e87cc1f0ca50e.
//
// Solidity: event Borrow(uint256 indexed pid, address indexed token, address indexed borrower, uint256 amount)
func (_PledgePool *PledgePoolFilterer) WatchBorrow(opts *bind.WatchOpts, sink chan<- *PledgePoolBorrow, pid []*big.Int, token []common.Address, borrower []common.Address) (event.Subscription, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _PledgePool.contract.WatchLogs(opts, "Borrow", pidRule, tokenRule, borrowerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PledgePoolBorrow)
				if err := _PledgePool.contract.UnpackLog(event, "Borrow", log); err != nil {
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

// ParseBorrow is a log parse operation binding the contract event 0xf9a33434428db5f0416c03e38307599ad0b9b9965d6c070eb08e87cc1f0ca50e.
//
// Solidity: event Borrow(uint256 indexed pid, address indexed token, address indexed borrower, uint256 amount)
func (_PledgePool *PledgePoolFilterer) ParseBorrow(log types.Log) (*PledgePoolBorrow, error) {
	event := new(PledgePoolBorrow)
	if err := _PledgePool.contract.UnpackLog(event, "Borrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PledgePoolClaimBorrowIterator is returned from FilterClaimBorrow and is used to iterate over the raw logs and unpacked data for ClaimBorrow events raised by the PledgePool contract.
type PledgePoolClaimBorrowIterator struct {
	Event *PledgePoolClaimBorrow // Event containing the contract specifics and raw log

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
func (it *PledgePoolClaimBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PledgePoolClaimBorrow)
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
		it.Event = new(PledgePoolClaimBorrow)
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
func (it *PledgePoolClaimBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PledgePoolClaimBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PledgePoolClaimBorrow represents a ClaimBorrow event raised by the PledgePool contract.
type PledgePoolClaimBorrow struct {
	Pid                   *big.Int
	Token                 common.Address
	Borrower              common.Address
	BorrowDebtTokenAmount *big.Int
	LendTokenAmount       *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterClaimBorrow is a free log retrieval operation binding the contract event 0x1226b632b60c533dd17d143c6f2616eb2415a03ce2ce3168a8311c8787e25556.
//
// Solidity: event ClaimBorrow(uint256 indexed pid, address indexed token, address indexed borrower, uint256 borrowDebtTokenAmount, uint256 lendTokenAmount)
func (_PledgePool *PledgePoolFilterer) FilterClaimBorrow(opts *bind.FilterOpts, pid []*big.Int, token []common.Address, borrower []common.Address) (*PledgePoolClaimBorrowIterator, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _PledgePool.contract.FilterLogs(opts, "ClaimBorrow", pidRule, tokenRule, borrowerRule)
	if err != nil {
		return nil, err
	}
	return &PledgePoolClaimBorrowIterator{contract: _PledgePool.contract, event: "ClaimBorrow", logs: logs, sub: sub}, nil
}

// WatchClaimBorrow is a free log subscription operation binding the contract event 0x1226b632b60c533dd17d143c6f2616eb2415a03ce2ce3168a8311c8787e25556.
//
// Solidity: event ClaimBorrow(uint256 indexed pid, address indexed token, address indexed borrower, uint256 borrowDebtTokenAmount, uint256 lendTokenAmount)
func (_PledgePool *PledgePoolFilterer) WatchClaimBorrow(opts *bind.WatchOpts, sink chan<- *PledgePoolClaimBorrow, pid []*big.Int, token []common.Address, borrower []common.Address) (event.Subscription, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _PledgePool.contract.WatchLogs(opts, "ClaimBorrow", pidRule, tokenRule, borrowerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PledgePoolClaimBorrow)
				if err := _PledgePool.contract.UnpackLog(event, "ClaimBorrow", log); err != nil {
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

// ParseClaimBorrow is a log parse operation binding the contract event 0x1226b632b60c533dd17d143c6f2616eb2415a03ce2ce3168a8311c8787e25556.
//
// Solidity: event ClaimBorrow(uint256 indexed pid, address indexed token, address indexed borrower, uint256 borrowDebtTokenAmount, uint256 lendTokenAmount)
func (_PledgePool *PledgePoolFilterer) ParseClaimBorrow(log types.Log) (*PledgePoolClaimBorrow, error) {
	event := new(PledgePoolClaimBorrow)
	if err := _PledgePool.contract.UnpackLog(event, "ClaimBorrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PledgePoolClaimLendDebtTokenIterator is returned from FilterClaimLendDebtToken and is used to iterate over the raw logs and unpacked data for ClaimLendDebtToken events raised by the PledgePool contract.
type PledgePoolClaimLendDebtTokenIterator struct {
	Event *PledgePoolClaimLendDebtToken // Event containing the contract specifics and raw log

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
func (it *PledgePoolClaimLendDebtTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PledgePoolClaimLendDebtToken)
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
		it.Event = new(PledgePoolClaimLendDebtToken)
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
func (it *PledgePoolClaimLendDebtTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PledgePoolClaimLendDebtTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PledgePoolClaimLendDebtToken represents a ClaimLendDebtToken event raised by the PledgePool contract.
type PledgePoolClaimLendDebtToken struct {
	Pid    *big.Int
	Token  common.Address
	Lender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterClaimLendDebtToken is a free log retrieval operation binding the contract event 0x535481ce60a76f235eeee0e6601988d1031b55eef7e85a043cd7a861f94fb59f.
//
// Solidity: event ClaimLendDebtToken(uint256 indexed pid, address indexed token, address indexed lender, uint256 amount)
func (_PledgePool *PledgePoolFilterer) FilterClaimLendDebtToken(opts *bind.FilterOpts, pid []*big.Int, token []common.Address, lender []common.Address) (*PledgePoolClaimLendDebtTokenIterator, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var lenderRule []interface{}
	for _, lenderItem := range lender {
		lenderRule = append(lenderRule, lenderItem)
	}

	logs, sub, err := _PledgePool.contract.FilterLogs(opts, "ClaimLendDebtToken", pidRule, tokenRule, lenderRule)
	if err != nil {
		return nil, err
	}
	return &PledgePoolClaimLendDebtTokenIterator{contract: _PledgePool.contract, event: "ClaimLendDebtToken", logs: logs, sub: sub}, nil
}

// WatchClaimLendDebtToken is a free log subscription operation binding the contract event 0x535481ce60a76f235eeee0e6601988d1031b55eef7e85a043cd7a861f94fb59f.
//
// Solidity: event ClaimLendDebtToken(uint256 indexed pid, address indexed token, address indexed lender, uint256 amount)
func (_PledgePool *PledgePoolFilterer) WatchClaimLendDebtToken(opts *bind.WatchOpts, sink chan<- *PledgePoolClaimLendDebtToken, pid []*big.Int, token []common.Address, lender []common.Address) (event.Subscription, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var lenderRule []interface{}
	for _, lenderItem := range lender {
		lenderRule = append(lenderRule, lenderItem)
	}

	logs, sub, err := _PledgePool.contract.WatchLogs(opts, "ClaimLendDebtToken", pidRule, tokenRule, lenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PledgePoolClaimLendDebtToken)
				if err := _PledgePool.contract.UnpackLog(event, "ClaimLendDebtToken", log); err != nil {
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

// ParseClaimLendDebtToken is a log parse operation binding the contract event 0x535481ce60a76f235eeee0e6601988d1031b55eef7e85a043cd7a861f94fb59f.
//
// Solidity: event ClaimLendDebtToken(uint256 indexed pid, address indexed token, address indexed lender, uint256 amount)
func (_PledgePool *PledgePoolFilterer) ParseClaimLendDebtToken(log types.Log) (*PledgePoolClaimLendDebtToken, error) {
	event := new(PledgePoolClaimLendDebtToken)
	if err := _PledgePool.contract.UnpackLog(event, "ClaimLendDebtToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PledgePoolCreatePledgePoolIterator is returned from FilterCreatePledgePool and is used to iterate over the raw logs and unpacked data for CreatePledgePool events raised by the PledgePool contract.
type PledgePoolCreatePledgePoolIterator struct {
	Event *PledgePoolCreatePledgePool // Event containing the contract specifics and raw log

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
func (it *PledgePoolCreatePledgePoolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PledgePoolCreatePledgePool)
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
		it.Event = new(PledgePoolCreatePledgePool)
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
func (it *PledgePoolCreatePledgePoolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PledgePoolCreatePledgePoolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PledgePoolCreatePledgePool represents a CreatePledgePool event raised by the PledgePool contract.
type PledgePoolCreatePledgePool struct {
	Pid *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterCreatePledgePool is a free log retrieval operation binding the contract event 0x67d5edbe9035b9e077288ae99fef26dd09577336cade83fbf1270ad01af587f9.
//
// Solidity: event CreatePledgePool(uint256 indexed pid)
func (_PledgePool *PledgePoolFilterer) FilterCreatePledgePool(opts *bind.FilterOpts, pid []*big.Int) (*PledgePoolCreatePledgePoolIterator, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _PledgePool.contract.FilterLogs(opts, "CreatePledgePool", pidRule)
	if err != nil {
		return nil, err
	}
	return &PledgePoolCreatePledgePoolIterator{contract: _PledgePool.contract, event: "CreatePledgePool", logs: logs, sub: sub}, nil
}

// WatchCreatePledgePool is a free log subscription operation binding the contract event 0x67d5edbe9035b9e077288ae99fef26dd09577336cade83fbf1270ad01af587f9.
//
// Solidity: event CreatePledgePool(uint256 indexed pid)
func (_PledgePool *PledgePoolFilterer) WatchCreatePledgePool(opts *bind.WatchOpts, sink chan<- *PledgePoolCreatePledgePool, pid []*big.Int) (event.Subscription, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _PledgePool.contract.WatchLogs(opts, "CreatePledgePool", pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PledgePoolCreatePledgePool)
				if err := _PledgePool.contract.UnpackLog(event, "CreatePledgePool", log); err != nil {
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

// ParseCreatePledgePool is a log parse operation binding the contract event 0x67d5edbe9035b9e077288ae99fef26dd09577336cade83fbf1270ad01af587f9.
//
// Solidity: event CreatePledgePool(uint256 indexed pid)
func (_PledgePool *PledgePoolFilterer) ParseCreatePledgePool(log types.Log) (*PledgePoolCreatePledgePool, error) {
	event := new(PledgePoolCreatePledgePool)
	if err := _PledgePool.contract.UnpackLog(event, "CreatePledgePool", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PledgePoolDestroyBorrowDebtTokenIterator is returned from FilterDestroyBorrowDebtToken and is used to iterate over the raw logs and unpacked data for DestroyBorrowDebtToken events raised by the PledgePool contract.
type PledgePoolDestroyBorrowDebtTokenIterator struct {
	Event *PledgePoolDestroyBorrowDebtToken // Event containing the contract specifics and raw log

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
func (it *PledgePoolDestroyBorrowDebtTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PledgePoolDestroyBorrowDebtToken)
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
		it.Event = new(PledgePoolDestroyBorrowDebtToken)
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
func (it *PledgePoolDestroyBorrowDebtTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PledgePoolDestroyBorrowDebtTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PledgePoolDestroyBorrowDebtToken represents a DestroyBorrowDebtToken event raised by the PledgePool contract.
type PledgePoolDestroyBorrowDebtToken struct {
	Pid          *big.Int
	Token        common.Address
	Borrower     common.Address
	BurnAmount   *big.Int
	RedeemAmount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDestroyBorrowDebtToken is a free log retrieval operation binding the contract event 0x6523113a4bcd4da7dc9c771b759cc100a53e2925903ef0789e7be97aa234160a.
//
// Solidity: event DestroyBorrowDebtToken(uint256 indexed pid, address indexed token, address indexed borrower, uint256 burnAmount, uint256 redeemAmount)
func (_PledgePool *PledgePoolFilterer) FilterDestroyBorrowDebtToken(opts *bind.FilterOpts, pid []*big.Int, token []common.Address, borrower []common.Address) (*PledgePoolDestroyBorrowDebtTokenIterator, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _PledgePool.contract.FilterLogs(opts, "DestroyBorrowDebtToken", pidRule, tokenRule, borrowerRule)
	if err != nil {
		return nil, err
	}
	return &PledgePoolDestroyBorrowDebtTokenIterator{contract: _PledgePool.contract, event: "DestroyBorrowDebtToken", logs: logs, sub: sub}, nil
}

// WatchDestroyBorrowDebtToken is a free log subscription operation binding the contract event 0x6523113a4bcd4da7dc9c771b759cc100a53e2925903ef0789e7be97aa234160a.
//
// Solidity: event DestroyBorrowDebtToken(uint256 indexed pid, address indexed token, address indexed borrower, uint256 burnAmount, uint256 redeemAmount)
func (_PledgePool *PledgePoolFilterer) WatchDestroyBorrowDebtToken(opts *bind.WatchOpts, sink chan<- *PledgePoolDestroyBorrowDebtToken, pid []*big.Int, token []common.Address, borrower []common.Address) (event.Subscription, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _PledgePool.contract.WatchLogs(opts, "DestroyBorrowDebtToken", pidRule, tokenRule, borrowerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PledgePoolDestroyBorrowDebtToken)
				if err := _PledgePool.contract.UnpackLog(event, "DestroyBorrowDebtToken", log); err != nil {
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

// ParseDestroyBorrowDebtToken is a log parse operation binding the contract event 0x6523113a4bcd4da7dc9c771b759cc100a53e2925903ef0789e7be97aa234160a.
//
// Solidity: event DestroyBorrowDebtToken(uint256 indexed pid, address indexed token, address indexed borrower, uint256 burnAmount, uint256 redeemAmount)
func (_PledgePool *PledgePoolFilterer) ParseDestroyBorrowDebtToken(log types.Log) (*PledgePoolDestroyBorrowDebtToken, error) {
	event := new(PledgePoolDestroyBorrowDebtToken)
	if err := _PledgePool.contract.UnpackLog(event, "DestroyBorrowDebtToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PledgePoolDestroyLendDebtTokenIterator is returned from FilterDestroyLendDebtToken and is used to iterate over the raw logs and unpacked data for DestroyLendDebtToken events raised by the PledgePool contract.
type PledgePoolDestroyLendDebtTokenIterator struct {
	Event *PledgePoolDestroyLendDebtToken // Event containing the contract specifics and raw log

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
func (it *PledgePoolDestroyLendDebtTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PledgePoolDestroyLendDebtToken)
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
		it.Event = new(PledgePoolDestroyLendDebtToken)
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
func (it *PledgePoolDestroyLendDebtTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PledgePoolDestroyLendDebtTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PledgePoolDestroyLendDebtToken represents a DestroyLendDebtToken event raised by the PledgePool contract.
type PledgePoolDestroyLendDebtToken struct {
	Pid          *big.Int
	Token        common.Address
	Lender       common.Address
	BurnAmount   *big.Int
	RedeemAmount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDestroyLendDebtToken is a free log retrieval operation binding the contract event 0xf2ddf7a85d68bf37293e4fd7e4f19e0309db0714a4ad5251906341cc9ebb90ba.
//
// Solidity: event DestroyLendDebtToken(uint256 indexed pid, address indexed token, address indexed lender, uint256 burnAmount, uint256 redeemAmount)
func (_PledgePool *PledgePoolFilterer) FilterDestroyLendDebtToken(opts *bind.FilterOpts, pid []*big.Int, token []common.Address, lender []common.Address) (*PledgePoolDestroyLendDebtTokenIterator, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var lenderRule []interface{}
	for _, lenderItem := range lender {
		lenderRule = append(lenderRule, lenderItem)
	}

	logs, sub, err := _PledgePool.contract.FilterLogs(opts, "DestroyLendDebtToken", pidRule, tokenRule, lenderRule)
	if err != nil {
		return nil, err
	}
	return &PledgePoolDestroyLendDebtTokenIterator{contract: _PledgePool.contract, event: "DestroyLendDebtToken", logs: logs, sub: sub}, nil
}

// WatchDestroyLendDebtToken is a free log subscription operation binding the contract event 0xf2ddf7a85d68bf37293e4fd7e4f19e0309db0714a4ad5251906341cc9ebb90ba.
//
// Solidity: event DestroyLendDebtToken(uint256 indexed pid, address indexed token, address indexed lender, uint256 burnAmount, uint256 redeemAmount)
func (_PledgePool *PledgePoolFilterer) WatchDestroyLendDebtToken(opts *bind.WatchOpts, sink chan<- *PledgePoolDestroyLendDebtToken, pid []*big.Int, token []common.Address, lender []common.Address) (event.Subscription, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var lenderRule []interface{}
	for _, lenderItem := range lender {
		lenderRule = append(lenderRule, lenderItem)
	}

	logs, sub, err := _PledgePool.contract.WatchLogs(opts, "DestroyLendDebtToken", pidRule, tokenRule, lenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PledgePoolDestroyLendDebtToken)
				if err := _PledgePool.contract.UnpackLog(event, "DestroyLendDebtToken", log); err != nil {
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

// ParseDestroyLendDebtToken is a log parse operation binding the contract event 0xf2ddf7a85d68bf37293e4fd7e4f19e0309db0714a4ad5251906341cc9ebb90ba.
//
// Solidity: event DestroyLendDebtToken(uint256 indexed pid, address indexed token, address indexed lender, uint256 burnAmount, uint256 redeemAmount)
func (_PledgePool *PledgePoolFilterer) ParseDestroyLendDebtToken(log types.Log) (*PledgePoolDestroyLendDebtToken, error) {
	event := new(PledgePoolDestroyLendDebtToken)
	if err := _PledgePool.contract.UnpackLog(event, "DestroyLendDebtToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PledgePoolEmergencyWithdrawBorrowIterator is returned from FilterEmergencyWithdrawBorrow and is used to iterate over the raw logs and unpacked data for EmergencyWithdrawBorrow events raised by the PledgePool contract.
type PledgePoolEmergencyWithdrawBorrowIterator struct {
	Event *PledgePoolEmergencyWithdrawBorrow // Event containing the contract specifics and raw log

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
func (it *PledgePoolEmergencyWithdrawBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PledgePoolEmergencyWithdrawBorrow)
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
		it.Event = new(PledgePoolEmergencyWithdrawBorrow)
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
func (it *PledgePoolEmergencyWithdrawBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PledgePoolEmergencyWithdrawBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PledgePoolEmergencyWithdrawBorrow represents a EmergencyWithdrawBorrow event raised by the PledgePool contract.
type PledgePoolEmergencyWithdrawBorrow struct {
	Pid      *big.Int
	Token    common.Address
	Borrower common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterEmergencyWithdrawBorrow is a free log retrieval operation binding the contract event 0xded74dfac71815a7a7980ee897eb0a579034909191f76db76a03e2f17b92c27b.
//
// Solidity: event EmergencyWithdrawBorrow(uint256 indexed pid, address indexed token, address indexed borrower, uint256 amount)
func (_PledgePool *PledgePoolFilterer) FilterEmergencyWithdrawBorrow(opts *bind.FilterOpts, pid []*big.Int, token []common.Address, borrower []common.Address) (*PledgePoolEmergencyWithdrawBorrowIterator, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _PledgePool.contract.FilterLogs(opts, "EmergencyWithdrawBorrow", pidRule, tokenRule, borrowerRule)
	if err != nil {
		return nil, err
	}
	return &PledgePoolEmergencyWithdrawBorrowIterator{contract: _PledgePool.contract, event: "EmergencyWithdrawBorrow", logs: logs, sub: sub}, nil
}

// WatchEmergencyWithdrawBorrow is a free log subscription operation binding the contract event 0xded74dfac71815a7a7980ee897eb0a579034909191f76db76a03e2f17b92c27b.
//
// Solidity: event EmergencyWithdrawBorrow(uint256 indexed pid, address indexed token, address indexed borrower, uint256 amount)
func (_PledgePool *PledgePoolFilterer) WatchEmergencyWithdrawBorrow(opts *bind.WatchOpts, sink chan<- *PledgePoolEmergencyWithdrawBorrow, pid []*big.Int, token []common.Address, borrower []common.Address) (event.Subscription, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _PledgePool.contract.WatchLogs(opts, "EmergencyWithdrawBorrow", pidRule, tokenRule, borrowerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PledgePoolEmergencyWithdrawBorrow)
				if err := _PledgePool.contract.UnpackLog(event, "EmergencyWithdrawBorrow", log); err != nil {
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

// ParseEmergencyWithdrawBorrow is a log parse operation binding the contract event 0xded74dfac71815a7a7980ee897eb0a579034909191f76db76a03e2f17b92c27b.
//
// Solidity: event EmergencyWithdrawBorrow(uint256 indexed pid, address indexed token, address indexed borrower, uint256 amount)
func (_PledgePool *PledgePoolFilterer) ParseEmergencyWithdrawBorrow(log types.Log) (*PledgePoolEmergencyWithdrawBorrow, error) {
	event := new(PledgePoolEmergencyWithdrawBorrow)
	if err := _PledgePool.contract.UnpackLog(event, "EmergencyWithdrawBorrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PledgePoolEmergencyWithdrawLendIterator is returned from FilterEmergencyWithdrawLend and is used to iterate over the raw logs and unpacked data for EmergencyWithdrawLend events raised by the PledgePool contract.
type PledgePoolEmergencyWithdrawLendIterator struct {
	Event *PledgePoolEmergencyWithdrawLend // Event containing the contract specifics and raw log

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
func (it *PledgePoolEmergencyWithdrawLendIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PledgePoolEmergencyWithdrawLend)
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
		it.Event = new(PledgePoolEmergencyWithdrawLend)
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
func (it *PledgePoolEmergencyWithdrawLendIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PledgePoolEmergencyWithdrawLendIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PledgePoolEmergencyWithdrawLend represents a EmergencyWithdrawLend event raised by the PledgePool contract.
type PledgePoolEmergencyWithdrawLend struct {
	Pid    *big.Int
	Token  common.Address
	Lender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEmergencyWithdrawLend is a free log retrieval operation binding the contract event 0x1925c2d32a3486e44382eebcaea2888a167e91fde83a533b9ca611b0fc7931d2.
//
// Solidity: event EmergencyWithdrawLend(uint256 indexed pid, address indexed token, address indexed lender, uint256 amount)
func (_PledgePool *PledgePoolFilterer) FilterEmergencyWithdrawLend(opts *bind.FilterOpts, pid []*big.Int, token []common.Address, lender []common.Address) (*PledgePoolEmergencyWithdrawLendIterator, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var lenderRule []interface{}
	for _, lenderItem := range lender {
		lenderRule = append(lenderRule, lenderItem)
	}

	logs, sub, err := _PledgePool.contract.FilterLogs(opts, "EmergencyWithdrawLend", pidRule, tokenRule, lenderRule)
	if err != nil {
		return nil, err
	}
	return &PledgePoolEmergencyWithdrawLendIterator{contract: _PledgePool.contract, event: "EmergencyWithdrawLend", logs: logs, sub: sub}, nil
}

// WatchEmergencyWithdrawLend is a free log subscription operation binding the contract event 0x1925c2d32a3486e44382eebcaea2888a167e91fde83a533b9ca611b0fc7931d2.
//
// Solidity: event EmergencyWithdrawLend(uint256 indexed pid, address indexed token, address indexed lender, uint256 amount)
func (_PledgePool *PledgePoolFilterer) WatchEmergencyWithdrawLend(opts *bind.WatchOpts, sink chan<- *PledgePoolEmergencyWithdrawLend, pid []*big.Int, token []common.Address, lender []common.Address) (event.Subscription, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var lenderRule []interface{}
	for _, lenderItem := range lender {
		lenderRule = append(lenderRule, lenderItem)
	}

	logs, sub, err := _PledgePool.contract.WatchLogs(opts, "EmergencyWithdrawLend", pidRule, tokenRule, lenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PledgePoolEmergencyWithdrawLend)
				if err := _PledgePool.contract.UnpackLog(event, "EmergencyWithdrawLend", log); err != nil {
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

// ParseEmergencyWithdrawLend is a log parse operation binding the contract event 0x1925c2d32a3486e44382eebcaea2888a167e91fde83a533b9ca611b0fc7931d2.
//
// Solidity: event EmergencyWithdrawLend(uint256 indexed pid, address indexed token, address indexed lender, uint256 amount)
func (_PledgePool *PledgePoolFilterer) ParseEmergencyWithdrawLend(log types.Log) (*PledgePoolEmergencyWithdrawLend, error) {
	event := new(PledgePoolEmergencyWithdrawLend)
	if err := _PledgePool.contract.UnpackLog(event, "EmergencyWithdrawLend", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PledgePoolFinishPoolIterator is returned from FilterFinishPool and is used to iterate over the raw logs and unpacked data for FinishPool events raised by the PledgePool contract.
type PledgePoolFinishPoolIterator struct {
	Event *PledgePoolFinishPool // Event containing the contract specifics and raw log

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
func (it *PledgePoolFinishPoolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PledgePoolFinishPool)
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
		it.Event = new(PledgePoolFinishPool)
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
func (it *PledgePoolFinishPoolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PledgePoolFinishPoolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PledgePoolFinishPool represents a FinishPool event raised by the PledgePool contract.
type PledgePoolFinishPool struct {
	Pid                *big.Int
	FinishAmountLend   *big.Int
	FinishAmountBorrow *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterFinishPool is a free log retrieval operation binding the contract event 0xe197bf82539ba04efe8d438ac6dc6e530ed505bb8442c1c746badc834f159777.
//
// Solidity: event FinishPool(uint256 indexed pid, uint256 finishAmountLend, uint256 finishAmountBorrow)
func (_PledgePool *PledgePoolFilterer) FilterFinishPool(opts *bind.FilterOpts, pid []*big.Int) (*PledgePoolFinishPoolIterator, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _PledgePool.contract.FilterLogs(opts, "FinishPool", pidRule)
	if err != nil {
		return nil, err
	}
	return &PledgePoolFinishPoolIterator{contract: _PledgePool.contract, event: "FinishPool", logs: logs, sub: sub}, nil
}

// WatchFinishPool is a free log subscription operation binding the contract event 0xe197bf82539ba04efe8d438ac6dc6e530ed505bb8442c1c746badc834f159777.
//
// Solidity: event FinishPool(uint256 indexed pid, uint256 finishAmountLend, uint256 finishAmountBorrow)
func (_PledgePool *PledgePoolFilterer) WatchFinishPool(opts *bind.WatchOpts, sink chan<- *PledgePoolFinishPool, pid []*big.Int) (event.Subscription, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _PledgePool.contract.WatchLogs(opts, "FinishPool", pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PledgePoolFinishPool)
				if err := _PledgePool.contract.UnpackLog(event, "FinishPool", log); err != nil {
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

// ParseFinishPool is a log parse operation binding the contract event 0xe197bf82539ba04efe8d438ac6dc6e530ed505bb8442c1c746badc834f159777.
//
// Solidity: event FinishPool(uint256 indexed pid, uint256 finishAmountLend, uint256 finishAmountBorrow)
func (_PledgePool *PledgePoolFilterer) ParseFinishPool(log types.Log) (*PledgePoolFinishPool, error) {
	event := new(PledgePoolFinishPool)
	if err := _PledgePool.contract.UnpackLog(event, "FinishPool", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PledgePoolLendIterator is returned from FilterLend and is used to iterate over the raw logs and unpacked data for Lend events raised by the PledgePool contract.
type PledgePoolLendIterator struct {
	Event *PledgePoolLend // Event containing the contract specifics and raw log

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
func (it *PledgePoolLendIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PledgePoolLend)
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
		it.Event = new(PledgePoolLend)
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
func (it *PledgePoolLendIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PledgePoolLendIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PledgePoolLend represents a Lend event raised by the PledgePool contract.
type PledgePoolLend struct {
	Pid    *big.Int
	Token  common.Address
	Lender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLend is a free log retrieval operation binding the contract event 0x77c494147cc26c9ccb1f3f1926fb5bffa128633d935fdf7147dff0da6f740db4.
//
// Solidity: event Lend(uint256 indexed pid, address indexed token, address indexed lender, uint256 amount)
func (_PledgePool *PledgePoolFilterer) FilterLend(opts *bind.FilterOpts, pid []*big.Int, token []common.Address, lender []common.Address) (*PledgePoolLendIterator, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var lenderRule []interface{}
	for _, lenderItem := range lender {
		lenderRule = append(lenderRule, lenderItem)
	}

	logs, sub, err := _PledgePool.contract.FilterLogs(opts, "Lend", pidRule, tokenRule, lenderRule)
	if err != nil {
		return nil, err
	}
	return &PledgePoolLendIterator{contract: _PledgePool.contract, event: "Lend", logs: logs, sub: sub}, nil
}

// WatchLend is a free log subscription operation binding the contract event 0x77c494147cc26c9ccb1f3f1926fb5bffa128633d935fdf7147dff0da6f740db4.
//
// Solidity: event Lend(uint256 indexed pid, address indexed token, address indexed lender, uint256 amount)
func (_PledgePool *PledgePoolFilterer) WatchLend(opts *bind.WatchOpts, sink chan<- *PledgePoolLend, pid []*big.Int, token []common.Address, lender []common.Address) (event.Subscription, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var lenderRule []interface{}
	for _, lenderItem := range lender {
		lenderRule = append(lenderRule, lenderItem)
	}

	logs, sub, err := _PledgePool.contract.WatchLogs(opts, "Lend", pidRule, tokenRule, lenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PledgePoolLend)
				if err := _PledgePool.contract.UnpackLog(event, "Lend", log); err != nil {
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

// ParseLend is a log parse operation binding the contract event 0x77c494147cc26c9ccb1f3f1926fb5bffa128633d935fdf7147dff0da6f740db4.
//
// Solidity: event Lend(uint256 indexed pid, address indexed token, address indexed lender, uint256 amount)
func (_PledgePool *PledgePoolFilterer) ParseLend(log types.Log) (*PledgePoolLend, error) {
	event := new(PledgePoolLend)
	if err := _PledgePool.contract.UnpackLog(event, "Lend", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PledgePoolLiquidatePoolIterator is returned from FilterLiquidatePool and is used to iterate over the raw logs and unpacked data for LiquidatePool events raised by the PledgePool contract.
type PledgePoolLiquidatePoolIterator struct {
	Event *PledgePoolLiquidatePool // Event containing the contract specifics and raw log

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
func (it *PledgePoolLiquidatePoolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PledgePoolLiquidatePool)
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
		it.Event = new(PledgePoolLiquidatePool)
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
func (it *PledgePoolLiquidatePoolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PledgePoolLiquidatePoolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PledgePoolLiquidatePool represents a LiquidatePool event raised by the PledgePool contract.
type PledgePoolLiquidatePool struct {
	Pid                     *big.Int
	LiquidationAmountLend   *big.Int
	LiquidationAmountBorrow *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterLiquidatePool is a free log retrieval operation binding the contract event 0x40a58cc34fb5788d695e1e808a735bf11c933c79a9c7af0726d365a2d10d29d3.
//
// Solidity: event LiquidatePool(uint256 indexed pid, uint256 liquidationAmountLend, uint256 liquidationAmountBorrow)
func (_PledgePool *PledgePoolFilterer) FilterLiquidatePool(opts *bind.FilterOpts, pid []*big.Int) (*PledgePoolLiquidatePoolIterator, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _PledgePool.contract.FilterLogs(opts, "LiquidatePool", pidRule)
	if err != nil {
		return nil, err
	}
	return &PledgePoolLiquidatePoolIterator{contract: _PledgePool.contract, event: "LiquidatePool", logs: logs, sub: sub}, nil
}

// WatchLiquidatePool is a free log subscription operation binding the contract event 0x40a58cc34fb5788d695e1e808a735bf11c933c79a9c7af0726d365a2d10d29d3.
//
// Solidity: event LiquidatePool(uint256 indexed pid, uint256 liquidationAmountLend, uint256 liquidationAmountBorrow)
func (_PledgePool *PledgePoolFilterer) WatchLiquidatePool(opts *bind.WatchOpts, sink chan<- *PledgePoolLiquidatePool, pid []*big.Int) (event.Subscription, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _PledgePool.contract.WatchLogs(opts, "LiquidatePool", pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PledgePoolLiquidatePool)
				if err := _PledgePool.contract.UnpackLog(event, "LiquidatePool", log); err != nil {
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

// ParseLiquidatePool is a log parse operation binding the contract event 0x40a58cc34fb5788d695e1e808a735bf11c933c79a9c7af0726d365a2d10d29d3.
//
// Solidity: event LiquidatePool(uint256 indexed pid, uint256 liquidationAmountLend, uint256 liquidationAmountBorrow)
func (_PledgePool *PledgePoolFilterer) ParseLiquidatePool(log types.Log) (*PledgePoolLiquidatePool, error) {
	event := new(PledgePoolLiquidatePool)
	if err := _PledgePool.contract.UnpackLog(event, "LiquidatePool", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PledgePoolOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PledgePool contract.
type PledgePoolOwnershipTransferredIterator struct {
	Event *PledgePoolOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PledgePoolOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PledgePoolOwnershipTransferred)
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
		it.Event = new(PledgePoolOwnershipTransferred)
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
func (it *PledgePoolOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PledgePoolOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PledgePoolOwnershipTransferred represents a OwnershipTransferred event raised by the PledgePool contract.
type PledgePoolOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PledgePool *PledgePoolFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PledgePoolOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PledgePool.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PledgePoolOwnershipTransferredIterator{contract: _PledgePool.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PledgePool *PledgePoolFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PledgePoolOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PledgePool.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PledgePoolOwnershipTransferred)
				if err := _PledgePool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_PledgePool *PledgePoolFilterer) ParseOwnershipTransferred(log types.Log) (*PledgePoolOwnershipTransferred, error) {
	event := new(PledgePoolOwnershipTransferred)
	if err := _PledgePool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PledgePoolRefundBorrowIterator is returned from FilterRefundBorrow and is used to iterate over the raw logs and unpacked data for RefundBorrow events raised by the PledgePool contract.
type PledgePoolRefundBorrowIterator struct {
	Event *PledgePoolRefundBorrow // Event containing the contract specifics and raw log

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
func (it *PledgePoolRefundBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PledgePoolRefundBorrow)
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
		it.Event = new(PledgePoolRefundBorrow)
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
func (it *PledgePoolRefundBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PledgePoolRefundBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PledgePoolRefundBorrow represents a RefundBorrow event raised by the PledgePool contract.
type PledgePoolRefundBorrow struct {
	Pid      *big.Int
	Token    common.Address
	Refunder common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRefundBorrow is a free log retrieval operation binding the contract event 0x445b80a962209ad38e06ec22380c7f6377277223b1ee8d229c2bcb60d80e3d54.
//
// Solidity: event RefundBorrow(uint256 indexed pid, address indexed token, address indexed refunder, uint256 amount)
func (_PledgePool *PledgePoolFilterer) FilterRefundBorrow(opts *bind.FilterOpts, pid []*big.Int, token []common.Address, refunder []common.Address) (*PledgePoolRefundBorrowIterator, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var refunderRule []interface{}
	for _, refunderItem := range refunder {
		refunderRule = append(refunderRule, refunderItem)
	}

	logs, sub, err := _PledgePool.contract.FilterLogs(opts, "RefundBorrow", pidRule, tokenRule, refunderRule)
	if err != nil {
		return nil, err
	}
	return &PledgePoolRefundBorrowIterator{contract: _PledgePool.contract, event: "RefundBorrow", logs: logs, sub: sub}, nil
}

// WatchRefundBorrow is a free log subscription operation binding the contract event 0x445b80a962209ad38e06ec22380c7f6377277223b1ee8d229c2bcb60d80e3d54.
//
// Solidity: event RefundBorrow(uint256 indexed pid, address indexed token, address indexed refunder, uint256 amount)
func (_PledgePool *PledgePoolFilterer) WatchRefundBorrow(opts *bind.WatchOpts, sink chan<- *PledgePoolRefundBorrow, pid []*big.Int, token []common.Address, refunder []common.Address) (event.Subscription, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var refunderRule []interface{}
	for _, refunderItem := range refunder {
		refunderRule = append(refunderRule, refunderItem)
	}

	logs, sub, err := _PledgePool.contract.WatchLogs(opts, "RefundBorrow", pidRule, tokenRule, refunderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PledgePoolRefundBorrow)
				if err := _PledgePool.contract.UnpackLog(event, "RefundBorrow", log); err != nil {
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

// ParseRefundBorrow is a log parse operation binding the contract event 0x445b80a962209ad38e06ec22380c7f6377277223b1ee8d229c2bcb60d80e3d54.
//
// Solidity: event RefundBorrow(uint256 indexed pid, address indexed token, address indexed refunder, uint256 amount)
func (_PledgePool *PledgePoolFilterer) ParseRefundBorrow(log types.Log) (*PledgePoolRefundBorrow, error) {
	event := new(PledgePoolRefundBorrow)
	if err := _PledgePool.contract.UnpackLog(event, "RefundBorrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PledgePoolRefundLendIterator is returned from FilterRefundLend and is used to iterate over the raw logs and unpacked data for RefundLend events raised by the PledgePool contract.
type PledgePoolRefundLendIterator struct {
	Event *PledgePoolRefundLend // Event containing the contract specifics and raw log

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
func (it *PledgePoolRefundLendIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PledgePoolRefundLend)
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
		it.Event = new(PledgePoolRefundLend)
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
func (it *PledgePoolRefundLendIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PledgePoolRefundLendIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PledgePoolRefundLend represents a RefundLend event raised by the PledgePool contract.
type PledgePoolRefundLend struct {
	Pid      *big.Int
	Token    common.Address
	Refunder common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRefundLend is a free log retrieval operation binding the contract event 0x366d54abb3194a1830bd9174cd4ada6d761f18ad3975cdb0895e36f5d2fae03e.
//
// Solidity: event RefundLend(uint256 indexed pid, address indexed token, address indexed refunder, uint256 amount)
func (_PledgePool *PledgePoolFilterer) FilterRefundLend(opts *bind.FilterOpts, pid []*big.Int, token []common.Address, refunder []common.Address) (*PledgePoolRefundLendIterator, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var refunderRule []interface{}
	for _, refunderItem := range refunder {
		refunderRule = append(refunderRule, refunderItem)
	}

	logs, sub, err := _PledgePool.contract.FilterLogs(opts, "RefundLend", pidRule, tokenRule, refunderRule)
	if err != nil {
		return nil, err
	}
	return &PledgePoolRefundLendIterator{contract: _PledgePool.contract, event: "RefundLend", logs: logs, sub: sub}, nil
}

// WatchRefundLend is a free log subscription operation binding the contract event 0x366d54abb3194a1830bd9174cd4ada6d761f18ad3975cdb0895e36f5d2fae03e.
//
// Solidity: event RefundLend(uint256 indexed pid, address indexed token, address indexed refunder, uint256 amount)
func (_PledgePool *PledgePoolFilterer) WatchRefundLend(opts *bind.WatchOpts, sink chan<- *PledgePoolRefundLend, pid []*big.Int, token []common.Address, refunder []common.Address) (event.Subscription, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var refunderRule []interface{}
	for _, refunderItem := range refunder {
		refunderRule = append(refunderRule, refunderItem)
	}

	logs, sub, err := _PledgePool.contract.WatchLogs(opts, "RefundLend", pidRule, tokenRule, refunderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PledgePoolRefundLend)
				if err := _PledgePool.contract.UnpackLog(event, "RefundLend", log); err != nil {
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

// ParseRefundLend is a log parse operation binding the contract event 0x366d54abb3194a1830bd9174cd4ada6d761f18ad3975cdb0895e36f5d2fae03e.
//
// Solidity: event RefundLend(uint256 indexed pid, address indexed token, address indexed refunder, uint256 amount)
func (_PledgePool *PledgePoolFilterer) ParseRefundLend(log types.Log) (*PledgePoolRefundLend, error) {
	event := new(PledgePoolRefundLend)
	if err := _PledgePool.contract.UnpackLog(event, "RefundLend", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PledgePoolSetFeeIterator is returned from FilterSetFee and is used to iterate over the raw logs and unpacked data for SetFee events raised by the PledgePool contract.
type PledgePoolSetFeeIterator struct {
	Event *PledgePoolSetFee // Event containing the contract specifics and raw log

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
func (it *PledgePoolSetFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PledgePoolSetFee)
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
		it.Event = new(PledgePoolSetFee)
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
func (it *PledgePoolSetFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PledgePoolSetFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PledgePoolSetFee represents a SetFee event raised by the PledgePool contract.
type PledgePoolSetFee struct {
	NewLendFee   *big.Int
	NewBorrowFee *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSetFee is a free log retrieval operation binding the contract event 0x032dc6a2d839eb179729a55633fdf1c41a1fc4739394154117005db2b354b9b5.
//
// Solidity: event SetFee(uint256 indexed newLendFee, uint256 indexed newBorrowFee)
func (_PledgePool *PledgePoolFilterer) FilterSetFee(opts *bind.FilterOpts, newLendFee []*big.Int, newBorrowFee []*big.Int) (*PledgePoolSetFeeIterator, error) {

	var newLendFeeRule []interface{}
	for _, newLendFeeItem := range newLendFee {
		newLendFeeRule = append(newLendFeeRule, newLendFeeItem)
	}
	var newBorrowFeeRule []interface{}
	for _, newBorrowFeeItem := range newBorrowFee {
		newBorrowFeeRule = append(newBorrowFeeRule, newBorrowFeeItem)
	}

	logs, sub, err := _PledgePool.contract.FilterLogs(opts, "SetFee", newLendFeeRule, newBorrowFeeRule)
	if err != nil {
		return nil, err
	}
	return &PledgePoolSetFeeIterator{contract: _PledgePool.contract, event: "SetFee", logs: logs, sub: sub}, nil
}

// WatchSetFee is a free log subscription operation binding the contract event 0x032dc6a2d839eb179729a55633fdf1c41a1fc4739394154117005db2b354b9b5.
//
// Solidity: event SetFee(uint256 indexed newLendFee, uint256 indexed newBorrowFee)
func (_PledgePool *PledgePoolFilterer) WatchSetFee(opts *bind.WatchOpts, sink chan<- *PledgePoolSetFee, newLendFee []*big.Int, newBorrowFee []*big.Int) (event.Subscription, error) {

	var newLendFeeRule []interface{}
	for _, newLendFeeItem := range newLendFee {
		newLendFeeRule = append(newLendFeeRule, newLendFeeItem)
	}
	var newBorrowFeeRule []interface{}
	for _, newBorrowFeeItem := range newBorrowFee {
		newBorrowFeeRule = append(newBorrowFeeRule, newBorrowFeeItem)
	}

	logs, sub, err := _PledgePool.contract.WatchLogs(opts, "SetFee", newLendFeeRule, newBorrowFeeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PledgePoolSetFee)
				if err := _PledgePool.contract.UnpackLog(event, "SetFee", log); err != nil {
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

// ParseSetFee is a log parse operation binding the contract event 0x032dc6a2d839eb179729a55633fdf1c41a1fc4739394154117005db2b354b9b5.
//
// Solidity: event SetFee(uint256 indexed newLendFee, uint256 indexed newBorrowFee)
func (_PledgePool *PledgePoolFilterer) ParseSetFee(log types.Log) (*PledgePoolSetFee, error) {
	event := new(PledgePoolSetFee)
	if err := _PledgePool.contract.UnpackLog(event, "SetFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PledgePoolSetFeeAddressIterator is returned from FilterSetFeeAddress and is used to iterate over the raw logs and unpacked data for SetFeeAddress events raised by the PledgePool contract.
type PledgePoolSetFeeAddressIterator struct {
	Event *PledgePoolSetFeeAddress // Event containing the contract specifics and raw log

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
func (it *PledgePoolSetFeeAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PledgePoolSetFeeAddress)
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
		it.Event = new(PledgePoolSetFeeAddress)
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
func (it *PledgePoolSetFeeAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PledgePoolSetFeeAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PledgePoolSetFeeAddress represents a SetFeeAddress event raised by the PledgePool contract.
type PledgePoolSetFeeAddress struct {
	OldFeeAddress common.Address
	NewFeeAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSetFeeAddress is a free log retrieval operation binding the contract event 0xd44190acf9d04bdb5d3a1aafff7e6dee8b40b93dfb8c5d3f0eea4b9f4539c3f7.
//
// Solidity: event SetFeeAddress(address indexed oldFeeAddress, address indexed newFeeAddress)
func (_PledgePool *PledgePoolFilterer) FilterSetFeeAddress(opts *bind.FilterOpts, oldFeeAddress []common.Address, newFeeAddress []common.Address) (*PledgePoolSetFeeAddressIterator, error) {

	var oldFeeAddressRule []interface{}
	for _, oldFeeAddressItem := range oldFeeAddress {
		oldFeeAddressRule = append(oldFeeAddressRule, oldFeeAddressItem)
	}
	var newFeeAddressRule []interface{}
	for _, newFeeAddressItem := range newFeeAddress {
		newFeeAddressRule = append(newFeeAddressRule, newFeeAddressItem)
	}

	logs, sub, err := _PledgePool.contract.FilterLogs(opts, "SetFeeAddress", oldFeeAddressRule, newFeeAddressRule)
	if err != nil {
		return nil, err
	}
	return &PledgePoolSetFeeAddressIterator{contract: _PledgePool.contract, event: "SetFeeAddress", logs: logs, sub: sub}, nil
}

// WatchSetFeeAddress is a free log subscription operation binding the contract event 0xd44190acf9d04bdb5d3a1aafff7e6dee8b40b93dfb8c5d3f0eea4b9f4539c3f7.
//
// Solidity: event SetFeeAddress(address indexed oldFeeAddress, address indexed newFeeAddress)
func (_PledgePool *PledgePoolFilterer) WatchSetFeeAddress(opts *bind.WatchOpts, sink chan<- *PledgePoolSetFeeAddress, oldFeeAddress []common.Address, newFeeAddress []common.Address) (event.Subscription, error) {

	var oldFeeAddressRule []interface{}
	for _, oldFeeAddressItem := range oldFeeAddress {
		oldFeeAddressRule = append(oldFeeAddressRule, oldFeeAddressItem)
	}
	var newFeeAddressRule []interface{}
	for _, newFeeAddressItem := range newFeeAddress {
		newFeeAddressRule = append(newFeeAddressRule, newFeeAddressItem)
	}

	logs, sub, err := _PledgePool.contract.WatchLogs(opts, "SetFeeAddress", oldFeeAddressRule, newFeeAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PledgePoolSetFeeAddress)
				if err := _PledgePool.contract.UnpackLog(event, "SetFeeAddress", log); err != nil {
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

// ParseSetFeeAddress is a log parse operation binding the contract event 0xd44190acf9d04bdb5d3a1aafff7e6dee8b40b93dfb8c5d3f0eea4b9f4539c3f7.
//
// Solidity: event SetFeeAddress(address indexed oldFeeAddress, address indexed newFeeAddress)
func (_PledgePool *PledgePoolFilterer) ParseSetFeeAddress(log types.Log) (*PledgePoolSetFeeAddress, error) {
	event := new(PledgePoolSetFeeAddress)
	if err := _PledgePool.contract.UnpackLog(event, "SetFeeAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PledgePoolSetGlobalPausedIterator is returned from FilterSetGlobalPaused and is used to iterate over the raw logs and unpacked data for SetGlobalPaused events raised by the PledgePool contract.
type PledgePoolSetGlobalPausedIterator struct {
	Event *PledgePoolSetGlobalPaused // Event containing the contract specifics and raw log

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
func (it *PledgePoolSetGlobalPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PledgePoolSetGlobalPaused)
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
		it.Event = new(PledgePoolSetGlobalPaused)
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
func (it *PledgePoolSetGlobalPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PledgePoolSetGlobalPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PledgePoolSetGlobalPaused represents a SetGlobalPaused event raised by the PledgePool contract.
type PledgePoolSetGlobalPaused struct {
	OldPaused bool
	NewPaused bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSetGlobalPaused is a free log retrieval operation binding the contract event 0x989810a81181eea4877dd45b6d6eea562c0ad665780a7c2738e0ea7504287693.
//
// Solidity: event SetGlobalPaused(bool indexed oldPaused, bool indexed newPaused)
func (_PledgePool *PledgePoolFilterer) FilterSetGlobalPaused(opts *bind.FilterOpts, oldPaused []bool, newPaused []bool) (*PledgePoolSetGlobalPausedIterator, error) {

	var oldPausedRule []interface{}
	for _, oldPausedItem := range oldPaused {
		oldPausedRule = append(oldPausedRule, oldPausedItem)
	}
	var newPausedRule []interface{}
	for _, newPausedItem := range newPaused {
		newPausedRule = append(newPausedRule, newPausedItem)
	}

	logs, sub, err := _PledgePool.contract.FilterLogs(opts, "SetGlobalPaused", oldPausedRule, newPausedRule)
	if err != nil {
		return nil, err
	}
	return &PledgePoolSetGlobalPausedIterator{contract: _PledgePool.contract, event: "SetGlobalPaused", logs: logs, sub: sub}, nil
}

// WatchSetGlobalPaused is a free log subscription operation binding the contract event 0x989810a81181eea4877dd45b6d6eea562c0ad665780a7c2738e0ea7504287693.
//
// Solidity: event SetGlobalPaused(bool indexed oldPaused, bool indexed newPaused)
func (_PledgePool *PledgePoolFilterer) WatchSetGlobalPaused(opts *bind.WatchOpts, sink chan<- *PledgePoolSetGlobalPaused, oldPaused []bool, newPaused []bool) (event.Subscription, error) {

	var oldPausedRule []interface{}
	for _, oldPausedItem := range oldPaused {
		oldPausedRule = append(oldPausedRule, oldPausedItem)
	}
	var newPausedRule []interface{}
	for _, newPausedItem := range newPaused {
		newPausedRule = append(newPausedRule, newPausedItem)
	}

	logs, sub, err := _PledgePool.contract.WatchLogs(opts, "SetGlobalPaused", oldPausedRule, newPausedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PledgePoolSetGlobalPaused)
				if err := _PledgePool.contract.UnpackLog(event, "SetGlobalPaused", log); err != nil {
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

// ParseSetGlobalPaused is a log parse operation binding the contract event 0x989810a81181eea4877dd45b6d6eea562c0ad665780a7c2738e0ea7504287693.
//
// Solidity: event SetGlobalPaused(bool indexed oldPaused, bool indexed newPaused)
func (_PledgePool *PledgePoolFilterer) ParseSetGlobalPaused(log types.Log) (*PledgePoolSetGlobalPaused, error) {
	event := new(PledgePoolSetGlobalPaused)
	if err := _PledgePool.contract.UnpackLog(event, "SetGlobalPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PledgePoolSetMinAmountIterator is returned from FilterSetMinAmount and is used to iterate over the raw logs and unpacked data for SetMinAmount events raised by the PledgePool contract.
type PledgePoolSetMinAmountIterator struct {
	Event *PledgePoolSetMinAmount // Event containing the contract specifics and raw log

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
func (it *PledgePoolSetMinAmountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PledgePoolSetMinAmount)
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
		it.Event = new(PledgePoolSetMinAmount)
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
func (it *PledgePoolSetMinAmountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PledgePoolSetMinAmountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PledgePoolSetMinAmount represents a SetMinAmount event raised by the PledgePool contract.
type PledgePoolSetMinAmount struct {
	OldMinAmount *big.Int
	NewMinAmount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSetMinAmount is a free log retrieval operation binding the contract event 0xfa6189b739625142c695478e9d0095a1cb9e6fad92ad8a727e0055a5cc85b06b.
//
// Solidity: event SetMinAmount(uint256 indexed oldMinAmount, uint256 indexed newMinAmount)
func (_PledgePool *PledgePoolFilterer) FilterSetMinAmount(opts *bind.FilterOpts, oldMinAmount []*big.Int, newMinAmount []*big.Int) (*PledgePoolSetMinAmountIterator, error) {

	var oldMinAmountRule []interface{}
	for _, oldMinAmountItem := range oldMinAmount {
		oldMinAmountRule = append(oldMinAmountRule, oldMinAmountItem)
	}
	var newMinAmountRule []interface{}
	for _, newMinAmountItem := range newMinAmount {
		newMinAmountRule = append(newMinAmountRule, newMinAmountItem)
	}

	logs, sub, err := _PledgePool.contract.FilterLogs(opts, "SetMinAmount", oldMinAmountRule, newMinAmountRule)
	if err != nil {
		return nil, err
	}
	return &PledgePoolSetMinAmountIterator{contract: _PledgePool.contract, event: "SetMinAmount", logs: logs, sub: sub}, nil
}

// WatchSetMinAmount is a free log subscription operation binding the contract event 0xfa6189b739625142c695478e9d0095a1cb9e6fad92ad8a727e0055a5cc85b06b.
//
// Solidity: event SetMinAmount(uint256 indexed oldMinAmount, uint256 indexed newMinAmount)
func (_PledgePool *PledgePoolFilterer) WatchSetMinAmount(opts *bind.WatchOpts, sink chan<- *PledgePoolSetMinAmount, oldMinAmount []*big.Int, newMinAmount []*big.Int) (event.Subscription, error) {

	var oldMinAmountRule []interface{}
	for _, oldMinAmountItem := range oldMinAmount {
		oldMinAmountRule = append(oldMinAmountRule, oldMinAmountItem)
	}
	var newMinAmountRule []interface{}
	for _, newMinAmountItem := range newMinAmount {
		newMinAmountRule = append(newMinAmountRule, newMinAmountItem)
	}

	logs, sub, err := _PledgePool.contract.WatchLogs(opts, "SetMinAmount", oldMinAmountRule, newMinAmountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PledgePoolSetMinAmount)
				if err := _PledgePool.contract.UnpackLog(event, "SetMinAmount", log); err != nil {
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

// ParseSetMinAmount is a log parse operation binding the contract event 0xfa6189b739625142c695478e9d0095a1cb9e6fad92ad8a727e0055a5cc85b06b.
//
// Solidity: event SetMinAmount(uint256 indexed oldMinAmount, uint256 indexed newMinAmount)
func (_PledgePool *PledgePoolFilterer) ParseSetMinAmount(log types.Log) (*PledgePoolSetMinAmount, error) {
	event := new(PledgePoolSetMinAmount)
	if err := _PledgePool.contract.UnpackLog(event, "SetMinAmount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PledgePoolSetOracleIterator is returned from FilterSetOracle and is used to iterate over the raw logs and unpacked data for SetOracle events raised by the PledgePool contract.
type PledgePoolSetOracleIterator struct {
	Event *PledgePoolSetOracle // Event containing the contract specifics and raw log

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
func (it *PledgePoolSetOracleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PledgePoolSetOracle)
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
		it.Event = new(PledgePoolSetOracle)
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
func (it *PledgePoolSetOracleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PledgePoolSetOracleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PledgePoolSetOracle represents a SetOracle event raised by the PledgePool contract.
type PledgePoolSetOracle struct {
	OldOracle common.Address
	NewOracle common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSetOracle is a free log retrieval operation binding the contract event 0xb7261e9c33aa7c56209c3bf60b424a8f9551ce28876c0ab3d0c487695e943487.
//
// Solidity: event SetOracle(address indexed oldOracle, address indexed newOracle)
func (_PledgePool *PledgePoolFilterer) FilterSetOracle(opts *bind.FilterOpts, oldOracle []common.Address, newOracle []common.Address) (*PledgePoolSetOracleIterator, error) {

	var oldOracleRule []interface{}
	for _, oldOracleItem := range oldOracle {
		oldOracleRule = append(oldOracleRule, oldOracleItem)
	}
	var newOracleRule []interface{}
	for _, newOracleItem := range newOracle {
		newOracleRule = append(newOracleRule, newOracleItem)
	}

	logs, sub, err := _PledgePool.contract.FilterLogs(opts, "SetOracle", oldOracleRule, newOracleRule)
	if err != nil {
		return nil, err
	}
	return &PledgePoolSetOracleIterator{contract: _PledgePool.contract, event: "SetOracle", logs: logs, sub: sub}, nil
}

// WatchSetOracle is a free log subscription operation binding the contract event 0xb7261e9c33aa7c56209c3bf60b424a8f9551ce28876c0ab3d0c487695e943487.
//
// Solidity: event SetOracle(address indexed oldOracle, address indexed newOracle)
func (_PledgePool *PledgePoolFilterer) WatchSetOracle(opts *bind.WatchOpts, sink chan<- *PledgePoolSetOracle, oldOracle []common.Address, newOracle []common.Address) (event.Subscription, error) {

	var oldOracleRule []interface{}
	for _, oldOracleItem := range oldOracle {
		oldOracleRule = append(oldOracleRule, oldOracleItem)
	}
	var newOracleRule []interface{}
	for _, newOracleItem := range newOracle {
		newOracleRule = append(newOracleRule, newOracleItem)
	}

	logs, sub, err := _PledgePool.contract.WatchLogs(opts, "SetOracle", oldOracleRule, newOracleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PledgePoolSetOracle)
				if err := _PledgePool.contract.UnpackLog(event, "SetOracle", log); err != nil {
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

// ParseSetOracle is a log parse operation binding the contract event 0xb7261e9c33aa7c56209c3bf60b424a8f9551ce28876c0ab3d0c487695e943487.
//
// Solidity: event SetOracle(address indexed oldOracle, address indexed newOracle)
func (_PledgePool *PledgePoolFilterer) ParseSetOracle(log types.Log) (*PledgePoolSetOracle, error) {
	event := new(PledgePoolSetOracle)
	if err := _PledgePool.contract.UnpackLog(event, "SetOracle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PledgePoolSetSwapRouterAddressIterator is returned from FilterSetSwapRouterAddress and is used to iterate over the raw logs and unpacked data for SetSwapRouterAddress events raised by the PledgePool contract.
type PledgePoolSetSwapRouterAddressIterator struct {
	Event *PledgePoolSetSwapRouterAddress // Event containing the contract specifics and raw log

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
func (it *PledgePoolSetSwapRouterAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PledgePoolSetSwapRouterAddress)
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
		it.Event = new(PledgePoolSetSwapRouterAddress)
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
func (it *PledgePoolSetSwapRouterAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PledgePoolSetSwapRouterAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PledgePoolSetSwapRouterAddress represents a SetSwapRouterAddress event raised by the PledgePool contract.
type PledgePoolSetSwapRouterAddress struct {
	OldSwapAddress common.Address
	NewSwapAddress common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSetSwapRouterAddress is a free log retrieval operation binding the contract event 0x4558149b3c5427365f76d4ff19bef30aba41f17e5e601d4661330d8d2b687627.
//
// Solidity: event SetSwapRouterAddress(address indexed oldSwapAddress, address indexed newSwapAddress)
func (_PledgePool *PledgePoolFilterer) FilterSetSwapRouterAddress(opts *bind.FilterOpts, oldSwapAddress []common.Address, newSwapAddress []common.Address) (*PledgePoolSetSwapRouterAddressIterator, error) {

	var oldSwapAddressRule []interface{}
	for _, oldSwapAddressItem := range oldSwapAddress {
		oldSwapAddressRule = append(oldSwapAddressRule, oldSwapAddressItem)
	}
	var newSwapAddressRule []interface{}
	for _, newSwapAddressItem := range newSwapAddress {
		newSwapAddressRule = append(newSwapAddressRule, newSwapAddressItem)
	}

	logs, sub, err := _PledgePool.contract.FilterLogs(opts, "SetSwapRouterAddress", oldSwapAddressRule, newSwapAddressRule)
	if err != nil {
		return nil, err
	}
	return &PledgePoolSetSwapRouterAddressIterator{contract: _PledgePool.contract, event: "SetSwapRouterAddress", logs: logs, sub: sub}, nil
}

// WatchSetSwapRouterAddress is a free log subscription operation binding the contract event 0x4558149b3c5427365f76d4ff19bef30aba41f17e5e601d4661330d8d2b687627.
//
// Solidity: event SetSwapRouterAddress(address indexed oldSwapAddress, address indexed newSwapAddress)
func (_PledgePool *PledgePoolFilterer) WatchSetSwapRouterAddress(opts *bind.WatchOpts, sink chan<- *PledgePoolSetSwapRouterAddress, oldSwapAddress []common.Address, newSwapAddress []common.Address) (event.Subscription, error) {

	var oldSwapAddressRule []interface{}
	for _, oldSwapAddressItem := range oldSwapAddress {
		oldSwapAddressRule = append(oldSwapAddressRule, oldSwapAddressItem)
	}
	var newSwapAddressRule []interface{}
	for _, newSwapAddressItem := range newSwapAddress {
		newSwapAddressRule = append(newSwapAddressRule, newSwapAddressItem)
	}

	logs, sub, err := _PledgePool.contract.WatchLogs(opts, "SetSwapRouterAddress", oldSwapAddressRule, newSwapAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PledgePoolSetSwapRouterAddress)
				if err := _PledgePool.contract.UnpackLog(event, "SetSwapRouterAddress", log); err != nil {
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

// ParseSetSwapRouterAddress is a log parse operation binding the contract event 0x4558149b3c5427365f76d4ff19bef30aba41f17e5e601d4661330d8d2b687627.
//
// Solidity: event SetSwapRouterAddress(address indexed oldSwapAddress, address indexed newSwapAddress)
func (_PledgePool *PledgePoolFilterer) ParseSetSwapRouterAddress(log types.Log) (*PledgePoolSetSwapRouterAddress, error) {
	event := new(PledgePoolSetSwapRouterAddress)
	if err := _PledgePool.contract.UnpackLog(event, "SetSwapRouterAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PledgePoolSettlePoolIterator is returned from FilterSettlePool and is used to iterate over the raw logs and unpacked data for SettlePool events raised by the PledgePool contract.
type PledgePoolSettlePoolIterator struct {
	Event *PledgePoolSettlePool // Event containing the contract specifics and raw log

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
func (it *PledgePoolSettlePoolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PledgePoolSettlePool)
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
		it.Event = new(PledgePoolSettlePool)
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
func (it *PledgePoolSettlePoolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PledgePoolSettlePoolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PledgePoolSettlePool represents a SettlePool event raised by the PledgePool contract.
type PledgePoolSettlePool struct {
	Pid                *big.Int
	SettleAmountLend   *big.Int
	SettleAmountBorrow *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterSettlePool is a free log retrieval operation binding the contract event 0xecb83b282735569068a276ac3b9f868bdb0048d952568d94f8a79a83eb001bbe.
//
// Solidity: event SettlePool(uint256 indexed pid, uint256 settleAmountLend, uint256 settleAmountBorrow)
func (_PledgePool *PledgePoolFilterer) FilterSettlePool(opts *bind.FilterOpts, pid []*big.Int) (*PledgePoolSettlePoolIterator, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _PledgePool.contract.FilterLogs(opts, "SettlePool", pidRule)
	if err != nil {
		return nil, err
	}
	return &PledgePoolSettlePoolIterator{contract: _PledgePool.contract, event: "SettlePool", logs: logs, sub: sub}, nil
}

// WatchSettlePool is a free log subscription operation binding the contract event 0xecb83b282735569068a276ac3b9f868bdb0048d952568d94f8a79a83eb001bbe.
//
// Solidity: event SettlePool(uint256 indexed pid, uint256 settleAmountLend, uint256 settleAmountBorrow)
func (_PledgePool *PledgePoolFilterer) WatchSettlePool(opts *bind.WatchOpts, sink chan<- *PledgePoolSettlePool, pid []*big.Int) (event.Subscription, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _PledgePool.contract.WatchLogs(opts, "SettlePool", pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PledgePoolSettlePool)
				if err := _PledgePool.contract.UnpackLog(event, "SettlePool", log); err != nil {
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

// ParseSettlePool is a log parse operation binding the contract event 0xecb83b282735569068a276ac3b9f868bdb0048d952568d94f8a79a83eb001bbe.
//
// Solidity: event SettlePool(uint256 indexed pid, uint256 settleAmountLend, uint256 settleAmountBorrow)
func (_PledgePool *PledgePoolFilterer) ParseSettlePool(log types.Log) (*PledgePoolSettlePool, error) {
	event := new(PledgePoolSettlePool)
	if err := _PledgePool.contract.UnpackLog(event, "SettlePool", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
