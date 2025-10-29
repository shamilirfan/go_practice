package handlers

import (
	"encoding/json"
	"net/http"
	"os"

	"go_practice/db"

	"github.com/gorilla/mux"
)

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	// Delete image from server
	var imagePath string
	db.DB.QueryRow("SELECT image_path FROM books WHERE id=$1", id).Scan(&imagePath)
	if imagePath != "" {
		os.Remove(imagePath)
	}

	_, err := db.DB.Exec("DELETE FROM books WHERE id=$1", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var message string = "Successfully Deleted!"
	json.NewEncoder(w).Encode(&message)
	w.WriteHeader(http.StatusOK)
}
