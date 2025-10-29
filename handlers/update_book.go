package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go_practice/db"
	"go_practice/models"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	/*
			এই কোডটা URL-এর path parameter (যেমন /books/{id}) থেকে id মানটা বের করে আনে।
			mux.Vars(r) একটা map রিটার্ন করে (যেখানে সব path parameter থাকে)
			যেমন: map[string]string{"id": "10"}
		তারপর params["id"] দিয়ে তুমি "10" মানটা বের করছো।
	*/
	params := mux.Vars(r)
	id := params["id"]

	r.ParseMultipartForm(10 << 20)

	name := r.FormValue("name")
	author := r.FormValue("author")
	priceStr := r.FormValue("price")
	price, _ := strconv.ParseFloat(priceStr, 64)
	file, handler, err := r.FormFile("image")

	var filePath string
	if err == nil {
		defer file.Close()
		os.MkdirAll("uploads", os.ModePerm)
		filePath = filepath.Join("uploads", handler.Filename)
		filePath = filepath.ToSlash(filePath)
		dst, _ := os.Create(filePath)
		defer dst.Close()
		io.Copy(dst, file)

		// Delete old image
		var oldImage string
		db.DB.QueryRow(
			`
		SELECT image_path 
		FROM books 
		WHERE id=$1
		`,
			id,
		).Scan(&oldImage)
		if oldImage != "" {
			os.Remove(oldImage)
		}
	} else {
		db.DB.QueryRow(
			`
			SELECT image_path 
			FROM books 
			WHERE id=$1
			`,
			id,
		).Scan(&filePath)
	}

	_, err = db.DB.Exec(
		`
		UPDATE books
		SET name=$1, author=$2, price=$3, image_path=$4 
		WHERE id=$5
		`,
		name, author, price, filePath, id,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	idInt, _ := strconv.Atoi(id)
	book := models.Book{
		ID:     idInt,
		Name:   name,
		Author: author,
		Price:  price,
		Image:  filePath,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}
