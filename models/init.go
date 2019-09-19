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
	/**
	根据运行模式选择对应数据库
	当运行模式为dev 时候 选项 section为dev下面的配置
	*/
	err := orm.RegisterDriver("mysql", orm.DRMySQL)
	if err != nil {
		fmt.Print("RegisterDriver Error")
	}

	runModel := beego.AppConfig.String("runmode")

	host := beego.AppConfig.String(fmt.Sprintf("%s_db::host", runModel))
	user := beego.AppConfig.String(fmt.Sprintf("%s_db::user", runModel))
	password := beego.AppConfig.String(fmt.Sprintf("%s_db::password", runModel))
	db := beego.AppConfig.String(fmt.Sprintf("%s_db::db", runModel))
	port := beego.AppConfig.String(fmt.Sprintf("%s_db::port", runModel))
	DbInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", user, password, host, port, db)
	RegisterErr := orm.RegisterDataBase("default", "mysql", DbInfo)

	if RegisterErr != nil {
		fmt.Print("RegisterDataBase Error")
	}
}

func createTable() {
	name := "default"
	// if true drop table 后再建表
	force, boolErr := beego.AppConfig.Bool("ForceCreateTable")
	if boolErr != nil {
		force = false
	}
	// 打印执行过程
	verbose := true
	// 遇到错误立即返回
	runModel := beego.AppConfig.String("runmode")

	//生产环境不能删除表重建
	if runModel == "prd" || runModel == "product" {
		force = false
	}
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}

}

func registerModel() {
	orm.RegisterModel(new(Resources))
}
