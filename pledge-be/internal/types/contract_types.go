package types

import (
	"time"

	"github.com/go-dev-frame/sponge/pkg/sgorm/query"
)

var _ time.Time

// Tip: suggested filling in the binding rules https://github.com/go-playground/validator in request struct fields tag.

// CreateContractRequest 创建合约请求参数
type CreateContractRequest struct {
	NodeURL          string `json:"nodeURL" binding:""`          // 合约发布网站地址，如127.0.0.1:8454
	ChainID          string `json:"chainID" binding:""`          // 链 ID: 56=BSC 主网 97=BSC 测试网
	ContractAddress  string `json:"contractAddress" binding:""`  // 合约地址
	ContractName     string `json:"contractName" binding:""`     // 合约名称
	TxHash           string `json:"txHash" binding:""`           // 部署交易哈希
	PublisherAddress string `json:"publisherAddress" binding:""` // 合约发布者地址
	IsToken          bool   `json:"isToken"`                     // 是否为代币合约
	TokenSymbol      string `json:"tokenSymbol"`                 // 代币符号
	TokenDecimals    int    `json:"tokenDecimals"`               // 代币精度(小数点位数)
}

// UpdateContractByIDRequest 更新合约请求参数
type UpdateContractByIDRequest struct {
	ID               uint64 `json:"id" binding:""`               // uint64 id
	NodeURL          string `json:"nodeURL" binding:""`          // 合约发布网站地址，如127.0.0.1:8454
	ChainID          string `json:"chainID" binding:""`          // 链 ID: 56=BSC 主网 97=BSC 测试网
	ContractAddress  string `json:"contractAddress" binding:""`  // 合约地址
	ContractName     string `json:"contractName" binding:""`     // 合约名称
	TxHash           string `json:"txHash" binding:""`           // 部署交易哈希
	PublisherAddress string `json:"publisherAddress" binding:""` // 合约发布者地址
	IsToken          bool   `json:"isToken"`                     // 是否为代币合约
	TokenSymbol      string `json:"tokenSymbol"`                 // 代币符号
	TokenDecimals    int    `json:"tokenDecimals"`               // 代币精度(小数点位数)
}

// ContractObjDetail 合约详情
type ContractObjDetail struct {
	ID               uint64     `json:"id"`               // convert to uint64 id
	NodeURL          string     `json:"nodeURL"`          // 合约发布网站地址，如127.0.0.1:8454
	ChainID          string     `json:"chainID"`          // 链 ID: 56=BSC 主网 97=BSC 测试网
	ContractAddress  string     `json:"contractAddress"`  // 合约地址
	ContractName     string     `json:"contractName"`     // 合约名称
	TxHash           string     `json:"txHash"`           // 部署交易哈希
	PublisherAddress string     `json:"publisherAddress"` // 合约发布者地址
	IsToken          bool       `json:"isToken"`          // 是否为代币合约
	TokenSymbol      string     `json:"tokenSymbol"`      // 代币符号
	TokenDecimals    int        `json:"tokenDecimals"`    // 代币精度(小数点位数)
	CreatedAt        *time.Time `json:"createdAt"`        // 创建时间
	UpdatedAt        *time.Time `json:"updatedAt"`        // 更新时间
}

// CreateContractReply 仅用于 API 文档
type CreateContractReply struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		ID uint64 `json:"id"` // id
	} `json:"data"` // return data
}

// DeleteContractByIDReply 仅用于 API 文档
type DeleteContractByIDReply struct {
	Code int      `json:"code"` // return code
	Msg  string   `json:"msg"`  // return information description
	Data struct{} `json:"data"` // return data
}

// UpdateContractByIDReply 仅用于 API 文档
type UpdateContractByIDReply struct {
	Code int      `json:"code"` // return code
	Msg  string   `json:"msg"`  // return information description
	Data struct{} `json:"data"` // return data
}

// GetContractByIDReply 仅用于 API 文档
type GetContractByIDReply struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		Contract ContractObjDetail `json:"contract"`
	} `json:"data"` // return data
}

// ListContractsRequest 查询合约列表请求参数
type ListContractsRequest struct {
	query.Params
}

// ListContractsReply 仅用于 API 文档
type ListContractsReply struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		Contracts []ContractObjDetail `json:"contracts"`
	} `json:"data"` // return data
}

