package company

type StaticController struct {
	baseController
}

func (c *StaticController)Promote()  {
	c.Data["Title"] = "能力提升"
	category := c.Input().Get("type")
	if category == ""{
		category = "dazhuan"
	}
	c.TplName = "company/promote_"+category+".html"
}

func (c *StaticController)Contact()  {
	c.Data["Title"] = "联系我们"
	c.TplName = "company/contact.html"
}

func (c *StaticController)Login()  {
	c.Data["Title"] = "登录"
	c.TplName = "company/login.html"
}

func (c *StaticController)Course()  {
	c.Data["Title"] = "课程列表"
	c.TplName = "company/course.html"
}

func (c *StaticController)Resource()  {
	c.Data["Title"] = "资料下载"
	c.TplName = "company/resource.html"
}

func (c *StaticController)Introduce()  {
	c.Data["Title"] = "详情介绍"
	c.TplName = "company/introduce.html"
}

func (c *StaticController)Special()  {
	c.Data["Title"] = "特色班级"
	c.TplName = "company/special.html"
}
