package mylinq

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuffix(t *testing.T) {
	lines := []string{"one", "two", "three"}
	actual := New(lines).Suffix(".txt")
	expected := []string{"one.txt", "two.txt", "three.txt"}
	assert.Equal(t, expected, actual)
}
