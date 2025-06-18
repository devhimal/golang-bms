
package models

import (
	"github.com/jinzhu/gorm"
	"github.com/devhimal/golang-bms/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Publication string `json:"publication"`
	Price	float64 `json:"price"`
}

func init(){
	config.ConnectDB()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func ( b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookByID(Id int) (*Book, *gorm.DB) {
	var getBook Book
	result := db.Where("ID=?", Id).Find(&getBook)
	return &getBook,result
}


func DeleteBook(Id int) Book {
	var book Book
	 db.Where("ID=?", Id).Delete(book)
	return book
}

