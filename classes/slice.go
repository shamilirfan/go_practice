package main

import "fmt"

func Slice() {

	// Slice is a part of array
	// 1.
	arr1 := [6]string{"This", "is", "a", "Go", "interview", "questions"}
	slice1 := arr1[1:4] // A slice holdes 3 things, 1. Pointer, 2. Length, 3. Capacity

	fmt.Println(slice1)

	// 2. Slice of slice
	slice2 := slice1[2:5]

	fmt.Println(slice2)

	// 3. Slice literal
	slice3 := []int{1, 2, 3}

	fmt.Println(slice3)

	// 4. Slice creating by make function with size
	slice4 := make([]int, 4)

	fmt.Println(slice4)

	// 5. Slice creating by make function with size and capacity
	slice5 := make([]int, 4, 6)

	fmt.Println(slice5)
	fmt.Println("Length =", len(slice5))
	fmt.Println("Capacity =", cap(slice5))

	// 6. Empty slice or nil slice
	var slice6 []int

	fmt.Println(slice6)

	// add new element in slice
	color := []string{"Red", "Green", "Blue"}
	color = append(color, "Yellow", "Pink")
	color = append([]string{"Black"}, color...)

	fmt.Println(color)
}
