// Calculates mean of integers slice, returns error for empty input.

package main

import "errors"

// CalculateMean returns the mean (average) of the numbers slice, or an error if the slice is empty or nil
func CalculateMean(numbers []int) (*float64, error) {
    
	// Check if the slice is nil or empty and return an error if true
	if numbers == nil || len(numbers) == 0 {
		return nil, errors.New("empty numbers") // Error message indicating that the slice is empty
	}
    
	// Initialize a variable to accumulate the sum of the numbers
	var sum int
	
	// Loop through each number in the slice and add it to the sum
	for _, num := range numbers {
		sum += num
	}

	// Calculate the mean by dividing the sum by the length of the slice
	mean := float64(sum) / float64(len(numbers))
	
	// Return a pointer to the mean value and nil for the error
	return &mean, nil
}