package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

/**
CREATE TABLE `resources` (
  `id` varchar(32) NOT NULL,
  `file_key` varchar(255) DEFAULT NULL,
  `type` varchar(64) DEFAULT NULL,
  `mime_type` varchar(64) DEFAULT NULL,
  `reference_id` varchar(32) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `original_file_name` varchar(255) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `extension` varchar(64) DEFAULT NULL,
  `storage_type` varchar(64) DEFAULT NULL,
  `storage_param` varchar(64) DEFAULT NULL,
  `size` int(10) unsigned DEFAULT NULL,
  `meta` longtext,
  `created_time` datetime(6) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
*/

type Resources struct {
	Id               string    `orm:"column(id);pk;size(32)" `
	FileKey          string    `orm:"column(file_key);size(255)"`
	Type             string    `orm:"column(type);size(54)"`
	MimeType         string    `orm:"column(mime_type);size(64)"`
	ReferenceId      string    `orm:"column(reference_id);size(32)"`
	Name             string    `orm:"column(name);size(255)"`
	OriginalFileName string    `orm:"column(original_file_name);size(255)"`
	Description      string    `orm:"column(description);size(255)"`
	Extension        string    `orm:"column(extension);size(64)"`
	StorageType      string    `orm:"column(storage_type);size(54)"`
	StorageParam     string    `orm:"column(storage_param);size(64)"`
	Size             int       `orm:"column(size)"`
	Meta             string    `orm:"column(meta);size(32)"`
	CreatedTime      time.Time `orm:"column(created_time);auto_now_add;type(datetime)"`
}

// 需要在init中注册定义的model
func init() {
	orm.RegisterModel(new(Resources))
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
	CreateTable()
}

func CreateTable() {
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
