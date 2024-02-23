package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/NishadVadgama/go-server-poc/models"
)

// Get handler for base route
// Route: /
func GetIndexRoute() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Default response
		var res = models.Response{
			Message: "Hello from the go API!",
		}

		w.WriteHeader(http.StatusOK) // set status code
		w.Header().Add("Content-Type", "application/json") // set response type
		json.NewEncoder(w).Encode(res) // response
	}
}