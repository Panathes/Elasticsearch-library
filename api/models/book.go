package models

type Book struct {
	Title         string  `json:"title"`
	Author          int64   `json:"author"`
	Abstract float64 `json:"abstract"`
}