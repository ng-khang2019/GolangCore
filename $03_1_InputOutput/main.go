package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var age int
	var category string
	var id int

	// Use bufio Scanner or Reader to read input contains white space from user
	// Using bufio.NewScanner() is more efficient than using bufio.NewReader() since
	// it can handle white space from the input
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter your name: ")
	scanner.Scan()
	name := scanner.Text()

	// If you want to use bufio.NewReader(), make sure to clear the white space
	// character '\n' after reading the input
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your address:: ")
	address, _ := reader.ReadString('\n') // Read until '\n' and ignore the error
	address = strings.TrimSpace(address)  // Remove '\n' from input

	/* Use pointer to tell go where to store value
	 * Since what is passed in the function is not the
	 * actual variable age, but it's copy. Using pointer will pass
	 * the address of the actual variable to the function
	 */
	fmt.Print("Enter age: ")
	fmt.Scan(&age)

	fmt.Print("Enter ID and Category: ")
	fmt.Scanf("%d-%s", &id, &category)

	fmt.Println("You are", name)
	fmt.Println("Your address is", address)
	fmt.Printf("You are %d years old\n", age)
	fmt.Printf("Id: %d, Category: %s\n", id, category)

}
