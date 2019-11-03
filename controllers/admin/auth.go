package admin

import (
	"beego_admin/models/admin"
	"beego_admin/utils"
	"encoding/json"
	"github.com/astaxie/beego/logs"
)

type AuthController struct {
	baseController
}

/*************************** 角色功能开始 ******************************/

// 角色列表
func (c *AuthController) GetRoleList() {
	search := c.KeepSearch()
	page := utils.MustInt(c.Input().Get("page"))
	pageSize := utils.MustInt(c.Input().Get("pageSize"))
	if page == 0 {
		page = PageDefault
	}
	if pageSize == 0 {
		pageSize = PageSizeDefault
	}
	roles, totalCount := admin.RoleList(page, pageSize, search)
	c.Data["Roles"] = roles
	c.Data["TotalCount"] = totalCount
	c.Data["Page"] = page
	c.Data["PageSize"] = pageSize
	c.Data["Title"] = "角色列表"
	// 模板
	c.TplName = "admin/auth/role.html"
}

// 添加角色
func (c *AuthController) PostAddRole() {
	returnJson := ResponseJson{}
	role := admin.Role{Name:c.Input().Get("name")}
	err := role.RoleCreate()
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

// 更新角色
func (c *AuthController) PutUpdateRole() {
	returnJson := ResponseJson{}
	role := admin.Role{ID:utils.MustInt(c.Input().Get("id")), Name:c.Input().Get("name")}
	err := role.RoleUpdate()
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

// 删除角色
func (c *AuthController) DeleteRole() {
	returnJson := ResponseJson{}
	role := admin.Role{ID:utils.MustInt(c.Input().Get("id"))}
	err := role.RoleDelete()
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

// 获取角色的权限情况
func (c *AuthController) GetRolePermissions() {
	returnJson := ResponseJson{}
	role := admin.Role{ID:utils.MustInt(c.Input().Get("id"))}
	data, err := role.RolePermissionList()
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

// 更新角色的行为情况
func (c *AuthController) PutRolePermissions() {
	returnJson := ResponseJson{}
	roleId := utils.MustInt(c.Input().Get("roleId"))
	// 数据绑定获取数组
	permissionIds := make([]int, 0)
	c.Ctx.Input.Bind(&permissionIds, "permissionIds")
	role := admin.Role{ID: roleId}
	err := role.AssignPermission(permissionIds)
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

/*************************** 角色功能结束 ******************************/

/*************************** 权限功能开始 ******************************/

// 权限列表
func (c *AuthController) GetPermissionList() {
	search := c.KeepSearch()
	page := utils.MustInt(c.Input().Get("page"))
	pageSize := utils.MustInt(c.Input().Get("pageSize"))
	if page == 0 {
		page = PageDefault
	}

	if pageSize == 0 {
		pageSize = PageSizeDefault
	}
	permissions, totalCount := admin.PermissionList(page, pageSize, search)
	c.Data["Permissions"] = permissions
	c.Data["TotalCount"] = totalCount
	c.Data["Page"] = page
	c.Data["PageSize"] = pageSize
	c.Data["Title"] = "权限列表"
	// 模板
	c.TplName = "admin/auth/permission.html"
}

// 添加权限
func (c *AuthController) PostAddPermission() {
	returnJson := ResponseJson{}
	permission := admin.Permission{Name:c.Input().Get("name")}
	err := permission.PermissionCreate()
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

// 更新权限
func (c *AuthController) PutUpdatePermission() {
	returnJson := ResponseJson{}
	permission := admin.Permission{ID:utils.MustInt(c.Input().Get("id")), Name:c.Input().Get("name")}

	err := permission.PermissionUpdate()
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

// 删除权限
func (c *AuthController) DeletePermission() {
	returnJson := ResponseJson{}
	permission := admin.Permission{ID:utils.MustInt(c.Input().Get("id"))}

	err := permission.PermissionDelete()
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

// 获取权限的行为情况
func (c *AuthController) GetPermissionActions() {
	returnJson := ResponseJson{}
	permission := admin.Permission{ID:utils.MustInt(c.Input().Get("id"))}
	data, err := permission.PermissionActionList()
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

// 更新权限的行为情况
func (c *AuthController) PutPermissionActions() {
	returnJson := ResponseJson{}
	permissionId := utils.MustInt(c.Input().Get("permissionId"))
	// 数据绑定获取数组
	actionIds := make([]int, 0)
	c.Ctx.Input.Bind(&actionIds, "actionIds")
	permission := admin.Permission{ID: permissionId}
	err := permission.AssignAction(actionIds)
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

/*************************** 权限功能结束 ******************************/


/*************************** 行为功能开始 ******************************/

// 行为列表
func (c *AuthController) GetActionList() {
	search := c.KeepSearch()
	page := utils.MustInt(c.Input().Get("page"))
	pageSize := utils.MustInt(c.Input().Get("pageSize"))
	if page == 0 {
		page = PageDefault
	}
	if pageSize == 0 {
		pageSize = PageSizeDefault
	}
	actions, totalCount := admin.ActionList(page, pageSize, search)
	c.Data["Actions"] = actions
	c.Data["TotalCount"] = totalCount
	c.Data["Page"] = page
	c.Data["PageSize"] = pageSize
	c.Data["Title"] = "权限列表"
	// 模板
	c.TplName = "admin/auth/action.html"
}

// 添加行为
func (c *AuthController) PostAddAction() {
	returnJson := ResponseJson{}
	action := admin.Action{Name:c.Input().Get("name"), Method:c.Input().Get("method"), Route:c.Input().Get("route")}
	err := action.ActionCreate()
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

// 更新行为
func (c *AuthController) PutUpdateAction() {
	returnJson := ResponseJson{}
	action := admin.Action{
		ID:utils.MustInt(c.Input().Get("id")),
		Name:c.Input().Get("name"),
		Method:c.Input().Get("method"),
		Route:c.Input().Get("route"),
	}

	err := action.ActionUpdate()
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

// 删除行为
func (c *AuthController) DeleteAction() {
	returnJson := ResponseJson{}
	action := admin.Action{
		ID:utils.MustInt(c.Input().Get("id")),
	}
	err := action.ActionDelete()
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

/*************************** 行为功能开始 ******************************/

/*************************** 菜单功能开始 ******************************/
// 菜单列表
func (c *AuthController) GetMenuList() {
	menu := admin.Menu{}
	pid := utils.MustInt(c.Input().Get("id"))
	menu.Pid = pid
	treeList := menu.MenuList([]string{})
	menus, err := json.Marshal(treeList)
	if err != nil {
		logs.Error("读取菜单列表错误", err)
	}
	c.Data["Menu"] = string(menus)
	c.Data["Title"] = "菜单列表"
	// 模板
	c.TplName = "admin/auth/menu.html"
}

// 删除菜单
func (c *AuthController) DeleteMenu() {
	returnJson := ResponseJson{}
	menu := &admin.Menu{ID: utils.MustInt(c.Input().Get("id"))}
	err := menu.MenuDelete()
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

// 添加菜单
func (c *AuthController) PostAddMenu() {
	returnJson := ResponseJson{}
	menu := &admin.Menu{}
	menu.Pid = utils.MustInt(c.Input().Get("pid"))
	menu.Name = c.Input().Get("name")
	menu.Sort = utils.MustInt(c.Input().Get("sort"))
	menu.Route = c.Input().Get("route")
	err := menu.MenuCreate()
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

// 更新菜单
func (c *AuthController) PutUpdateMenu() {
	returnJson := ResponseJson{}
	menu := &admin.Menu{}
	menu.Pid = utils.MustInt(c.Input().Get("pid"))
	menu.Name = c.Input().Get("name")
	menu.Sort = utils.MustInt(c.Input().Get("sort"))
	menu.Route = c.Input().Get("route")
	menu.ID = utils.MustInt(c.Input().Get("id"))
	err := menu.MenuUpdate()
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

/*************************** 菜单功能结束 ******************************/
