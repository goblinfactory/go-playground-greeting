package concurrencypatterns

import (
	"fmt"
	"net/http"
	"sync"
)

// DemoUsingDoneChannel ...
func DemoUsingDoneChannel() {
	ws := []string{
		"https://www.goblinfactory.co.uk",
		"https://www.google.co.uk",
		"https://cnn.com",
	}

	var wg sync.WaitGroup
	f := returnFastestWebsiteToRespond(&wg, ws)
	fmt.Println("fastest", f.status, f.url)
	wg.Wait()
}

type result struct {
	status int
	url    string
}

func returnFastestWebsiteToRespond(wg *sync.WaitGroup, urls []string) result {
	done := make(chan struct{})
	results := make(chan result)

	for _, ws := range urls {
		wg.Add(1)
		fmt.Println("GET ", ws)
		go func(url string) {
			select {
			case results <- getStatus(wg, url):
			case <-done: // interestingly we're not blocking waiting to read from the channel, we're blocking until channel is closed, when then allows us to read default zero value.
			}
		}(ws)
		// it was tempting at first to put wg.Done here, and the program would exit
		// faster, however that would have left the goRoutines running and the
		// threads killed without any defer's (cleanups) being called.
		// Almost always this is the wrong approach and can cause memory leaks and-or data inconsistencies.
	}
	r := <-results // now we read the first result, and then close the done chanel
	close(done)
	return r
}

func getStatus(wg *sync.WaitGroup, url string) result {
	defer wg.Done()
	r, _ := http.Get(url) // in VScode setting breakpoint here appears to stop on a random thread. Need to test this in goland.
	fmt.Println("response", url, "status", r.StatusCode)
	return result{r.StatusCode, url}
}

/*

Running this code produces

GET  https://www.goblinfactory.co.uk
GET  https://www.google.co.uk
GET  https://cnn.com
response https://www.goblinfactory.co.uk status 200
fastest 200 https://www.goblinfactory.co.uk
response https://www.google.co.uk status 200
response https://cnn.com status 200

*/
