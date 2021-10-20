package errorhandling

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

// ErrFormat example sentinel error value
var ErrFormat = errors.New("Not an image file")

// DemoCorrectUseOfSentinelErrors ...
func DemoCorrectUseOfSentinelErrors() {
	compressed, err := CompressJpeg("123.gif")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(compressed)
}

// CompressJpeg ...
func CompressJpeg(file string) (string, error) {
	if !strings.Contains(file, ".jpeg") {
		return "", ErrFormat
	}
	return "compressed.zip", nil
}
