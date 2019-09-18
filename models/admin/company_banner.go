package admin

import (
	"beego_admin/models"
	"errors"
	"github.com/astaxie/beego/logs"
	"github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type Banner struct {
	models.Model
	Name string
	Url string
	Sort int
	IsShow int
}

func (Banner)TableName() string {
	return "company_banner"
}


func (banner Banner)Validate() error {
	return validation.ValidateStruct(&banner,
		// 名称不得为空,且大小为1-20字
		validation.Field(
			&banner.Name,
			validation.Required.Error("轮播图名称不得为空"),
			validation.Length(1, 60).Error("名称为1-20字")),
		// 路由
		validation.Field(
			&banner.Sort,
			validation.Required.Error("排序不得为空"),
		),
		
		// url
		validation.Field(
			&banner.Url,
			validation.Required.Error("轮播图地址名称不得为空"),
			is.URL.Error("轮播图地址错误"),
		),
	)
}

// 轮播图列表
func (banner *Banner)List(page, pageSize int) (banners []Banner, totalCount int64) {
	models.DB.Model(&banners).Unscoped().Count(&totalCount)
	err := models.DB.Unscoped().Offset((page - 1) * pageSize).Limit(pageSize).Find(&banners).Error
	if err != nil{
		logs.Error("查询轮播图列表报错", err)
		return nil, 0
	}
	return
}

// 前台轮播图列表
func (banner *Banner)CompanyList() (banners []Banner) {
	err := models.DB.Order("sort").Find(&banners).Error
	if err != nil{
		logs.Error("查询轮播图列表报错", err)
		return nil
	}
	return
}

// 添加轮播图信息
func (banner *Banner)Create() (err error) {
	// 数据验证
	err = banner.Validate()
	if err != nil {
		return err
	}
	// 判断名称未使用
	if !models.DB.Unscoped().Where("name = ?", banner.Name).Find(&Banner{}).RecordNotFound(){
		return errors.New("轮播图名称已经存在")
	}

	if err = models.DB.Create(&banner).Error; err != nil {
		logs.Error("轮播图创建失败", err)
		return errors.New("信息保存失败")
	}
	return nil
}

// 更新轮播图信息
func (banner *Banner)Update() (err error) {
	// 数据验证
	err = banner.Validate()
	if err != nil {
		return err
	}
	// 判断名称未使用
	if !models.DB.Unscoped().Select("name").Where("id <> ?", banner.ID).Where("name = ?", banner.Name).Find(&Banner{}).RecordNotFound(){
		return errors.New("轮播图名称已经存在")
	}

	if err = models.DB.Select("name", "url", "is_show", "sort").Save(&banner).Error; err != nil {
		logs.Error("轮播图创建失败", err)
		return errors.New("信息保存失败")
	}
	return nil
}

// 删除轮播图
func (banner *Banner)Delete() (err error) {
	if err := models.DB.Delete(&banner).Error; err != nil {
		logs.Error("删除轮播图失败：", err)
		return errors.New("删除失败")
	}
	return nil
}

// 恢复轮播图
func (banner *Banner)Recover() (err error) {
	banner.DeletedAt = nil
	if err := models.DB.Unscoped().Model(&banner).Update("deleted_at", nil).Error; err != nil {
		logs.Error("恢复轮播图失败：", err)
		return errors.New("恢复轮播图失败")
	}
	return nil
}