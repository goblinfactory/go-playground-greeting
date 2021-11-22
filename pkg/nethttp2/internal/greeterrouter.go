package internal

// import (
// 	"github.com/goblinfactory/greeting/pkg/consolespikes"
// 	"github.com/mum4k/termdash/widgets/text"
// )

// // NewGreeterRouter ...
// func NewGreeterRouter(con *text.Text) NewServeMux {

// 	greeter := http.NewServeMux()

// 	greeter.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
// 		right.Write(r.RequestURI)
// 	})

// 	greeter.HandleFunc("/cat", func(rw http.ResponseWriter, r *http.Request) {
// 		right.Green(r.RequestURI)
// 		rw.Write([]byte("Meeoow!\n"))
// 		fmt.Println("/cat meeow")
// 	})
// 	greeter.HandleFunc("/dog", func(rw http.ResponseWriter, r *http.Request) {
// 		right.Green(r.RequestURI)
// 		rw.Write([]byte("Wooof!\n"))
// 		fmt.Println("/dog woof")
// 	})

// 	mux := http.NewServeMux()
// 	mux.Handle("/greet/", http.StripPrefix("/greet/", greeter))
// }
