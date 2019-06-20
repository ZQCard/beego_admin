package admin

import (
	"fmt"
	"github.com/astaxie/beego"
	"net"
	"runtime"
	"strings"
)

type IndexController struct {
	baseController
}

// 后台首页
func (c *IndexController)Get(){
	// 模板
	c.TplName = "admin/index/index.html"
	// 获取系统信息

	// 项目名称
	c.Data["AppName"] = beego.AppConfig.String("appname")
	// 框架版本
	c.Data["FrameWorkVersion"] = "Beego : " + beego.VERSION
	// Go版本
	c.Data["Version"] = runtime.Version()
	// 操作系统
	c.Data["GOOS"] = runtime.GOOS
	// 系统架构
	c.Data["GOARCH"] = runtime.GOARCH
	// CPU核数
	c.Data["GOMAXPROCS"] = runtime.GOMAXPROCS(0)

	// 当前goroutine数目
	c.Data["Goroutine"] = runtime.NumGoroutine()

	// 获取ip地址
	conn, err := net.Dial("udp", "google.com:80")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer conn.Close()
	c.Data["IP"] = strings.Split(conn.LocalAddr().String(), ":")[0]
	c.Data["Title"] = "后台首页"
	c.TplName = "admin/index/index.html"
}

// 退出登录
func (c *IndexController)Logout()  {
	c.DestroySession()
	c.Redirect("/login", 302)
}