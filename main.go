package main

import (
	"beego_admin/controllers/admin"
	"beego_admin/handlers"
	_ "beego_admin/routers"
	"beego_admin/utils"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"time"
)

func init()  {
	// 支持表单伪造PUT,DELETE,PATCH,OPTIONS请求
	beego.InsertFilter("*", beego.BeforeRouter, handlers.RestfulHandler())
	//beego.InsertFilter("*", beego.BeforeExec, handlers.Auth())
}

func main() {
	// 输入日志文件设定
	logFileName := "./storage/logs/"+time.Now().Format("2006-01-02")+".log"
	logs.SetLevel(logs.LevelError)
	logs.SetLogger(logs.AdapterFile, `{"filename":"`+logFileName+`"}`)
	// 输入行号
	logs.EnableFuncCallDepth(true)
	// 异步输出
	logs.Async()
	// 自定义错误处理
	beego.ErrorController(&admin.ErrorController{})
	// 自定义模板函数
	beego.AddFuncMap("TimestampToDate", utils.TimestampToDate)
	beego.Run()
}

