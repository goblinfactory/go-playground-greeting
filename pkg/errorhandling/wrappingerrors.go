package errorhandling

import (
	"errors"
	"fmt"
)

//type

// TestWrappingErrors calls checkCar that will return a wrapped
func TestWrappingErrors() {
	d := driver{}
	de := checkDriver(d)
	if de != nil {
		printDriverErrors(de)
		return
	}

	fmt.Println("Valid driver.")
}

func printDriverErrors(e error) {
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

type driver struct {
}

func checkDriver(d driver) error {
	de := checkDriverAddress(d)
	if de != nil {
		return fmt.Errorf("Invalid driver: %w", de)
	}
	return nil
}

func checkDriverAddress(d driver) error {
	ce := checkDriverAddressCountry(d)
	if ce != nil {
		return fmt.Errorf("Invalid address: %w", ce)
	}
	return nil
}

func checkDriverAddressCountry(d driver) error {
	return errors.New("driver not in the UK")
}
