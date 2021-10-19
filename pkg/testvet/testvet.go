package testvet

import (
	"fmt"
	"sync"
)

// TestThatVetRunsOnSave minimal code to show vet not running on save in VS code
func TestThatVetRunsOnSave() {

	ch := make(chan string)
	var wg sync.WaitGroup
	wg.Add(1)
	go printit(&wg, ch)

	ch <- "one"
	ch <- "two"
	close(ch)
	wg.Wait()
	fmt.Println("done.")

}

func printit(wg *sync.WaitGroup, ch chan string) {
	defer wg.Done()
	for t := range ch {
		fmt.Println(t)
	}
}
