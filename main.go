package main

import (
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/zhangCan112/stm_server_project/routers"
)

func init() {
	orm.DefaultTimeLoc = time.UTC
	orm.RegisterDriver("mysql", orm.DRMySQL)
	dbname := beego.AppConfig.String("mysqldb")
	url := beego.AppConfig.String("mysqlurls")
	port := beego.AppConfig.String("mysqlport")
	user := beego.AppConfig.String("mysqluser")
	pw := beego.AppConfig.String("mysqlpass")
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", user, pw, url, port, dbname)
	orm.RegisterDataBase("default", "mysql", dataSource)
}

func main() {

	syncdb()

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}

func syncdb() {
	// 数据库别名
	name := "default"

	// drop table 后再建表
	force := true

	// 打印执行过程
	verbose := true

	// 遇到错误立即返回
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}
}
