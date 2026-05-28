package ecode

import (
	"github.com/go-dev-frame/sponge/pkg/errcode"
)

// user 业务层 HTTP 错误码定义。
// userNO 的取值范围为 1~999，若使用了重复的错误码将触发 panic。
var (
	userNO       = 68
	userName     = "user"
	userBaseCode = errcode.HCode(userNO)

	ErrCreateUser     = errcode.NewError(userBaseCode+1, "failed to create "+userName)
	ErrDeleteByIDUser = errcode.NewError(userBaseCode+2, "failed to delete "+userName)
	ErrUpdateByIDUser = errcode.NewError(userBaseCode+3, "failed to update "+userName)
	ErrGetByIDUser    = errcode.NewError(userBaseCode+4, "failed to get "+userName+" details")
	ErrListUser       = errcode.NewError(userBaseCode+5, "failed to list of "+userName)
	// --- 登录/注册相关错误码 ---
	ErrLogin               = errcode.NewError(userBaseCode+6, "invalid login credentials")
	ErrRegisterLoginExists = errcode.NewError(userBaseCode+7, "login already exists")
	ErrRegisterLoginFormat = errcode.NewError(userBaseCode+8, "登录账号必须为3-10位字母或数字")
	ErrRegisterPwdFormat   = errcode.NewError(userBaseCode+9, "密码必须为8-20位")
	ErrRegisterNikeFormat  = errcode.NewError(userBaseCode+10, "昵称必须为1-12位")

	// 错误码全局唯一，请在上一个错误码基础上加 1
)
