package admin

import (
	"beego_admin/models"
	"errors"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/utils"
	"github.com/go-ozzo/ozzo-validation"
	"strconv"
	"strings"
)

// 角色结构体
type Role struct {
	ID int
	Name string
}


func (Role)TableName() string {
	return "auth_role"
}

func (role Role)Validate() error {

	return validation.ValidateStruct(&role,
		// 名称不得为空,且大小为1-20字
		validation.Field(
			&role.Name,
			validation.Required.Error("角色名称不得为空"),
			validation.Length(1, 60).Error("名称为1-20字")),
	)
}

// 角色列表
func RoleList(page, pageSize int) (roles []Role, totalCount int64) {
	models.DB.Model(&roles).Count(&totalCount)
	err := models.DB.Offset((page - 1) * pageSize).Limit(pageSize).Find(&roles).Error
	if err != nil{
		logs.Error("查询角色列表报错", err)
		return nil, 0
	}
	return
}

// 添加角色信息
func (role *Role)RoleCreate() (err error) {
	// 数据验证
	err = role.Validate()
	if err != nil {
		return err
	}
	// 判断用户名或者昵称未使用
	if !models.DB.Where("name = ?", role.Name).Find(&Role{}).RecordNotFound(){
		return errors.New("角色名称已经存在")
	}

	if err = models.DB.Create(&role).Error; err != nil {
		logs.Error("角色创建失败", err)
		return errors.New("信息保存失败")
	}
	return nil
}

// 添加角色信息
func (role *Role)RoleUpdate() (err error) {
	// 数据验证
	err = role.Validate()
	if err != nil {
		return err
	}
	// 判断用户名或者昵称未使用
	if !models.DB.Where("id <> ?", role.ID).Where("name = ?", role.Name).Find(&Role{}).RecordNotFound(){
		return errors.New("角色名称已经存在")
	}

	if err = models.DB.Save(&role).Error; err != nil {
		logs.Error("角色创建失败", err)
		return errors.New("信息保存失败")
	}
	return nil
}

// 删除角色
func (role *Role)RoleDelete() (err error) {
	tx := models.DB.Begin()
	if err := tx.Delete(&role).Error; err != nil {
		logs.Error("删除角色失败：", err)
		return errors.New("删除失败")
	}
	// 删除角色下的权限
	tx.Exec("DELETE FROM auth_role_permission WHERE role_id = ?", role.ID)
	tx.Commit()
	return nil
}

// 角色拥有的权限列表
func (role *Role)RolePermissionList() (map[string][]Permission, error) {
	// 查询所有权限
	var permissionAll []Permission
	err := models.DB.Find(&permissionAll).Error
	if err != nil {
		logs.Error("查询权限失败：", err)
		return nil, errors.New("查询错误")
	}

	// 查询该权限已经拥有的权限
	rows, err := models.DB.Raw("SELECT permission_id FROM auth_role_permission WHERE role_id = ?", role.ID).Rows()
	defer rows.Close()
	if err != nil{
		logs.Error("查询权限失败：", err)
		return nil, errors.New("查询错误")
	}
	var permissionHasIds []int
	var temp int
	for rows.Next() {
		if err = rows.Scan(&temp); err != nil {
			logs.Error("查询权限失败：", err)
			return nil, errors.New("查询错误")
		}
		permissionHasIds = append(permissionHasIds, temp)
	}

	// 分开处理已有权限和未有权限
	var permissionHas []Permission
	var permissionHasNot []Permission

	for _, permission := range permissionAll{
		// 判断id是否已经拥有
		var flag = false
		for _,id := range permissionHasIds {
			if id == permission.ID{
				flag = true
			}
		}
		if flag {
			permissionHas = append(permissionHas, permission)
		}else {
			permissionHasNot = append(permissionHasNot, permission)
		}
	}
	data := make(map[string][]Permission)
	data["has"] = permissionHas
	data["not"] = permissionHasNot
	return data, nil
}

func (role *Role)AssignPermission(permissionIds []int) error {
	if models.DB.Where(role).First(&Role{}).RecordNotFound(){
		return errors.New("角色不存在")
	}
	if len(permissionIds) > 0 {
		num := models.DB.Where("id in (?)", permissionIds).Find(&Permission{}).RowsAffected
		if num != int64(len(permissionIds)) {
			return errors.New("权限不存在")
		}
	}
	// 事务操作 先删除该角色的所有权限  然后赋值
	tx := models.DB.Begin()
	delNum := tx.Exec("DELETE FROM auth_role_permission WHERE role_id = ?", role.ID).RowsAffected
	var addNum int64 = 0
	if len(permissionIds) > 0 {
		sql := "INSERT INTO auth_role_permission(role_id, permission_id) VALUES"
		for _,id := range permissionIds {
			sql += "("+ strconv.Itoa(role.ID) + "," + strconv.Itoa(id) + "),"
		}

		sql = strings.Trim(sql, ",")
		addNum = tx.Exec(sql).RowsAffected
	}

	if delNum > 0 || addNum > 0{
		tx.Commit()
		return nil
	}else {
		tx.Rollback()
		return errors.New("操作失败")
	}
}

func (role *Role)AuthList()(authList map[string][]string, err error) {
	sql := `SELECT act.method,act.route
			FROM auth_role r 
			INNER JOIN auth_role_permission rp ON r.id = rp.role_id
			INNER JOIN auth_permission p ON p.id = rp.permission_id
			INNER JOIN auth_permission_action pa ON p.id = pa.permission_id
			INNER JOIN auth_action act ON pa.action_id = act.id
			WHERE r.name = ?`
	// 查询该用户已经拥有的角色
	rows, err := models.DB.Raw(sql, role.Name).Rows()
	defer rows.Close()
	if err != nil{
		logs.Error("查询公共行为列表失败：", err)
		return nil, errors.New("查询错误")
	}
	// 初始化map
	authList = make(map[string][]string)

	var methodTemp string
	var routeTemp string
	for rows.Next() {
		err = rows.Scan(&methodTemp, &routeTemp)
		if err != nil{
			logs.Error("查询公共行为列表解析参数失败：", err)
			return nil, errors.New("查询错误")
		}
		if routes,ok := authList[methodTemp]; ok {
			// 如果请求方式中无该路由 则加入
			if !utils.InSlice(routeTemp, routes) {
				authList[methodTemp] = append(routes, routeTemp)
			}
		} else {
			// 无该请求方式
			authList[methodTemp] = []string{routeTemp}
		}
	}
	return authList, nil
}

