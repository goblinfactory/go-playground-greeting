package muxyidiomatic

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// EvenOddWithWaitGroup foo
func EvenOddWithWaitGroup() {

	type stop struct{}
	type toggleBold struct{}

	sleep := func(min int, max int) {
		ms := time.Duration(rand.Intn(min)+max) * time.Millisecond
		time.Sleep(ms)
	}

	// Exit this routine by closing the nums channel.
	_printNums := func(wg *sync.WaitGroup, nums chan int, toggleBold chan toggleBold) {
		defer wg.Done()
		bold := false
		for n := range nums {
			select {
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

	printNums := func(wg *sync.WaitGroup, nums chan int, toggleBold chan toggleBold) {
		wg.Add(1)
		go _printNums(wg, nums, toggleBold)
	}

	_streamNums := func(wg *sync.WaitGroup, nums chan int, start int, stop chan stop) {
		i := start
		defer wg.Done()
		for {
			select {
			case <-stop:
				{
					fmt.Printf("stopping service [%d]\n", i)
					sleep(500, 1000)
					fmt.Printf("service stopped [%d]\n", i)
					return
				}
			default:
				nums <- i
				i = i + 2
				sleep(100, 500)
			}

		}
	}

	streamNums := func(wg *sync.WaitGroup, nums chan int, start int, stop chan stop) {
		wg.Add(1)
		go _streamNums(wg, nums, start, stop)
	}

	var wg sync.WaitGroup

	evens := make(chan stop)
	odds := make(chan stop)
	nums := make(chan int, 100)
	tb := make(chan toggleBold)

	streamNums(&wg, nums, 1, odds)
	streamNums(&wg, nums, 2, evens)
	printNums(&wg, nums, tb)

	// print numbers for 2 second in default (non bold)
	time.Sleep(2000 * time.Millisecond)

	// toggle bold
	tb <- toggleBold{}

	// print numbers for a further 2 seconds
	time.Sleep(2000 * time.Millisecond)

	// toggle bold
	tb <- toggleBold{}

	// print numbers for a further 2 seconds
	time.Sleep(2000 * time.Millisecond)

	fmt.Println("Stopping all")
	close(evens)
	close(odds)
	close(nums)
	wg.Wait()
	fmt.Println("FINISHED")
}
