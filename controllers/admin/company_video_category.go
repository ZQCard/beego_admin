package admin

import (
	"beego_admin/models"
	"beego_admin/models/admin"
	"beego_admin/utils"
)

type VideoCategoryControl struct {
	baseController
}


// 视频分类列表
func (c *VideoCategoryControl) GetCategoryList() {
	page := utils.MustInt(c.Input().Get("page"))
	pageSize := utils.MustInt(c.Input().Get("pageSize"))
	if page == 0 {
		page = PageDefault
	}
	if pageSize == 0 {
		pageSize = PageSizeDefault
	}
	category := admin.VideoCategory{}
	categories, totalCount := category.List(page, pageSize)
	c.Data["Categories"] = categories
	c.Data["TotalCount"] = totalCount
	c.Data["Page"] = page
	c.Data["PageSize"] = pageSize
	c.Data["Title"] = "视频分类列表"
	// 模板
	c.TplName = "admin/company_video_category/list.html"
}

// 添加视频分类
func (c *VideoCategoryControl) PostAddCategory() {
	returnJson := ResponseJson{}
	category := admin.VideoCategory{Name:c.Input().Get("name"), Sort:utils.MustInt(c.Input().Get("sort"))}
	err := category.Create()
	if err == nil {
		returnJson.StatusCode = Success
		returnJson.Message = AddSuccess
		returnJson.UrlType = Reload
	} else {
		returnJson.StatusCode = Fail
		returnJson.Message = err.Error()
	}
	c.Data["json"] = &returnJson
	c.ServeJSON()
}

// 更新视频分类
func (c *VideoCategoryControl) PutUpdateCategory() {
	returnJson := ResponseJson{}
	category := admin.VideoCategory{
		Model:models.Model{
			ID:utils.MustInt(c.Input().Get("id")),
		},
		Name:c.Input().Get("name"),
		Sort:utils.MustInt(c.Input().Get("sort")),
	}

	err := category.Update()
	if err == nil {
		returnJson.StatusCode = Success
		returnJson.Message = SaveSuccess
		returnJson.UrlType = Reload
	} else {
		returnJson.StatusCode = Fail
		returnJson.Message = err.Error()
	}
	c.Data["json"] = &returnJson
	c.ServeJSON()
}

// 删除视频分类
func (c *VideoCategoryControl) DeleteCategory() {
	returnJson := ResponseJson{}
	category := admin.VideoCategory{
		Model:models.Model{
			ID:utils.MustInt(c.Input().Get("id")),
		},
	}
	err := category.Delete()
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

// 恢复视频分类
func (c *VideoCategoryControl) RecoveryCategory() {
	returnJson := ResponseJson{}
	category := admin.VideoCategory{
		Model:models.Model{
			ID:utils.MustInt(c.Input().Get("id")),
		},
	}
	err := category.Recover()
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
