package main

import (
	"bytes"
	"fmt"
	"github.com/astaxie/beego/config"
	"github.com/robfig/cron"
	"log"
	"os/exec"
	"time"
)

func main()  {
	startCron()
}

// 开始定时任务
func startCron()  {
	log.Println("starting....")
	c := cron.New()

	c.AddFunc("0 0 1 * * ?", func(){
		mysqlDump()
	})

	c.Start()

	t1 := time.NewTimer(time.Second * 1)
	for{
		select{
		case <- t1.C:
			t1.Reset(time.Second * 1)
		}
	}
}

// 备份docker中的item_fmg数据库 备份到/usr/local/applications/data/mysqlbak/ 目录下
func mysqlDump()  {
	fileName := "item_fmg-"+time.Now().Format("2006-01-02-15:04:05")+".sql"
	filePath := "/usr/local/applications/data/mysqlbak/"+fileName
	// 读取配置文件内容
	conf,err := config.NewConfig("ini", "./../conf/my.ini")
	if err != nil {
		fmt.Println("new config failed, err:", err)
		return
	}

	// 读取配置文件数据库账号和密码
	dbName := conf.String("db_username")

	dbPassword := conf.String("db_password")

	commandDump := "sudo docker exec -i mysql_master mysqldump --opt -u"+dbName+" -p"+dbPassword+" item_fmg > " + filePath
	cmd := exec.Command("/bin/sh", "-c", commandDump)
	// 获取报错输出内容
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err != nil {
		fmt.Println(err)
		fmt.Println(stderr.String())
	}else {
		fmt.Println(out.String())
	}
	fmt.Println("success", filePath)
	return
}
