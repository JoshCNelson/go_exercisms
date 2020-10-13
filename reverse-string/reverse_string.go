// Package reverse provides a function for reversing a string
package reverse

// Reverse takes a string and returns it in reverse order
func Reverse(str string) string {
	ret := ""
	runes := []rune(str)

	for i := len(runes) - 1; i >= 0; i-- {
		ret += string(runes[i])
	}

	return ret
}
