package main

import (
	"go_practice/db"
	"go_practice/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	db.Connect()

	r := mux.NewRouter()

	r.HandleFunc("/books", handlers.CreateBook).Methods("POST")
	// r.HandleFunc("/books", handlers.GetBooks).Methods("GET")
	// r.HandleFunc("/books/{id}", handlers.GetBook).Methods("GET")
	// r.HandleFunc("/books/{id}", handlers.UpdateBook).Methods("PUT")
	// r.HandleFunc("/books/{id}", handlers.DeleteBook).Methods("DELETE")

	// Serve uploaded images
	r.PathPrefix("/uploads/").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads"))))

	log.Println("Server running at :8080")
	http.ListenAndServe(":8080", r)
}
