package role

import (
	"github.com/gin-gonic/gin"
	"github.com/yhkl-dev/turtle-dove-beego/gin-turtle-api/utils/response"
)

// ListAllUsers return user list
func ListAllRoles(c *gin.Context) {
	//	service := UserListService{}
	//	if err := c.ShouldBind(&service); err == nil {
	//		res := service.ListAllUsers()
	//		c.JSON(200, res)
	//	} else {
	//		c.JSON(200, response.ErrorResponse(err))
	//}
	c.JSON(200, "role")
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
