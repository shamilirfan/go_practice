package handler

import (
	"encoding/json"
	"go_practice/connection"
	"net/http"
	"strconv"
)

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	strId := r.PathValue("id")
	id, _ := strconv.Atoi(strId)
	query := `DELETE FROM books WHERE id = $1`

	_, err := connection.DB.Exec(query, id)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode("✅ Successfully Deleted")
}
