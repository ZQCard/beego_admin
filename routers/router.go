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
	beego.Router("/auth/administrator", &admin.AdministratorController{}, "get:GetAdministratorList;post:PostAddAdministrator;put:PutAdministratorInfo;delete:DeleteAdministrator;patch:RecoverAdministrator")
	beego.Router("/administrator/info", &admin.AdministratorController{}, "get:GetAdministratorInfo;put:PutAdministratorInfo")
	// 角色列表
	beego.Router("/auth/administrator/roles", &admin.AdministratorController{}, "get:GetAdministratorRoles;put:PutAdministratorRoles")
	beego.Router("/auth/administrator/refreshAuth", &admin.AdministratorController{}, "get:RefreshAuth")

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

	/************ 首页模块功能开始 ***********/
	beego.Router("/homepage/setting", &admin.HomepageControl{}, "get:ModuleList;patch:ModulePatch")
	/************ 首页模块功能设置 ***********/

	/************ 工具箱功能开始 ***********/
	// 文件上传
	beego.Router("/tools/uploadFile", &admin.ToolsController{}, "get:GetUploadFile;post:PostUploadFile")
	// 邮件发送
	beego.Router("/tools/sendEmail", &admin.ToolsController{}, "get:GetSendEmail;post:PostSendEmail")
	// excel导入导出
	// 页面
	beego.Router("/tools/excel", &admin.ToolsController{}, "get:GetExcel")
	// 导入
	beego.Router("/tools/excelImport", &admin.ToolsController{}, "post:PostExcelImport")
	// 导出
	beego.Router("/tools/excelExport", &admin.ToolsController{}, "get:GetExcelExport")

	/************ 工具箱功能结束 ***********/
}
