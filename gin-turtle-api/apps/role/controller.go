package role

import (
	"github.com/gin-gonic/gin"
	"github.com/yhkl-dev/turtle-dove-beego/gin-turtle-api/utils/response"
)

// ListAllUsers return user list
func ListAllRoles(c *gin.Context) {
	service := RoleListService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.ListAllRoles()
		c.JSON(200, res)
	} else {
		c.JSON(200, response.ErrorResponse(err))
	}
}

// AddRole
// @Titile Add a role
// @success 200
// @router / [post]
func AddRole(c *gin.Context) {
	var service AddRoleService

	if err := c.ShouldBind(&service); err == nil {
		res := service.AddRole()
		c.JSON(200, res)

	} else {
		c.JSON(200, response.ErrorResponse(err))
	}
}

// UpdateRole
// @router /:id [put]
func UpdateRole(c *gin.Context) {
	var service UpdateRoelService

	if err := c.ShouldBind(&service); err == nil {
		res := service.UpdateRole(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, response.ErrorResponse(err))
	}
}

// DeleteRole
// @router /:id [delete]
func DeleteRole(c *gin.Context) {
	var service DeleteRoleService
	if err := c.ShouldBind(&service); err == nil {
		res := service.DeleteRole(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, response.ErrorResponse(err))
	}
}
