package concurrencypatterns

import (
	"fmt"
	"time"
)

// DemoUsingCancelFuncToStopBackgroundGenerators ...
func DemoUsingCancelFuncToStopBackgroundGenerators() {

	results, cancel := queueExpensiveAndSlowCalculations()
	for r := range results {
		fmt.Println(r)
		if r.value > 5 {
			cancel()
			break
		}
	}
}

type expensiveResult struct {
	value float64
}

func queueExpensiveAndSlowCalculations() (<-chan expensiveResult, func()) {
	ch := make(chan expensiveResult)
	done := make(chan struct{})
	cancel := func() {
		close(done)
	}

	go func() {
		defer close(ch)
		start := 0

		for {
			start = start + 1
			select {
			case <-done:
				return
			case ch <- expensiveResult{expensiveCalculation(start)}:
			}
		}
	}()
	return ch, cancel
}

func expensiveCalculation(i int) float64 {
	time.Sleep(1 * time.Second)
	return float64(i) + 0.1
}

/*
running this code produces
{1.1}
{2.1}
{3.1}
{4.1}
{5.1}
*/
