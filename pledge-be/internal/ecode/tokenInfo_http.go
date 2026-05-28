package ecode

import (
	"github.com/go-dev-frame/sponge/pkg/errcode"
)

// tokenInfo 业务层 HTTP 错误码定义。
// tokenInfoNO 的取值范围为 1~999，若使用了重复的错误码将触发 panic。
var (
	tokenInfoNO       = 80
	tokenInfoName     = "tokenInfo"
	tokenInfoBaseCode = errcode.HCode(tokenInfoNO)

	ErrCreateTokenInfo     = errcode.NewError(tokenInfoBaseCode+1, "failed to create "+tokenInfoName)
	ErrDeleteByIDTokenInfo = errcode.NewError(tokenInfoBaseCode+2, "failed to delete "+tokenInfoName)
	ErrUpdateByIDTokenInfo = errcode.NewError(tokenInfoBaseCode+3, "failed to update "+tokenInfoName)
	ErrGetByIDTokenInfo    = errcode.NewError(tokenInfoBaseCode+4, "failed to get "+tokenInfoName+" details")
	ErrListTokenInfo       = errcode.NewError(tokenInfoBaseCode+5, "failed to list of "+tokenInfoName)

	// 错误码全局唯一，请在上一个错误码基础上加 1
)
