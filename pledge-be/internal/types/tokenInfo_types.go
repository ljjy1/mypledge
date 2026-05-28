package types

import (
	"time"

	"github.com/go-dev-frame/sponge/pkg/sgorm/query"
)

var _ time.Time

// Tip: suggested filling in the binding rules https://github.com/go-playground/validator in request struct fields tag.

// CreateTokenInfoRequest 创建代币信息请求参数
type CreateTokenInfoRequest struct {
	Symbol   string `json:"symbol" binding:""`   // 代币符号，如 BUSD/BTC
	Logo     string `json:"logo" binding:""`     // 代币 logo URL
	Price    string `json:"price" binding:""`    // 价格(精度值，用于估值与清算)
	Token    string `json:"token" binding:""`    // 代币合约地址
	ChainID  string `json:"chainID" binding:""`  // 链 ID: 56=BSC 97=测试网
	Decimals int    `json:"decimals" binding:""` // 代币精度(小数位数)
}

// UpdateTokenInfoByIDRequest 更新代币信息请求参数
type UpdateTokenInfoByIDRequest struct {
	ID uint64 `json:"id" binding:""` // uint64 id
	// 主键ID
	Symbol   string `json:"symbol" binding:""`   // 代币符号，如 BUSD/BTC
	Logo     string `json:"logo" binding:""`     // 代币 logo URL
	Price    string `json:"price" binding:""`    // 价格(精度值，用于估值与清算)
	Token    string `json:"token" binding:""`    // 代币合约地址
	ChainID  string `json:"chainID" binding:""`  // 链 ID: 56=BSC 97=测试网
	Decimals int    `json:"decimals" binding:""` // 代币精度(小数位数)
}

// TokenInfoObjDetail 代币信息详情
type TokenInfoObjDetail struct {
	ID uint64 `json:"id"` // convert to uint64 id
	// 主键ID
	Symbol    string     `json:"symbol"`    // 代币符号，如 BUSD/BTC
	Logo      string     `json:"logo"`      // 代币 logo URL
	Price     string     `json:"price"`     // 价格(精度值，用于估值与清算)
	Token     string     `json:"token"`     // 代币合约地址
	ChainID   string     `json:"chainID"`   // 链 ID: 56=BSC 97=测试网
	Decimals  int        `json:"decimals"`  // 代币精度(小数位数)
	CreatedAt *time.Time `json:"createdAt"` // 创建时间
	UpdatedAt *time.Time `json:"updatedAt"` // 更新时间
}

// CreateTokenInfoReply 仅用于 API 文档
type CreateTokenInfoReply struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		ID uint64 `json:"id"` // id
	} `json:"data"` // return data
}

// DeleteTokenInfoByIDReply 仅用于 API 文档
type DeleteTokenInfoByIDReply struct {
	Code int      `json:"code"` // return code
	Msg  string   `json:"msg"`  // return information description
	Data struct{} `json:"data"` // return data
}

// UpdateTokenInfoByIDReply 仅用于 API 文档
type UpdateTokenInfoByIDReply struct {
	Code int      `json:"code"` // return code
	Msg  string   `json:"msg"`  // return information description
	Data struct{} `json:"data"` // return data
}

// GetTokenInfoByIDReply 仅用于 API 文档
type GetTokenInfoByIDReply struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		TokenInfo TokenInfoObjDetail `json:"tokenInfo"`
	} `json:"data"` // return data
}

// ListTokenInfosRequest 查询代币信息列表请求参数
type ListTokenInfosRequest struct {
	query.Params
}

// ListTokenInfosReply 仅用于 API 文档
type ListTokenInfosReply struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		TokenInfos []TokenInfoObjDetail `json:"tokenInfos"`
	} `json:"data"` // return data
}
