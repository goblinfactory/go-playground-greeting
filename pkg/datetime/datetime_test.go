package datetime

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLayout(t *testing.T) {
	d1, err := time.Parse(time.Layout, "03/01 10:00:00PM '21 -0000")
	fmt.Println(d1)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 2021, d1.Year())

	tzName, tzOffset := d1.Zone()
	assert.Equal(t, "GMT", tzName)
	assert.Equal(t, 0, tzOffset)
}

func TestTimeZone(t *testing.T) {
	d1, err := time.Parse(time.Layout, "03/01 10:00:00PM '21 -0700")
	fmt.Println(d1)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 2021, d1.Year())

	tzName, tzOffset := d1.Zone()
	assert.Equal(t, "", tzName)
	assert.Equal(t, -7*60*60, tzOffset)
}
