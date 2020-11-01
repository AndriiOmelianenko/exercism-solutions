// Package raindrops convert a number to a string, the contents of which depend on the number's factors
package raindrops

import (
	//"fmt"
	"strconv"
)

const testVersion = 3

// Convert convert a number to a string, the contents of which depend on the number's factors
func Convert(num int) string {
	var conv string
	var check int
	if num%3 == 0 {
		conv += "Pling"
		check = 1
	}
	if num%5 == 0 {
		conv += "Plang"
		check = 1
	}
	if num%7 == 0 {
		conv += "Plong"
		check = 1
	}
	if check == 0 {
		conv = strconv.FormatInt(int64(num), 10)
	}
	return conv
}
