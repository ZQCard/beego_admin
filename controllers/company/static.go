package company

type StaticController struct {
	baseController
}

func (c *StaticController)Promote()  {
	c.Data["Title"] = "学历提升-无锡菲墨"
	category := c.Input().Get("type")
	if category == ""{
		category = "dazhuan"
	}
	c.TplName = "company/promote_"+category+".html"
}

func (c *StaticController)English()  {
	c.Data["Title"] = "外语教学-无锡菲墨"
	category := c.Input().Get("type")
	if category == ""{
		category = "jqyy"
	}
	c.TplName = "company/english_"+category+".html"
}

func (c *StaticController)Finance()  {
	c.Data["Title"] = "经济金融-无锡菲墨"
	category := c.Input().Get("type")
	if category == ""{
		category = "jjs"
	}
	c.TplName = "company/finance_"+category+".html"
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