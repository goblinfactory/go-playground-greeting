package bloggy

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type toggleBold struct{}

// RunTestContext runs a simple test ot test using context to cancel a background operation
func RunTestContext() {

	ct := context.Background()
	ctx, done := context.WithTimeout(ct, 3000*time.Millisecond)

	fmt.Println(ct, done)

	var wg sync.WaitGroup

	nums := make(chan int, 100)
	tb := make(chan toggleBold)

	streamNums(ctx, &wg, nums, 1)
	streamNums(ctx, &wg, nums, 2)
	printNums(ctx, &wg, nums, tb)

	// print numbers for 1 second in default (non bold)
	time.Sleep(1000 * time.Millisecond)

	// toggle bold
	tb <- toggleBold{}

	// print numbers for 1 second
	time.Sleep(1000 * time.Millisecond)

	// toggle bold
	tb <- toggleBold{}

	// lets wait for a further 2 seconds which should cause the context to timeout (max at 2 seconds and we'd be at 3 so definitely overlaps past deadline)
	time.Sleep(2000 * time.Millisecond)

	// if not already signalled stop
	// signal we're done
	// if already signalled this line has not effect.
	done()

	close(nums)
	wg.Wait()
	fmt.Println("FINISHED")
}

func printNums(ct context.Context, wg *sync.WaitGroup, nums chan int, toggleBold chan toggleBold) {
	wg.Add(1)
	go _printNums(ct, wg, nums, toggleBold)
}

func _printNums(ct context.Context, wg *sync.WaitGroup, nums chan int, toggleBold chan toggleBold) {
	defer wg.Done()
	bold := false
	for n := range nums {
		select {
		case <-ct.Done():
			return
		case <-toggleBold:
			bold = !bold
		default:
			{
				if bold {
					fmt.Printf("** num:%d **\n", n)
				} else {
					fmt.Printf("num:%d\n", n)
				}

			}
		}

	}
	fmt.Println("nums channel closed, printing has stopped.")
}

func sleep(min int, max int) {
	ms := time.Duration(rand.Intn(min)+max) * time.Millisecond
	time.Sleep(ms)
}

func streamNums(ctx context.Context, wg *sync.WaitGroup, nums chan int, start int) {
	wg.Add(1)
	go _streamNums(ctx, wg, nums, start)
}

// may not need the wait group?
func _streamNums(ctx context.Context, wg *sync.WaitGroup, nums chan int, start int) {
	i := start

	defer func() {
		fmt.Printf("service stopped [%d]\n", i)
		wg.Done()
	}()

	for {
		select {
		case <-ctx.Done():
			return
		default:
			nums <- i
			i = i + 2
			sleep(100, 500)
		}

	}
}
