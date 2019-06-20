package auth

import (
	"beego_admin/utils"
	"errors"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"strings"
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

// 数据验证
func (a *Action)Valid(v *validation.Validation){
	if len(a.Name) < 1 || len(a.Name) > 60 {
		v.SetError("Name", "权限名称长度为1-20个字")
	}

	if len(a.Method) < 1 || len(a.Method) > 50 {
		v.SetError("Method", "权限名称长度为1-50个字")
	}

	supportMethod := [6]string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	// 判断当前请求是否在允许请求内
	flag := false
	for _, method := range supportMethod{
		if method == strings.ToUpper(a.Method) {
			a.Method = strings.ToUpper(a.Method)
			flag = true
			break
		}
	}

	if !flag {
		v.SetError("Method", "请求方法必须在为合法HTTP请求")

	}

	if len(a.Route) < 1 || len(a.Route) > 50 {
		v.SetError("Route", "权限名称长度为1-50个字")
	}

	if !strings.HasPrefix(a.Route, "/") {
		v.SetError("Route", "路由必须以 / 开头")
	}
}

// 行为列表
func ActionList(page, pageSize int) (actions []Action, totalCount int64) {
	o := orm.NewOrm()
	totalCount, err := o.QueryTable("auth_action").Count()
	_,err = o.QueryTable("auth_action").Limit(pageSize, (page - 1) * pageSize).All(&actions)
	if err != nil {
		logs.Error("读取行为列表错误", err)
		return nil, 0
	}
	return
}

// 创建行为
func (a *Action)CreateAction() (bool, error) {
	// 检测角色名是否存在
	o := orm.NewOrm()
	// 数据验证
	valid := validation.Validation{}
	b, err := valid.Valid(a)
	if err != nil {
		logs.Error("创建行为错误", err)
		return false, errors.New("系统错误")
	}

	if !b {
		for _, err := range valid.Errors {
			return false, err
		}
	}
	if created, _, err := o.ReadOrCreate(a, "Name", "Method", "Route"); err == nil {
		if created {
			return true, nil
		}
	}
	return false, errors.New("权限已存在")
}

// 更新行为
func (a *Action)UpdateAction() (bool, error) {
	o := orm.NewOrm()

	// 数据验证
	valid := validation.Validation{}
	b, err := valid.Valid(a)
	if err != nil {
		logs.Error("创建行为数据验证错误", err)
		return false, errors.New("系统错误")
	}

	if !b {
		for _, err := range valid.Errors {
			return false, err
		}
	}
	actionBak := &Action{Name:a.Name, Method:a.Method, Route:a.Route}

	// 判断路由名是否已经存在
	err = o.Read(actionBak, "name", "method", "route")
	if err != nil  && err != orm.ErrNoRows{
		return false, err
	}
	if actionBak.Id != 0 && actionBak.Id != a.Id {
		return false, errors.New("路由已存在")
	}

	// 更新
	if o.Read(a) == nil{
		a.Name = actionBak.Name
		a.Method = actionBak.Method
		a.Route = actionBak.Route
		num, err := o.Update(a)
		if err != nil {
			return false, err
		}
		if num == 0{
			return  false, errors.New("更新失败")
		}
		return true, nil
	}
	return false, errors.New("数据不存在")
}

// 删除行为
func (a *Action)DeleteAction() bool {
	o := orm.NewOrm()

	// 删除角色
	num, err := o.Delete(a)
	if err != nil {
		logs.Error("删除角色错误", err)
		return false
	}
	if num == 0 {
		return false
	}
	return true
}

// 根据id切片判断是否为合法action
func IsAction(ids []int) error {
	if len(ids) == 0{
		 return nil
	}
	o := orm.NewOrm()
	num, err := o.QueryTable("auth_action").Filter("id__in", ids).Count()
	if err != nil {
		return err
	}
	if utils.Int64ToInt(num) != len(ids) {
		return errors.New("非法数据")
	}
	return nil
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