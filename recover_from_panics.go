// Adds persons with age validation, handles panic, returns operation result.

package main

import (
	"fmt"
)


const (
	SuccessMsg string = "Operation successful!"  // Message for successful operation
	PanicMsg   string = "Operation panicked!"    // Message for panic situations
)


// Person defines a structure with Name and Age.
type Person struct {
	Name string  // Person's name
	Age  int     // Person's age
}


// Persons stores a map of Person objects, keyed by their name
type Persons struct {
	m map[string]Person  // A map where keys are names and values are Person objects
}


// NewPersons initializes and returns a new Persons object
func NewPersons() *Persons {
	return &Persons{
		// Initializes an empty map for storing Persons
		m: make(map[string]Person),  
	}
}


// PanickingAdd adds a new Person to the Persons receiver and
// panics if the age of the person is invalid (less than 0 or greater than 100)
func (p *Persons) PanickingAdd(np Person) {
	// Check for valid age (between 0 and 100)
	if np.Age < 0 || np.Age > 100 {
		// Panic with a message if the age is invalid
		panic(fmt.Sprintf("add:invalid age %d for name %s", np.Age, np.Name))
	}
	
	// Add the person to the map
	p.m[np.Name] = np
}


// PopulateData adds a slice of Person objects to the Persons type
// If any invalid ages are encountered, the function handles the panic and returns a result message
func PopulateData(data []Person) (result string) {
	result = SuccessMsg  // Default result message is success
	// Defer block to recover from a panic and set the result message to PanicMsg
	defer func() {
		if r := recover(); r != nil {
			// If a panic occurred, set the result to panic message
			result = PanicMsg  
		}
	}()
	
	// Create a new Persons object to store the data
	p := NewPersons()
	
	// Loop through the data and attempt to add each Person to the Persons object
	for _, np := range data {
		// Add each person, may panic if invalid age
		p.PanickingAdd(np)  
	}
	
	// Return the result (SuccessMsg or PanicMsg)
	return  
}