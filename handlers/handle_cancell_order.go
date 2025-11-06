package handlers

import (
	"encoding/json"
	"go_practice/db"
	"net/http"
)

func HandleCancellOrder(w http.ResponseWriter, r *http.Request) {
	if db.DB == nil {
		db.Connect()
	}

	if err := UpdateOrder(db.DB, r); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("✅Successfully Cancelled")
}
