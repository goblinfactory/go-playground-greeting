package concurrencypatterns

import (
	"fmt"
	"net/http"

	"github.com/goblinfactory/greeting/pkg/checkablewaitgroup.go"
)

// DemoUsingDoneChannel ...
func DemoUsingDoneChannel() {
	ws := []string{
		"https://www.goblinfactory.co.uk",
		"https://www.google.co.uk",
		"https://cnn.com",
	}

	wg := checkablewaitgroup.New()
	f := returnFastestWebsiteToRespond(wg, ws)
	fmt.Println("fastest", f.url)
	wg.Wait()
}

const indent = "   "

type result struct {
	status int
	url    string
}

func returnFastestWebsiteToRespond(wg *checkablewaitgroup.WaitGroup, urls []string) result {
	done := make(chan struct{})
	results := make(chan result)

	for _, ws := range urls {
		wg.Add(1)
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

func getStatus(wg *checkablewaitgroup.WaitGroup, url string) result {
	defer wg.Done()
	fmt.Println("GET ", url)
	// dont use http.Get etc in production, it has no timeout.
	r, _ := http.Get(url) // in VScode setting breakpoint here appears to stop on a random thread. Need to test this in goland.
	printResponse(wg, r, url)
	return result{r.StatusCode, url}
}

func printResponse(wg *checkablewaitgroup.WaitGroup, r *http.Response, url string) {
	if wg.IsDone() {
		fmt.Println(indent, "slower", url, "status", r.StatusCode)
		return
	}
	fmt.Println(indent, "first", url, "status", r.StatusCode)
}

/*

Running this code produces

GET  https://cnn.com
GET  https://www.google.co.uk
GET  https://www.goblinfactory.co.uk
    first https://www.goblinfactory.co.uk status 200
fastest https://www.goblinfactory.co.uk
    slower https://www.google.co.uk status 200
    slower https://cnn.com status 200

*/
