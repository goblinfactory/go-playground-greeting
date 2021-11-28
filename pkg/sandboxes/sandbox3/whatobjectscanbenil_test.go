package sandbox3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProveMapPropBeSetToNil(t *testing.T) {
	m := mappy{}
	var nc map[string]int // this is a nil "empty" map that can't have items added.
	assert.Equal(t, m.topncities, nc)

	// setting to nil doesnt change anything, prove nil == var nc map[string]int
	m.topncities = nil
	assert.Equal(t, m.topncities, nc)

	// while map is set to nil accessing it's len will not throw exception!
	assert.Equal(t, 0, len(m.topncities))

	// code below panics with error <-- assignment to entry in nil map
	assert.Panics(t, func() {
		m.topncities["a"] = 1
	})
}

type mappy struct {
	topncities map[string]int
}

func TestHowToCreateDefaultstructWithMapPropNotNil(t *testing.T) {
	// in this example I actually want a map I can add to
	m := mappy{make(map[string]int)}
	m.topncities["a"] = 1
}

type manager struct {
	empno string
	employee
}

type employee struct {
	person
}

type person struct {
	name          string
	add           address
	prevAddresses map[address]int
}
