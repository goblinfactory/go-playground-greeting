package nethttp

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"
)

type Handler interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}

type ResponseWriter interface {
	Header() http.Header
	Write([]byte) (int, error)
	WriteHeader(statusCode int)
}

type GreetHandler struct{}

// ServeHttp ..
func (gh GreetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Greetings!"))
}

func TestMinimalHttpServer(t *testing.T) {
	s := http.Server{
		Addr:         ":8080",
		ReadTimeout:  50 * time.Millisecond,
		WriteTimeout: 50 * time.Millisecond,
		IdleTimeout:  100 * time.Millisecond,
		Handler:      GreetHandler{},
	}

	bg := context.Background()
	ctx, cancel := context.WithCancel(bg)

	defer s.Shutdown(ctx)

	go func() {
		fmt.Println("starting server")
		err := s.ListenAndServe()
		if err != nil {
			t.Error(err)
		}
	}()

	time.Sleep(5 * time.Second)
	fmt.Println("closing server")
	cancel()
}
