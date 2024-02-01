package config

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

// Connect initializes a database connection.
func Connect() {
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, dbName)

	d, err := gorm.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}

	db = d
}

// CloseDB closes the database connection.
func CloseDB() {
	if db != nil {
		sqlDB := db.DB()
		err := sqlDB.Close()
		if err != nil {
			fmt.Println("Error closing database connection:", err)
			return
		}
	}
}

// GetDB returns the current database instance.
func GetDB() *gorm.DB {
	return db
}
