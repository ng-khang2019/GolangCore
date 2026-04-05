package main

import "fmt"

/* In Golang, there is only for loop, Go does not have while or do-while
 * For loop in Go is quite flexible and could be used for anything
 */
func main() {
	// Most basic for type with a single condition - while loop
	fmt.Print("i: ")
	i := 1
	for i <= 3 {
		fmt.Print(i, " ")
		i++
	}
	fmt.Println()

	// Classic type with initial/condition/post for loop
	fmt.Print("j: ")
	for j := 1; j <= 3; j++ {
		fmt.Print(j, " ")
	}
	fmt.Println()

	// For loop with "continue"
	fmt.Print("m: ")
	for m := 1; m <= 5; m++ {
		if m == 3 {
			continue
		}
		fmt.Print(m, " ")
	}
	fmt.Println()

	// Infinite loop with break
	fmt.Print("k: ")
	for {
		fmt.Print("k ")
		break
	}
	fmt.Println()

	/* For loop with post-function inside the loop
	 * Notice that this form is exactly like the
	 * first example - a while loop but with the variable n
	 * declared inside for statement (local variable)
	 */
	fmt.Print("n: ")
	for n := 1; n < 10; {
		fmt.Print(n, " ")
		n *= 2
	}
	fmt.Println()
}
