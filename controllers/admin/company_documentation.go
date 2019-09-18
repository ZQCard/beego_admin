package admin

import (
	"beego_admin/models"
	"beego_admin/models/admin"
	"beego_admin/utils"
)

type DocumentationControl struct {
	baseController
}

// 资料列表
func (c *DocumentationControl) GetDocumentationList() {
	page := utils.MustInt(c.Input().Get("page"))
	pageSize := utils.MustInt(c.Input().Get("pageSize"))
	if page == 0 {
		page = PageDefault
	}
	if pageSize == 0 {
		pageSize = PageSizeDefault
	}
	documentation := admin.Documentation{}
	documentations, totalCount := documentation.List(page, pageSize)
	c.Data["Documentations"] = documentations
	c.Data["TotalCount"] = totalCount
	c.Data["Page"] = page
	c.Data["PageSize"] = pageSize
	c.Data["Title"] = "资料列表"
	// 模板
	c.TplName = "admin/company_documentation/list.html"
}

// 资料详情(添加/编辑页面)
func (c *DocumentationControl) GetDocumentationInfo() {

	id := utils.MustInt(c.Input().Get("id"))

	documentation := admin.Documentation{
		Model:models.Model{
			ID:id,
		},
	}
	title := "资料详情查看"
	if id != 0{
		title = "资料详情编辑"
		v, err := documentation.Find()
		if err != nil {
			c.Data["Error"] = err.Error()
			c.TplName = "admin/common/error.html"
			return
		}
		documentation = v
	}

	category := admin.DocumentationCategory{}
	c.Data["Categories"],_ = category.List(1, 300)
	c.Data["Documentation"] = documentation
	c.Data["Title"] = title
	// 模板
	c.TplName = "admin/company_documentation/info.html"
}

// 添加资料
func (c *DocumentationControl) PostAddDocumentation() {
	returnJson := ResponseJson{}
	documentation := admin.Documentation{}
	documentation.Name = c.Input().Get("name")
	documentation.CompanyDocumentationCategoryId = utils.MustInt(c.Input().Get("company_documentation_category_id"))
	documentation.Description = c.Input().Get("description")
	documentation.Url = c.Input().Get("url")
	err := documentation.Create()
	if err == nil {
		returnJson.StatusCode = Success
		returnJson.Message = AddSuccess
		returnJson.UrlType = Jump
		returnJson.Url = "/admin/documentation"
	} else {
		returnJson.StatusCode = Fail
		returnJson.Message = err.Error()
	}
	c.Data["json"] = &returnJson
	c.ServeJSON()
}

// 更新资料
func (c *DocumentationControl) PutUpdateDocumentation() {
	returnJson := ResponseJson{}
	documentation := admin.Documentation{
		Model:models.Model{
			ID:utils.MustInt(c.Input().Get("id")),
		},
	}
	documentation.Name = c.Input().Get("name")
	documentation.CompanyDocumentationCategoryId = utils.MustInt(c.Input().Get("company_documentation_category_id"))
	documentation.Description = c.Input().Get("description")
	documentation.Url = c.Input().Get("url")
	err := documentation.Update()
	if err == nil {
		returnJson.StatusCode = Success
		returnJson.Message = SaveSuccess
		returnJson.UrlType = Jump
		returnJson.Url = "/admin/documentation"
	} else {
		returnJson.StatusCode = Fail
		returnJson.Message = err.Error()
	}
	c.Data["json"] = &returnJson
	c.ServeJSON()
}

// 删除资料
func (c *DocumentationControl) DeleteDocumentation() {
	returnJson := ResponseJson{}
	documentation := admin.Documentation{
		Model:models.Model{
			ID:utils.MustInt(c.Input().Get("id")),
		},
	}
	err := documentation.Delete()
	if err == nil {
		returnJson.StatusCode = Success
		returnJson.Message = DeleteSuccess
	} else {
		returnJson.StatusCode = Fail
		returnJson.Message = DeleteFail
	}
	returnJson.UrlType = Reload
	c.Data["json"] = &returnJson
	c.ServeJSON()
}

// 恢复资料
func (c *DocumentationControl) RecoveryDocumentation() {
	returnJson := ResponseJson{}
	documentation := admin.Documentation{
		Model:models.Model{
			ID:utils.MustInt(c.Input().Get("id")),
		},
	}
	err := documentation.Recover()
	if err == nil {
		returnJson.StatusCode = Success
		returnJson.Message = SaveSuccess
	} else {
		returnJson.StatusCode = Fail
		returnJson.Message = SaveFail
	}
	returnJson.UrlType = Reload
	c.Data["json"] = &returnJson
	c.ServeJSON()
}

