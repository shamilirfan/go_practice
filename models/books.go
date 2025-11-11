package models

type Book struct {
	ID          int      `json:"id" db:"id"`
	Title       string   `json:"title" db:"title"`
	Author      string   `json:"author" db:"author"`
	Price       int      `json:"price" db:"price"`
	Description string   `json:"description" db:"description"`
	Category    string   `json:"category" db:"category"`
	Brand       string   `json:"brand" db:"brand"`
	IsStock     bool     `json:"is_stock" db:"is_stock"`
	ImagesUrl   []string `json:"images_url" db:"images_url"`
}
