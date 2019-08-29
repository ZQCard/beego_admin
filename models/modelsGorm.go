package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

var DB *gorm.DB

// 数据库基础模型
type ModelGORM struct {
	ID        int `gorm:"primary_key"json:"Id"`
	CreatedAt time.Time `json:"CreatedAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
	DeletedAt *time.Time `json:"DeletedAt"`
}


func init()  {
	var (
		err error
		dbType, dbName, dbUser, dbPassword, dbHost string
	)

	dbType = beego.AppConfig.String("db_type")
	dbName = beego.AppConfig.String("db_database")
	dbUser = beego.AppConfig.String("db_username")
	dbPassword = beego.AppConfig.String("db_password")
	dbHost = beego.AppConfig.String("db_host")

	DB, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dbUser,
		dbPassword,
		dbHost,
		dbName,
	))
	if err != nil {
		logs.Error("注册数据库错误", err)
	}

	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)

	// 禁止使用复数表明
	DB.SingularTable(true)
	// 查看原生sql
	DB.LogMode(true)
}
