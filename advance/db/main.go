package main

import (
	"github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	UserName string
}

type Animal struct {
	ID int64
	Name string
	Age int64
}

func main() {
	db, err := gorm.Open("mysql", "laravel:laravel@(localhost:33085)/laravel?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connection database")
	}

	defer db.Close()

	// db.CreateTable(&User{})
	// db.HasTable(&User{})

	var animal = Animal{Age: 14, Name: "Lion"}
	db.Create(&animal)

	// get first
	db.First(&Animal())
}
