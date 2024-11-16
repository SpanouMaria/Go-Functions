// Calculates variance using concurrency for efficient summation of numbers.

package main

import "errors"
import "math"


// CalculateVariance returns the variance of the numbers slice, or an error
// Variance is a measure of how spread out the numbers are.
func CalculateVariance(numbers []int) (*float64, error) {
	// Check if the input slice is nil or empty. If so, return an error
	if numbers == nil || len(numbers) == 0 {
		// Return error indicating an empty slice
		return nil, errors.New("empty numbers") 
	}

	// Convert the length of the slice to float64 for precise calculations
	n := float64(len(numbers))

	// sum is a helper function that applies a function (f) to each element of numbers 
	// and sums the results concurrently using goroutines
	sum := func(f func(int) float64) float64 {
		var s float64
		// Create a buffered channel to hold results of goroutines
		results := make(chan float64, len(numbers))

		// Iterate through the numbers, spawning a goroutine for each element
		for _, num := range numbers {
			go func(i int) {
				// Send the result of applying f(i) to the channel
				results <- f(i) 
			}(num)
		}

		// Wait for all goroutines to finish and accumulate the results
		for i := 0; i < len(numbers); i++ {
			// Sum up the results received from the channel
			s += <-results 
		}
		return s
	}

	// Calculate the mean (average) of the numbers by summing them and dividing by n
	mean := sum(func(i int) float64 {
		// Convert each number to float64
		return float64(i) 
	}) / n

	// Calculate the sum of squared differences from the mean, which is used in variance calculation
	sumRes := sum(func(i int) float64 {
		// For each number, subtract the mean, square the result, and sum them up
		return math.Pow(float64(i)-mean, 2)
	})

	// Variance is the sum of squared differences divided by the number of elements
	variance := sumRes / n
	
	// Return a pointer to the variance value and no error
	return &variance, nil
}