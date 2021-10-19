package dependencyinjection

import (
	"fmt"
	"net/http"
)

// Main ..
func Main() {

	l := LoggerAdapter(LogOutput)
	ds := NewSimpleDataStore()
	logic := NewSimpleLogic(l, ds)
	c := NewController(l, logic)
	http.HandleFunc("/hello", c.SayHello)
	fmt.Println("starting")
	http.ListenAndServe(":8080", nil)
	fmt.Println("stopped.")
}
