package admin

import (
	"github.com/astaxie/beego"
	"html/template"
	"strings"
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
	c.Data["AdminId"] = c.GetSession("adminId")
	c.LayoutSections["HtmlHeader"] = "admin/layouts/header.html"
	c.LayoutSections["Menu"] = "admin/layouts/menu.html"
}

// 保持url搜索参数数据回显
func (c *baseController)KeepSearch() map[string]interface{} {
	// 循环搜索参数,然后赋值保持回显
	search := c.Input()
	searchMap := make(map[string]interface{})
	length := len("Search_")
	for k, v := range search{
		// 如果搜索参数以search_开头则赋值到模板当中
		if strings.HasPrefix(k, "Search_") {
			c.Data[k] = v[0]
			if v[0] != ""{
				searchMap[strings.ToLower(beego.Substr(k, length, len(k) - length))] = v[0]
			}
		}
	}
	return searchMap
}

