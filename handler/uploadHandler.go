package handler

import (
	"encoding/json"
	"strconv"

	"go_practice/connection"
	"go_practice/models"
	"net/http"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

var Cld *cloudinary.Cloudinary

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}

	title := r.FormValue("title")
	author := r.FormValue("author")
	price, _ := strconv.Atoi(r.FormValue("price"))
	description := r.FormValue("description")
	category := r.FormValue("category")
	brand := r.FormValue("brand")
	isStock, _ := strconv.ParseBool(r.FormValue("is_stock"))

	var bookID int
	query := `
        INSERT INTO books 
        (title, author, price, description, category, brand, is_stock)
        VALUES ($1, $2, $3, $4, $5, $6, $7)
        RETURNING id
    `
	err = connection.DB.QueryRow(
		query,
		title,
		author,
		price,
		description,
		category,
		brand,
		isStock,
	).Scan(&bookID)
	if err != nil {
		http.Error(w, "Database insert failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	files := r.MultipartForm.File["images_url"]
	var imageUrls []string

	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			http.Error(w, "Error reading file", http.StatusInternalServerError)
			return
		}
		defer file.Close()

		uploadResult, err := Cld.Upload.Upload(r.Context(), file, uploader.UploadParams{
			Folder: "books",
		})
		if err != nil {
			http.Error(w, "Cloudinary upload failed", http.StatusInternalServerError)
			return
		}

		imageUrls = append(imageUrls, uploadResult.SecureURL)
	}

	// Update the book row with images
	query = `UPDATE books SET images_url=$1 WHERE id=$2`
	_, err = connection.DB.Exec(query, pq.Array(imageUrls), bookID)
	if err != nil {
		http.Error(w, "Failed to update images", http.StatusInternalServerError)
		return
	}

	book := models.Book{
		ID:          bookID,
		Title:       title,
		Author:      author,
		Price:       price,
		Description: description,
		Category:    category,
		Brand:       brand,
		IsStock:     isStock,
		ImagesUrl:   imageUrls,
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}
