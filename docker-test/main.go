package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHello(c *gin.Context) {
	fmt.Println("hello")
	c.JSON(http.StatusOK, gin.H{
		"msg": "hello from docker",
	})
}
func main() {
	router := gin.Default()
	router.GET("/hello", GetHello)
	router.Run(":8080")
}
