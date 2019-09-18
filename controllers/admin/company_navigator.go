package admin

import (
	"beego_admin/models"
	"beego_admin/models/admin"
	"beego_admin/utils"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
)

type NavigatorControl struct {
	baseController
}

func (c *NavigatorControl) GetNavigatorList() {
	navigator := admin.Navigator{}
	pid := utils.MustInt(c.Input().Get("id"))
	navigator.Pid = pid
	treeList := navigator.List([]string{})
	navigators, err := json.Marshal(treeList)
	if err != nil {
		logs.Error("读取导航栏列表错误", err)
	}
	c.Data["Navigator"] = string(navigators)
	c.Data["Title"] = "导航栏列表"
	// 模板
	c.TplName = "admin/company_navigator/list.html"
}


// 删除导航栏
func (c *NavigatorControl) DeleteNavigator() {
	returnJson := ResponseJson{}
	navigator := &admin.Navigator{
		Model:models.Model{
			ID:utils.MustInt(c.Input().Get("id")),
		},
	}
	err := navigator.Delete()
	if err == nil {
		returnJson.StatusCode = Success
		returnJson.Message = DeleteSuccess
	} else {
		returnJson.StatusCode = Fail
		returnJson.Message = err.Error()
	}
	returnJson.UrlType = Reload
	c.Data["json"] = &returnJson
	c.ServeJSON()
}

// 添加导航栏
func (c *NavigatorControl) PostAddNavigator() {
	returnJson := ResponseJson{}
	navigator := &admin.Navigator{}
	navigator.Pid = utils.MustInt(c.Input().Get("pid"))
	navigator.Name = c.Input().Get("name")
	navigator.Sort = utils.MustInt(c.Input().Get("sort"))
	navigator.Url = c.Input().Get("url")
	err := navigator.Create()
	if err == nil {
		returnJson.StatusCode = Success
		returnJson.Message = AddSuccess
		returnJson.UrlType = Reload
	} else {
		returnJson.StatusCode = Fail
		returnJson.Message = err.Error()
	}
	c.Data["json"] = &returnJson
	c.ServeJSON()
}

// 更新导航栏
func (c *NavigatorControl) PutUpdateNavigator() {
	returnJson := ResponseJson{}
	navigator := &admin.Navigator{}
	navigator.Pid = utils.MustInt(c.Input().Get("pid"))
	navigator.Name = c.Input().Get("name")
	navigator.Sort = utils.MustInt(c.Input().Get("sort"))
	navigator.Url = c.Input().Get("url")
	navigator.ID = utils.MustInt(c.Input().Get("id"))
	err := navigator.Update()
	if err == nil {
		returnJson.StatusCode = Success
		returnJson.Message = SaveSuccess
		returnJson.UrlType = Reload
	} else {
		returnJson.StatusCode = Fail
		returnJson.Message = err.Error()
	}
	c.Data["json"] = &returnJson
	c.ServeJSON()
}
