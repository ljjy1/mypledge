package handler

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"

	"github.com/go-dev-frame/sponge/pkg/gin/response"
	"github.com/go-dev-frame/sponge/pkg/logger"

	"pledge-be/internal/contract/bindcode"
	"pledge-be/internal/ecode"
	"pledge-be/internal/types"
)

// PoolLend 用户出借资产到指定池
func (h *contractHandler) PoolLend(c *gin.Context) {
	form := &types.PoolAmountRequest{}
	if err := c.ShouldBindJSON(form); err != nil {
		response.Error(c, ecode.InvalidParams)
		return
	}
	pid, err := parseBigInt(form.PoolID)
	if err != nil {
		response.Error(c, ecode.ErrPoolInvalidPoolID)
		return
	}
	lendAmount, err := parseBigInt(form.Amount)
	if err != nil {
		response.Error(c, ecode.ErrPoolInvalidAmount)
		return
	}

	client, auth, _, cleanup, err := prepareTransactor(c, form.RpcURL, form.PrivateKey)
	if err != nil {
		response.Error(c, ecode.ErrPoolRPCConnect)
		return
	}
	defer cleanup()

	poolContract, err := bindcode.NewPledgePool(common.HexToAddress(form.PoolContractAddress), client)
	if err != nil {
		response.Error(c, ecode.ErrPoolInvalidAddr)
		return
	}

	tx, err := poolContract.Lend(auth, pid, lendAmount)
	if err != nil {
		logger.Error("PoolLend error", logger.Err(err))
		response.Error(c, ecode.ErrPoolSendTx)
		return
	}
	response.Success(c, gin.H{"txHash": tx.Hash().Hex()})
}

// PoolBorrow 用户从指定池借入资产
func (h *contractHandler) PoolBorrow(c *gin.Context) {
	form := &types.PoolBorrowRequest{}
	if err := c.ShouldBindJSON(form); err != nil {
		response.Error(c, ecode.InvalidParams)
		return
	}
	pid, err := parseBigInt(form.PoolID)
	if err != nil {
		response.Error(c, ecode.ErrPoolInvalidPoolID)
		return
	}
	borrowAmount, err := parseBigInt(form.BorrowTokenAmount)
	if err != nil {
		response.Error(c, ecode.ErrPoolInvalidAmount)
		return
	}

	client, auth, _, cleanup, err := prepareTransactor(c, form.RpcURL, form.PrivateKey)
	if err != nil {
		response.Error(c, ecode.ErrPoolRPCConnect)
		return
	}
	defer cleanup()

	poolContract, err := bindcode.NewPledgePool(common.HexToAddress(form.PoolContractAddress), client)
	if err != nil {
		response.Error(c, ecode.ErrPoolInvalidAddr)
		return
	}

	tx, err := poolContract.Borrow(auth, pid, borrowAmount)
	if err != nil {
		logger.Error("PoolBorrow error", logger.Err(err))
		response.Error(c, ecode.ErrPoolSendTx)
		return
	}
	response.Success(c, gin.H{"txHash": tx.Hash().Hex()})
}

// PoolSettle 结算指定池
func (h *contractHandler) PoolSettle(c *gin.Context) {
	form := &types.PoolWriteRequest{}
	if err := c.ShouldBindJSON(form); err != nil {
		response.Error(c, ecode.InvalidParams)
		return
	}
	pid, err := parseBigInt(form.PoolID)
	if err != nil {
		response.Error(c, ecode.ErrPoolInvalidPoolID)
		return
	}

	client, auth, _, cleanup, err := prepareTransactor(c, form.RpcURL, form.PrivateKey)
	if err != nil {
		response.Error(c, ecode.ErrPoolRPCConnect)
		return
	}
	defer cleanup()

	poolContract, err := bindcode.NewPledgePool(common.HexToAddress(form.PoolContractAddress), client)
	if err != nil {
		response.Error(c, ecode.ErrPoolInvalidAddr)
		return
	}

	tx, err := poolContract.SettlePool(auth, pid)
	if err != nil {
		logger.Error("PoolSettle error", logger.Err(err))
		response.Error(c, ecode.ErrPoolSendTx)
		return
	}
	response.Success(c, gin.H{"txHash": tx.Hash().Hex()})
}

