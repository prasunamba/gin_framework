package api

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// User model
type User struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Name  string `json:"name"`
	Email string `gorm:"unique" json:"email"`
	Age   int    `json:"age"`
}

var db *gorm.DB

// creating  database
func init() {
	var err error
	// mysql username:myuser,password:mypassword,prasuna-databasename
	dsn := "myuser:mypassword@tcp(127.0.0.1:3306)/prasuna"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil || db == nil {
		log.Fatal("Failed to connect to MySQL:", err, db)
	}

	// Auto-create the users table if not exists
	db.AutoMigrate(&User{})
	fmt.Println("Database connected & Users table created!")

	// Insert a record
	users := []User{
		{Name: "Alice", Email: "alice@example.com", Age: 25},
		{Name: "Bob", Email: "bob@example.com", Age: 30},
		{Name: "Charlie", Email: "charlie@example.com", Age: 22},
		{Name: "David", Email: "david@example.com", Age: 28},
	}

	db.Create(&users) // Batch insert
	fmt.Println("User inserted successfully!")
}
