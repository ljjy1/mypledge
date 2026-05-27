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

// ==================== ERC20 Operations ====================

// TokenApprove 授权代币
func (h *contractHandler) TokenApprove(c *gin.Context) {
	form := &types.TokenApproveRequest{}
	if err := c.ShouldBindJSON(form); err != nil {
		response.Error(c, ecode.InvalidParams)
		return
	}
	amount, err := parseBigInt(form.Amount)
	if err != nil {
		response.Error(c, ecode.ErrTokenSendTx)
		return
	}

	client, auth, _, cleanup, err := prepareTransactor(c, form.RpcURL, form.PrivateKey)
	if err != nil {
		response.Error(c, ecode.ErrTokenRPCConnect)
		return
	}
	defer cleanup()

	token, err := bindcode.NewMockTestERC20(common.HexToAddress(form.TokenAddress), client)
	if err != nil {
		response.Error(c, ecode.ErrPoolInvalidAddr)
		return
	}

	tx, err := token.Approve(auth, common.HexToAddress(form.Spender), amount)
	if err != nil {
		logger.Error("TokenApprove error", logger.Err(err))
		response.Error(c, ecode.ErrTokenSendTx)
		return
	}
	response.Success(c, gin.H{"txHash": tx.Hash().Hex()})
}

// TokenTransfer 转账代币
func (h *contractHandler) TokenTransfer(c *gin.Context) {
	form := &types.TokenTransferRequest{}
	if err := c.ShouldBindJSON(form); err != nil {
		response.Error(c, ecode.InvalidParams)
		return
	}
	amount, err := parseBigInt(form.Amount)
	if err != nil {
		response.Error(c, ecode.ErrTokenSendTx)
		return
	}

	client, auth, _, cleanup, err := prepareTransactor(c, form.RpcURL, form.PrivateKey)
	if err != nil {
		response.Error(c, ecode.ErrTokenRPCConnect)
		return
	}
	defer cleanup()

	token, err := bindcode.NewMockTestERC20(common.HexToAddress(form.TokenAddress), client)
	if err != nil {
		response.Error(c, ecode.ErrPoolInvalidAddr)
		return
	}

	tx, err := token.Transfer(auth, common.HexToAddress(form.To), amount)
	if err != nil {
		logger.Error("TokenTransfer error", logger.Err(err))
		response.Error(c, ecode.ErrTokenSendTx)
		return
	}
	response.Success(c, gin.H{"txHash": tx.Hash().Hex()})
}

// TokenBalanceOf 查询代币余额
func (h *contractHandler) TokenBalanceOf(c *gin.Context) {
	form := &types.TokenBalanceOfRequest{}
	if err := c.ShouldBindJSON(form); err != nil {
		response.Error(c, ecode.InvalidParams)
		return
	}

	client, cleanup, err := prepareCaller(form.RpcURL)
	if err != nil {
		response.Error(c, ecode.ErrTokenRPCConnect)
		return
	}
	defer cleanup()

	token, err := bindcode.NewMockTestERC20Caller(common.HexToAddress(form.TokenAddress), client)
	if err != nil {
		response.Error(c, ecode.ErrPoolInvalidAddr)
		return
	}

	balance, err := token.BalanceOf(nil, common.HexToAddress(form.Account))
	if err != nil {
		logger.Error("TokenBalanceOf error", logger.Err(err))
		response.Error(c, ecode.ErrTokenReadCall)
		return
	}
	response.Success(c, gin.H{"balance": balance.String()})
}

// TokenAllowance 查询授权额度
func (h *contractHandler) TokenAllowance(c *gin.Context) {
	form := &types.TokenAllowanceRequest{}
	if err := c.ShouldBindJSON(form); err != nil {
		response.Error(c, ecode.InvalidParams)
		return
	}

	client, cleanup, err := prepareCaller(form.RpcURL)
	if err != nil {
		response.Error(c, ecode.ErrTokenRPCConnect)
		return
	}
	defer cleanup()

	token, err := bindcode.NewMockTestERC20Caller(common.HexToAddress(form.TokenAddress), client)
	if err != nil {
		response.Error(c, ecode.ErrPoolInvalidAddr)
		return
	}

	allowance, err := token.Allowance(nil, common.HexToAddress(form.Owner), common.HexToAddress(form.Spender))
	if err != nil {
		logger.Error("TokenAllowance error", logger.Err(err))
		response.Error(c, ecode.ErrTokenReadCall)
		return
	}
	response.Success(c, gin.H{"allowance": allowance.String()})
}

