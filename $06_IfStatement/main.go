package main

import "fmt"

func main() {
	/*
		Note that in Golang, we don't need parentheses around the condition(s)
		But curly braces are still required
	*/
	var number int
	fmt.Println("Enter a number: ")
	fmt.Scan(&number)
	if number%2 == 0 {
		fmt.Println("The number is even")
	} else {
		fmt.Println("The number is odd")
	}

	// If statement without an else
	if number%4 == 0 {
		fmt.Println("The number is divisible by 4")
	}

	/*
		A statement can precede conditionals; any variables declared in this statement
		are available in the current and all later branches
	*/
	fmt.Println("Given a number 4")
	if num := 4; num < 0 {
		fmt.Println("The number is negative")
	} else if num >= 0 && num < 10 {
		fmt.Println("The number has 1 digit")
	} else {
		fmt.Println("The number has more than 1 digit")
	}
	/*
		One more thing needs to be noted is that there is no ternary operator in Go.
		You have use if statement for even basic conditions
	*/
}
