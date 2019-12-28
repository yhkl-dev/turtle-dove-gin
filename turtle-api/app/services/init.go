package services

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/yhkl-dev/turtle-dove-beego/turtle-api/app/tables"
)

var (
	o                 orm.Ormer
	r                 redis.Conn
	RedisService      *redisService
	UserService       *userService
	RoleService       *roleService
	PermissionService *permissionService
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
		new(tables.Permission),
	)
	orm.RunSyncdb("default", false, true)

	o = orm.NewOrm()
	orm.RunCommand()

	redisHost := beego.AppConfig.String("redis.host")
	redisPort := beego.AppConfig.String("redis.port")
	redisDB := beego.AppConfig.String("redis.db")
	if redisDB == "" {
		redisDB = "0"
	}
	if redisPort == "" {
		redisPort = "6379"
	}
	rsdn := fmt.Sprintf("redis://%s:%s/%s", redisHost, redisPort, redisDB)
	r, _ = redis.DialURL(rsdn)
	//defer r.Close()

	initService()
}

func initService() {
	UserService = &userService{}
	RoleService = &roleService{}
	PermissionService = &permissionService{}
	RedisService = &redisService{}
}

func tableName(name string) string {
	return name
}