// DeployRequest 部署合约请求参数
type DeployRequest struct {
	RpcURL     string `json:"rpcUrl" binding:"required"`     // RPC URL（必填）
	PrivateKey string `json:"privateKey" binding:"required"` // 部署者私钥（必填）
	ChainName  string `json:"chainName"`                     // 链名称（如 hardhat）

	PoolOwner    string `json:"poolOwner"`    // PledgePool owner（默认部署者地址）
	OracleOwner  string `json:"oracleOwner"`  // BscPledgeOracle owner（默认部署者地址）
	FeeAddress   string `json:"feeAddress"`   // 手续费接收地址（默认部署者地址）
	FactoryFeeTo string `json:"factoryFeeTo"` // Factory feeToSetter（默认部署者地址）

	LendTokenName   string `json:"lendTokenName"`   // MockLendToken 名称
	LendTokenSym    string `json:"lendTokenSym"`    // MockLendToken 符号
	BorrowTokenName string `json:"borrowTokenName"` // MockBorrowToken 名称
	BorrowTokenSym  string `json:"borrowTokenSym"`  // MockBorrowToken 符号
	TokenSupply     string `json:"tokenSupply"`     // MockToken 初始发行量
	DebtTokenName   string `json:"debtTokenName"`   // 债务代币名称
	DebtTokenSym    string `json:"debtTokenSym"`    // 债务代币符号
}

// DeployReply 部署响应
type DeployReply struct {
	Code int        `json:"code"`
	Msg  string     `json:"msg"`
	Data DeployData `json:"data"`
}

// DeployData 部署数据
type DeployData struct {
	ChainName      string               `json:"chainName"`
	RpcURL         string               `json:"rpcUrl"`
	Deployer       string               `json:"deployer"`
	Contracts      []DeployContractItem `json:"contracts"`
	PledgePoolAddr string               `json:"pledgePoolAddr"`
}

// DeployContractItem 单个合约部署信息
type DeployContractItem struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	TxHash  string `json:"txHash"`
	Status  string `json:"status"`
}

// ListDeployRecordsRequest 查询部署记录请求
type ListDeployRecordsRequest struct {
	query.Params
}

// CreatePoolRequest 创建借贷池请求参数
type CreatePoolRequest struct {
	RpcURL              string `json:"rpcUrl" binding:"required"`
	PrivateKey          string `json:"privateKey" binding:"required"`
	PoolContractAddress string `json:"poolContractAddress" binding:"required"`
	ChainID             string `json:"chainID" binding:"required"`

	SettleTime             string `json:"settleTime" binding:"required"`
	EndTime                string `json:"endTime" binding:"required"`
	InterestRate           string `json:"interestRate" binding:"required"`
	MaxSupply              string `json:"maxSupply" binding:"required"`
	MortgageRate           string `json:"mortgageRate" binding:"required"`
	LendToken              string `json:"lendToken" binding:"required"`
	BorrowToken            string `json:"borrowToken" binding:"required"`
	LendDebtToken          string `json:"lendDebtToken" binding:"required"`
	BorrowDebtToken        string `json:"borrowDebtToken" binding:"required"`
	AutoLiquidateThreshold string `json:"autoLiquidateThreshold" binding:"required"`
	LendTokenSymbol        string `json:"lendTokenSymbol"`
	BorrowTokenSymbol      string `json:"borrowTokenSymbol"`
}

// CreatePoolReply 创建借贷池响应
type CreatePoolReply struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		PoolID uint64 `json:"poolID"`
		TxHash string `json:"txHash"`
	} `json:"data"`
}

// ListDeployRecordsReply 部署记录列表响应
type ListDeployRecordsReply struct {
	Code int            `json:"code"`
	Msg  string         `json:"msg"`
	Data ListDeployData `json:"data"`
}

// ListDeployData 部署记录列表数据
type ListDeployData struct {
	Records []DeployRecordItem `json:"records"`
	Total   int64              `json:"total"`
}

// DeployRecordItem 部署记录条目
type DeployRecordItem struct {
	ID              uint64 `json:"id"`
	ChainName       string `json:"chainName"`
	RpcURL          string `json:"rpcUrl"`
	ContractName    string `json:"contractName"`
	ContractAddress string `json:"contractAddress"`
	TxHash          string `json:"txHash"`
	DeployerAddress string `json:"deployerAddress"`
	Status          int    `json:"status"`
	CreatedAt       string `json:"createdAt"`
}

