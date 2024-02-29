package utils

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/NishadVadgama/go-server-poc/models"
	_ "github.com/lib/pq"
)

type PostgresDB struct {
	Instance *sql.DB
}

// To initialize postgres db
func (pg *PostgresDB) InitializeDB(connectionString string) (*sql.DB, error) {
	// Connect to db
	var err error
	pg.Instance, err = sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	// Ping db and throw error if not connected
	if err = pg.Instance.Ping(); err != nil {
		return nil, err
	}

	log.Println("Database connected.")
	return pg.Instance, nil
}

// Seed articles
func (pg *PostgresDB) SeedArticles() error {
	var articles []models.Article

	// Open the JSON file
	file, err := os.Open("pkg/db/seed/articles.json")
	if err != nil {
		return err
	}
	defer file.Close() // it will close the file at the end of the function

	// Read the file content
	bytes, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	// Unmarshal the JSON data into the Person struct
	err = json.Unmarshal(bytes, &articles)
	if err != nil {
		return err
	}

	// Return if not articles to seed
	if len(articles) == 0 {
		return nil
	}

	// Prepare SQL statement
	log.Println("Seeding articles...")
	sqlStr := "INSERT INTO articles(title, description, tags) VALUES "
	values := []interface{}{}

	// Feed in values
	for i, row := range articles {
		sqlStr += fmt.Sprintf("($%d, $%d, $%d),", (i*3)+1, (i*3)+2, (i*3)+3)
		values = append(values, row.Title, row.Description, row.Tags)
	}

	// Trim the last ,
	sqlStr = sqlStr[0 : len(sqlStr)-1]

	stmt, err := pg.Instance.Prepare(sqlStr)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(values...)

	// Return error if any
	return err
}
