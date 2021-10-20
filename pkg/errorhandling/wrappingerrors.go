package errorhandling

import (
	"fmt"
	"os"
)

// DemoWrappingErrors ..
func DemoWrappingErrors() {
	fmt.Println(checkFile("readme.md"))
	fmt.Println(checkFile("not-exist.md"))
}

func checkFile(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return fmt.Errorf("in GetFile: %w", err)
	}
	f.Close()
	return nil
}
