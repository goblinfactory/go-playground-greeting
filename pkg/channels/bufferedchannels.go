package channels

import (
	"fmt"
	"sync"
	"time"

	"github.com/goblinfactory/greeting/pkg/rnd"
)

// DemoBufferedChannels shows an example of using a buffered channel to make concurrent calls to a micro-service
func DemoBufferedChannels() {

	products := []string{
		"kbd-01",
		"kbd-02",
		"mon-23",
		"mon-13",
		"ssd-16",
		"cas-66",
		"cpu-91",
		"hea-72",
	}

	// The calls to the service will be concurrent so getting all the weights for all the products shouldnt take more than
	// the time for the longest call to complete fetching all the values;  average max is 2000 ms at the bottom.

	for p := range lookupProductWeights(products) {
		rnd.SleepMinMaxMs(200, 600)

		fmt.Printf("%s - %00.2fkg \t: %4dms\n", p.id, p.weight, p.requestDuration)
	}
}

type productWeight struct {
	id              string
	weight          float64
	requestDuration int64
}

// lookupProductWeights (using a buffered channel) this is somewhat similar to C#'s new IAsyncEnumerable
// example below excludes important code for requesting cancellation. See cancelfunction for simple pattern that can easily be added.
func lookupProductWeights(IDs []string) <-chan productWeight {
	var wg sync.WaitGroup

	cnt := len(IDs)
	wg.Add(cnt)
	ch := make(chan productWeight, cnt)
	start := time.Now()

	for _, id := range IDs {
		go func(id string) {
			defer wg.Done()
			weight := lookupProductWeight(id)
			ms := time.Since(start).Milliseconds()
			ch <- productWeight{id, weight, ms}
		}(id)
	}
	// no idea if this is clever or not, but it seems to get the job done?
	// wait for all the functions to complete then close the channel.
	// I did it this way to follow best practice of producer is responsible
	// for closing the channel.
	go func() {
		wg.Wait()
		close(ch)
		fmt.Println("all done")
	}()
	return ch
}

// fake "service" that looks up a product weight from product code
func lookupProductWeight(productID string) float64 {
	rnd.SleepMinMaxMs(40, 2000)
	return rnd.RandFloat(20, 2)
}

/*
spike demonstrates consuming the feed before all the results are ready
and the services then completing before the whole feed is processed.

Running above code produces

kbd-01 - 1.94kg 	:   81ms
hea-72 - 10.30kg 	:  203ms
mon-13 - 16.27kg 	:  220ms
cpu-91 - 4.29kg 	:  364ms
all done
kbd-02 - 6.36kg 	:  848ms
ssd-16 - 5.66kg 	: 1011ms
mon-23 - 5.86kg 	: 1068ms
cas-66 - 4.06kg 	: 1419ms

*/
