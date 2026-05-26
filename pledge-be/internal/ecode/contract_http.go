package ecode

import (
	"github.com/go-dev-frame/sponge/pkg/errcode"
)

// contract business-level http error codes.
// the contractNO value range is 1~999, if the same error code is used, it will cause panic.
var (
	contractNO       = 27
	contractName     = "contract"
	contractBaseCode = errcode.HCode(contractNO)

	ErrCreateContract     = errcode.NewError(contractBaseCode+1, "failed to create "+contractName)
	ErrDeleteByIDContract = errcode.NewError(contractBaseCode+2, "failed to delete "+contractName)
	ErrUpdateByIDContract = errcode.NewError(contractBaseCode+3, "failed to update "+contractName)
	ErrGetByIDContract    = errcode.NewError(contractBaseCode+4, "failed to get "+contractName+" details")
	ErrListContract       = errcode.NewError(contractBaseCode+5, "failed to list of "+contractName)

	// error codes are globally unique, adding 1 to the previous error code
)
