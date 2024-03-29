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

// Get article by id
func GetArticle(conn *sql.DB, id int) (Article, error) {
	// Article
	var article = Article{}

	// Fetch article
	row := conn.QueryRow(`SELECT * from articles WHERE id = $1`, id)

	// Using Scan to assign column data to struct fields
	err := row.Scan(&article.Id, &article.Title, &article.Description, &article.Tags)
	if err != nil && err == sql.ErrNoRows {
		return article, nil
	}

	// Return article or error
	return article, err
}

// Get all articles
func GetArticles(conn *sql.DB) ([]Article, error) {
	// Articles
	var articles = []Article{}

	// Fetch articles
	rows, err := conn.Query(`SELECT * from articles ORDER BY id DESC`)
	if err != nil {
		return articles, err
	}
	defer rows.Close()

	// Loop through rows, using Scan to assign column data to struct fields
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

// Create new article
func CreateArticle(conn *sql.DB, article Article) (int, error) {
	var newArticleId int

	// Insert article
	err := conn.QueryRow(
		`INSERT INTO articles (title, description, tags) VALUES ($1, $2, $3) RETURNING id`,
		article.Title,
		article.Description,
		article.Tags).Scan(&newArticleId)
	if err != nil {
		return 0, err
	}

	return newArticleId, err
}

// Delete article
func DeleteArticle(conn *sql.DB, id int) (int64, error) {
	res, err := conn.Exec(`DELETE FROM articles WHERE id = $1`, id)
	if err != nil {
		return 0, err
	}

	// get affected rows count
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, err
}

// Update article
func UpdateArticle(conn *sql.DB, id int, article Article) (int64, error) {
	res, err := conn.Exec(
		`UPDATE articles SET title = $1, description = $2, tags = $3 WHERE id = $4`,
		article.Title,
		article.Description,
		article.Tags,
		id)
	if err != nil {
		return 0, err
	}

	// get affected rows count
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, err
}
