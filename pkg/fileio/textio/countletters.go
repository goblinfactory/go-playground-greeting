package textio

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// DemoCountLettersUsingStringReader ..
func DemoCountLettersUsingStringReader() {

	//countLetters("readme.md")
	s := "this is a test string that I want to count letters from"
	sr := strings.NewReader(s)

	m, err := countLetters(sr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("counted letters")
	fmt.Println(m)
}

// DemoCountLettersUsingOsOpen ..
func DemoCountLettersUsingOsOpen() {

	file, err := os.Open("readme.md")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	m, err := countLetters(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("counted letters")
	fmt.Println(m)
}

func countLetters(r io.Reader) (map[string]int, error) {
	buffer := make([]byte, 1024)
	out := map[string]int{}
	for {
		n, err := r.Read(buffer)
		for _, b := range buffer[:n] {
			if (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') {
				out[string(b)]++
			}
		}
		if err == io.EOF {
			return out, nil
		}
		if err != nil {
			return nil, err
		}
	}
}

func countWords(r io.Reader) (map[string]int, error) {
	buffer := make([]byte, 1024)
	out := map[string]int{}
	for {
		n, err := r.Read(buffer)
		for _, b := range buffer[:n] {
			if (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') {
				out[string(b)]++
			}
		}
		if err == io.EOF {
			return out, nil
		}
		if err != nil {
			return nil, err
		}
	}
}
