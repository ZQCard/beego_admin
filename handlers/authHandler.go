package handlers

import (
	"beego_admin/controllers/admin"
	"fmt"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/utils"
)

func Auth() func(ctx *context.Context) {
	var PermissionSupport = func(ctx *context.Context) {
		// 未登录状态
		if nil == ctx.Input.Session("adminId") {
			if ctx.Input.URL() != "/admin/login" {
				ctx.Redirect(302, "/admin/login")
			}
		} else { // 已登录
			// 读取公共权限
			customerAuth := ctx.Input.Session("AUTH_COMMON"+ctx.Input.Method())
			var authSlice []string
			switch t := customerAuth.(type) {
			case []string:
				authSlice = append(authSlice, t...)
			}

			// 读取当前用户权限路由
			userAuth:= ctx.Input.Session("AUTH_"+ctx.Input.Method())
			switch t := userAuth.(type) {
			case []string:
				authSlice = append(authSlice, t...)
			}
			fmt.Println(ctx.Input.URL())
			fmt.Println(authSlice)
			if !utils.InSlice(ctx.Input.URL(), authSlice) {
				// ajax请求返回json
				if ctx.Input.IsAjax() {
					responseJson := admin.ResponseJson{}
					responseJson.StatusCode = admin.Fail
					responseJson.Message = "权限不足"
					//data,_ := json.Marshal(responseJson)
					ctx.Output.JSON(responseJson, false, false)
				} else {
					fmt.Println(1111111111111)

					ctx.Output.Session("error", "权限不足")
					ctx.Redirect(302, "/admin/error")
				}
			}

		}
	}
	return PermissionSupport
}
