// class - 3
package main

import "fmt"

func ifelse() {
	var mark int = 70

	if mark > 100 || mark < 0 {
		fmt.Println("Invalid marks")
	} else if mark >= 80 {
		fmt.Println("A+")
	} else if mark >= 70 {
		fmt.Println("A")
	} else if mark >= 60 {
		fmt.Println("A-")
	} else if mark >= 50 {
		fmt.Println("B")
	} else if mark >= 40 {
		fmt.Println("C")
	} else if mark >= 33 {
		fmt.Println("D")
	} else {
		fmt.Println("Fail")
	}

}
