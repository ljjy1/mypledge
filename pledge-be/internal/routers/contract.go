package routers

import (
	"github.com/gin-gonic/gin"

	"pledge-be/internal/handler"
)

func init() {
	apiV1RouterFns = append(apiV1RouterFns, func(group *gin.RouterGroup) {
		contractRouter(group, handler.NewContractHandler())
	})
}

func contractRouter(group *gin.RouterGroup, h handler.ContractHandler) {
	g := group.Group("/contract")

	// JWT authentication reference: https://go-sponge.com/component/transport/gin.html#jwt-authorization-middleware

	// All the following routes use jwt authentication, you also can use middleware.Auth(middleware.WithExtraVerify(fn))
	//g.Use(middleware.Auth())

	// If jwt authentication is not required for all routes, authentication middleware can be added
	// separately for only certain routes. In this case, g.Use(middleware.Auth()) above should not be used.

	// --- Contract CRUD ---
	g.POST("/", h.Create)                // [post] /api/v1/contract
	g.POST("/deploy", h.Deploy)          // [post] /api/v1/contract/deploy
	g.POST("/create-pool", h.CreatePool) // [post] /api/v1/contract/create-pool
	g.DELETE("/:id", h.DeleteByID)       // [delete] /api/v1/contract/:id
	g.PUT("/:id", h.UpdateByID)          // [put] /api/v1/contract/:id
	g.GET("/:id", h.GetByID)             // [get] /api/v1/contract/:id
	g.POST("/list", h.List)              // [post] /api/v1/contract/list

	// --- PledgePool Write ---
	g.POST("/pool/lend", h.PoolLend)                                         // [post] /api/v1/contract/pool/lend
	g.POST("/pool/borrow", h.PoolBorrow)                                     // [post] /api/v1/contract/pool/borrow
	g.POST("/pool/settle", h.PoolSettle)                                     // [post] /api/v1/contract/pool/settle
	g.POST("/pool/finish", h.PoolFinish)                                     // [post] /api/v1/contract/pool/finish
	g.POST("/pool/liquidate", h.PoolLiquidate)                               // [post] /api/v1/contract/pool/liquidate
	g.POST("/pool/refund-borrow", h.PoolRefundBorrow)                        // [post] /api/v1/contract/pool/refund-borrow
	g.POST("/pool/refund-lend", h.PoolRefundLend)                            // [post] /api/v1/contract/pool/refund-lend
	g.POST("/pool/claim-borrow", h.PoolClaimBorrow)                          // [post] /api/v1/contract/pool/claim-borrow
	g.POST("/pool/claim-lend-debt", h.PoolClaimLendDebtToken)                // [post] /api/v1/contract/pool/claim-lend-debt
	g.POST("/pool/destroy-borrow-debt", h.PoolDestroyBorrowDebtToken)        // [post] /api/v1/contract/pool/destroy-borrow-debt
	g.POST("/pool/destroy-lend-debt", h.PoolDestroyLendDebtToken)            // [post] /api/v1/contract/pool/destroy-lend-debt
	g.POST("/pool/set-fee", h.PoolSetFee)                                    // [post] /api/v1/contract/pool/set-fee
	g.POST("/pool/set-fee-address", h.PoolSetFeeAddress)                     // [post] /api/v1/contract/pool/set-fee-address
	g.POST("/pool/set-oracle", h.PoolSetOracle)                              // [post] /api/v1/contract/pool/set-oracle
	g.POST("/pool/set-swap-router", h.PoolSetSwapRouter)                     // [post] /api/v1/contract/pool/set-swap-router
	g.POST("/pool/set-min-amount", h.PoolSetMinAmount)                       // [post] /api/v1/contract/pool/set-min-amount
	g.POST("/pool/set-global-paused", h.PoolSetGlobalPaused)                 // [post] /api/v1/contract/pool/set-global-paused
	g.POST("/pool/emergency-withdraw-borrow", h.PoolEmergencyWithdrawBorrow) // [post] /api/v1/contract/pool/emergency-withdraw-borrow
	g.POST("/pool/emergency-withdraw-lend", h.PoolEmergencyWithdrawLend)     // [post] /api/v1/contract/pool/emergency-withdraw-lend
	g.POST("/pool/transfer-ownership", h.PoolTransferOwnership)              // [post] /api/v1/contract/pool/transfer-ownership

	// --- PledgePool Read ---
	g.POST("/pool/state", h.PoolGetState)                        // [post] /api/v1/contract/pool/state
	g.POST("/pool/info", h.PoolGetInfo)                          // [post] /api/v1/contract/pool/info
	g.POST("/pool/data", h.PoolGetData)                          // [post] /api/v1/contract/pool/data
	g.POST("/pool/check-can-settle", h.PoolCheckCanSettle)       // [post] /api/v1/contract/pool/check-can-settle
	g.POST("/pool/check-can-finish", h.PoolCheckCanFinish)       // [post] /api/v1/contract/pool/check-can-finish
	g.POST("/pool/check-can-liquidate", h.PoolCheckCanLiquidate) // [post] /api/v1/contract/pool/check-can-liquidate
	g.POST("/pool/config", h.PoolGetConfig)                      // [post] /api/v1/contract/pool/config

	// --- BscPledgeOracle ---
	g.POST("/oracle/set-price", h.OracleSetPrice)                   // [post] /api/v1/contract/oracle/set-price
	g.POST("/oracle/set-prices", h.OracleSetPrices)                 // [post] /api/v1/contract/oracle/set-prices (暂未完整实现)
	g.POST("/oracle/set-aggregator", h.OracleSetAggregator)         // [post] /api/v1/contract/oracle/set-aggregator
	g.POST("/oracle/transfer-ownership", h.OracleTransferOwnership) // [post] /api/v1/contract/oracle/transfer-ownership
	g.POST("/oracle/get-price", h.OracleGetPrice)                   // [post] /api/v1/contract/oracle/get-price

	// --- DebtToken ---
	g.POST("/debt-token/mint", h.DebtTokenMint)                // [post] /api/v1/contract/debt-token/mint
	g.POST("/debt-token/burn", h.DebtTokenBurn)                // [post] /api/v1/contract/debt-token/burn
	g.POST("/debt-token/set-minter", h.DebtTokenSetMinter)     // [post] /api/v1/contract/debt-token/set-minter
	g.POST("/debt-token/balance-of", h.DebtTokenBalanceOf)     // [post] /api/v1/contract/debt-token/balance-of
	g.POST("/debt-token/total-supply", h.DebtTokenTotalSupply) // [post] /api/v1/contract/debt-token/total-supply

	// --- ERC20 ---
	g.POST("/token/approve", h.TokenApprove)      // [post] /api/v1/contract/token/approve
	g.POST("/token/transfer", h.TokenTransfer)    // [post] /api/v1/contract/token/transfer
	g.POST("/token/balance-of", h.TokenBalanceOf) // [post] /api/v1/contract/token/balance-of
	g.POST("/token/allowance", h.TokenAllowance)  // [post] /api/v1/contract/token/allowance

	// --- WETH ---
	g.POST("/weth/deposit", h.WETHDeposit)      // [post] /api/v1/contract/weth/deposit
	g.POST("/weth/withdraw", h.WETHWithdraw)    // [post] /api/v1/contract/weth/withdraw
	g.POST("/weth/balance-of", h.WETHBalanceOf) // [post] /api/v1/contract/weth/balance-of

	// --- UniswapV2Factory ---
	g.POST("/factory/create-pair", h.FactoryCreatePair) // [post] /api/v1/contract/factory/create-pair
	g.POST("/factory/set-fee-to", h.FactorySetFeeTo)    // [post] /api/v1/contract/factory/set-fee-to
	g.POST("/factory/get-pair", h.FactoryGetPair)       // [post] /api/v1/contract/factory/get-pair

	// --- UniswapV2Router02 ---
	g.POST("/router/add-liquidity", h.RouterAddLiquidity)                            // [post] /api/v1/contract/router/add-liquidity
	g.POST("/router/swap-exact-tokens-for-tokens", h.RouterSwapExactTokensForTokens) // [post] /api/v1/contract/router/swap-exact-tokens-for-tokens
	g.POST("/router/get-amounts-out", h.RouterGetAmountsOut)                         // [post] /api/v1/contract/router/get-amounts-out
}
