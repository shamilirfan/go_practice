package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func ResetPassword(w http.ResponseWriter, r *http.Request, db *sql.DB) {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
        return
    }

    token := r.FormValue("token")
    password := r.FormValue("password")

    if len(password) < 8 {
        http.Error(w, "Password must be at least 8 characters", http.StatusBadRequest)
        return
    }

    var userID int
    var expires time.Time
    err := db.QueryRow("SELECT user_id, expires_at FROM password_resets WHERE token=$1", token).Scan(&userID, &expires)
    if err != nil {
        http.Error(w, "Invalid token", http.StatusBadRequest)
        return
    }

    if time.Now().After(expires) {
        http.Error(w, "Token expired", http.StatusBadRequest)
        return
    }

    hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    _, err = db.Exec("UPDATE users SET password=$1 WHERE id=$2", string(hash), userID)
    if err != nil {
        fmt.Println("DB update error:", err)
        http.Error(w, "Failed to update password", http.StatusInternalServerError)
        return
    }

    _, _ = db.Exec("DELETE FROM password_resets WHERE token=$1", token)

    w.Write([]byte("✅ Password reset successful"))
}
