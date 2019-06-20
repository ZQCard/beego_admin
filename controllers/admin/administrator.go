package admin

import (
	"beego_admin/models"
	"beego_admin/models/admin"
	"beego_admin/models/common/auth"
	"beego_admin/utils"
)

type AdministratorController struct {
	baseController
}

// 管理员列表
func (c *AdministratorController) GetAdministratorList() {
	page := utils.MustInt(c.Input().Get("page"))
	pageSize := utils.MustInt(c.Input().Get("pageSize"))
	if page == 0 {
		page = PageDefault
	}

	if pageSize == 0 {
		pageSize = PageSizeDefault
	}
	administrators, totalCount := admin.AdministratorList(page, pageSize)
	c.Data["Administrators"] = administrators
	c.Data["TotalCount"] = totalCount
	c.Data["Page"] = page
	c.Data["PageSize"] = pageSize
	c.Data["Title"] = "管理员列表"
	// 模板
	c.TplName = "admin/auth/administrator.html"
}

// 添加管理员
func (c *AdministratorController) PostAddAdministrator() {
	returnJson:= ResponseJson{}
	administrator := admin.Administrator{}
	administrator.Username = c.Input().Get("username")
	administrator.Password = c.Input().Get("password")
	administrator.Nickname = c.Input().Get("nickname")
	administrator.Email = c.Input().Get("email")
	ok, err := administrator.AddAdministrator()
	if ok == true {
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

// 更新管理员
func (c *AdministratorController) PutUpdateAdministrator() {
	returnJson := ResponseJson{}
	// 查找管理员
	administrator, err := admin.FindAdministratorById(utils.MustInt(c.Input().Get("id")))
	if err != nil {
		returnJson.StatusCode = Fail
		returnJson.Message = err.Error()
	} else {
		// 用户存在
		if c.Input().Get("password") != "" {
			administrator.Password = utils.GenerateMD5String(c.Input().Get("password"))
		}
		administrator.Nickname = c.Input().Get("nickname")
		administrator.Email = c.Input().Get("email")
		ok, err := administrator.UpdateAdministrator()
		if ok == true {
			returnJson.StatusCode = Success
			returnJson.Message = SaveSuccess
			returnJson.UrlType = Reload
		} else {
			returnJson.StatusCode = Fail
			returnJson.Message = err.Error()
		}
	}
	c.Data["json"] = &returnJson
	c.ServeJSON()
}

// 删除管理员
func (c *AdministratorController) DeleteAdministrator() {
	returnJson:= ResponseJson{}
	administrator := admin.Administrator{}
	administrator.Id = utils.MustInt(c.Input().Get("id"))
	ok := administrator.DeleteAdministrator()
	if ok == true {
		returnJson.StatusCode = Success
		returnJson.Message = DeleteSuccess
	} else {
		returnJson.StatusCode = Fail
		returnJson.Message = DeleteFail
	}
	returnJson.UrlType = Reload
	c.Data["json"] = &returnJson
	c.ServeJSON()
}

// 恢复管理员
func (c *AdministratorController) RecoverAdministrator() {
	returnJson:= ResponseJson{}
	administrator := admin.Administrator{}
	administrator.Id = utils.MustInt(c.Input().Get("id"))
	ok := administrator.RecoverAdministrator()
	if ok == true {
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

// 获取管理员的角色情况
func (c *AdministratorController) GetAdministratorRoles() {
	returnJson:= ResponseJson{}
	data, err := auth.AdministratorRoleList(utils.MustInt(c.Input().Get("id")))
	if err != nil {
		returnJson.StatusCode = Fail
		returnJson.Message = err.Error()
	} else {
		returnJson.StatusCode = Success
		returnJson.Data = data
	}

	c.Data["json"] = &returnJson
	c.ServeJSON()
}

// 更新管理员的角色情况
func (c *AdministratorController) PutAdministratorRoles() {
	returnJson:= ResponseJson{}
	adminId := utils.MustInt(c.Input().Get("adminId"))
	// 数据绑定获取数组
	roleIds := make([]int, 0)
	c.Ctx.Input.Bind(&roleIds, "roleIds")

	administrator := &admin.Administrator{
		Model:models.Model{
			Id:adminId,
		},
	}

	err := administrator.AssignRole(roleIds)
	if err != nil {
		returnJson.StatusCode = Fail
		returnJson.Message = err.Error()
	} else {
		returnJson.StatusCode = Success
		returnJson.Message = SaveSuccess
		returnJson.UrlType = Reload
	}
	c.Data["json"] = &returnJson
	c.ServeJSON()
}
