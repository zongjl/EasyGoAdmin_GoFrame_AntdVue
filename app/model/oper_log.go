// ==========================================================================
// This is auto-generated by gf cli tool. Fill this file as you wish.
// ==========================================================================

package model

import (
	"easygoadmin/app/model/internal"
)

// OperLog is the golang structure for table sys_oper_log.
type OperLog internal.OperLog

// Fill with you ideas below.

// 分页查询条件
type OperLogPageReq struct {
	Username string `p:"username"` // 操作账号
	Model    string `p:"model"`    // 操作模块
	Page     int    `p:"page"`     // 页码
	Limit    int    `p:"limit"`    // 每页数
}
