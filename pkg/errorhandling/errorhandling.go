package errorhandling

import (
	"errors"
	"fmt"
)

// Demo1 ...
func Demo1() {
	DemoMostCommonPattern(-1, 2)
}

// Demo2 ...
func Demo2() {
	DemoMostCommonPattern(1, 2)
}

// IAmAnError ..
type IAmAnError struct{}

func (e IAmAnError) Error() string {
	return "I am an error"
}

// DemoPassingErrorToPrintln ..
func DemoPassingErrorToPrintln() {
	fmt.Println(IAmAnError{})
}

// DemoMostCommonPattern ...
func DemoMostCommonPattern(a int, b int) {
	r, err := DoSomethingRisky(a, b)
	if err != nil {
		fmt.Println("error", err.Error())
	}
	fmt.Println("1 + 2 =", r)
}

// DoSomethingRisky ...
func DoSomethingRisky(a int, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, errors.New("negative numbers not supported")
	}
	return a + b, nil
}
