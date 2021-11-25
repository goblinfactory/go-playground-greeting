package nethttp2

import "net/http"

// HistoryWriter ...
type HistoryWriter struct {
	history []byte
	bytes   []byte
}

// Header ...
func (fw *HistoryWriter) Header() http.Header {
	return nil
}

// Write...
func (fw *HistoryWriter) Write(bytes []byte) (int, error) {
	// not sure if this needs to be threadsafe?
	fw.bytes = bytes
	fw.history = append(fw.history, []byte(string("\n"))...)
	fw.history = append(fw.history, bytes...)
	return len(bytes), nil
}

// WriteHeader ...
func (fw *HistoryWriter) WriteHeader(statusCode int) {

}
