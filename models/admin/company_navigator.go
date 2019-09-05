package admin

import (
	"beego_admin/models"
)

type Navigator struct {
	models.Model
	Pid int
	Name string
	Url string
	Sort int
}

func (Navigator)TableName() string {
	return "company_navigator"
}

