package database

import (
	"github.com/yhkl-dev/turtle-dove-beego/gin-turtle-api/database/models"
)

func migrate() {
	//DB.AutoMigrate(&user.User{})
	//DB.AutoMigrate(&role.Role{})
	DB.AutoMigrate(&models.User{}, &models.Role{}, &models.Permission{})
	//	DB.AutoMigrate(&models.User{})
}
