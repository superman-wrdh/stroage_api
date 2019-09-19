package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// 注册模型 初始化数据库表结构
func init() {
	//注册模型
	registerModel()
	//初始化数据库连接
	initDb()
	//创建表结构
	createTable()
}

func initDb() {
	err := orm.RegisterDriver("mysql", orm.DRMySQL)
	if err != nil {
		fmt.Print("RegisterDriver Error")
	}
	host := beego.AppConfig.String("dev_db::host")
	user := beego.AppConfig.String("dev_db::user")
	password := beego.AppConfig.String("dev_db::password")
	db := beego.AppConfig.String("dev_db::db")
	port := beego.AppConfig.String("dev_db::port")
	DbInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", user, password, host, port, db)
	RegisterErr := orm.RegisterDataBase("default", "mysql", DbInfo)

	if RegisterErr != nil {
		fmt.Print("RegisterDataBase Error")
	}
}

func createTable() {
	name := "default"
	// if true drop table 后再建表
	force := false
	// 打印执行过程
	verbose := true
	// 遇到错误立即返回
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}
}

func registerModel() {
	orm.RegisterModel(new(Resources))
}
