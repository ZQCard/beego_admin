package demo

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type Excel struct {
	Id int
	Name string
	Age int
	Birthday int64
}

func init()  {
	// 注册模型(注册模型必须在引导(new ORM())之前运行)
	orm.RegisterModel(new(Excel))
}

func (e *Excel)TableName() string {
	return "demo_excel"
}

// 导入数据到数据库
func Import(content []Excel) (num int64, err error) {
	o := orm.NewOrm()
	num, err =  o.InsertMulti(len(content), content)
	if err != nil {
		logs.Error("数据导入失败: ", err.Error())
		return 0, err
	}
	return num, nil
}

// 导出数据
func Export()  {
	
}
