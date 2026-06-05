package main

import "fmt"

func maxOf2(a, b int) int {
	switch {
	case a > b:
		return a
	default:
		return b
	}
}

func main() {

	var a, b int
	fmt.Println("Enter two numbers: ")
	fmt.Scan(&a, &b)
	fmt.Println("Max of", a, "and", b, "is", maxOf2(a, b))
}
