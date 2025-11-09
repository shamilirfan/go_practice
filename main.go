package main

// bo12ok34sh56op78

import (
	"go_practice/connection"
	"go_practice/handlers"
	"net/http"
)

func main() {
	if connection.DB == nil {
		connection.InitDB()
	}

	http.HandleFunc("/request-reset", func(w http.ResponseWriter, r *http.Request) {
		handlers.RequestPasswordReset(w, r, connection.DB)
	})

	http.HandleFunc("/reset-password-form", handlers.ResetPasswordForm)

	http.HandleFunc("/reset-password", func(w http.ResponseWriter, r *http.Request) {
		handlers.ResetPassword(w, r, connection.DB)
	})

	http.ListenAndServe(":8080", nil)
}
