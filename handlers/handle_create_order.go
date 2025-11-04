package handlers

import (
	"encoding/json"
	"go_practice/db"
	"go_practice/models"
	"net/http"
)

func HandleCreateOrder(w http.ResponseWriter, r *http.Request) {
	if db.DB == nil {
		db.Connect()
	}

	var order models.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := CreateOrder(db.DB, &order); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}
