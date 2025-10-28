package models

type Book struct {
	ID     int     `json:"id" db:"id"`
	Name   string  `json:"name" db:"name"`
	Author string  `json:"author" db:"author"`
	Price  float64 `json:"price" db:"price"`
	Image  string  `json:"image" db:"image_path"` // Path in server
}
