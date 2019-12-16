package main

import (
	_ "github.com/yhkl-dev/turtle-dove-beego/turtle-api/app/services"
	_ "github.com/yhkl-dev/turtle-dove-beego/turtle-api/routers"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
