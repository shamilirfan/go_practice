package classes

import (
	"encoding/base64"
	"fmt"
)

func Base() {
	var a string = "DD"
	b := []byte(a)

	base := base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(b)
	decode, _ := base64.URLEncoding.WithPadding(base64.NoPadding).DecodeString(base)

	fmt.Println("Byte =", b)
	fmt.Println("Base64 =", base)
	fmt.Println("Decode =", decode)
	fmt.Println("Decode =", string(decode))
}
