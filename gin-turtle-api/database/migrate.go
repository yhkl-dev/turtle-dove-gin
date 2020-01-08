package database

import (
	"fmt"

	"github.com/yhkl-dev/turtle-dove-gin/gin-turtle-api/database/models"
)

type initPermission struct {
	CodeName       string
	PermissionName string
}

func migrate() {
	fmt.Println("Migrating tables....")
	DB.AutoMigrate(&models.Permission{}, &models.User{}, &models.Role{})

	fmt.Println("Init Permissions")
	DB.Exec("truncate table sys_permission")
	models.User{}.RegisterPermission(DB)
	models.Role{}.RegisterPermission(DB)
}
