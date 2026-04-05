package main

import (
	"fmt"
)

/*
In Golang, the "var" statement can declare one or more variables; the var keyword comes first
then variable name and its type. A var statement can also be at the package or function level
*/

/*
IMPORTANT: var and const can be used anywhere in the program, in or outside a function,
but ":=" can only be used inside a function
*/

func main() {
	// Declare without initializing
	var i int                  // Default value for integer is 0
	var cpp, python, java bool // boolean variables are false as default
	var job string             // string default literal is ""
	fmt.Println("Declare without initializing:")
	fmt.Println(i, cpp, python, java)
	fmt.Println(job)

	// Declare and initialization or complete declaration
	var a string = "Hello"
	var b, c int = 1, 2
	fmt.Println(a, b, c)

	// Declare without its type. Go will automatically infer the type
	var d = "World"
	var e = 3.14
	fmt.Println(d)
	fmt.Println("Pi value is ", e)

	// := operator is used for shorthand declaration. It is commonly used inside a function
	fruit := "Banana"
	prime1, prime2, prime3 := 2, 3, 5
	fmt.Println(fruit)
	fmt.Println(prime1, prime2, prime3)
	fmt.Println(printApple())

	// Declare, initialize and pass values to a function
	x, y := swap(10, 5)
	fmt.Println(x, y)

	var1, var2 := 1, 2
	// var1, var2 := 7, 8 - Reassign with this will cause error
	var1, var2, var3 := 7, 8, 9 // But this works
	fmt.Printf("var1: %d\nvar2: %d\nvar3: %d\n ", var1, var2, var3)
}

func swap(x, y int) (int, int) {
	return y, x
}

func printApple() string {
	apple := "Apple" // It is similar to var apple string = "Apple"
	return apple
}
