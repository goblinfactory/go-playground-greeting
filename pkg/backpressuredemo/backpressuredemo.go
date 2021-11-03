package backpressuredemo

// ReadingFasterThanWritingWithoutPressure is a demo that shows how buffer keeps growing when your reader reads faster than your writer without
// any backpressure technique.
// func ReadingFasterThanWritingWithoutPressure() {
// 	r := NewReader(100)
// 	w := NewWriter(200)

// }

// ReadingFasterThanWritingWithBackPressure is a demo of using backpressure to reduce the speed of reading so that it matches the speed of 	writing, to keep the size of buffer small.
func ReadingFasterThanWritingWithBackPressure() {

}

// FakeReader ...
type FakeReader struct {
	BytesPerSecond int
}

// FakeWriter ...
type FakeWriter struct {
	BytesPerSecond int
}

// NewReader returns a fake reader that will simulate reading at specified bytes per second
func NewReader(bytesPerSecond int, filename string) FakeReader {
	return FakeReader{bytesPerSecond}
}

// NewWriter returns a fake writer that will simulate reading at specified bytes per second
func NewWriter(bytesPerSecond int, filename string) FakeWriter {
	return FakeWriter{bytesPerSecond}
}

// StartReading
func (r FakeReader) StartReading() {

}
