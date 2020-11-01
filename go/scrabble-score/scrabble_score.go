package scrabble

import (
	"strings"
)

var prices = map[string]int{
	"AEIOULNRST": 1,
	"DG":         2,
	"BCMP":       3,
	"FHVWY":      4,
	"K":          5,
	"JX":         8,
	"QZ":         10,
}

// Score counts the number of occurances
func Score(word string) int {
	var score int
	word = strings.ToUpper(word)
	for _, letter := range word {
		for lettersWithPrice := range prices {
			if strings.ContainsRune(lettersWithPrice, letter) {
				score += prices[lettersWithPrice]
			}
		}
	}
	return score
}
