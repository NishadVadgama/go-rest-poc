package controllers

import (
	"database/sql"
	"encoding/json"
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
			utils.JSONResponse(w, http.StatusInternalServerError, models.Response{Message: err.Error()})
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

		// Make article id
		article.Id = 1
		if len(models.Articles) > 0 {
			article.Id = models.Articles[len(models.Articles)-1].Id + 1
		}
		models.Articles = append(models.Articles, article)

		// Return response
		utils.JSONResponse(w, http.StatusOK, models.Response{Message: "Article added successfully!"})
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
			utils.JSONResponse(w, http.StatusBadRequest, models.Response{Message: "Invalid parameter!"})
			return
		}

		// Finding article
		var index = -1
		for i, v := range models.Articles {
			if v.Id == articleId {
				index = i
			}
		}

		// Check if we found article
		if index == -1 {
			utils.JSONResponse(w, http.StatusNotFound, models.Response{Message: "No article found!"})
			return
		}

		// Return article
		utils.JSONResponse(w, http.StatusOK, models.Articles[index])
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
			utils.JSONResponse(w, http.StatusBadRequest, models.Response{Message: "Invalid parameter!"})
			return
		}

		// Finding article
		var index = -1
		for i, v := range models.Articles {
			if v.Id == articleId {
				index = i
			}
		}

		// Check if we found article
		if index == -1 {
			utils.JSONResponse(w, http.StatusInternalServerError, models.Response{Message: "Unable to delete article!"})
			return
		}

		// Remove article from a slice
		models.Articles = append(models.Articles[:index], models.Articles[index+1:]...)

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

		// Parse article id
		articleId, err := strconv.Atoi(id)
		if err != nil || articleId == 0 {
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

		// Finding article
		var index = -1
		for i, v := range models.Articles {
			if v.Id == articleId {
				index = i
			}
		}

		// Check if we found article
		if index == -1 {
			utils.JSONResponse(w, http.StatusNotFound, models.Response{Message: "Unable to find article!"})
			return
		}

		// Update article
		models.Articles[index].Title = article.Title
		models.Articles[index].Description = article.Description
		models.Articles[index].Tags = article.Tags

		// Return response
		utils.JSONResponse(w, http.StatusOK, models.Response{Message: "Article updated successfully!"})
	}
}
