package marshalling

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
	"time"
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

// func TestUnMarshalling(t *testing.T) {
// 	var d1 time.Time
// 	_ := time.Parse("2018-05-01T13:01:02Z", d1)
// 	horses := []Horse{
// 		{"b1", "buttons", d1, "sl1", []Locations{}},
// 	}
// }
