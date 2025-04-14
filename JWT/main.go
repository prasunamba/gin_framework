package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.POST("/login", Login)
	router.GET("/home", GetHome)
	router.GET("/refresh", Refreshtoken)
	router.Run(":8080")
}
