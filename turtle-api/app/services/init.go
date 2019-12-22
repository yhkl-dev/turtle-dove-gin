package services

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/yhkl-dev/turtle-dove-beego/turtle-api/app/tables"
)

var (
	o           orm.Ormer
	UserService *userService
	AuthService *authService
)

func init() {
	dbHost := beego.AppConfig.String("db.host")
	dbPort := beego.AppConfig.String("db.port")
	dbUser := beego.AppConfig.String("db.user")
	dbPassword := beego.AppConfig.String("db.password")
	dbName := beego.AppConfig.String("db.name")

	if dbPort == "" {
		dbPort = "3306"
	}

	dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8"

	orm.RegisterDataBase("default", "mysql", dsn)
	orm.RegisterModel(
		new(tables.User),
		new(tables.Role),
		new(tables.AuthToken),
	)
	orm.RunSyncdb("default", false, true)

	o = orm.NewOrm()
	orm.RunCommand()

	initService()
}

func initService() {
	UserService = &userService{}
	AuthService = &authService{}
}

func tableName(name string) string {
	return name
}
