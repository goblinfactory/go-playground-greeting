package nethttp2

import (
	"net/http"

	"github.com/goblinfactory/greeting/pkg/ansi"
	"github.com/goblinfactory/greeting/pkg/consolespikes"
	"github.com/goblinfactory/greeting/pkg/nethttp2/logrequests"
	"github.com/mum4k/termdash/linestyle"
)

// DemoServerWithLoggingMiddleware ...
func DemoServerWithLoggingMiddleware() {

	handler := func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("Hello world.")) }
	win, wg, _, _, kb := consolespikes.NewWindow("main", linestyle.Round)

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	loggedMux := logrequests.ToConsole(win, ansi.Green)
	s := http.Server{Addr: ":8080", Handler: loggedMux(mux)}

	kb.OnQuit = func() {
		s.Close()
	}
	win.WriteLine("starting server, requests will be logged to console, press 'q' to quit")
	s.ListenAndServe()
	wg.Wait()
}
