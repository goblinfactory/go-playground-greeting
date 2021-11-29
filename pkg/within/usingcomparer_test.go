package within

import (
	"math"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/assert"
)

func TestCheckingIfFloatIsCloseEnoughDoingitManually(t *testing.T) {

	// float comparer
	fc := cmp.Comparer(func(x, y float64) bool {
		delta := math.Abs(x - y)
		mean := math.Abs(x+y) / 2.0
		return delta/mean < 0.0001
	})

	f1 := 1.00001
	f2 := 1.00002

	assert.True(t, cmp.Equal(f1, f2, fc))

	f1 = 1.1
	f2 = 1.2

	assert.False(t, cmp.Equal(f1, f2, fc))
}

func TestCheckingIfFloatIsCloseEnoughUsingEquateApproxWithFraction(t *testing.T) {

	// float comparer comparing with an absolute difference of 0.00001
	fc := cmpopts.EquateApprox(0.00001, 0)

	f1 := 1.00001
	f2 := 1.00002

	assert.True(t, cmp.Equal(f1, f2, fc))

	f1 = 1.1
	f2 = 1.2

	assert.False(t, cmp.Equal(f1, f2, fc))
}

func TestCheckingIfFloatIsCloseEnoughUsingEquateApproxWithMargin(t *testing.T) {

	// float comparer, using a 1% margin
	fc := cmpopts.EquateApprox(0, 0.01)

	f1 := 1.01
	f2 := 1.02

	assert.False(t, cmp.Equal(f1, f2, fc))

	f1 = 1.01
	f2 = 1.015

	assert.True(t, cmp.Equal(f1, f2, fc))
}
