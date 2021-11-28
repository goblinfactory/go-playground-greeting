package sandbox3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type MyInt int
type YourInt MyInt

func TestAddingCustomTypesTogether(t *testing.T) {
	a1 := MyInt(1)
	a2 := MyInt(2)

	// b1 := YourInt(4)
	// b2 := YourInt(8)

	// prove you can do maths on custom types just like you can with the original
	assert.Equal(t, a1+a2, MyInt(3))

	// but you can't mix int's and myInts
	assert.NotEqual(t, a1+a2, 3)

	// to mix, convert back to int
	assert.Equal(t, int(a1+a2), 3)

	// you can't mix
	assert.NotEqual(t, MyInt(1), YourInt(1))

	// so, how is this valid? time.Hour + 30 * time.Minute
	// assume it's int constants?

	// lets test that
	const i10 MyInt = 10
	const i20 MyInt = 20

	assert.Equal(t, MyInt(10), i10)
	assert.Equal(t, MyInt(20), i20)

}
