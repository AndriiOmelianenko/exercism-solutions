package diffsquares

const testVersion = 1

func SquareOfSums(n int) int {
	var square int
	for i := 1; i <= n; i++ {
		square += i
	}
	return square * square
}

func SumOfSquares(n int) int {
	var sum int
	for i := 1; i <= n; i++ {
		sum += i * i
	}
	return sum
}

func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
