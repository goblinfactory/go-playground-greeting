package internal

import (
	"fmt"
	"net/http"

	"github.com/goblinfactory/greeting/pkg/consolespikes"
)

// MyConsoleEchoHandler ...
type MyConsoleEchoHandler struct {
	con consolespikes.Konsole
}

// NewMyConsoleEchoHandler ...
func NewMyConsoleEchoHandler(con consolespikes.Konsole) MyConsoleEchoHandler {
	return MyConsoleEchoHandler{
		con,
	}
}

func (h *MyConsoleEchoHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.con.Green(fmt.Sprintf("%s ", r.Method))
	h.con.WriteLine(r.RequestURI)

	rw.Write([]byte(r.RequestURI))
}
