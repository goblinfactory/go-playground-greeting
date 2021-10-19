package arrs

import (
	"fmt"
	"sort"
)

// TestSorting yknow ...
func TestSorting() {

	a := []int{1, 2, 3, 4}
	b := []int{1, 2, 3, 3}
	c := []int{1, 2, 3, 4}
	d := []int{2, 3, 4, 1}
	fmt.Println("test if arrays are equal and haveSameElements")
	fmt.Println("a == b", equal(a, b))
	fmt.Println("a == c", equal(a, c))
	fmt.Println("a == c (ignore order)", equalIgnoreOrder(a, c))
	fmt.Println("a == d", equal(a, d))
	fmt.Println("a == d (ignore order)", equalIgnoreOrder(a, d))
}

// has the same elements and the same order
func equal(lhs []int, rhs []int) bool {
	for i, o := range lhs {
		if o != rhs[i] {
			return false
		}
	}
	return true
}

// has the same elements ignoring order
func equalIgnoreOrder(lhs []int, rhs []int) bool {

	return equal(
		sortSafe(lhs),
		sortSafe(rhs),
	)
}

// clones the array and then sorts it. This is safe/side effect free.
func sortSafe(items []int) []int {
	n := make([]int, len(items))
	for i := range n {
		n[i] = items[i]
	}
	sort.Ints(n)
	return n
}
