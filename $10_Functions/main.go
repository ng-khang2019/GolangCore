package main

import (
	"fmt"
)

func main() {
	fmt.Print("Sum of 1, 2 and 3 is ", add(1, 2, 3))
	fmt.Println()

	a, b := swap("Hello", "World") // Same as var a, b string = "Hello", "World"
	fmt.Println(a, b)

	fmt.Println(split(add(5, 8, 4)))

	// A defer statement defers the execution of a
	// function until the surrounding function returns.
	defer fmt.Print("Done!\n")

	/*
		Deferred function calls are pushed onto a stack. When a
		function returns, its deferred calls are executed in
		last-in-first-out order.
	*/
	fmt.Println("Counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	array := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("Sum of 1,2,3", sumArr(1, 2, 3))
	fmt.Println("Sum of elements:", sumArr(array...))

	// Closure function call
	count := counter()

	fmt.Println("Counter:")
	fmt.Println(count())
	fmt.Println(count())
	fmt.Println(count())

	// More complex closure function call
	posSum, negSum := sum(), sum()
	for i := 0; i < 10; i++ {
		fmt.Println(posSum(i*2), negSum(i*(-1)))
	}
}

/*
When two or more consecutive named function parameters share a type,
you can omit the type from all but the last. The example below
shortened (x int, y int, z int) to (x, y, z int)
*/
func add(x, y, z int) int {
	return x + y + z
}

/*
A function can also returns multiple values
*/
func swap(x, y string) (string, string) {
	return y, x
}

/*
A return statement without arguments returns the named return values.
*/
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // This will return the named return values which are x and y
}

// Variadic function
func sumArr(numbers ...int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// This function returns a closure.
// The returned anonymous function retains direct reference to the local variables
// of its parent scope (state preservation). Any changes made to these variables
// will persist across multiple function calls.
func counter() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func sum() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
