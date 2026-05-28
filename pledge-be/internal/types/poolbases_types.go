package types

import (
	"time"

	"github.com/go-dev-frame/sponge/pkg/sgorm/query"
)

var _ time.Time

// Tip: suggested filling in the binding rules https://github.com/go-playground/validator in request struct fields tag.

// CreatePoolbasesRequest 创建池子基础信息请求参数
type CreatePoolbasesRequest struct {
	PoolID                 int    `json:"poolID" binding:""`                 // 业务池子 ID，与链上 pool 对应
	SettleTime             string `json:"settleTime" binding:""`             // 结算开始时间戳
	EndTime                string `json:"endTime" binding:""`                // 池子结束时间戳
	InterestRate           string `json:"interestRate" binding:""`           // 利率(精度值，如 10000000 表示 1%)
	MaxSupply              string `json:"maxSupply" binding:""`              // 池子最大可借/可存额度(最小单位)
	LendSupply             string `json:"lendSupply" binding:""`             // 当前已出借总量
	BorrowSupply           string `json:"borrowSupply" binding:""`           // 当前已借入(抵押)总量
	MortgageRate           string `json:"mortgageRate" binding:""`           // 抵押率(精度值，如 10000000=100%)
	LendToken              string `json:"lendToken" binding:""`              // 出借资产合约地址(用户借出的币)
	BorrowToken            string `json:"borrowToken" binding:""`            // 借入(抵押)资产合约地址(用户抵押的币)
	State                  string `json:"state" binding:""`                  // 池子状态: 0未开启 1进行中 2已结算 3清算中 4未开启等
	LendDebtToken          string `json:"lendDebtToken" binding:""`          // 出借侧池子代币/合约地址(出借凭证)
	BorrowDebtToken        string `json:"borrowDebtToken" binding:""`        // 借入(抵押)侧池子代币/合约地址(质押资产)
	AutoLiquidateThreshold string `json:"autoLiquidateThreshold" binding:""` // 自动清算阈值(精度值)
	BorrowTokenInfo        string `json:"borrowTokenInfo" binding:""`        // 借入(抵押)代币信息: tokenName, tokenLogo, tokenPrice, borrowFee 等
	LendTokenInfo          string `json:"lendTokenInfo" binding:""`          // 出借代币信息: tokenName, tokenLogo, tokenPrice, lendFee 等
	ChainID                string `json:"chainID" binding:""`                // 链 ID: 56=BSC 主网 97=BSC 测试网
	LendTokenSymbol        string `json:"lendTokenSymbol" binding:""`        // 出借代币符号，如 BUSD
	BorrowTokenSymbol      string `json:"borrowTokenSymbol" binding:""`      // 借入(抵押)代币符号，如 BTC
}

// UpdatePoolbasesByIDRequest 更新池子基础信息请求参数
type UpdatePoolbasesByIDRequest struct {
	ID uint64 `json:"id" binding:""` // uint64 id
	// 主键ID
	PoolID                 int    `json:"poolID" binding:""`                 // 业务池子 ID，与链上 pool 对应
	SettleTime             string `json:"settleTime" binding:""`             // 结算开始时间戳
	EndTime                string `json:"endTime" binding:""`                // 池子结束时间戳
	InterestRate           string `json:"interestRate" binding:""`           // 利率(精度值，如 10000000 表示 1%)
	MaxSupply              string `json:"maxSupply" binding:""`              // 池子最大可借/可存额度(最小单位)
	LendSupply             string `json:"lendSupply" binding:""`             // 当前已出借总量
	BorrowSupply           string `json:"borrowSupply" binding:""`           // 当前已借入(抵押)总量
	MortgageRate           string `json:"mortgageRate" binding:""`           // 抵押率(精度值，如 10000000=100%)
	LendToken              string `json:"lendToken" binding:""`              // 出借资产合约地址(用户借出的币)
	BorrowToken            string `json:"borrowToken" binding:""`            // 借入(抵押)资产合约地址(用户抵押的币)
	State                  string `json:"state" binding:""`                  // 池子状态: 0未开启 1进行中 2已结算 3清算中 4未开启等
	LendDebtToken          string `json:"lendDebtToken" binding:""`          // 出借侧池子代币/合约地址(出借凭证)
	BorrowDebtToken        string `json:"borrowDebtToken" binding:""`        // 借入(抵押)侧池子代币/合约地址(质押资产)
	AutoLiquidateThreshold string `json:"autoLiquidateThreshold" binding:""` // 自动清算阈值(精度值)
	BorrowTokenInfo        string `json:"borrowTokenInfo" binding:""`        // 借入(抵押)代币信息: tokenName, tokenLogo, tokenPrice, borrowFee 等
	LendTokenInfo          string `json:"lendTokenInfo" binding:""`          // 出借代币信息: tokenName, tokenLogo, tokenPrice, lendFee 等
	ChainID                string `json:"chainID" binding:""`                // 链 ID: 56=BSC 主网 97=BSC 测试网
	LendTokenSymbol        string `json:"lendTokenSymbol" binding:""`        // 出借代币符号，如 BUSD
	BorrowTokenSymbol      string `json:"borrowTokenSymbol" binding:""`      // 借入(抵押)代币符号，如 BTC
}

