package model

import (
	"time"
)

// PoolState 资金池状态枚举
type PoolState string

const (
	PoolStateMatch       PoolState = "MATCH"       // 撮合中/匹配中，允许存款和质押
	PoolStateExecution   PoolState = "EXECUTION"   // 已结算进入执行期
	PoolStateFinish      PoolState = "FINISH"      // 到期正常结束
	PoolStateLiquidation PoolState = "LIQUIDATION" // 已触发清算
	PoolStateUndone      PoolState = "UNDONE"      // 极端情况下未成立
)

// Poolbases 借贷池主表(按链+pool_id)，包含基础信息与结算清算数据
type Poolbases struct {
	ID                     uint64     `gorm:"column:id;type:bigint(20) unsigned;primary_key;AUTO_INCREMENT" json:"id"`             // 主键ID
	ContractID             uint64     `gorm:"column:contract_id;type:bigint(20) unsigned;default:0" json:"contractID"`             // 关联的合约 ID，对应 contract.id
	PoolID                 int        `gorm:"column:pool_id;type:int(11)" json:"poolID"`                                           // 业务池子 ID，与链上 pool 对应
	SettleTime             string     `gorm:"column:settle_time;type:varchar(100)" json:"settleTime"`                              // 结算开始时间戳
	EndTime                string     `gorm:"column:end_time;type:varchar(100)" json:"endTime"`                                    // 池子结束时间戳
	InterestRate           string     `gorm:"column:interest_rate;type:varchar(100)" json:"interestRate"`                          // 利率(精度值，如 10000000 表示 1%)
	MaxSupply              string     `gorm:"column:max_supply;type:varchar(100)" json:"maxSupply"`                                // 池子最大可借/可存额度(最小单位)
	LendSupply             string     `gorm:"column:lend_supply;type:varchar(100)" json:"lendSupply"`                              // 当前已出借总量
	BorrowSupply           string     `gorm:"column:borrow_supply;type:varchar(100)" json:"borrowSupply"`                          // 当前已借入(抵押)总量
	MortgageRate           string     `gorm:"column:mortgage_rate;type:varchar(100)" json:"mortgageRate"`                          // 抵押率(精度值，如 10000000=100%)
	LendToken              string     `gorm:"column:lend_token;type:varchar(100)" json:"lendToken"`                                // 出借资产合约地址(用户借出的币)
	BorrowToken            string     `gorm:"column:borrow_token;type:varchar(100)" json:"borrowToken"`                            // 借入(抵押)资产合约地址(用户抵押的币)
	State                  PoolState  `gorm:"column:state;type:varchar(100)" json:"state"`                                         // 池子状态: MATCH=撮合中 EXECUTION=执行期 FINISH=结束 LIQUIDATION=清算 UNDONE=未成立
	LendDebtToken          string     `gorm:"column:lend_debt_token;type:varchar(100)" json:"lendDebtToken"`                       // 出借侧池子代币/合约地址(出借凭证)
	BorrowDebtToken        string     `gorm:"column:borrow_debt_token;type:varchar(100)" json:"borrowDebtToken"`                   // 借入(抵押)侧池子代币/合约地址(质押资产)
	AutoLiquidateThreshold string     `gorm:"column:auto_liquidate_threshold;type:varchar(100)" json:"autoLiquidateThreshold"`     // 自动清算阈值(精度值)
	ChainID                string     `gorm:"column:chain_id;type:varchar(20);default:56" json:"chainID"`                          // 链 ID: 56=BSC 主网 97=BSC 测试网
	SettleAmountLend       string     `gorm:"column:settle_amount_lend;type:varchar(100)" json:"settleAmountLend"`                 // 结算时借出侧(贷方)金额
	SettleAmountBorrow     string     `gorm:"column:settle_amount_borrow;type:varchar(100)" json:"settleAmountBorrow"`             // 结算时借入侧(抵押)金额
	FinishAmountLend       string     `gorm:"column:finish_amount_lend;type:varchar(100)" json:"finishAmountLend"`                 // 已完成/归还的借出侧金额
	FinishAmountBorrow     string     `gorm:"column:finish_amount_borrow;type:varchar(100)" json:"finishAmountBorrow"`             // 已完成/归还的借入侧金额
	LiquidationAmounLend   string     `gorm:"column:liquidation_amoun_lend;type:varchar(100)" json:"liquidationAmounLend"`         // 清算产生的借出侧金额
	LiquidationAmounBorrow string     `gorm:"column:liquidation_amoun_borrow;type:varchar(100)" json:"liquidationAmounBorrow"`     // 清算产生的借入侧金额
	CreatedAt              *time.Time `gorm:"column:created_at;type:datetime;default:CURRENT_TIMESTAMP;not null" json:"createdAt"` // 创建时间
	UpdatedAt              *time.Time `gorm:"column:updated_at;type:datetime;default:CURRENT_TIMESTAMP;not null" json:"updatedAt"` // 更新时间
}

// PoolbasesColumnNames 自定义查询字段白名单，防止 SQL 注入攻击
var PoolbasesColumnNames = map[string]bool{
	"id":                       true,
	"contract_id":              true,
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
	"chain_id":                 true,
	"settle_amount_lend":       true,
	"settle_amount_borrow":     true,
	"finish_amount_lend":       true,
	"finish_amount_borrow":     true,
	"liquidation_amoun_lend":   true,
	"liquidation_amoun_borrow": true,
	"created_at":               true,
	"updated_at":               true,
}
