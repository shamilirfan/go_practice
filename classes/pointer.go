// class - 14
package main

import "fmt"

func printArray(a *[4]string) {
	fmt.Println(a)
}

func Pointer() {
	// Example of pointer - 1. Pointer means address/reference of memory
	x := 10
	pointer := &x // & => ampersand

	fmt.Println(pointer)
	fmt.Println(*pointer) // value at pointer

	// Example of pointer - 2
	var flowers = [4]string{"Sunflower", "Jasmine", "Rose", "Marigold"}

	printArray(&flowers)
}