// PoolbasesObjDetail 池子基础信息详情
type PoolbasesObjDetail struct {
	ID uint64 `json:"id"` // convert to uint64 id
	// 主键ID
	PoolID                 int        `json:"poolID"`                 // 业务池子 ID，与链上 pool 对应
	SettleTime             string     `json:"settleTime"`             // 结算开始时间戳
	EndTime                string     `json:"endTime"`                // 池子结束时间戳
	InterestRate           string     `json:"interestRate"`           // 利率(精度值，如 10000000 表示 1%)
	MaxSupply              string     `json:"maxSupply"`              // 池子最大可借/可存额度(最小单位)
	LendSupply             string     `json:"lendSupply"`             // 当前已出借总量
	BorrowSupply           string     `json:"borrowSupply"`           // 当前已借入(抵押)总量
	MortgageRate           string     `json:"mortgageRate"`           // 抵押率(精度值，如 10000000=100%)
	LendToken              string     `json:"lendToken"`              // 出借资产合约地址(用户借出的币)
	BorrowToken            string     `json:"borrowToken"`            // 借入(抵押)资产合约地址(用户抵押的币)
	State                  string     `json:"state"`                  // 池子状态: 0未开启 1进行中 2已结算 3清算中 4未开启等
	LendDebtToken          string     `json:"lendDebtToken"`          // 出借侧池子代币/合约地址(出借凭证)
	BorrowDebtToken        string     `json:"borrowDebtToken"`        // 借入(抵押)侧池子代币/合约地址(质押资产)
	AutoLiquidateThreshold string     `json:"autoLiquidateThreshold"` // 自动清算阈值(精度值)
	BorrowTokenInfo        string     `json:"borrowTokenInfo"`        // 借入(抵押)代币信息: tokenName, tokenLogo, tokenPrice, borrowFee 等
	LendTokenInfo          string     `json:"lendTokenInfo"`          // 出借代币信息: tokenName, tokenLogo, tokenPrice, lendFee 等
	ChainID                string     `json:"chainID"`                // 链 ID: 56=BSC 主网 97=BSC 测试网
	LendTokenSymbol        string     `json:"lendTokenSymbol"`        // 出借代币符号，如 BUSD
	BorrowTokenSymbol      string     `json:"borrowTokenSymbol"`      // 借入(抵押)代币符号，如 BTC
	CreatedAt              *time.Time `json:"createdAt"`              // 创建时间
	UpdatedAt              *time.Time `json:"updatedAt"`              // 更新时间
}

// CreatePoolbasesReply 仅用于 API 文档
type CreatePoolbasesReply struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		ID uint64 `json:"id"` // id
	} `json:"data"` // return data
}

// DeletePoolbasesByIDReply 仅用于 API 文档
type DeletePoolbasesByIDReply struct {
	Code int      `json:"code"` // return code
	Msg  string   `json:"msg"`  // return information description
	Data struct{} `json:"data"` // return data
}

// UpdatePoolbasesByIDReply 仅用于 API 文档
type UpdatePoolbasesByIDReply struct {
	Code int      `json:"code"` // return code
	Msg  string   `json:"msg"`  // return information description
	Data struct{} `json:"data"` // return data
}

// GetPoolbasesByIDReply 仅用于 API 文档
type GetPoolbasesByIDReply struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		Poolbases PoolbasesObjDetail `json:"poolbases"`
	} `json:"data"` // return data
}

// ListPoolbasessRequest 查询池子基础信息列表请求参数
type ListPoolbasessRequest struct {
	query.Params
}

// ListPoolbasessReply 仅用于 API 文档
type ListPoolbasessReply struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		Poolbasess []PoolbasesObjDetail `json:"poolbasess"`
	} `json:"data"` // return data
}
