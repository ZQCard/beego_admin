package admin

import (
	"beego_admin/models/demo"
	"beego_admin/utils"
	"fmt"
	"github.com/astaxie/beego"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type ToolsController struct {
	baseController
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
	// 读取文件查看是否有目录设置
	dir := c.Input().Get("type")
	if dir == ""{
		 dir = "common"
	}

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
/***************** 文件上传模拟功能结束 ******************/

/***************** 邮件模拟功能开始 ******************/
func (c *ToolsController)GetSendEmail()  {
	c.Data["Title"] = "邮件发送功能"
	// 模板
	c.TplName = "admin/tools/sendEmail.html"
}

func (c *ToolsController)PostSendEmail()  {
	// 获取邮箱列表
	emails := c.Input().Get("emailList")
	content := c.Input().Get("content")
	// 将邮件字符串进行分割,并放入slice中
	emailSlice := strings.Split(emails, "\r\n")
	successNum, failNum, unreachable := utils.SendEmail("邮件发送测试", emailSlice, content)
	returnJson:= ResponseJson{}
	returnJson.StatusCode = Success
	returnJson.Message = "请求成功,发送成功" +
		strconv.Itoa(successNum) + " 封;发送失败 " +
		strconv.Itoa(failNum)+ " 封;无效" +
		strconv.Itoa(unreachable) + "封"
	c.Data["json"] = &returnJson
	c.ServeJSON()
}
/***************** 邮件模拟功能结束 ******************/


/***************** excel导入导出功能结束 ******************/
//func (c *ToolsController)GetExcel()  {
//	c.Data["Title"] = "excel导入导出功能"
//	c.Data["ExampleExcelUrl"] = beego.AppConfig.String("base_url") + "static/example/excelDemo.xlsx"
//	// 模板
//	c.TplName = "admin/tools/excel.html"
//}
//
//
//func (c *ToolsController)PostExcelImport()  {
//	// 接收文件
//	file, handler, err := c.GetFile("file")
//	if err != nil {
//		log.Fatal("接收文件失败：", err)
//	}
//	// 关闭句柄
//	defer file.Close()
//	// 存放临时文件
//	directory := "storage/temp/"
//	_, err = os.Stat(directory)
//	if err != nil {
//		// 建立文件夹
//		err = os.MkdirAll(directory, os.ModePerm)
//		if err != nil {
//			log.Fatal("建立文件夹失败：", err)
//		}
//	}
//
//	filePath := directory + utils.RandString(10)+"_"+handler.Filename
//	// 混淆文件名前缀，并保存文件到指定地点
//	err = c.SaveToFile("file", filePath)
//	if err != nil {
//		log.Fatal("保存文件失败：", err)
//	}
//	// 删除文件
//	defer os.Remove(filePath)
//
//	content := utils.GetExcelContent(filePath)
//	// 将excel内容除第一行 插入数据库
//	insertContent := []demo.Excel{}
//
//	for index,row := range content{
//		if index == 0 {
//			continue
//		}
//		// 将每一行设置成数据结构体
//		demo := demo.Excel{}
//		// 空值标志符
//		empty := false
//		for key,value := range row{
//			// 去除空值
//			if value == ""{
//				 empty = true
//				 continue
//			}
//			// 设置姓名
//			if key == 0 {
//				demo.Name = value
//			}
//			// 设置年龄
//			if key == 1 {
//				demo.Age,_ = strconv.Atoi(value)
//			}
//			// 将年月日改变成时间戳
//			if key == 2 {
//				// excel日期需要单独处理
//				theTime, _ := time.Parse("2006-01-02", utils.ConvertToFormatDay(value))
//				demo.Birthday = theTime.Unix()
//			}
//		}
//		if empty == false{
//			insertContent = append(insertContent, demo)
//		}
//	}
//	successNums, err := demo.Import(insertContent)
//	returnJson:= ResponseJson{}
//	if err != nil {
//		returnJson.StatusCode = Fail
//		returnJson.Message = "导入失败"
//	}else {
//		returnJson.StatusCode = Success
//		returnJson.Message = "导入成功" + strconv.Itoa(int(successNums)) + "条"
//	}
//	returnJson.UrlType = Reload
//	c.Data["json"] = &returnJson
//	c.ServeJSON()
//}
//
//func (c *ToolsController)GetExcelExport()  {
//
//	title := &[]interface{}{"姓名", "年龄", "出生日期"}
//	// 读取数据库测试数据
//	excelData := demo.GetAllData()
//	var content [][]interface{}
//	// 将数据结构体放入 interface slice
//	for _, v := range excelData{
//		content = append(content, []interface{}{v.Name, v.Age, time.Unix(v.Birthday, 0).Format("2006-01-02")})
//	}
//	path, err := utils.SetExcelContent("excelDownload", title, &content)
//	returnJson:= ResponseJson{}
//	if err != nil {
//		returnJson.StatusCode = Fail
//		returnJson.Message = "文件生成失败"
//	}else {
//		returnJson.StatusCode = Success
//		returnJson.UrlType = Jump
//		returnJson.Url = beego.AppConfig.String("base_url") + path
//	}
//	returnJson.UrlType = Reload
//	c.Data["json"] = &returnJson
//	c.ServeJSON()
//}
/***************** excel导入导出功能结束 ******************/