package jwtauth

import (
	"github.com/gin-gonic/gin"
)

// JWTAuthMiddleware middle ware for auth
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//	token := c.Request.Header.Get("Authorization")
		if c.Request.RequestURI == "/api/v1/user/login" || c.Request.RequestURI == "/api/v1/user/register" {
			c.Next()
			return
		}
		//		if token == "" {
		//			c.JSON(200, serializers.ErrorHandler(403, "Authtication failed. Please login", errors.New("Authtication failed")))
		//			c.Abort()
		//			return
		//
		//		}
		//		fmt.Println("token", token)

	}
}
