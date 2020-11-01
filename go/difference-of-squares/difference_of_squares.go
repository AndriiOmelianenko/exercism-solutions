package diffsquares

const testVersion = 1

// SquareOfSum calculates SquareOfSums
func SquareOfSum(n int) int {
	var square int
	for i := 1; i <= n; i++ {
		square += i
	}
	return square * square
}

// SumOfSquares calculates SumOfSquares
func SumOfSquares(n int) int {
	var sum int
	for i := 1; i <= n; i++ {
		sum += i * i
	}
	return sum
}

// Difference calculates difference between SquareOfSums and SumOfSquares
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
