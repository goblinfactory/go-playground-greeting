package byteio

import (
	"log"
	"os"
)

// ReadAllBytes ...
// func ReadAllBytes(path string) []byte {
// 	file, err := os.Open(path)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer file.Close()
// 	r := bufio.NewReader(file)
// 	b := file.Read(r)
// }

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

// references
// https://stackoverflow.com/questions/1821811/how-to-read-write-from-to-a-file-using-go
