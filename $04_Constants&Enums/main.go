package main

import (
	"fmt"
)

// Declare constants without type (Global Scope)
const Pi = 3.14159
const ApplicationName = "Golang App"

// Grouped constants
const (
	StatusActive  = 1
	StatusPending = 2
	StatusClosed  = 3
)

// Using iota for auto-increment
const (
	CategoryTech  = iota //0
	CategoryArt          // 1
	CategoryMusic        // 2
)

func main() {
	// Declare constants without a type inside a function (Local Scope)
	const LocalLimit = 1000
	const WelcomeMessage = "Welcome to Golang"
	// Declare a constant with type
	const SecondPerMinute int = 60

	fmt.Println("Message:", WelcomeMessage)
	fmt.Println("App Name:", ApplicationName)
	fmt.Println("Pi Value:", Pi)
	fmt.Println("Category Art ID:", CategoryArt)
	// A constant can hold no type until it is used or through a conversion
	fmt.Println("Local Limit:", int64(LocalLimit))

	// Create enums with custom Type, const and iota
	type Day int
	const (
		/* The first line can be declared as '_ = iota' without a
		 * type. Its type will be inferred as "untyped integer".
		 * Therefore, the following constants will be of type int.
		 */
		_         Day = iota // Skip 0
		Monday               // 1
		Tuesday              // 2
		Wednesday            // 3
		Thursday             // 4
		Friday               // 5
		Saturday             // 6
		Sunday               // 7
	)

	fmt.Println("First day of the week has the value of", Monday)

}
