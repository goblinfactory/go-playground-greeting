package byteio

import (
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadingWholeFileIntoMemory(t *testing.T) {

	b, err := os.ReadFile("./testdata/lisa1.gif")
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, 153295, len(b))
}

// see https://chmod-calculator.com/
const OwnerDefault = 0754 // rwxr-xr-- (octal)

func TestWritingBytesToFile(t *testing.T) {
	file1 := "./testdata/lisa1.gif"
	file2 := "./testdata/lisa2.gif"

	b, _ := os.ReadFile(file1)
	err := ioutil.WriteFile(file2, b, fs.FileMode(OwnerDefault))
	if err != nil {
		log.Fatal(err)
	}
}

// references

// os.xyz : low level access, control of file descriptors and permissions.
// https://medium.com/rungo/working-with-files-and-file-system-a-low-level-introduction-825ea3bac5f9

// ioutil : simple read and write
// https://medium.com/rungo/working-with-files-using-ioutil-package-2c526064febb
