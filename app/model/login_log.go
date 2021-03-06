// ==========================================================================
// This is auto-generated by gf cli tool. Fill this file as you wish.
// ==========================================================================

package model

import (
	"easygoadmin/app/model/internal"
)

// LoginLog is the golang structure for table sys_login_log.
type LoginLog internal.LoginLog

// Fill with you ideas below.

// 分页信息查询条件
type LoginLogPageReq struct {
	Username string `p:"username"` // 操作账号
	Page     int    `p:"page"`     // 页码
	Limit    int    `p:"limit"`    // 每页数
}

// 删除登录日志
type LoginLogDeleteReq struct {
	Ids string `p:"ids"  v:"required#请选择要删除的数据记录"`
}
