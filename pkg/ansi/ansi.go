package ansi

import (
	"fmt"
	"strings"
)

// Reset resets the colors back to console default
var Reset = "\033[0m"

// DarkYellow ...
var DarkYellow = "\033[33m"

// Red ...
var Red = "\033[31m"

// Green ...
var Green = "\033[32m"

// Cls clears the terminal screen
var Cls = "\033[H\033[2J"

// White ...
var White = "\033[29m"

// Gray ...
var Gray = "\033[90m"

// PrintColors prints a few colors. At a quick glance of what this produces it doesnt look correct. So just an interesting spike
// there are proper ansi libraries for doing this.
// for example see : https://github.com/gookit/color
func PrintColors() {
	defer fmt.Println(Reset)
	color := "\033[XXm"
	for i := 0; i < 100; i++ {
		nc := strings.Replace(color, "XX", fmt.Sprintf("%00d", i), -1)
		fmt.Println(nc, "color", i)
	}
}

// reference : https://medium.com/@inhereat/terminal-color-rendering-tool-library-support-8-16-colors-256-colors-by-golang-a68fb8deee86
