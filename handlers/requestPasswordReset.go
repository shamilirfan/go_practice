package handlers

import (
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"go_practice/utils"
	"net/http"
	"time"
)

type ResetRequest struct {
	Email string `json:"email"`
}

func generateToken() string {
	/*
		১. b := make([]byte, 32) -
			👉 এটা ৩২ বাইট লম্বা একটা byte slice তৈরি করছে।
				মানে, এখানে আমরা ৩২ বাইটের একটা জায়গা নিচ্ছি random ডেটা রাখার জন্য।

		২. rand.Read(b) -
			👉 এটা crypto/rand প্যাকেজ থেকে আসছে।
				এটা	(CSPRNG) দিয়ে ওই b slice টা random bytes দিয়ে পূরণ করে।

		৩. return hex.EncodeToString(b) -
			👉 যেহেতু b এখন binary data (৩২ বাইট), এটাকে hexadecimal string
				এ রূপান্তর করা হচ্ছে।
				ফলাফল হবে ৬৪ অক্ষরের একটি নিরাপদ টোকেন (কারণ ১ বাইট = ২ হেক্স ক্যারেক্টার)।
	*/

	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println(err)
	}

	return hex.EncodeToString(b)
}

func RequestPasswordReset(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	w.Header().Set("Content-Type", "application/json")

	var req ResetRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Check if user exists
	var userID int
	err := db.QueryRow("SELECT id FROM users WHERE email=$1", req.Email).Scan(&userID)
	if err != nil {
		http.Error(w, "No user found with this email", http.StatusNotFound)
		return
	}

	// Generate token & expiry
	token := generateToken()
	expires := time.Now().Add(15 * time.Minute)

	// Store token in DB
	_, err = db.Exec(`
		INSERT INTO password_resets (user_id, token, expires_at)
		VALUES ($1, $2, $3)
	`, userID, token, expires)
	if err != nil {
		http.Error(w, "Failed to create reset token", http.StatusInternalServerError)
		return
	}

	// Create the reset link
	resetLink := fmt.Sprintf("http://localhost:8080/reset-password-form?token=%s", token)

	// Send email using updated SendEmail (HTML formatted)
	subject := "Reset your BookShop password"
	if err := utils.SendEmail(req.Email, subject, resetLink); err != nil {
		fmt.Println("Email error:", err)
		http.Error(w, "Failed to send email", http.StatusInternalServerError)
		return
	}

	// Success response
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Password reset link sent to your email",
	})
}
