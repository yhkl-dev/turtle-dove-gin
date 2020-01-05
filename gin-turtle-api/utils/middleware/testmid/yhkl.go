package testmid

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func TestMidware(c *gin.Context) {
	fmt.Println("yhkl  in middle ware")

	c.Next()
}
