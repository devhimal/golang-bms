
package routes

import (
	"github.com/gorilla/mux"
	"github.com/devhimal/golang-bms/internal/controllers"
)

var RegisterBooks = func (router *mux.Router){
	router.HandleFunc("/books", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/books/{id}", controllers.GetBookByID).Methods("GET")
	router.HandleFunc("/books", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/books/{id}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", controllers.DeleteBook).Methods("DELETE")
}
