package common

import (
	"github.com/astaxie/beego"
	"strings"
)

type ErrorController struct {
	beego.Controller
}

func (c *ErrorController) Error404()  {
	if strings.HasPrefix(c.Ctx.Request.URL.String(), "/admin") {
		c.Data["content"] = "page not found"
		c.TplName = "admin/common/404.html"
	}else {
		c.TplName = "common/404.html"
	}
}

func (c *ErrorController) ErrorTip()  {
	tips := "好像出错了呢"
	flash := beego.ReadFromRequest(&c.Controller)
	if n,ok := flash.Data["error"];ok{
		tips = n
	}
	// 显示错误
	c.Data["Error"] = tips
	c.TplName = "common/error.html"
	return
}
