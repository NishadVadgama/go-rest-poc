package main

import (
	"fmt"
	"net/http"

	"github.com/NishadVadgama/go-server-poc/controllers"
	"github.com/NishadVadgama/go-server-poc/utils"
)

func main() {
	// Init db
	var pgdb utils.PostgresDB
	conn, err := pgdb.InitializeDB("user=postgres dbname=go-rest-poc sslmode=disable")
	if err != nil {
		fmt.Println("Error while initializing db: ", err.Error())
		return
	}

	// Initialize router
	var mux = http.NewServeMux()

	// Bind Routes
	//
	// base route
	mux.HandleFunc("GET /", utils.Handler(controllers.GetIndexRoute()))

	// articles route
	mux.HandleFunc("GET /articles", utils.Handler(controllers.GetArticlesRoute(conn)))
	mux.HandleFunc("POST /articles", utils.Handler(controllers.CreateArticleRoute(conn)))
	mux.HandleFunc("GET /articles/{id}", utils.Handler(controllers.GetArticleByIdRoute(conn)))
	mux.HandleFunc("DELETE /articles/{id}", utils.Handler(controllers.DeleteArticleByIdRoute(conn)))
	mux.HandleFunc("PUT /articles/{id}", utils.Handler(controllers.UpdateArticleByIdRoute(conn)))

	// Starting listener
	fmt.Println("Server starting at http://localhost:3333/")
	http.ListenAndServe(":3333", mux)
}
