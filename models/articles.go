package models

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

var Articles []Article

// Load dummy articles from json
func init() {
	// Open the JSON file
	file, err := os.Open("data/sample.json")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close() // it will close the file at the end of the function

	// Read the file content
	bytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Unmarshal the JSON data into the Person struct
	err = json.Unmarshal(bytes, &Articles)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}

// Get all articles from db
func GetArticles(conn *sql.DB) ([]Article, error) {
	// An album slice to hold data from returned rows.
	var articles = []Article{}

	// fetch articles
	rows, err := conn.Query(`SELECT * from articles LIMIT 1`)
	if err != nil {
		return articles, err
	}
	defer rows.Close()

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var article Article
		if err := rows.Scan(&article.Id, &article.Title, &article.Description,
			&article.Tags); err != nil {
			return articles, err
		}
		articles = append(articles, article)
	}
	if err = rows.Err(); err != nil {
		return articles, err
	}
	return articles, nil
}
