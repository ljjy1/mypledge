package handler

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"

	"github.com/go-dev-frame/sponge/pkg/copier"
	"github.com/go-dev-frame/sponge/pkg/gin/middleware"
	"github.com/go-dev-frame/sponge/pkg/gin/response"
	"github.com/go-dev-frame/sponge/pkg/logger"
	"github.com/go-dev-frame/sponge/pkg/utils"

	"pledge-be/internal/cache"
	"pledge-be/internal/contract/bindcode"
	"pledge-be/internal/dao"
	"pledge-be/internal/database"
	"pledge-be/internal/ecode"
	"pledge-be/internal/model"
	"pledge-be/internal/types"
)

var _ ContractHandler = (*contractHandler)(nil)

// ContractHandler 合约操作处理器接口，包含合约部署、借贷池操作、预言机、债务代币、ERC20、WETH、Uniswap 等全部功能
type ContractHandler interface {
	Create(c *gin.Context)
	DeleteByID(c *gin.Context)
	UpdateByID(c *gin.Context)
	GetByID(c *gin.Context)
	List(c *gin.Context)
	Deploy(c *gin.Context)
	CreatePool(c *gin.Context)

	// PledgePool operations
	PoolLend(c *gin.Context)
	PoolBorrow(c *gin.Context)
	PoolSettle(c *gin.Context)
	PoolFinish(c *gin.Context)
	PoolLiquidate(c *gin.Context)
	PoolRefundBorrow(c *gin.Context)
	PoolRefundLend(c *gin.Context)
	PoolClaimBorrow(c *gin.Context)
	PoolClaimLendDebtToken(c *gin.Context)
	PoolDestroyBorrowDebtToken(c *gin.Context)
	PoolDestroyLendDebtToken(c *gin.Context)
	PoolSetFee(c *gin.Context)
	PoolSetFeeAddress(c *gin.Context)
	PoolSetOracle(c *gin.Context)
	PoolSetSwapRouter(c *gin.Context)
	PoolSetMinAmount(c *gin.Context)
	PoolSetGlobalPaused(c *gin.Context)
	PoolEmergencyWithdrawBorrow(c *gin.Context)
	PoolEmergencyWithdrawLend(c *gin.Context)
	PoolTransferOwnership(c *gin.Context)
	PoolGetState(c *gin.Context)
	PoolGetInfo(c *gin.Context)
	PoolGetData(c *gin.Context)
	PoolCheckCanSettle(c *gin.Context)
	PoolCheckCanFinish(c *gin.Context)
	PoolCheckCanLiquidate(c *gin.Context)
	PoolGetConfig(c *gin.Context)

	// BscPledgeOracle operations
	OracleSetPrice(c *gin.Context)
	OracleSetPrices(c *gin.Context)
	OracleSetAggregator(c *gin.Context)
	OracleTransferOwnership(c *gin.Context)
	OracleGetPrice(c *gin.Context)

	// DebtToken operations
	DebtTokenMint(c *gin.Context)
	DebtTokenBurn(c *gin.Context)
	DebtTokenSetMinter(c *gin.Context)
	DebtTokenBalanceOf(c *gin.Context)
	DebtTokenTotalSupply(c *gin.Context)

	// ERC20 operations
	TokenApprove(c *gin.Context)
	TokenTransfer(c *gin.Context)
	TokenBalanceOf(c *gin.Context)
	TokenAllowance(c *gin.Context)

	// WETH operations
	WETHDeposit(c *gin.Context)
	WETHWithdraw(c *gin.Context)
	WETHBalanceOf(c *gin.Context)

	// UniswapV2Factory operations
	FactoryCreatePair(c *gin.Context)
	FactorySetFeeTo(c *gin.Context)
	FactoryGetPair(c *gin.Context)

	// UniswapV2Router02 operations
	RouterAddLiquidity(c *gin.Context)
	RouterSwapExactTokensForTokens(c *gin.Context)
	RouterGetAmountsOut(c *gin.Context)
}

// contractHandler 合约操作处理器实现，聚合了合约、代币、借贷池基础信息和借贷池运行数据的 DAO
type contractHandler struct {
	iDao         dao.ContractDao
	tokenDao     dao.TokenInfoDao
	poolbasesDao dao.PoolbasesDao
	pooldataDao  dao.PooldataDao
}

// NewContractHandler 创建合约操作处理器实例，初始化合约、代币、借贷池基础信息和运行数据的 DAO
func NewContractHandler() ContractHandler {
	return &contractHandler{
		iDao: dao.NewContractDao(
			database.GetDB(), // db driver is mysql
			cache.NewContractCache(database.GetCacheType()),
		),
		tokenDao: dao.NewTokenInfoDao(
			database.GetDB(),
			nil, // deploy 场景无需缓存
		),
		poolbasesDao: dao.NewPoolbasesDao(
			database.GetDB(),
			cache.NewPoolbasesCache(database.GetCacheType()),
		),
		pooldataDao: dao.NewPooldataDao(
			database.GetDB(),
			cache.NewPooldataCache(database.GetCacheType()),
		),
	}
}

