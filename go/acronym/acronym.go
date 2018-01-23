// Package acronym converts a phrase to its acronym
package acronym

import "strings"

const testVersion = 3

func Abbreviate(phrase string) string {
	var acronym string
	for i := range phrase {
		if i == 0 || string(phrase[i-1]) == " " || string(phrase[i-1]) == "-" {
			acronym = acronym + strings.ToUpper(string(phrase[i]))
		}
	}
	return acronym
}
