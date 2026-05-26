package types

import (
	"time"

	"github.com/go-dev-frame/sponge/pkg/sgorm/query"
)

var _ time.Time

// Tip: suggested filling in the binding rules https://github.com/go-playground/validator in request struct fields tag.

// CreateContractRequest request params
type CreateContractRequest struct {
	NodeURL          string `json:"nodeURL" binding:""`          // 合约发布网站地址，如127.0.0.1:8454
	ChainID          string `json:"chainID" binding:""`          // 链 ID: 56=BSC 主网 97=BSC 测试网
	ContractAddress  string `json:"contractAddress" binding:""`  // 合约地址
	ContractAbi      string `json:"contractAbi" binding:""`      // 合约ABI
	ContractBin      string `json:"contractBin" binding:""`      // 合约BIN
	PublisherAddress string `json:"publisherAddress" binding:""` // 合约发布者地址
}

// UpdateContractByIDRequest request params
type UpdateContractByIDRequest struct {
	ID uint64 `json:"id" binding:""` // uint64 id
	// 主键ID
	NodeURL          string `json:"nodeURL" binding:""`          // 合约发布网站地址，如127.0.0.1:8454
	ChainID          string `json:"chainID" binding:""`          // 链 ID: 56=BSC 主网 97=BSC 测试网
	ContractAddress  string `json:"contractAddress" binding:""`  // 合约地址
	ContractAbi      string `json:"contractAbi" binding:""`      // 合约ABI
	ContractBin      string `json:"contractBin" binding:""`      // 合约BIN
	PublisherAddress string `json:"publisherAddress" binding:""` // 合约发布者地址
}

// ContractObjDetail detail
type ContractObjDetail struct {
	ID uint64 `json:"id"` // convert to uint64 id
	// 主键ID
	NodeURL          string     `json:"nodeURL"`          // 合约发布网站地址，如127.0.0.1:8454
	ChainID          string     `json:"chainID"`          // 链 ID: 56=BSC 主网 97=BSC 测试网
	ContractAddress  string     `json:"contractAddress"`  // 合约地址
	ContractAbi      string     `json:"contractAbi"`      // 合约ABI
	ContractBin      string     `json:"contractBin"`      // 合约BIN
	PublisherAddress string     `json:"publisherAddress"` // 合约发布者地址
	CreatedAt        *time.Time `json:"createdAt"`        // 创建时间
	UpdatedAt        *time.Time `json:"updatedAt"`        // 更新时间
}

// CreateContractReply only for api docs
type CreateContractReply struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		ID uint64 `json:"id"` // id
	} `json:"data"` // return data
}

// DeleteContractByIDReply only for api docs
type DeleteContractByIDReply struct {
	Code int      `json:"code"` // return code
	Msg  string   `json:"msg"`  // return information description
	Data struct{} `json:"data"` // return data
}

// UpdateContractByIDReply only for api docs
type UpdateContractByIDReply struct {
	Code int      `json:"code"` // return code
	Msg  string   `json:"msg"`  // return information description
	Data struct{} `json:"data"` // return data
}

// GetContractByIDReply only for api docs
type GetContractByIDReply struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		Contract ContractObjDetail `json:"contract"`
	} `json:"data"` // return data
}

// ListContractsRequest request params
type ListContractsRequest struct {
	query.Params
}

// ListContractsReply only for api docs
type ListContractsReply struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		Contracts []ContractObjDetail `json:"contracts"`
	} `json:"data"` // return data
}
