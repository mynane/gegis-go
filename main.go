package main

import (
	_ "gegis/routers"

	"github.com/astaxie/beego"

	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
		)

func init()  {
	// start mysql
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:12qwaszx@tcp(localhost:3306)/gegis?charset=utf8")
	orm.SetMaxIdleConns("default", 1000)
	orm.SetMaxOpenConns("default", 2000)

	// start cron
	//utils.Cron()
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
