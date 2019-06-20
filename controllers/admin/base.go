package admin

import (
	"github.com/astaxie/beego"
	"html/template"
)

// 基础控制器
type baseController struct {
	beego.Controller
}

func (c *baseController)Prepare()  {
	// 后台全局CSRF防范
	c.Data["CSRF_TOKEN"] = c.XSRFToken()

	// 查看是否有缓存菜单
	if nil == c.GetSession("MENU_LEFT") {
		// 清空session
		c.DestroySession()
		c.Data["Error"] = "解析权限失败"
		c.Data["CsrfData"] = template.HTML(c.XSRFFormHTML())
		c.TplName = "admin/common/login.html"
		return
	}

	// 模板菜单
	c.Data["MenuLeft"] = c.GetSession("MENU_LEFT")
	// 后台模板基础布局
	c.Layout = "admin/layouts/common.html"
	c.LayoutSections = make(map[string]string)
	c.Data["Nickname"] = c.GetSession("nickname")
	c.LayoutSections["HtmlHeader"] = "admin/layouts/header.html"
	c.LayoutSections["Menu"] = "admin/layouts/menu.html"
}

