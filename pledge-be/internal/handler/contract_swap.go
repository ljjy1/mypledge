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

// ==================== UniswapV2Factory Operations ====================

// FactoryCreatePair 创建交易对
func (h *contractHandler) FactoryCreatePair(c *gin.Context) {
	form := &types.FactoryCreatePairRequest{}
	if err := c.ShouldBindJSON(form); err != nil {
		response.Error(c, ecode.InvalidParams)
		return
	}

	client, auth, _, cleanup, err := prepareTransactor(c, form.RpcURL, form.PrivateKey)
	if err != nil {
		response.Error(c, ecode.ErrSwapRPCConnect)
		return
	}
	defer cleanup()

	factory, err := bindcode.NewUniswapV2Factory(common.HexToAddress(form.FactoryAddress), client)
	if err != nil {
		response.Error(c, ecode.ErrPoolInvalidAddr)
		return
	}

	tx, err := factory.CreatePair(auth, common.HexToAddress(form.TokenA), common.HexToAddress(form.TokenB))
	if err != nil {
		logger.Error("FactoryCreatePair error", logger.Err(err))
		response.Error(c, ecode.ErrSwapSendTx)
		return
	}

	_ = tx
	// CreatePair 返回交易对地址，类型为 *types.Transaction，但没有直接的 address 返回值
	// 需要等待 receipt 然后解析事件，简化实现返回 txHash
	response.Success(c, gin.H{"txHash": tx.Hash().Hex()})
}

// FactorySetFeeTo 设置工厂手续费地址
func (h *contractHandler) FactorySetFeeTo(c *gin.Context) {
	form := &types.FactorySetFeeToRequest{}
	if err := c.ShouldBindJSON(form); err != nil {
		response.Error(c, ecode.InvalidParams)
		return
	}

	client, auth, _, cleanup, err := prepareTransactor(c, form.RpcURL, form.PrivateKey)
	if err != nil {
		response.Error(c, ecode.ErrSwapRPCConnect)
		return
	}
	defer cleanup()

	factory, err := bindcode.NewUniswapV2Factory(common.HexToAddress(form.FactoryAddress), client)
	if err != nil {
		response.Error(c, ecode.ErrPoolInvalidAddr)
		return
	}

	tx, err := factory.SetFeeTo(auth, common.HexToAddress(form.FeeTo))
	if err != nil {
		logger.Error("FactorySetFeeTo error", logger.Err(err))
		response.Error(c, ecode.ErrSwapSendTx)
		return
	}
	response.Success(c, gin.H{"txHash": tx.Hash().Hex()})
}

// FactoryGetPair 查询交易对地址
func (h *contractHandler) FactoryGetPair(c *gin.Context) {
	form := &types.FactoryGetPairRequest{}
	if err := c.ShouldBindJSON(form); err != nil {
		response.Error(c, ecode.InvalidParams)
		return
	}

	client, cleanup, err := prepareCaller(form.RpcURL)
	if err != nil {
		response.Error(c, ecode.ErrSwapRPCConnect)
		return
	}
	defer cleanup()

	factory, err := bindcode.NewUniswapV2FactoryCaller(common.HexToAddress(form.FactoryAddress), client)
	if err != nil {
		response.Error(c, ecode.ErrPoolInvalidAddr)
		return
	}

	pairAddr, err := factory.GetPair(nil, common.HexToAddress(form.TokenA), common.HexToAddress(form.TokenB))
	if err != nil {
		logger.Error("FactoryGetPair error", logger.Err(err))
		response.Error(c, ecode.ErrSwapReadCall)
		return
	}
	response.Success(c, gin.H{"pairAddress": pairAddr.Hex()})
}

// ==================== UniswapV2Router02 Operations ====================

