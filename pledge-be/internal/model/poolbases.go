package model

import (
	"gorm.io/datatypes"
	"time"
)

// Poolbases 借贷池主表(按链+pool_id)
type Poolbases struct {
	ID                     uint64          `gorm:"column:id;type:bigint(20) unsigned;primary_key;AUTO_INCREMENT" json:"id"`             // 主键ID
	PoolID                 int             `gorm:"column:pool_id;type:int(11)" json:"poolID"`                                           // 业务池子 ID，与链上 pool 对应
	SettleTime             string          `gorm:"column:settle_time;type:varchar(100)" json:"settleTime"`                              // 结算开始时间戳
	EndTime                string          `gorm:"column:end_time;type:varchar(100)" json:"endTime"`                                    // 池子结束时间戳
	InterestRate           string          `gorm:"column:interest_rate;type:varchar(100)" json:"interestRate"`                          // 利率(精度值，如 10000000 表示 1%)
	MaxSupply              string          `gorm:"column:max_supply;type:varchar(100)" json:"maxSupply"`                                // 池子最大可借/可存额度(最小单位)
	LendSupply             string          `gorm:"column:lend_supply;type:varchar(100)" json:"lendSupply"`                              // 当前已出借总量
	BorrowSupply           string          `gorm:"column:borrow_supply;type:varchar(100)" json:"borrowSupply"`                          // 当前已借入(抵押)总量
	MortgageRate           string          `gorm:"column:mortgage_rate;type:varchar(100)" json:"mortgageRate"`                          // 抵押率(精度值，如 10000000=100%)
	LendToken              string          `gorm:"column:lend_token;type:varchar(100)" json:"lendToken"`                                // 出借资产合约地址(用户借出的币)
	BorrowToken            string          `gorm:"column:borrow_token;type:varchar(100)" json:"borrowToken"`                            // 借入(抵押)资产合约地址(用户抵押的币)
	State                  string          `gorm:"column:state;type:varchar(100)" json:"state"`                                         // 池子状态: 0未开启 1进行中 2已结算 3清算中 4未开启等
	LendDebtToken          string          `gorm:"column:lend_debt_token;type:varchar(100)" json:"lendDebtToken"`                       // 出借侧池子代币/合约地址(出借凭证)
	BorrowDebtToken        string          `gorm:"column:borrow_debt_token;type:varchar(100)" json:"borrowDebtToken"`                   // 借入(抵押)侧池子代币/合约地址(质押资产)
	AutoLiquidateThreshold string          `gorm:"column:auto_liquidate_threshold;type:varchar(100)" json:"autoLiquidateThreshold"`     // 自动清算阈值(精度值)
	BorrowTokenInfo        *datatypes.JSON `gorm:"column:borrow_token_info;type:json" json:"borrowTokenInfo"`                           // 借入(抵押)代币信息: tokenName, tokenLogo, tokenPrice, borrowFee 等
	LendTokenInfo          *datatypes.JSON `gorm:"column:lend_token_info;type:json" json:"lendTokenInfo"`                               // 出借代币信息: tokenName, tokenLogo, tokenPrice, lendFee 等
	ChainID                string          `gorm:"column:chain_id;type:varchar(20);default:56" json:"chainID"`                          // 链 ID: 56=BSC 主网 97=BSC 测试网
	LendTokenSymbol        string          `gorm:"column:lend_token_symbol;type:varchar(100)" json:"lendTokenSymbol"`                   // 出借代币符号，如 BUSD
	BorrowTokenSymbol      string          `gorm:"column:borrow_token_symbol;type:varchar(100)" json:"borrowTokenSymbol"`               // 借入(抵押)代币符号，如 BTC
	CreatedAt              *time.Time      `gorm:"column:created_at;type:datetime;default:CURRENT_TIMESTAMP;not null" json:"createdAt"` // 创建时间
	UpdatedAt              *time.Time      `gorm:"column:updated_at;type:datetime;default:CURRENT_TIMESTAMP;not null" json:"updatedAt"` // 更新时间
}

// PoolbasesColumnNames Whitelist for custom query fields to prevent sql injection attacks
var PoolbasesColumnNames = map[string]bool{
	"id":                       true,
	"pool_id":                  true,
	"settle_time":              true,
	"end_time":                 true,
	"interest_rate":            true,
	"max_supply":               true,
	"lend_supply":              true,
	"borrow_supply":            true,
	"mortgage_rate":            true,
	"lend_token":               true,
	"borrow_token":             true,
	"state":                    true,
	"lend_debt_token":          true,
	"borrow_debt_token":        true,
	"auto_liquidate_threshold": true,
	"borrow_token_info":        true,
	"lend_token_info":          true,
	"chain_id":                 true,
	"lend_token_symbol":        true,
	"borrow_token_symbol":      true,
	"created_at":               true,
	"updated_at":               true,
}
