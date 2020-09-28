// Package accumulate contains a function for taking a collection and performing a provided operation on each member of the collection
package accumulate

// Accumulate takes the provided collection and performs the provided function on each member of the collection
func Accumulate(collection []string, transform func(string) string) []string {
	for i := range collection {
		collection[i] = transform(string(collection[i]))
	}

	return collection
}
