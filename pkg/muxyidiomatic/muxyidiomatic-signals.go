package muxyidiomatic

import (
	"fmt"
	"time"

	"github.com/goblinfactory/greeting/pkg/rnd"
)

// EvenOddWithSignals foo
func EvenOddWithSignals() {

	var evens = make(chan bool)
	var evensFinished = make(chan bool)

	var odds = make(chan bool)
	var oddsFinished = make(chan bool)

	nums := make(chan int, 100)

	go streamNums(nums, 1, odds, oddsFinished)
	go streamNums(nums, 2, evens, evensFinished)

	go printNums(nums)
	rnd.SleepMs(2000)

	fmt.Println("Stopping all")
	evens <- true
	odds <- true

	// wait for each service to shutdown
	<-evensFinished
	<-oddsFinished
	close(nums)
	fmt.Println("FINISHED")

}

func printNums(nums chan int) {
	for n := range nums {
		fmt.Printf("num:%d\n", n)
	}
}

func streamNums(nums chan int, start int, stop chan bool, finished chan bool) {
	i := start
	for {
		select {
		case <-stop:
			{
				time.Sleep(2000)
				fmt.Printf("stopping service [%d]\n", i)
				finished <- true
				return
			}
		default:
			nums <- i
			i = i + 2
			rnd.Sleep(100, 500)
		}

	}
}
