package auth

import (
	"errors"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/utils"
	"github.com/astaxie/beego/validation"
)

type Menu struct {
	Id int			`json:"id"`
	Pid int			`json:"pid"`
	Name string 	`json:"name"`
	Sort int 		`json:"sort"`
	Route string	`json:"route"`
}

type TreeList struct {
	Id int			`json:"id"`
	Name string		`json:"name"`
	Pid int			`json:"pid"`
	Sort int 		`json:"sort"`
	Route string	`json:"route"`
	Children []*TreeList	`json:"children"`
}

var menuRoute []string

var authRoute []string

func init()  {
	// 注册模型
	orm.RegisterModel(new(Menu))
}

func (m *Menu)TableName() string {
	return "auth_menu"
}

// 数据验证
func (m *Menu)Valid(v *validation.Validation){
	if len(m.Name) < 1 || len(m.Name) > 60 {
		v.SetError("Name", "权限名称长度为1-20个字")
	}

	if m.Sort < 0 || m.Sort > 100 {
		v.SetError("Sort", "排序值为0-100之间")
	}

	// 检测路由是否存在
	if m.Route != "" {
		_, err := findActionByName(m.Route)
		if err != nil {
			v.SetError("Route", err.Error())
		}
	}
}

/**
菜单列表
 */
func (m *Menu)MenuList(route []string) []*TreeList{
	authRoute = route
	return m.getMenu(m.Pid)
}

/**
	递归获取树形菜单
 */
func (m *Menu)getMenu(pid int) []*TreeList {
	o := orm.NewOrm()
	var menu []Menu
	_,_ = o.QueryTable("auth_menu").Filter("pid", pid).OrderBy("sort").All(&menu)
	treeList := []*TreeList{}
	for _, v := range menu{
		if len(authRoute) != 0{
			// 根据当前菜单Pid判断是否与有路由在权限路由有交集，如果没有continue
			// 如果route为空,有下级  不为空,无下级
			if v.Route == "" {
				menuRoute = findActionByMenuId(v.Id)
			} else {
				tmpMenu := &Menu{Id: v.Id}
				o.Read(tmpMenu)
				menuRoute = []string{tmpMenu.Route}
			}
			flag := false
			for _,v := range menuRoute{
				if utils.InSlice(v, authRoute) {
					flag = true
					break
				}
			}
			// 重置menuRoute
			menuRoute = []string{}
			if flag == false{
				continue
			}
		}
		child := v.getMenu(v.Id)
		node := &TreeList{
			Id:v.Id,
			Name:v.Name,
			Sort:v.Sort,
			Route:v.Route,
			Pid:v.Pid,
		}
		node.Children = child
		treeList = append(treeList, node)
	}
	return treeList
}

// 根据id找到所有的action_route
func findActionByMenuId(id int) []string {
	o := orm.NewOrm()
	var menu []Menu
	_, _ = o.QueryTable("auth_menu").Filter("pid", id).All(&menu)
	for _, v := range menu {
		findActionByMenuId(v.Id)
		menuRoute = append(menuRoute, v.Route)
	}
	return menuRoute
}

// 创建菜单
func (m *Menu)CreateMenu() error {
	// 检测角色名是否存在
	o := orm.NewOrm()

	// 数据验证
	valid := validation.Validation{}
	b, err := valid.Valid(m)
	if err != nil {
		logs.Error("创建行为数据验证错误", err)
		return errors.New("系统错误")
	}

	if !b {
		for _, err := range valid.Errors {
			return err
		}
	}

	// 如果pid不存在,报错
	if m.Pid != 0 {
		tmp := &Menu{Id:m.Pid}
		err := o.Read(tmp, "id")
		if err != nil {
			return err
		}
	}

	if created, _, err := o.ReadOrCreate(m, "Name"); err == nil {
		if created {
			return nil
		}
	}
	return errors.New("菜单名已存在")
}

// 删除菜单
func (m *Menu)DeleteMenu() error {
	o := orm.NewOrm()
	// 如果有下级菜单无法删除
	temp := &Menu{Pid:m.Id}
	err := o.Read(temp, "pid")
	if err != nil  && err != orm.ErrNoRows{
		return err
	}
	if temp.Id != 0 {
		return errors.New("拥有子菜单，无法删除")
	}

	// 删除角色
	num, err := o.Delete(m)
	if err != nil {
		logs.Error("删除角色错误", err)
		return errors.New("系统错误")
	}
	if num == 0 {
		return errors.New("删除失败")
	}
	return nil
}

// 更新行为
func (m *Menu)UpdateMenu() error {
	o := orm.NewOrm()
	// 数据验证
	valid := validation.Validation{}
	b, err := valid.Valid(m)
	if err != nil {
		return err
	}

	if !b {
		for _, err := range valid.Errors {
			return err
		}
	}

	// 判断角色名是否已经存在
	tmp := &Menu{Name:m.Name}
	err = o.Read(tmp, "name")
	if err != nil && err != orm.ErrNoRows{
		return err
	}

	if tmp.Id != m.Id && tmp.Id != 0{
		return errors.New("菜单名已存在")
	}

	// 更新
	num, err := o.Update(m)
	if err != nil {
		return err
	}
	if num == 0{
		return errors.New("更新失败")
	}
	return nil
}

