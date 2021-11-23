package backpressure

import "errors"

// Limiter limits concurrency
type Limiter struct {
	ch chan struct{}
}

// NewLimiter returns a new concurrency limiter
func NewLimiter(maxConnections int) Limiter {
	ch := make(chan struct{}, maxConnections)
	for i := 0; i < maxConnections; i++ {
		ch <- struct{}{}
	}
	return Limiter{ch}
}

// RunWithMaxConcurrency ...
func (l Limiter) RunWithMaxConcurrency(f func()) error {
	select {
	case <-l.ch:
		f()
		l.ch <- struct{}{}
		return nil
	default:
		return errors.New(("429, Too many requests."))
	}
}