// PoolFinish 完成指定池
func (h *contractHandler) PoolFinish(c *gin.Context) {
	form := &types.PoolWriteRequest{}
	if err := c.ShouldBindJSON(form); err != nil {
		response.Error(c, ecode.InvalidParams)
		return
	}
	pid, err := parseBigInt(form.PoolID)
	if err != nil {
		response.Error(c, ecode.ErrPoolInvalidPoolID)
		return
	}

	client, auth, _, cleanup, err := prepareTransactor(c, form.RpcURL, form.PrivateKey)
	if err != nil {
		response.Error(c, ecode.ErrPoolRPCConnect)
		return
	}
	defer cleanup()

	poolContract, err := bindcode.NewPledgePool(common.HexToAddress(form.PoolContractAddress), client)
	if err != nil {
		response.Error(c, ecode.ErrPoolInvalidAddr)
		return
	}

	tx, err := poolContract.FinishPool(auth, pid)
	if err != nil {
		logger.Error("PoolFinish error", logger.Err(err))
		response.Error(c, ecode.ErrPoolSendTx)
		return
	}
	response.Success(c, gin.H{"txHash": tx.Hash().Hex()})
}

// PoolLiquidate 清算指定池
func (h *contractHandler) PoolLiquidate(c *gin.Context) {
	form := &types.PoolWriteRequest{}
	if err := c.ShouldBindJSON(form); err != nil {
		response.Error(c, ecode.InvalidParams)
		return
	}
	pid, err := parseBigInt(form.PoolID)
	if err != nil {
		response.Error(c, ecode.ErrPoolInvalidPoolID)
		return
	}

	client, auth, _, cleanup, err := prepareTransactor(c, form.RpcURL, form.PrivateKey)
	if err != nil {
		response.Error(c, ecode.ErrPoolRPCConnect)
		return
	}
	defer cleanup()

	poolContract, err := bindcode.NewPledgePool(common.HexToAddress(form.PoolContractAddress), client)
	if err != nil {
		response.Error(c, ecode.ErrPoolInvalidAddr)
		return
	}

	tx, err := poolContract.LiquidatePool(auth, pid)
	if err != nil {
		logger.Error("PoolLiquidate error", logger.Err(err))
		response.Error(c, ecode.ErrPoolSendTx)
		return
	}
	response.Success(c, gin.H{"txHash": tx.Hash().Hex()})
}

// PoolRefundBorrow 归还借入资产
func (h *contractHandler) PoolRefundBorrow(c *gin.Context) {
	form := &types.PoolWriteRequest{}
	if err := c.ShouldBindJSON(form); err != nil {
		response.Error(c, ecode.InvalidParams)
		return
	}
	pid, err := parseBigInt(form.PoolID)
	if err != nil {
		response.Error(c, ecode.ErrPoolInvalidPoolID)
		return
	}

	client, auth, _, cleanup, err := prepareTransactor(c, form.RpcURL, form.PrivateKey)
	if err != nil {
		response.Error(c, ecode.ErrPoolRPCConnect)
		return
	}
	defer cleanup()

	poolContract, err := bindcode.NewPledgePool(common.HexToAddress(form.PoolContractAddress), client)
	if err != nil {
		response.Error(c, ecode.ErrPoolInvalidAddr)
		return
	}

	tx, err := poolContract.RefundBorrow(auth, pid)
	if err != nil {
		logger.Error("PoolRefundBorrow error", logger.Err(err))
		response.Error(c, ecode.ErrPoolSendTx)
		return
	}
	response.Success(c, gin.H{"txHash": tx.Hash().Hex()})
}

// PoolRefundLend 归还出借资产
func (h *contractHandler) PoolRefundLend(c *gin.Context) {
	form := &types.PoolWriteRequest{}
	if err := c.ShouldBindJSON(form); err != nil {
		response.Error(c, ecode.InvalidParams)
		return
	}
	pid, err := parseBigInt(form.PoolID)
	if err != nil {
		response.Error(c, ecode.ErrPoolInvalidPoolID)
		return
	}

	client, auth, _, cleanup, err := prepareTransactor(c, form.RpcURL, form.PrivateKey)
	if err != nil {
		response.Error(c, ecode.ErrPoolRPCConnect)
		return
	}
	defer cleanup()

	poolContract, err := bindcode.NewPledgePool(common.HexToAddress(form.PoolContractAddress), client)
	if err != nil {
		response.Error(c, ecode.ErrPoolInvalidAddr)
		return
	}

	tx, err := poolContract.RefundLend(auth, pid)
	if err != nil {
		logger.Error("PoolRefundLend error", logger.Err(err))
		response.Error(c, ecode.ErrPoolSendTx)
		return
	}
	response.Success(c, gin.H{"txHash": tx.Hash().Hex()})
}

