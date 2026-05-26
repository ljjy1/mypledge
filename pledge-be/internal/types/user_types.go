package types

import (
	"time"

	"github.com/go-dev-frame/sponge/pkg/sgorm/query"
)

var _ time.Time

// Tip: suggested filling in the binding rules https://github.com/go-playground/validator in request struct fields tag.

// CreateUserRequest request params
type CreateUserRequest struct {
	Login    string `json:"login" binding:""`    // 用户账号
	Nike     string `json:"nike" binding:""`     // 用户昵称
	Password string `json:"password" binding:""` // 加密后的密码
}

// UpdateUserByIDRequest request params
type UpdateUserByIDRequest struct {
	ID uint64 `json:"id" binding:""` // uint64 id

	Login    string `json:"login" binding:""`    // 用户账号
	Nike     string `json:"nike" binding:""`     // 用户昵称
	Password string `json:"password" binding:""` // 加密后的密码
}

// UserObjDetail detail
type UserObjDetail struct {
	ID uint64 `json:"id"` // convert to uint64 id

	Login    string `json:"login"`    // 用户账号
	Nike     string `json:"nike"`     // 用户昵称
	Password string `json:"password"` // 加密后的密码
}

// CreateUserReply only for api docs
type CreateUserReply struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		ID uint64 `json:"id"` // id
	} `json:"data"` // return data
}

// DeleteUserByIDReply only for api docs
type DeleteUserByIDReply struct {
	Code int      `json:"code"` // return code
	Msg  string   `json:"msg"`  // return information description
	Data struct{} `json:"data"` // return data
}

// UpdateUserByIDReply only for api docs
type UpdateUserByIDReply struct {
	Code int      `json:"code"` // return code
	Msg  string   `json:"msg"`  // return information description
	Data struct{} `json:"data"` // return data
}

// GetUserByIDReply only for api docs
type GetUserByIDReply struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		User UserObjDetail `json:"user"`
	} `json:"data"` // return data
}

// ListUsersRequest request params
type ListUsersRequest struct {
	query.Params
}

// RegisterRequest request params
type RegisterRequest struct {
	Login    string `json:"login" binding:"required"`    // 用户账号 (3-10位字母或数字)
	Password string `json:"password" binding:"required"` // 密码 (8-20位)
	Nike     string `json:"nike" binding:"required"`     // 用户昵称 (1-12位)
}

// RegisterReply only for api docs
type RegisterReply struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		ID uint64 `json:"id"` // 用户ID
	} `json:"data"` // return data
}

// LoginRequest request params
type LoginRequest struct {
	Login    string `json:"login" binding:"required"`    // 用户账号
	Password string `json:"password" binding:"required"` // 密码
}

// LoginReply only for api docs
type LoginReply struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		Token string `json:"token"` // jwt token
	} `json:"data"` // return data
}

// ListUsersReply only for api docs
type ListUsersReply struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		Users []UserObjDetail `json:"users"`
	} `json:"data"` // return data
}
