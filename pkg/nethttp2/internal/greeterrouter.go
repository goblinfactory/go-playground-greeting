package internal

import (
	"net/http"

	"github.com/goblinfactory/greeting/pkg/consolespikes"
)

// func aboutHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("<h1>This is the about page</h1>"))
// }

// NewGreeterMux ...
func NewGreeterMux(con consolespikes.Konsole) *http.ServeMux {

	// con.WriteLine("listening to routes:")
	// con.WriteLine("  /cat/greet")
	// con.WriteLine("  /dog/greet")
	// con.Gray("  /greeter/cat\n")
	// con.Gray("  /dog/greet\n")

	// func catHandler(w http.ResponseWriter, r *http.Request) {
	// 	con.Green(r.RequestURI)
	// 	w.Write([]byte("Meeoow!\n"))
	// }

	// // dog := http.NewServeMux()

	// // dog.HandleFunc("/dog", func(w http.ResponseWriter, r *http.Request) {
	// // 	con.Green(r.RequestURI)
	// // 	w.Write([]byte("Wooof!\n"))
	// // })

	// mux := http.NewServeMux()

	// // mux.Handle("/greet/cat/1", http.StripPrefix("/greet/", cat))
	// // mux.Handle("/greet/dog/1", http.StripPrefix("/greet/", dog))

	// mux.Handle("/cat", cat)

	return nil
}

// references : https://www.honeybadger.io/blog/go-web-services/
