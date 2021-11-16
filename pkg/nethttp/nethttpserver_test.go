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

func TestMinimalHttpServer(t *testing.T) {

	greeter := http.NewServeMux()
	greeter.HandleFunc("/cat", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("Meeoow!\n"))
		fmt.Println("/cat meeow")
	})
	greeter.HandleFunc("/dog", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("Wooof!\n"))
		fmt.Println("/dog woof")
	})

	mux := http.NewServeMux()
	mux.Handle("/greet/", http.StripPrefix("/greet/", greeter))

	s := http.Server{
		Addr:         ":8080",
		ReadTimeout:  50 * time.Millisecond,
		WriteTimeout: 50 * time.Millisecond,
		IdleTimeout:  100 * time.Millisecond,
		Handler:      mux,
	}

	bg := context.Background()
	ctx, cancel := context.WithCancel(bg)

	defer s.Shutdown(ctx)

	go func() {
		fmt.Println("starting server for 20 seconds")
		err := s.ListenAndServe()
		if err != nil {
			t.Error(err)
		}
	}()

	time.Sleep(20 * time.Second)
	fmt.Println("closing server")
	cancel()
}
