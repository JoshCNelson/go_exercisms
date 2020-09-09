// Package proverb provides a function for turning a slice of strings into a short proverb
package proverb

import "fmt"

// Proverb turns each string in []string into a line of a short proverb
func Proverb(rhyme []string) []string {
	ret := []string{}

	for i := 0; i < len(rhyme)-1; i++ {
		line := fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1])
		ret = append(ret, line)
	}

	if len(rhyme) > 0 {
		ret = append(ret, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
	}

	return ret
}
