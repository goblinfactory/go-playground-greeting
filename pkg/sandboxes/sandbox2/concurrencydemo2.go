package sandbox2

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/goblinfactory/greeting/pkg/money"
)

var apiCallTime = 20 * time.Millisecond // 20 x 3 = 60ms, this should cause the final stage of the processing pipeline to exceed the context timeout. run the code, see this fail, then set the value to 10, to see it all pass.

// DemoGatherAndProcess shows a combination of concurrency techniques working together to satisfy the following requirement
// 1. call 2 webservices
// 		shipping := get_shipping data ("KEY01")
// 		price := get_price ("KEY01")
// 		quote := generate_quote (shipping, price)
// 3. the entire process must take less than 50ms or an error is returned.
// GatherAndProcess ...
func DemoGatherAndProcess() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	rfq := RFQ{"KEY01", "cb11bb"}
	q, err := gatherAndProcess(ctx, rfq)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Woot, quote received in time:")
	fmt.Println("product:", rfq.PartNumber, q.Product.Short())
	fmt.Println("shipping to:", rfq.DeliveryPostCode, q.Shipping.Short())
	fmt.Println("total", q.Total.Short())
}

// PartNumber ...
type PartNumber string

// Quote ...
type Quote struct {
	Shipping money.GBP
	Product  money.GBP
	Total    money.GBP
}

// RFQ ...
type RFQ struct {
	PartNumber       string
	DeliveryPostCode string
}

type costs struct {
	shipping getShippingResult
	product  getUnitCostResult
}

type webapi struct {
	ctx context.Context
}

type getShippingResult struct{ amount money.GBP }
type getUnitCostResult struct{ amount money.GBP }
type getFinalQuoteResult struct{ quote Quote }

var maxRequestTime = 50 * time.Millisecond

func gatherAndProcess(ctx context.Context, rfq RFQ) (Quote, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Millisecond)
	defer cancel()

	chShip := make(chan getShippingResult)
	chCost := make(chan getUnitCostResult)
	chQuote := make(chan getFinalQuoteResult)
	chError := make(chan error, 2)

	api := webapi{ctx}

	// fail fast on the first error

	// lookup shipping costs
	// ---------------------
	go func() {
		shipping, err := api.getShippingCost(rfq.DeliveryPostCode)
		if err != nil {
			chError <- fmt.Errorf("getShippingCost: %w", err)
			return
		}
		chShip <- shipping
	}()

	// lookup product unit cost
	// ------------------------
	go func() {
		unitcost, err := api.getUnitCost(rfq.PartNumber)
		if err != nil {
			chError <- fmt.Errorf("getUnitCost: %w", err)
			return
		}
		chCost <- unitcost
	}()

	// wait for both shipping and unit cost to finish
	// and ensure that if ctx is cancelled or timeout is exceeded
	// it must timeout with error
	// --------------------------
	var qcosts costs
	for i := 0; i < 2; i++ {
		select {
		case <-ctx.Done():
			return Quote{}, fmt.Errorf("timeout exceeded %v", maxRequestTime)
		case err := <-chError:
			return Quote{}, err
		case shipping := <-chShip:
			qcosts.shipping = shipping
			continue
		case unitcost := <-chCost:
			qcosts.product = unitcost
			continue
		}
	}

	fmt.Println("shipping and unit cost returned")
	fmt.Println("product cost:", qcosts.product.amount.Short())
	fmt.Println("shipping cost:", qcosts.shipping.amount.Short())

	// last step is to get the quote
	// and ensure that if ctx is cancelled or timeout is exceeded
	// it must timeout with error
	// --------------------------
	go func() {
		quote, err := api.getFinalQuote(qcosts)
		if err != nil {
			chError <- fmt.Errorf("getFinalQuote: %w", err)
			return
		}
		chQuote <- quote
	}()

	select {
	case <-ctx.Done():
		return Quote{}, fmt.Errorf("timeout exceeded %v", maxRequestTime)
	case err := <-chError:
		return Quote{}, err
	case q := <-chQuote:
		return q.quote, nil
	}
}

// -------- fake web API client -----------
// api client will not take channels as parameters
// because it typically would be an external package and channels are not appropriate concurrency idiom for an api.
// need to check what happens when context times out here during the time.Sleep?

func (api *webapi) getShippingCost(postcode string) (getShippingResult, error) {
	fmt.Println("calling api/shipping")
	// fake call out to web API
	time.Sleep(apiCallTime)
	return getShippingResult{money.NewGBP(10)}, nil
}

func (api *webapi) getUnitCost(partNumber string) (getUnitCostResult, error) {
	fmt.Println("calling api/inventory")
	// fake call out to web API
	time.Sleep(apiCallTime)
	return getUnitCostResult{money.NewGBP(20)}, nil
}

func (api *webapi) getFinalQuote(costs costs) (getFinalQuoteResult, error) {
	// fake call out to web API
	fmt.Println("calling api/quotes")
	time.Sleep(apiCallTime)
	// add 30% and 40% markup respectively
	ns := costs.shipping.amount.Multiply(1.3)
	np := costs.product.amount.Multiply(1.4)
	total := ns.Add(np)
	quote := Quote{ns, np, total}

	return getFinalQuoteResult{quote}, nil
}
