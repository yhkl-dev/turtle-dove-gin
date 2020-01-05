package main

import (
	"github.com/yhkl-dev/turtle-dove-beego/gin-turtle-api/conf"
	_ "github.com/yhkl-dev/turtle-dove-beego/gin-turtle-api/database"
	_ "github.com/yhkl-dev/turtle-dove-beego/gin-turtle-api/docs"
	"github.com/yhkl-dev/turtle-dove-beego/gin-turtle-api/router"
)

// @title Turtel Dove API
// @version 1.0
// @description Turtle Dove API for rds manager
// @termsOfService http://yangkai.org.cn

// @contact.name API Support
// @contact.url https://github.com/yhkl-dev
// @contact.email kaiyang939325@gmail.com

func main() {
	port, err := conf.SysConfig.GetValue("Base", "port")
	if err != nil {
		panic("Loading config file error")
	}

	r := router.InitRouter()
	r.Run(":" + port)
}
