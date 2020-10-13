// Package pangram contains a function for determining whether a string is a pangram
package pangram

import "unicode"

// IsPangram takes a string and returns whether or not it is a pangram
func IsPangram(str string) bool {
	var letterMap = map[rune]bool{
		'a': false, 'b': false, 'c': false, 'd': false,
		'e': false, 'f': false, 'g': false, 'h': false,
		'i': false, 'j': false, 'k': false, 'l': false,
		'm': false, 'n': false, 'o': false, 'p': false,
		'q': false, 'r': false, 's': false, 't': false,
		'u': false, 'v': false, 'w': false, 'x': false,
		'y': false, 'z': false,
	}

	for _, letter := range str {
		letterMap[unicode.ToLower(letter)] = true
	}

	for key := range letterMap {
		if !letterMap[key] {
			return false
		}
	}

	return true
}
