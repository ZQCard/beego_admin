package company

import (
	"beego_admin/models/admin"
	"fmt"
)

type IndexController struct {
	baseController
}

func (c *IndexController)HomePage()  {
	// 读取首页设置
	homePageSetting := admin.Homepage{}
	c.Data["HomePageSetting"] = homePageSetting.ModuleList()

	// 读取轮播图
	banner := admin.Banner{}
	c.Data["Banners"] = banner.CompanyList()
	fmt.Println(c.Data["Banners"])

	// 读取特色课程以及分类
	// 读取
	c.TplName = "company/index.html"
}