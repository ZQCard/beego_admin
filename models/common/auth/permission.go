package auth

import (
	"beego_admin/utils"
	"errors"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
)

// 权限结构体
type Permission struct {
	Id int
	Name string
}


func init()  {
	// 注册模型
	orm.RegisterModel(new(Permission))
}


func (p *Permission)TableName() string {
	return "auth_permission"
}

// 数据验证
func (p *Permission)Valid(v *validation.Validation){
	if p.Name == "" {
		v.SetError("Name", "权限名称不得为空")
	}
	if len(p.Name) < 1 || len(p.Name) > 60 {
		v.SetError("Name", "权限名称长度为1-20个字")
	}
}

// 权限列表
func PermissionList(page, pageSize int) (permissions []Permission, totalCount int64) {
	o := orm.NewOrm()
	totalCount, err := o.QueryTable("auth_permission").Count()
	_,err = o.QueryTable("auth_permission").Limit(pageSize, (page - 1) * pageSize).All(&permissions)
	if err != nil {
		logs.Error("读取权限列表错误", err)
		return nil, 0
	}
	return
}

// 创建权限
func (p *Permission)CreatePermission() (bool, error) {
	// 检测权限名是否存在
	o := orm.NewOrm()

	// 数据验证
	valid := validation.Validation{}
	b, err := valid.Valid(p)
	if err != nil {
		logs.Error("权限数据验证错误", err)
		return false, nil
	}

	if !b {
		for _, err := range valid.Errors {
			return false, err
		}
	}
	
	
	if created, _, err := o.ReadOrCreate(p, "Name"); err == nil {
		if created {
			return true, nil
		}
	}
	return false, errors.New("权限已存在")
}

// 更新权限
func (p *Permission)UpdatePermission() (bool, error) {
	o := orm.NewOrm()

	// 数据验证
	valid := validation.Validation{}
	b, err := valid.Valid(p)
	if err != nil {
		logs.Error("权限数据验证错误", err)
		return false, nil
	}

	if !b {
		for _, err := range valid.Errors {
			return false, err
		}
	}
	permissionBak := &Permission{Name:p.Name}
	// 判断权限名是否已经存在
	err = o.Read(permissionBak, "name")
	if err != nil  && err != orm.ErrNoRows{
		return false, err
	}
	if permissionBak.Id != 0 && permissionBak.Id != p.Id {
		return false, errors.New("权限名已存在")
	}

	// 更新
	if o.Read(p) == nil{
		p.Name = permissionBak.Name
		num, err := o.Update(p)
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

// 删除权限
func (p *Permission)DeletePermission() bool {
	// 开启事务
	o := orm.NewOrm()
	err := o.Begin()
	if err != nil {
		logs.Error("删除权限开启事务错误", err)
		return false
	}
	// 删除权限
	num, err := o.Delete(p)
	if err != nil {
		logs.Error("删除权限错误", err)
		return false
	}
	if num == 0 {
		o.Rollback()
		return false
	}
	// 删除权限与权限的关联
	_, err1 := o.Raw("DELETE FROM auth_role_permission WHERE permission_id = ?", p.Id).Exec()
	// 删除操作与权限的关联
	_, err2 := o.Raw("DELETE FROM auth_permission_action WHERE permission_id = ?", p.Id).Exec()
	if err != nil || err1 != nil || err2 != nil {
		o.Rollback()
		return false
	}
	o.Commit()
	return true
}


// 根据id切片判断是否为合法permission
func IsPermission(ids []int) error {
	if len(ids) == 0{
		return nil
	}
	o := orm.NewOrm()
	num, err := o.QueryTable("auth_permission").Filter("id__in", ids).Count()
	if err != nil {
		return err
	}
	if utils.Int64ToInt(num) != len(ids) {
		return errors.New("非法权限")
	}
	return nil
}