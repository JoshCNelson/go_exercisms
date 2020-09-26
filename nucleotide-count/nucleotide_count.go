// Package dna provides functionality for outputting nucleotide counts in a strand of dna
package dna

import (
	"fmt"
)

type Histogram map[rune]int
type DNA string

// Counts returns a Histogram of nucleotides in a string of DNA
func (d DNA) Counts() (Histogram, error) {
	h := Histogram{'A': 0, 'C': 0, 'T': 0, 'G': 0}

	for _, letter := range d {
		if letter == 'A' || letter == 'C' || letter == 'G' || letter == 'T' {
			h[letter]++
		} else {
			return nil, fmt.Errorf("strang with invalid nucleotide provided")
		}
	}

	return h, nil
}
