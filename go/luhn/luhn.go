package luhn

import (
	"math"
	"strconv"
	"strings"
	"unicode"
)

// Valid checks if card is corrrect
func Valid(input string) bool {
	var doubling int
	cardNum := strings.ReplaceAll(input, " ", "")
	if len(cardNum) <= 1 {
		return false
	}
	for i, v := range cardNum {

		if !unicode.IsDigit(v) {
			return false
		}

		vi, err := strconv.Atoi(string(v))
		if err != nil {
			panic(err)
		}
		if (i%2 == 0 && len(cardNum)%2 == 0) || (i%2 != 0 && len(cardNum)%2 != 0) {
			vi *= 2
			if vi > 9 {
				vi -= 9
			}
		}
		doubling += vi * int(math.Pow10((len(cardNum) - i - 1)))
	}
	if sumDigits(doubling)%10 == 0 {
		return true
	}
	return false
}

func sumDigits(number int) int {
	var sumResult int
	for number != 0 {
		sumResult += number % 10
		number = number / 10
	}
	return sumResult
}
