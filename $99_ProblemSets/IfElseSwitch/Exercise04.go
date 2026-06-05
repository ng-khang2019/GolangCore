package main

import "fmt"

func standing(grade float32) {
	if grade < 0 || grade > 10 {
		fmt.Println("Invalid grade")
		return
	}
	switch {
	case grade >= 9.0 && grade <= 10:
		fmt.Println("Excellent")
	case grade >= 8.0:
		fmt.Println("Very good")
	case grade >= 7.0:
		fmt.Println("Good")
	case grade >= 6.0:
		fmt.Println("Fair")
	case grade >= 5.0:
		fmt.Println("Pass")
	default:
		fmt.Println("Failed")
	}
}

func main() {
	standing(5.9)
}
