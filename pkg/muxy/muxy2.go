package muxy

import (
	"fmt"

	"github.com/goblinfactory/greeting/pkg/rnd"
)

// EvenOdd runs two goroutines and signal to routines using signal channels.
func EvenOdd() {

	var stop bool = false
	var evens = make(chan struct{})
	var odds = make(chan struct{})

	nums := make(chan int, 100)
	go getNums(&stop, nums, 1, evens)
	go getNums(&stop, nums, 2, odds)

	go printNums(nums)
	rnd.SleepMs(2000)
	stop = true

	<-evens
	<-odds

}

func printNums(nums chan int) {
	for n := range nums {
		fmt.Printf("num:%d\n", n)
	}
}

func getNums(stop *bool, nums chan int, start int, fin chan struct{}) {
	// how to signal to this service to stop idiomatically without using * bool?
	i := start
	for {
		if *stop {
			fmt.Printf("stopping service [%d]\n", i)
			close(fin)
			return
		}
		nums <- i
		i = i + 2
		rnd.Sleep(100, 500)
	}
}
