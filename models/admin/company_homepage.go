package admin

import (
	"beego_admin/models"
	"errors"
	"github.com/astaxie/beego/logs"
)

// 管理员表GORM
type Homepage struct {
	ID int
	Name string
	IsShow int
}

func (Homepage)TableName() string {
	return "company_homepage"
}

// 模块列表
func (homepage Homepage)ModuleList() (h []Homepage) {
	err := models.DB.Find(&h).Error
	if err != nil{
		logs.Error("查询首页模块列表报错", err)
		return nil
	}
	return
}

// 模块设置
func (homepage *Homepage)ModuleUpdate() (err error) {
	if err := models.DB.Unscoped().Model(&homepage).Update("is_show", homepage.IsShow).Error; err != nil {
		logs.Error("首页模块设置失败：", err)
		return errors.New("首页模块设置失败")
	}
	return nil
}


