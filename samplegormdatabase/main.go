package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

type family struct {
	ID       int `gorm:"primarykey"`
	Name     string
	Relation string
	Age      int
}

func main() {
	var err error
	dsn := "myuser:yourpassword@tcp(127.0.0.1:3306)/prasuna" //step1
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})     // step2 -connect to database
	if err != nil || db == nil {
		fmt.Println("")
	}
	fmt.Println("connected successfully")
	if err = db.AutoMigrate(&family{}); err != nil { // step3 -create table
		fmt.Println("failed to create table")
	}
	fmt.Println("table created sucessfully")
	fam := []family{
		{Name: "basavamma", Age: 45, Relation: "mom"},
		{Name: "krishna rao", Age: 53, Relation: "dad"},
		{Name: "mahesh", Age: 28, Relation: "bro"},
	}
	id := db.Create(&fam) // step4 -insert data into table
	fmt.Println("Insert result:", id.RowsAffected, id.Error)

	var famslice []family
	result := db.Where("Age=?", 28).Find(&famslice)
	if result.Error != nil {
		fmt.Println("Error fetching data:", result.Error)
		return
	}
	for _, val := range famslice {
		fmt.Println("result", val)
	}

}
