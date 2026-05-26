package types

import (
	"time"

	"github.com/go-dev-frame/sponge/pkg/sgorm/query"
)

var _ time.Time

// Tip: suggested filling in the binding rules https://github.com/go-playground/validator in request struct fields tag.

// CreateTokenInfoRequest request params
type CreateTokenInfoRequest struct {
	Symbol      string `json:"symbol" binding:""`      // 代币符号，如 BUSD/BTC
	Logo        string `json:"logo" binding:""`        // 代币 logo URL
	Price       string `json:"price" binding:""`       // 价格(精度值，用于估值与清算)
	Token       string `json:"token" binding:""`       // 代币合约地址
	ChainID     string `json:"chainID" binding:""`     // 链 ID: 56=BSC 97=测试网
	ContractAbi string `json:"contractAbi" binding:""` // 代币ABI
	ContractBin string `json:"contractBin" binding:""` // 代币BIN
	Decimals    int    `json:"decimals" binding:""`    // 代币精度(小数位数)
}

// UpdateTokenInfoByIDRequest request params
type UpdateTokenInfoByIDRequest struct {
	ID uint64 `json:"id" binding:""` // uint64 id
	// 主键ID
	Symbol      string `json:"symbol" binding:""`      // 代币符号，如 BUSD/BTC
	Logo        string `json:"logo" binding:""`        // 代币 logo URL
	Price       string `json:"price" binding:""`       // 价格(精度值，用于估值与清算)
	Token       string `json:"token" binding:""`       // 代币合约地址
	ChainID     string `json:"chainID" binding:""`     // 链 ID: 56=BSC 97=测试网
	ContractAbi string `json:"contractAbi" binding:""` // 代币ABI
	ContractBin string `json:"contractBin" binding:""` // 代币BIN
	Decimals    int    `json:"decimals" binding:""`    // 代币精度(小数位数)
}

// TokenInfoObjDetail detail
type TokenInfoObjDetail struct {
	ID uint64 `json:"id"` // convert to uint64 id
	// 主键ID
	Symbol      string     `json:"symbol"`      // 代币符号，如 BUSD/BTC
	Logo        string     `json:"logo"`        // 代币 logo URL
	Price       string     `json:"price"`       // 价格(精度值，用于估值与清算)
	Token       string     `json:"token"`       // 代币合约地址
	ChainID     string     `json:"chainID"`     // 链 ID: 56=BSC 97=测试网
	ContractAbi string     `json:"contractAbi"` // 代币ABI
	ContractBin string     `json:"contractBin"` // 代币BIN
	Decimals    int        `json:"decimals"`    // 代币精度(小数位数)
	CreatedAt   *time.Time `json:"createdAt"`   // 创建时间
	UpdatedAt   *time.Time `json:"updatedAt"`   // 更新时间
}

// CreateTokenInfoReply only for api docs
type CreateTokenInfoReply struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		ID uint64 `json:"id"` // id
	} `json:"data"` // return data
}

// DeleteTokenInfoByIDReply only for api docs
type DeleteTokenInfoByIDReply struct {
	Code int      `json:"code"` // return code
	Msg  string   `json:"msg"`  // return information description
	Data struct{} `json:"data"` // return data
}

// UpdateTokenInfoByIDReply only for api docs
type UpdateTokenInfoByIDReply struct {
	Code int      `json:"code"` // return code
	Msg  string   `json:"msg"`  // return information description
	Data struct{} `json:"data"` // return data
}

// GetTokenInfoByIDReply only for api docs
type GetTokenInfoByIDReply struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		TokenInfo TokenInfoObjDetail `json:"tokenInfo"`
	} `json:"data"` // return data
}

// ListTokenInfosRequest request params
type ListTokenInfosRequest struct {
	query.Params
}

// ListTokenInfosReply only for api docs
type ListTokenInfosReply struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		TokenInfos []TokenInfoObjDetail `json:"tokenInfos"`
	} `json:"data"` // return data
}
