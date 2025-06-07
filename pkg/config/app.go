package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv" // Import godotenv
	"log"                       // For logging errors from godotenv
	"os"
)

var (
	db *gorm.DB
)

func Connect() {
	// Load .env file first
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, using default environment variables or hardcoded values.")
	}

	// Get database connection info from environment variable
	databaseUSER := os.Getenv("DBUSERINFO")

	// If environment variable is not set, use default connection string
	if databaseUSER == "" {
		databaseUSER = "root:@tcp(localhost:3306)/bookstore?charset=utf8&parseTime=True&loc=Local"
		log.Println("DBUSERINFO not set, using default connection string.")
	}

	d, err := gorm.Open("mysql", databaseUSER)

	if err != nil {
		panic(err)
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}