// PoolClaimBorrow 领取借入资产
func (h *contractHandler) PoolClaimBorrow(c *gin.Context) {
	form := &types.PoolWriteRequest{}
	if err := c.ShouldBindJSON(form); err != nil {
		response.Error(c, ecode.InvalidParams)
		return
	}
	pid, err := parseBigInt(form.PoolID)
	if err != nil {
		response.Error(c, ecode.ErrPoolInvalidPoolID)
		return
	}

	client, auth, _, cleanup, err := prepareTransactor(c, form.RpcURL, form.PrivateKey)
	if err != nil {
		response.Error(c, ecode.ErrPoolRPCConnect)
		return
	}
	defer cleanup()

	poolContract, err := bindcode.NewPledgePool(common.HexToAddress(form.PoolContractAddress), client)
	if err != nil {
		response.Error(c, ecode.ErrPoolInvalidAddr)
		return
	}

	tx, err := poolContract.ClaimBorrow(auth, pid)
	if err != nil {
		logger.Error("PoolClaimBorrow error", logger.Err(err))
		response.Error(c, ecode.ErrPoolSendTx)
		return
	}
	response.Success(c, gin.H{"txHash": tx.Hash().Hex()})
}

// PoolClaimLendDebtToken 领取出借债务代币
func (h *contractHandler) PoolClaimLendDebtToken(c *gin.Context) {
	form := &types.PoolWriteRequest{}
	if err := c.ShouldBindJSON(form); err != nil {
		response.Error(c, ecode.InvalidParams)
		return
	}
	pid, err := parseBigInt(form.PoolID)
	if err != nil {
		response.Error(c, ecode.ErrPoolInvalidPoolID)
		return
	}

	client, auth, _, cleanup, err := prepareTransactor(c, form.RpcURL, form.PrivateKey)
	if err != nil {
		response.Error(c, ecode.ErrPoolRPCConnect)
		return
	}
	defer cleanup()

	poolContract, err := bindcode.NewPledgePool(common.HexToAddress(form.PoolContractAddress), client)
	if err != nil {
		response.Error(c, ecode.ErrPoolInvalidAddr)
		return
	}

	tx, err := poolContract.ClaimLendDebtToken(auth, pid)
	if err != nil {
		logger.Error("PoolClaimLendDebtToken error", logger.Err(err))
		response.Error(c, ecode.ErrPoolSendTx)
		return
	}
	response.Success(c, gin.H{"txHash": tx.Hash().Hex()})
}

// PoolDestroyBorrowDebtToken 销毁借入债务代币
func (h *contractHandler) PoolDestroyBorrowDebtToken(c *gin.Context) {
	form := &types.PoolDestroyDebtRequest{}
	if err := c.ShouldBindJSON(form); err != nil {
		response.Error(c, ecode.InvalidParams)
		return
	}
	pid, err := parseBigInt(form.PoolID)
	if err != nil {
		response.Error(c, ecode.ErrPoolInvalidPoolID)
		return
	}
	amount, err := parseBigInt(form.Amount)
	if err != nil {
		response.Error(c, ecode.ErrPoolInvalidAmount)
		return
	}

	client, auth, _, cleanup, err := prepareTransactor(c, form.RpcURL, form.PrivateKey)
	if err != nil {
		response.Error(c, ecode.ErrPoolRPCConnect)
		return
	}
	defer cleanup()

	poolContract, err := bindcode.NewPledgePool(common.HexToAddress(form.PoolContractAddress), client)
	if err != nil {
		response.Error(c, ecode.ErrPoolInvalidAddr)
		return
	}

	tx, err := poolContract.DestroyBorrowDebtToken(auth, pid, amount)
	if err != nil {
		logger.Error("PoolDestroyBorrowDebtToken error", logger.Err(err))
		response.Error(c, ecode.ErrPoolSendTx)
		return
	}
	response.Success(c, gin.H{"txHash": tx.Hash().Hex()})
}

