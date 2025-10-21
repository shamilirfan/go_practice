// class - 7
package classes

import "fmt"

func inputName() string {
	fmt.Print("Enter your name: ")
	var name string
	fmt.Scanln(&name)

	return name
}

func getTwoNumbers() (int, int) {
	fmt.Print("Enter two number: ")
	var num1, num2 int
	fmt.Scanln(&num1, &num2) // & is ampersand

	return num1, num2
}

func sum() int {
	num1, num2 := getTwoNumbers()
	add := num1 + num2

	return add
}

func printName() { fmt.Println("Hi, " + inputName() + "!") }
func printSum()  { fmt.Println("Sum =", sum()) }

func input() {
	printName()
	printName()

	printSum()
	printSum()
}
