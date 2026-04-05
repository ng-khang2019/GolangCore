package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		Go only runs the selected case, not all the cases that follow.
		In effect, the break statement is provided automatically in Go.
	*/
	// Basic switch
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// Switch with no condition is like an if-else statement
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}

	// Use fallthrough to proceed to the next case
	num := 2
	switch num {
	case 1:
		fmt.Println("Case 1")
	case 2:
		fmt.Println("Case 2")
		fallthrough // Forces execution of the next case
	case 3:
		fmt.Println("Case 3") // This prints even though num != 3
	default:
		fmt.Println("Default")
	}

	// Multiple cases with the same expression
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("It's a weekday!")
	}

	// Include expressions in cases
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	// A type switch compares type instead of value
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		case string:
			fmt.Println("I'm a string")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(13)
	whatAmI("Hello")
	whatAmI(13.37)
}
