package nethttp2

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/goblinfactory/greeting/pkg/consolespikes"
)

// SpikeMinimalHTTPServer ...
func SpikeMinimalHTTPServer() {

	_ = os.Mkdir("logs", 0700)
	f, err := os.Create("logs/httpserver.log")
	if err != nil {
		panic(err)
	}
	log.SetOutput(f)

	left, right, wg, ctx, cancel, k := consolespikes.SplitLeftRight("server", "requests")

	cat := func(w http.ResponseWriter, r *http.Request) {

		right.Write(time.Now().Format(time.RFC3339))
		right.Green(" Meeoow! ")
		right.WriteLine(r.RequestURI)

		w.Write([]byte("Meeoow!\n"))
	}

	dog := func(w http.ResponseWriter, r *http.Request) {

		right.Write(time.Now().Format(time.RFC3339))
		right.Green(" Woof!   ")
		right.WriteLine(r.RequestURI)

		w.Write([]byte("Woof!!\n"))
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/cat/", cat)
	mux.HandleFunc("/dog/", dog)

	s := http.Server{
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  30 * time.Second,
		Handler:      mux,
	}

	defer func() {
		log.Printf("defer: shutting down")
		s.Shutdown(ctx)
	}()

	left.WriteLine("starting server")
	left.Gray("listening on port 8080 for http requests\n")
	left.Gray("/cat/...\n")
	left.Gray("/dog/...\n")
	left.WriteLine("press q to quit\n")

	k.OnQuit = func() {
		log.Printf("Shutting down server")
		s.Shutdown(ctx)
		cancel()
	}

	log.Printf("server starting.")

	err = s.ListenAndServe()

	if err != nil {
		left.Red(err)
		if err != http.ErrServerClosed {
			panic(err)
		}
	}
	log.Printf("server closed")
	wg.Wait()
}
