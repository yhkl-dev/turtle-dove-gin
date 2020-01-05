package permission

import (
	"github.com/gin-gonic/gin"
	"github.com/yhkl-dev/turtle-dove-beego/gin-turtle-api/utils/response"
)

// ListAllPermissions return permissions list
func ListAllPermissions(c *gin.Context) {
	service := ListPermissionsService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.ListAllPermissions()
		c.JSON(200, res)
	} else {
		c.JSON(200, response.ErrorResponse(err))

	}
}
