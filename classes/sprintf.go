package classes

import "fmt"

func sprintf() {
	name := "Rupa"
	city := "Khulna"
	age := 25

	result := fmt.Sprintf("My name is %s. I'm %d years old. I live in %s.", name, age, city)
	fmt.Println(result)
}