// Create 创建新的合约记录
// @Summary Create a new contract
// @Description Creates a new contract entity using the provided data in the request body.
// @Tags contract
// @Accept json
// @Produce json
// @Param data body types.CreateContractRequest true "contract information"
// @Success 200 {object} types.CreateContractReply{}
// @Router /api/v1/contract [post]
// @Security BearerAuth
func (h *contractHandler) Create(c *gin.Context) {
	form := &types.CreateContractRequest{}
	err := c.ShouldBindJSON(form)
	if err != nil {
		logger.Warn("ShouldBindJSON error: ", logger.Err(err), middleware.GCtxRequestIDField(c))
		response.Error(c, ecode.InvalidParams)
		return
	}

	contract := &model.Contract{}
	err = copier.Copy(contract, form)
	if err != nil {
		response.Error(c, ecode.ErrCreateContract)
		return
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here

	ctx := middleware.WrapCtx(c)
	err = h.iDao.Create(ctx, contract)
	if err != nil {
		logger.Error("Create error", logger.Err(err), logger.Any("form", form), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())
		return
	}

	response.Success(c, gin.H{"id": contract.ID})
}

// DeleteByID 根据 ID 删除合约记录
// @Summary Delete a contract by id
// @Description Deletes a existing contract identified by the given id in the path.
// @Tags contract
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} types.DeleteContractByIDReply{}
// @Router /api/v1/contract/{id} [delete]
// @Security BearerAuth
func (h *contractHandler) DeleteByID(c *gin.Context) {
	_, id, isAbort := getContractIDFromPath(c)
	if isAbort {
		response.Error(c, ecode.InvalidParams)
		return
	}

	ctx := middleware.WrapCtx(c)
	err := h.iDao.DeleteByID(ctx, id)
	if err != nil {
		logger.Error("DeleteByID error", logger.Err(err), logger.Any("id", id), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())
		return
	}

	response.Success(c)
}

// UpdateByID 根据 ID 更新合约记录
// @Summary Update a contract by id
// @Description Updates the specified contract by given id in the path, support partial update.
// @Tags contract
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param data body types.UpdateContractByIDRequest true "contract information"
// @Success 200 {object} types.UpdateContractByIDReply{}
// @Router /api/v1/contract/{id} [put]
// @Security BearerAuth
func (h *contractHandler) UpdateByID(c *gin.Context) {
	_, id, isAbort := getContractIDFromPath(c)
	if isAbort {
		response.Error(c, ecode.InvalidParams)
		return
	}

	form := &types.UpdateContractByIDRequest{}
	err := c.ShouldBindJSON(form)
	if err != nil {
		logger.Warn("ShouldBindJSON error: ", logger.Err(err), middleware.GCtxRequestIDField(c))
		response.Error(c, ecode.InvalidParams)
		return
	}
	form.ID = id

	contract := &model.Contract{}
	err = copier.Copy(contract, form)
	if err != nil {
		response.Error(c, ecode.ErrUpdateByIDContract)
		return
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here

	ctx := middleware.WrapCtx(c)
	err = h.iDao.UpdateByID(ctx, contract)
	if err != nil {
		logger.Error("UpdateByID error", logger.Err(err), logger.Any("form", form), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())
		return
	}

	response.Success(c)
}

// GetByID 根据 ID 查询合约记录
// @Summary Get a contract by id
// @Description Gets detailed information of a contract specified by the given id in the path.
// @Tags contract
// @Param id path string true "id"
// @Accept json
// @Produce json
// @Success 200 {object} types.GetContractByIDReply{}
// @Router /api/v1/contract/{id} [get]
// @Security BearerAuth
func (h *contractHandler) GetByID(c *gin.Context) {
	_, id, isAbort := getContractIDFromPath(c)
	if isAbort {
		response.Error(c, ecode.InvalidParams)
		return
	}

	ctx := middleware.WrapCtx(c)
	contract, err := h.iDao.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, database.ErrRecordNotFound) {
			logger.Warn("GetByID not found", logger.Err(err), logger.Any("id", id), middleware.GCtxRequestIDField(c))
			response.Error(c, ecode.NotFound)
		} else {
			logger.Error("GetByID error", logger.Err(err), logger.Any("id", id), middleware.GCtxRequestIDField(c))
			response.Output(c, ecode.InternalServerError.ToHTTPCode())
		}
		return
	}

	data := &types.ContractObjDetail{}
	err = copier.Copy(data, contract)
	if err != nil {
		response.Error(c, ecode.ErrGetByIDContract)
		return
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here

	response.Success(c, gin.H{"contract": data})
}

