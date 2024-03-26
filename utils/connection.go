package connection

import (
	"fmt"
	"os"

	// Import the driver for your specific database (e.g., mysql, postgres)
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func connectionToMySql() {

	// user := os.Getenv("DB_USER")
	// password := os.Getenv("DB_PASSWORD")
	// host := os.Getenv("DB_HOST")
	// port := os.Getenv("DB_PORT")
	// database := os.Getenv("DB_NAME")

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai",
		host, port, user, password, dbname)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("connection", db)
	// return db

	// Use the db instance for database operations
}

func connectionToMySql() (*gorm.DB, error) {

	// user := os.Getenv("DB_USER")
	// password := os.Getenv("DB_PASSWORD")
	// host := os.Getenv("DB_HOST")
	// port := os.Getenv("DB_PORT")
	// database := os.Getenv("DB_NAME")
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
	}

	// host := os.Getenv("DB_HOST")
	// port := os.Getenv("DB_PORT")
	// user := os.Getenv("DB_USER")
	// password := os.Getenv("DB_PASSWORD")
	// dbname := os.Getenv("DB_NAME")

	// host := "db"
	// port := "3306"
	// user := "root"
	// password := "admin123"
	// dbname := "merodiscount-booking"

	dsn := "root:admin123@tcp(db:3306)/merodiscount-booking?charset=utf8&parseTime=true" // Replace with your details

	// Connect to the database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.Facility{})
	fmt.Println("Successfully connected to database!")
	return db, nil
}
