package errorhandling

import (
	"errors"
	"fmt"
	"log"

	"github.com/goblinfactory/greeting/pkg/erroraddress"
)

// DemoUsingErrorsAsToCheckIfAnErrorContainsAnyErrorOfSpecificType ...
func DemoUsingErrorsAsToCheckIfAnErrorContainsAnyErrorOfSpecificType() {

	e1 := someFunctionReturningWrappedErrorsWithoutAddressError()
	// check e1 should return no errors "as" AddressError
	var ae1 erroraddress.AddressError
	if errors.As(e1, &ae1) {
		log.Fatal("should not get here, e1 does not contain an address error.")
	} else {
		fmt.Println("correct; no address errors found.")
	}

	e2 := returnADeeplyNestedErrorWithAnAddressErrorBuriedSomewhere()

	var ae2 erroraddress.AddressError
	if errors.As(e2, &ae2) {
		fmt.Println("correct; An address error was found.")
	} else {
		log.Fatal("should not get here, e2 does actually contain an address error!")
	}

}

func someFunctionReturningWrappedErrorsWithoutAddressError() error {
	e1 := errors.New("file not found")
	e2 := fmt.Errorf("io error: %w", e1)
	return e2
}

func returnADeeplyNestedErrorWithAnAddressErrorBuriedSomewhere() error {
	e1 := errors.New("file not found")
	e2 := erroraddress.NewAddressError(
		erroraddress.PostCode,
		"could not find address file",
		e1,
	)
	e3 := fmt.Errorf("invalid driver :%w", e2)
	return e3
}
