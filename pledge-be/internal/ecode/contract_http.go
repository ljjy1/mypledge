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

	ErrDeployRPCConn     = errcode.NewError(contractBaseCode+6, "连接RPC节点失败")
	ErrDeployTxSign      = errcode.NewError(contractBaseCode+7, "交易签名失败")
	ErrDeploySend        = errcode.NewError(contractBaseCode+8, "发送部署交易失败")
	ErrDeployWaitReceipt = errcode.NewError(contractBaseCode+9, "等待交易确认超时")
	ErrDeployInvalidAddr = errcode.NewError(contractBaseCode+10, "无效的地址参数")

	ErrCreatePoolSend    = errcode.NewError(contractBaseCode+11, "发送创建池子交易失败")
	ErrCreatePoolReceipt = errcode.NewError(contractBaseCode+12, "等待创建池子交易确认失败")
	ErrCreatePoolEvent   = errcode.NewError(contractBaseCode+13, "解析创建池子事件失败")
	ErrCreatePoolSaveDB  = errcode.NewError(contractBaseCode+14, "保存池子信息到数据库失败")

	// PledgePool operations
	ErrPoolRPCConnect    = errcode.NewError(contractBaseCode+15, "连接RPC节点失败")
	ErrPoolTxSign        = errcode.NewError(contractBaseCode+16, "交易签名失败")
	ErrPoolSendTx        = errcode.NewError(contractBaseCode+17, "发送交易失败")
	ErrPoolWaitReceipt   = errcode.NewError(contractBaseCode+18, "等待交易确认失败")
	ErrPoolInvalidAddr   = errcode.NewError(contractBaseCode+19, "无效的合约地址")
	ErrPoolReadCall      = errcode.NewError(contractBaseCode+20, "链上查询失败")
	ErrPoolInvalidPoolID = errcode.NewError(contractBaseCode+21, "无效的池子ID")
	ErrPoolInvalidAmount = errcode.NewError(contractBaseCode+22, "无效的金额参数")

	// Oracle operations
	ErrOracleRPCConnect    = errcode.NewError(contractBaseCode+23, "连接Oracle RPC节点失败")
	ErrOracleSendTx        = errcode.NewError(contractBaseCode+24, "发送Oracle交易失败")
	ErrOracleReadCall      = errcode.NewError(contractBaseCode+25, "Oracle链上查询失败")
	ErrOracleInvalidParams = errcode.NewError(contractBaseCode+26, "Oracle无效的参数")

	// Token operations
	ErrTokenRPCConnect = errcode.NewError(contractBaseCode+27, "连接Token RPC节点失败")
	ErrTokenSendTx     = errcode.NewError(contractBaseCode+28, "发送Token交易失败")
	ErrTokenReadCall   = errcode.NewError(contractBaseCode+29, "Token链上查询失败")

	// Swap operations
	ErrSwapRPCConnect    = errcode.NewError(contractBaseCode+30, "连接Swap RPC节点失败")
	ErrSwapSendTx        = errcode.NewError(contractBaseCode+31, "发送Swap交易失败")
	ErrSwapReadCall      = errcode.NewError(contractBaseCode+32, "Swap链上查询失败")
	ErrSwapInvalidParams = errcode.NewError(contractBaseCode+33, "Swap无效的参数")

	// error codes are globally unique, adding 1 to the previous error code
)
