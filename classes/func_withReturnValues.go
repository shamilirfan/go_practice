// class - 6
package main

import "fmt"

// single value return
func Add(num1 int, num2 int) int {
	var sum int = num1 + num2

	return sum
}

// multi value return
func getNumbers(num1 int, num2 int) (int, int, int) {
	subtraction := num1 - num2
	multiplication := num1 * num2
	devide := num1 / num2

	return subtraction, multiplication, devide
}

func func_withReturnValues() {
	// Add function
	fmt.Println("Add function")
	fmt.Println(Add(10, 20))

	sum1 := Add(10, 30)
	sum2 := Add(10, 40)

	fmt.Println(sum1)
	fmt.Println(sum2)

	// getNumbers function
	subtraction, multiplication, devide := getNumbers(20, 10)

	fmt.Println()
	fmt.Println("getNumbers function")
	fmt.Println(subtraction)
	fmt.Println(multiplication)
	fmt.Println(devide)
}