// ==================== DebtToken Operations ====================

// DebtTokenMint 铸造债务代币（管理员）
func (h *contractHandler) DebtTokenMint(c *gin.Context) {
	form := &types.DebtTokenMintRequest{}
	if err := c.ShouldBindJSON(form); err != nil {
		response.Error(c, ecode.InvalidParams)
		return
	}
	amount, err := parseBigInt(form.Amount)
	if err != nil {
		response.Error(c, ecode.ErrTokenSendTx)
		return
	}

	client, auth, _, cleanup, err := prepareTransactor(c, form.RpcURL, form.PrivateKey)
	if err != nil {
		response.Error(c, ecode.ErrTokenRPCConnect)
		return
	}
	defer cleanup()

	token, err := bindcode.NewDebtToken(common.HexToAddress(form.TokenAddress), client)
	if err != nil {
		response.Error(c, ecode.ErrPoolInvalidAddr)
		return
	}

	tx, err := token.Mint(auth, common.HexToAddress(form.Account), amount)
	if err != nil {
		logger.Error("DebtTokenMint error", logger.Err(err))
		response.Error(c, ecode.ErrTokenSendTx)
		return
	}
	response.Success(c, gin.H{"txHash": tx.Hash().Hex()})
}

// DebtTokenBurn 销毁债务代币（管理员）
func (h *contractHandler) DebtTokenBurn(c *gin.Context) {
	form := &types.DebtTokenBurnRequest{}
	if err := c.ShouldBindJSON(form); err != nil {
		response.Error(c, ecode.InvalidParams)
		return
	}
	amount, err := parseBigInt(form.Amount)
	if err != nil {
		response.Error(c, ecode.ErrTokenSendTx)
		return
	}

	client, auth, _, cleanup, err := prepareTransactor(c, form.RpcURL, form.PrivateKey)
	if err != nil {
		response.Error(c, ecode.ErrTokenRPCConnect)
		return
	}
	defer cleanup()

	token, err := bindcode.NewDebtToken(common.HexToAddress(form.TokenAddress), client)
	if err != nil {
		response.Error(c, ecode.ErrPoolInvalidAddr)
		return
	}

	tx, err := token.Burn(auth, common.HexToAddress(form.Account), amount)
	if err != nil {
		logger.Error("DebtTokenBurn error", logger.Err(err))
		response.Error(c, ecode.ErrTokenSendTx)
		return
	}
	response.Success(c, gin.H{"txHash": tx.Hash().Hex()})
}

// DebtTokenSetMinter 设置铸造者（管理员）
func (h *contractHandler) DebtTokenSetMinter(c *gin.Context) {
	form := &types.DebtTokenSetMinterRequest{}
	if err := c.ShouldBindJSON(form); err != nil {
		response.Error(c, ecode.InvalidParams)
		return
	}

	client, auth, _, cleanup, err := prepareTransactor(c, form.RpcURL, form.PrivateKey)
	if err != nil {
		response.Error(c, ecode.ErrTokenRPCConnect)
		return
	}
	defer cleanup()

	token, err := bindcode.NewDebtToken(common.HexToAddress(form.TokenAddress), client)
	if err != nil {
		response.Error(c, ecode.ErrPoolInvalidAddr)
		return
	}

	tx, err := token.SetMinter(auth, common.HexToAddress(form.Minter), form.Status)
	if err != nil {
		logger.Error("DebtTokenSetMinter error", logger.Err(err))
		response.Error(c, ecode.ErrTokenSendTx)
		return
	}
	response.Success(c, gin.H{"txHash": tx.Hash().Hex()})
}

// DebtTokenBalanceOf 查询债务代币余额
func (h *contractHandler) DebtTokenBalanceOf(c *gin.Context) {
	form := &types.TokenBalanceOfRequest{}
	if err := c.ShouldBindJSON(form); err != nil {
		response.Error(c, ecode.InvalidParams)
		return
	}

	client, cleanup, err := prepareCaller(form.RpcURL)
	if err != nil {
		response.Error(c, ecode.ErrTokenRPCConnect)
		return
	}
	defer cleanup()

	token, err := bindcode.NewDebtTokenCaller(common.HexToAddress(form.TokenAddress), client)
	if err != nil {
		response.Error(c, ecode.ErrPoolInvalidAddr)
		return
	}

	balance, err := token.BalanceOf(nil, common.HexToAddress(form.Account))
	if err != nil {
		logger.Error("DebtTokenBalanceOf error", logger.Err(err))
		response.Error(c, ecode.ErrTokenReadCall)
		return
	}
	response.Success(c, gin.H{"balance": balance.String()})
}

