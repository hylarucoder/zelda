package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GlobalMiddleware() gin.HandlerFunc {
	fmt.Println("globalMiddleware...1")

	return func(c *gin.Context) {
		fmt.Println("globalMiddleware...2")
		c.Next()
		fmt.Println("globalMiddleware...3")
	}
}
