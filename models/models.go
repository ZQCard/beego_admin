package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

// 数据库基础模型
type Model struct {
	Id        int `orm:"column(id);"`
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime);description(创建时间)"`
	UpdatedAt time.Time `orm:"column(updated_at);auto_now;type(datetime);description(更新时间)"`
	DeletedAt int	`orm:"column(deleted_at);description(删除时间,0为未删除)"`
}

// 初始化数据库句柄 切换数据库，或者，进行事务处理，都会作用于这个 Ormer 对象，以及其进行的任何查询。
// 所以：需要 切换数据库 和 事务处理 的话，不要使用全局保存的 Ormer 对象。
// var DB orm.Ormer

func init()  {
	// 注册数据库驱动
	err := orm.RegisterDriver("mysql", orm.DRMySQL)
	if err != nil {
		logs.Error("数据库连接错误", err)
		panic(err)
	}
	// 读取数据库配置文件
	host := beego.AppConfig.String("db_host")
	port := beego.AppConfig.String("db_port")
	username := beego.AppConfig.String("db_username")
	password := beego.AppConfig.String("db_password")
	database := beego.AppConfig.String("db_database")
	charset := beego.AppConfig.String("db_charset")
	// 拼接数据库参数
	dataSource := username + ":" + password + "@tcp(" +host + ":" + port + ")/" + database + "?charset=" + charset

	// 注册数据库
	err = orm.RegisterDataBase("default", "mysql", dataSource)
	if err != nil {
		logs.Error("注册数据库错误", err)
		panic(err)
	}
	orm.SetMaxIdleConns("default", 30)

	// 开启调试模式
	orm.Debug = true
}
