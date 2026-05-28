package types

import (
	"time"

	"github.com/go-dev-frame/sponge/pkg/sgorm/query"
)

var _ time.Time

// Tip: suggested filling in the binding rules https://github.com/go-playground/validator in request struct fields tag.

// CreateUserRequest 创建用户请求参数
type CreateUserRequest struct {
	Login    string `json:"login" binding:""`    // 用户账号
	Nike     string `json:"nike" binding:""`     // 用户昵称
	Password string `json:"password" binding:""` // 加密后的密码
}

// UpdateUserByIDRequest 更新用户请求参数
type UpdateUserByIDRequest struct {
	ID uint64 `json:"id" binding:""` // uint64 id

	Login    string `json:"login" binding:""`    // 用户账号
	Nike     string `json:"nike" binding:""`     // 用户昵称
	Password string `json:"password" binding:""` // 加密后的密码
}

// UserObjDetail 用户详情
type UserObjDetail struct {
	ID uint64 `json:"id"` // convert to uint64 id

	Login    string `json:"login"`    // 用户账号
	Nike     string `json:"nike"`     // 用户昵称
	Password string `json:"password"` // 加密后的密码
}

// CreateUserReply 仅用于 API 文档
type CreateUserReply struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		ID uint64 `json:"id"` // id
	} `json:"data"` // return data
}

// DeleteUserByIDReply 仅用于 API 文档
type DeleteUserByIDReply struct {
	Code int      `json:"code"` // return code
	Msg  string   `json:"msg"`  // return information description
	Data struct{} `json:"data"` // return data
}

// UpdateUserByIDReply 仅用于 API 文档
type UpdateUserByIDReply struct {
	Code int      `json:"code"` // return code
	Msg  string   `json:"msg"`  // return information description
	Data struct{} `json:"data"` // return data
}

// GetUserByIDReply 仅用于 API 文档
type GetUserByIDReply struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		User UserObjDetail `json:"user"`
	} `json:"data"` // return data
}

// ListUsersRequest 查询用户列表请求参数
type ListUsersRequest struct {
	query.Params
}

// RegisterRequest 用户注册请求参数
type RegisterRequest struct {
	Login    string `json:"login" binding:"required"`    // 用户账号 (3-10位字母或数字)
	Password string `json:"password" binding:"required"` // 密码 (8-20位)
	Nike     string `json:"nike" binding:"required"`     // 用户昵称 (1-12位)
}

// RegisterReply 仅用于 API 文档
type RegisterReply struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		ID uint64 `json:"id"` // 用户ID
	} `json:"data"` // return data
}

// LoginRequest 用户登录请求参数
type LoginRequest struct {
	Login    string `json:"login" binding:"required"`    // 用户账号
	Password string `json:"password" binding:"required"` // 密码
}

// LoginReply 仅用于 API 文档
type LoginReply struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		Token string `json:"token"` // jwt token
	} `json:"data"` // return data
}

// ListUsersReply 仅用于 API 文档
type ListUsersReply struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		Users []UserObjDetail `json:"users"`
	} `json:"data"` // return data
}
