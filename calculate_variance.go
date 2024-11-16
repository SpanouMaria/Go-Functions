// Calculates variance of a slice of numbers, handling errors appropriately.

package main

import "errors"
import "math"


// CalculateVariance returns the variance of the numbers slice, or an error
// Variance is a measure of how far a set of numbers are spread out from their average value
func CalculateVariance(numbers []int) (*float64, error) {
	// Check if the input slice is nil or empty, return an error if true
	if numbers == nil || len(numbers) == 0 {
		// Error indicating invalid input
		return nil, errors.New("not implemented") 
	}
	
	// Convert the length of the numbers slice to a float for precision in calculation
	n := float64(len(numbers))

	// Sum is a helper function that applies a function (f) to each element of numbers 
	// and sums the results. It takes a function that operates on an int and returns a float64
	sum := func(f func(int) float64) float64 {
		var s float64
		for _, num := range numbers {
			// Add the result of applying f to each number to the sum
			s += f(num)
		}
		return s
	}

	// Calculate the mean (average) of the numbers by summing the numbers and dividing by n
	mean := sum(func(i int) float64 {
		return float64(i) // Convert each number to float64
	}) / n

	// Calculate the sum of squared differences from the mean, used in variance calculation
	sumRes := sum(func(i int) float64 {
		// For each number, subtract the mean, square the result, and sum them up
		return math.Pow(float64(i)-mean, 2)
	})

	// The variance is the sum of squared differences divided by the number of elements
	variance := sumRes / n
	
	// Return a pointer to the variance value and no error
	return &variance, nil
}