// ==================== PledgePool Operations ====================

// PoolWriteRequest PledgePool写入操作通用请求
type PoolWriteRequest struct {
	RpcURL              string `json:"rpcUrl" binding:"required"`              // RPC URL
	PrivateKey          string `json:"privateKey" binding:"required"`          // 私钥
	PoolContractAddress string `json:"poolContractAddress" binding:"required"` // PledgePool合约地址
	PoolID              string `json:"poolID" binding:"required"`              // 池子ID
}

// PoolWriteReply PledgePool写入操作通用响应
type PoolWriteReply struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		TxHash string `json:"txHash"` // 交易哈希
	} `json:"data"`
}

// PoolLendRequest 出借请求
type PoolAmountRequest struct {
	RpcURL              string `json:"rpcUrl" binding:"required"`              // RPC URL
	PrivateKey          string `json:"privateKey" binding:"required"`          // 私钥
	PoolContractAddress string `json:"poolContractAddress" binding:"required"` // PledgePool合约地址
	PoolID              string `json:"poolID" binding:"required"`              // 池子ID
	Amount              string `json:"amount" binding:"required"`              // 金额(最小单位)
}

// PoolSetFeeRequest 设置费率请求
type PoolSetFeeRequest struct {
	RpcURL              string `json:"rpcUrl" binding:"required"`              // RPC URL
	PrivateKey          string `json:"privateKey" binding:"required"`          // 私钥
	PoolContractAddress string `json:"poolContractAddress" binding:"required"` // PledgePool合约地址
	LendFee             string `json:"lendFee" binding:"required"`             // 出借费率
	BorrowFee           string `json:"borrowFee" binding:"required"`           // 借入费率
}

// PoolSetAddressRequest 设置地址请求
type PoolSetAddressRequest struct {
	RpcURL              string `json:"rpcUrl" binding:"required"`              // RPC URL
	PrivateKey          string `json:"privateKey" binding:"required"`          // 私钥
	PoolContractAddress string `json:"poolContractAddress" binding:"required"` // PledgePool合约地址
	NewAddress          string `json:"newAddress" binding:"required"`          // 新地址
}

// PoolSetMinAmountRequest 设置最小金额请求
type PoolSetMinAmountRequest struct {
	RpcURL              string `json:"rpcUrl" binding:"required"`              // RPC URL
	PrivateKey          string `json:"privateKey" binding:"required"`          // 私钥
	PoolContractAddress string `json:"poolContractAddress" binding:"required"` // PledgePool合约地址
	MinAmount           string `json:"minAmount" binding:"required"`           // 最小金额
}

// PoolTransferOwnershipRequest 转移所有权请求
type PoolTransferOwnershipRequest struct {
	RpcURL              string `json:"rpcUrl" binding:"required"`              // RPC URL
	PrivateKey          string `json:"privateKey" binding:"required"`          // 私钥
	PoolContractAddress string `json:"poolContractAddress" binding:"required"` // PledgePool合约地址
	NewOwner            string `json:"newOwner" binding:"required"`            // 新所有者地址
}

// PoolDestroyDebtRequest 销毁债务代币请求
type PoolDestroyDebtRequest struct {
	RpcURL              string `json:"rpcUrl" binding:"required"`              // RPC URL
	PrivateKey          string `json:"privateKey" binding:"required"`          // 私钥
	PoolContractAddress string `json:"poolContractAddress" binding:"required"` // PledgePool合约地址
	PoolID              string `json:"poolID" binding:"required"`              // 池子ID
	Amount              string `json:"amount" binding:"required"`              // 销毁数量
}

// PoolBorrowRequest 借入请求(可能 payable)
type PoolBorrowRequest struct {
	RpcURL              string `json:"rpcUrl" binding:"required"`              // RPC URL
	PrivateKey          string `json:"privateKey" binding:"required"`          // 私钥
	PoolContractAddress string `json:"poolContractAddress" binding:"required"` // PledgePool合约地址
	PoolID              string `json:"poolID" binding:"required"`              // 池子ID
	BorrowTokenAmount   string `json:"borrowTokenAmount" binding:"required"`   // 借入代币数量
}

