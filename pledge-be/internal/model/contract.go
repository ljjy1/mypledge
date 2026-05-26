package model

import (
	"time"

	"gorm.io/datatypes"
)

// Contract 合约表
type Contract struct {
	ID               uint64          `gorm:"column:id;type:bigint(20) unsigned;primary_key;AUTO_INCREMENT" json:"id"`             // 主键ID
	NodeURL          string          `gorm:"column:node_url;type:varchar(255);not null" json:"nodeURL"`                           // 合约发布网站地址，如127.0.0.1:8454
	ChainID          string          `gorm:"column:chain_id;type:varchar(20);default:56" json:"chainID"`                          // 链 ID: 56=BSC 主网 97=BSC 测试网
	ContractAddress  string          `gorm:"column:contract_address;type:varchar(255);not null" json:"contractAddress"`           // 合约地址
	ContractAbi      *datatypes.JSON `gorm:"column:contract_abi;type:json;not null" json:"contractAbi"`                           // 合约ABI
	ContractBin      *datatypes.JSON `gorm:"column:contract_bin;type:json;not null" json:"contractBin"`                           // 合约BIN
	PublisherAddress string          `gorm:"column:publisher_address;type:varchar(255);not null" json:"publisherAddress"`         // 合约发布者地址
	CreatedAt        *time.Time      `gorm:"column:created_at;type:datetime;default:CURRENT_TIMESTAMP;not null" json:"createdAt"` // 创建时间
	UpdatedAt        *time.Time      `gorm:"column:updated_at;type:datetime;default:CURRENT_TIMESTAMP;not null" json:"updatedAt"` // 更新时间
}

// TableName table name
func (m *Contract) TableName() string {
	return "contract"
}

// ContractColumnNames Whitelist for custom query fields to prevent sql injection attacks
var ContractColumnNames = map[string]bool{
	"id":                true,
	"node_url":          true,
	"chain_id":          true,
	"contract_address":  true,
	"contract_abi":      true,
	"contract_bin":      true,
	"publisher_address": true,
	"created_at":        true,
	"updated_at":        true,
}
