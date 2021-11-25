package nethttp2

import (
	"net/http"
	"testing"
)

func TestCreatingOurOwnHandler(t *testing.T) {

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from test."))
	}

	s := http.Server{Addr: ":8081", Handler: handler}

}
