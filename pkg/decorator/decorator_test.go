package decorator

import (
	"fmt"
	"testing"
)

type DecoratedAdd func(a int, b int) int

// this is such a quick way to implement a decorator
func (f DecoratedAdd) Add(a int, b int) int {
	fmt.Println("before")
	defer fmt.Println("after")
	r := f(a, b)
	fmt.Println("result", r)
	return r
}

func Test_howToUseDecorator(t *testing.T) {
	ad := DecoratedAdd(AddNums)
	r := ad.Add(1, 3)
	if r != 4 {
		t.Error("expected 4")
	}

}

func AddNums(a int, b int) int {
	return a + b
}

// here's what the same decorator pattern looks like in C#
// https://dotnetfiddle.net/9bRx4e
