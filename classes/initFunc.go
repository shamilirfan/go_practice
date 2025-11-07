// class - 10
package classes

import "fmt"

var a = 10

func abc() {
	a = 20
	fmt.Println(a)
}

func initFunction() {
	a = 30
	abc()

	fmt.Println(a) // print will be 20
	a = 30
	fmt.Println(a) // print will be 30
}

func init() {
	fmt.Println(a)
}
