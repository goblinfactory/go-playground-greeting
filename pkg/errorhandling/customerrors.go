package errorhandling

import (
	"fmt"
)

// DemoCustomErrors ..
func DemoCustomErrors() {
	loadItemHandleErrors("cat")
	fmt.Println("---")
	loadItemHandleErrors("car")
	fmt.Println("---")
	loadItemHandleErrorsReportExtraErrorDetails("car")
	loadItemHandleErrorsReportExtraErrorDetails("carrot")
}

func loadItemHandleErrors(item string) {
	r, err := LoadCatOrCar(item)
	if err != nil {
		fmt.Println("error loading item: %w", err)

	} else {
		fmt.Printf("loaded:%s\n", r)
	}
}

func loadItemHandleErrorsReportExtraErrorDetails(item string) {
	r, err := LoadCatOrCar(item)
	if err != nil {
		fmt.Println(err)
		le, ok := err.(LoadError)
		if ok {
			switch le.Status {
			case NotACat:
				fmt.Println("Not a cat")
			case IsACar:
				fmt.Println("Is a car")
			default:
				fmt.Println("unknown status", le.Status)
			}
		}
	} else {
		fmt.Printf("loaded:%s\n", r)
	}
}

// LoadStatus ..
type LoadStatus int

const (
	// NotACat ..
	NotACat LoadStatus = iota + 1
	// IsACar ...
	IsACar
	// Unknown ...
	Unknown
)

// LoadError ..
type LoadError struct {
	Status  LoadStatus
	Message string
	err     error
}

func (le LoadError) Error() string {
	return le.Message
}

// LoadCatOrCar ...
func LoadCatOrCar(name string) (string, error) {
	if name == "cat" {
		return "cat has a hat now", nil
	}

	if name == "car" {
		// really hacky error message, normally it would
		return "", LoadError{IsACar, "Error loading " + name, nil}
	}
	return "", LoadError{Unknown, "Error loading  " + name, nil}
}
