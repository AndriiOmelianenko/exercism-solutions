// Package accumulate implement the accumulate operation, which, given a collection and an operation to perform on each element of the collection, returns a new collection containing the result of applying that operation to each element of the input collection.
package accumulate

const testVersion = 1

func Accumulate(collection []string, customFunc func(string) string) []string {
	var output []string
	for _, value := range collection {
		output = append(output, customFunc(value))
	}
	return output
}
