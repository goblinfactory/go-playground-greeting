package simon

import "fmt"

// Greet will say hello
func Greet() {
	notExported()
}

func notExported() {
	fmt.Println("Hello")
}
