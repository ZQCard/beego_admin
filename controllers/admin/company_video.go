package admin

import (
	"beego_admin/models"
	"beego_admin/models/admin"
	"beego_admin/utils"
	"fmt"
)

type VideoControl struct {
	baseController
}

// 视频列表
func (c *VideoControl) GetVideoList() {
	page := utils.MustInt(c.Input().Get("page"))
	pageSize := utils.MustInt(c.Input().Get("pageSize"))
	if page == 0 {
		page = PageDefault
	}
	if pageSize == 0 {
		pageSize = PageSizeDefault
	}
	video := admin.Video{}
	videos, totalCount := video.List(page, pageSize)
	c.Data["Videos"] = videos
	c.Data["TotalCount"] = totalCount
	c.Data["Page"] = page
	c.Data["PageSize"] = pageSize
	c.Data["Title"] = "视频列表"
	// 模板
	c.TplName = "admin/company_video/list.html"
}

// 视频详情(添加/编辑页面)
func (c *VideoControl) GetVideoInfo() {

	id := utils.MustInt(c.Input().Get("id"))

	video := admin.Video{
		Model:models.Model{
			ID:id,
		},
	}
	title := "视频详情查看"
	if id != 0{
		title = "视频详情编辑"
		v, err := video.Find()
		if err != nil {
			c.Data["Error"] = err.Error()
			c.TplName = "admin/common/error.html"
			return
		}
		video = v
	}

	category := admin.VideoCategory{}
	c.Data["Categories"],_ = category.List(1, 300)
	c.Data["Video"] = video
	c.Data["Title"] = title
	// 模板
	c.TplName = "admin/company_video/info.html"
}

// 添加视频
func (c *VideoControl) PostAddVideo() {
	returnJson := ResponseJson{}
	video := admin.Video{}
	video.Title = c.Input().Get("title")
	video.CompanyVideoCategoryId = utils.MustInt(c.Input().Get("company_video_category_id"))
	video.IsOnSale = utils.MustInt(c.Input().Get("is_on_sale"))
	video.IsShowHomePage = utils.MustInt(c.Input().Get("is_show_home_page"))
	video.Description = c.Input().Get("description")
	video.Url = c.Input().Get("url")
	fmt.Println(video)
	err := video.Create()
	if err == nil {
		returnJson.StatusCode = Success
		returnJson.Message = AddSuccess
		returnJson.UrlType = Jump
		returnJson.Url = "/video"
	} else {
		returnJson.StatusCode = Fail
		returnJson.Message = err.Error()
	}
	c.Data["json"] = &returnJson
	c.ServeJSON()
}

// 更新视频
func (c *VideoControl) PutUpdateVideo() {
	returnJson := ResponseJson{}
	video := admin.Video{
		Model:models.Model{
			ID:utils.MustInt(c.Input().Get("id")),
		},
	}
	video.Title = c.Input().Get("title")
	video.CompanyVideoCategoryId = utils.MustInt(c.Input().Get("company_video_category_id"))
	video.IsOnSale = utils.MustInt(c.Input().Get("is_on_sale"))
	video.IsShowHomePage = utils.MustInt(c.Input().Get("is_show_home_page"))
	video.Description = c.Input().Get("description")
	video.Url = c.Input().Get("url")
	err := video.Update()
	if err == nil {
		returnJson.StatusCode = Success
		returnJson.Message = AddSuccess
		returnJson.UrlType = Jump
		returnJson.Url = "/video"
	} else {
		returnJson.StatusCode = Fail
		returnJson.Message = err.Error()
	}
	c.Data["json"] = &returnJson
	c.ServeJSON()
}

// 删除视频
func (c *VideoControl) DeleteVideo() {
	returnJson := ResponseJson{}
	video := admin.Video{
		Model:models.Model{
			ID:utils.MustInt(c.Input().Get("id")),
		},
	}
	err := video.Delete()
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

// 恢复视频
func (c *VideoControl) RecoveryVideo() {
	returnJson := ResponseJson{}
	video := admin.Video{
		Model:models.Model{
			ID:utils.MustInt(c.Input().Get("id")),
		},
	}
	err := video.Recover()
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

