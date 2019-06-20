package routers

import (
	"beego_admin/controllers/admin"
	"github.com/astaxie/beego"
)

func init() {
	// 根目录
	beego.Router("/", &admin.IndexController{})
	// 测试
	beego.Router("/test", &admin.CommonController{}, "*:Test")
	beego.Router("/error", &admin.CommonController{}, "get:Error")

	// 登录页面
	beego.Router("/login", &admin.CommonController{}, "get:LoginPage;post:Login")

	// 退出
	beego.Router("/logout", &admin.IndexController{}, "*:Logout")

	// 后台首页
    beego.Router("/index", &admin.IndexController{})

	/************ 权限管理开始 ***********/
	// 管理员操作
	beego.Router("/auth/administrator", &admin.AdministratorController{}, "get:GetAdministratorList;post:PostAddAdministrator;put:PutUpdateAdministrator;delete:DeleteAdministrator;patch:RecoverAdministrator")
	beego.Router("/auth/administrator/roles", &admin.AdministratorController{}, "get:GetAdministratorRoles;put:PutAdministratorRoles")

	// 角色操作
	beego.Router("/auth/role", &admin.AuthController{}, "get:GetRoleList;post:PostAddRole;put:PutUpdateRole;delete:DeleteRole")
	beego.Router("/auth/role/permissions", &admin.AuthController{}, "get:GetRolePermissions;put:PutRolePermissions")

	// 权限操作
	beego.Router("/auth/permission", &admin.AuthController{}, "get:GetPermissionList;post:PostAddPermission;put:PutUpdatePermission;delete:DeletePermission")
	beego.Router("/auth/permission/actions", &admin.AuthController{}, "get:GetPermissionActions;put:PutPermissionActions")
	// 行为操作
	beego.Router("/auth/action", &admin.AuthController{}, "get:GetActionList;post:PostAddAction;put:PutUpdateAction;delete:DeleteAction")
	// 菜单操作
	beego.Router("/auth/menu", &admin.AuthController{}, "get:GetMenuList;post:PostAddMenu;put:PutUpdateMenu;delete:DeleteMenu")
	/************ 权限管理结束 ***********/


}
