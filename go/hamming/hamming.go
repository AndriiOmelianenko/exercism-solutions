//Package hamming calculates the Hamming difference between two DNA strands
package hamming

import (
	"errors"
)

const testVersion = 6

// Distance calculates the Hamming difference between two DNA strands
func Distance(a, b string) (int, error) {
	var counter int
	if len(a) != len(b) {
		return -1, errors.New("string lenghts are not the same")
	}
	for i := range a {
		if a[i] != b[i] {
			counter++
		}
	}
	return counter, nil
}
