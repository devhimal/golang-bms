package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func ParseRequestBody(r *http.Request, v interface{}) error {
	body, err := io.ReadAll(r.Body) // Read the entire request body
	if err != nil {
		return err // Return error if reading failed
	}

	// Parse JSON from body into v (which is usually a pointer to a struct)
	if err := json.Unmarshal(body, v); err != nil {
		return err // Return error if JSON unmarshalling failed
	}

	return nil // Success, no errors
}

