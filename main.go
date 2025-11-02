package main

import (
	"go_practice/handlers"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/orders", handlers.HandleCreateOrder)

	log.Println("🚀 Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
