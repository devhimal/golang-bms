
package controllers

import (
	"encoding/json"
	"net/http"
)

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode([]string{"Book 1", "Book 2", "Book 3"})
}
