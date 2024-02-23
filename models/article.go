package models

type Article struct {
	Id    int      `json:"id"`
	Title string   `json:"title"`
	Description  string   `json:"description"`
	Tags  []string `json:"tags"`
}