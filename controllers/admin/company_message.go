package admin

import (
	"beego_admin/models/admin"
	"beego_admin/utils"
)

type MessageControl struct {
	baseController
}


// 留言列表
func (c *MessageControl) GetMessageList() {
	page := utils.MustInt(c.Input().Get("page"))
	pageSize := utils.MustInt(c.Input().Get("pageSize"))
	if page == 0 {
		page = PageDefault
	}
	if pageSize == 0 {
		pageSize = PageSizeDefault
	}
	message := admin.Message{}
	messages, totalCount := message.List(page, pageSize)
	c.Data["Messages"] = messages
	c.Data["TotalCount"] = totalCount
	c.Data["Page"] = page
	c.Data["PageSize"] = pageSize
	c.Data["Title"] = "留言列表"
	// 模板
	c.TplName = "admin/company_message/list.html"
}

