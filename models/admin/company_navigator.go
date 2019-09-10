package admin

import (
	"beego_admin/models"
	"errors"
	"github.com/astaxie/beego/logs"
	"github.com/go-ozzo/ozzo-validation"
)

type Navigator struct {
	models.Model
	Pid int
	Name string
	Url string
	Sort int
}

type NavigatorTreeList struct {
	ID int			`json:"id"`
	Name string		`json:"name"`
	Pid int			`json:"pid"`
	Sort int 		`json:"sort"`
	Url string		`json:"url"`
	Children []*NavigatorTreeList	`json:"children"`
}

func (Navigator)TableName() string {
	return "company_navigator"
}

func (navigator Navigator)Validate() error {

	return validation.ValidateStruct(&navigator,
		// 名称不得为空,且大小为1-20字
		validation.Field(
			&navigator.Name,
			validation.Required.Error("导航栏名称不得为空"),
			validation.Length(1, 60).Error("名称为1-20字")),

		// 排序
		validation.Field(
			&navigator.Sort,
			validation.Required.Error("排序不得为空"),
		),
	)
}

/**
导航栏列表
 */
func (navigator *Navigator)List(route []string) []*NavigatorTreeList{
	return navigator.getNavigator(navigator.Pid)
}

func (navigator *Navigator)getNavigator(pid int) []*NavigatorTreeList {
	// 查找除所有pid的子导航栏
	var navigators []Navigator
	models.DB.Where("pid = ?", pid).Order("sort").Find(&navigators)
	treeList := []*NavigatorTreeList{}
	for _, v := range navigators{
		child := v.getNavigator(v.ID)
		node := &NavigatorTreeList{
			ID:v.ID,
			Name:v.Name,
			Sort:v.Sort,
			Url:v.Url,
			Pid:v.Pid,
		}
		node.Children = child
		treeList = append(treeList, node)
	}
	return treeList
}

// 创建导航栏
func (navigator *Navigator)Create()  (err error) {
	// 数据验证
	err = navigator.Validate()
	if err != nil {
		return err
	}

	// 判断用户名或者昵称未使用
	if !models.DB.Where("name = ?", navigator.Name).Find(&Navigator{}).RecordNotFound(){
		return errors.New("导航栏名称已经存在")
	}

	if err = models.DB.Create(&navigator).Error; err != nil {
		logs.Error("导航栏创建失败", err)
		return errors.New("信息保存失败")
	}
	return nil
}

// 更新导航栏
func (navigator *Navigator)Update() (err error) {
	// 数据验证
	err = navigator.Validate()
	if err != nil {
		return err
	}

	// 判断名称未使用
	if !models.DB.Where("id <> ?", navigator.ID).Where("name = ?", navigator.Name).Find(&Navigator{}).RecordNotFound(){
		return errors.New("导航栏名称已经存在")
	}

	if err = models.DB.Select("pid", "name", "url", "sort").Save(&navigator).Error; err != nil {
		logs.Error("导航栏保存失败", err)
		return errors.New("信息保存失败")
	}
	return nil
}

// 删除导航栏
func (navigator *Navigator)Delete() (err error) {
	if err := models.DB.Delete(&navigator).Error; err != nil {
		logs.Error("删除导航栏失败：", err)
		return errors.New("删除失败")
	}
	return nil
}