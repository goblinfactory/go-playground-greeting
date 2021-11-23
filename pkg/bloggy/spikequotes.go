package bloggy

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/goblinfactory/greeting/pkg/bloggy/quotestream"
)

// DemoCallingAPIsWithCircuitBreaker starts a test quotestream that uses a circuit breaker
func DemoCallingAPIsWithCircuitBreaker() {
	fmt.Printf("\nCIRCUIT BREAKER\n")
	fmt.Printf("---------------\n\n")

	var wg sync.WaitGroup
	bg := context.Background()
	ctx, cancel := context.WithCancel(bg)

	ch := quotestream.StartQuoteService(ctx, &wg, 100)

	// start goroutine to print the quotes as they arrive sd
	go func(quotes chan string) {
		for q := range quotes {
			fmt.Println(q)
		}
	}(ch)

	// now wait for service to print some quotes
	time.Sleep(6 * time.Second)

	cancel()
	wg.Wait()
	// wait for all to finish
}