// DebtTokenTotalSupply 查询债务代币总供应量
func (h *contractHandler) DebtTokenTotalSupply(c *gin.Context) {
	form := &types.TokenReadRequest{}
	if err := c.ShouldBindJSON(form); err != nil {
		response.Error(c, ecode.InvalidParams)
		return
	}

	client, cleanup, err := prepareCaller(form.RpcURL)
	if err != nil {
		response.Error(c, ecode.ErrTokenRPCConnect)
		return
	}
	defer cleanup()

	token, err := bindcode.NewDebtTokenCaller(common.HexToAddress(form.TokenAddress), client)
	if err != nil {
		response.Error(c, ecode.ErrPoolInvalidAddr)
		return
	}

	supply, err := token.TotalSupply(nil)
	if err != nil {
		logger.Error("DebtTokenTotalSupply error", logger.Err(err))
		response.Error(c, ecode.ErrTokenReadCall)
		return
	}
	response.Success(c, gin.H{"totalSupply": supply.String()})
}

// ==================== WETH Operations ====================

// WETHDeposit 存入ETH换取WETH
func (h *contractHandler) WETHDeposit(c *gin.Context) {
	form := &types.WETHDepositRequest{}
	if err := c.ShouldBindJSON(form); err != nil {
		response.Error(c, ecode.InvalidParams)
		return
	}
	amount, err := parseBigInt(form.Amount)
	if err != nil {
		response.Error(c, ecode.ErrTokenSendTx)
		return
	}

	client, auth, _, cleanup, err := prepareTransactor(c, form.RpcURL, form.PrivateKey)
	if err != nil {
		response.Error(c, ecode.ErrTokenRPCConnect)
		return
	}
	defer cleanup()

	// Deposit 是 payable 方法，需要设置 Value
	auth.Value = amount

	weth, err := bindcode.NewWETH(common.HexToAddress(form.TokenAddress), client)
	if err != nil {
		response.Error(c, ecode.ErrPoolInvalidAddr)
		return
	}

	tx, err := weth.Deposit(auth)
	if err != nil {
		logger.Error("WETHDeposit error", logger.Err(err))
		response.Error(c, ecode.ErrTokenSendTx)
		return
	}
	response.Success(c, gin.H{"txHash": tx.Hash().Hex()})
}

// WETHWithdraw 提取WETH换回ETH
func (h *contractHandler) WETHWithdraw(c *gin.Context) {
	form := &types.WETHDepositRequest{}
	if err := c.ShouldBindJSON(form); err != nil {
		response.Error(c, ecode.InvalidParams)
		return
	}
	amount, err := parseBigInt(form.Amount)
	if err != nil {
		response.Error(c, ecode.ErrTokenSendTx)
		return
	}

	client, auth, _, cleanup, err := prepareTransactor(c, form.RpcURL, form.PrivateKey)
	if err != nil {
		response.Error(c, ecode.ErrTokenRPCConnect)
		return
	}
	defer cleanup()

	weth, err := bindcode.NewWETH(common.HexToAddress(form.TokenAddress), client)
	if err != nil {
		response.Error(c, ecode.ErrPoolInvalidAddr)
		return
	}

	tx, err := weth.Withdraw(auth, amount)
	if err != nil {
		logger.Error("WETHWithdraw error", logger.Err(err))
		response.Error(c, ecode.ErrTokenSendTx)
		return
	}
	response.Success(c, gin.H{"txHash": tx.Hash().Hex()})
}

// WETHBalanceOf 查询WETH余额
func (h *contractHandler) WETHBalanceOf(c *gin.Context) {
	form := &types.TokenBalanceOfRequest{}
	if err := c.ShouldBindJSON(form); err != nil {
		response.Error(c, ecode.InvalidParams)
		return
	}

	client, cleanup, err := prepareCaller(form.RpcURL)
	if err != nil {
		response.Error(c, ecode.ErrTokenRPCConnect)
		return
	}
	defer cleanup()

	weth, err := bindcode.NewWETHCaller(common.HexToAddress(form.TokenAddress), client)
	if err != nil {
		response.Error(c, ecode.ErrPoolInvalidAddr)
		return
	}

	balance, err := weth.BalanceOf(nil, common.HexToAddress(form.Account))
	if err != nil {
		logger.Error("WETHBalanceOf error", logger.Err(err))
		response.Error(c, ecode.ErrTokenReadCall)
		return
	}
	response.Success(c, gin.H{"balance": balance.String()})
}