// List 分页查询合约列表，支持自定义筛选条件
// @Summary Get a paginated list of contracts by custom conditions
// @Description Returns a paginated list of contract based on query filters, including page number and size.
// @Tags contract
// @Accept json
// @Produce json
// @Param data body types.Params true "query parameters"
// @Success 200 {object} types.ListContractsReply{}
// @Router /api/v1/contract/list [post]
// @Security BearerAuth
func (h *contractHandler) List(c *gin.Context) {
	form := &types.ListContractsRequest{}
	err := c.ShouldBindJSON(form)
	if err != nil {
		logger.Warn("ShouldBindJSON error: ", logger.Err(err), middleware.GCtxRequestIDField(c))
		response.Error(c, ecode.InvalidParams)
		return
	}

	ctx := middleware.WrapCtx(c)
	contracts, total, err := h.iDao.GetByColumns(ctx, &form.Params)
	if err != nil {
		logger.Error("GetByColumns error", logger.Err(err), logger.Any("form", form), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())
		return
	}

	data, err := convertContracts(contracts)
	if err != nil {
		response.Error(c, ecode.ErrListContract)
		return
	}

	response.Success(c, gin.H{
		"contracts": data,
		"total":     total,
	})
}

// Deploy 部署整个借贷合约套件
// @Summary 部署所有合约
// @Description 将完整的借贷合约套件部署到指定的区块链网络。
// @Tags contract
// @Accept json
// @Produce json
// @Param data body types.DeployRequest true "部署参数"
// @Success 200 {object} types.DeployReply{}
// @Router /api/v1/contract/deploy [post]
// @Security BearerAuth
func (h *contractHandler) Deploy(c *gin.Context) {
	form := &types.DeployRequest{}
	if err := c.ShouldBindJSON(form); err != nil {
		response.Error(c, ecode.InvalidParams)
		return
	}

	// 1. 连接 RPC 节点
	client, err := ethclient.Dial(form.RpcURL)
	if err != nil {
		logger.Error("ethclient.Dial error", logger.Err(err))
		response.Error(c, ecode.ErrDeployRPCConn)
		return
	}
	defer client.Close()

	// 2. 从私钥创建钱包对象
	privateKey, err := crypto.HexToECDSA(strings.TrimPrefix(form.PrivateKey, "0x"))
	if err != nil {
		logger.Error("crypto.HexToECDSA error", logger.Err(err))
		response.Error(c, ecode.ErrDeployTxSign)
		return
	}
	fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)

	// 3. 获取 nonce 和 gas 价格
	nonce, err := client.PendingNonceAt(c, fromAddress)
	if err != nil {
		logger.Error("PendingNonceAt error", logger.Err(err))
		response.Error(c, ecode.ErrDeployRPCConn)
		return
	}
	gasPrice, err := client.SuggestGasPrice(c)
	if err != nil {
		logger.Error("SuggestGasPrice error", logger.Err(err))
		response.Error(c, ecode.ErrDeployRPCConn)
		return
	}
	// 从 RPC 节点自动获取链 ID
	chainID, err := client.ChainID(c)
	if err != nil {
		logger.Error("client.ChainID error", logger.Err(err))
		response.Error(c, ecode.ErrDeployRPCConn)
		return
	}

	// 5. 创建交易签名器 (Transactor)
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		logger.Error("NewKeyedTransactorWithChainID error", logger.Err(err))
		response.Error(c, ecode.ErrDeployTxSign)
		return
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasPrice = gasPrice
	auth.GasLimit = uint64(8000000)

	// 默认地址回退到部署者地址
	oracleOwner := defaultAddr(form.OracleOwner, fromAddress)
	poolOwner := defaultAddr(form.PoolOwner, fromAddress)
	feeAddress := defaultAddr(form.FeeAddress, fromAddress)
	factoryFeeTo := defaultAddr(form.FactoryFeeTo, fromAddress)

	// 校验地址不能为空
	for name, addr := range map[string]common.Address{
		"oracleOwner":  oracleOwner,
		"poolOwner":    poolOwner,
		"feeAddress":   feeAddress,
		"factoryFeeTo": factoryFeeTo,
	} {
		if addr == (common.Address{}) {
			logger.Warn("invalid address", logger.String("field", name))
			response.Error(c, ecode.ErrDeployInvalidAddr)
			return
		}
	}

	ctx := middleware.WrapCtx(c)
	result := types.DeployData{
		ChainName: form.ChainName,
		RpcURL:    form.RpcURL,
		Deployer:  fromAddress.Hex(),
		Contracts: []types.DeployContractItem{},
	}

	// 5. 部署 BscPledgeOracle（价格预言机）
	logger.Info("deploying BscPledgeOracle...")
	oracleAddr, oracleTx, _, err := bindcode.DeployBscPledgeOracle(auth, client, oracleOwner)
	if err != nil {
		logger.Error("DeployBscPledgeOracle error", logger.Err(err))
		response.Error(c, ecode.ErrDeploySend)
		return
	}
	_, err = waitReceipt(c, client, oracleTx)
	if err != nil {
		logger.Error("BscPledgeOracle waitReceipt error", logger.Err(err))
		response.Error(c, ecode.ErrDeploySend)
		return
	}
	err = saveDeploy(ctx, h, "BscPledgeOracle", oracleAddr.Hex(), oracleTx.Hash().Hex(), chainID, form, "", 0)
	if err != nil {
		return
	}
	result.Contracts = append(result.Contracts, types.DeployContractItem{
		Name: "BscPledgeOracle", Address: oracleAddr.Hex(), TxHash: oracleTx.Hash().Hex(), Status: "success",
	})
	nonce, err = client.PendingNonceAt(c, fromAddress)
	if err != nil {
		response.Error(c, ecode.ErrDeployRPCConn)
		return
	}
	auth.Nonce = big.NewInt(int64(nonce))

	// 6. 部署 DebtToken（借出债务代币）
	debtTokenName := ifEmpty(form.DebtTokenName, "Lend Debt Token")
	debtTokenSym := ifEmpty(form.DebtTokenSym, "LDT")
	logger.Info("deploying DebtToken (lend)...")
	lendDebtAddr, lendDebtTx, _, err := bindcode.DeployDebtToken(auth, client, debtTokenName, debtTokenSym, poolOwner)
	if err != nil {
		logger.Error("DeployDebtToken(lend) error", logger.Err(err))
		response.Error(c, ecode.ErrDeploySend)
		return
	}
	_, err = waitReceipt(c, client, lendDebtTx)
	if err != nil {
		logger.Error("DebtToken(lend) waitReceipt error", logger.Err(err))
		response.Error(c, ecode.ErrDeploySend)
		return
	}
	err = saveDeploy(ctx, h, "DebtToken(lend)", lendDebtAddr.Hex(), lendDebtTx.Hash().Hex(), chainID, form, debtTokenSym, 18)
	if err != nil {
		return
	}
	result.Contracts = append(result.Contracts, types.DeployContractItem{
		Name: "LendDebtToken", Address: lendDebtAddr.Hex(), TxHash: lendDebtTx.Hash().Hex(), Status: "success",
	})
	nonce, err = client.PendingNonceAt(c, fromAddress)
	if err != nil {
		response.Error(c, ecode.ErrDeployRPCConn)
		return
	}
	auth.Nonce = big.NewInt(int64(nonce))

	// 7. 部署 DebtToken（借入债务代币）
	debtTokenName2 := ifEmpty(form.DebtTokenName+" Borrow", "Borrow Debt Token")
	debtTokenSym2 := ifEmpty(form.DebtTokenSym, "BDT")
	logger.Info("deploying DebtToken (borrow)...")
	borrowDebtAddr, borrowDebtTx, _, err := bindcode.DeployDebtToken(auth, client, debtTokenName2, debtTokenSym2, poolOwner)
	if err != nil {
		logger.Error("DeployDebtToken(borrow) error", logger.Err(err))
		response.Error(c, ecode.ErrDeploySend)
		return
	}
	_, err = waitReceipt(c, client, borrowDebtTx)
	if err != nil {
		logger.Error("DebtToken(borrow) waitReceipt error", logger.Err(err))
		response.Error(c, ecode.ErrDeploySend)
		return
	}
	err = saveDeploy(ctx, h, "DebtToken(borrow)", borrowDebtAddr.Hex(), borrowDebtTx.Hash().Hex(), chainID, form, debtTokenSym2, 18)
	if err != nil {
		return
	}
	result.Contracts = append(result.Contracts, types.DeployContractItem{
		Name: "BorrowDebtToken", Address: borrowDebtAddr.Hex(), TxHash: borrowDebtTx.Hash().Hex(), Status: "success",
	})
	nonce, err = client.PendingNonceAt(c, fromAddress)
	if err != nil {
		response.Error(c, ecode.ErrDeployRPCConn)
		return
	}
	auth.Nonce = big.NewInt(int64(nonce))

	// 8. 部署 MockTestERC20（借出代币，测试用）
	lendName := ifEmpty(form.LendTokenName, "Lend Token")
	lendSym := ifEmpty(form.LendTokenSym, "LEND")
	supply := new(big.Int)
	supply.SetString(ifEmpty(form.TokenSupply, "1000000000000000000000000"), 10)
	logger.Info("deploying MockTestERC20 (lend)...")
	lendTokenAddr, lendTokenTx, _, err := bindcode.DeployMockTestERC20(auth, client, lendName, lendSym, supply)
	if err != nil {
		logger.Error("DeployMockTestERC20(lend) error", logger.Err(err))
		response.Error(c, ecode.ErrDeploySend)
		return
	}
	_, err = waitReceipt(c, client, lendTokenTx)
	if err != nil {
		logger.Error("MockTestERC20(lend) waitReceipt error", logger.Err(err))
		response.Error(c, ecode.ErrDeploySend)
		return
	}
	err = saveDeploy(ctx, h, "MockTestERC20(lend)", lendTokenAddr.Hex(), lendTokenTx.Hash().Hex(), chainID, form, lendSym, 18)
	if err != nil {
		return
	}
	result.Contracts = append(result.Contracts, types.DeployContractItem{
		Name: "MockLendToken", Address: lendTokenAddr.Hex(), TxHash: lendTokenTx.Hash().Hex(), Status: "success",
	})
	nonce, err = client.PendingNonceAt(c, fromAddress)
	if err != nil {
		response.Error(c, ecode.ErrDeployRPCConn)
		return
	}
	auth.Nonce = big.NewInt(int64(nonce))

	// 9. 部署 MockTestERC20（借入代币，测试用）
	borrowName := ifEmpty(form.BorrowTokenName, "Borrow Token")
	borrowSym := ifEmpty(form.BorrowTokenSym, "BRW")
	logger.Info("deploying MockTestERC20 (borrow)...")
	borrowTokenAddr, borrowTokenTx, _, err := bindcode.DeployMockTestERC20(auth, client, borrowName, borrowSym, supply)
	if err != nil {
		logger.Error("DeployMockTestERC20(borrow) error", logger.Err(err))
		response.Error(c, ecode.ErrDeploySend)
		return
	}
	_, err = waitReceipt(c, client, borrowTokenTx)
	if err != nil {
		logger.Error("MockTestERC20(borrow) waitReceipt error", logger.Err(err))
		response.Error(c, ecode.ErrDeploySend)
		return
	}
	err = saveDeploy(ctx, h, "MockTestERC20(borrow)", borrowTokenAddr.Hex(), borrowTokenTx.Hash().Hex(), chainID, form, borrowSym, 18)
	if err != nil {
		return
	}
	result.Contracts = append(result.Contracts, types.DeployContractItem{
		Name: "MockBorrowToken", Address: borrowTokenAddr.Hex(), TxHash: borrowTokenTx.Hash().Hex(), Status: "success",
	})
	nonce, err = client.PendingNonceAt(c, fromAddress)
	if err != nil {
		response.Error(c, ecode.ErrDeployRPCConn)
		return
	}
	auth.Nonce = big.NewInt(int64(nonce))

	// 10. 部署 WETH（包装 ETH）
	logger.Info("deploying WETH...")
	wethAddr, wethTx, _, err := bindcode.DeployWETH(auth, client)
	if err != nil {
		logger.Error("DeployWETH error", logger.Err(err))
		response.Error(c, ecode.ErrDeploySend)
		return
	}
	_, err = waitReceipt(c, client, wethTx)
	if err != nil {
		logger.Error("WETH waitReceipt error", logger.Err(err))
		response.Error(c, ecode.ErrDeploySend)
		return
	}
	err = saveDeploy(ctx, h, "WETH", wethAddr.Hex(), wethTx.Hash().Hex(), chainID, form, "WETH", 18)
	if err != nil {
		return
	}
	result.Contracts = append(result.Contracts, types.DeployContractItem{
		Name: "WETH", Address: wethAddr.Hex(), TxHash: wethTx.Hash().Hex(), Status: "success",
	})
	nonce, err = client.PendingNonceAt(c, fromAddress)
	if err != nil {
		response.Error(c, ecode.ErrDeployRPCConn)
		return
	}
	auth.Nonce = big.NewInt(int64(nonce))

	// 11. 部署 UniswapV2Factory（交易对工厂）
	logger.Info("deploying UniswapV2Factory...")
	factoryAddr, factoryTx, _, err := bindcode.DeployUniswapV2Factory(auth, client, factoryFeeTo)
	if err != nil {
		logger.Error("DeployUniswapV2Factory error", logger.Err(err))
		response.Error(c, ecode.ErrDeploySend)
		return
	}
	_, err = waitReceipt(c, client, factoryTx)
	if err != nil {
		logger.Error("UniswapV2Factory waitReceipt error", logger.Err(err))
		response.Error(c, ecode.ErrDeploySend)
		return
	}
	err = saveDeploy(ctx, h, "UniswapV2Factory", factoryAddr.Hex(), factoryTx.Hash().Hex(), chainID, form, "", 0)
	if err != nil {
		return
	}
	result.Contracts = append(result.Contracts, types.DeployContractItem{
		Name: "UniswapV2Factory", Address: factoryAddr.Hex(), TxHash: factoryTx.Hash().Hex(), Status: "success",
	})
	nonce, err = client.PendingNonceAt(c, fromAddress)
	if err != nil {
		response.Error(c, ecode.ErrDeployRPCConn)
		return
	}
	auth.Nonce = big.NewInt(int64(nonce))

	// 12. 部署 UniswapV2Router02（去中心化交易所路由）
	logger.Info("deploying UniswapV2Router02...")
	routerAddr, routerTx, _, err := bindcode.DeployUniswapV2Router02(auth, client, factoryAddr, wethAddr)
	if err != nil {
		logger.Error("DeployUniswapV2Router02 error", logger.Err(err))
		response.Error(c, ecode.ErrDeploySend)
		return
	}
	_, err = waitReceipt(c, client, routerTx)
	if err != nil {
		logger.Error("UniswapV2Router02 waitReceipt error", logger.Err(err))
		response.Error(c, ecode.ErrDeploySend)
		return
	}
	err = saveDeploy(ctx, h, "UniswapV2Router02", routerAddr.Hex(), routerTx.Hash().Hex(), chainID, form, "", 0)
	if err != nil {
		return
	}
	result.Contracts = append(result.Contracts, types.DeployContractItem{
		Name: "UniswapV2Router02", Address: routerAddr.Hex(), TxHash: routerTx.Hash().Hex(), Status: "success",
	})
	nonce, err = client.PendingNonceAt(c, fromAddress)
	if err != nil {
		response.Error(c, ecode.ErrDeployRPCConn)
		return
	}
	auth.Nonce = big.NewInt(int64(nonce))

	// 13. 部署 PledgePool（借贷池主合约）
	logger.Info("deploying PledgePool...")
	poolAddr, poolTx, _, err := bindcode.DeployPledgePool(auth, client, oracleAddr, routerAddr, feeAddress, poolOwner)
	if err != nil {
		logger.Error("DeployPledgePool error", logger.Err(err))
		response.Error(c, ecode.ErrDeploySend)
		return
	}
	_, err = waitReceipt(c, client, poolTx)
	if err != nil {
		logger.Error("PledgePool waitReceipt error", logger.Err(err))
		response.Error(c, ecode.ErrDeploySend)
		return
	}
	err = saveDeploy(ctx, h, "PledgePool", poolAddr.Hex(), poolTx.Hash().Hex(), chainID, form, "", 0)
	if err != nil {
		return
	}
	result.Contracts = append(result.Contracts, types.DeployContractItem{
		Name: "PledgePool", Address: poolAddr.Hex(), TxHash: poolTx.Hash().Hex(), Status: "success",
	})
	result.PledgePoolAddr = poolAddr.Hex()

	logger.Info("all contracts deployed successfully, PledgePool at", logger.String("address", poolAddr.Hex()))
	response.Success(c, result)
}

