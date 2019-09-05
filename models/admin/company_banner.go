package admin

import "beego_admin/models"

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
