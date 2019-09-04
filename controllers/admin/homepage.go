package admin

type HomepageControl struct {
	baseController
}

// 首页模块列表
func (c *HomepageControl)ModuleList()  {
	c.Data["Title"] = "首页模块设置"
	// 模板
	c.TplName = "admin/homepage/index.html"
}

// 首页模块设置
func (c *HomepageControl)ModulePatch()  {

}
