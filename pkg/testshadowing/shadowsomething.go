package testshadowing

import "fmt"

// ShadowSomething contains some code that shadows a variable that will compile but should
// trigger our make file which runs shadow ./... and should detect this and flag an error.
func ShadowSomething() {

	x := 10
	fmt.Println("x is", x)
	for i := 0; i < 10; i++ {
		// shadow x below
		x := i + 1
		fmt.Println(i, x)
	}
}
