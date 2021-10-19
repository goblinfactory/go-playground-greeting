package testwaitgroup

import (
	"fmt"
	"net/http"
	"sync"
)

type httpPkg struct{}

func (httpPkg) Get(url string) {}

// TestWaitGroup runs a simplified test that shows how to use wg to wait for goroutines.
func TestWaitGroup() {
	ch := make(chan string, 5)
	var wg sync.WaitGroup
	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.somestupidname.com/",
	}

	for _, url := range urls {
		wg.Add(1)
		// Increment the WaitGroup counter
		// Launch a goroutine to fetch the URL.
		go func(ch chan string, url string) {
			// Decrement the counter when the goroutine completes.
			fmt.Printf("fetching '%s'\n", url)
			defer wg.Done()
			// Fetch the URL.
			r, _ := http.Get(url)
			ch <- fmt.Sprintf("[%s] %s", r.Status, url)
		}(ch, url)
	}
	// Wait for all HTTP fetches to complete.

	go func(wg *sync.WaitGroup, ch chan string) {
		for t := range ch {
			fmt.Println(t)
		}
	}(&wg, ch)

	fmt.Println("waiting for dowloads to finish")
	wg.Wait()
	fmt.Println("finished!")
}
