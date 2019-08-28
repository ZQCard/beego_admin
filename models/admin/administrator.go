package admin

import (
	"beego_admin/models"
	"beego_admin/models/common/auth"
	"beego_admin/utils"
	"encoding/json"
	"errors"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	utils2 "github.com/astaxie/beego/utils"
	beegoValidation "github.com/astaxie/beego/validation"
	"github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"time"
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

// GORM根据条件查找管理员信息
func (administrator *AdministratorGORM)FindAdministratorGORM() (admin AdministratorGORM, err error) {
	err = models.DB.Where(administrator).First(&admin).Error
	if err != nil {
		return admin, err
	}
	return admin, nil
}

// GORM更新管理员信息
func (administrator *AdministratorGORM)UpdateAdministratorGORM() (err error) {
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



func init()  {
	// 注册模型(注册模型必须在引导(new ORM())之前运行)
	orm.RegisterModel(new(Administrator))
}

// 数据验证
func (administrator *Administrator)Valid(v *beegoValidation.Validation){

	if administrator.Username == "" {
		v.SetError("Username", "用户名不得为空")
	}

	if administrator.Nickname == "" {
		v.SetError("Nickname", "昵称不得为空")
	}

	if len(administrator.Username) < 1 || len(administrator.Username) > 60 {
		v.SetError("Username", "用户名长度为1-20个字")
	}

	if len(administrator.Nickname) < 1 || len(administrator.Nickname) > 60 {
		v.SetError("Username", "用户名长度为1-20个字")
	}

	if len(administrator.Password) < 6 {
		v.SetError("Password", "密码不得小于6位")
	}

	if !utils.VerifyEmailFormat(administrator.Email) {
		v.SetError("Email", "邮箱格式错误")
	}
}

// 管理员列表
func AdministratorList(page, pageSize int) (administrator []Administrator, totalCount int64) {
	o := orm.NewOrm()
	totalCount, err := o.QueryTable("administrator").Count()
	_,err = o.QueryTable("administrator").Limit(pageSize, (page - 1) * pageSize).All(&administrator)
	if err != nil {
		logs.Error("查询管理员报错", err)
		return nil, 0
	}
	return
}

// 增加管理员
func (administrator *Administrator)AddAdministrator() (bool, error) {
	// 数据验证
	valid := beegoValidation.Validation{}
	b, err := valid.Valid(administrator)
	if err != nil {
		logs.Error("验证管理员数据失败：", err)
		return false, errors.New("系统错误")
	}

	if !b {
		for _, err := range valid.Errors {
			return false, err
		}
	}

	o := orm.NewOrm()
	administrator.Password = utils.GenerateMD5String(administrator.Password)
	_, err = o.Insert(administrator)
	if err != nil {
		logs.Error("插入管理员数据失败：", err)
		return false, errors.New("添加管理员失败")
	}
	return true, nil
}

// 根据用户名密码查找管理员信息
func (administrator *Administrator)FindAdministratorByUsername() (*Administrator, error) {
	o := orm.NewOrm()
	err := o.QueryTable("administrator").Filter("username", administrator.Username).Filter("password",  administrator.Password).One(administrator)
	if err == orm.ErrNoRows {
		return administrator, errors.New("用户或密码错误")
	}

	if administrator.DeletedAt != 0 {
		return administrator, errors.New("该用户已冻结")
	}
	return administrator, nil
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

// 更新管理员
func (administrator *Administrator)UpdateAdministrator() (bool, error) {
	o := orm.NewOrm()
	// 数据验证
	valid := beegoValidation.Validation{}
	b, err := valid.Valid(administrator)
	if err != nil {
		logs.Error("验证管理员数据失败：", err)
		return false, errors.New("系统错误")
	}

	if !b {
		for _, err := range valid.Errors {
			return false, err
		}
	}

	// 判断用户名或者昵称是否已经存在
	var ad Administrator
	err = o.Raw("SELECT id FROM administrator WHERE id != ? AND (username = ? OR nickname = ?)", administrator.Id, administrator.Username, administrator.Nickname).QueryRow(&ad)
	if err != orm.ErrNoRows{
		return false, errors.New("用户名或昵称已存在")
	}
	if _, err := o.Update(administrator); err != nil {
		return false, err
	}
	return true, nil
}

// 删除管理员
func (administrator *Administrator)DeleteAdministrator() bool {
	// 开启事务
	o := orm.NewOrm()
	err := o.Begin()
	if err != nil {
		logs.Error("删除管理员报错", err)
		return false
	}
	// 将管理员表deteled_at置为当前时间戳
	administrator.DeletedAt = utils.Int64ToInt(time.Now().Unix())

	_, err = o.Update(administrator, "DeletedAt")
	// 删除权限与用户的关联角色
	_, err2 := o.Raw("DELETE FROM auth_administrator_role WHERE administrator_id = ?", administrator.Id).Exec()
	if err != nil || err2 != nil{
		logs.Error("删除管理员报错", err, "   ",err2)
		o.Rollback()
		return false
	}
	o.Commit()
	return true
}

// 恢复管理员
func (administrator *Administrator)RecoverAdministrator() bool {
	// 开启事务
	o := orm.NewOrm()

	// 将管理员表deteled_at置为0
	administrator.DeletedAt = 0

	_, err := o.Update(administrator, "DeletedAt")

	if err != nil {
		return false
	}
	return true
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
