package demo

import "beego_admin/models"

type Excel struct {
	models.Model
	Name string
	Age int
	Birthday int
}

// 导入数据
func Import()  {
	
}

// 导出数据
func Export()  {
	
}
