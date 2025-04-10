package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Student struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Course string `json:"course"`
}

func getStudents(c *gin.Context) {
	var Studentslice []Student
	DB.Find(&Studentslice)
	// for _, val := range Studentslice {

	// }
	c.JSON(http.StatusOK, Studentslice)
}
func init() {
	var err error
	dsn := "myuser:yourpassword@tcp(127.0.0.1:3306)/prasuna"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("failed to connect to database")
	}
	fmt.Println("connected to db")
	DB.AutoMigrate(&Student{})
	var Students = []Student{
		{Id: "one", Name: "prasuna", Course: "Mtech"},
		{Id: "two", Name: "rahul", Course: "primary"},
		{Id: "three", Name: "satwik", Course: "playschool"},
	}
	DB.Create(&Students)

}
