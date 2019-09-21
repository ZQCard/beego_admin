package admin

import (
	"beego_admin/models"
	"errors"
	"github.com/astaxie/beego/logs"
	"github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type Video struct {
	models.Model
	CompanyVideoCategoryId int
	Title string
	Url string
	ViewTimes int
	IsShowHomePage int
	IsOnSale int
	Description string
}

func (Video)TableName() string {
	return "company_video"
}

func (video Video)Validate() error {
	return validation.ValidateStruct(&video,
		// 名称不得为空,且大小为1-20字
		validation.Field(
			&video.Title,
			validation.Required.Error("视频标题名称不得为空"),
			validation.Length(1, 60).Error("名称为1-20字")),
		validation.Field(
			&video.Url,
			validation.Required.Error("视频地址名称不得为空"),
			is.URL.Error("视频地址错误"),
			),
	)
}

type VideoInfo struct {
	models.Model
	CompanyVideoId int
	Title string
	Url string
	ViewTimes int
	IsShowHomePage int
	IsOnSale int
	Description string
	Category string
}

// 视频列表
func (video *Video)List(page, pageSize int) (results []VideoInfo, totalCount int64) {
	models.DB.Model(&Video{}).Unscoped().Count(&totalCount)
	// 关联分类表查询信息
	err := models.DB.Table("company_video as v").
		Select("v.*, c.name as category").
		Joins("left join company_video_category as c on v.company_video_category_id = c.id").
		Limit(pageSize).
		Order("id").
		Offset(page - 1).
		Scan(&results).Error
	if err != nil{
		logs.Error("查询视频列表报错", err)
		return nil, 0
	}
	return
}

// 视频列表(根据分类)
func (video *Video)ListFront(page, pageSize int) (results []Video, totalCount int64) {
	models.DB.Model(&video).Where("company_video_category_id = ?", video.CompanyVideoCategoryId).Where("is_on_sale = 1").Count(&totalCount)
	err := models.DB.Where("company_video_category_id = ?", video.CompanyVideoCategoryId).Where("is_on_sale = 1").Offset((page - 1) * pageSize).Limit(pageSize).Find(&results).Error
	if err != nil{
		logs.Error("查询课程列表报错", err)
		return nil, 0
	}
	return
}

// 添加视频信息
func (video *Video)Create() (err error) {
	// 数据验证
	err = video.Validate()
	if err != nil {
		return err
	}
	// 判断用户名或者昵称未使用
	if !models.DB.Unscoped().Where("title = ?", video.Title).Find(&Video{}).RecordNotFound(){
		return errors.New("视频标题已经存在")
	}

	if err = models.DB.Create(&video).Error; err != nil {
		logs.Error("视频创建失败", err)
		return errors.New("信息保存失败")
	}
	return nil
}

// 更新视频信息
func (video *Video)Update() (err error) {
	// 数据验证
	err = video.Validate()
	if err != nil {
		return err
	}
	// 判断用户名或者昵称未使用
	if !models.DB.Unscoped().Select("title").Where("id <> ?", video.ID).Where("title = ?", video.Title).Find(&Video{}).RecordNotFound(){
		return errors.New("视频名称已经存在")
	}
	if err = models.DB.Select("company_video_category_id", "title", "url", "is_show_home_page", "is_on_sale", "description").Save(&video).Error; err != nil {
		logs.Error("视频创建失败", err)
		return errors.New("信息保存失败")
	}
	return nil
}

// 删除视频
func (video *Video)Delete() (err error) {
	if err := models.DB.Delete(&video).Error; err != nil {
		logs.Error("删除视频失败：", err)
		return errors.New("删除失败")
	}
	return nil
}

// 恢复视频
func (video *Video)Recover() (err error) {
	video.DeletedAt = nil
	if err := models.DB.Unscoped().Model(&video).Update("deleted_at", nil).Error; err != nil {
		logs.Error("恢复视频失败：", err)
		return errors.New("恢复视频失败")
	}
	return nil
}

func (video *Video)Find() (v Video, err error) {
	err = models.DB.Unscoped().Where(video).First(&v).Error
	if err != nil {
		return v, errors.New("用户信息错误")
	}
	return v, nil
}
