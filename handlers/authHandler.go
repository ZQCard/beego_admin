package handlers

import (
	"beego_admin/controllers/admin"
	"encoding/json"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/utils"
)

func Auth() func(ctx *context.Context) {
	var PermissionSupport = func(ctx *context.Context) {
		// 未登录状态
		if nil == ctx.Input.Session("adminId") {
			if ctx.Input.URL() != "/login" {
				ctx.Redirect(302, "/login")
			}
		} else { // 已登录
			// 读取游客权限
			customerAuth := ctx.Input.Session("AUTH_COMMON"+ctx.Input.Method())
			var authSlice []string
			switch t := customerAuth.(type) {
			case []string:
				for _, v := range t {
					authSlice = append(authSlice, v)
				}
			}

			// 读取当前用户权限路由
			userAuth:= ctx.Input.Session("AUTH_"+ctx.Input.Method())
			switch t := userAuth.(type) {
			case []string:
				for _, v := range t {
					authSlice = append(authSlice, v)
				}
			}
			if !utils.InSlice(ctx.Input.URL(), authSlice) {
				ctx.ResponseWriter.WriteHeader(403)
				// ajax请求返回json
				if ctx.Input.IsAjax() {
					responseJson := admin.ResponseJson{}
					responseJson.StatusCode = admin.Fail
					responseJson.Message = "权限不足"
					data,_ := json.Marshal(responseJson)
					ctx.Output.Body(data)
				} else {
					ctx.Output.Body([]byte("forbidden！权限不足"))
				}
			}

		}
	}
	return PermissionSupport
}
