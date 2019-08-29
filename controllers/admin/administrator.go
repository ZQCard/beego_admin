package admin

import (
	"beego_admin/models"
	"beego_admin/models/admin"
	"beego_admin/models/common/auth"
	"beego_admin/utils"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
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
	administrator := admin.AdministratorGORM{}
	administrator.Username = c.Input().Get("username")
	administrator.Password = c.Input().Get("password")
	administrator.Nickname = c.Input().Get("nickname")
	administrator.Email = c.Input().Get("email")
	err := administrator.AdministratorCreate()
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

// 删除管理员
func (c *AdministratorController) DeleteAdministrator() {
	returnJson:= ResponseJson{}
	administrator := admin.AdministratorGORM{}
	administrator.ModelGORM.ID = utils.MustInt(c.Input().Get("id"))
	err := administrator.AdministratorDelete()
	if err == nil {
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
	administrator := admin.AdministratorGORM{}
	administrator.ModelGORM.ID = utils.MustInt(c.Input().Get("id"))
	err := administrator.AdministratorRecover()
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

// 刷新用户的权限和菜单
func (c *AdministratorController)RefreshAuth()  {
	returnJson:= ResponseJson{}

	adminId := c.GetSession("adminId")
	id := 0
	switch adminId.(type) {
	case int:
		id = adminId.(int)
	}
	administrator, err := admin.FindAdministratorById(id)
	// 读取权限map
	authList, err := administrator.AuthList()
	if err != nil {
		returnJson.StatusCode = Fail
		returnJson.Message = "刷新失败"
		c.Data["json"] = &returnJson
		c.ServeJSON()
		return
	}
	// session无法读取存储map[string][]string,所以分为请求方式存储
	for k, v := range authList{
		c.SetSession("AUTH_" + k, v)
	}

	// 读取公共权限
	authCommonList, err := auth.RoleAuthList(beego.AppConfig.String("customer_role_name"))
	if err != nil {
		logs.Error("刷新权限,读取用户公共权限失败")
		returnJson.StatusCode = Fail
		returnJson.Message = "刷新失败"
		c.Data["json"] = &returnJson
		c.ServeJSON()
		return
	}
	// session无法读取存储map[string][]string,所以分为请求方式存储
	for k, v := range authCommonList{
		c.SetSession("AUTH_COMMON" + k, v)
	}

	// 根据GET权限读取后台菜单列表
	authGet := c.GetSession("AUTH_GET")
	var authGetSlice []string
	switch t := authGet.(type) {
	case []string:
		for _, v := range t {
			authGetSlice = append(authGetSlice, v)
		}
	default:
		logs.Error("刷新权限,读取用户菜单失败")
		returnJson.StatusCode = Fail
		returnJson.Message = "刷新失败"
		c.Data["json"] = &returnJson
		c.ServeJSON()
		return
	}
	// 读取管理员的菜单
	c.SetSession("MENU_LEFT", administrator.MenuList(authGetSlice))
	returnJson.StatusCode = Success
	returnJson.Message = "刷新成功"
	c.Data["json"] = &returnJson
	c.ServeJSON()
}

func (c *AdministratorController)GetAdministratorInfo(){
	returnJson := ResponseJson{}
	// 查找管理员
	administrator := admin.AdministratorGORM{
		ModelGORM:models.ModelGORM{
			ID:utils.MustInt(c.Input().Get("id")),
		},
	}
	administrator, err := administrator.FindAdministrator()
	if err != nil {
		returnJson.StatusCode = Fail
		returnJson.Message = err.Error()
	} else {
		// 用户信息
		returnJson.StatusCode = Success
		returnJson.Data = administrator
		returnJson.UrlType = Reload
	}
	c.Data["json"] = &returnJson
	c.ServeJSON()
	return
}

func (c *AdministratorController)PutAdministratorInfo()  {
	returnJson := ResponseJson{}
	administrator := admin.AdministratorGORM{
		ModelGORM:models.ModelGORM{
			ID:utils.MustInt(c.Input().Get("id")),
		},
	}
	administrator, err := administrator.FindAdministrator()
	if err != nil {
		returnJson.StatusCode = Fail
		returnJson.Message = err.Error()

	}else {
		// 用户存在
		if c.Input().Get("password") != "" {
			administrator.Password = utils.GenerateMD5String(c.Input().Get("password"))
		}
		administrator.Nickname = c.Input().Get("nickname")
		administrator.Email = c.Input().Get("email")
		err := administrator.AdministratorUpdate()
		if err == nil {
			returnJson.StatusCode = Success
			returnJson.Message = SaveSuccess
			returnJson.UrlType = Reload
			// 更新完信息 变更昵称
			c.SetSession("nickname", administrator.Nickname)
		} else {
			returnJson.StatusCode = Fail
			returnJson.Message = err.Error()
		}
	}
	c.Data["json"] = &returnJson
	c.ServeJSON()
	return
}
