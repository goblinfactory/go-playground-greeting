package controlproducer

import (
	"log"
	"time"
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
	log.Println("adding customer", correlationID)
	err := db.limiter.RunWithMaxConcurrency(func() {
		time.Sleep(500 * time.Millisecond)
	})
	if err == nil {
		log.Println("sucess", correlationID)
		return FakeCustomer{"F001", "Fred Flintstone "}, nil

	}

	log.Println("error", correlationID, err.Error())
	time.Sleep(500 * time.Millisecond)
	return FakeCustomer{"", ""}, err

}
