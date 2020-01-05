package user

import (
	"github.com/gin-gonic/gin"
	"github.com/yhkl-dev/turtle-dove-beego/gin-turtle-api/utils/response"
)

// ListAllUsers return user list
func ListAllUsers(c *gin.Context) {
	service := UserListService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.ListAllUsers()
		c.JSON(200, res)
	} else {
		c.JSON(200, response.ErrorResponse(err))
	}
}

// @Title user register
// UserRegister user register
// @router /register [post]
func UserRegister(c *gin.Context) {
	var service UserRegieterService

	if err := c.ShouldBind(&service); err == nil {
		res := service.Register()
		c.JSON(200, res)
	} else {
		c.JSON(200, response.ErrorResponse(err))
	}
}

// GetUserProfile return user profile
func GetUserProfile(c *gin.Context) {
	service := GetUserProfileService{}
	res := service.GetUserProfile(c.Param("id"))
	c.JSON(200, res)
}

// UpdateUserInfo update user profile
func UpdateUserInfo(c *gin.Context) {
	var service UpdateUserService

	if err := c.ShouldBind(&service); err == nil {
		res := service.UpdateUser(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, response.ErrorResponse(err))
	}
}

// DeleteUser set user is_deleted 1
func DeleteUser(c *gin.Context) {
	var service DeleteUserService
	if err := c.ShouldBind(&service); err == nil {
		res := service.DeleteUser(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, response.ErrorResponse(err))
	}
}

// UserLogin for user login
func UserLogin(c *gin.Context) {
	var service UserLoginService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Login()
		c.JSON(200, res)
	} else {
		c.JSON(200, response.ErrorResponse(err))
	}

}
