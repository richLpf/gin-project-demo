package middleware

import (
	"net/http"
	"fmt"

	"github.com/gin-gonic/gin"
)

//Cors 跨域中间件
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Origin, Authorization, Content-Type,remote_user,X-Requested-With,*")
		c.Header("Access-Control-Allow-Headers", "Content-Type, remote_user")
		c.Header("Access-Control-Allow-Methods", "POST,OPTIONS, GET,PUT")
		c.Header("Access-Control-Expose-Headers", "Accept , Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "True")
		c.Header("Content-Type", "application/json; charset=utf-8")
		fmt.Println("method", method)
		if method == "OPTIONS" {
			c.AbortWithStatus(204)
			c.AbortWithStatus(http.StatusNoContent)
		}
		submitUser := c.Request.Header.Get("remoteUser")
		c.Set("submitUser", submitUser)
		c.Next()
	}
}
