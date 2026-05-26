package ecode

import (
	"github.com/go-dev-frame/sponge/pkg/errcode"
)

// poolbases business-level http error codes.
// the poolbasesNO value range is 1~999, if the same error code is used, it will cause panic.
var (
	poolbasesNO       = 4
	poolbasesName     = "poolbases"
	poolbasesBaseCode = errcode.HCode(poolbasesNO)

	ErrCreatePoolbases     = errcode.NewError(poolbasesBaseCode+1, "failed to create "+poolbasesName)
	ErrDeleteByIDPoolbases = errcode.NewError(poolbasesBaseCode+2, "failed to delete "+poolbasesName)
	ErrUpdateByIDPoolbases = errcode.NewError(poolbasesBaseCode+3, "failed to update "+poolbasesName)
	ErrGetByIDPoolbases    = errcode.NewError(poolbasesBaseCode+4, "failed to get "+poolbasesName+" details")
	ErrListPoolbases       = errcode.NewError(poolbasesBaseCode+5, "failed to list of "+poolbasesName)

	// error codes are globally unique, adding 1 to the previous error code
)
