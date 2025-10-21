// class - 11
package main

import "fmt"

func anonymousFunc() {
	// anonymous function
	func(num1 int, num2 int) {
		sum := num1 + num2

		fmt.Println(sum)
	}(5, 5) // IIFE - Immediatly Invoked(Call) Function Expression

	// or
	// anonymous function
	sum := func(num1 int, num2 int) {
		sum := num1 + num2

		fmt.Println(sum)
	}

	sum(5, 5)
}
