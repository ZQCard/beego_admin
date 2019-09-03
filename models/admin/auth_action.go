package admin

import (
	"beego_admin/models"
	"errors"
	"github.com/astaxie/beego/logs"
	"github.com/go-ozzo/ozzo-validation"
	"regexp"
)

// 行为结构体
type Action struct {
	ID int
	Name string
	Method string
	Route string
}

func (Action)TableName() string {
	return "auth_action"
}

func (action Action)Validate() error {

	return validation.ValidateStruct(&action,
		// 名称不得为空,且大小为1-20字
		validation.Field(
			&action.Name,
			validation.Required.Error("行为名称不得为空"),
			validation.Length(1, 60).Error("名称为1-20字")),

		// 请求方式必须是GET,POST,PUT,PATCH,DELETE,OPTIONS中的一种
		validation.Field(
			&action.Method,
			validation.Required.Error("请求方式不得为空"),
			validation.In("GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS").Error("请求方式不合法"),
			),
		// 路由
		validation.Field(
			&action.Route,
			validation.Required.Error("路由不得为空"),
			// 必须以‘/’开头
			validation.Match(regexp.MustCompile("^/")).Error("路由必须以‘/’开头"),
		),
	)
}

// 行为列表
func ActionList(page, pageSize int) (actions []Action, totalCount int64) {
	models.DB.Model(&actions).Count(&totalCount)
	err := models.DB.Offset((page - 1) * pageSize).Limit(pageSize).Find(&actions).Error
	if err != nil{
		logs.Error("查询行为列表报错", err)
		return nil, 0
	}
	return
}

// 添加行为信息
func (action *Action)ActionCreate() (err error) {
	// 数据验证
	err = action.Validate()
	if err != nil {
		return err
	}
	// 判断用户名或者昵称未使用
	if !models.DB.Where("name = ?", action.Name).Find(&Action{}).RecordNotFound(){
		return errors.New("行为名称已经存在")
	}

	if err = models.DB.Create(&action).Error; err != nil {
		logs.Error("行为创建失败", err)
		return errors.New("信息保存失败")
	}
	return nil
}

// 添加行为信息
func (action *Action)ActionUpdate() (err error) {
	// 数据验证
	err = action.Validate()
	if err != nil {
		return err
	}
	// 判断用户名或者昵称未使用
	if !models.DB.Where("id <> ?", action.ID).Where("name = ?", action.Name).Find(&Action{}).RecordNotFound(){
		return errors.New("行为名称已经存在")
	}

	if err = models.DB.Save(&action).Error; err != nil {
		logs.Error("行为创建失败", err)
		return errors.New("信息保存失败")
	}
	return nil
}

// 删除行为
func (action *Action)ActionDelete() (err error) {
	if err := models.DB.Delete(&action).Error; err != nil {
		logs.Error("删除行为失败：", err)
		return errors.New("删除失败")
	}
	return nil
}

// 根据条件查找管理员信息
func (action *Action)FindAction() (act Action, err error) {
	err = models.DB.Where(action).First(&act).Error
	if err != nil {
		return act, errors.New("用户信息错误")
	}
	return act, nil
}
