// Package anagram provides function for determining if a slice of words is an anagram
package anagram

import (
	"sort"
	"strings"
)

// Detect takes a string and a slice of candidates and returns a slice of strings that are anagrams of the string
func Detect(str string, candidates []string) []string {
	var ret []string
	sortedStr := sortString(strings.ToLower(str))

	for _, candidate := range candidates {
		if len(candidate) != len(sortedStr) || strings.ToLower(candidate) == strings.ToLower(str) {
			continue
		}

		sortedCandidate := sortString(strings.ToLower(candidate))

		if sortedCandidate == sortedStr {
			ret = append(ret, candidate)
		}
	}

	return ret
}

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
