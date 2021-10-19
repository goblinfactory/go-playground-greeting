package quoteapi

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/paulbellamy/ratecounter"
)

// QuoteState for a newly created api.
type QuoteState struct {
	counter *ratecounter.RateCounter
	maxRps  float64
}

// NewQuoteAPI returns a test quote api that returns 429 when rate exceeds maxRps.
func NewQuoteAPI(maxRps float64) QuoteState {
	counter := ratecounter.NewRateCounter(1 * time.Second)
	return QuoteState{counter, maxRps}
}

//var cnt = 0

// GetQuote gets a market quote returns either -> 200,0, "Quote: £ 429.32 [17 rps]" or -> 429, {backoffMilliseconds}, "429 (too many requests)"
func (state QuoteState) GetQuote(context context.Context) (string, error) {
	// not sure this type of access is thread safe since we're not accessing it in a select
	// if the GetQuote is being accessed in a select, then I guess it IS threadsafe?
	state.counter.Incr(1)
	// cnt++
	// if cnt > 18 {
	// 	return "", errors.New("service stopped")
	// }
	rps := float64(state.counter.Rate())
	if rps > state.maxRps {
		return "", fmt.Errorf("429 (too many requests) %4.1f rps > %4.1f rps", rps, state.maxRps)
	}
	return fmt.Sprintf("Quote: £%7.2f", rand.Float64()*1000), nil
}
