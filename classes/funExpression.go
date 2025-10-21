package classes

import "fmt"

func funcExpression() {
    // ফাংশন এক্সপ্রেশন 
    greet := func(name string) string {
        return "Hello, " + name
    }

    fmt.Println(greet("Shamil"))
}
