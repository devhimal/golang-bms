
package models

import (
	"github.com/jinzhu/gorm"
	"github.com/devhimal/golang-bms/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	ID     int    `json:"id"`
	Title  string `gorm:""json:"title"`
	Author string `json:"author"`
	Publication string `json:"publication"`
}

func init(){
	config.ConnectDB()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

