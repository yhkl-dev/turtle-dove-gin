package role

import (
	"github.com/gin-gonic/gin"
)

// RegisterRouter 注册路由
func RegisterRouter(r *gin.RouterGroup) {
	r.GET("", ListAllRoles)
	r.POST("", AddRole)
	r.PUT("/:id", UpdateRole)
	r.DELETE("/:id", DeleteRole)
}
