package gzips

import (
	"archive/zip"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestReadingAZipWithMultipleFiles ..
func TestReadingZipContents(t *testing.T) {

	// read a zip file with multiple files
	expected := []string{
		"foozipped.go",
		"lisa1.gif",
	}

	z, err := zip.OpenReader("testdata/Multiple.zip")
	if err != nil {
		t.Error(err)
	}
	defer z.Close()

	assert.Equal(t, z.File[0].Name, expected[0])
	assert.Equal(t, z.File[1].Name, expected[1])
}

// TestReadingAZipWithMultipleFiles ..
func TestReadingAZipFile(t *testing.T) {

	fn := "iamcompressed.txt"
	z, err := zip.OpenReader("testdata/" + fn + ".zip")
	if err != nil {
		t.Error(err)
	}
	defer z.Close()

	assert.Equal(t, 1, len(z.File))

	// now read one of the files
	zf := z.File[0]
	assert.Equal(t, fn, zf.Name)

	f, err := zf.Open()
	if err != nil {
		t.Error(err)
	}

	b, err := io.ReadAll(f)
	if err != nil {
		t.Error(err)
	}

	text := string(b)

	te := "If you are reading this, then my hamster is dead."
	assert.Equal(t, te, text)
}
