package ignorebackpressure

// ReadingFasterThanWritingWithoutPressure is a demo that shows how buffer keeps growing when your reader reads faster than your writer without
// any backpressure technique.
// func ReadingFasterThanWritingWithoutPressure() {
// 	r := NewReader(100)
// 	w := NewWriter(200)

// }

// ControlTheProducer is a demo of using backpressure to reduce the speed of reading so that it matches the speed of writing
func ControlTheProducer() {
	// r := NewReader(100, "lisa1.gif", 200)
	// w := NewWriter(10, "lisa2.gif", 200)

}

// StartReading ...
func (r FakeReader) StartReading() {

}

// FakeReader ...
type FakeReader struct {
	BytesPerSecond int
	file           string
	buffersize     int
}

// FakeWriter ...
type FakeWriter struct {
	BytesPerSecond int
	file           string
	buffersize     int
}

// NewReader returns a fake reader that will simulate reading at specified bytes per second
func NewReader(bytesPerSecond int, filename string, buffersize int) FakeReader {
	return FakeReader{bytesPerSecond, filename, buffersize}
}

// NewWriter returns a fake writer that will simulate reading at specified bytes per second
func NewWriter(bytesPerSecond int, filename string, buffersize int) FakeWriter {
	return FakeWriter{bytesPerSecond, filename, buffersize}
}
