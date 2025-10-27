package classes

import (
	"crypto/sha256"
	"fmt"
)

func Sha() {
	// sha - secure hash algorithm
	var a string = "Hello"
	b := []byte(a)

	hash := sha256.Sum256(b)

	fmt.Println("Sha-256 =", hash)
}
