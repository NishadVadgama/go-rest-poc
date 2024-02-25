package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/NishadVadgama/go-server-poc/models"
)

// Get articles route
//
// Route: /articles
func GetArticlesRoute() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK) 
		json.NewEncoder(w).Encode(models.Articles)
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
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnprocessableEntity) 
			json.NewEncoder(w).Encode(models.Response{
				Message: "Invalid data!",
			})
			return
		}
		
		// Make article id
		article.Id = 1
		if len(models.Articles) > 0 {
			article.Id = models.Articles[len(models.Articles)-1].Id + 1
		}
		models.Articles = append(models.Articles, article)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK) 
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
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest) 
			json.NewEncoder(w).Encode(models.Response{
				Message: "Invalid parameter!",
			})
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
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound) 
			json.NewEncoder(w).Encode(models.Response{
				Message: "No article found!",
			})
			return
		}

		// Return article
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK) 
		json.NewEncoder(w).Encode(models.Articles[index])
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
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest) 
			json.NewEncoder(w).Encode(models.Response{
				Message: "Invalid parameter!",
			})
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
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound) 
			json.NewEncoder(w).Encode(models.Response{
				Message: "Unable to delete article!",
			})
			return
		}
		
		// Remove article from a slice
		models.Articles = append(models.Articles[:index], models.Articles[index+1:]...)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK) 
		json.NewEncoder(w).Encode(models.Response{
			Message: "Article deleted successfully!",
		})
	}
}

// Update article by id route
//
// Route: /articles/{id}
func UpdateArticleByIdRoute() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var id = r.PathValue("id")

		// Parse article id
		articleId, err := strconv.Atoi(id)
		if err != nil || articleId == 0 {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest) 
			json.NewEncoder(w).Encode(models.Response{
				Message: "Invalid parameter!",
			})
			return
		}

		// Validate article data
		var article models.Article
        json.NewDecoder(r.Body).Decode(&article)
		if article.Title == "" {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnprocessableEntity) 
			json.NewEncoder(w).Encode(models.Response{
				Message: "Invalid data!",
			})
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
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound) 
			json.NewEncoder(w).Encode(models.Response{
				Message: "Unable to find article!",
			})
			return
		}
		
		// Update article
		models.Articles[index].Title = article.Title
		models.Articles[index].Description = article.Description
		models.Articles[index].Tags = article.Tags
		
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK) 
		json.NewEncoder(w).Encode(models.Response{
			Message: "Article updated successfully!",
		})
	}
}