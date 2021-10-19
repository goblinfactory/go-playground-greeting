package quotestream

import (
	"context"
	"fmt"
	"sync"

	"github.com/goblinfactory/greeting/pkg/bloggy/breaker"
	"github.com/goblinfactory/greeting/pkg/bloggy/quoteapi"
	"github.com/goblinfactory/greeting/pkg/rnd"
)

// StartQuoteService starts a quote service that gets quotes 10 times a second, returns a channel with the quotes
func StartQuoteService(ctx context.Context, wg *sync.WaitGroup, msPause int) chan string {

	wg.Add(1)
	// count request per 1/10th of a second

	ch := make(chan string)
	api := quoteapi.NewQuoteAPI(6.0)
	getQuote := breaker.NewBreaker(api.GetQuote, 3, 100)

	fmt.Println("starting")
	go func() {
		defer func() {
			fmt.Println("closing")
			close(ch)
			wg.Done()
		}()
		for {
			select {
			case <-ctx.Done():
				ch <- "done."
				return
			default:
				rnd.SleepMs(msPause)
				quote, err := getQuote(ctx)
				if err != nil {
					ch <- err.Error()
				} else {
					ch <- quote
				}
			}
		}
	}()
	return ch
}
