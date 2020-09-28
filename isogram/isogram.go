// Package isogram contains a function for determining whether or not a string is an isogram
package isogram

import (
	"strings"
)

type letterMap map[string]bool

// IsIsogram takes a string and returns whether or not the input is an isogram
func IsIsogram(str string) bool {
	if len(str) <= 1 {
		return true
	}

	lm := letterMap{}

	for _, r := range str {
		if r == ' ' || r == '-' {
			continue
		}

		s := strings.ToUpper(string(r))
		if _, ok := lm[s]; ok {
			return false
		}

		lm[s] = true
	}

	return true
}
