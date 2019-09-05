package admin

import (
	"beego_admin/models"
	"errors"
	"github.com/astaxie/beego/logs"
	"github.com/go-ozzo/ozzo-validation"
)

type VideoCategory struct {
	models.Model
	Name string		`json:"name"`
	Sort int 		`json:"sort"`
}

func (VideoCategory)TableName() string {
	return "company_video_category"
}


func (category VideoCategory)Validate() error {
	return validation.ValidateStruct(&category,
		// 名称不得为空,且大小为1-20字
		validation.Field(
			&category.Name,
			validation.Required.Error("视频分类名称不得为空"),
			validation.Length(1, 60).Error("名称为1-20字")),
		// 路由
		validation.Field(
			&category.Sort,
			validation.Required.Error("排序不得为空"),
		),
	)
}

// 视频分类列表
func (category *VideoCategory)List(page, pageSize int) (categories []VideoCategory, totalCount int64) {
	models.DB.Model(&categories).Unscoped().Count(&totalCount)
	err := models.DB.Unscoped().Offset((page - 1) * pageSize).Limit(pageSize).Find(&categories).Error
	if err != nil{
		logs.Error("查询视频分类列表报错", err)
		return nil, 0
	}
	return
}

// 添加视频分类信息
func (category *VideoCategory)Create() (err error) {
	// 数据验证
	err = category.Validate()
	if err != nil {
		return err
	}
	// 判断用户名或者昵称未使用
	if !models.DB.Unscoped().Where("name = ?", category.Name).Find(&VideoCategory{}).RecordNotFound(){
		return errors.New("视频分类名称已经存在")
	}

	if err = models.DB.Create(&category).Error; err != nil {
		logs.Error("视频分类创建失败", err)
		return errors.New("信息保存失败")
	}
	return nil
}

// 更新视频分类信息
func (category *VideoCategory)Update() (err error) {
	// 数据验证
	err = category.Validate()
	if err != nil {
		return err
	}
	// 判断用户名或者昵称未使用
	if !models.DB.Unscoped().Select("name", "sort").Where("id <> ?", category.ID).Where("name = ?", category.Name).Find(&VideoCategory{}).RecordNotFound(){
		return errors.New("视频分类名称已经存在")
	}

	if err = models.DB.Save(&category).Error; err != nil {
		logs.Error("视频分类创建失败", err)
		return errors.New("信息保存失败")
	}
	return nil
}

// 删除视频分类
func (category *VideoCategory)Delete() (err error) {
	if err := models.DB.Delete(&category).Error; err != nil {
		logs.Error("删除视频分类失败：", err)
		return errors.New("删除失败")
	}
	return nil
}

// 恢复视频分类
func (category *VideoCategory)Recover() (err error) {
	category.DeletedAt = nil
	if err := models.DB.Unscoped().Model(&category).Update("deleted_at", nil).Error; err != nil {
		logs.Error("恢复视频分类失败：", err)
		return errors.New("恢复视频分类失败")
	}
	return nil
}