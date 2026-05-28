// Package ecode 统一管理 HTTP 和 gRPC 错误码定义。
package ecode

import (
	"github.com/go-dev-frame/sponge/pkg/errcode"
)

// HTTP 系统级错误码，错误码范围 10000~20000
var (
	Success = errcode.Success

	InvalidParams       = errcode.InvalidParams
	Unauthorized        = errcode.Unauthorized
	InternalServerError = errcode.InternalServerError
	NotFound            = errcode.NotFound
	Timeout             = errcode.Timeout
	TooManyRequests     = errcode.TooManyRequests
	Forbidden           = errcode.Forbidden
	LimitExceed         = errcode.LimitExceed
	Conflict            = errcode.Conflict
	TooEarly            = errcode.TooEarly

	DeadlineExceeded   = errcode.DeadlineExceeded
	AccessDenied       = errcode.AccessDenied
	MethodNotAllowed   = errcode.MethodNotAllowed
	ServiceUnavailable = errcode.ServiceUnavailable

	Canceled           = errcode.Canceled
	Unknown            = errcode.Unknown
	PermissionDenied   = errcode.PermissionDenied
	ResourceExhausted  = errcode.ResourceExhausted
	FailedPrecondition = errcode.FailedPrecondition
	Aborted            = errcode.Aborted
	OutOfRange         = errcode.OutOfRange
	Unimplemented      = errcode.Unimplemented
	DataLoss           = errcode.DataLoss
)

// SkipResponse 跳过响应包装的错误码，用于无需统一响应格式的场景
var SkipResponse = errcode.SkipResponse

// GetErrorCode 从错误对象中提取错误码
var GetErrorCode = errcode.GetErrorCode
