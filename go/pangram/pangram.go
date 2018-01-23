package pangram

import (
	"strings"
)

const testVersion = 2

func toChar(i int) string {
	return string('a' - 1 + i)
}
func IsPangram(sentence string) bool {
	for i := 1; i <= 26; i++ {
		if strings.Contains(strings.ToLower(sentence), toChar(i)) == false {
			return false
		}
	}
	return true
}
