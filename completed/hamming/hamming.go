// Package hamming provides methods for calculating the Hamming Distance between strands of DNA.
package hamming

import "fmt"

// Distance calculates the Hamming Distance between two strands of equal length DNA
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, fmt.Errorf("Cannot compare strands of unequal lengths.\nLength of a: %d\nLength of b: %d\n", len(a), len(b))
	}

	hammingDistance := 0
	for i := range a {
		if a[i] != b[i] {
			hammingDistance++
		}
	}

	return hammingDistance, nil
}
