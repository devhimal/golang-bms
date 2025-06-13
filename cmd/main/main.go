
package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/devhimal/golang-bms/internal/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBooks(r)
	http.Handle("/", r)
	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

