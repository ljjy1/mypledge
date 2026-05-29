package model

import (
	"time"
)

// Contract 合约表
type Contract struct {
	ID               uint64     `gorm:"column:id;type:bigint(20) unsigned;primary_key;AUTO_INCREMENT" json:"id"`             // 主键ID
	NodeURL          string     `gorm:"column:node_url;type:varchar(255);not null" json:"nodeURL"`                           // 合约发布网站地址，如127.0.0.1:8454
	ChainID          string     `gorm:"column:chain_id;type:varchar(20);default:56" json:"chainID"`                          // 链 ID: 56=BSC 主网 97=BSC 测试网
	ContractAddress  string     `gorm:"column:contract_address;type:varchar(255);not null" json:"contractAddress"`           // 合约地址
	ContractName     string     `gorm:"column:contract_name;type:varchar(100)" json:"contractName"`                          // 合约名称（如 PledgePool、WETH）
	TxHash           string     `gorm:"column:tx_hash;type:varchar(100)" json:"txHash"`                                      // 部署交易哈希
	PublisherAddress string     `gorm:"column:publisher_address;type:varchar(255);not null" json:"publisherAddress"`         // 合约发布者地址
	IsToken          bool       `gorm:"column:is_token;type:tinyint(1);default:0" json:"isToken"`                            // 是否为代币合约
	TokenSymbol      string     `gorm:"column:token_symbol;type:varchar(50)" json:"tokenSymbol"`                             // 代币符号
	TokenDecimals    int        `gorm:"column:token_decimals;type:int(11);default:0" json:"tokenDecimals"`                   // 代币精度(小数点位数)
	CreatedAt        *time.Time `gorm:"column:created_at;type:datetime;default:CURRENT_TIMESTAMP;not null" json:"createdAt"` // 创建时间
	UpdatedAt        *time.Time `gorm:"column:updated_at;type:datetime;default:CURRENT_TIMESTAMP;not null" json:"updatedAt"` // 更新时间
}

// TableName 返回合约表名
func (m *Contract) TableName() string {
	return "contract"
}

// ContractColumnNames 自定义查询字段白名单，防止 SQL 注入攻击
var ContractColumnNames = map[string]bool{
	"id":                true,
	"node_url":          true,
	"chain_id":          true,
	"contract_address":  true,
	"contract_name":     true,
	"tx_hash":           true,
	"publisher_address": true,
	"is_token":          true,
	"token_symbol":      true,
	"token_decimals":    true,
	"created_at":        true,
	"updated_at":        true,
}
