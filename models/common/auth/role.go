package auth

import (
	"errors"
	"github.com/astaxie/beego/orm"
	utils2 "github.com/astaxie/beego/utils"
)

// 角色结构体
type Role struct {
	Id int
	Name string `valid:"Required"`
}

func (r *Role)TableName() string {
	return "auth_role"
}

func init()  {
	// 注册模型
	orm.RegisterModel(new(Role))
}
// 获取角色角色列表
func RoleAuthList(name string) (authList map[string][]string, err error) {
	var lists []orm.ParamsList
	o := orm.NewOrm()
	num, err := o.Raw(`SELECT act.method,act.route
			FROM auth_role r 
			INNER JOIN auth_role_permission rp ON r.id = rp.role_id
			INNER JOIN auth_permission p ON p.id = rp.permission_id
			INNER JOIN auth_permission_action pa ON p.id = pa.permission_id
			INNER JOIN auth_action act ON pa.action_id = act.id
			WHERE r.name = ?`, name).ValuesList(&lists)
	if err != nil {
		return nil, err
	}
	if num == 0{
		return nil, errors.New("无任何菜单角色")
	}
	// 初始化map
	authList = make(map[string][]string)
	// 将角色列表放到map中
	for _,v := range lists {
		// map中有该请求方式
		if method,ok := authList[v[0].(string)]; ok {
			// 如果请求方式中无该行为路由 则加入

			if !utils2.InSlice(v[1].(string),method) {
				authList[v[0].(string)] = append(method, v[1].(string))
			}
		} else {
			// 无该请求方式
			authList[v[0].(string)] = []string{v[1].(string)}
		}
	}
	return authList, err
}