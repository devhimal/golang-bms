
package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/devhimal/golang-bms/internal/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterRoutes(r)

	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

