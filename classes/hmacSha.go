package classes

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
)

func hmacSha() {
	secret_key := []byte("secret-key")
	message := []byte("Hello")

	h := hmac.New(sha256.New, secret_key)
	h.Write(message)
	hash := h.Sum(nil)

	fmt.Println("HmacSha-256 = ", hash)
}