// CreatePool 在已部署的 PledgePool 合约上创建借贷池
// @Summary 创建借贷池
// @Description 在已部署的 PledgePool 合约上创建新的借贷池，并将池子信息存入数据库。
// @Tags contract
// @Accept json
// @Produce json
// @Param data body types.CreatePoolRequest true "创建池子参数"
// @Success 200 {object} types.CreatePoolReply{}
// @Router /api/v1/contract/create-pool [post]
// @Security BearerAuth
func (h *contractHandler) CreatePool(c *gin.Context) {
	form := &types.CreatePoolRequest{}
	if err := c.ShouldBindJSON(form); err != nil {
		response.Error(c, ecode.InvalidParams)
		return
	}

	// 1. 连接 RPC 节点
	client, err := ethclient.Dial(form.RpcURL)
	if err != nil {
		logger.Error("ethclient.Dial error", logger.Err(err))
		response.Error(c, ecode.ErrCreatePoolSend)
		return
	}
	defer client.Close()

	// 2. 从私钥创建钱包对象
	privateKey, err := crypto.HexToECDSA(strings.TrimPrefix(form.PrivateKey, "0x"))
	if err != nil {
		logger.Error("crypto.HexToECDSA error", logger.Err(err))
		response.Error(c, ecode.ErrCreatePoolSend)
		return
	}
	fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)

	// 3. 获取 nonce、gas price 和 chain ID
	nonce, err := client.PendingNonceAt(c, fromAddress)
	if err != nil {
		logger.Error("PendingNonceAt error", logger.Err(err))
		response.Error(c, ecode.ErrCreatePoolSend)
		return
	}
	gasPrice, err := client.SuggestGasPrice(c)
	if err != nil {
		logger.Error("SuggestGasPrice error", logger.Err(err))
		response.Error(c, ecode.ErrCreatePoolSend)
		return
	}
	chainID, err := client.ChainID(c)
	if err != nil {
		logger.Error("client.ChainID error", logger.Err(err))
		response.Error(c, ecode.ErrCreatePoolSend)
		return
	}

	// 4. 创建交易签名器
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		logger.Error("NewKeyedTransactorWithChainID error", logger.Err(err))
		response.Error(c, ecode.ErrCreatePoolSend)
		return
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasPrice = gasPrice
	auth.GasLimit = uint64(8000000)

	// 5. 实例化 PledgePool 合约（绑定到已部署的地址）
	poolContract, err := bindcode.NewPledgePool(common.HexToAddress(form.PoolContractAddress), client)
	if err != nil {
		logger.Error("NewPledgePool error", logger.Err(err))
		response.Error(c, ecode.ErrCreatePoolSend)
		return
	}

	// 6. 构造 CreatePool 参数
	settleTime := new(big.Int)
	settleTime.SetString(form.SettleTime, 10)
	endTime := new(big.Int)
	endTime.SetString(form.EndTime, 10)
	interestRate := new(big.Int)
	interestRate.SetString(form.InterestRate, 10)
	maxSupply := new(big.Int)
	maxSupply.SetString(form.MaxSupply, 10)
	mortgageRate := new(big.Int)
	mortgageRate.SetString(form.MortgageRate, 10)
	autoLiquidateThreshold := new(big.Int)
	autoLiquidateThreshold.SetString(form.AutoLiquidateThreshold, 10)

	params := bindcode.PledgePoolCreatePoolParams{
		SettleTime:             settleTime,
		EndTime:                endTime,
		InterestRate:           interestRate,
		MaxSupply:              maxSupply,
		MortgageRate:           mortgageRate,
		LendToken:              common.HexToAddress(form.LendToken),
		BorrowToken:            common.HexToAddress(form.BorrowToken),
		LendDebtToken:          common.HexToAddress(form.LendDebtToken),
		BorrowDebtToken:        common.HexToAddress(form.BorrowDebtToken),
		AutoLiquidateThreshold: autoLiquidateThreshold,
	}

	// 7. 调用合约的 createPledgePool
	tx, err := poolContract.CreatePledgePool(auth, params)
	if err != nil {
		logger.Error("CreatePledgePool error", logger.Err(err))
		response.Error(c, ecode.ErrCreatePoolSend)
		return
	}

	// 8. 等待交易确认
	receipt, err := waitReceipt(c, client, tx)
	if err != nil {
		logger.Error("CreatePledgePool waitReceipt error", logger.Err(err))
		response.Error(c, ecode.ErrCreatePoolReceipt)
		return
	}

	// 9. 从 receipt 日志中解析 CreatePledgePool 事件获取 poolID
	var poolID uint64
	for _, log := range receipt.Logs {
		event, err := poolContract.ParseCreatePledgePool(*log)
		if err == nil {
			poolID = event.Pid.Uint64()
			break
		}
	}
	if poolID == 0 {
		logger.Error("ParseCreatePledgePool event not found in receipt logs")
		response.Error(c, ecode.ErrCreatePoolEvent)
		return
	}
	logger.Info("pool created on chain", logger.Uint64("poolID", poolID), logger.String("txHash", tx.Hash().Hex()))

	// 10. 保存 poolbases 记录
	ctx := middleware.WrapCtx(c)
	poolbases := &model.Poolbases{
		PoolID:                 int(poolID),
		SettleTime:             form.SettleTime,
		EndTime:                form.EndTime,
		InterestRate:           form.InterestRate,
		MaxSupply:              form.MaxSupply,
		MortgageRate:           form.MortgageRate,
		LendToken:              form.LendToken,
		BorrowToken:            form.BorrowToken,
		State:                  "1", // 1=进行中
		LendDebtToken:          form.LendDebtToken,
		BorrowDebtToken:        form.BorrowDebtToken,
		AutoLiquidateThreshold: form.AutoLiquidateThreshold,
		ChainID:                form.ChainID,
		LendTokenSymbol:        form.LendTokenSymbol,
		BorrowTokenSymbol:      form.BorrowTokenSymbol,
	}
	if err := h.poolbasesDao.Create(ctx, poolbases); err != nil {
		logger.Error("poolbasesDao.Create error", logger.Err(err))
		response.Error(c, ecode.ErrCreatePoolSaveDB)
		return
	}

	// 11. 保存 pooldata 记录（初始清算数据为零）
	pooldata := &model.Pooldata{
		ChainID: form.ChainID,
		PoolID:  strconv.FormatUint(poolID, 10),
	}
	if err := h.pooldataDao.Create(ctx, pooldata); err != nil {
		logger.Error("pooldataDao.Create error", logger.Err(err))
		response.Error(c, ecode.ErrCreatePoolSaveDB)
		return
	}

	response.Success(c, gin.H{
		"poolID": poolID,
		"txHash": tx.Hash().Hex(),
	})
}

