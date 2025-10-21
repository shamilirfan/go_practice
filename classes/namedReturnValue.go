package main

import "fmt"

func returnValue(a int, b int) int {
	result := a + b
	return result
}

func namedReturnValue(a int, b int) (result int) {
	result = a + b
	return
}

func named_Return_Value() {

	fmt.Println("Normal return value =", returnValue(5, 5))
	fmt.Println("Named return value =", namedReturnValue(5, 5))
}
