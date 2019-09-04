package admin

import (
	"beego_admin/models/admin"
	"beego_admin/utils"
)

type HomepageControl struct {
	baseController
}

// 首页模块列表
func (c *HomepageControl)ModuleList()  {
	homepage := admin.Homepage{}
	modules := homepage.ModuleList()
	c.Data["Homepages"] = modules
	c.Data["Title"] = "首页模块设置"
	// 模板
	c.TplName = "admin/homepage/index.html"
}

// 首页模块设置
func (c *HomepageControl)ModulePatch()  {
	returnJson:= ResponseJson{}
	homepage := admin.Homepage{}
	homepage.ID = utils.MustInt(c.Input().Get("id"))
	homepage.IsShow = utils.MustInt(c.Input().Get("is_show"))
	err := homepage.ModuleUpdate()
	if err == nil {
		returnJson.StatusCode = Success
		returnJson.Message = SaveSuccess
	} else {
		returnJson.StatusCode = Fail
		returnJson.Message = SaveFail
	}
	returnJson.UrlType = Reload
	c.Data["json"] = &returnJson
	c.ServeJSON()
}
