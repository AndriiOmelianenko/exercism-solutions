package reverse

// Reverse reverses the input
func Reverse(input string) string {
	var reverse []rune
	for _, v := range input {
		reverse = append([]rune{v}, reverse...)
	}
	return string(reverse)
}
