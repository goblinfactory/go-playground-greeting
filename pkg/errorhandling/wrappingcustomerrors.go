package errorhandling

import (
	"errors"
	"fmt"
)

// TestWrappingCustomErrors calls checkCar that will return a wrapped
func TestWrappingCustomErrors() {
	d := driver{}
	de := checkDriver(d)
	if de != nil {
		printCustomDriverErrors(de)
		return
	}
	fmt.Println("Valid driver.")
}

func printCustomDriverErrors(e error) {
	fmt.Println("DRIVER ERROR")
	fmt.Println("------------")
	fmt.Println("Invalid driver")
	er := e
	i := "   "
	for {
		er = errors.Unwrap(er)
		if er == nil {
			return
		}
		i = i + "   "
		fmt.Printf("%s| %s\n", i, er)
	}
}