// PoolReadRequest PledgePool读取操作通用请求
type PoolReadRequest struct {
	RpcURL              string `json:"rpcUrl" binding:"required"`              // RPC URL
	PoolContractAddress string `json:"poolContractAddress" binding:"required"` // PledgePool合约地址
	PoolID              string `json:"poolID" binding:"required"`              // 池子ID
}

// PoolStateReply 池子状态响应
type PoolStateReply struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		State string `json:"state"` // 池子状态
	} `json:"data"`
}

// PoolCheckReply 检查结果响应
type PoolCheckReply struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Result bool `json:"result"` // 检查结果
	} `json:"data"`
}

// PoolInfoData PledgePoolInfoList返回数据
type PoolInfoData struct {
	SettleTime             string `json:"settleTime"`
	EndTime                string `json:"endTime"`
	InterestRate           string `json:"interestRate"`
	MaxSupply              string `json:"maxSupply"`
	LendSupply             string `json:"lendSupply"`
	BorrowSupply           string `json:"borrowSupply"`
	MortgageRate           string `json:"mortgageRate"`
	LendToken              string `json:"lendToken"`
	BorrowToken            string `json:"borrowToken"`
	State                  uint8  `json:"state"`
	LendDebtToken          string `json:"lendDebtToken"`
	BorrowDebtToken        string `json:"borrowDebtToken"`
	AutoLiquidateThreshold string `json:"autoLiquidateThreshold"`
}

// PoolInfoReply 池子详情响应
type PoolInfoReply struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Info PoolInfoData `json:"info"`
	} `json:"data"`
}

// PoolDataInfo 链上PoolDataInfo数据
type PoolDataInfo struct {
	SettleAmountLend        string `json:"settleAmountLend"`
	SettleAmountBorrow      string `json:"settleAmountBorrow"`
	FinishAmountLend        string `json:"finishAmountLend"`
	FinishAmountBorrow      string `json:"finishAmountBorrow"`
	LiquidationAmountLend   string `json:"liquidationAmountLend"`
	LiquidationAmountBorrow string `json:"liquidationAmountBorrow"`
}

// PoolDataReply 池子清算数据响应
type PoolDataReply struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Data PoolDataInfo `json:"data"`
	} `json:"data"`
}

// PoolConfigReply 池子配置查询响应
type PoolConfigReply struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Oracle     string `json:"oracle"`
		FeeAddress string `json:"feeAddress"`
		SwapRouter string `json:"swapRouter"`
		LendFee    string `json:"lendFee"`
		BorrowFee  string `json:"borrowFee"`
		MinAmount  string `json:"minAmount"`
		Owner      string `json:"owner"`
	} `json:"data"`
}

// ==================== BscPledgeOracle Operations ====================

// OracleWriteRequest Oracle写入操作通用请求
type OracleWriteRequest struct {
	RpcURL        string `json:"rpcUrl" binding:"required"`
	PrivateKey    string `json:"privateKey" binding:"required"`
	OracleAddress string `json:"oracleAddress" binding:"required"` // Oracle合约地址
	NewAddress    string `json:"newAddress"`                       // 新地址(TransferOwnership用)
}

// OracleSetPriceRequest 设置价格请求
type OracleSetPriceRequest struct {
	RpcURL        string `json:"rpcUrl" binding:"required"`
	PrivateKey    string `json:"privateKey" binding:"required"`
	OracleAddress string `json:"oracleAddress" binding:"required"`
	AssetAddress  string `json:"assetAddress" binding:"required"` // 资产地址
	Price         string `json:"price" binding:"required"`        // 价格
}

// OracleSetAggregatorRequest 设置聚合器请求
type OracleSetAggregatorRequest struct {
	RpcURL        string `json:"rpcUrl" binding:"required"`
	PrivateKey    string `json:"privateKey" binding:"required"`
	OracleAddress string `json:"oracleAddress" binding:"required"`
	AssetAddress  string `json:"assetAddress" binding:"required"`  // 资产地址
	Aggregator    string `json:"aggregator" binding:"required"`    // 聚合器地址
	TokenDecimals int    `json:"tokenDecimals" binding:"required"` // 代币精度
}

