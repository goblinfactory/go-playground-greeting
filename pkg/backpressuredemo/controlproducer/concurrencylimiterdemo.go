package controlproducer

import (
	"context"
	"fmt"
	"time"

	"github.com/goblinfactory/greeting/pkg/consolespikes"
)

// DemoConcurrencyLimiter creates an instance of a fake database, passing in a limiter to limit concurrency to 2 simultaneous requests
// then shows three goroutines trying to access the database, and shows visually how one of the goroutines blocked while 2 are allowed access.
// also visually shows how different attempts to "load balance" or manage concurrency can be a lot trickier to achieve expected outcomes.
// in this demo, we run a concurrency limited that ensure that no more than 2 concurrent requests simultaenously, and see visually that
// while we are assured no more than 2 requests run at a time, there are often many times when we dont get an even spread of the work
// and can even have periods where we have multiple requests from all three clients returning 429 at the same time in a row. This is
// counter intuitive. A really nice example of interesting distributed system choreography. See screenshot.
func DemoConcurrencyLimiter() {

	db := NewFakeDatabase(2)
	status, c1, c2, c3, wg, ctx, _, _ := consolespikes.SplitColumns1234("status", "client-1", "client-2", "client-3")

	status.Write("starting 3 clients, press 'q' to quit.\n")

	go readWriteDemo(ctx, c1, db, "client1")
	go readWriteDemo(ctx, c2, db, "client2")
	go readWriteDemo(ctx, c3, db, "client3")
	wg.Wait()
}

func readWriteDemo(ctx context.Context, k consolespikes.Konsole, db FakeDatabase, name string) {
	i := 0
	k.WriteLine("starting ", name)
	for {
		i++
		select {
		case <-ctx.Done():
			k.Green("--finished--\n")
			return
		default:
			cid := fmt.Sprintf("%s:%d", name, i)
			_, err := db.AddCustomer(cid, FakeCustomer{})
			if err != nil {
				k.Red(err.Error(), "\n")
				time.Sleep(20 * time.Millisecond)
			} else {
				k.Green(cid, "write ok\n")
			}
		}
	}
}
