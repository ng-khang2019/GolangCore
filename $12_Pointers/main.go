package main

import "fmt"

func main() {
	a, b := 10, 15
	fmt.Println("Value of a and b: ", a, b)

	// Declaring a pointer
	var ptr1 *int = &a
	ptr2 := &b
	fmt.Println("Pointer value of a and b: ", ptr1, "|", ptr2) // Output reference of the pointers

	// Accessing the value pointed by the pointers
	fmt.Println("Accessing the value pointed by the pointers:")
	fmt.Println(*ptr1, *ptr2)

	// De-referencing the pointers
	*ptr1 = 20 // a = 20 as ptr1 stores the address of a
	*ptr2 = 30 // b = 30 as ptr2 stores the address of b
	fmt.Println("De-referencing the pointers:")
	fmt.Println(a, b)

}
