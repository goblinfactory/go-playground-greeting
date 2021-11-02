package checkablewaitgroup

// My Naive implementation; pretty certain I may need some sync locks here. But will use for now until I learn more.

import "sync"

// WaitGroup is the waitgroup state
type WaitGroup struct {
	done bool
	wg   sync.WaitGroup
}

// New returns a waitgroup that can be used to check whether any goRoutine has signalled done yet.
// This is useful for scenarios where you have multiple functions racing against each other
// and only the fastest value is used. This avoids having to use an additional channel to signal cancel-abort
// allowing for goroutines to properly do any cleanup.
func New() *WaitGroup {
	return &WaitGroup{}
}

// Add adds delta, which may be negative, to the WaitGroup counter.
// If the counter becomes zero, all goroutines blocked on Wait are released.
// If the counter goes negative, Add panics.
//
// Note that calls with a positive delta that occur when the counter is zero
// must happen before a Wait. Calls with a negative delta, or calls with a
// positive delta that start when the counter is greater than zero, may happen
// at any time.
// Typically this means the calls to Add should execute before the statement
// creating the goroutine or other event to be waited for.
// If a WaitGroup is reused to wait for several independent sets of events,
// new Add calls must happen after all previous Wait calls have returned.
// See the WaitGroup example.
func (wg *WaitGroup) Add(delta int) {
	wg.wg.Add(delta)
}

// Done decrements the WaitGroup counter by one.
func (wg *WaitGroup) Done() {
	wg.done = true
	wg.wg.Done()
}

// Wait blocks until the WaitGroup counter is zero.
func (wg *WaitGroup) Wait() {
	wg.wg.Wait()
}

// IsDone returns whether any goroutine has signalled done yet.
func (wg *WaitGroup) IsDone() bool {
	return wg.done
}
