package arrs

import (
	"fmt"
	"sort"

	"github.com/goblinfactory/greeting/pkg/clone"
)

// TestSlicesArePassedByCopy yknow ...
func TestSlicesArePassedByCopy() {
	nums := []int{10, 5, 15, 20}
	printSortedNumsWithoutSideEffect(nums)
	fmt.Println("finally", nums)
	fmt.Println("---")
	printSortedNumsWithSideEffect(nums)
	fmt.Println("finally", nums)

}

func printSortedNumsWithSideEffect(nums []int) {
	fmt.Println("before", nums)
	// while the slice was copied,
	// modifying the contents of the
	// slice modifies what was
	// "protected" by the "pass by value"
	// default of Go.
	sort.Ints(nums)
	fmt.Println("after", nums)
}

func printSortedNumsWithoutSideEffect(nums []int) {
	fmt.Println("before", nums)
	sort.Ints(clone.Int(nums))
	fmt.Println("after", nums)
}
