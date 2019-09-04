package admin

import (
	"beego_admin/models"
	"errors"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/utils"
	"github.com/go-ozzo/ozzo-validation"
)

type Menu struct {
	ID int
	Pid int
	Name string
	Sort int
	Route string
}

type TreeList struct {
	Id int			`json:"id"`
	Name string		`json:"name"`
	Pid int			`json:"pid"`
	Sort int 		`json:"sort"`
	Route string	`json:"route"`
	Children []*TreeList	`json:"children"`
}

func (Menu)TableName() string {
	return "auth_menu"
}

func (menu Menu)Validate() error {

	return validation.ValidateStruct(&menu,
		// 名称不得为空,且大小为1-20字
		validation.Field(
			&menu.Name,
			validation.Required.Error("菜单名称不得为空"),
			validation.Length(1, 60).Error("名称为1-20字")),

		// 路由
		validation.Field(
			&menu.Sort,
			validation.Required.Error("排序不得为空"),
		),
	)
}

var menuRoute []string

var authRoute []string

/**
菜单列表
 */
func (menu *Menu)MenuList(route []string) []*TreeList{
	authRoute = route
	return menu.getMenu(menu.Pid)
}

/**
	递归获取树形菜单
 */
func (m *Menu)getMenu(pid int) []*TreeList {
	// 查找除所有pid的子菜单
	var menu []Menu
	models.DB.Where("pid = ?", pid).Order("sort").Find(&menu)
	treeList := []*TreeList{}
	for _, v := range menu{
		if len(authRoute) != 0{
			// 根据当前菜单Pid判断是否与有路由在权限路由有交集，如果没有continue
			// 如果route为空,有下级  不为空,无下级
			if v.Route == "" {
				menuRoute = findActionByMenuId(v.ID)
			} else {
				menuRoute = []string{v.Route}
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
		child := v.getMenu(v.ID)
		node := &TreeList{
			Id:v.ID,
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
	var menu []Menu
	models.DB.Where("pid = ?", id).Order("sort").Find(&menu)
	for _, v := range menu {
		findActionByMenuId(v.ID)
		menuRoute = append(menuRoute, v.Route)
	}
	return menuRoute
}

// 创建菜单
func (menu *Menu)MenuCreate()  (err error) {
	// 数据验证
	err = menu.Validate()
	if err != nil {
		return err
	}
	// 路由是否正确
	if menu.Route != "" {
		menu := Action{Route:menu.Route}
		_, err = menu.FindAction()
		if err != nil {
			return errors.New("路由填写不正确")
		}
	}

	// 判断用户名或者昵称未使用
	if !models.DB.Where("name = ?", menu.Name).Find(&Menu{}).RecordNotFound(){
		return errors.New("菜单名称已经存在")
	}

	if err = models.DB.Create(&menu).Error; err != nil {
		logs.Error("菜单创建失败", err)
		return errors.New("信息保存失败")
	}
	return nil
}

// 更新菜单
func (menu *Menu)MenuUpdate() (err error) {
	// 数据验证
	err = menu.Validate()
	if err != nil {
		return err
	}
	// 路由是否正确
	if menu.Route != "" {
		menu := Action{Route:menu.Route}
		_, err = menu.FindAction()
		if err != nil {
			return errors.New("路由填写不正确")
		}
	}

	// 判断用户名或者昵称未使用
	if !models.DB.Where("id <> ?", menu.ID).Where("name = ?", menu.Name).Find(&Action{}).RecordNotFound(){
		return errors.New("菜单名称已经存在")
	}

	if err = models.DB.Save(&menu).Error; err != nil {
		logs.Error("菜单保存失败", err)
		return errors.New("信息保存失败")
	}
	return nil
}

// 删除菜单
func (menu *Menu)MenuDelete() (err error) {
	if err := models.DB.Delete(&menu).Error; err != nil {
		logs.Error("删除菜单失败：", err)
		return errors.New("删除失败")
	}
	return nil
}