// PoolDestroyLendDebtToken 销毁出借债务代币
func (h *contractHandler) PoolDestroyLendDebtToken(c *gin.Context) {
	form := &types.PoolDestroyDebtRequest{}
	if err := c.ShouldBindJSON(form); err != nil {
		response.Error(c, ecode.InvalidParams)
		return
	}
	pid, err := parseBigInt(form.PoolID)
	if err != nil {
		response.Error(c, ecode.ErrPoolInvalidPoolID)
		return
	}
	amount, err := parseBigInt(form.Amount)
	if err != nil {
		response.Error(c, ecode.ErrPoolInvalidAmount)
		return
	}

	client, auth, _, cleanup, err := prepareTransactor(c, form.RpcURL, form.PrivateKey)
	if err != nil {
		response.Error(c, ecode.ErrPoolRPCConnect)
		return
	}
	defer cleanup()

	poolContract, err := bindcode.NewPledgePool(common.HexToAddress(form.PoolContractAddress), client)
	if err != nil {
		response.Error(c, ecode.ErrPoolInvalidAddr)
		return
	}

	tx, err := poolContract.DestroyLendDebtToken(auth, pid, amount)
	if err != nil {
		logger.Error("PoolDestroyLendDebtToken error", logger.Err(err))
		response.Error(c, ecode.ErrPoolSendTx)
		return
	}
	response.Success(c, gin.H{"txHash": tx.Hash().Hex()})
}

// PoolSetFee 设置费率（管理员）
func (h *contractHandler) PoolSetFee(c *gin.Context) {
	form := &types.PoolSetFeeRequest{}
	if err := c.ShouldBindJSON(form); err != nil {
		response.Error(c, ecode.InvalidParams)
		return
	}
	lendFee, err := parseBigInt(form.LendFee)
	if err != nil {
		response.Error(c, ecode.ErrPoolInvalidAmount)
		return
	}
	borrowFee, err := parseBigInt(form.BorrowFee)
	if err != nil {
		response.Error(c, ecode.ErrPoolInvalidAmount)
		return
	}

	client, auth, _, cleanup, err := prepareTransactor(c, form.RpcURL, form.PrivateKey)
	if err != nil {
		response.Error(c, ecode.ErrPoolRPCConnect)
		return
	}
	defer cleanup()

	poolContract, err := bindcode.NewPledgePool(common.HexToAddress(form.PoolContractAddress), client)
	if err != nil {
		response.Error(c, ecode.ErrPoolInvalidAddr)
		return
	}

	tx, err := poolContract.SetFee(auth, lendFee, borrowFee)
	if err != nil {
		logger.Error("PoolSetFee error", logger.Err(err))
		response.Error(c, ecode.ErrPoolSendTx)
		return
	}
	response.Success(c, gin.H{"txHash": tx.Hash().Hex()})
}

// PoolSetFeeAddress 设置手续费接收地址（管理员）
func (h *contractHandler) PoolSetFeeAddress(c *gin.Context) {
	form := &types.PoolSetAddressRequest{}
	if err := c.ShouldBindJSON(form); err != nil {
		response.Error(c, ecode.InvalidParams)
		return
	}

	client, auth, _, cleanup, err := prepareTransactor(c, form.RpcURL, form.PrivateKey)
	if err != nil {
		response.Error(c, ecode.ErrPoolRPCConnect)
		return
	}
	defer cleanup()

	poolContract, err := bindcode.NewPledgePool(common.HexToAddress(form.PoolContractAddress), client)
	if err != nil {
		response.Error(c, ecode.ErrPoolInvalidAddr)
		return
	}

	tx, err := poolContract.SetFeeAddress(auth, common.HexToAddress(form.NewAddress))
	if err != nil {
		logger.Error("PoolSetFeeAddress error", logger.Err(err))
		response.Error(c, ecode.ErrPoolSendTx)
		return
	}
	response.Success(c, gin.H{"txHash": tx.Hash().Hex()})
}

// PoolSetOracle 设置预言机地址（管理员）
func (h *contractHandler) PoolSetOracle(c *gin.Context) {
	form := &types.PoolSetAddressRequest{}
	if err := c.ShouldBindJSON(form); err != nil {
		response.Error(c, ecode.InvalidParams)
		return
	}

	client, auth, _, cleanup, err := prepareTransactor(c, form.RpcURL, form.PrivateKey)
	if err != nil {
		response.Error(c, ecode.ErrPoolRPCConnect)
		return
	}
	defer cleanup()

	poolContract, err := bindcode.NewPledgePool(common.HexToAddress(form.PoolContractAddress), client)
	if err != nil {
		response.Error(c, ecode.ErrPoolInvalidAddr)
		return
	}

	tx, err := poolContract.SetOracle(auth, common.HexToAddress(form.NewAddress))
	if err != nil {
		logger.Error("PoolSetOracle error", logger.Err(err))
		response.Error(c, ecode.ErrPoolSendTx)
		return
	}
	response.Success(c, gin.H{"txHash": tx.Hash().Hex()})
}

