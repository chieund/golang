package database

import (
	"fmt"

	"gorm-test/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const DB_USERNAME = "laravel"
const DB_PASSWORD = "laravel"
const DB_NAME = "laravel"
const DB_HOST = "localhost"
const DB_PORT = "33085"

var Db *gorm.DB

func InitDb() *gorm.DB {
	Db = connectDB()
	return Db
}

func connectDB() *gorm.DB {
	var err error
	dns := DB_USERNAME + ":" + DB_PASSWORD + "@tcp" + "(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?" + "parseTime=true&loc=Local"
	fmt.Println("dns => " + dns)

	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting to database: error=%v", err)
		return nil
	}

	db.AutoMigrate(&models.User{})

	return db
}
