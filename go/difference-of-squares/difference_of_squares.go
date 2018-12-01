package diffsquares

// Difference calculates the absolute distance between the sum of squares of n and the square of the sum of n
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}

// SumOfSquares is the sum of squares in the range 1..n
func SumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

// SquareOfSum is the square of the nth triangle number
func SquareOfSum(n int) int {
	nthTriangle := (n * (n + 1)) / 2
	return nthTriangle * nthTriangle
}
