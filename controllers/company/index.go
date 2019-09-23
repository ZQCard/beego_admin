package company

import (
	"beego_admin/models/admin"
)

type IndexController struct {
	baseController
}

func (c *IndexController)HomePage()  {
	// 读取首页设置
	homePageSetting := admin.Homepage{}
	homePageSettingList := homePageSetting.ModuleList()
	homeBanner := homePageSettingList[0].IsShow
	homeClass := homePageSettingList[1].IsShow
	homeVideo := homePageSettingList[2].IsShow
	homeHighLine := homePageSettingList[3].IsShow
	c.Data["HomeBanner"] = homeBanner
	c.Data["HomeClass"] = homeClass
	c.Data["HomeVideo"] = homeVideo
	c.Data["HighLine"] = homeHighLine

	// 读取轮播图
	banner := admin.Banner{}
	c.Data["Banners"] = banner.CompanyList()

	// 读取特色课程以及分类
	// 读取
	c.TplName = "company/index.html"
}
