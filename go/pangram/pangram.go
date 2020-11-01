package pangram

import (
	"strings"
)

// IsPangram check is sentence is pangram
func IsPangram(sentence string) bool {
	for i := 'a'; i <= 'z'; i++ {
		if strings.Contains(strings.ToLower(sentence), string(i)) == false {
			return false
		}
	}
	return true
}
