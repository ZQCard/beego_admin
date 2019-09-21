package company

import (
	"beego_admin/models/admin"
	"beego_admin/utils"
)

type StaticController struct {
	baseController
}

// 学历提升
func (c *StaticController)Promote()  {
	c.Data["Title"] = "学历提升-无锡菲墨"
	category := c.Input().Get("type")
	if category == ""{
		category = "dazhuan"
	}
	c.TplName = "company/promote_"+category+".html"
}
// 英语能力考试
func (c *StaticController)English()  {
	c.Data["Title"] = "外语教学-无锡菲墨"
	category := c.Input().Get("type")
	if category == ""{
		category = "jqyy"
	}
	c.TplName = "company/english_"+category+".html"
}

// 经济金融考试
func (c *StaticController)Finance()  {
	c.Data["Title"] = "经济金融-无锡菲墨"
	category := c.Input().Get("type")
	if category == ""{
		category = "jjs"
	}
	c.TplName = "company/finance_"+category+".html"
}

// 关于我们
func (c *StaticController)Contact()  {
	c.Data["Title"] = "联系我们"
	c.TplName = "company/contact.html"
}

// 登录
func (c *StaticController)Login()  {
	c.Data["Title"] = "登录"
	c.TplName = "company/login.html"
}

// 课程
func (c *StaticController)Course()  {
	// 读取所有课程分类
	category := admin.VideoCategory{}
	categories,_ := category.ListFront(1, 1000)
	id := utils.MustInt(c.Input().Get("id"))
	name := ""
	if id == 0{
		id = categories[0].ID
		name = categories[0].Name
	}else {
		for _,category := range categories{
			if id == category.ID{
				name = category.Name
			}
		}
	}
	if name == ""{
		c.Redirect("/error", 302)
	}
	// 查找分类下的课程
	page := utils.MustInt(c.Input().Get("page"))
	pageSize := utils.MustInt(c.Input().Get("pageSize"))
	if page == 0 {
		page = 1
	}
	if pageSize == 0 {
		pageSize = 10
	}
	video := admin.Video{
		CompanyVideoCategoryId:id,
	}
	videos, totalCount := video.ListFront(page, pageSize)
	c.Data["Videos"] = videos
	c.Data["TotalCount"] = totalCount
	c.Data["Page"] = page
	c.Data["PageSize"] = pageSize


	c.Data["CurrentID"] = id
	c.Data["CurrentName"] = name
	c.Data["Categories"] = categories
	c.Data["Title"] = "课程列表"
	c.TplName = "company/course.html"
}

// 资源下载
func (c *StaticController)Resource()  {
	// 读取所有课程分类
	category := admin.DocumentationCategory{}
	categories,_ := category.ListFront(1, 1000)
	id := utils.MustInt(c.Input().Get("id"))
	name := ""
	if id == 0{
		id = categories[0].ID
		name = categories[0].Name
	}else {
		for _,category := range categories{
			if id == category.ID{
				name = category.Name
			}
		}
	}
	if name == ""{
		c.Redirect("/error", 302)
	}
	// 查找分类下的课程
	page := utils.MustInt(c.Input().Get("page"))
	pageSize := utils.MustInt(c.Input().Get("pageSize"))
	if page == 0 {
		page = 1
	}
	if pageSize == 0 {
		pageSize = 1
	}
	document := admin.Documentation{
		CompanyDocumentationCategoryId:id,
	}
	documents, totalCount := document.ListFront(page, pageSize)
	c.Data["Documents"] = documents
	c.Data["TotalCount"] = totalCount
	c.Data["Page"] = page
	c.Data["PageSize"] = pageSize

	c.Data["CurrentID"] = id
	c.Data["CurrentName"] = name
	c.Data["Categories"] = categories

	c.Data["Title"] = "资料下载"
	c.TplName = "company/resource.html"
}

func (c *StaticController)Introduce()  {
	c.Data["Title"] = "详情介绍"
	c.TplName = "company/introduce.html"
}