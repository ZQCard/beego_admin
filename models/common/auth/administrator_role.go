package auth

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/utils"
)

// 权限结构体
type AdministratorRole struct {
	Id int
	AdministratorId int
	RoleId int
}


func init()  {
	// 注册模型
	orm.RegisterModel(new(AdministratorRole))
}

func (ar *AdministratorRole)TableName() string {
	return "auth_administrator_role"
}

// 管理员拥有的角色列表
func AdministratorRoleList(administratorId int) (map[string][]Role, error) {
	// 所有行为
	roleAll,_ := RoleList(0, 3000)
	o := orm.NewOrm()
	// 已拥有行为id
	type RoleId struct {
		RoleId int
	}
	var roleIds []RoleId
	_,err := o.Raw("SELECT role_id FROM auth_administrator_role WHERE administrator_id = ?", administratorId).QueryRows(&roleIds)
	if err != nil {
		return nil, err
	}
	var roleIdArray []interface{}
	for _, role := range roleIds {
		roleIdArray = append(roleIdArray, role.RoleId)
	}

	// 分开处理已有行为和未有行为
	var roleHas []Role
	var roleHasNot []Role
	for _, role := range roleAll{
		// 判断id是否已经拥有
		if utils.InSliceIface(role.Id, roleIdArray) {
			roleHas = append(roleHas, role)
		} else {
			roleHasNot = append(roleHasNot, role)
		}
	}
	data := make(map[string][]Role)
	data["has"] = roleHas
	data["not"] = roleHasNot
	return data, nil
}

func AdministratorRoleAssign(administratorId int, roleIds []int) error {
	o := orm.NewOrm()
	err := o.Begin()
	if err != nil {
		return err
	}

	_, err = o.Raw("DELETE FROM auth_administrator_role WHERE administrator_id = ?", administratorId).Exec()

	if err != nil {
		o.Rollback()
		return errors.New("数据处理失败")
	}

	var administratorRole []AdministratorRole
	for _, roleId := range roleIds{
		administratorRole = append(administratorRole, AdministratorRole{AdministratorId:administratorId, RoleId:roleId})
	}
	// 排除为空的情况
	if len(roleIds) != 0{
		_, err = o.InsertMulti(len(roleIds), administratorRole)
		if err != nil && err != orm.ErrArgs{
			o.Rollback()
			return errors.New("数据处理失败")
		}
	}
	o.Commit()
	return  nil
}