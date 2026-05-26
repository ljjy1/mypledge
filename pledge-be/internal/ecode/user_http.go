package ecode

import (
	"github.com/go-dev-frame/sponge/pkg/errcode"
)

// user business-level http error codes.
// the userNO value range is 1~999, if the same error code is used, it will cause panic.
var (
	userNO       = 68
	userName     = "user"
	userBaseCode = errcode.HCode(userNO)

	ErrCreateUser          = errcode.NewError(userBaseCode+1, "failed to create "+userName)
	ErrDeleteByIDUser      = errcode.NewError(userBaseCode+2, "failed to delete "+userName)
	ErrUpdateByIDUser      = errcode.NewError(userBaseCode+3, "failed to update "+userName)
	ErrGetByIDUser         = errcode.NewError(userBaseCode+4, "failed to get "+userName+" details")
	ErrListUser            = errcode.NewError(userBaseCode+5, "failed to list of "+userName)
	ErrLogin               = errcode.NewError(userBaseCode+6, "invalid login credentials")
	ErrRegisterLoginExists = errcode.NewError(userBaseCode+7, "login already exists")
	ErrRegisterLoginFormat = errcode.NewError(userBaseCode+8, "登录账号必须为3-10位字母或数字")
	ErrRegisterPwdFormat   = errcode.NewError(userBaseCode+9, "密码必须为8-20位")
	ErrRegisterNikeFormat  = errcode.NewError(userBaseCode+10, "昵称必须为1-12位")

	// error codes are globally unique, adding 1 to the previous error code
)
