package main

import (
	"fmt"
	"github.com/robfig/cron"
	"log"
	"os/exec"
	"time"
)

func main()  {
	mysqlDump()
	// startCron()
}

// 开始定时任务
func startCron()  {
	log.Println("starting....")
	c := cron.New()

	c.AddFunc("* * * * * *", func(){
		fmt.Print(111)
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
	fileName := "item_fmg-"+time.Now().Format("2006-01-02")+".sql"
	filePath := "/usr/local/applications/data/mysqlbak/"+fileName



	// 执行mysql导入
	commandDump := "docker exec -it mysql_master mysqldump --opt -t -uroot -pzhouqi445864742. item_fmg > " + filePath
	fmt.Println(commandDump)
	out, _ := exec.Command("/bin/sh", "-c", commandDump).Output()
	fmt.Printf("%s", out)
	//commandDump = commandDump
	// docker exec -it mysql_master mysqldump --opt -t -uroot -pzhouqi445864742. item_fmg
	//cmd := exec.Command("docker", "exec" , "-it", "mysql_master", "mysqldump", "--opt", "-t", "-uroot", "-pzhouqi445864742.", "item_fmg")
	//stdout, err := cmd.StdoutPipe()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//if err := cmd.Start(); err != nil {
	//	log.Fatal(err)
	//}
	//
	//bytes, err := ioutil.ReadAll(stdout)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(bytes)
	//err = ioutil.WriteFile(filePath, bytes, 0644)
	//
	//if err != nil {
	//	panic(err)
	//}
}
