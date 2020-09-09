// Package raindrops provides a function for returning a pre-defined string pattern based on the the provided integer being a factor of 3,5, or 7
package raindrops

import (
	"strconv"
)

// Convert outputs a pre-defined string pattern depending on whether the int argument is a factor or 3,5, or 7
func Convert(num int) string {
	b := ""

	if num%3 == 0 {
		b += "Pling"
	}

	if num%5 == 0 {
		b += "Plang"
	}

	if num%7 == 0 {
		b += "Plong"
	}

	if len(b) == 0 {
		b += string(strconv.Itoa(num))
	}

	return b
}
