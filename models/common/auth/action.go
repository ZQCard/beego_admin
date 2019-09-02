package auth

import (
	"errors"
	"github.com/astaxie/beego/orm"
)

// 权限结构体
type Action struct {
	Id int
	Name string
	Method string
	Route string
}

func init()  {
	// 注册模型
	orm.RegisterModel(new(Action))
}

func (a *Action)TableName() string {
	return "auth_action"
}

// 根据路由查找id
func findActionByName(route string) (*Action, error) {
	o := orm.NewOrm()
	action := &Action{Route:route}
	err := o.Read(action, "route")
	if err != nil  && err != orm.ErrNoRows{
		return action, err
	}

	if err == orm.ErrNoRows {
		return action, errors.New("行为不存在")
	}
	return action, nil
}