package utils

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

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

	log.Println("Database connected!")
	return pg.Instance, nil
}

// To push schema db
func (pg *PostgresDB) PushSchema(file string) error {
	sqlFilePath := filepath.Join(file)

	// Read SQL file
	sqlBytes, err := os.ReadFile(sqlFilePath)
	if err != nil {
		return err
	}

	// Convert bytes to string
	sql := string(sqlBytes)

	// Execute SQL statements
	_, err = pg.Instance.Exec(sql)
	if err != nil {
		return err
	}

	return nil
}

// Seed articles
func (pg *PostgresDB) SeedArticles() error {
	var articles []models.Article

	// Open the JSON file
	file, err := os.Open("data/sample.json")
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
	sqlStr := "INSERT INTO articles(title, description) VALUES "
	values := []interface{}{}

	// Feed in values
	for _, row := range articles {
		sqlStr += "(?, ?, ?),"
		fmt.Printf("%T", row.Tags)
		values = append(values, row.Title, row.Description, row.Tags)
	}

	// Trim the last ,
	sqlStr = sqlStr[0 : len(sqlStr)-1]
	fmt.Print(sqlStr)
	fmt.Print(values...)

	stmt, err := pg.Instance.Prepare(sqlStr)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(values...)

	// Return error if any
	return err
}
