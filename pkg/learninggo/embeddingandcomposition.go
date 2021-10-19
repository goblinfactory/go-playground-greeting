package learninggo

import "fmt"

// Employee ...
type Employee struct {
	Name string
	ID   string
}

// Box1 ...
type Box1 struct {
	name string
}

// Box2 ...
type Box2 struct {
	Box1
	desc string
}

// Manager ...
type Manager struct {
	Employee
	Reports []Employee
}

// TestEmbeddingAndComposition ...
func TestEmbeddingAndComposition() {

	b1 := Box2{Box1{"fred"}, "desc1"}

	fmt.Println(b1)
	// m := Manager{
	// 	Employee{"Grant", "G1"},
	// 	[]Employee{
	// 		{"Kevin", "K1"},
	// 	},
	// }

	// fmt.Println(m)
}
