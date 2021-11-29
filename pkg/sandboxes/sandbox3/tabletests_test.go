package sandbox3

import (
	"errors"
	"testing"
)

func TestUsingTableTests(t *testing.T) {

	data := []struct {
		name string
		age  int
	}{
		{"fred", 21},
		{"mark", 22},
		//{"grant", -1},
	}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			_, err := doSomethingwithData(d)
			if err != nil {
				t.Error(err)
			}
		})
	}
}

// uncomment line above to cause one of the table tests to fail, and see the output in the command line go test ./...

func doSomethingwithData(data struct {
	name string
	age  int
}) (bool, error) {
	if data.age < 0 {
		return false, errors.New("invalid age")
	}
	return true, nil
}
