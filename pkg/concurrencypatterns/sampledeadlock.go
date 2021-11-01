package concurrencypatterns

import "fmt"

// DemoDeadlocking shows an example of two goroutines that access the same two channels
// in different order, causing a deadlock
func DemoDeadlocking() {

	num := make(chan int)
	result := make(chan int)

	go add10(num, result)

	num <- 5                       // write to nums
	fmt.Println("5+10=", <-result) // read from result

	result2 := <-result

	num <- 10
	fmt.Println(result2)

	// running this code in an IDE will hang (deadlock) forever.
	// running this code from a release build will give you the following error
	// fatal error: all goroutines are asleep - deadlock!
}

func add10(nums chan int, result chan int) {
	n := <-nums      // read from nums
	result <- n + 10 // write to result
}

// Authors note;
// I struggled with this a few hours, trying to figure out a realistic tempting scenario where you could unwittingly find yourself
//
