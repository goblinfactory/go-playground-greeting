package breaker

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Circuit function that the circuit breaker will protect.
type Circuit func(context context.Context) (string, error)

// NewBreaker returns a circuit breaker that will allow you  to fail n times in a row. if that happens, then it trips the circuit in increasing lengths of time (no limit)
// any succesful request after tripped resets (untrips) the circuit.
func NewBreaker(circuit Circuit, failureThreshold uint, msBreakFor uint) Circuit {

	var consecutiveFailures int = 0
	var lastAttempt = time.Now()
	var m sync.RWMutex

	return func(ctx context.Context) (string, error) {

		m.RLock() // establish a read locks

		d := consecutiveFailures - int(failureThreshold)

		if d >= 0 { // backoff will only start to apply after d becomes positive and will then continue indefinitely.
			tripDuration := time.Millisecond * 20 << d
			shouldRetryAt := lastAttempt.Add(tripDuration)

			if !time.Now().After(shouldRetryAt) {
				m.RUnlock()
				return "", fmt.Errorf("service tripped Pausing %v", tripDuration)
			}
		}

		m.RUnlock() // Release read lock

		response, err := circuit(ctx)

		m.Lock() // lock around shared resource
		defer m.Unlock()

		lastAttempt = time.Now()
		if err != nil {
			consecutiveFailures++
			return response, err
		}

		consecutiveFailures = 0
		return response, nil
	}
}

// todo; it would be nice if I can show a simple demo where this breaks if we dont include the RW mutex locks
// and this is called in 10 parallel threads.