// OracleReadRequest Oracle读取操作请求
type OracleReadRequest struct {
	RpcURL        string `json:"rpcUrl" binding:"required"`
	OracleAddress string `json:"oracleAddress" binding:"required"`
}

// OracleGetPriceRequest 查询价格请求
type OracleGetPriceRequest struct {
	RpcURL        string `json:"rpcUrl" binding:"required"`
	OracleAddress string `json:"oracleAddress" binding:"required"`
	AssetAddress  string `json:"assetAddress" binding:"required"` // 资产合约地址
}

// OraclePriceReply 价格响应
type OraclePriceReply struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Price string `json:"price"`
	} `json:"data"`
}

// ==================== ERC20 / DebtToken / WETH Operations ====================

// TokenWriteRequest Token写入操作通用请求
type TokenWriteRequest struct {
	RpcURL       string `json:"rpcUrl" binding:"required"`
	PrivateKey   string `json:"privateKey" binding:"required"`
	TokenAddress string `json:"tokenAddress" binding:"required"` // 代币合约地址
}

// TokenApproveRequest 授权请求
type TokenApproveRequest struct {
	RpcURL       string `json:"rpcUrl" binding:"required"`
	PrivateKey   string `json:"privateKey" binding:"required"`
	TokenAddress string `json:"tokenAddress" binding:"required"`
	Spender      string `json:"spender" binding:"required"` // 授权地址
	Amount       string `json:"amount" binding:"required"`  // 授权数量
}

// TokenTransferRequest 转账请求
type TokenTransferRequest struct {
	RpcURL       string `json:"rpcUrl" binding:"required"`
	PrivateKey   string `json:"privateKey" binding:"required"`
	TokenAddress string `json:"tokenAddress" binding:"required"`
	To           string `json:"to" binding:"required"`     // 接收地址
	Amount       string `json:"amount" binding:"required"` // 转账数量
}

// TokenReadRequest Token读取操作请求
type TokenReadRequest struct {
	RpcURL       string `json:"rpcUrl" binding:"required"`
	TokenAddress string `json:"tokenAddress" binding:"required"`
}

// TokenBalanceOfRequest 查询余额请求
type TokenBalanceOfRequest struct {
	RpcURL       string `json:"rpcUrl" binding:"required"`
	TokenAddress string `json:"tokenAddress" binding:"required"`
	Account      string `json:"account" binding:"required"` // 账户地址
}

// TokenAllowanceRequest 查询授权请求
type TokenAllowanceRequest struct {
	RpcURL       string `json:"rpcUrl" binding:"required"`
	TokenAddress string `json:"tokenAddress" binding:"required"`
	Owner        string `json:"owner" binding:"required"`   // 所有者
	Spender      string `json:"spender" binding:"required"` // 被授权者
}

// TokenBalanceReply 余额响应
type TokenBalanceReply struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Balance string `json:"balance"`
	} `json:"data"`
}

// TokenAllowanceReply 授权额度响应
type TokenAllowanceReply struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Allowance string `json:"allowance"`
	} `json:"data"`
}

// TokenSupplyReply 总供应量响应
type TokenSupplyReply struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		TotalSupply string `json:"totalSupply"`
	} `json:"data"`
}

// DebtTokenMintRequest 铸造债务代币请求
type DebtTokenMintRequest struct {
	RpcURL       string `json:"rpcUrl" binding:"required"`
	PrivateKey   string `json:"privateKey" binding:"required"`
	TokenAddress string `json:"tokenAddress" binding:"required"`
	Account      string `json:"account" binding:"required"` // 接收地址
	Amount       string `json:"amount" binding:"required"`  // 铸造数量
}

// DebtTokenBurnRequest 销毁债务代币请求
type DebtTokenBurnRequest struct {
	RpcURL       string `json:"rpcUrl" binding:"required"`
	PrivateKey   string `json:"privateKey" binding:"required"`
	TokenAddress string `json:"tokenAddress" binding:"required"`
	Account      string `json:"account" binding:"required"` // 销毁地址
	Amount       string `json:"amount" binding:"required"`  // 销毁数量
}

