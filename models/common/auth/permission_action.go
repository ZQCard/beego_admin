package auth

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/utils"
)

// 权限结构体
type PermissionAction struct {
	Id int
	PermissionId int
	ActionId int
}


func init()  {
	// 注册模型
	orm.RegisterModel(new(PermissionAction))
}

func (pa *PermissionAction)TableName() string {
	return "auth_permission_action"
}

// 权限拥有的行为列表
func PermissionActionList(permissionId int) (map[string][]Action, error) {
	// 所有行为
	actionAll,_ := ActionList(0, 3000)
	o := orm.NewOrm()
	// 已拥有行为id
	type ActionId struct {
		ActionId int
	}
	var actionIds []ActionId
	_,err := o.Raw("SELECT action_id FROM auth_permission_action WHERE permission_id = ?", permissionId).QueryRows(&actionIds)
	if err != nil {
		return nil, err
	}
	var actionIdArray []interface{}
	for _, action := range actionIds {
		actionIdArray = append(actionIdArray, action.ActionId)
	}

	// 分开处理已有行为和未有行为
	var actionHas []Action
	var actionHasNot []Action
	for _, action := range actionAll{
		// 判断id是否已经拥有
		if utils.InSliceIface(action.Id, actionIdArray) {
			actionHas = append(actionHas, action)
		} else {
			actionHasNot = append(actionHasNot, action)
		}
	}
	data := make(map[string][]Action)
	data["has"] = actionHas
	data["not"] = actionHasNot
	return data, nil
}

// 授予权限行为
func PermissionActionAssign(permissionId int, actionIds []int) error {
	o := orm.NewOrm()
	err := o.Begin()
	if err != nil {
		return err
	}

	_, err = o.Raw("DELETE FROM auth_permission_action WHERE permission_id = ?", permissionId).Exec()

	if err != nil {
		o.Rollback()
		return errors.New("数据处理失败")
	}

	var actionPermission []PermissionAction
	for _, actionId := range actionIds{
		actionPermission = append(actionPermission, PermissionAction{ActionId:actionId, PermissionId:permissionId})
	}
	// 排除为空的情况
	if len(actionIds) != 0{
		_, err = o.InsertMulti(len(actionIds), actionPermission)
		if err != nil && err != orm.ErrArgs{
			o.Rollback()
			return errors.New("数据处理失败")
		}
	}
	o.Commit()
	return  nil
}