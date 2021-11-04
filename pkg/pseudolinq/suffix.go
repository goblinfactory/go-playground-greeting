package pseudolinq

// Suffix appends a suffix to each of the strings
func (l Strings) Suffix(suffix string) []string {
	arr := []string{}
	for _, r := range l.lines {
		arr = append(arr, r+suffix)
	}
	return arr
}
