package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// pre-middleware
func Authenticate(c *gin.Context) {
	if !(c.Request.Header.Get("Token") != "Auth") {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "token not present",
		})
		return
	}
	c.Next()
}

// 2nd type of writing middleware
// to apply custom logic to specific routes or pass parameters to the middleware
func Authenticate2() gin.HandlerFunc {
	// write custom logic to be applied  before my middleware is executed
	//used in realtime usecases
	return func(c *gin.Context) {
		if !(c.Request.Header.Get("Token") != "Auth") {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "token not present",
			})
			return
		}
		c.Next()
	}
}

// post middleware
func AddHeader(c *gin.Context) {
	c.Writer.Header().Set("key", "value")
	c.Next()
}
