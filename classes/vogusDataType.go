package main

import "fmt"

func vogusDataType() {
	// int bit depend on the PC bit. If PC is 32 bit, int 32 bit. If PC is 64 bit, int 64 bit
	var a int = 10
	var b int8 = 127 // int8 range is -128 to 127

	// u means unsigned. It will not contain negative(-255) number. Range: 0 to 255.
	var c uint8 = 255
	var d float32 = 10.23348
	var e bool = true // bool is 8 bit

	// Go এ rune আসলে int32 এর আরেকটা নাম। তাই rune এর সাইজ 4 byte = 32 bit
	var smile rune = '😊'
	var f rune = 'ঠ'

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%.2f\n", d) // %f is floating format.
	fmt.Print(e, "\n")
	fmt.Printf("%c\n", smile) // %c is rune/character format.
	fmt.Printf("%c\n", f)
	fmt.Printf("%T", f) // %T is type format.

}
