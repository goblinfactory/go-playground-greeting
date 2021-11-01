package concurrencypatterns

import "fmt"

// WowThisIsSubtle shows examples of noobie Go code that can really trip you up. Code examples show using a goRoutine to generate a list; unless generating the list is expensive or slow. Normally this is something you shouldnt do, (wrong use of concurrency) but is show here because it's interesting to look at the details as well as compare how the similar code can be written in C#. see subtle.md for a spike of a side by side comparison.
func WowThisIsSubtle() {
	for n := range generateNumbers(10) {
		fmt.Println(n)
	}
}

func generateNumbers(cnt int) <-chan int {
	ch := make(chan int)
	// start a goroutine to push 10 ints to the channel
	go func() {
		for i := 0; i < cnt; i++ {
			fmt.Printf("Adding: %d\n", i)
			ch <- i
		}
		close(ch)
	}()
	// line below doesnt actually close the channel
	// until the code gets here.
	// this goroutine runs and blocks t the first ch<- i line, above, line [12]
	// until it's actually read by a consumer
	return ch
}
