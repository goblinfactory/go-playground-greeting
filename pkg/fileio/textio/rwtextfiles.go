package textio

import (
	"bufio"
	"log"
	"os"
)

// DemoTextio ...
func DemoTextio() {

}

// ReadAllLines ...
func ReadAllLines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	lines := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return lines
}

// func readWholeFileBuffered() {

// }

// WriteAllLines writes all lines to new file, overwrites (truncates) if already exists
func WriteAllLines(path string, lines []string) {
	file, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	for _, l := range lines {
		_, err := file.WriteString(l)
		if err != nil {
			log.Fatal(err)
		}
	}
}

// func writeWholeFileBuffered() {
// }
