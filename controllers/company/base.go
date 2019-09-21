package company

import (
	"beego_admin/models/admin"
	"github.com/astaxie/beego"
	cache2 "github.com/astaxie/beego/cache"
	"sync"
	"time"
)

// 全局简单同步等待锁
var wg sync.WaitGroup

var bc cache2.Cache

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
	// 业务逻辑 读取导航菜单栏 为两级 第一级pid为0 第二级pid不为0 将结果放入文件缓存中
	navigator := admin.Navigator{}
	bc, _ = cache2.NewCache("file", `{"CachePath":"./storage/cache","FileSuffix":".cache","DirectoryLevel":2,"EmbedExpiry":3600}`)
	if bc.Get("NavigatorParentCache") == "" {
		_ = bc.Put("NavigatorParentCache", navigator.FindNavigatorByLevel(0), 3600 * time.Second)
		_ = bc.Put("NavigatorChild", navigator.FindNavigatorByLevel(1), 3600 * time.Second)
	}
	// 这里可以放到每个控制器方法的最后
	c.Data["NavigatorParent"] = bc.Get("NavigatorParentCache")
	c.Data["NavigatorChild"] = bc.Get("NavigatorChild")
	// 后台模板基础布局
	c.Layout = "company/layouts/common.html"

}

