package auth

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/utils"
)

// 权限结构体
type RolePermission struct {
	Id int
	RoleId int
	PermissionId int
}


func init()  {
	// 注册模型
	orm.RegisterModel(new(RolePermission))
}

func (pa *RolePermission)TableName() string {
	return "auth_role_permission"
}

// 权限拥有的行为列表
func RolePermissionList(roleId int) (map[string][]Permission, error) {
	// 所有行为
	permissionAll,_ := PermissionList(0, 3000)
	o := orm.NewOrm()
	// 已拥有权限id
	type PermissionId struct {
		PermissionId int
	}
	var permissionIds []PermissionId
	_,err := o.Raw("SELECT permission_id FROM auth_role_permission WHERE role_id = ?", roleId).QueryRows(&permissionIds)
	if err != nil {
		return nil, err
	}
	var permissionIdArray []interface{}
	for _, permission := range permissionIds {
		permissionIdArray = append(permissionIdArray, permission.PermissionId)
	}

	// 分开处理已有行为和未有行为
	var permissionHas []Permission
	var permissionHasNot []Permission
	for _, permission := range permissionAll{
		// 判断id是否已经拥有
		if utils.InSliceIface(permission.Id, permissionIdArray) {
			permissionHas = append(permissionHas, permission)
		} else {
			permissionHasNot = append(permissionHasNot, permission)
		}
	}
	data := make(map[string][]Permission)
	data["has"] = permissionHas
	data["not"] = permissionHasNot
	return data, nil
}

func RolePermissionAssign(roleId int, permissionIds []int) error {
	o := orm.NewOrm()
	err := o.Begin()
	if err != nil {
		return err
	}

	_, err = o.Raw("DELETE FROM auth_role_permission WHERE role_id = ?", roleId).Exec()

	if err != nil {
		o.Rollback()
		return errors.New("数据处理失败")
	}

	var rolePermission []RolePermission
	for _, permissionId := range permissionIds{
		rolePermission = append(rolePermission, RolePermission{RoleId:roleId, PermissionId:permissionId})
	}
	// 排除为空的情况
	if len(permissionIds) != 0{
		_, err = o.InsertMulti(len(permissionIds), rolePermission)
		if err != nil && err != orm.ErrArgs{
			o.Rollback()
			return errors.New("数据处理失败")
		}
	}
	o.Commit()
	return  nil
}