// class - 8
package classes

import "fmt"

var number int = 10

func varShadowing() {
	age := 30

	if age > 18 {
		number := 47
		fmt.Println(number)
	}

	fmt.Println(number)
}
