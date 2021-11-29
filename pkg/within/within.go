package within

import (
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

// Tolerance returns whether two floats are within an acceptable tolerance
func Tolerance(a float64, b float64, tolerance float64) bool {
	fc := cmpopts.EquateApprox(tolerance, 0)
	return cmp.Equal(a, b, fc)
}
