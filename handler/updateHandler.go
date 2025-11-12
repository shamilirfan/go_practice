package handler

import (
	"encoding/json"
	"go_practice/connection"
	"go_practice/models"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/lib/pq"
)

func extractPublicID(imageURL string) string {
	parts := strings.Split(imageURL, "/upload/")
	if len(parts) < 2 {
		return ""
	}

	publicPath := parts[1]
	publicPath = strings.SplitN(publicPath, ".", 2)[0]
	publicPath = strings.TrimSuffix(publicPath, filepath.Ext(publicPath))
	return publicPath
}

func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	idStr := r.PathValue("id")
	id, _ := strconv.Atoi(idStr)
	title := r.FormValue("title")
	author := r.FormValue("author")
	price, _ := strconv.Atoi(r.FormValue("price"))
	description := r.FormValue("description")
	category := r.FormValue("category")
	brand := r.FormValue("brand")
	isStock, _ := strconv.ParseBool(r.FormValue("is_stock"))

	if err != nil {
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}

	query := `
	UPDATE books SET 
	title=$1, author=$2, price=$3, description=$4, category=$5, brand=$6, is_stock=$7 
	WHERE id=$8
	`
	_, err = connection.DB.Exec(query,
		title,
		author,
		price,
		description,
		category,
		brand,
		isStock,
		id,
	)
	if err != nil {
		http.Error(w, "Failed to update book info", http.StatusInternalServerError)
		return
	}

	var oldImages []string
	query = `SELECT images_url FROM books WHERE id=$1`
	connection.DB.QueryRow(query, id).Scan(pq.Array(&oldImages))

	for _, imgURL := range oldImages {
		publicID := extractPublicID(imgURL)
		if publicID == "" {
			continue
		}

		_, err := Cld.Upload.Destroy(r.Context(), uploader.DestroyParams{
			PublicID: publicID,
		})
		if err != nil {
			http.Error(w, "Failed to delete:", http.StatusInternalServerError)
			return
		}
	}

	files := r.MultipartForm.File["images_url"]
	var newImageUrls []string

	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			http.Error(w, "Error reading file", http.StatusInternalServerError)
			return
		}

		uploadResult, err := Cld.Upload.Upload(r.Context(), file, uploader.UploadParams{
			Folder: "books",
		})
		file.Close()

		if err != nil {
			http.Error(w, "Cloudinary upload failed", http.StatusInternalServerError)
			return
		}

		newImageUrls = append(newImageUrls, uploadResult.SecureURL)
	}

	query = `UPDATE books SET images_url=$1 WHERE id=$2`
	_, err = connection.DB.Exec(query, pq.Array(newImageUrls), id)
	if err != nil {
		http.Error(w, "Failed to update images", http.StatusInternalServerError)
		return
	}

	book := models.Book{
		ID:          id,
		Title:       title,
		Author:      author,
		Price:       price,
		Description: description,
		Category:    category,
		Brand:       brand,
		IsStock:     isStock,
		ImagesUrl:   newImageUrls,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&book)
}
