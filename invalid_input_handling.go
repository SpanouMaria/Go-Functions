// Calculates mean of a slice, handling errors for invalid input.

package main

import "errors"

// MeanInput is a type constraint that allows int, uint, or float64 types.
// It ensures that the input type of numbers is one of these types.
type MeanInput interface {
    int | uint | float64
}


// CalculateMean returns the mean of the numbers slice, or an error
// It computes the average of the numbers, or returns an error if input is invalid
func CalculateMean[T any](numbers []T) (*float64, error) {
    // Check if the input slice is nil or empty. Return an error if so
    if numbers == nil || len(numbers) == 0 {
        // Error indicating invalid input
        return nil, errors.New("empty numbers") 
    }

    // Initialize sum variable to accumulate the sum of numbers
    var sum float64
    
    // Iterate through the numbers and accumulate their sum
    for _, num := range numbers {
        // Convert each number to float64 before adding
        sum += float64(num) 
    }

    // Calculate the mean by dividing the sum by the number of elements
    mean := sum / float64(len(numbers))
    
    // Return a pointer to the mean value and no error
    return &mean, nil
}