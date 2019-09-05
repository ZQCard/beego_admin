package admin

import (
	"beego_admin/models"
	"errors"
	"github.com/astaxie/beego/logs"
	"github.com/go-ozzo/ozzo-validation"
)

type DocumentationCategory struct {
	models.Model
	Name string		`json:"name"`
	Sort int 		`json:"sort"`
}

func (DocumentationCategory)TableName() string {
	return "company_documentation_category"
}


func (category DocumentationCategory)Validate() error {
	return validation.ValidateStruct(&category,
		// 名称不得为空,且大小为1-20字
		validation.Field(
			&category.Name,
			validation.Required.Error("资料分类名称不得为空"),
			validation.Length(1, 60).Error("名称为1-20字")),
		// 路由
		validation.Field(
			&category.Sort,
			validation.Required.Error("排序不得为空"),
		),
	)
}

// 资料分类列表
func (category *DocumentationCategory)List(page, pageSize int) (categories []DocumentationCategory, totalCount int64) {
	models.DB.Model(&categories).Unscoped().Count(&totalCount)
	err := models.DB.Unscoped().Offset((page - 1) * pageSize).Limit(pageSize).Find(&categories).Error
	if err != nil{
		logs.Error("查询资料分类列表报错", err)
		return nil, 0
	}
	return
}

// 添加资料分类信息
func (category *DocumentationCategory)Create() (err error) {
	// 数据验证
	err = category.Validate()
	if err != nil {
		return err
	}
	// 判断用户名或者昵称未使用
	if !models.DB.Unscoped().Where("name = ?", category.Name).Find(&DocumentationCategory{}).RecordNotFound(){
		return errors.New("资料分类名称已经存在")
	}

	if err = models.DB.Create(&category).Error; err != nil {
		logs.Error("资料分类创建失败", err)
		return errors.New("信息保存失败")
	}
	return nil
}

// 更新资料分类信息
func (category *DocumentationCategory)Update() (err error) {
	// 数据验证
	err = category.Validate()
	if err != nil {
		return err
	}
	// 判断用户名或者昵称未使用
	if !models.DB.Unscoped().Where("id <> ?", category.ID).Where("name = ?", category.Name).Find(&DocumentationCategory{}).RecordNotFound(){
		return errors.New("资料分类名称已经存在")
	}

	if err = models.DB.Select("name", "sort").Save(&category).Error; err != nil {
		logs.Error("资料分类创建失败", err)
		return errors.New("信息保存失败")
	}
	return nil
}

// 删除资料分类
func (category *DocumentationCategory)Delete() (err error) {
	if err := models.DB.Delete(&category).Error; err != nil {
		logs.Error("删除资料分类失败：", err)
		return errors.New("删除失败")
	}
	return nil
}

// 恢复资料分类
func (category *DocumentationCategory)Recover() (err error) {
	category.DeletedAt = nil
	if err := models.DB.Unscoped().Model(&category).Update("deleted_at", nil).Error; err != nil {
		logs.Error("恢复资料分类失败：", err)
		return errors.New("恢复资料分类失败")
	}
	return nil
}