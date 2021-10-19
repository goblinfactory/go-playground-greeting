package learninggo

import (
	"encoding/json"
	"fmt"
)

type person struct {
	name   string
	age    int
	parent *person
}

// TestJSON does...
func TestJSON() {

	p1 := person{"fred", 10, nil}

	fmt.Println(p1)

	if p1.parent == nil {
		fmt.Println("no parent")
	}
	var xx = struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{}

	json.Unmarshal([]byte(`{
		"name":"Bob", 
		"age":30
	}`), &xx)
	fmt.Println(xx)
}
