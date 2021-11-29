package within_test

import (
	"testing"

	"github.com/goblinfactory/greeting/pkg/within"
	"github.com/stretchr/testify/assert"
)

func TestWithinTolerance(t *testing.T) {

	f1 := 1.00001
	f2 := 1.00002

	assert.True(t, within.Tolerance(f1, f2, 0.00001))

	f1 = 1.1
	f2 = 1.2

	assert.False(t, within.Tolerance(f1, f2, 0.00001))
}
