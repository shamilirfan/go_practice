package main

import (
	"fmt"
	"go_practice/connection"
	"go_practice/handler"
	"log"
	"net/http"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/joho/godotenv"
)

func main() {
	// ✅ Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// ✅ Connect Database
	if connection.DB == nil {
		connection.Connection()
	}

	// ✅ Initialize Cloudinary
	handler.Cld, err = cloudinary.NewFromURL(os.Getenv("CLOUDINARY_URL"))
	if err != nil {
		log.Fatalf("Cloudinary initialization failed: %v", err)
	}

	fmt.Println("✅ Cloudinary initialized successfully")

	// ✅ Routes
	http.HandleFunc("/books", handler.UploadHandler)

	fmt.Println("🚀 Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
