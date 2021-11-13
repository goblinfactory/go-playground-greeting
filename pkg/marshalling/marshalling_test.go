package marshalling

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// DemoMarhsalling ...
func DemoMarhsalling() {

}

// Horse ...
type Horse struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Born      time.Time  `json:"born"`
	ParentID  string     `json:"parent-id"`
	Locations []Location `json:"locations"`
}

// Location for Horse
type Location struct {
	ID      int    `json:"id"`
	Paddock string `json:"paddock"`
}

func TestMarshalling(t *testing.T) {
	f, _ := os.Open("testdata/test.json")
	defer f.Close()
	b, _ := ioutil.ReadAll(f)
	var h Horse
	_ = json.Unmarshal(b, &h)
	fmt.Println(h)
}

// see https://chmod-calculator.com/
const OwnerDefault = 0754 // rwxr-xr-- (octal)

func TestMarshallingAndUnmarshalling(t *testing.T) {

	const file = "testdata/test_readwrite.json"

	// Marshalling (serialising) to JSON
	// ---------------------------------
	d1, err := time.Parse(time.RFC3339, "2000-05-01T13:01:02Z")
	if err != nil {
		t.Error(err)
	}

	horse := Horse{"GAR01", "Gary", d1, "but01",
		[]Location{
			{10, "PortElizabeth"},
			{20, "Cambridge"},
		},
	}
	b, err := json.Marshal(horse)
	if err != nil {
		t.Error(err)
	}

	err = ioutil.WriteFile(file, b, fs.FileMode(OwnerDefault))
	if err != nil {
		t.Error(err)
	}

	// UnMarshalling (deserialising) from JSON
	// ---------------------------------------
	f, err := os.Open(file)
	defer f.Close()
	if err != nil {
		t.Error(err)
	}

	hb, err := ioutil.ReadAll(f)
	if err != nil {
		t.Error(err)
	}

	var horse2 Horse
	err = json.Unmarshal(hb, &horse2)

	assert.Equal(t, horse, horse2)

	// need to check that this is doing a deep equal.
	horse.Locations[1].Paddock = "new paddock"
	assert.NotEqual(t, horse, horse2)
}
