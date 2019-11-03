package common

import (
	"beego_admin/controllers/admin"
	"beego_admin/utils"
	"fmt"
	"github.com/astaxie/beego"
	"log"
	"os"
	"time"
)
type UploadFileController struct {
	beego.Controller
}

// 公共请求不需要CSRF
func (c *UploadFileController)Prepare()  {
	c.EnableXSRF = false
}

func (c *UploadFileController)PostUploadFile()  {

	// 接收文件
	file, handler, err := c.GetFile("file")
	if err != nil {
		log.Fatal("接收文件失败：", err)
	}
	// 关闭句柄
	defer file.Close()
	// 根据日期保存文件目录
	today := time.Now().Format("2006-01-02")
	// 读取文件查看是否有目录设置
	dir := c.Input().Get("type")
	if dir == ""{
		dir = "common"
	}
	// 当前项目路径
	appDir, _ := os.Getwd()
	fmt.Println(appDir)
	directory := "static/uploadFile/" +dir+"/"+ today + "/"
	// 建立文件夹
	err = os.MkdirAll(directory, os.ModePerm)
	if err != nil {
		log.Fatal("建立文件夹失败：", err)
	}

	filePath := directory + utils.RandString(10)+"_"+handler.Filename
	// 混淆文件名前缀，并保存文件到指定地点
	err = c.SaveToFile("file", filePath)
	if err != nil {
		log.Fatal("接收文件失败：", err)
	}

	returnJson:= admin.ResponseJson{}
	if err != nil {
		returnJson.StatusCode = admin.Fail
		returnJson.Message = "文件上传失败"
	}

	returnJson.StatusCode = admin.Success
	returnJson.Url = beego.AppConfig.String("base_url") +"/"+ filePath
	c.Data["json"] = &returnJson
	c.ServeJSON()
}
