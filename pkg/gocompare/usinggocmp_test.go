package gocompare

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/assert"
)

func TestHowToUseGoCmpPackage(t *testing.T) {
	c1 := Company{"company1", []Person{
		{"fred", Address{1, "street1", "city1"}},
		{"fred", Address{2, "street2", "city2"}},
	}}

	// c2.company.employee[1].address.city changed to city3 s

	c2 := Company{"company1", []Person{
		{"fred", Address{1, "street1", "city1"}},
		{"fred", Address{2, "street2", "city3"}},
	}}
	diff := cmp.Diff(c1, c2)
	assert.NotNil(t, diff)

}

// Company ...
type Company struct {
	Name      string
	Employees []Person
}

// Person ...
type Person struct {
	Name    string
	Address Address
}

// Address ...
type Address struct {
	Houseno int
	Street  string
	City    string
}

// cmp can't handle unexported types
func TestComparingWithUnexportedFields(t *testing.T) {
	c1 := Cat{10, "Garfield"}
	c2 := Cat{10, "Garfield"}
	diff := cmp.Diff(c1, c2, cmpopts.IgnoreUnexported(c1, c2))
	assert.Equal(t, diff, "")

	// now compare without ignoring unexpected and it will panic
	assert.Panics(t, func() {
		cmp.Diff(c1, c2)
	})
}

type Cat struct {
	age  int
	Name string
}

// references : https://mariocarrion.com/2021/01/22/go-package-equality-google-go-cmp.html
