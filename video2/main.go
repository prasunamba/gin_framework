package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type student struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Course string `json:"course"`
}

var students = []student{
	{Id: "one", Name: "prasuna", Course: "Mtech"},
	{Id: "two", Name: "rahul", Course: "primary"},
	{Id: "three", Name: "satwik", Course: "playschool"},
}

func Getstudents(c *gin.Context) {
	c.JSON(http.StatusOK, students)
}

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
		admin.GET("/students", Getstudents)
	}
	server := &http.Server{
		Addr:         ":9091",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	server.ListenAndServe()
}
