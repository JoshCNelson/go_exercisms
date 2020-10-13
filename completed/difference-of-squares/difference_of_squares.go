// Package diffsquares provides functions for determining the sum of squares, the square of sums, and the difference between the two
package diffsquares

import "math"

// SquareOfSum takes a number and returns its SquareOfSum
func SquareOfSum(num int) int {
	sum := 0

	for i := 1; i < num+1; i++ {
		sum += i
	}

	return square(sum)
}

// SumOfSquares takes a number and returns its SumOfSquares
func SumOfSquares(num int) int {
	sum := 0

	for i := 1; i < num+1; i++ {
		sum += square(i)
	}

	return sum
}

// Difference returns the difference between the SumOfSquares and SquareOfSums for the supplied number
func Difference(num int) int {
	return SquareOfSum(num) - SumOfSquares(num)
}

func square(num int) int {
	return int(math.Pow(float64(num), 2))
}
