package routers

import (
	"beego_admin/controllers/admin"
	"beego_admin/controllers/common"
	"beego_admin/controllers/company"
	"beego_admin/handlers"
	"github.com/astaxie/beego"
)

func init() {
	// 前台路由
	beego.Router("/", &company.IndexController{}, "get:Index")
	// 前台错误提示页面
	beego.Router("/error", &common.ErrorController{}, "get:ErrorTip")

	// 微信接口路由
	WeChatNameSpace :=
		beego.NewNamespace("/wechat",
			// 微信网页授权,获取用户信息
			beego.NSRouter("/login", &common.WeChatController{}, "get:GetUserLogin"),
		)
	//注册 namespace
	beego.AddNamespace(WeChatNameSpace)

	// 公共路由
	commonNameSpace :=
		beego.NewNamespace("/common",
			// 阿里云OSS WEB直传 公共接口
			beego.NSRouter("/uploadFile/OSS", &common.AliyunOSSController{}, "get:GetPolicy;post:PostCallback"),
			// 文件上传 公共接口
			beego.NSRouter("/uploadFile", &common.UploadFileController{}, "post:PostUploadFile"),
			)
	//注册 namespace
	beego.AddNamespace(commonNameSpace)

	// 管理后台路由
	ns :=
		beego.NewNamespace("/admin",
			// 后台权限过滤器
			beego.NSBefore(handlers.Auth()),
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


			/************ 视频功能开始 ***********/
			// 视频管理
			beego.NSRouter("/video", &admin.VideoControl{}, "get:GetVideoList;post:PostAddVideo;put:PutUpdateVideo;delete:DeleteVideo;patch:RecoveryVideo"),
			beego.NSRouter("/video/info", &admin.VideoControl{}, "get:GetVideoInfo"),
			/************ 视频功能结束 ***********/

			// 退出
			beego.NSRouter("/logout", &admin.IndexController{}, "*:Logout"),
			// 登录页面
			beego.NSRouter("/login", &admin.CommonController{}, "get:LoginPage;post:Login"))
	//注册 namespace
	beego.AddNamespace(ns)
}
