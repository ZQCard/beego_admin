package common

import (
	"beego_admin/models"
	"errors"
	"github.com/astaxie/beego/logs"
	"github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/jinzhu/gorm"
)

type Video struct {
	models.Model
	Title string
	Url string
	ViewTimes int
	Description string
	Poster string
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
		validation.Field(
			&video.Poster,
			validation.Required.Error("视频封面名称不得为空"),
			is.URL.Error("视频封面地址错误"),
		),
	)
}

// 视频列表
func (video *Video)List(page, pageSize int, where map[string]interface{}) (results []Video, totalCount int64) {
	tempDb := models.DB.Model(&Video{})
	// 如果有搜索条件 加入where条件中
	if title, ok := where["title"]; ok != false{
		tempDb = tempDb.Where( "title LIKE ? ", "%"+title.(string)+"%")
	}
	tempDb.Unscoped().Count(&totalCount)
	err := tempDb.Unscoped().Offset((page - 1) * pageSize).Limit(pageSize).Find(&results).Error
	if err != nil{
		logs.Error("查询视频列表报错", err)
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
	if err = models.DB.Select("title", "url", "description", "poster").Save(&video).Error; err != nil {
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

// 查找单条记录
func (video *Video)Find() (v Video, err error) {
	err = models.DB.Where(video).First(&v).Error
	if err != nil {
		return v, errors.New("视频不存在")
	}
	return v, nil
}

// 增加观看次数
func (video *Video)AddViewTimes()  {
	models.DB.Model(&video).UpdateColumn("view_times", gorm.Expr("view_times + ?", 1))
}
