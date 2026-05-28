package ecode

import (
	"github.com/go-dev-frame/sponge/pkg/errcode"
)

// pooldata 业务层 HTTP 错误码定义。
// pooldataNO 的取值范围为 1~999，若使用了重复的错误码将触发 panic。
var (
	pooldataNO       = 88
	pooldataName     = "pooldata"
	pooldataBaseCode = errcode.HCode(pooldataNO)

	ErrCreatePooldata     = errcode.NewError(pooldataBaseCode+1, "failed to create "+pooldataName)
	ErrDeleteByIDPooldata = errcode.NewError(pooldataBaseCode+2, "failed to delete "+pooldataName)
	ErrUpdateByIDPooldata = errcode.NewError(pooldataBaseCode+3, "failed to update "+pooldataName)
	ErrGetByIDPooldata    = errcode.NewError(pooldataBaseCode+4, "failed to get "+pooldataName+" details")
	ErrListPooldata       = errcode.NewError(pooldataBaseCode+5, "failed to list of "+pooldataName)

	// 错误码全局唯一，请在上一个错误码基础上加 1
)
