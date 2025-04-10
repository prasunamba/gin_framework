package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

/*
	This session covers:

1. Custom HTTP configuration with GIN
2. Route Grouping in GIN.
3. Using Basic Auth functionality with GIN
*/
func main() {
	router := gin.Default()

	auth := gin.BasicAuth(gin.Accounts{
		"user": "pass",
	})
	admin := router.Group("/admin", auth)
	{
		admin.GET("/Students", getStudents)
	}
	server := &http.Server{
		Addr:         ":9091",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	server.ListenAndServe()
}
