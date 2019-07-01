package admin

import (
	"beego_admin/utils"
	"fmt"
	"github.com/astaxie/beego"
	"log"
	"os"
	"time"
)

type ToolsController struct {
	baseController
}

func (c *ToolsController)GetUploadFile()  {
	c.Data["Title"] = "文件上传功能"
	// 模板
	c.TplName = "admin/tools/uploadFile.html"
}

func (c *ToolsController)PostUploadFile()  {

	// 接收文件
	file, handler, err := c.GetFile("file")
	if err != nil {
		log.Fatal("接收文件失败：", err)
	}
	// 关闭句柄
	defer file.Close()
	// 根据日期保存文件目录
	today := time.Now().Format("2006-01-02")
	directory := "static/uploadFile/" + today + "/"
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

	returnJson:= ResponseJson{}
	if err != nil {
		returnJson.StatusCode = Fail
		returnJson.Message = "文件上传失败"
	}

	returnJson.StatusCode = Success
	fmt.Println(beego.AppConfig.String("base_url"))
	returnJson.Url = beego.AppConfig.String("base_url") + filePath
	c.Data["json"] = &returnJson
	c.ServeJSON()
}