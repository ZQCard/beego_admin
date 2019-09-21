package admin

import (
	"beego_admin/models"
	"errors"
	"github.com/astaxie/beego/logs"
	"github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type Documentation struct {
	models.Model
	CompanyDocumentationCategoryId int
	Name string
	Url string
	Description string
}

type DocumentationInfo struct {
	models.Model
	CompanyDocumentationId int
	Name string
	Url string
	ViewTimes int
	IsShowHomePage int
	IsOnSale int
	Description string
	Category string
}


func (Documentation)TableName() string {
	return "company_documentation"
}


func (documentation Documentation)Validate() error {
	return validation.ValidateStruct(&documentation,
		// 名称不得为空,且大小为1-20字
		validation.Field(
			&documentation.Name,
			validation.Required.Error("资料名称名称不得为空"),
			validation.Length(1, 60).Error("名称为1-20字")),
		validation.Field(
			&documentation.Url,
			validation.Required.Error("资料地址名称不得为空"),
			is.URL.Error("资料地址错误"),
		),
	)
}


// 资料列表
func (documentation *Documentation)List(page, pageSize int) (results []DocumentationInfo, totalCount int64) {
	models.DB.Model(&Documentation{}).Unscoped().Count(&totalCount)
	// 关联分类表查询信息
	err := models.DB.Table("company_documentation as d").
		Select("d.*, c.name as category").
		Joins("left join company_documentation_category as c on d.company_documentation_category_id = c.id").
		Limit(pageSize).
		Order("id").
		Offset(page - 1).
		Scan(&results).Error
	if err != nil{
		logs.Error("查询资料列表报错", err)
		return nil, 0
	}
	return
}

// 资料列表(根据分类)
func (documentation *Documentation)ListFront(page, pageSize int) (results []Documentation, totalCount int64) {
	models.DB.Model(&documentation).Where("company_documentation_category_id = ?", documentation.CompanyDocumentationCategoryId).Count(&totalCount)
	err := models.DB.Where("company_documentation_category_id = ?", documentation.CompanyDocumentationCategoryId).Offset((page - 1) * pageSize).Limit(pageSize).Find(&results).Error
	if err != nil{
		logs.Error("查询资料列表报错", err)
		return nil, 0
	}
	return
}

// 添加资料信息
func (documentation *Documentation)Create() (err error) {
	// 数据验证
	err = documentation.Validate()
	if err != nil {
		return err
	}
	// 判断用户名或者昵称未使用
	if !models.DB.Unscoped().Where("name = ?", documentation.Name).Find(&Documentation{}).RecordNotFound(){
		return errors.New("资料名称已经存在")
	}

	if err = models.DB.Create(&documentation).Error; err != nil {
		logs.Error("资料创建失败", err)
		return errors.New("信息保存失败")
	}
	return nil
}

// 更新资料信息
func (documentation *Documentation)Update() (err error) {
	// 数据验证
	err = documentation.Validate()
	if err != nil {
		return err
	}
	// 判断用户名或者昵称未使用
	if !models.DB.Unscoped().Select("name").Where("id <> ?", documentation.ID).Where("name = ?", documentation.Name).Find(&Documentation{}).RecordNotFound(){
		return errors.New("资料名称已经存在")
	}
	if err = models.DB.Select("company_documentation_category_id", "name", "url", "description").Save(&documentation).Error; err != nil {
		logs.Error("资料创建失败", err)
		return errors.New("信息保存失败")
	}
	return nil
}

// 删除资料
func (documentation *Documentation)Delete() (err error) {
	if err := models.DB.Delete(&documentation).Error; err != nil {
		logs.Error("删除资料失败：", err)
		return errors.New("删除失败")
	}
	return nil
}

// 恢复资料
func (documentation *Documentation)Recover() (err error) {
	documentation.DeletedAt = nil
	if err := models.DB.Unscoped().Model(&documentation).Update("deleted_at", nil).Error; err != nil {
		logs.Error("恢复资料失败：", err)
		return errors.New("恢复资料失败")
	}
	return nil
}

func (documentation *Documentation)Find() (v Documentation, err error) {
	err = models.DB.Unscoped().Where(documentation).First(&v).Error
	if err != nil {
		return v, errors.New("用户信息错误")
	}
	return v, nil
}
