package ecode

import (
	"github.com/go-dev-frame/sponge/pkg/errcode"
)

// poolbases 业务层 HTTP 错误码定义。
// poolbasesNO 的取值范围为 1~999，若使用了重复的错误码将触发 panic。
var (
	poolbasesNO       = 4
	poolbasesName     = "poolbases"
	poolbasesBaseCode = errcode.HCode(poolbasesNO)

	ErrCreatePoolbases     = errcode.NewError(poolbasesBaseCode+1, "failed to create "+poolbasesName)
	ErrDeleteByIDPoolbases = errcode.NewError(poolbasesBaseCode+2, "failed to delete "+poolbasesName)
	ErrUpdateByIDPoolbases = errcode.NewError(poolbasesBaseCode+3, "failed to update "+poolbasesName)
	ErrGetByIDPoolbases    = errcode.NewError(poolbasesBaseCode+4, "failed to get "+poolbasesName+" details")
	ErrListPoolbases       = errcode.NewError(poolbasesBaseCode+5, "failed to list of "+poolbasesName)

	// 错误码全局唯一，请在上一个错误码基础上加 1
)
