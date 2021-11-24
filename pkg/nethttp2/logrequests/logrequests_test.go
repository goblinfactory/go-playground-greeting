package logrequests

import (
	"net/http"
	"time"

	"github.com/goblinfactory/greeting/pkg/ansi"
	"github.com/goblinfactory/greeting/pkg/consolespikes"
)

// ToConsole logs requests to console. Register this middleware using mux.Handle(),
// returns a factory "HandlerFunc" that creates a new handler for each request
// that logs the time and requestURI to the console.
func ToConsole2(con consolespikes.Konsole, ansiColor string) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			con.WriteLine(ansiColor, time.Now().Format(time.RFC3339), r.RequestURI, ansi.Reset)
			h.ServeHTTP(w, r)
		})
	}
}
