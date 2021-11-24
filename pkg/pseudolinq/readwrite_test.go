package pseudolinq

import (
	"fmt"
	"testing"

	"github.com/goblinfactory/greeting/pkg/fileio/textio"
)

func TestReadingWritingActualFiles(t *testing.T) {
	const file = "testreadwrite.txt"
	textio.WriteAllLines(file, []string{"one", "two", "three"})
	fmt.Println("lines", textio.ReadAllLines(file))
	textio.WriteAllLines(file, New(textio.ReadAllLines(file)).Suffix(" * done"))
	fmt.Println("lines", textio.ReadAllLines(file))
}

func TestReadingFiles(t *testing.T) {
	const file = "./testdata/testreadwrite.txt"
	lines := textio.ReadAllLines(file)
	fmt.Println("lines", lines)
}
