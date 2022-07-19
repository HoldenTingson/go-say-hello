package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// SetupDB : initializing mysql database

func main() {
	SetupDB()
}

func SetupDB() *gorm.DB {
	USER := "root"
	PASS := ""
	HOST := "localhost"
	PORT := "3306"
	DBNAME := "arclight"
	URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)
	db, err := gorm.Open("mysql", URL)
	if err != nil {
		panic(err.Error())
	}
	return db
}
