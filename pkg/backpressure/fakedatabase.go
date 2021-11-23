package backpressure

import (
	"github.com/goblinfactory/greeting/pkg/rnd"
)

// FakeDatabase is a fake database that uses a limited to throttle concurrency limit to max (n) connections at a time.
type FakeDatabase struct {
	limiter Limiter
}

// FakeCustomer ...
type FakeCustomer struct {
	ID   string
	Name string
}

// Processor interface for the database. Allows for pre and post processing of requests.
type Processor interface {
	Process()
}

// NewFakeDatabase ..
func NewFakeDatabase(maxConnections int) FakeDatabase {
	return FakeDatabase{NewLimiter(maxConnections)}
}

// AddCustomer ..
func (db FakeDatabase) AddCustomer(correlationID string, customer FakeCustomer) (FakeCustomer, error) {
	err := db.limiter.RunWithMaxConcurrency(func() {
		rnd.SleepMinMaxMs(100, 200)
	})
	if err == nil {
		return FakeCustomer{"F001", "Fred Flintstone "}, nil

	}
	return FakeCustomer{"", ""}, err

}
