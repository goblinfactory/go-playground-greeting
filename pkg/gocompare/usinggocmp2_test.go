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
	assert.True(t, diff != "")
	// you can't compare or assume anything about the returned text, below text is apparently deliberately randomised by the dev
	// team to make sure you dont rely on the se values.
	//"\u00a0\u00a0gocompare.Company{\n\u00a0\u00a0\tName: \"company1\",\n\u00a0\u00a0\tEmployees: []gocompare.Person{\n\u00a0\u00a0\t\t{Name: \"fred\", Address: {Houseno: 1, Street: \"street1\", City: \"city1\"}},\n\u00a0\u00a0\t\t{\n\u00a0\u00a0\t\t\tName: \"fred\",\n\u00a0\u00a0\t\t\tAddress: gocompare.Address{\n\u00a0\u00a0\t\t\t\tHouseno: 2,\n\u00a0\u00a0\t\t\t\tStreet:  \"street2\",\n-\u00a0\t\t\t\tCity:    \"city2\",\n+\u00a0\t\t\t\tCity:    \"city3\",\n\u00a0\u00a0\t\t\t},\n\u00a0\u00a0\t\t},\n\u00a0\u00a0\t},\n\u00a0\u00a0}\n"
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

// cmp can't handle unexported type
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
