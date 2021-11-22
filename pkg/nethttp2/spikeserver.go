package nethttp2

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/goblinfactory/greeting/pkg/consolespikes"
	"github.com/mum4k/termdash/widgets/text"
)

// MyConsoleEchoHandler ...
type MyConsoleEchoHandler struct {
	con consolespikes.Konsole
}

func newMyConsoleEchoHandler(con *text.Text) MyConsoleEchoHandler {
	return MyConsoleEchoHandler{
		consolespikes.NewKonsole(con),
	}
}

func (h *MyConsoleEchoHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.con.GreenLine(r.RequestURI)
}

// SpikeMinimalHTTPServer ...
func SpikeMinimalHTTPServer() {

	_ = os.Mkdir("logs", 0700)
	f, err := os.Create("logs/httpserver.log")
	if err != nil {
		panic(err)
	}
	log.SetOutput(f)

	_left, right, wg, ctx, cancel, k := consolespikes.SplitLeftRight("server", "requests")
	left := consolespikes.NewKonsole(_left)

	echoHandler := newMyConsoleEchoHandler(right)
	// greeter := http.NewServeMux()

	// greeter.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
	// 	right.Write(r.RequestURI)
	// })

	// greeter.HandleFunc("/cat", func(rw http.ResponseWriter, r *http.Request) {
	// 	right.Green(r.RequestURI)
	// 	rw.Write([]byte("Meeoow!\n"))
	// 	fmt.Println("/cat meeow")
	// })
	// greeter.HandleFunc("/dog", func(rw http.ResponseWriter, r *http.Request) {
	// 	right.Green(r.RequestURI)
	// 	rw.Write([]byte("Wooof!\n"))
	// 	fmt.Println("/dog woof")
	// })

	// mux := http.NewServeMux()
	// mux.Handle("/greet/", http.StripPrefix("/greet/", greeter))

	s := http.Server{
		Addr:         ":8080",
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
		IdleTimeout:  10 * time.Second,
		Handler:      &echoHandler,
	}

	defer func() {
		log.Printf("defer: shutting down")
		s.Shutdown(ctx)
	}()

	left.Write("starting server,press q to quit\n")

	k.OnQuit = func() {
		log.Printf("Shutting down server")
		s.Shutdown(ctx)
		cancel()
	}

	log.Printf("server started.")
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
