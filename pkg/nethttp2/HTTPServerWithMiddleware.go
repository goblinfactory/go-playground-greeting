package nethttp2

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// HTTPServerWithMiddleware ...
func HTTPServerWithMiddleware() {
	runFor := 21 * time.Second
	start := time.Now()

	history := &HistoryWriter{}

	indexrequest := func(w http.ResponseWriter, r *http.Request) {

		left := runFor - time.Since(start)
		msg := fmt.Sprintf("Hello from test2. %0.1f seconds left, till test server stops.", left.Seconds())
		w.Write([]byte(msg))
	}

	middleware := applyMiddelware(history)
	handlerFunc := middleware(http.HandlerFunc(indexrequest))

	s := http.Server{Addr: ":8081", Handler: handlerFunc}

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Duration(runFor))

	go s.ListenAndServe()
	<-ctx.Done()
	cancel()
	s.Close()
}

func applyMiddelware(hw *HistoryWriter) func(http.Handler) http.Handler { // this gives us access to historyWriter via closure
	return func(h http.Handler) http.Handler { // this gives us access to h.Serve in the closure
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { // this gives us access to the request in the closure

			// capture the bytes written by calling serveHTTP passing in
			// our fake responseWriter
			h.ServeHTTP(hw, r)
			// render the bytes we want displayed - full history (record of all requests)
			// this is because we want to see the count down the seconds left during each http request.
			// this will record ALL users making requests since we are not (for this demo) limiting this to a per-user session history.
			w.Write(hw.history)
		})
	}
}

// in english ... the above reads as ?
// a history writer middle ware factory, that returns a handler that when called returns a handlerFunc of the required middleware code