// prepareTransactor 创建已签名的交易器：连接 RPC 节点、加载私钥、获取 nonce/gas price/chain ID，返回 client、auth、cleanup 函数和错误
func prepareTransactor(c *gin.Context, rpcURL, privateKey string) (
	client *ethclient.Client,
	auth *bind.TransactOpts,
	chainID *big.Int,
	cleanup func(),
	err error,
) {
	client, err = ethclient.Dial(rpcURL)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	cleanup = func() { client.Close() }

	privKey, err := crypto.HexToECDSA(strings.TrimPrefix(privateKey, "0x"))
	if err != nil {
		client.Close()
		return nil, nil, nil, nil, err
	}
	fromAddr := crypto.PubkeyToAddress(privKey.PublicKey)

	nonce, err := client.PendingNonceAt(c, fromAddr)
	if err != nil {
		client.Close()
		return nil, nil, nil, nil, err
	}
	gasPrice, err := client.SuggestGasPrice(c)
	if err != nil {
		client.Close()
		return nil, nil, nil, nil, err
	}
	chainID, err = client.ChainID(c)
	if err != nil {
		client.Close()
		return nil, nil, nil, nil, err
	}

	auth, err = bind.NewKeyedTransactorWithChainID(privKey, chainID)
	if err != nil {
		client.Close()
		return nil, nil, nil, nil, err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasPrice = gasPrice
	auth.GasLimit = uint64(8000000)
	return client, auth, chainID, cleanup, nil
}

// prepareCaller 创建只读调用器，连接 RPC 节点并返回 client 和清理函数
func prepareCaller(rpcURL string) (*ethclient.Client, func(), error) {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, nil, err
	}
	return client, func() { client.Close() }, nil
}

