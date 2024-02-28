package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	_ "github.com/lib/pq"

	"github.com/NishadVadgama/go-server-poc/models"
	"github.com/NishadVadgama/go-server-poc/utils"
)

// Get articles route
//
// Route: /articles
func GetArticlesRoute(conn *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var articles, err = models.GetArticles(conn)
		if err != nil {
			log.Printf("Error: %v\n", err)
			utils.JSONResponse(w, http.StatusInternalServerError, models.Response{Message: "Error occurred while fetching articles!"})
			return
		}

		// Return articles
		utils.JSONResponse(w, http.StatusOK, articles)
	}
}

// Create articles route
//
// Route: /articles
func CreateArticleRoute(conn *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var article models.Article
		json.NewDecoder(r.Body).Decode(&article)

		if article.Title == "" {
			utils.JSONResponse(w, http.StatusUnprocessableEntity, models.Response{Message: "Invalid data!"})
			return
		}

		insertedId, err := models.CreateArticle(conn, article)
		if err != nil || insertedId == 0 {
			utils.JSONResponse(w, http.StatusInternalServerError, models.Response{Message: "Error occurred while creating article!"})
			return
		}

		// Set inserted
		article.Id = int(insertedId)

		// Return response
		utils.JSONResponse(w, http.StatusOK, article)
	}
}

// Get article by id route
//
// Route: /articles/{id}
func GetArticleByIdRoute(conn *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var id = r.PathValue("id")

		// Parse article id
		articleId, err := strconv.Atoi(id)
		if err != nil || articleId == 0 {
			if err != nil {
				log.Printf("Error parsing string to integer: %v\n", err)
			}
			utils.JSONResponse(w, http.StatusBadRequest, models.Response{Message: "Invalid parameter!"})
			return
		}

		// Check if we found article
		article, err := models.GetArticle(conn, articleId)
		if err != nil {
			log.Printf("Error: %v\n", err)
			utils.JSONResponse(w, http.StatusInternalServerError, models.Response{Message: "Error occurred while finding article!"})
			return
		}

		if article.Id == 0 {
			utils.JSONResponse(w, http.StatusNotFound, models.Response{Message: "No article found!"})
			return
		}

		// Return article
		utils.JSONResponse(w, http.StatusOK, article)
	}
}

// Delete article by id route
//
// Route: /articles/{id}
func DeleteArticleByIdRoute(conn *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var id = r.PathValue("id")

		// Parse article id
		articleId, err := strconv.Atoi(id)
		if err != nil || articleId == 0 {
			if err != nil {
				log.Printf("Error parsing string to integer: %v\n", err)
			}
			utils.JSONResponse(w, http.StatusBadRequest, models.Response{Message: "Invalid parameter!"})
			return
		}

		rowsAffected, err := models.DeleteArticle(conn, articleId)
		if err != nil {
			log.Printf("Error: %v\n", err)
			utils.JSONResponse(w, http.StatusInternalServerError, models.Response{Message: "Error occurred while deleting article!"})
			return
		}
		if rowsAffected == 0 {
			utils.JSONResponse(w, http.StatusNotFound, models.Response{Message: "No article found, that can be deleted!"})
			return
		}

		// Return response
		utils.JSONResponse(w, http.StatusOK, models.Response{Message: "Article deleted successfully!"})
	}
}

// Update article by id route
//
// Route: /articles/{id}
func UpdateArticleByIdRoute(conn *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var id = r.PathValue("id")

		var (
			articleId int
			err       error
		)

		// Parse article id
		articleId, err = strconv.Atoi(id)
		if err != nil || articleId == 0 {
			if err != nil {
				log.Printf("Error parsing string to integer: %v\n", err)
			}
			utils.JSONResponse(w, http.StatusBadRequest, models.Response{Message: "Invalid parameter!"})
			return
		}

		// Validate article data
		var article models.Article
		json.NewDecoder(r.Body).Decode(&article)
		if article.Title == "" {
			utils.JSONResponse(w, http.StatusUnprocessableEntity, models.Response{Message: "Invalid data!"})
			return
		}

		_, err = models.UpdateArticle(conn, articleId, article)
		if err != nil {
			log.Printf("Error: %v\n", err)
			utils.JSONResponse(w, http.StatusInternalServerError, models.Response{Message: "Error occurred while updating article!"})
			return
		}

		// Return response
		utils.JSONResponse(w, http.StatusOK, models.Response{Message: "Article updated successfully!"})
	}
}
