package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var database *gorm.DB
var e error

func DatabaseInit() {

	// host := "db"
	// port := 3306
	// user := "root"
	// password := "admin123"
	// dbName := "merodiscount-booking"

	// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta", host, user, password, dbName, port)
	// database, e = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	dsn := "root:admin123@tcp(db:3306)/merodiscount-booking?charset=utf8&parseTime=true" // Replace with your details
	// Connect to the database
	database, e = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if e != nil {
		panic(e)
	}
	fmt.Println("database connected")
}

func DB() *gorm.DB {
	return database
}
