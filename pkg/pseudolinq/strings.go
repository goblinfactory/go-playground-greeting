package pseudolinq

// Strings provides psuedo-linq like projections over strings, they're not short circuitable
// e.g. AllEvenNumbers.Take(10) would hang indefinitely with current implementation.
type Strings struct {
	lines []string
}

// New ..
func New(strings []string) Strings {
	return Strings{strings}
}