// PoolSetSwapRouter 设置DEX路由地址（管理员）
func (h *contractHandler) PoolSetSwapRouter(c *gin.Context) {
	form := &types.PoolSetAddressRequest{}
	if err := c.ShouldBindJSON(form); err != nil {
		response.Error(c, ecode.InvalidParams)
		return
	}

	client, auth, _, cleanup, err := prepareTransactor(c, form.RpcURL, form.PrivateKey)
	if err != nil {
		response.Error(c, ecode.ErrPoolRPCConnect)
		return
	}
	defer cleanup()

	poolContract, err := bindcode.NewPledgePool(common.HexToAddress(form.PoolContractAddress), client)
	if err != nil {
		response.Error(c, ecode.ErrPoolInvalidAddr)
		return
	}

	tx, err := poolContract.SetSwapRouter(auth, common.HexToAddress(form.NewAddress))
	if err != nil {
		logger.Error("PoolSetSwapRouter error", logger.Err(err))
		response.Error(c, ecode.ErrPoolSendTx)
		return
	}
	response.Success(c, gin.H{"txHash": tx.Hash().Hex()})
}

// PoolSetMinAmount 设置最小金额（管理员）
func (h *contractHandler) PoolSetMinAmount(c *gin.Context) {
	form := &types.PoolSetMinAmountRequest{}
	if err := c.ShouldBindJSON(form); err != nil {
		response.Error(c, ecode.InvalidParams)
		return
	}
	minAmount, err := parseBigInt(form.MinAmount)
	if err != nil {
		response.Error(c, ecode.ErrPoolInvalidAmount)
		return
	}

	client, auth, _, cleanup, err := prepareTransactor(c, form.RpcURL, form.PrivateKey)
	if err != nil {
		response.Error(c, ecode.ErrPoolRPCConnect)
		return
	}
	defer cleanup()

	poolContract, err := bindcode.NewPledgePool(common.HexToAddress(form.PoolContractAddress), client)
	if err != nil {
		response.Error(c, ecode.ErrPoolInvalidAddr)
		return
	}

	tx, err := poolContract.SetMinAmount(auth, minAmount)
	if err != nil {
		logger.Error("PoolSetMinAmount error", logger.Err(err))
		response.Error(c, ecode.ErrPoolSendTx)
		return
	}
	response.Success(c, gin.H{"txHash": tx.Hash().Hex()})
}

// PoolSetGlobalPaused 暂停/恢复全局（管理员）
func (h *contractHandler) PoolSetGlobalPaused(c *gin.Context) {
	form := &types.PoolWriteRequest{}
	if err := c.ShouldBindJSON(form); err != nil {
		response.Error(c, ecode.InvalidParams)
		return
	}

	client, auth, _, cleanup, err := prepareTransactor(c, form.RpcURL, form.PrivateKey)
	if err != nil {
		response.Error(c, ecode.ErrPoolRPCConnect)
		return
	}
	defer cleanup()

	poolContract, err := bindcode.NewPledgePool(common.HexToAddress(form.PoolContractAddress), client)
	if err != nil {
		response.Error(c, ecode.ErrPoolInvalidAddr)
		return
	}

	tx, err := poolContract.SetGlobalPaused(auth)
	if err != nil {
		logger.Error("PoolSetGlobalPaused error", logger.Err(err))
		response.Error(c, ecode.ErrPoolSendTx)
		return
	}
	response.Success(c, gin.H{"txHash": tx.Hash().Hex()})
}

// PoolEmergencyWithdrawBorrow 紧急提取借入资产
func (h *contractHandler) PoolEmergencyWithdrawBorrow(c *gin.Context) {
	form := &types.PoolWriteRequest{}
	if err := c.ShouldBindJSON(form); err != nil {
		response.Error(c, ecode.InvalidParams)
		return
	}
	pid, err := parseBigInt(form.PoolID)
	if err != nil {
		response.Error(c, ecode.ErrPoolInvalidPoolID)
		return
	}

	client, auth, _, cleanup, err := prepareTransactor(c, form.RpcURL, form.PrivateKey)
	if err != nil {
		response.Error(c, ecode.ErrPoolRPCConnect)
		return
	}
	defer cleanup()

	poolContract, err := bindcode.NewPledgePool(common.HexToAddress(form.PoolContractAddress), client)
	if err != nil {
		response.Error(c, ecode.ErrPoolInvalidAddr)
		return
	}

	tx, err := poolContract.EmergencyWithdrawBorrow(auth, pid)
	if err != nil {
		logger.Error("PoolEmergencyWithdrawBorrow error", logger.Err(err))
		response.Error(c, ecode.ErrPoolSendTx)
		return
	}
	response.Success(c, gin.H{"txHash": tx.Hash().Hex()})
}

