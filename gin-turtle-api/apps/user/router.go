package user

import (
	"github.com/gin-gonic/gin"
)

// RegisterRouter 注册路由
func RegisterRouter(r *gin.RouterGroup) {
	r.POST("/login", UserLogin)
	r.GET("", ListAllUsers)
	r.GET("/:id", GetUserProfile)
	r.POST("/register", UserRegister)
	r.PUT("/:id", UpdateUserInfo)
	r.DELETE("/:id", DeleteUser)
}
