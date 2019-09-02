package admin

import (
	"beego_admin/models"
	"errors"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/utils"
	"github.com/go-ozzo/ozzo-validation"
	"strconv"
	"strings"
)

// 权限结构体
type Permission struct {
	ID int
	Name string
}


func (Permission)TableName() string {
	return "auth_permission"
}

func (permission Permission)Validate() error {

	return validation.ValidateStruct(&permission,
		// 名称不得为空,且大小为1-20字
		validation.Field(
			&permission.Name,
			validation.Required.Error("权限名称不得为空"),
			validation.Length(1, 60).Error("名称为1-20字")),
	)
}

// 权限列表
func PermissionList(page, pageSize int) (permissions []Permission, totalCount int64) {
	models.DB.Model(&permissions).Count(&totalCount)
	err := models.DB.Offset((page - 1) * pageSize).Limit(pageSize).Find(&permissions).Error
	if err != nil{
		logs.Error("查询权限列表报错", err)
		return nil, 0
	}
	return
}

// 添加权限信息
func (permission *Permission)PermissionCreate() (err error) {
	// 数据验证
	err = permission.Validate()
	if err != nil {
		return err
	}
	// 判断用户名或者昵称未使用
	rowAffected := models.DB.Where("name = ?", permission.Name).Find(&Permission{}).RowsAffected
	if rowAffected != 0{
		return errors.New("权限名称已经存在")
	}

	if err = models.DB.Create(&permission).Error; err != nil {
		logs.Error("权限创建失败", err)
		return errors.New("信息保存失败")
	}
	return nil
}

// 添加权限信息
func (permission *Permission)PermissionUpdate() (err error) {
	// 数据验证
	err = permission.Validate()
	if err != nil {
		return err
	}
	// 判断用户名或者昵称未使用
	rowAffected := models.DB.Where("id <> ?", permission.ID).Where("name = ?", permission.Name).Find(&Permission{}).RowsAffected
	if rowAffected != 0{
		return errors.New("权限名称已经存在")
	}

	if err = models.DB.Save(&permission).Error; err != nil {
		logs.Error("权限创建失败", err)
		return errors.New("信息保存失败")
	}
	return nil
}

// 删除权限
func (permission *Permission)PermissionDelete() (err error) {
	if err := models.DB.Delete(&permission).Error; err != nil {
		logs.Error("删除权限失败：", err)
		return errors.New("删除失败")
	}
	return nil
}

// 权限拥有的行为列表
func (permission *Permission)PermissionActionList() (map[string][]Action, error) {
	// 所有行为
	actionAll,_ := ActionList(0, 3000)
	o := orm.NewOrm()
	// 已拥有行为id
	type ActionId struct {
		ActionId int
	}
	var actionIds []ActionId
	_,err := o.Raw("SELECT action_id FROM auth_permission_action WHERE permission_id = ?", permission.ID).QueryRows(&actionIds)
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
		if utils.InSliceIface(action.ID, actionIdArray) {
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

func (permission *Permission)AssignAction(actionIds []int) error {
	if models.DB.Where(permission).First(&Permission{}).RecordNotFound(){
		return errors.New("权限不存在")
	}

	num := models.DB.Where("id in (?)", actionIds).Find(&Action{}).RowsAffected
	if num != int64(len(actionIds)){
		return errors.New("行为不存在")
	}

	// 事务操作 先删除该权限的所有行为  然后赋值
	tx := models.DB.Begin()
	delNum := tx.Delete(&permission).RowsAffected
	sql := "INSERT INTO auth_permission_action(permission_id, action_id) VALUES"
	for _,id := range actionIds {
		fmt.Println(permission.ID)
		fmt.Println(id)
		sql += "("+ strconv.Itoa(permission.ID) + "," + strconv.Itoa(id) + "),"
	}
	sql = strings.Trim(sql, ",")
	addNum := tx.Exec(sql).RowsAffected
	if delNum >0 || addNum > 0{
		tx.Commit()
		return nil
	}else {
		tx.Rollback()
		return errors.New("操作失败")
	}
}

