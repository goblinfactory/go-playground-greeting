package nethttp2

import (
	"net/http"
	"testing"
)

func TestInterfaceIsImplemented(t *testing.T) {
	doesItWrite := func(w http.ResponseWriter) {}
	hw := &HistoryWriter{}
	doesItWrite(hw)
}
