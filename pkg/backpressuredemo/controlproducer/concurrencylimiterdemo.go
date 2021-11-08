package controlproducer

import (
	"context"
	"fmt"
	"time"

	"github.com/goblinfactory/greeting/pkg/consolespikes"
	"github.com/mum4k/termdash/widgets/text"
)

const reset = string("\033[0m")
const red = string("\033[31m")
const green = string("\033[32m")

// DemoConcurrencyLimiter creates an instance of a fake database, passing in a limiter to limit concurrency to 2 simultaneous requests
// then shows three goroutines trying to access the database, and shows visually how one of the goroutines blocked while 2 are allowed access.
func DemoConcurrencyLimiter() {

	db := NewFakeDatabase(2)
	status, c1, c2, c3, wg, ctx := consolespikes.SplitColumns1234("status", "client-1", "client-2", "client-3")

	status.Write("starting 3 clients\n")

	go readWriteDemo(ctx, c1, db, "client1")
	go readWriteDemo(ctx, c2, db, "client2")
	go readWriteDemo(ctx, c3, db, "client3")
	wg.Wait()
}

func readWriteDemo(ctx context.Context, con *text.Text, db FakeDatabase, name string) {
	k := consolespikes.NewKonsole(con)
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
				time.Sleep(200 * time.Millisecond)
			} else {
				k.Green(cid, "write ok\n")
			}
		}
	}
}
