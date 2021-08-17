// ==========================================================================
// This is auto-generated by gf cli tool. Fill this file as you wish.
// ==========================================================================

package model

import (
	"easygoadmin/app/model/internal"
)

// Item is the golang structure for table sys_item.
type Item internal.Item

// Fill with you ideas below.

// 分页查询条件
type ItemPageReq struct {
	Name  string `p:"name"` // 站点名称
	Type  int    `p:type`   // 站点类型
	Page  int    `p:page`   // 页码
	Limit int    `p:limit`  // 每页数
}

// 站点信息Vo
type ItemInfoVo struct {
	Item
	TypeName string `json:"typeName"` // 站点类型
}

// 添加站点
type ItemAddReq struct {
	Name   string `p:"name"        v:"required#站点名称不能为空"`  // 站点名称
	Type   int    `p:"type"        v:"required#请选择站点类型"`   // 站点类型:1普通站点 2其他
	Url    string `p:"url"         v:"required#站点地址不能为空"`  // 站点地址
	Image  string `p:"image"`                              // 站点图片
	Status int    `p:"status"      v:"required#请选择站点状态"`   // 状态：1在用 2停用
	Note   string `p:"note"`                               // 站点备注
	Sort   int    `p:"sort"        v:"required#站点排序号不能为空"` // 显示顺序
}

// 更新站点
type ItemUpdateReq struct {
	Id     int    `p:"id" v:"required#主键ID不能为空"`
	Name   string `p:"name"        v:"required#站点名称不能为空"`  // 站点名称
	Type   int    `p:"type"        v:"required#请选择站点类型"`   // 站点类型:1普通站点 2其他
	Url    string `p:"url"         v:"required#站点地址不能为空"`  // 站点地址
	Image  string `p:"image"`                              // 站点图片
	Status int    `p:"status"      v:"required#请选择站点状态"`   // 状态：1在用 2停用
	Note   string `p:"note"`                               // 站点备注
	Sort   int    `p:"sort"        v:"required#站点排序号不能为空"` // 显示顺序
}

// 删除站点
type ItemDeleteReq struct {
	Ids string `p:"ids"  v:"required#请选择要删除的数据记录"`
}

// 设置状态
type ItemStatusReq struct {
	Id     int `p:"id" v:"required#主键ID不能为空"`
	Status int `p:"status"    v:"required#状态不能为空"`
}
