package company

import (
	"beego_admin/handlers"
	"beego_admin/models"
	"beego_admin/models/common"
	"beego_admin/utils"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/logs"
	"time"
)
var bm cache.Cache
func init()  {
	// 缓存对象
	bm,_ = cache.NewCache("file", `{"CachePath":"./storage/cache","FileSuffix":".cache","DirectoryLevel":2,"EmbedExpiry":3600}`)
}
// 基础控制器
type IndexController struct {
	beego.Controller
}

func (c *IndexController)Index()  {
	flash := beego.NewFlash()

	id := utils.MustInt(c.Input().Get("id"))
	video := common.Video{
		Model:models.Model{
			ID:id,
		},
	}

	video,err := video.Find()
	if err != nil || id == 0{
		logs.Error("视频首页错误：", err)
		flash.Error("视频不存在")
		flash.Store(&c.Controller)
		c.Redirect("/error", 302)
		return
	}
	video.AddViewTimes()

	// 微信配置
	appId := beego.AppConfig.String("wechat_app_id")
	appSecret := beego.AppConfig.String("wechat_app_secret")
	c.Data["WeChatAppId"] = appId
	c.Data["WeChatAppSecret"] = appSecret

	// 微信js票据
	ticket := getTicket(appId, appSecret)
	if ticket == ""{
		logs.Error("获取JsApiTicket错误")
		flash.Error("获取信息错误")
		flash.Store(&c.Controller)
		c.Redirect("/error", 302)
		return
	}
	// 随机按字符串
	noncestr := utils.RandString(10)
	// 时间戳
	timestamp := time.Now().Unix()

	// 分享当前页面
	url := beego.AppConfig.String("base_url") + c.Ctx.Request.RequestURI
	signature := handlers.GetSignature(ticket, noncestr, url, timestamp)
	c.Data["WeChatTimestamp"] = timestamp
	c.Data["WeChatNoncestr"] = noncestr
	c.Data["WeChatSignature"] = signature
	c.Data["WeChatShareUrl"] = url
	c.Data["Title"] = "太格摄影团队"
	c.Data["Video"] = video
	c.TplName = "company/index.html"
}

// 获取微信授权access_token
func getAccessToken(appId, appSecret string) string {
	var accessToken = ""
	// 缓存access_token,access_token不存在则重新读取
	if (!bm.IsExist("wechat_access_token")) || (bm.Get("wechat_access_token").(string) == ""){
		// 获取accessToken
		token, expire, err := handlers.JsApiGetAccessToken(appId, appSecret)
		if err != nil {
			logs.Error("微信access_token获取错误：", err)
			return ""
		}
		err = bm.Put("wechat_access_token", token, expire * time.Second)
		if err != nil {
			logs.Error("wechat_access_token数据缓存失败", err)
			return ""
		}
		accessToken = token
	}else {
		// 直接读取缓存内容
		accessToken = bm.Get("wechat_access_token").(string)
	}
	return accessToken
}

func getTicket(appId, appSecret string) string {

	accessToken := getAccessToken(appId, appSecret)
	// 缓存ticket
	var jsApiTicket = ""
	if (!bm.IsExist("wechat_js_api_tiket")) || (bm.Get("wechat_js_api_tiket").(string) == "") {
		// 获取accessToken
		ticket, expire, err := handlers.GetJsApiTicket(accessToken)
		if err != nil {
			logs.Error("微信wechat_js_api_tiket获取错误：", err)
			return ""
		}
		err = bm.Put("wechat_js_api_tiket", ticket, expire * time.Second)
		if err != nil {
			logs.Error("wechat_js_api_tiket数据缓存失败", err)
			return ""
		}
		jsApiTicket = ticket
	}else {
		// 直接读取缓存内容
		jsApiTicket = bm.Get("wechat_js_api_tiket").(string)
	}
	return jsApiTicket
}
