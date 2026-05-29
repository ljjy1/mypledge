package model

import (
	"time"
)

// Pooldata 池子结算与清算数据表
type Pooldata struct {
	ID                     uint64     `gorm:"column:id;type:bigint(20) unsigned;primary_key;AUTO_INCREMENT" json:"id"`             // 主键ID
	ContractID             uint64     `gorm:"column:contract_id;type:bigint(20) unsigned;default:0" json:"contractID"`             // 关联的合约 ID，对应 contract.id
	ChainID                string     `gorm:"column:chain_id;type:varchar(20);default:56" json:"chainID"`                          // 链 ID，与 poolbases 一致
	PoolID                 string     `gorm:"column:pool_id;type:varchar(50)" json:"poolID"`                                       // 池子 ID，与 poolbases.pool_id 对应
	SettleAmountLend       string     `gorm:"column:settle_amount_lend;type:varchar(100)" json:"settleAmountLend"`                 // 结算时借出侧(贷方)金额
	SettleAmountBorrow     string     `gorm:"column:settle_amount_borrow;type:varchar(100)" json:"settleAmountBorrow"`             // 结算时借入侧(抵押)金额
	FinishAmountLend       string     `gorm:"column:finish_amount_lend;type:varchar(100)" json:"finishAmountLend"`                 // 已完成/归还的借出侧金额
	FinishAmountBorrow     string     `gorm:"column:finish_amount_borrow;type:varchar(100)" json:"finishAmountBorrow"`             // 已完成/归还的借入侧金额
	LiquidationAmounLend   string     `gorm:"column:liquidation_amoun_lend;type:varchar(100)" json:"liquidationAmounLend"`         // 清算产生的借出侧金额
	LiquidationAmounBorrow string     `gorm:"column:liquidation_amoun_borrow;type:varchar(100)" json:"liquidationAmounBorrow"`     // 清算产生的借入侧金额
	CreatedAt              *time.Time `gorm:"column:created_at;type:datetime;default:CURRENT_TIMESTAMP;not null" json:"createdAt"` // 创建时间
	UpdatedAt              *time.Time `gorm:"column:updated_at;type:datetime;default:CURRENT_TIMESTAMP;not null" json:"updatedAt"` // 更新时间
}

// PooldataColumnNames 自定义查询字段白名单，防止 SQL 注入攻击
var PooldataColumnNames = map[string]bool{
	"id":                       true,
	"contract_id":              true,
	"chain_id":                 true,
	"pool_id":                  true,
	"settle_amount_lend":       true,
	"settle_amount_borrow":     true,
	"finish_amount_lend":       true,
	"finish_amount_borrow":     true,
	"liquidation_amoun_lend":   true,
	"liquidation_amoun_borrow": true,
	"created_at":               true,
	"updated_at":               true,
}