// PoolEmergencyWithdrawLend 紧急提取出借资产
func (h *contractHandler) PoolEmergencyWithdrawLend(c *gin.Context) {
	form := &types.PoolWriteRequest{}
	if err := c.ShouldBindJSON(form); err != nil {
		response.Error(c, ecode.InvalidParams)
		return
	}
	pid, err := parseBigInt(form.PoolID)
	if err != nil {
		response.Error(c, ecode.ErrPoolInvalidPoolID)
		return
	}

	client, auth, _, cleanup, err := prepareTransactor(c, form.RpcURL, form.PrivateKey)
	if err != nil {
		response.Error(c, ecode.ErrPoolRPCConnect)
		return
	}
	defer cleanup()

	poolContract, err := bindcode.NewPledgePool(common.HexToAddress(form.PoolContractAddress), client)
	if err != nil {
		response.Error(c, ecode.ErrPoolInvalidAddr)
		return
	}

	tx, err := poolContract.EmergencyWithdrawLend(auth, pid)
	if err != nil {
		logger.Error("PoolEmergencyWithdrawLend error", logger.Err(err))
		response.Error(c, ecode.ErrPoolSendTx)
		return
	}
	response.Success(c, gin.H{"txHash": tx.Hash().Hex()})
}

// PoolTransferOwnership 转移PledgePool所有权（管理员）
func (h *contractHandler) PoolTransferOwnership(c *gin.Context) {
	form := &types.PoolTransferOwnershipRequest{}
	if err := c.ShouldBindJSON(form); err != nil {
		response.Error(c, ecode.InvalidParams)
		return
	}

	client, auth, _, cleanup, err := prepareTransactor(c, form.RpcURL, form.PrivateKey)
	if err != nil {
		response.Error(c, ecode.ErrPoolRPCConnect)
		return
	}
	defer cleanup()

	poolContract, err := bindcode.NewPledgePool(common.HexToAddress(form.PoolContractAddress), client)
	if err != nil {
		response.Error(c, ecode.ErrPoolInvalidAddr)
		return
	}

	tx, err := poolContract.TransferOwnership(auth, common.HexToAddress(form.NewOwner))
	if err != nil {
		logger.Error("PoolTransferOwnership error", logger.Err(err))
		response.Error(c, ecode.ErrPoolSendTx)
		return
	}
	response.Success(c, gin.H{"txHash": tx.Hash().Hex()})
}

// ==================== PledgePool 只读操作 ====================

// PoolGetState 查询池子状态
func (h *contractHandler) PoolGetState(c *gin.Context) {
	form := &types.PoolReadRequest{}
	if err := c.ShouldBindJSON(form); err != nil {
		response.Error(c, ecode.InvalidParams)
		return
	}
	pid, err := parseBigInt(form.PoolID)
	if err != nil {
		response.Error(c, ecode.ErrPoolInvalidPoolID)
		return
	}

	client, cleanup, err := prepareCaller(form.RpcURL)
	if err != nil {
		response.Error(c, ecode.ErrPoolRPCConnect)
		return
	}
	defer cleanup()

	poolContract, err := bindcode.NewPledgePoolCaller(common.HexToAddress(form.PoolContractAddress), client)
	if err != nil {
		response.Error(c, ecode.ErrPoolInvalidAddr)
		return
	}

	state, err := poolContract.GetPoolState(nil, pid)
	if err != nil {
		logger.Error("PoolGetState error", logger.Err(err))
		response.Error(c, ecode.ErrPoolReadCall)
		return
	}
	response.Success(c, gin.H{"state": state.String()})
}