// parseBigInt 安全地将十进制数字符串转换为 *big.Int，转换失败时返回错误
func parseBigInt(s string) (*big.Int, error) {
	n := new(big.Int)
	if _, ok := n.SetString(s, 10); !ok {
		return nil, fmt.Errorf("invalid big int: %s", s)
	}
	return n, nil
}

// defaultAddr 如果地址字符串为空则返回默认地址，否则解析为 common.Address
func defaultAddr(addrStr string, fallback common.Address) common.Address {
	if addrStr == "" {
		return fallback
	}
	return common.HexToAddress(addrStr)
}

// ifEmpty 如果字符串为空则返回备用值，否则返回原字符串
func ifEmpty(s, fallback string) string {
	if s == "" {
		return fallback
	}
	return s
}

// waitReceipt 等待链上交易确认并检查交易状态，状态为 0（回滚）时返回错误
func waitReceipt(c *gin.Context, client *ethclient.Client, tx *ethtypes.Transaction) (*ethtypes.Receipt, error) {
	receipt, err := bind.WaitMined(c, client, tx)
	if err != nil {
		return nil, err
	}
	if receipt.Status == 0 {
		return nil, errors.New("transaction reverted, status=0")
	}
	logger.Info("tx confirmed", logger.Any("blockNumber", receipt.BlockNumber), logger.String("txHash", tx.Hash().Hex()))
	return receipt, nil
}

