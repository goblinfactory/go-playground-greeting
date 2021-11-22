package nethttp2

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/goblinfactory/greeting/pkg/consolespikes"
	"github.com/goblinfactory/greeting/pkg/nethttp2/internal"
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

	//echoHandler := internal.NewMyConsoleEchoHandler(right)
	handler := internal.NewGreeterMux(right)

	s := http.Server{
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  30 * time.Second,
		Handler:      handler,
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
