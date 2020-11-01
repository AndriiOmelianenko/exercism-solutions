package isogram

import (
	"strings"
)

// IsIsogram checks if word is isogram
func IsIsogram(word string) bool {
	var alphabet string = "abcdefghijklmnopqrstuvwxyz"

	removeCharacters := []string{" ", "-"}
	for _, rc := range removeCharacters {
		word = strings.Replace(word, rc, "", -1)
	}

	word = strings.ToLower(word)

	for _, v := range word {
		if strings.ContainsRune(alphabet, v) {
			alphabet = strings.Replace(alphabet, string(v), "", 1)
		} else {
			return false
		}
	}
	return true
}
