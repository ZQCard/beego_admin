package common

import (
	"beego_admin/handlers"
	"beego_admin/models/common"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"strconv"
)

type WeChatController struct {
	beego.Controller
}

func (c *WeChatController)GetUserLogin()  {
	redirect := c.Input().Get("redirect")
	code := c.Input().Get("code")

	// 微信配置
	appId := beego.AppConfig.String("wechat_app_id")
	appSecret := beego.AppConfig.String("wechat_app_secret")
	accessToken,_,openId, err := handlers.GrantGetAccessToken(appId, appSecret, code)
	if err != nil {
		logs.Error(err.Error())
		c.Redirect("/"+redirect, 302)
		return
	}
	userInfo,err := handlers.GrantGetUserInfo(accessToken, openId)
	if err != nil {
		logs.Error(err.Error())
		c.Ctx.WriteString("<h1>系统异常</h1>")
		return
	}
	user := common.User{}
	user.Nickname = userInfo.Nickname
	user.WechatOpenId = userInfo.OpenId
	user.Sex = userInfo.Sex
	user.Province = userInfo.Province
	user.City = userInfo.City
	user.HeadImgUrl = userInfo.HeadImgUrl
	user.LoginIp = c.Ctx.Request.Header.Get("X-Real-ip")
	user, err = user.FindUserByOpenId()
	if err != nil {
		logs.Error(err.Error())
		c.Ctx.WriteString("用户信息保存失败")
		return
	}

	if redirect != ""{
		c.Redirect("/"+redirect, 302)
		return
	}

	// 根据openId保存用户信息 如果不存在,则添加
	c.Ctx.WriteString("user_id : " +  strconv.Itoa(user.ID))
	c.Ctx.WriteString("nickname : " +  user.Nickname)
	return
}