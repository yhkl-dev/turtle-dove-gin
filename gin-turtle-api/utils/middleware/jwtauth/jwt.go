package jwtauth

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/yhkl-dev/turtle-dove-beego/gin-turtle-api/database"
	"github.com/yhkl-dev/turtle-dove-beego/gin-turtle-api/database/models"
	"github.com/yhkl-dev/turtle-dove-beego/gin-turtle-api/utils/response"
	"github.com/yhkl-dev/turtle-dove-beego/gin-turtle-api/utils/serializers"
)

// JWTAuthMiddleware middle ware for auth
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if c.Request.RequestURI == "/api/v1/user/login" || c.Request.RequestURI == "/api/v1/user/register" {
			c.Next()
			return
		}

		fmt.Println("token", token)
		if token == "" {
			c.JSON(200, serializers.ErrorHandler(403, "Authtication failed. Please login", errors.New("Authtication failed")))
			c.Abort()
			return
		}
		claims, err := ParseToken(token)
		if err != nil {

			c.JSON(200, response.ErrorResponse(err))
			c.Abort()
			return
		}
		perList := PermissionJudge(claims.RoleList)
		permissionURI := c.Request.Method + ":" + c.FullPath()

		var permissionObject models.Permission

		err = database.DB.Where("permission_name = ?", permissionURI).First(&permissionObject).Error
		if err != nil || !isExist(permissionObject.ID, perList) {
			c.JSON(200, serializers.ErrorHandler(403, "no permission", errors.New("Permission Denied")))
			c.Abort()
			return
		}
	}
}

func isExist(element int, eList []int) bool {
	for _, e := range eList {
		if element == e {
			return true
		}
	}
	return false
}

// PermissionJudge
func PermissionJudge(roleList []int) []int {
	var (
		permissionList []int
		distinctList   []int
		roles          []models.Role
	)
	err := database.DB.Preload("Permissions").Where("id in (?)", roleList).Find(&roles).Error
	if err != nil {
		fmt.Println(err)
	}

	for _, role := range roles {
		for _, permission := range role.Permissions {
			permissionList = append(permissionList, permission.ID)
		}
	}

	distinctList = RemoveDuplicatesAndEmpty(permissionList)
	return distinctList
}

// RemoveDuplicatesAndEmpty
func RemoveDuplicatesAndEmpty(a []int) (ret []int) {
	aLen := len(a)
	for i := 0; i < aLen; i++ {
		if i > 0 && a[i-1] == a[i] {
			continue
		}
		ret = append(ret, a[i])
	}
	return
}
