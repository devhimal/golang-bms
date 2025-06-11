
package services

import "github.com/devhimal/golang-bms/internal/models"

func GetBooks() []models.Book {
	return []models.Book{
		{ID: 1, Title: "1984", Author: "George Orwell"},
		{ID: 2, Title: "The Hobbit", Author: "J.R.R. Tolkien"},
	}
}
