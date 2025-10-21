// class - 2
package main

import "fmt"

func variable() {
	a := 5
	var b = 10
	var c string = "Hello" // var is changeable
	c = "Hi"
	const d string = "Sunflower" // const is not changeable
	var e int = 15
	var f bool = true
	const g bool = false
	var h float32 = 20.55

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
}
