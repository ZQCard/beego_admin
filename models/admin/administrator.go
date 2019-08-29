package admin

import (
	"beego_admin/models"
	"beego_admin/models/common/auth"
	"encoding/json"
	"errors"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	utils2 "github.com/astaxie/beego/utils"
	"github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

// 管理员表
type Administrator struct {
	models.Model
	Username string
	Nickname string
	Password string
	Email string
}

// 管理员表GORM
type AdministratorGORM struct {
	models.ModelGORM
	Username string
	Nickname string
	Password string
	Email string
}

func (AdministratorGORM)TableName() string {
	return "administrator"
}

func (administrator AdministratorGORM)Validate() error {
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
func AdministratorList(page, pageSize int) (administrator []AdministratorGORM, totalCount int64) {
	models.DB.Unscoped().Model(&administrator).Count(&totalCount)
	err := models.DB.Unscoped().Offset((page - 1) * pageSize).Limit(pageSize).Find(&administrator).Error
	if err != nil{
		logs.Error("查询管理员列表报错", err)
		return nil, 0
	}
	return
}

// 根据条件查找管理员信息
func (administrator *AdministratorGORM)FindAdministrator() (admin AdministratorGORM, err error) {
	err = models.DB.Where(administrator).First(&admin).Error
	if err != nil {
		return admin, errors.New("用户信息错误")
	}
	return admin, nil
}

// 添加管理员信息
func (administrator *AdministratorGORM)AdministratorCreate() (err error) {
	// 数据验证
	err = administrator.Validate()
	if err != nil {
		return err
	}
	// 判断用户名或者昵称未使用
	rowAffected := models.DB.Unscoped().Where("username = ?", administrator.Username).Or("nickname = ?", administrator.Nickname).Find(&AdministratorGORM{}).RowsAffected
	if rowAffected != 0{
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
func (administrator *AdministratorGORM)AdministratorUpdate() (err error) {
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
func (administrator *AdministratorGORM)AdministratorDelete() (err error) {
	if administrator.ModelGORM.ID == 1{
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
func (administrator *AdministratorGORM)AdministratorRecover() (err error) {
	administrator.DeletedAt = nil
	if err := models.DB.Unscoped().Model(&administrator).Update("deleted_at", nil).Error; err != nil {
		logs.Error("恢复管理员失败：", err)
		return errors.New("恢复管理员失败")
	}
	return nil
}



func init()  {
	// 注册模型(注册模型必须在引导(new ORM())之前运行)
	orm.RegisterModel(new(Administrator))
}

// 根据用户名密码查找管理员信息
func FindAdministratorById(id int) ( *Administrator, error) {
	o := orm.NewOrm()
	administrator := &Administrator{}
	err := o.QueryTable("administrator").Filter("id", id).One(administrator)
	if err == orm.ErrNoRows {
		return administrator, errors.New("用户不存在")
	}

	if administrator.DeletedAt != 0 {
		return administrator, errors.New("该用户已冻结")
	}
	return administrator, nil
}

// 授予管理员角色
func (administrator *Administrator)AssignRole(roleIds []int) error {
	o := orm.NewOrm()
	err := o.Read(administrator)
	if err == orm.ErrNoRows {
		return errors.New("权限不存在")
	}
	if err != nil {
		return err
	}
	err = auth.IsRole(roleIds)
	if err != nil {
		return err
	}

	err = auth.AdministratorRoleAssign(administrator.Id, roleIds)
	if err != nil {
		return err
	}
	return nil
}

// 获取管理员菜单列表
func (administrator *Administrator)MenuList(actionRoute []string) string {
	m := auth.Menu{Pid:0}
	treeList := m.MenuList(actionRoute)
	menus,err := json.Marshal(treeList)
	if err != nil {
		logs.Error("删除管理员数据失败：", err)
		return ""
	}
	return string(menus)
}

// 获取管理员行为列表(权限结点)
func (administrator *Administrator)AuthList() (authList map[string][]string, err error) {
	var lists []orm.ParamsList
	o := orm.NewOrm()
	num, err := o.Raw(`SELECT act.method,act.route
			FROM auth_administrator_role a
			INNER JOIN auth_role r ON a.role_id = r.id
			INNER JOIN auth_role_permission rp ON r.id = rp.role_id
			INNER JOIN auth_permission p ON p.id = rp.permission_id
			INNER JOIN auth_permission_action pa ON p.id = pa.permission_id
			INNER JOIN auth_action act ON pa.action_id = act.id
			WHERE administrator_id = ?`, administrator.Id).ValuesList(&lists)
	if err != nil {
		return nil, err
	}
	if num == 0{
		return nil, errors.New("无任何菜单权限")
	}
	// 初始化map
	authList = make(map[string][]string)
	// 将权限列表放到map中
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
