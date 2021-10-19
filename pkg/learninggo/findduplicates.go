package learninggo

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// FindDuplicateLines prints the contents of a file and prints any duplicate lines in red.
func FindDuplicateLines() {
	findDuplicates(os.Args[1])
}

func findDuplicates(path string) {

	reset := string("\033[0m")
	red := string("\033[31m")
	green := string("\033[32m")

	defer fmt.Println(reset)

	lines := readAllLines(path)
	counts := countLines(lines)

	for i, l := range lines {
		fmt.Print(green, i, "\t", reset)
		if counts[l] > 1 {
			fmt.Println(red, l, reset)
		} else {
			fmt.Println(l)
		}
	}
}

func countLines(lines []string) map[string]int {
	counts := map[string]int{}
	for _, l := range lines {
		counts[l]++
	}
	return counts
}

func readAllLines(path string) []string {
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

// colorYellow := "\033[33m"
// colorBlue := "\033[34m"
// colorPurple := "\033[35m"
// colorCyan := "\033[36m"
// colorWhite := "\033[37m"
