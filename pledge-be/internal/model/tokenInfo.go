package model

import (
	"gorm.io/datatypes"
	"time"
)

// TokenInfo 代币信息表(按链)
type TokenInfo struct {
	ID          uint64          `gorm:"column:id;type:bigint(20) unsigned;primary_key;AUTO_INCREMENT" json:"id"`             // 主键ID
	Symbol      string          `gorm:"column:symbol;type:varchar(100)" json:"symbol"`                                       // 代币符号，如 BUSD/BTC
	Logo        string          `gorm:"column:logo;type:varchar(150)" json:"logo"`                                           // 代币 logo URL
	Price       string          `gorm:"column:price;type:varchar(50)" json:"price"`                                          // 价格(精度值，用于估值与清算)
	Token       string          `gorm:"column:token;type:varchar(100)" json:"token"`                                         // 代币合约地址
	ChainID     string          `gorm:"column:chain_id;type:varchar(20);default:56" json:"chainID"`                          // 链 ID: 56=BSC 97=测试网
	ContractAbi *datatypes.JSON `gorm:"column:contract_abi;type:json;not null" json:"contractAbi"`                           // 代币ABI
	ContractBin *datatypes.JSON `gorm:"column:contract_bin;type:json;not null" json:"contractBin"`                           // 代币BIN
	Decimals    int             `gorm:"column:decimals;type:int(11);not null" json:"decimals"`                               // 代币精度(小数位数)
	CreatedAt   *time.Time      `gorm:"column:created_at;type:datetime;default:CURRENT_TIMESTAMP;not null" json:"createdAt"` // 创建时间
	UpdatedAt   *time.Time      `gorm:"column:updated_at;type:datetime;default:CURRENT_TIMESTAMP;not null" json:"updatedAt"` // 更新时间
}

// TableName table name
func (m *TokenInfo) TableName() string {
	return "token_info"
}

// TokenInfoColumnNames Whitelist for custom query fields to prevent sql injection attacks
var TokenInfoColumnNames = map[string]bool{
	"id":           true,
	"symbol":       true,
	"logo":         true,
	"price":        true,
	"token":        true,
	"chain_id":     true,
	"contract_abi": true,
	"contract_bin": true,
	"decimals":     true,
	"created_at":   true,
	"updated_at":   true,
}
