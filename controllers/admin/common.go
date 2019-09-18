package admin

import (
	"beego_admin/models/admin"
	"beego_admin/utils"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	captcha2 "github.com/astaxie/beego/utils/captcha"
	"html/template"
)

// ajax返回信息
type ResponseJson struct {
	StatusCode int `json:"statusCode"` // 返回状态值
	Data interface{} `json:"data"` // 返回数据
	Message string `json:"message"` // 提示信息
	Url string `json:"url"`	// 跳转url
	UrlType int `json:"urlType"`	// url操作类型
}

// Message 提示中文信息
const SaveSuccess  = "保存成功"
const SaveFail  = "保存失败"
const AddSuccess  = "添加成功"
const AddFail = "添加失败"
const DeleteSuccess = "删除成功"
const DeleteFail = "删除失败"

// StatusCode 请求返回值
const Fail = 0  	// 表示操作失败
const Success = 1	// 表示操作成功
const Unknow = 2	// 表示未知错误

// urlType
const Nothing = 0  // 不做任何操作
const Jump = 1		// 跳转
const Reload = 2   // 刷新页面

// 分页数据信息
const PageDefault = 1
const PageSizeDefault = 10

// 公共controller
type CommonController struct {
	beego.Controller
}

var captcha *captcha2.Captcha

func init()  {
	// 验证码功能
	// 使用Beego缓存存储验证码数据
	store := cache.NewMemoryCache()
	captcha = captcha2.NewWithFilter("/captcha", store)
	captcha.ChallengeNums = 4
	captcha.StdHeight = 50
	captcha.StdWidth = 120
}

// 登录页面
func (c *CommonController)LoginPage()  {
	// 登录判断
	adminId := c.GetSession("adminId")
	if adminId != nil {
		c.Redirect("/", 302)
	}
	c.Data["Error"] = ""
	// csrf
	c.Data["CsrfData"] = template.HTML(c.XSRFFormHTML())

	c.TplName = "admin/common/login.html"
}

// 登录
func (c *CommonController)Login()  {
	// 验证码验证
	if !captcha.VerifyReq(c.Ctx.Request) {
		c.Data["Error"] = "验证码错误"
		// csrf
		c.Data["CsrfData"] = template.HTML(c.XSRFFormHTML())
		c.TplName = "admin/common/login.html"
		return
	}
	// 读取post参数
	username := c.Input().Get("username")
	password := utils.GenerateMD5String(c.Input().Get("password"))
	Administrator := admin.Administrator{
		Username:username,
		Password:password,
	}

	Administrator, err := Administrator.FindAdministrator()
	// 用户验证失败
	if err != nil{
		c.Data["Error"] = err.Error()
		// csrf
		c.Data["CsrfData"] = template.HTML(c.XSRFFormHTML())
		c.TplName = "admin/common/login.html"
		return
	}
	// 保存session
	// 保存用户信息
	c.SetSession("adminId", Administrator.Model.ID)
	c.SetSession("adminName", Administrator.Username)
	c.SetSession("nickname", Administrator.Nickname)

	// 读取权限map
	authList, err := Administrator.AuthList()
	if err != nil {
		c.Data["Error"] = err.Error()
		c.Data["CsrfData"] = template.HTML(c.XSRFFormHTML())
		c.TplName = "admin/common/login.html"
		return
	}
	// session无法读取存储map[string][]string,所以分为请求方式存储
	for k, v := range authList{
		c.SetSession("AUTH_" + k, v)
	}

	// 读取公共权限
	role := admin.Role{Name:beego.AppConfig.String("customer_role_name")}
	// 读取公共权限
	authCommonList, err := role.AuthList()
	if err != nil {
		c.Data["Error"] = err.Error()
		c.Data["CsrfData"] = template.HTML(c.XSRFFormHTML())
		c.TplName = "admin/common/login.html"
		return
	}
	// session无法读取存储map[string][]string,所以分为请求方式存储
	for k, v := range authCommonList{
		c.SetSession("AUTH_COMMON" + k, v)
	}

	// 根据GET权限读取后台菜单列表
	authGet := c.GetSession("AUTH_GET")
	var authGetSlice []string
	switch t := authGet.(type) {
	case []string:
		for _, v := range t {
			authGetSlice = append(authGetSlice, v)
		}
	default:
		c.Data["Error"] = "解析权限失败"
		c.Data["CsrfData"] = template.HTML(c.XSRFFormHTML())
		c.TplName = "admin/common/login.html"
		return
	}
	// 读取管理员的菜单
	c.SetSession("MENU_LEFT", Administrator.MenuList(authGetSlice))
	c.Redirect("/admin/", 302)
	return
}

// 错误提示页面
func (c *CommonController)Error()  {
	errMsg := c.GetSession("error")
	c.SetSession("error", nil)
	if errMsg == nil {
		errMsg = "好像出错了呢^ - ^"
	}
	c.Data["Error"] = errMsg
	c.TplName = "admin/common/error.html"
	return
}