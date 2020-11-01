// Package acronym converts a phrase to its acronym
package acronym

import (
	"strings"
	"unicode"
)

const testVersion = 3

// Abbreviate converts a phrase to its acronym
func Abbreviate(phrase string) string {
	var acronym string
	for i := range phrase {
		if (i == 0 || (string(phrase[i-1]) == " ") || string(phrase[i-1]) == "-" || string(phrase[i-1]) == "_") && unicode.IsLetter(rune(phrase[i])) {
			acronym = acronym + strings.ToUpper(string(phrase[i]))
		}
	}
	return acronym
}
