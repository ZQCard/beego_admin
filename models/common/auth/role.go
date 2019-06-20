package auth

import (
	"beego_admin/utils"
	"errors"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	utils2 "github.com/astaxie/beego/utils"
	"github.com/astaxie/beego/validation"
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


// 数据验证
func (r *Role)Valid(v *validation.Validation){
	if r.Name == "" {
		v.SetError("Name", "角色名称不得为空")
	}
	if len(r.Name) < 1 || len(r.Name) > 60 {
		v.SetError("Name", "角色名称长度为1-20个字")
	}
}

// 角色列表
func RoleList(page, pageSize int) (roles []Role, totalCount int64) {
	o := orm.NewOrm()
	totalCount, err := o.QueryTable("auth_role").Count()
	_,err = o.QueryTable("auth_role").Limit(pageSize, (page - 1) * pageSize).All(&roles)
	if err != nil {
		logs.Error("读取菜单列表错误", err)
		return nil, 0
	}
	return
}

// 创建角色
func (r *Role)CreateRole() (bool, error) {
	// 检测角色名是否存在
	o := orm.NewOrm()

	// 数据验证
	valid := validation.Validation{}
	b, err := valid.Valid(r)
	if err != nil {
		logs.Error("创建角色数据验证错误", err)
		return false, errors.New("系统错误")
	}

	if !b {
		for _, err := range valid.Errors {
			return false, err
		}
	}

	if created, _, err := o.ReadOrCreate(r, "Name"); err == nil {
		if created {
			return true, nil
		}
	}
	return false, errors.New("角色已存在")
}

// 更新角色
func (r *Role)UpdateRole() (bool, error) {
	o := orm.NewOrm()
	// 数据验证
	valid := validation.Validation{}
	b, err := valid.Valid(r)
	if err != nil {
		logs.Error("更新角色数据验证错误", err)
		return false, errors.New("系统错误")
	}

	if !b {
		for _, err := range valid.Errors {
			return false, err
		}
	}
	// 判断角色名是否已经存在
	rolebak := &Role{Name:r.Name}
	err = o.Read(rolebak, "name")
	if err != nil  && err != orm.ErrNoRows{
		return false, err
	}
	if rolebak.Id != 0 && rolebak.Id != r.Id {
		return false, errors.New("角色名已存在")
	}
	// 更新
	if o.Read(r) == nil{
		r.Name = rolebak.Name
		num, err := o.Update(r)
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

// 删除角色
func (r *Role)DeleteRole() bool{
	// 开启事务
	o := orm.NewOrm()
	err := o.Begin()
	if err != nil {
		logs.Error("删除角色事务开启错误", err)
		return false
	}
	// 删除角色
	num, err := o.Delete(r)
	if err != nil {
		logs.Error("删除角色错误", err)
		return false
	}
	if num == 0 {
		o.Rollback()
		return false
	}
	// 删除角色与角色的关联
	_, err1 := o.Raw("DELETE FROM auth_administrator_role WHERE role_id = ?", r.Id).Exec()
	// 删除角色与用户的关联
	_, err2 := o.Raw("DELETE FROM auth_role_permission WHERE role_id = ?", r.Id).Exec()
	if err != nil || err1 != nil || err2 != nil{
		o.Rollback()
		return false
	}
	o.Commit()
	return true
}

// 授予角色行为
func (r *Role)AssignPermission(permissionIds []int) error {
	o := orm.NewOrm()
	err := o.Read(r)
	if err == orm.ErrNoRows {
		return errors.New("角色不存在")
	}
	if err != nil {
		return err
	}

	err = IsPermission(permissionIds)
	if err != nil {
		return err
	}

	err = RolePermissionAssign(r.Id, permissionIds)
	if err != nil {
		return err
	}
	return nil
}

// 根据id切片判断是否为合法role
func IsRole(ids []int) error {
	if len(ids) == 0{
		return nil
	}
	o := orm.NewOrm()
	num, err := o.QueryTable("auth_role").Filter("id__in", ids).Count()
	if err != nil {
		return err
	}
	if utils.Int64ToInt(num) != len(ids) {
		return errors.New("非法角色")
	}
	return nil
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