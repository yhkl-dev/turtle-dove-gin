package router

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	v1 "github.com/yhkl-dev/turtle-dove-beego/gin-turtle-api/router/api/v1"
	"github.com/yhkl-dev/turtle-dove-beego/gin-turtle-api/utils/middleware/jwtauth"
)

// InitRouter ini router
func InitRouter() *gin.Engine {
	router := gin.Default()

	setUpConfig(router)
	setUpRouter(router)
	return router
}

func setUpConfig(router *gin.Engine) {
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Use(jwtauth.JWTAuthMiddleware())
}

func setUpRouter(router *gin.Engine) {
	api := router.Group("/api")
	{
		v1.RegisterRouter(api)
	}

}
