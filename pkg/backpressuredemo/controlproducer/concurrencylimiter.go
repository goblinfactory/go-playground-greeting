package controlproducer

import "time"

// DemoConcurrencyLimiter creates an instance of a fake database, passing in a limiter to limit concurrency to 2 simultaneous requests
// then shows three goroutines trying to access the database, and shows visually how one of the goroutines blocked while 2 are allowed access.
func DemoConcurrencyLimiter() {
	// db := NewFakeDatabase(2)
	// col1, col2, col3 =
}

// func LimitMe() int {

// }

// Limiter limits concurrency
type Limiter struct {
	maxConnections int
}

// ------ fake database

// FakeDatabase is a fake database that uses a limited to throttle concurrency limit to max (n) connections at a time.
type FakeDatabase struct {
	limiter Limiter
}

// FakeCustomer ...
type FakeCustomer struct {
	ID   string
	Name string
}

// NewFakeDatabase returns a new fake database
func NewFakeDatabase(maxConnections int) FakeDatabase {
	return FakeDatabase{Limiter{maxConnections}}
}

// AddCustomer ...
func AddCustomer() (bool, error) {
	time.Sleep(500)
	return true, nil
}

// GetCustomer ..
func GetCustomer(id string) (FakeCustomer, error) {
	time.Sleep(500)
	return FakeCustomer{"F001", "Fred Flintstone "}, nil
}
