package bob

import (
	"strings"
	"unicode"
)

const testVersion = 3

// Hey greets you
func Hey(input string) string {
	var upperLetters, lowerLetters int
	removeCharacters := []string{" ", "\t", "\n", "\r"}
	for _, rc := range removeCharacters {
		input = strings.Replace(input, rc, "", -1)
	}
	if input == "" {
		return "Fine. Be that way!"
	}
	for _, c := range input {
		if unicode.IsUpper(c) {
			upperLetters++
		} else if unicode.IsLower(c) {
			lowerLetters++
		}
	}
	if string(input[len(input)-1]) == "?" && upperLetters > lowerLetters {
		return "Calm down, I know what I'm doing!"
	}
	if upperLetters > lowerLetters {
		return "Whoa, chill out!"
	}
	if string(input[len(input)-1]) == "?" {
		return "Sure."
	}
	return "Whatever."
}
