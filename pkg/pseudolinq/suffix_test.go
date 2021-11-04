package pseudolinq

import (
	"fmt"
	"testing"

	"github.com/goblinfactory/greeting/pkg/fileio/textio"
	"github.com/stretchr/testify/assert"
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

func TestSuffix(t *testing.T) {
	lines := []string{"one", "two", "three"}
	actual := New(lines).Suffix(".")
	expected := []string{"one.", "two.", "three."}
	assert.Equal(t, expected, actual)
}
