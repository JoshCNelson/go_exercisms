// Package strain contains methods for keeping or discarding values in a collection
package strain

type Ints []int
type Lists [][]int
type Strings []string

// Ints.Keep keeps integers in the collection that return true based on the passed in function
func (in Ints) Keep(keepIt func(int) bool) Ints {
	var ret Ints

	for _, i := range in {
		if keepIt(i) {
			ret = append(ret, i)
		}
	}

	return ret
}

// Ints.Discard removes integers in the collection that return true based on the passed in function
func (in Ints) Discard(keepIt func(int) bool) Ints {
	var ret Ints

	for _, i := range in {
		if !keepIt(i) {
			ret = append(ret, i)
		}
	}

	return ret
}

// Lists.Keep keeps lists in the collection that return true based on the passed in function
func (l Lists) Keep(keepIt func([]int) bool) Lists {
	var ret Lists

	for _, i := range l {
		if keepIt(i) {
			ret = append(ret, i)
		}
	}

	return ret
}

// Strings.Keep keeps strings in the collection that return true based on the passed in function
func (s Strings) Keep(keepIt func(string) bool) Strings {
	var ret Strings

	for _, i := range s {
		if keepIt(i) {
			ret = append(ret, i)
		}
	}

	return ret
}
