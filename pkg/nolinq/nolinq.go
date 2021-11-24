package nolinq

import "fmt"

type car struct {
	year                int
	owner, model, color string
	id                  int
}

type garage []car

// DemoQueryingObjectsAndCollectionsWithoutLinq tests raw Go instead of Linq.
func DemoQueryingObjectsAndCollectionsWithoutLinq() {

	var cars = []car{
		{1950, "Jay Leno", "buic", "red", 1},
		{1965, "Jay Leno", "chrysler", "blue", 2},
		{2000, "Fred", "bmw", "black", 3},
		{2010, "Dan", "volvo", "red", 4},
	}

	var g = garage(cars)

	fmt.Println("\nVintage\n", g.vintage())
	fmt.Println("\nRed\n", g.color("red"))
	fmt.Println("\nVintage AND red\n", g.vintage().color("red"))
	fmt.Println("\nVintage OR red\n", g.vintage().or(g.color("red")))
}

func (cars garage) color(color string) garage {
	matches := garage([]car{})
	for _, c := range cars {
		if c.color == color {
			matches = append(matches, c)
		}
	}
	return matches
}

func (cars garage) or(rhs garage) garage {
	union := garage([]car{})
	for i := range cars {
		union = append(union, cars[i])
	}
	for _, c := range rhs {
		if !cars.contains(c) {
			union = append(union, c)
		}

	}
	return union
}

func (cars garage) contains(car car) bool {
	// if this was a super large set, we'd use an additional hashset so that
	// we didnt have to loop through entire collection
	// but would need to intercept "add/remove"
	for _, c := range cars {
		if c.id == car.id {
			return true
		}
	}
	return false
}

func (cars garage) vintage() garage {
	matches := garage([]car{})
	for _, c := range cars {
		if c.year <= 1970 {
			matches = append(matches, c)
		}
	}
	return matches
}
