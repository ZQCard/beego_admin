package admin

import (
	"beego_admin/models"
	"encoding/json"
	"errors"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/utils"
	"github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"strconv"
	"strings"
)

// 管理员表GORM
type Administrator struct {
	models.Model
	Username string
	Nickname string
	Password string
	Email string
}

func (Administrator)TableName() string {
	return "administrator"
}

func (administrator Administrator)Validate() error {
	return validation.ValidateStruct(&administrator,
		// 名称不得为空,且大小为1-20字
		validation.Field(
			&administrator.Username,
			validation.Required.Error("名称不得为空"),
			validation.Length(1, 20).Error("名称为1-20字")),
		// 邮箱验证
		validation.Field(&administrator.Email,
			validation.Required.Error("邮箱不得为空"),
			is.Email.Error("邮箱格式不正确"),
			),
	)
}


// 管理员列表
func AdministratorList(page, pageSize int) (administrator []Administrator, totalCount int64) {
	models.DB.Unscoped().Model(&administrator).Count(&totalCount)
	err := models.DB.Unscoped().Offset((page - 1) * pageSize).Limit(pageSize).Find(&administrator).Error
	if err != nil{
		logs.Error("查询管理员列表报错", err)
		return nil, 0
	}
	return
}

// 根据条件查找管理员信息
func (administrator *Administrator)FindAdministrator() (admin Administrator, err error) {
	err = models.DB.Where(administrator).First(&admin).Error
	if err != nil {
		return admin, errors.New("用户信息错误")
	}
	return admin, nil
}

// 添加管理员信息
func (administrator *Administrator)AdministratorCreate() (err error) {
	// 数据验证
	err = administrator.Validate()
	if err != nil {
		return err
	}
	// 判断用户名或者昵称未使用
	if !models.DB.Unscoped().Where("username = ?", administrator.Username).Or("nickname = ?", administrator.Nickname).Find(&Administrator{}).RecordNotFound(){
		return errors.New("用户名或昵称已经存在")
	}

	// 查看密码
	if administrator.Password == "" {
		return errors.New("密码不得为空")
	}

	if err = models.DB.Create(&administrator).Error; err != nil {
		logs.Error("管理员信息创建失败", err)
		return errors.New("信息保存失败")
	}
	return nil
}

// 更新管理员信息
func (administrator *Administrator)AdministratorUpdate() (err error) {
	// 数据验证
	err = administrator.Validate()
	if err != nil {
		return err
	}
	// 判断用户名或者昵称未使用
	row1 := models.DB.Unscoped().Where("id <> ?", administrator.ID).Where("username = ?", administrator.Username).RowsAffected
	row2 := models.DB.Unscoped().Where("id <> ?", administrator.ID).Where("nickname = ?", administrator.Nickname).RowsAffected
	if row1 != 0 || row2 != 0{
		return errors.New("用户名或昵称已经存在")
	}

	
	if err = models.DB.Save(&administrator).Error; err != nil {
		logs.Error("管理员信息保存失败", err)
		return errors.New("信息保存失败")
	}
	return nil
}

// 删除管理员
func (administrator *Administrator)AdministratorDelete() (err error) {
	if administrator.Model.ID == 1{
		logs.Error("试图删除1号超级管理员")
		return errors.New("删除失败")
	}
	if err := models.DB.Delete(&administrator).Error; err != nil {
		logs.Error("删除管理员失败：", err)
		return errors.New("删除失败")
	}
	return nil
}

// 恢复管理员
func (administrator *Administrator)AdministratorRecover() (err error) {
	administrator.DeletedAt = nil
	if err := models.DB.Unscoped().Model(&administrator).Update("deleted_at", nil).Error; err != nil {
		logs.Error("恢复管理员失败：", err)
		return errors.New("恢复管理员失败")
	}
	return nil
}

// 管理员拥有的用户列表
func (administrator *Administrator)AdministratorRoleList() (map[string]interface{}, error) {
	// 查询所有角色
	var roleAll []Role
	err := models.DB.Find(&roleAll).Error
	if err != nil {
		logs.Error("查询角色失败：", err)
		return nil, errors.New("查询错误")
	}

	// 查询该用户已经拥有的角色
	rows, err := models.DB.Raw("SELECT role_id FROM auth_administrator_role WHERE administrator_id = ?", administrator.ID).Rows()
	defer rows.Close()
	if err != nil{
		logs.Error("查询角色失败：", err)
		return nil, errors.New("查询错误")
	}
	var roleHasIds []int
	var temp int
	for rows.Next() {
		if err = rows.Scan(&temp); err != nil {
			logs.Error("查询角色失败：", err)
			return nil, errors.New("查询错误")
		}
		roleHasIds = append(roleHasIds, temp)
	}
	data := make(map[string]interface{})
	data["all"] = roleAll
	data["has"] = roleHasIds
	return data, nil
}

// 授予管理员用户
func (administrator *Administrator)AssignRole(roleIds []int) error {
	if models.DB.Unscoped().Where(administrator).First(&Administrator{}).RecordNotFound(){
		return errors.New("管理员不存在")
	}

	if len(roleIds) > 0{
		num := models.DB.Where("id in (?)", roleIds).Find(&Role{}).RowsAffected
		if num != int64(len(roleIds)){
			return errors.New("角色不存在")
		}
	}

	// 事务操作 先删除该用户的所有权限  然后赋值
	tx := models.DB.Begin()
	delNum := tx.Exec("DELETE FROM auth_administrator_role WHERE administrator_id = ?", administrator.ID).RowsAffected
	var addNum int64 = 0
	if len(roleIds) > 0 {
		sql := "INSERT INTO auth_administrator_role(administrator_id, role_id) VALUES"
		for _, id := range roleIds {
			sql += "(" + strconv.Itoa(administrator.ID) + "," + strconv.Itoa(id) + "),"
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

// 获取管理员角色列表(用户结点)
func (administrator *Administrator)AuthList() (authList map[string][]string, err error) {
	sql := `SELECT act.method,act.route
	FROM auth_administrator_role a
	INNER JOIN auth_role r ON a.role_id = r.id
	INNER JOIN auth_role_permission rp ON r.id = rp.role_id
	INNER JOIN auth_permission p ON p.id = rp.permission_id
	INNER JOIN auth_permission_action pa ON p.id = pa.permission_id
	INNER JOIN auth_action act ON pa.action_id = act.id
	WHERE administrator_id = ?`
	// 查询该用户已经拥有的角色
	rows, err := models.DB.Raw(sql, administrator.ID).Rows()
	defer rows.Close()
	if err != nil{
		logs.Error("查询当前用户行为列表失败：", err)
		return nil, errors.New("查询错误")
	}
	// 初始化map
	authList = make(map[string][]string)

	var methodTemp string
	var routeTemp string
	for rows.Next() {
		err = rows.Scan(&methodTemp, &routeTemp)
		if err != nil{
			logs.Error("查询当前用户行为列表解析参数失败：", err)
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

// 获取管理员菜单列表
func (administrator *Administrator)MenuList(roleRoute []string) string {
	m := Menu{Pid:0}
	treeList := m.MenuList(roleRoute)
	menus,err := json.Marshal(treeList)
	if err != nil {
		logs.Error("删除管理员数据失败：", err)
		return ""
	}
	return string(menus)
}

