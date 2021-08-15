// +----------------------------------------------------------------------
// | EasyGoAdmin敏捷开发框架 [ EasyGoAdmin ]
// +----------------------------------------------------------------------
// | 版权所有 2021 EasyGoAdmin深圳研发中心
// +----------------------------------------------------------------------
// | 官方网站: http://www.easygoadmin.vip
// +----------------------------------------------------------------------
// | Author: 半城风雨 <easygoadmin@163.com>
// +----------------------------------------------------------------------

/**
 * 菜单管理-控制器
 * @author 半城风雨
 * @since 2021/7/19
 * @File : menu
 */
package controller

import (
	"easygoadmin/app/dao"
	"easygoadmin/app/model"
	"easygoadmin/app/service"
	"easygoadmin/app/utils"
	"easygoadmin/app/utils/common"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gutil"
)

var Menu = new(menuCtl)

type menuCtl struct{}

func (c *menuCtl) List(r *ghttp.Request) {
	// 参数验证
	var req *model.MenuQueryReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJsonExit(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}

	// 调用查询方法
	list := service.Menu.GetList(req)
	// 返回结果
	r.Response.WriteJsonExit(common.JsonResult{
		Code: 0,
		Data: list,
		Msg:  "操作成功",
	})
}

func (c *menuCtl) Detail(r *ghttp.Request) {
	// 查询记录
	id := r.GetQueryInt64("id")
	if id > 0 {
		info, err := dao.Menu.FindOne("id=?", id)
		if err != nil || info == nil {
			r.Response.WriteJsonExit(common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
		}
		// 菜单信息
		menu := model.MenuInfoVo{}
		menu.Menu = *info
		// 获取权限节点
		if info.Type == 0 {
			// 获取角色菜单权限列表
			var menuList []model.Menu
			dao.Menu.Where("parent_id=?", info.Id).Where("type=1").Where("mark=1").Structs(&menuList)
			checkedList := gutil.ListItemValuesUnique(&menuList, "Sort")
			if len(checkedList) > 0 {
				menu.CheckedList = checkedList
			} else {
				menu.CheckedList = make([]interface{}, 0)
			}
		}

		// 返回结果
		r.Response.WriteJsonExit(common.JsonResult{
			Code: 0,
			Msg:  "查询成功",
			Data: menu,
		})
	}
}

func (c *menuCtl) Add(r *ghttp.Request) {
	// 参数验证
	var req *model.MenuAddReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJsonExit(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}

	// 调用添加方法
	id, err := service.Menu.Add(req, utils.Uid(r))
	if err != nil || id == 0 {
		r.Response.WriteJsonExit(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}

	// 返回成功提示
	r.Response.WriteJsonExit(common.JsonResult{
		Code: 0,
		Msg:  "添加成功",
	})
}

func (c *menuCtl) Update(r *ghttp.Request) {
	// 参数验证
	var req *model.MenuUpdateReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJsonExit(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}

	// 调用更新方法
	rows, err := service.Menu.Update(req, utils.Uid(r))
	if err != nil || rows == 0 {
		r.Response.WriteJsonExit(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}

	// 返回成功提示
	r.Response.WriteJsonExit(common.JsonResult{
		Code: 0,
		Msg:  "更新成功",
	})
}

func (c *menuCtl) Delete(r *ghttp.Request) {
	// 参数验证
	var req *model.MenuDeleteReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJsonExit(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}

	// 调用删除方法
	rows, err := service.Menu.Delete(req.Ids)
	if err != nil || rows == 0 {
		r.Response.WriteJsonExit(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}

	// 返回结果
	r.Response.WriteJsonExit(common.JsonResult{
		Code: 0,
		Msg:  "删除成功",
	})
}