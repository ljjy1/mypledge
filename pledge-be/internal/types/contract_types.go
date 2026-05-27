package types

import (
	"time"

	"github.com/go-dev-frame/sponge/pkg/sgorm/query"
)

var _ time.Time

// Tip: suggested filling in the binding rules https://github.com/go-playground/validator in request struct fields tag.

// CreateContractRequest request params
type CreateContractRequest struct {
	NodeURL          string `json:"nodeURL" binding:""`          // 合约发布网站地址，如127.0.0.1:8454
	ChainID          string `json:"chainID" binding:""`          // 链 ID: 56=BSC 主网 97=BSC 测试网
	ContractAddress  string `json:"contractAddress" binding:""`  // 合约地址
	ContractName     string `json:"contractName" binding:""`     // 合约名称
	TxHash           string `json:"txHash" binding:""`           // 部署交易哈希
	PublisherAddress string `json:"publisherAddress" binding:""` // 合约发布者地址
}

// UpdateContractByIDRequest request params
type UpdateContractByIDRequest struct {
	ID uint64 `json:"id" binding:""` // uint64 id
	// 主键ID
	NodeURL          string `json:"nodeURL" binding:""`          // 合约发布网站地址，如127.0.0.1:8454
	ChainID          string `json:"chainID" binding:""`          // 链 ID: 56=BSC 主网 97=BSC 测试网
	ContractAddress  string `json:"contractAddress" binding:""`  // 合约地址
	ContractName     string `json:"contractName" binding:""`     // 合约名称
	TxHash           string `json:"txHash" binding:""`           // 部署交易哈希
	PublisherAddress string `json:"publisherAddress" binding:""` // 合约发布者地址
}

// ContractObjDetail detail
type ContractObjDetail struct {
	ID uint64 `json:"id"` // convert to uint64 id
	// 主键ID
	NodeURL          string     `json:"nodeURL"`          // 合约发布网站地址，如127.0.0.1:8454
	ChainID          string     `json:"chainID"`          // 链 ID: 56=BSC 主网 97=BSC 测试网
	ContractAddress  string     `json:"contractAddress"`  // 合约地址
	ContractName     string     `json:"contractName"`     // 合约名称
	TxHash           string     `json:"txHash"`           // 部署交易哈希
	PublisherAddress string     `json:"publisherAddress"` // 合约发布者地址
	CreatedAt        *time.Time `json:"createdAt"`        // 创建时间
	UpdatedAt        *time.Time `json:"updatedAt"`        // 更新时间
}

// CreateContractReply only for api docs
type CreateContractReply struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		ID uint64 `json:"id"` // id
	} `json:"data"` // return data
}

// DeleteContractByIDReply only for api docs
type DeleteContractByIDReply struct {
	Code int      `json:"code"` // return code
	Msg  string   `json:"msg"`  // return information description
	Data struct{} `json:"data"` // return data
}

// UpdateContractByIDReply only for api docs
type UpdateContractByIDReply struct {
	Code int      `json:"code"` // return code
	Msg  string   `json:"msg"`  // return information description
	Data struct{} `json:"data"` // return data
}

// GetContractByIDReply only for api docs
type GetContractByIDReply struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		Contract ContractObjDetail `json:"contract"`
	} `json:"data"` // return data
}

// ListContractsRequest request params
type ListContractsRequest struct {
	query.Params
}

// ListContractsReply only for api docs
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
