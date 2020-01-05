package database

import (
	"fmt"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/yhkl-dev/turtle-dove-beego/gin-turtle-api/conf"
)

var (
	DB      *gorm.DB
	DBError error
)

func init() {
	dbHost, err := conf.SysConfig.GetValue("Database", "host")
	dbPort, err := conf.SysConfig.GetValue("Database", "port")
	dbUserName, err := conf.SysConfig.GetValue("Database", "username")
	dbPassword, err := conf.SysConfig.GetValue("Database", "password")
	dbName, err := conf.SysConfig.GetValue("Database", "db")
	dbConnetionsLimit, err := conf.SysConfig.GetValue("Database", "connections")

	if err != nil {
		panic(err)
	}
	if dbPort == "" {
		dbPort = "3306"
	}

	dbLimits, _ := strconv.Atoi(dbConnetionsLimit)

	dsn := dbUserName + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8&parseTime=True&loc=Local"

	DB, DBError = gorm.Open("mysql", dsn)
	if DBError != nil {
		fmt.Println(DBError)
	}
	DB.DB().SetMaxIdleConns(dbLimits)
	DB.DB().SetMaxOpenConns(dbLimits)
	DB.DB().SetConnMaxLifetime(time.Second * 30)
	migrate()
}
