package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get all users
func GetUsers(c *gin.Context) {
	var users []User
	result := db.Find(&users)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

// Create a new user (POST /register)
func CreateUser(c *gin.Context) {
	var user User

	// Bind form data (JSON or x-www-form-urlencoded)
	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println("Bind Error:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Insert into database
	if err := db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully", "user": user})
}
