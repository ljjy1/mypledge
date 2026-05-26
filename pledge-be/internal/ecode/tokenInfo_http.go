package ecode

import (
	"github.com/go-dev-frame/sponge/pkg/errcode"
)

// tokenInfo business-level http error codes.
// the tokenInfoNO value range is 1~999, if the same error code is used, it will cause panic.
var (
	tokenInfoNO       = 80
	tokenInfoName     = "tokenInfo"
	tokenInfoBaseCode = errcode.HCode(tokenInfoNO)

	ErrCreateTokenInfo     = errcode.NewError(tokenInfoBaseCode+1, "failed to create "+tokenInfoName)
	ErrDeleteByIDTokenInfo = errcode.NewError(tokenInfoBaseCode+2, "failed to delete "+tokenInfoName)
	ErrUpdateByIDTokenInfo = errcode.NewError(tokenInfoBaseCode+3, "failed to update "+tokenInfoName)
	ErrGetByIDTokenInfo    = errcode.NewError(tokenInfoBaseCode+4, "failed to get "+tokenInfoName+" details")
	ErrListTokenInfo       = errcode.NewError(tokenInfoBaseCode+5, "failed to list of "+tokenInfoName)

	// error codes are globally unique, adding 1 to the previous error code
)
