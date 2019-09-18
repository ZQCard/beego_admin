package company

import (
	"beego_admin/models/admin"
	"github.com/astaxie/beego"
	"sync"
)

// 全局简单同步等待锁
var wg sync.WaitGroup

// 基础控制器
type baseController struct {
	beego.Controller
}

func (c *baseController)Prepare()  {
	// 前台全局CSRF防范
	c.Data["CSRF_TOKEN"] = c.XSRFToken()
	c.Data["Title"] = "无锡菲墨企业管理有限公司"
	c.Data["Keyword"] = "菲墨,教育培训"
	c.Data["Description"] = "菲墨,教育培训"

	// 业务逻辑 读取导航菜单栏 为两级 第一级pid为0 第二级pid不为0
	navigator := admin.Navigator{}
	// 这里可以放到每个控制器方法的最后
	c.Data["NavigatorParent"] = navigator.FindNavigatorByLevel(0)
	c.Data["NavigatorChild"] = navigator.FindNavigatorByLevel(1)
	// 后台模板基础布局
	c.Layout = "company/layouts/common.html"

}

