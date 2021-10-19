package structsandmaps

import "fmt"

// TestStructs yknow ...
func TestStructs() {

	type firstPerson struct {
		name string
		age  int
	}

	type secondPerson struct {
		name string
		age  int
	}

	p1 := firstPerson{"Fred", 10}
	p2 := secondPerson(p1)

	fmt.Println(p2)

}

// TestMaps because...
func TestMaps() {

	abbMonths := map[string]string{

		"jan": "January",
		"feb": "February",
		"mar": "March",
		"apr": "April",
		"may": "May",
		"jun": "June",
		"jul": "July",
		"aug": "August",
		"sep": "September",
		"oct": "October",
		"nov": "November",
		"dec": "December",
	}

	months := []string{
		"jan",
		"feb",
		"mar",
		"apr",
		"may",
		"jun",
		"jul",
		"aug",
		"sep",
		"oct",
		"nov",
		"dec",
	}

	for _, v := range months {
		fmt.Println(v, abbMonths[v])
	}

	q1 := months[0:3]
	q2 := months[3:6]
	q3 := months[6:9]
	q4 := months[9:12]

	fmt.Println("q1", q1)
	fmt.Println("q2", q2)
	fmt.Println("q3", q3)
	fmt.Println("q4", q4)

}
