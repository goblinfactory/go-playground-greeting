package nethttp2

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

// Handler ...
type Handler interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}

// ResponseWriter ...
type ResponseWriter interface {
	Header() http.Header
	Write([]byte) (int, error)
	WriteHeader(statusCode int)
}

// TestMinimalHTTPServer ...
func TestMinimalHTTPServer() {

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
			log.Fatal(err)
		}
	}()

	time.Sleep(5 * time.Second)
	fmt.Println("closing server")
	cancel()
}
