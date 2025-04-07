package main

import (
	api "maingorm/APIfolder"

	"github.com/gin-gonic/gin"
)

func main() {

	// Initialize Gin router
	r := gin.Default()

	r.POST("/register", api.CreateUser) // to register user
	// Define GET API endpoint
	r.GET("/users", api.GetUsers) // to get users

	// Serve HTML form - to register user using the html form [it indirectly calls the api.CreateUser api]
	r.GET("/fromform", func(c *gin.Context) {
		c.File("./APIfolder/index.html")
	})
	// Start the server
	r.Run(":8080") // Runs on http://localhost:8080
}
