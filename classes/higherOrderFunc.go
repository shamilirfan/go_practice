package main

import "fmt"

// Higher order function. Recieve a function by the paramiter
func higherOrder(c int, d int, e func(x int, y int)) {
	e(c, d)
}

func add(a int, b int) {
	sum := a + b

	fmt.Println(sum)
}

// Higher order function. Recieve a function by the return
func returnHigherOrder() func(a int, b int) {
	return add
}

func higherOrderFunc() {

	higherOrder(5, 5, add)
	sum := returnHigherOrder()

	sum(5, 5)
}
