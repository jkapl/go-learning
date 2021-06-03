package diffsquares

// SquareOfSum returns the square of the sum of the first n natural numbers
func SquareOfSum(n int) int {
	sum := (n * (n + 1)) / 2
	return sum * sum
}

// SumOfSquares returns the sum of the squares of the first n natural numbers
func SumOfSquares(n int) int {
	sum := (n * (n + 1) * (2*n + 1)) / 6
	return sum
}

// Difference returns the difference between the square of the sum of the first 
// n natural numbers and the sum of the squares of the first n natural numbers
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
