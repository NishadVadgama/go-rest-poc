package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/NishadVadgama/go-server-poc/controllers"
	"github.com/NishadVadgama/go-server-poc/pkg/config"
	"github.com/NishadVadgama/go-server-poc/utils"
)

func main() {
	// For DB seed
	var shouldDBSeed = flag.Bool("db_seed", false, "Perform database seed?")

	// Parse the flags
	flag.Parse()

	// Init db
	var pgdb utils.PostgresDB
	conn, err := pgdb.InitializeDB(config.GetPGConnectionString())
	if err != nil {
		log.Println("Error while initializing db: ", err.Error())
		return
	}

	// Seed articles
	if *shouldDBSeed {
		err = pgdb.SeedArticles()
		if err != nil {
			log.Println("Error while seeding articles: ", err.Error())
			return
		}
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
	log.Printf("Server starting at http://localhost:%s/\n", config.GetPort())
	http.ListenAndServe(":"+config.GetPort(), mux)
}
