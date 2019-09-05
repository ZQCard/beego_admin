package admin

import "beego_admin/models"

type Documentation struct {
	models.Model
	CompanyDocumentationCategoryId int
	Name string
	Url string
	Description string
}

func (Documentation)TableName() string {
	return "company_documentation"
}
