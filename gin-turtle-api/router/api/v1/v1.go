package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/yhkl-dev/turtle-dove-beego/gin-turtle-api/apps/role"
	"github.com/yhkl-dev/turtle-dove-beego/gin-turtle-api/apps/user"
)

// RegisterRouter 注册路由
func RegisterRouter(router *gin.RouterGroup) {
	v1 := router.Group("/v1")
	{
		user.RegisterRouter(v1.Group("/user"))
		role.RegisterRouter(v1.Group("/role"))
	}
}
