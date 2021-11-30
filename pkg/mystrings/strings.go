package mystrings

import "strings"

// ContainsAny returns true if any of the strings are in the string.
func ContainsAny(src string, matches ...string) bool {
	for _, s := range matches {
		if strings.Contains(src, s) {
			return true
		}
	}
	return false
}
