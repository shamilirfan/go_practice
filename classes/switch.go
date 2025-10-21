// class - 4
package classes

import "fmt"

func switcH() {
	var mark int = 56

	switch true {
	case (mark > 100 || mark < 0):
		fmt.Println("Invalid marks")
	case mark >= 80:
		fmt.Println("A+")
	case mark >= 70:
		fmt.Println("A")
	case mark >= 60:
		fmt.Println("A-")
	case mark >= 50:
		fmt.Println("B")
	case mark >= 40:
		fmt.Println("C")
	case mark >= 33:
		fmt.Println("D")
	default:
		fmt.Println("Fail")
	}
}
