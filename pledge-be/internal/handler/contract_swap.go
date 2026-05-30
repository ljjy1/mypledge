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

// ==================== UniswapV2Factory 操作 ====================

// FactoryCreatePair 创建交易对
// @Summary Create a token pair on UniswapV2 factory
// @Description Create a new token trading pair on the UniswapV2 factory contract
// @Tags contract
// @Accept json
// @Produce json
// @Param data body types.FactoryCreatePairRequest true "create pair params"
// @Success 200 {object} types.FactoryCreatePairReply{}
// @Router /api/v1/contract/factory/create-pair [post]
// @Security BearerAuth
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
	// CreatePair 返回交易对地址，但 go 绑定返回类型为 *types.Transaction，没有直接的 address 返回值
	// 需要等待 receipt 然后解析事件；此处简化实现，直接返回交易哈希
	response.Success(c, gin.H{"txHash": tx.Hash().Hex()})
}

// FactorySetFeeTo 设置工厂手续费地址
// @Summary Set factory fee address
// @Description Set the fee address for the UniswapV2 factory contract
// @Tags contract
// @Accept json
// @Produce json
// @Param data body types.FactorySetFeeToRequest true "set fee to params"
// @Success 200 {object} types.PoolWriteReply{}
// @Router /api/v1/contract/factory/set-fee-to [post]
// @Security BearerAuth
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
// @Summary Get token pair address
// @Description Query the trading pair address on UniswapV2 factory
// @Tags contract
// @Accept json
// @Produce json
// @Param data body types.FactoryGetPairRequest true "get pair params"
// @Success 200 {object} types.FactoryPairReply{}
// @Router /api/v1/contract/factory/get-pair [post]
// @Security BearerAuth
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

// ==================== UniswapV2Router02 操作 ====================

// RouterAddLiquidity 添加流动性
// @Summary Add liquidity to a token pair
// @Description Add liquidity to a UniswapV2 trading pair
// @Tags contract
// @Accept json
// @Produce json
// @Param data body types.RouterAddLiquidityRequest true "add liquidity params"
// @Success 200 {object} types.PoolWriteReply{}
// @Router /api/v1/contract/router/add-liquidity [post]
// @Security BearerAuth
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
// @Summary Swap exact tokens for tokens
// @Description Swap an exact amount of input tokens for a minimum amount of output tokens
// @Tags contract
// @Accept json
// @Produce json
// @Param data body types.RouterSwapRequest true "swap params"
// @Success 200 {object} types.PoolWriteReply{}
// @Router /api/v1/contract/router/swap-exact-tokens-for-tokens [post]
// @Security BearerAuth
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
// @Summary Get swap output amounts
// @Description Query expected output amounts for a given input along a swap path
// @Tags contract
// @Accept json
// @Produce json
// @Param data body types.RouterGetAmountsRequest true "get amounts params"
// @Success 200 {object} types.RouterAmountsReply{}
// @Router /api/v1/contract/router/get-amounts-out [post]
// @Security BearerAuth
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
