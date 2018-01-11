// Package diffsquares gives us the difference between the sum of squares vs the squared sum of
// the first N natural numbers
package diffsquares

// SquareOfSums returns the squared value of the sum of the first N natural numbers
func SquareOfSums(num int) (total int) {
	if num <= 0 {
		return 0
	}

	for i := 1; i <= num; i++ {
		total += i
	}

	return total * total
}

// SumOfSquares returns the sum of the squares of the first N natural numbers
func SumOfSquares(num int) (total int) {
	if num <= 0 {
		return 0
	}

	for i := 1; i <= num; i++ {
		total += i * i
	}

	return total
}

// Difference yields the gap between two numbers
func Difference(num int) (diff int) {
	if num <= 0 {
		return 0
	}

	return SquareOfSums(num) - SumOfSquares(num)
}
