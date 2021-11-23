package backpressure

// fakereaderwriter is a demo of using backpressure to reduce the speed of reading so that it matches the speed of writing, not yet implemented

// func ControlTheProducer() {
// 	r := fakeReader{100, "lisa1.gif", 200}
// 	w := fakeWriter{10, "lisa2.gif", 200}

// }

// type fakeReader struct {
// 	BytesPerSecond int
// 	file           string
// 	buffersize     int
// }

// type fakeWriter struct {
// 	BytesPerSecond int
// 	file           string
// 	buffersize     int
// }

// // StartReading ...
// func (r fakeReader) StartReading(file string) {
// 	f, err := io.rea
// }