// PoolGetInfo 查询池子详情
func (h *contractHandler) PoolGetInfo(c *gin.Context) {
	form := &types.PoolReadRequest{}
	if err := c.ShouldBindJSON(form); err != nil {
		response.Error(c, ecode.InvalidParams)
		return
	}
	pid, err := parseBigInt(form.PoolID)
	if err != nil {
		response.Error(c, ecode.ErrPoolInvalidPoolID)
		return
	}

	client, cleanup, err := prepareCaller(form.RpcURL)
	if err != nil {
		response.Error(c, ecode.ErrPoolRPCConnect)
		return
	}
	defer cleanup()

	poolContract, err := bindcode.NewPledgePoolCaller(common.HexToAddress(form.PoolContractAddress), client)
	if err != nil {
		response.Error(c, ecode.ErrPoolInvalidAddr)
		return
	}

	info, err := poolContract.PledgePoolInfoList(nil, pid)
	if err != nil {
		logger.Error("PoolGetInfo error", logger.Err(err))
		response.Error(c, ecode.ErrPoolReadCall)
		return
	}
	response.Success(c, gin.H{
		"settleTime":             info.SettleTime.String(),
		"endTime":                info.EndTime.String(),
		"interestRate":           info.InterestRate.String(),
		"maxSupply":              info.MaxSupply.String(),
		"lendSupply":             info.LendSupply.String(),
		"borrowSupply":           info.BorrowSupply.String(),
		"mortgageRate":           info.MortgageRate.String(),
		"lendToken":              info.LendToken.Hex(),
		"borrowToken":            info.BorrowToken.Hex(),
		"state":                  info.State,
		"lendDebtToken":          info.LendDebtToken.Hex(),
		"borrowDebtToken":        info.BorrowDebtToken.Hex(),
		"autoLiquidateThreshold": info.AutoLiquidateThreshold.String(),
	})
}

// PoolGetData 查询池子清算数据
func (h *contractHandler) PoolGetData(c *gin.Context) {
	form := &types.PoolReadRequest{}
	if err := c.ShouldBindJSON(form); err != nil {
		response.Error(c, ecode.InvalidParams)
		return
	}
	pid, err := parseBigInt(form.PoolID)
	if err != nil {
		response.Error(c, ecode.ErrPoolInvalidPoolID)
		return
	}

	client, cleanup, err := prepareCaller(form.RpcURL)
	if err != nil {
		response.Error(c, ecode.ErrPoolRPCConnect)
		return
	}
	defer cleanup()

	poolContract, err := bindcode.NewPledgePoolCaller(common.HexToAddress(form.PoolContractAddress), client)
	if err != nil {
		response.Error(c, ecode.ErrPoolInvalidAddr)
		return
	}

	data, err := poolContract.PoolDataInfoList(nil, pid)
	if err != nil {
		logger.Error("PoolGetData error", logger.Err(err))
		response.Error(c, ecode.ErrPoolReadCall)
		return
	}
	response.Success(c, gin.H{
		"settleAmountLend":        data.SettleAmountLend.String(),
		"settleAmountBorrow":      data.SettleAmountBorrow.String(),
		"finishAmountLend":        data.FinishAmountLend.String(),
		"finishAmountBorrow":      data.FinishAmountBorrow.String(),
		"liquidationAmountLend":   data.LiquidationAmountLend.String(),
		"liquidationAmountBorrow": data.LiquidationAmountBorrow.String(),
	})
}

// PoolCheckCanSettle 检查池子是否可以结算
func (h *contractHandler) PoolCheckCanSettle(c *gin.Context) {
	form := &types.PoolReadRequest{}
	if err := c.ShouldBindJSON(form); err != nil {
		response.Error(c, ecode.InvalidParams)
		return
	}
	pid, err := parseBigInt(form.PoolID)
	if err != nil {
		response.Error(c, ecode.ErrPoolInvalidPoolID)
		return
	}

	client, cleanup, err := prepareCaller(form.RpcURL)
	if err != nil {
		response.Error(c, ecode.ErrPoolRPCConnect)
		return
	}
	defer cleanup()

	poolContract, err := bindcode.NewPledgePoolCaller(common.HexToAddress(form.PoolContractAddress), client)
	if err != nil {
		response.Error(c, ecode.ErrPoolInvalidAddr)
		return
	}

	result, err := poolContract.CheckCanSettle(nil, pid)
	if err != nil {
		logger.Error("PoolCheckCanSettle error", logger.Err(err))
		response.Error(c, ecode.ErrPoolReadCall)
		return
	}
	response.Success(c, gin.H{"result": result})
}

// PoolCheckCanFinish 检查池子是否可以完成
func (h *contractHandler) PoolCheckCanFinish(c *gin.Context) {
	form := &types.PoolReadRequest{}
	if err := c.ShouldBindJSON(form); err != nil {
		response.Error(c, ecode.InvalidParams)
		return
	}
	pid, err := parseBigInt(form.PoolID)
	if err != nil {
		response.Error(c, ecode.ErrPoolInvalidPoolID)
		return
	}

	client, cleanup, err := prepareCaller(form.RpcURL)
	if err != nil {
		response.Error(c, ecode.ErrPoolRPCConnect)
		return
	}
	defer cleanup()

	poolContract, err := bindcode.NewPledgePoolCaller(common.HexToAddress(form.PoolContractAddress), client)
	if err != nil {
		response.Error(c, ecode.ErrPoolInvalidAddr)
		return
	}

	result, err := poolContract.CheckCanFinish(nil, pid)
	if err != nil {
		logger.Error("PoolCheckCanFinish error", logger.Err(err))
		response.Error(c, ecode.ErrPoolReadCall)
		return
	}
	response.Success(c, gin.H{"result": result})
}

