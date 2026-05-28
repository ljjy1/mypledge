package types

import (
	"time"

	"github.com/go-dev-frame/sponge/pkg/sgorm/query"
)

var _ time.Time

// Tip: suggested filling in the binding rules https://github.com/go-playground/validator in request struct fields tag.

// CreatePooldataRequest 创建池子动态数据请求参数
type CreatePooldataRequest struct {
	ChainID                string `json:"chainID" binding:""`                // 链 ID，与 poolbases 一致
	PoolID                 string `json:"poolID" binding:""`                 // 池子 ID，与 poolbases.pool_id 对应
	SettleAmountLend       string `json:"settleAmountLend" binding:""`       // 结算时借出侧(贷方)金额
	SettleAmountBorrow     string `json:"settleAmountBorrow" binding:""`     // 结算时借入侧(抵押)金额
	FinishAmountLend       string `json:"finishAmountLend" binding:""`       // 已完成/归还的借出侧金额
	FinishAmountBorrow     string `json:"finishAmountBorrow" binding:""`     // 已完成/归还的借入侧金额
	LiquidationAmounLend   string `json:"liquidationAmounLend" binding:""`   // 清算产生的借出侧金额
	LiquidationAmounBorrow string `json:"liquidationAmounBorrow" binding:""` // 清算产生的借入侧金额
}

// UpdatePooldataByIDRequest 更新池子动态数据请求参数
type UpdatePooldataByIDRequest struct {
	ID uint64 `json:"id" binding:""` // uint64 id
	// 主键ID
	ChainID                string `json:"chainID" binding:""`                // 链 ID，与 poolbases 一致
	PoolID                 string `json:"poolID" binding:""`                 // 池子 ID，与 poolbases.pool_id 对应
	SettleAmountLend       string `json:"settleAmountLend" binding:""`       // 结算时借出侧(贷方)金额
	SettleAmountBorrow     string `json:"settleAmountBorrow" binding:""`     // 结算时借入侧(抵押)金额
	FinishAmountLend       string `json:"finishAmountLend" binding:""`       // 已完成/归还的借出侧金额
	FinishAmountBorrow     string `json:"finishAmountBorrow" binding:""`     // 已完成/归还的借入侧金额
	LiquidationAmounLend   string `json:"liquidationAmounLend" binding:""`   // 清算产生的借出侧金额
	LiquidationAmounBorrow string `json:"liquidationAmounBorrow" binding:""` // 清算产生的借入侧金额
}

// PooldataObjDetail 池子动态数据详情
type PooldataObjDetail struct {
	ID uint64 `json:"id"` // convert to uint64 id
	// 主键ID
	ChainID                string     `json:"chainID"`                // 链 ID，与 poolbases 一致
	PoolID                 string     `json:"poolID"`                 // 池子 ID，与 poolbases.pool_id 对应
	SettleAmountLend       string     `json:"settleAmountLend"`       // 结算时借出侧(贷方)金额
	SettleAmountBorrow     string     `json:"settleAmountBorrow"`     // 结算时借入侧(抵押)金额
	FinishAmountLend       string     `json:"finishAmountLend"`       // 已完成/归还的借出侧金额
	FinishAmountBorrow     string     `json:"finishAmountBorrow"`     // 已完成/归还的借入侧金额
	LiquidationAmounLend   string     `json:"liquidationAmounLend"`   // 清算产生的借出侧金额
	LiquidationAmounBorrow string     `json:"liquidationAmounBorrow"` // 清算产生的借入侧金额
	CreatedAt              *time.Time `json:"createdAt"`              // 创建时间
	UpdatedAt              *time.Time `json:"updatedAt"`              // 更新时间
}

// CreatePooldataReply 仅用于 API 文档
type CreatePooldataReply struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		ID uint64 `json:"id"` // id
	} `json:"data"` // return data
}

// DeletePooldataByIDReply 仅用于 API 文档
type DeletePooldataByIDReply struct {
	Code int      `json:"code"` // return code
	Msg  string   `json:"msg"`  // return information description
	Data struct{} `json:"data"` // return data
}

// UpdatePooldataByIDReply 仅用于 API 文档
type UpdatePooldataByIDReply struct {
	Code int      `json:"code"` // return code
	Msg  string   `json:"msg"`  // return information description
	Data struct{} `json:"data"` // return data
}

// GetPooldataByIDReply 仅用于 API 文档
type GetPooldataByIDReply struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		Pooldata PooldataObjDetail `json:"pooldata"`
	} `json:"data"` // return data
}

// ListPooldatasRequest 查询池子动态数据列表请求参数
type ListPooldatasRequest struct {
	query.Params
}

// ListPooldatasReply 仅用于 API 文档
type ListPooldatasReply struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		Pooldatas []PooldataObjDetail `json:"pooldatas"`
	} `json:"data"` // return data
}
