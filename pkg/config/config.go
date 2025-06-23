package config

import (
	"github.com/jinzhu/gorm"
)

// Placeholder for config vars

var (
	db *gorm.DB
)

func ConnectDB() {
	d, err := gorm.Open("mysql", "devhimal:devhimal123/simplerest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect to database")
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}

var AppName = "Golang Book Management System"
