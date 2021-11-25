package nethttp2

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestCreatingOurOwnHandler(t *testing.T) {
	runFor := 2 * time.Second
	start := time.Now()

	handler := func(w http.ResponseWriter, r *http.Request) {

		left := runFor - time.Since(start)
		msg := fmt.Sprintf("Hello from test2. %0.1f seconds left, till test server stops.", left.Seconds())
		w.Write([]byte(msg))
	}

	s := http.Server{Addr: ":8081", Handler: http.HandlerFunc(handler)}

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Duration(runFor))

	go s.ListenAndServe()
	<-ctx.Done()
	cancel()
	s.Close()
}
