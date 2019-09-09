package admin

import (
	"beego_admin/models"
	"errors"
	"fmt"
	"github.com/astaxie/beego/logs"
)

type User struct {
	models.Model
	UnionId string
	WeixinOpenId string
	Username string
	Nickname string
	Sex string
	Province string
	City string
	Headimgurl string
	Mobile string
	Email string
	LoginTimes int
}

func (User)TableName() string {
	return "company_user"
}

// 留言板列表
func (user *User)List(page, pageSize int) (users []User, totalCount int64) {
	fmt.Println(users)
	models.DB.Unscoped().Model(&users).Count(&totalCount)
	err := models.DB.Unscoped().Offset((page - 1) * pageSize).Limit(pageSize).Find(&users).Error
	if err != nil{
		logs.Error("查询用户列表报错", err)
		return nil, 0
	}
	return
}


// 删除用户
func (user *User)Delete() (err error) {
	if err := models.DB.Delete(&user).Error; err != nil {
		logs.Error("删除用户失败：", err)
		return errors.New("删除失败")
	}
	return nil
}

// 恢复用户
func (user *User)Recover() (err error) {
	user.DeletedAt = nil
	if err := models.DB.Unscoped().Model(&user).Update("deleted_at", nil).Error; err != nil {
		logs.Error("恢复用户失败：", err)
		return errors.New("恢复用户失败")
	}
	return nil
}