// saveDeploy 将已部署的合约记录保存到数据库，如果是代币合约还会同时写入 token_info 表
func saveDeploy(ctx context.Context, h *contractHandler, name, addr, txHash string, chainID *big.Int, form *types.DeployRequest, tokenSymbol string, tokenDecimals int) error {
	privKey, err := crypto.HexToECDSA(strings.TrimPrefix(form.PrivateKey, "0x"))
	if err != nil {
		logger.Error("private key error", logger.Err(err))
		return err
	}
	publisher := crypto.PubkeyToAddress(privKey.PublicKey).Hex()
	record := &model.Contract{
		NodeURL:          form.RpcURL,
		ChainID:          chainID.String(),
		ContractAddress:  addr,
		ContractName:     name,
		TxHash:           txHash,
		PublisherAddress: publisher,
	}
	if err := h.iDao.Create(ctx, record); err != nil {
		logger.Error("save deploy record error", logger.Err(err))
		return err
	}
	// 如果是代币合约，同时写入 token_info
	if tokenSymbol != "" {
		token := &model.TokenInfo{
			Symbol:   tokenSymbol,
			Token:    addr,
			ChainID:  chainID.String(),
			Decimals: tokenDecimals,
		}
		if err := h.tokenDao.Create(ctx, token); err != nil {
			logger.Warn("save token info error", logger.Err(err), logger.String("symbol", tokenSymbol))
		}
	}
	return nil
}

// getContractIDFromPath 从请求路径中解析合约 ID，返回 (id字符串, id数值, 是否应中止请求)
func getContractIDFromPath(c *gin.Context) (string, uint64, bool) {
	idStr := c.Param("id")
	id, err := utils.StrToUint64E(idStr)
	if err != nil || id == 0 {
		logger.Warn("StrToUint64E error: ", logger.String("idStr", idStr), middleware.GCtxRequestIDField(c))
		return "", 0, true
	}

	return idStr, id, false
}

// convertContract 将模型层的 Contract 结构体转换为响应层结构体
func convertContract(contract *model.Contract) (*types.ContractObjDetail, error) {
	data := &types.ContractObjDetail{}
	err := copier.Copy(data, contract)
	if err != nil {
		return nil, err
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here

	return data, nil
}

// convertContracts 将模型层的 Contract 列表批量转换为响应层结构体
func convertContracts(fromValues []*model.Contract) ([]*types.ContractObjDetail, error) {
	toValues := []*types.ContractObjDetail{}
	for _, v := range fromValues {
		data, err := convertContract(v)
		if err != nil {
			return nil, err
		}
		toValues = append(toValues, data)
	}

	return toValues, nil
}
