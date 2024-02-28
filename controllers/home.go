package controllers

import (
	"net/http"

	"github.com/NishadVadgama/go-server-poc/models"
	"github.com/NishadVadgama/go-server-poc/utils"
)

// Get handler for base route
// Route: /
func GetIndexRoute() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Default response
		utils.JSONResponse(w, http.StatusOK, models.Response{
			Message: "API is live!",
		})
	}
}