// RouterAddLiquidity 添加流动性
func (h *contractHandler) RouterAddLiquidity(c *gin.Context) {
	form := &types.RouterAddLiquidityRequest{}
	if err := c.ShouldBindJSON(form); err != nil {
		response.Error(c, ecode.InvalidParams)
		return
	}

	amountADesired, _ := parseBigInt(form.AmountADesired)
	amountBDesired, _ := parseBigInt(form.AmountBDesired)
	amountAMin, _ := parseBigInt(form.AmountAMin)
	amountBMin, _ := parseBigInt(form.AmountBMin)
	deadline, _ := parseBigInt(form.Deadline)

	client, auth, _, cleanup, err := prepareTransactor(c, form.RpcURL, form.PrivateKey)
	if err != nil {
		response.Error(c, ecode.ErrSwapRPCConnect)
		return
	}
	defer cleanup()

	router, err := bindcode.NewUniswapV2Router02(common.HexToAddress(form.RouterAddress), client)
	if err != nil {
		response.Error(c, ecode.ErrPoolInvalidAddr)
		return
	}

	tx, err := router.AddLiquidity(auth,
		common.HexToAddress(form.TokenA),
		common.HexToAddress(form.TokenB),
		amountADesired, amountBDesired,
		amountAMin, amountBMin,
		common.HexToAddress(form.To),
		deadline,
	)
	if err != nil {
		logger.Error("RouterAddLiquidity error", logger.Err(err))
		response.Error(c, ecode.ErrSwapSendTx)
		return
	}
	response.Success(c, gin.H{"txHash": tx.Hash().Hex()})
}

// RouterSwapExactTokensForTokens 精确兑换(指定输入金额)
func (h *contractHandler) RouterSwapExactTokensForTokens(c *gin.Context) {
	form := &types.RouterSwapRequest{}
	if err := c.ShouldBindJSON(form); err != nil {
		response.Error(c, ecode.InvalidParams)
		return
	}

	amountIn, _ := parseBigInt(form.AmountIn)
	amountOutMin, _ := parseBigInt(form.AmountOutMin)
	deadline, _ := parseBigInt(form.Deadline)

	path := make([]common.Address, len(form.Path))
	for i, p := range form.Path {
		path[i] = common.HexToAddress(p)
	}

	client, auth, _, cleanup, err := prepareTransactor(c, form.RpcURL, form.PrivateKey)
	if err != nil {
		response.Error(c, ecode.ErrSwapRPCConnect)
		return
	}
	defer cleanup()

	router, err := bindcode.NewUniswapV2Router02(common.HexToAddress(form.RouterAddress), client)
	if err != nil {
		response.Error(c, ecode.ErrPoolInvalidAddr)
		return
	}

	tx, err := router.SwapExactTokensForTokens(auth, amountIn, amountOutMin, path, common.HexToAddress(form.To), deadline)
	if err != nil {
		logger.Error("RouterSwapExactTokensForTokens error", logger.Err(err))
		response.Error(c, ecode.ErrSwapSendTx)
		return
	}
	response.Success(c, gin.H{"txHash": tx.Hash().Hex()})
}

// RouterGetAmountsOut 查询兑换输出金额
func (h *contractHandler) RouterGetAmountsOut(c *gin.Context) {
	form := &types.RouterGetAmountsRequest{}
	if err := c.ShouldBindJSON(form); err != nil {
		response.Error(c, ecode.InvalidParams)
		return
	}

	amountIn, _ := parseBigInt(form.AmountIn)
	path := make([]common.Address, len(form.Path))
	for i, p := range form.Path {
		path[i] = common.HexToAddress(p)
	}

	client, cleanup, err := prepareCaller(form.RpcURL)
	if err != nil {
		response.Error(c, ecode.ErrSwapRPCConnect)
		return
	}
	defer cleanup()

	router, err := bindcode.NewUniswapV2Router02Caller(common.HexToAddress(form.RouterAddress), client)
	if err != nil {
		response.Error(c, ecode.ErrPoolInvalidAddr)
		return
	}

	amounts, err := router.GetAmountsOut(nil, amountIn, path)
	if err != nil {
		logger.Error("RouterGetAmountsOut error", logger.Err(err))
		response.Error(c, ecode.ErrSwapReadCall)
		return
	}

	result := make([]string, len(amounts))
	for i, v := range amounts {
		result[i] = v.String()
	}
	response.Success(c, gin.H{"amounts": result})
}
