package di

import (
	"testing"
)

type Adder interface {
	Add(lhs int, rhs int) int
	Subtract(lhs int, rhs int) int
}

func Test_thatYouCanPassANakeFunctionToAMethod(t *testing.T) {
	r := PassMeNakedFunc(AddNaked, 1, 2)
	if r != 3 {
		t.Error("expected 3")
	}
}

func Test_thatYouCanPassInstanceThatImplementsTheInterfaceToAMethodExpectingInterface(t *testing.T) {
	adder := AdderAdapter{}
	r := PassMeAdder(adder, 1, 2)
	if r != 3 {
		t.Error("expected 3")
	}
}

func PassMeNakedFunc(adder func(int, int) int, a int, b int) int {
	r := adder(a, b)
	return r
}

func PassMeAdder(adder Adder, a int, b int) int {
	r := adder.Add(a, b)
	return r
}

type AdderAdapter struct {
}

func AddNaked(lhs int, rhs int) int {
	return lhs + rhs
}
func (ad AdderAdapter) Add(lhs int, rhs int) int {
	return lhs + rhs
}

func (ad AdderAdapter) Subtract(lhs int, rhs int) int {
	return lhs + rhs
}
