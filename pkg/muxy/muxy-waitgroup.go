package muxy

import (
	"fmt"
	"time"

	"github.com/goblinfactory/greeting/pkg/rnd"
)

// EvenOddWaitGroup test multiplexing using select
func EvenOddWaitGroup() {

	printNums := func(finished chan struct{}, nums chan int) {
		for n := range nums {
			fmt.Printf("num:%d\n", n)
		}
		fmt.Println("Num printer stopped.")
		close(finished)
	}

	var stopEvens = make(chan struct{})
	var evensStopped = make(chan struct{})

	var stopOdds = make(chan struct{})
	var oddsStopped = make(chan struct{})

	var numsStopped = make(chan struct{})

	nums := make(chan int, 100)

	go getNumsIdiomatic(nums, 1, stopEvens, evensStopped)
	go getNumsIdiomatic(nums, 2, stopOdds, oddsStopped)

	go printNums(numsStopped, nums)
	rnd.SleepMs(2000)

	fmt.Println("Stopping all")

	close(stopEvens)
	close(stopOdds)

	// wait for each service to shutdown
	<-evensStopped
	<-oddsStopped

	// close UX service last
	close(nums)
	<-numsStopped

	fmt.Println("FINISHED")
}

func getNumsIdiomatic(nums chan int, start int, stop chan struct{}, finished chan struct{}) {
	// how to signal to this service to stop idiomatically without using * bool?
	i := start
	for {
		select {
		case <-stop:
			{
				time.Sleep(2000)
				fmt.Printf("stopping service [%d]\n", i)
				close(finished)
				return
			}
		default:
			nums <- i
			i = i + 2
			rnd.SleepMinMaxMs(100, 500)
		}

	}
}

func sleepMs(ms int) {
	time.Sleep(time.Duration(ms * int(time.Millisecond)))
}
