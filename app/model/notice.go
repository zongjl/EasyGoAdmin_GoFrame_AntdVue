// ==========================================================================
// This is auto-generated by gf cli tool. Fill this file as you wish.
// ==========================================================================

package model

import (
	"easygoadmin/app/model/internal"
)

// Notice is the golang structure for table sys_notice.
type Notice internal.Notice

// Fill with you ideas below.

// 分页查询
type NoticePageReq struct {
	Title  string `p:"title"`  // 通知标题
	Source int    `p:"source"` // 通知来源
	Page   int    `p:"page"`   // 页码
	Limit  int    `p:"limit"`  // 每页数
}

// 添加通知公告
type NoticeAddReq struct {
	Title   string `p:"title"       v:"required#通知标题不能为空"` // 通知标题
	Content string `p:"content"     v:"required#通知内容不能为空"` // 通知内容
	Source  int    `p:"source"      v:"required#通知来源不能为空"` // 来源：1内部通知 2外部新闻
	IsTop   int    `p:"is_top"      v:"required#请选择是否置顶"`  // 是否置顶：1是 2否
	Status  int    `p:"status"      v:"required#请选择通知状态"`  // 状态：1已发布 2待发布
}

// 更新通知公告
type NoticeUpdateReq struct {
	Id      int    `p:"id" 			v:"required#主键ID不能为空"`
	Title   string `p:"title"       v:"required#通知标题不能为空"` // 通知标题
	Content string `p:"content"     v:"required#通知内容不能为空"` // 通知内容
	Source  int    `p:"source"      v:"required#通知来源不能为空"` // 来源：1内部通知 2外部新闻
	IsTop   int    `p:"is_top"      v:"required#请选择是否置顶"`  // 是否置顶：1是 2否
	Status  int    `p:"status"      v:"required#请选择通知状态"`  // 状态：1已发布 2待发布
}

// 删除通知公告
type NoticeDeleteReq struct {
	Ids string `p:"ids"  v:"required#请选择要删除的数据记录"`
}

// 设置状态
type NoticeStatusReq struct {
	Id     int `p:"id" v:"required#主键ID不能为空"`
	Status int `p:"status"    v:"required#状态不能为空"`
}

// 通知公告Vo
type NoticeInfoVo struct {
	Notice
	SourceName string `json:"sourceName"` // 通知来源
}