// PoolCheckCanLiquidate 检查池子是否可以清算
func (h *contractHandler) PoolCheckCanLiquidate(c *gin.Context) {
	form := &types.PoolReadRequest{}
	if err := c.ShouldBindJSON(form); err != nil {
		response.Error(c, ecode.InvalidParams)
		return
	}
	pid, err := parseBigInt(form.PoolID)
	if err != nil {
		response.Error(c, ecode.ErrPoolInvalidPoolID)
		return
	}

	client, cleanup, err := prepareCaller(form.RpcURL)
	if err != nil {
		response.Error(c, ecode.ErrPoolRPCConnect)
		return
	}
	defer cleanup()

	poolContract, err := bindcode.NewPledgePoolCaller(common.HexToAddress(form.PoolContractAddress), client)
	if err != nil {
		response.Error(c, ecode.ErrPoolInvalidAddr)
		return
	}

	result, err := poolContract.CheckCanLiquidate(nil, pid)
	if err != nil {
		logger.Error("PoolCheckCanLiquidate error", logger.Err(err))
		response.Error(c, ecode.ErrPoolReadCall)
		return
	}
	response.Success(c, gin.H{"result": result})
}

// PoolGetConfig 查询池子全局配置
// 依次调用合约的多个只读方法获取全局配置信息，包括：
// 预言机地址、手续费地址、DEX路由地址、出借费率、借入费率、最小金额、合约所有者
func (h *contractHandler) PoolGetConfig(c *gin.Context) {
	form := &types.PoolReadRequest{}
	if err := c.ShouldBindJSON(form); err != nil {
		response.Error(c, ecode.InvalidParams)
		return
	}

	client, cleanup, err := prepareCaller(form.RpcURL)
	if err != nil {
		response.Error(c, ecode.ErrPoolRPCConnect)
		return
	}
	defer cleanup()

	poolContract, err := bindcode.NewPledgePool(common.HexToAddress(form.PoolContractAddress), client)
	if err != nil {
		response.Error(c, ecode.ErrPoolInvalidAddr)
		return
	}

	// 查询预言机地址
	oracle, err := poolContract.Oracle(nil)
	if err != nil {
		logger.Error("PoolGetConfig Oracle error", logger.Err(err))
		response.Error(c, ecode.ErrPoolReadCall)
		return
	}
	// 查询手续费接收地址
	feeAddr, err := poolContract.FeeAddress(nil)
	if err != nil {
		logger.Error("PoolGetConfig FeeAddress error", logger.Err(err))
		response.Error(c, ecode.ErrPoolReadCall)
		return
	}
	// 查询 DEX 路由合约地址
	swapRouter, err := poolContract.SwapRouter(nil)
	if err != nil {
		logger.Error("PoolGetConfig SwapRouter error", logger.Err(err))
		response.Error(c, ecode.ErrPoolReadCall)
		return
	}
	// 查询出借手续费率
	lendFee, err := poolContract.LendFee(nil)
	if err != nil {
		response.Error(c, ecode.ErrPoolReadCall)
		return
	}
	// 查询借入手续费率
	borrowFee, err := poolContract.BorrowFee(nil)
	if err != nil {
		response.Error(c, ecode.ErrPoolReadCall)
		return
	}
	// 查询最小操作金额
	minAmount, err := poolContract.MinAmount(nil)
	if err != nil {
		response.Error(c, ecode.ErrPoolReadCall)
		return
	}
	// 查询合约所有者地址
	owner, err := poolContract.Owner(nil)
	if err != nil {
		response.Error(c, ecode.ErrPoolReadCall)
		return
	}

	response.Success(c, gin.H{
		"oracle":     oracle.Hex(),
		"feeAddress": feeAddr.Hex(),
		"swapRouter": swapRouter.Hex(),
		"lendFee":    lendFee.String(),
		"borrowFee":  borrowFee.String(),
		"minAmount":  minAmount.String(),
		"owner":      owner.Hex(),
	})
}
