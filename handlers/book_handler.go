package handlers

import (
	"encoding/json"
	//"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	//"github.com/gorilla/mux"
	"go_practice/db"
	"go_practice/models"
)

// CreateBook with image upload
func CreateBook(w http.ResponseWriter, r *http.Request) {
	/* এই লাইনটা মূলত Go তে ফাইল বা ফর্ম ডেটা আপলোড করার সময় ব্যবহৃত হয়।
	10 << 20 মানে হলো —
	“এই রিকোয়েস্টের ফর্মের ডেটা (ফাইলসহ) পার্স করো,
	তুমি সার্ভারকে বলছো,
	“ফর্ম ডেটা পার্স করার সময় ১০ মেগাবাইট পর্যন্ত RAM ব্যবহার করো।” */
	r.ParseMultipartForm(10 << 20) // 10 MB max

	name := r.FormValue("name")
	author := r.FormValue("author")
	priceStr := r.FormValue("price")
	price, _ := strconv.ParseFloat(priceStr, 64)

	/*
		file(multipart.File) - ফাইলের ডেটা স্ট্রিম
		(যেটা তুমি পড়তে পারো বা সেভ করতে পারো)
		   handler(*multipart.FileHeader) - ফাইল সম্পর্কিত তথ্য
		   (নাম, সাইজ, টাইপ ইত্যাদি)
	*/
	file, handler, err := r.FormFile("image")

	var filePath string
	if err == nil {
		defer file.Close()
		// Create uploads folder if not exist
		os.MkdirAll("uploads", os.ModePerm)
		filePath = filepath.Join("uploads", handler.Filename) // filePath = uploads/image-1.jpg

		/*
			মানে হলো —
			একটা নতুন ফাইল তৈরি করো,
			আর যদি ওই নামে আগে থেকেই ফাইল থাকে,
			তাহলে পুরনোটা মুছে দিয়ে নতুনটা লেখার জন্য খুলে দাও।
			dst(*os.File) - তৈরি হওয়া ফাইলের অবজেক্ট।
		*/
		dst, _ := os.Create(filePath)
		defer dst.Close()

		/* যখন তুমি ইউজারের আপলোড করা ফাইল সার্ভারে সেভ করতে চাও,
		   তখন এই লাইনটা আবশ্যক
		   আপলোড করা ফাইলের সব বাইট dst তে কপি করে
		*/
		io.Copy(dst, file)
	}

	query :=
		`
	INSERT INTO books (name, author, price, image_path) 
	VALUES ($1, $2, $3, $4) 
	RETURNING id
	`

	var id int
	err = db.DB.QueryRow(query, name, author, price, filePath).Scan(&id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	book := models.Book{
		ID:     id,
		Name:   name,
		Author: author,
		Price:  price,
		Image:  filePath,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}
