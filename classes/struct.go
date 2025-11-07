// class - 12
package classes

import "fmt"

type User struct { // User means instance or object
	Name string
	Age  int
}

func printInstance(a User) {
	fmt.Println("Name:", a.Name+", "+"Age:", a.Age)
}

// Receiver Function
func (a User) receiverFunc() {
	fmt.Println("Name:", a.Name+", "+"Age:", a.Age)
}

func Struct() {

	var user1 User = User{Name: "Alice", Age: 20}
	var user2 User = User{Name: "Bob", Age: 35}
	var user3 User = User{Name: "Rupa", Age: 28}

	printInstance(user1)
	printInstance(user2)

	// Receiver Function Call
	user3.receiverFunc()
}
