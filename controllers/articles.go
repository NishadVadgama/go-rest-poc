package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/NishadVadgama/go-server-poc/data"
	"github.com/NishadVadgama/go-server-poc/models"
)

// Get articles route
//
// Route: /articles
func GetArticlesRoute() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK) 
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data.Articles)
	}
}

// Create articles route
//
// Route: /articles
func CreateArticleRoute() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var article models.Article
        json.NewDecoder(r.Body).Decode(&article)
		
		if article.Title == "" {
			w.WriteHeader(http.StatusUnprocessableEntity) 
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(models.Response{
				Message: "Invalid data!",
			})
			return
		}
		
		article.Id = len(data.Articles) + 1
		data.Articles = append(data.Articles, article)

		w.WriteHeader(http.StatusOK) 
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(models.Response{
			Message: "Article added successfully!",
		})
	}
}

// Get article by id route
//
// Route: /articles/{id}
func GetArticleByIdRoute() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var id = r.PathValue("id")

		// Parse article id
		articleId, err := strconv.Atoi(id)
		if err != nil || articleId == 0 {
			w.WriteHeader(http.StatusBadRequest) 
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(models.Response{
				Message: "Invalid parameter!",
			})
			return
		}

		// Finding article
		var article models.Article
		for _, v := range data.Articles {
			if v.Id == articleId {
				article = v
			}
		}

		// Check if we found article
		if article.Id == 0 {
			w.WriteHeader(http.StatusNotFound) 
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(models.Response{
				Message: "No article found!",
			})
			return
		}

		// Return article
		w.WriteHeader(http.StatusOK) 
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(article)
	}
}

// Delete article by id route
//
// Route: /articles/{id}
func DeleteArticleByIdRoute() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var id = r.PathValue("id")

		// Parse article id
		articleId, err := strconv.Atoi(id)
		if err != nil || articleId == 0 {
			w.WriteHeader(http.StatusBadRequest) 
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(models.Response{
				Message: "Invalid parameter!",
			})
			return
		}

		// Finding article
		var article models.Article
		for _, v := range data.Articles {
			if v.Id == articleId {
				article = v
			}
		}

		// Check if we found article
		if article.Id == 0 {
			w.WriteHeader(http.StatusNotFound) 
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(models.Response{
				Message: "No article found!",
			})
			return
		}

		w.WriteHeader(http.StatusOK) 
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(models.Response{
			Message: "Article deleted successfully!",
		})
	}
}