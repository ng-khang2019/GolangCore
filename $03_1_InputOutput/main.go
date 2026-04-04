package main

import "fmt"

func main() {
	var age int
	var category string
	var id int
	/* Use pointer to tell go where to store value
	 * Since what is passed in the function is not the
	 * actual variable age, but it's copy. Using pointer will pass
	 * the address of the actual variable to the function
	 */
	fmt.Print("Enter age: ")
	fmt.Scan(&age)

	fmt.Print("Enter ID and Category: ")
	fmt.Scanf("%d-%s", &id, &category)

	fmt.Printf("You are %d years old\n", age)
	fmt.Printf("Id: %d, Category: %s\n", id, category)

}
