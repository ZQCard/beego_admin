package admin

import "beego_admin/models"

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
	return "company_documentation"
}
