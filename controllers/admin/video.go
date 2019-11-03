package admin

import (
	"beego_admin/models"
	"beego_admin/models/common"
	"beego_admin/utils"
	"fmt"
	"github.com/astaxie/beego"
)

type VideoControl struct {
	baseController
}

// 视频列表
func (c *VideoControl) GetVideoList() {
	// 读取搜索条件
	search := c.KeepSearch()
	page := utils.MustInt(c.Input().Get("page"))
	pageSize := utils.MustInt(c.Input().Get("pageSize"))
	if page == 0 {
		page = PageDefault
	}
	if pageSize == 0 {
		pageSize = PageSizeDefault
	}

	video := common.Video{}
	videos, totalCount := video.List(page, pageSize, search)
	c.Data["Videos"] = videos
	c.Data["TotalCount"] = totalCount
	c.Data["Page"] = page
	c.Data["PageSize"] = pageSize
	c.Data["Title"] = "视频列表"

	// 联图二维码生成
	c.Data["QRcode"] = beego.AppConfig.String("liantu_qr_code")
	// 网站host
	c.Data["Host"] = beego.AppConfig.String("base_url")
	// 模板
	c.TplName = "admin/video/list.html"
}

// 视频详情(添加/编辑页面)
func (c *VideoControl) GetVideoInfo() {

	id := utils.MustInt(c.Input().Get("id"))

	video := common.Video{
		Model:models.Model{
			ID:id,
		},
	}
	title := "视频详情查看"
	if id != 0{
		title = "视频详情编辑"
		v, err := video.Find()
		if err != nil {
			flash := beego.NewFlash()
			fmt.Println(err.Error())
			flash.Error(err.Error())
			flash.Store(&c.Controller)
			c.Redirect("/error", 302)
			return
		}
		video = v
	}
	// 根据阿里云OSS回调地址判断是否需要OSS还是原生文件上传
	callBack := beego.AppConfig.String("aliyun_oss_callBack")
	bucket := beego.AppConfig.String("aliyun_oss_bucket")
	c.Data["AliyunOSSCallback"] = callBack

	ossAddress := bucket+"."+beego.AppConfig.String("end_point")
	c.Data["OssAddress"] = "http://"+ossAddress
	c.Data["Bucket"] = bucket

	c.Data["Video"] = video
	c.Data["Title"] = title
	// 模板
	c.TplName = "admin/video/info.html"
}

// 添加视频
func (c *VideoControl) PostAddVideo() {
	returnJson := ResponseJson{}
	video := common.Video{}
	video.Title = c.Input().Get("title")
	video.Description = c.Input().Get("description")
	video.Url = c.Input().Get("url")
	video.Poster = c.Input().Get("poster")
	err := video.Create()
	if err == nil {
		returnJson.StatusCode = Success
		returnJson.Message = AddSuccess
		returnJson.UrlType = Jump
		returnJson.Url = "/admin/video"
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
	video := common.Video{
		Model:models.Model{
			ID:utils.MustInt(c.Input().Get("id")),
		},
	}
	video.Title = c.Input().Get("title")
	video.Description = c.Input().Get("description")
	video.Url = c.Input().Get("url")
	video.Poster = c.Input().Get("poster")

	err := video.Update()
	if err == nil {
		returnJson.StatusCode = Success
		returnJson.Message = SaveSuccess
		returnJson.UrlType = Jump
		returnJson.Url = "/admin/video"
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
	video := common.Video{
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
	video := common.Video{
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

