package switchy

import "fmt"

// TestSwitchy becauuse...
func TestSwitchy() {
loop:
	for i := 1; i <= 15; i++ {
		switch {
		case i%5 == 0 && i%3 == 0:
			fmt.Println(i, "fizz-buzz")
		case i%3 == 0:
			fmt.Println(i, "fizz")
		case i%5 == 0:
			fmt.Println(i, "buzz")
		default:
			fmt.Println(i, "boring")
		case i >= 10:
			break loop
		}
	}
}