// DebtTokenSetMinterRequest 设置铸造者请求
type DebtTokenSetMinterRequest struct {
	RpcURL       string `json:"rpcUrl" binding:"required"`
	PrivateKey   string `json:"privateKey" binding:"required"`
	TokenAddress string `json:"tokenAddress" binding:"required"`
	Minter       string `json:"minter" binding:"required"` // 铸造者地址
	Status       bool   `json:"status" binding:"required"` // 是否启用
}

// WETHDepositRequest WETH充值请求
type WETHDepositRequest struct {
	RpcURL       string `json:"rpcUrl" binding:"required"`
	PrivateKey   string `json:"privateKey" binding:"required"`
	TokenAddress string `json:"tokenAddress" binding:"required"`
	Amount       string `json:"amount" binding:"required"` // 充值ETH数量(wei)
}

// ==================== UniswapV2Factory Operations ====================

// FactoryCreatePairRequest 创建交易对请求
type FactoryCreatePairRequest struct {
	RpcURL         string `json:"rpcUrl" binding:"required"`
	PrivateKey     string `json:"privateKey" binding:"required"`
	FactoryAddress string `json:"factoryAddress" binding:"required"`
	TokenA         string `json:"tokenA" binding:"required"` // 代币A地址
	TokenB         string `json:"tokenB" binding:"required"` // 代币B地址
}

// FactorySetFeeToRequest 设置手续费地址请求
type FactorySetFeeToRequest struct {
	RpcURL         string `json:"rpcUrl" binding:"required"`
	PrivateKey     string `json:"privateKey" binding:"required"`
	FactoryAddress string `json:"factoryAddress" binding:"required"`
	FeeTo          string `json:"feeTo" binding:"required"` // 新手续费地址
}

// FactoryGetPairRequest 查询交易对请求
type FactoryGetPairRequest struct {
	RpcURL         string `json:"rpcUrl" binding:"required"`
	FactoryAddress string `json:"factoryAddress" binding:"required"`
	TokenA         string `json:"tokenA" binding:"required"`
	TokenB         string `json:"tokenB" binding:"required"`
}

// FactoryPairReply 交易对查询响应
type FactoryPairReply struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		PairAddress string `json:"pairAddress"`
	} `json:"data"`
}

// FactoryCreatePairReply 创建交易对响应
type FactoryCreatePairReply struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		TxHash      string `json:"txHash"`
		PairAddress string `json:"pairAddress"`
	} `json:"data"`
}

// ==================== UniswapV2Router02 Operations ====================

// RouterAddLiquidityRequest 添加流动性请求
type RouterAddLiquidityRequest struct {
	RpcURL         string `json:"rpcUrl" binding:"required"`
	PrivateKey     string `json:"privateKey" binding:"required"`
	RouterAddress  string `json:"routerAddress" binding:"required"`
	TokenA         string `json:"tokenA" binding:"required"`
	TokenB         string `json:"tokenB" binding:"required"`
	AmountADesired string `json:"amountADesired" binding:"required"`
	AmountBDesired string `json:"amountBDesired" binding:"required"`
	AmountAMin     string `json:"amountAMin" binding:"required"`
	AmountBMin     string `json:"amountBMin" binding:"required"`
	To             string `json:"to" binding:"required"`
	Deadline       string `json:"deadline" binding:"required"`
}

// RouterSwapRequest 兑换请求
type RouterSwapRequest struct {
	RpcURL        string   `json:"rpcUrl" binding:"required"`
	PrivateKey    string   `json:"privateKey" binding:"required"`
	RouterAddress string   `json:"routerAddress" binding:"required"`
	AmountIn      string   `json:"amountIn" binding:"required"`
	AmountOutMin  string   `json:"amountOutMin" binding:"required"`
	Path          []string `json:"path" binding:"required"`
	To            string   `json:"to" binding:"required"`
	Deadline      string   `json:"deadline" binding:"required"`
}

// RouterGetAmountsRequest 查询金额请求
type RouterGetAmountsRequest struct {
	RpcURL        string   `json:"rpcUrl" binding:"required"`
	RouterAddress string   `json:"routerAddress" binding:"required"`
	AmountIn      string   `json:"amountIn" binding:"required"`
	Path          []string `json:"path" binding:"required"`
}

// RouterAmountsReply 金额查询响应
type RouterAmountsReply struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Amounts []string `json:"amounts"`
	} `json:"data"`
}
