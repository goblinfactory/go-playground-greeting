package clone

// Int returns a copy of the []int (useful for writing code in a more functional side effect free style)
func Int(nums []int) []int {
	dest := make([]int, len(nums))
	copy(dest, nums)
	return dest
}
