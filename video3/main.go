package main

import (
	"example/video3/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
	This session covers:

1.What is a middleware.
2. Custom (Request-Response) Middlewares in GIN.
3. Different ways to write middleware functions in Go and their use.
4. How to apply Middleware to individual routes, routes group and whole application at once in GIN.
*/
func GetHello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Say hello",
	})
}
func GetMessage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Message": "2nd type of middleware",
	})
}
func GetGrouphandler(c *gin.Context) {
	c.JSON(http.StatusAccepted, gin.H{
		"message": "group middleware handler",
	})
}

// 2. Custom (Request-Response) Middlewares in GIN.
func main() {
	router := gin.New()
	securityauth := gin.BasicAuth(gin.Accounts{
		"prasuna": "1212",
	})
	/* 	router.Use(middleware.Authenticate, middleware.AddHeader) // applies to all API's
	   	router.GET("/hello", GetHello) */
	router.Use(gin.Logger()) // to get logs in the terminal
	router.GET("/hello", GetHello, middleware.Authenticate, middleware.AddHeader, securityauth)
	router.GET("/hello2", GetMessage, middleware.Authenticate2()) // other type of writing middleware
	//applying the middleware to the route groups
	admin := router.Group("/admin", securityauth)
	{
		admin.GET("/hellogroup", GetGrouphandler)
	}

	router.Run()
}
