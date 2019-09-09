package routers

import (
	"beego_admin/controllers/admin"
	"github.com/astaxie/beego"
)

func init() {
	ns :=
		beego.NewNamespace("/admin",
			// 根目录
			beego.NSRouter("/", &admin.IndexController{}),
			// 后台首页
			beego.NSRouter("/index", &admin.IndexController{}),
			///************ 权限管理开始 ***********/
			// 管理员操作
			beego.NSRouter("/auth/administrator", &admin.AdministratorController{}, "get:GetAdministratorList;post:PostAddAdministrator;put:PutAdministratorInfo;delete:DeleteAdministrator;patch:RecoverAdministrator"),
			beego.NSRouter("/administrator/info", &admin.AdministratorController{}, "get:GetAdministratorInfo;put:PutAdministratorInfo"),
			// 角色列表
			beego.NSRouter("/auth/administrator/roles", &admin.AdministratorController{}, "get:GetAdministratorRoles;put:PutAdministratorRoles"),
			beego.NSRouter("/auth/administrator/refreshAuth", &admin.AdministratorController{}, "get:RefreshAuth"),

			// 角色操作
			beego.NSRouter("/auth/role", &admin.AuthController{}, "get:GetRoleList;post:PostAddRole;put:PutUpdateRole;delete:DeleteRole"),
			beego.NSRouter("/auth/role/permissions", &admin.AuthController{}, "get:GetRolePermissions;put:PutRolePermissions"),

			// 权限操作
			beego.NSRouter("/auth/permission", &admin.AuthController{}, "get:GetPermissionList;post:PostAddPermission;put:PutUpdatePermission;delete:DeletePermission"),
			beego.NSRouter("/auth/permission/actions", &admin.AuthController{}, "get:GetPermissionActions;put:PutPermissionActions"),
			// 行为操作
			beego.NSRouter("/auth/action", &admin.AuthController{}, "get:GetActionList;post:PostAddAction;put:PutUpdateAction;delete:DeleteAction"),
			// 菜单操作
			beego.NSRouter("/auth/menu", &admin.AuthController{}, "get:GetMenuList;post:PostAddMenu;put:PutUpdateMenu;delete:DeleteMenu"),
			/************ 权限管理结束 ***********/
			/************ 首页模块功能开始 ***********/
			beego.NSRouter("/homepage/setting", &admin.HomepageControl{}, "get:ModuleList;patch:ModulePatch"),
			/************ 首页模块功能设置 ***********/
			/************ 视频功能开始 ***********/
			// 视频分类
			beego.NSRouter("/video/category", &admin.VideoCategoryControl{}, "get:GetCategoryList;post:PostAddCategory;put:PutUpdateCategory;delete:DeleteCategory;patch:RecoveryCategory"),

			// 视频管理
			beego.NSRouter("/video", &admin.VideoControl{}, "get:GetVideoList;post:PostAddVideo;put:PutUpdateVideo;delete:DeleteVideo;patch:RecoveryVideo"),
			beego.NSRouter("/video/info", &admin.VideoControl{}, "get:GetVideoInfo"),
			/************ 视频功能结束 ***********/


			/************ 资料分类功能开始 ***********/
			beego.NSRouter("/documentation/category", &admin.DocumentationCategoryControl{}, "get:GetCategoryList;post:PostAddCategory;put:PutUpdateCategory;delete:DeleteCategory;patch:RecoveryCategory"),
			// 资料管理
			beego.NSRouter("/documentation", &admin.DocumentationControl{}, "get:GetDocumentationList;post:PostAddDocumentation;put:PutUpdateDocumentation;delete:DeleteDocumentation;patch:RecoveryDocumentation"),
			beego.NSRouter("/documentation/info", &admin.DocumentationControl{}, "get:GetDocumentationInfo"),
			/************ 资料分类功能结束 ***********/

			/************ 轮播图功能开始 ***********/
			beego.NSRouter("/banner", &admin.BannerControl{}, "get:GetBannerList;post:PostAddBanner;put:PutUpdateBanner;delete:DeleteBanner;patch:RecoveryBanner"),
			/************ 轮播图功结束 ***********/

			/************ 留言板功能开始 ***********/
			beego.NSRouter("/message", &admin.MessageControl{}, "get:GetMessageList"),
			/************ 留言板功能开始 ***********/

			/************ 用户功能开始 ***********/
			beego.NSRouter("/user", &admin.UserControl{}, "get:GetUserList;delete:DeleteUser;patch:RecoveryUser"),
			/************ 用户功能开始 ***********/

			/************ 工具箱功能开始 ***********/
			// 文件上传
			beego.NSRouter("/tools/uploadFile", &admin.ToolsController{}, "post:PostUploadFile"),
			// 邮件发送
			beego.NSRouter("/tools/sendEmail", &admin.ToolsController{}, "get:GetSendEmail;post:PostSendEmail"),
			// excel导入导出
			// 页面
			beego.NSRouter("/tools/excel", &admin.ToolsController{}, "get:GetExcel"),
			// 导入
			beego.NSRouter("/tools/excelImport", &admin.ToolsController{}, "post:PostExcelImport"),
			// 导出
			beego.NSRouter("/tools/excelExport", &admin.ToolsController{}, "get:GetExcelExport"),

			/************ 工具箱功能结束 ***********/
			// 退出
			beego.NSRouter("/logout", &admin.IndexController{}, "*:Logout"),
			// 登录页面
			beego.NSRouter("/login", &admin.CommonController{}, "get:LoginPage;post:Login"))
	//注册 namespace
	beego.AddNamespace(ns)


	// 根目录
	// beego.Router("/", &admin.IndexController{})

	// 登录页面
	//beego.Router("/login", &admin.CommonController{}, "get:LoginPage;post:Login")
	//
	//// 退出
	//beego.Router("/logout", &admin.IndexController{}, "*:Logout")
	//
	//// 后台首页
    //beego.Router("/index", &admin.IndexController{})
	//
	///************ 权限管理开始 ***********/
	//// 管理员操作
	//beego.Router("/auth/administrator", &admin.AdministratorController{}, "get:GetAdministratorList;post:PostAddAdministrator;put:PutAdministratorInfo;delete:DeleteAdministrator;patch:RecoverAdministrator")
	//beego.Router("/administrator/info", &admin.AdministratorController{}, "get:GetAdministratorInfo;put:PutAdministratorInfo")
	//// 角色列表
	//beego.Router("/auth/administrator/roles", &admin.AdministratorController{}, "get:GetAdministratorRoles;put:PutAdministratorRoles")
	//beego.Router("/auth/administrator/refreshAuth", &admin.AdministratorController{}, "get:RefreshAuth")
	//
	//// 角色操作
	//beego.Router("/auth/role", &admin.AuthController{}, "get:GetRoleList;post:PostAddRole;put:PutUpdateRole;delete:DeleteRole")
	//beego.Router("/auth/role/permissions", &admin.AuthController{}, "get:GetRolePermissions;put:PutRolePermissions")
	//
	//// 权限操作
	//beego.Router("/auth/permission", &admin.AuthController{}, "get:GetPermissionList;post:PostAddPermission;put:PutUpdatePermission;delete:DeletePermission")
	//beego.Router("/auth/permission/actions", &admin.AuthController{}, "get:GetPermissionActions;put:PutPermissionActions")
	//// 行为操作
	//beego.Router("/auth/action", &admin.AuthController{}, "get:GetActionList;post:PostAddAction;put:PutUpdateAction;delete:DeleteAction")
	//// 菜单操作
	//beego.Router("/auth/menu", &admin.AuthController{}, "get:GetMenuList;post:PostAddMenu;put:PutUpdateMenu;delete:DeleteMenu")
	///************ 权限管理结束 ***********/
	//
	///************ 首页模块功能开始 ***********/
	//beego.Router("/homepage/setting", &admin.HomepageControl{}, "get:ModuleList;patch:ModulePatch")
	///************ 首页模块功能设置 ***********/
	//
	///************ 视频功能开始 ***********/
	//// 视频分类
	//beego.Router("/video/category", &admin.VideoCategoryControl{}, "get:GetCategoryList;post:PostAddCategory;put:PutUpdateCategory;delete:DeleteCategory;patch:RecoveryCategory")
	//
	//// 视频管理
	//beego.Router("/video", &admin.VideoControl{}, "get:GetVideoList;post:PostAddVideo;put:PutUpdateVideo;delete:DeleteVideo;patch:RecoveryVideo")
	//beego.Router("/video/info", &admin.VideoControl{}, "get:GetVideoInfo")
	//
	///************ 视频功能设置 ***********/
	//
	///************ 资料分类功能开始 ***********/
	//beego.Router("/documentation/category", &admin.DocumentationCategoryControl{}, "get:GetCategoryList;post:PostAddCategory;put:PutUpdateCategory;delete:DeleteCategory;patch:RecoveryCategory")
	//// 资料管理
	//beego.Router("/documentation", &admin.DocumentationControl{}, "get:GetDocumentationList;post:PostAddDocumentation;put:PutUpdateDocumentation;delete:DeleteDocumentation;patch:RecoveryDocumentation")
	//beego.Router("/documentation/info", &admin.DocumentationControl{}, "get:GetDocumentationInfo")
	///************ 资料分类功能结束 ***********/
	//
	///************ 轮播图功能开始 ***********/
	//beego.Router("/banner", &admin.BannerControl{}, "get:GetBannerList;post:PostAddBanner;put:PutUpdateBanner;delete:DeleteBanner;patch:RecoveryBanner")
	///************ 轮播图功结束 ***********/
	//
	///************ 留言板功能开始 ***********/
	//beego.Router("/message", &admin.MessageControl{}, "get:GetMessageList")
	///************ 留言板功能开始 ***********/
	//
	///************ 用户功能开始 ***********/
	//beego.Router("/user", &admin.UserControl{}, "get:GetUserList;delete:DeleteUser;patch:RecoveryUser")
	///************ 用户功能开始 ***********/
	//
	///************ 工具箱功能开始 ***********/
	//// 文件上传
	//beego.Router("/tools/uploadFile", &admin.ToolsController{}, "post:PostUploadFile")
	//// 邮件发送
	//beego.Router("/tools/sendEmail", &admin.ToolsController{}, "get:GetSendEmail;post:PostSendEmail")
	//// excel导入导出
	//// 页面
	//beego.Router("/tools/excel", &admin.ToolsController{}, "get:GetExcel")
	//// 导入
	//beego.Router("/tools/excelImport", &admin.ToolsController{}, "post:PostExcelImport")
	//// 导出
	//beego.Router("/tools/excelExport", &admin.ToolsController{}, "get:GetExcelExport")
	//
	///************ 工具箱功能结束 ***********/
}
