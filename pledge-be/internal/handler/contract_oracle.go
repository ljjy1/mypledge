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

// ==================== BscPledgeOracle 写入操作 ====================

// OracleSetPrice 设置资产价格（管理员）
// @Summary Set asset price in oracle
// @Description Admin sets the price of an asset in the price oracle
// @Tags contract
// @Accept json
// @Produce json
// @Param data body types.OracleSetPriceRequest true "set price params"
// @Success 200 {object} types.PoolWriteReply{}
// @Router /api/v1/contract/oracle/set-price [post]
// @Security BearerAuth
func (h *contractHandler) OracleSetPrice(c *gin.Context) {
	form := &types.OracleSetPriceRequest{}
	if err := c.ShouldBindJSON(form); err != nil {
		response.Error(c, ecode.InvalidParams)
		return
	}
	price, err := parseBigInt(form.Price)
	if err != nil {
		response.Error(c, ecode.ErrOracleInvalidParams)
		return
	}

	client, auth, _, cleanup, err := prepareTransactor(c, form.RpcURL, form.PrivateKey)
	if err != nil {
		response.Error(c, ecode.ErrOracleRPCConnect)
		return
	}
	defer cleanup()

	oracle, err := bindcode.NewBscPledgeOracle(common.HexToAddress(form.OracleAddress), client)
	if err != nil {
		response.Error(c, ecode.ErrPoolInvalidAddr)
		return
	}

	tx, err := oracle.SetPrice(auth, common.HexToAddress(form.AssetAddress), price)
	if err != nil {
		logger.Error("OracleSetPrice error", logger.Err(err))
		response.Error(c, ecode.ErrOracleSendTx)
		return
	}
	response.Success(c, gin.H{"txHash": tx.Hash().Hex()})
}

// OracleSetPrices 批量设置价格（需要 assets, prices 数组参数）
// @Summary Batch set asset prices (not yet implemented)
// @Description Batch set prices for multiple assets - currently not fully implemented
// @Tags contract
// @Accept json
// @Produce json
// @Param data body types.OracleWriteRequest true "set prices params"
// @Success 200 {object} types.PoolWriteReply{}
// @Router /api/v1/contract/oracle/set-prices [post]
// @Security BearerAuth
func (h *contractHandler) OracleSetPrices(c *gin.Context) {
	form := &types.OracleWriteRequest{}
	if err := c.ShouldBindJSON(form); err != nil {
		response.Error(c, ecode.InvalidParams)
		return
	}
	response.Error(c, ecode.ErrOracleInvalidParams)
}

// OracleSetAggregator 设置预言机聚合器（管理员）
// @Summary Set oracle aggregator
// @Description Admin sets the price aggregator for an asset
// @Tags contract
// @Accept json
// @Produce json
// @Param data body types.OracleSetAggregatorRequest true "set aggregator params"
// @Success 200 {object} types.PoolWriteReply{}
// @Router /api/v1/contract/oracle/set-aggregator [post]
// @Security BearerAuth
func (h *contractHandler) OracleSetAggregator(c *gin.Context) {
	form := &types.OracleSetAggregatorRequest{}
	if err := c.ShouldBindJSON(form); err != nil {
		response.Error(c, ecode.InvalidParams)
		return
	}

	client, auth, _, cleanup, err := prepareTransactor(c, form.RpcURL, form.PrivateKey)
	if err != nil {
		response.Error(c, ecode.ErrOracleRPCConnect)
		return
	}
	defer cleanup()

	oracle, err := bindcode.NewBscPledgeOracle(common.HexToAddress(form.OracleAddress), client)
	if err != nil {
		response.Error(c, ecode.ErrPoolInvalidAddr)
		return
	}

	tx, err := oracle.SetAssetsAggregator(auth, common.HexToAddress(form.AssetAddress), common.HexToAddress(form.Aggregator), uint8(form.TokenDecimals))
	if err != nil {
		logger.Error("OracleSetAggregator error", logger.Err(err))
		response.Error(c, ecode.ErrOracleSendTx)
		return
	}
	response.Success(c, gin.H{"txHash": tx.Hash().Hex()})
}

// OracleTransferOwnership 转移Oracle所有权（管理员）
// @Summary Transfer oracle ownership
// @Description Admin transfers ownership of the oracle contract to a new address
// @Tags contract
// @Accept json
// @Produce json
// @Param data body types.OracleWriteRequest true "transfer ownership params"
// @Success 200 {object} types.PoolWriteReply{}
// @Router /api/v1/contract/oracle/transfer-ownership [post]
// @Security BearerAuth
func (h *contractHandler) OracleTransferOwnership(c *gin.Context) {
	form := &types.OracleWriteRequest{}
	if err := c.ShouldBindJSON(form); err != nil {
		response.Error(c, ecode.InvalidParams)
		return
	}

	client, auth, _, cleanup, err := prepareTransactor(c, form.RpcURL, form.PrivateKey)
	if err != nil {
		response.Error(c, ecode.ErrOracleRPCConnect)
		return
	}
	defer cleanup()

	oracle, err := bindcode.NewBscPledgeOracle(common.HexToAddress(form.OracleAddress), client)
	if err != nil {
		response.Error(c, ecode.ErrPoolInvalidAddr)
		return
	}

	tx, err := oracle.TransferOwnership(auth, common.HexToAddress(form.NewAddress))
	if err != nil {
		logger.Error("OracleTransferOwnership error", logger.Err(err))
		response.Error(c, ecode.ErrOracleSendTx)
		return
	}
	response.Success(c, gin.H{"txHash": tx.Hash().Hex()})
}

// ==================== BscPledgeOracle 只读操作 ====================

// OracleGetPrice 查询资产价格
// @Summary Get asset price from oracle
// @Description Query the current price of an asset from the price oracle
// @Tags contract
// @Accept json
// @Produce json
// @Param data body types.OracleGetPriceRequest true "get price params"
// @Success 200 {object} types.OraclePriceReply{}
// @Router /api/v1/contract/oracle/get-price [post]
// @Security BearerAuth
func (h *contractHandler) OracleGetPrice(c *gin.Context) {
	form := &types.OracleGetPriceRequest{}
	if err := c.ShouldBindJSON(form); err != nil {
		response.Error(c, ecode.InvalidParams)
		return
	}

	client, cleanup, err := prepareCaller(form.RpcURL)
	if err != nil {
		response.Error(c, ecode.ErrOracleRPCConnect)
		return
	}
	defer cleanup()

	oracle, err := bindcode.NewBscPledgeOracleCaller(common.HexToAddress(form.OracleAddress), client)
	if err != nil {
		response.Error(c, ecode.ErrPoolInvalidAddr)
		return
	}

	price, err := oracle.GetPrice(nil, common.HexToAddress(form.AssetAddress))
	if err != nil {
		logger.Error("OracleGetPrice error", logger.Err(err))
		response.Error(c, ecode.ErrOracleReadCall)
		return
	}
	response.Success(c, gin.H{"price": price.String()})
}
