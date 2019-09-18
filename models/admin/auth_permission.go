package admin

import (
	"beego_admin/models"
	"errors"
	"github.com/astaxie/beego/logs"
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
	if !models.DB.Where("name = ?", permission.Name).Find(&Permission{}).RecordNotFound(){
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
	if !models.DB.Where("id <> ?", permission.ID).Where("name = ?", permission.Name).Find(&Permission{}).RecordNotFound(){
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
	tx := models.DB.Begin()
	if err := tx.Delete(&permission).Error; err != nil {
		logs.Error("删除权限失败：", err)
		return errors.New("删除失败")
	}
	// 删除权限下的行为
	tx.Exec("DELETE FROM auth_permission_action WHERE permission_id = ?", permission.ID)
	tx.Commit()
	return nil
}

// 权限拥有的行为列表
func (permission *Permission)PermissionActionList() (map[string]interface{}, error) {
	// 查询所有行为
	var actionAll []Action
	err := models.DB.Find(&actionAll).Error
	if err != nil {
		logs.Error("查询行为失败：", err)
		return nil, errors.New("查询错误")
	}

	// 查询该权限已经拥有的行为
	rows, err := models.DB.Raw("SELECT action_id FROM auth_permission_action WHERE permission_id = ?", permission.ID).Rows()
	defer rows.Close()
	if err != nil{
		logs.Error("查询行为失败：", err)
		return nil, errors.New("查询错误")
	}

	var actionHasIds []int
	var temp int
	for rows.Next() {
		if err = rows.Scan(&temp); err != nil {
			logs.Error("查询行为失败：", err)
			return nil, errors.New("查询错误")
		}
		actionHasIds = append(actionHasIds, temp)
	}
	//
	//// 分开处理已有行为和未有行为
	//var actionHas []Action
	//var actionHasNot []Action
	//
	//for _, action := range actionAll{
	//	// 判断id是否已经拥有
	//	var flag = false
	//	for _,id := range actionHasIds {
	//		if id == action.ID{
	//			flag = true
	//		}
	//	}
	//	if flag {
	//		actionHas = append(actionHas, action)
	//	}else {
	//		actionHasNot = append(actionHasNot, action)
	//	}
	//}
	data := make(map[string]interface{})
	data["has"] = actionHasIds
	data["all"] = actionAll
	return data, nil
}

// 授予权限行为
func (permission *Permission)AssignAction(actionIds []int) error {
	if models.DB.Where(permission).First(&Permission{}).RecordNotFound(){
		return errors.New("权限不存在")
	}

	if len(actionIds) > 0{
		num := models.DB.Where("id in (?)", actionIds).Find(&Action{}).RowsAffected
		if num != int64(len(actionIds)){
			return errors.New("行为不存在")
		}
	}


	// 事务操作 先删除该权限的所有行为  然后赋值
	tx := models.DB.Begin()
	delNum := tx.Exec("DELETE FROM auth_permission_action WHERE permission_id = ?", permission.ID).RowsAffected
	var addNum int64 = 0
	if len(actionIds) > 0 {
		sql := "INSERT INTO auth_permission_action(permission_id, action_id) VALUES"
		for _, id := range actionIds {
			sql += "(" + strconv.Itoa(permission.ID) + "," + strconv.Itoa(id) + "),"
		}
		sql = strings.Trim(sql, ",")
		addNum = tx.Exec(sql).RowsAffected
	}
	if delNum >0 || addNum > 0{
		tx.Commit()
		return nil
	}else {
		tx.Rollback()
		return errors.New("操作失败")
	}
}

