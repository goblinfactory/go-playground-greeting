package concurrencypatterns

import "fmt"

//DemoForSelectLoop ...
func DemoForSelectLoop() {
	fmt.Println("Lets print some stuff")
	con := make(chan string)
	fin := make(chan done)

	go printIt(con, fin)
	fmt.Println("all set!")
	con <- "one"
	con <- "two"
	con <- "three"

	// line below blocks until fin is read
	// simpler pattern to use instead of using syncWait
	fin <- done{}

	fmt.Println("this will not be reached before printIt finishes printing.")
}

type done struct{}

func printIt(ch chan string, done chan done) {
	for {
		select {
		case <-done:
			return
		case text := <-ch:
			fmt.Println(text)
		}
	}
}
