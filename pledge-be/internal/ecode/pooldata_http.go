package ecode

import (
	"github.com/go-dev-frame/sponge/pkg/errcode"
)

// pooldata business-level http error codes.
// the pooldataNO value range is 1~999, if the same error code is used, it will cause panic.
var (
	pooldataNO       = 88
	pooldataName     = "pooldata"
	pooldataBaseCode = errcode.HCode(pooldataNO)

	ErrCreatePooldata     = errcode.NewError(pooldataBaseCode+1, "failed to create "+pooldataName)
	ErrDeleteByIDPooldata = errcode.NewError(pooldataBaseCode+2, "failed to delete "+pooldataName)
	ErrUpdateByIDPooldata = errcode.NewError(pooldataBaseCode+3, "failed to update "+pooldataName)
	ErrGetByIDPooldata    = errcode.NewError(pooldataBaseCode+4, "failed to get "+pooldataName+" details")
	ErrListPooldata       = errcode.NewError(pooldataBaseCode+5, "failed to list of "+pooldataName)

	// error codes are globally unique, adding 1 to the previous error code
)
