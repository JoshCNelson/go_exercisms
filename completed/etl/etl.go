// Package etl provides a method for transforming legacy scrabble score template to a newer format
package etl

import (
	"strings"
)

// Transform takes a legacy scrabble score template and transforms it to the new format
func Transform(extract map[int][]string) map[string]int {
	ret := make(map[string]int)

	for k, v := range extract {
		for _, letter := range v {
			ret[strings.ToLower(letter)] = k
		}
	}

	return ret
}
