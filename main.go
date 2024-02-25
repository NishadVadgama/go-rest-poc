package main

import (
	"fmt"
	"net/http"

	"github.com/NishadVadgama/go-server-poc/controllers"
	"github.com/NishadVadgama/go-server-poc/utils"
)

func main() {
	mux := http.NewServeMux()
	
	// Routes
	//
	// base route
	mux.HandleFunc("GET /", utils.Logger(controllers.GetIndexRoute()))

	// articles route
	mux.HandleFunc("GET /articles", utils.Logger(controllers.GetArticlesRoute()))
	mux.HandleFunc("POST /articles", utils.Logger(controllers.CreateArticleRoute()))
	mux.HandleFunc("GET /articles/{id}", utils.Logger(controllers.GetArticleByIdRoute()))
	mux.HandleFunc("DELETE /articles/{id}", utils.Logger(controllers.DeleteArticleByIdRoute()))
	mux.HandleFunc("PUT /articles/{id}", utils.Logger(controllers.UpdateArticleByIdRoute()))

	// Starting listener
	fmt.Println("Server starting at http://localhost:3333/")
	http.ListenAndServe(":3333", mux)
}