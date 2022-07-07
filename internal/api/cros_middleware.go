package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func crosMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		fmt.Println("method", method)
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization, Host")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Host, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		if method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}
