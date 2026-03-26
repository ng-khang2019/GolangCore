package main

import (
	"fmt"
)

func main() {

	s := "gopher"
	var x string = "Khang"
	fmt.Println("Hi " + x)
	fmt.Printf("Hello and welcome, %s!\n", s)

	for i := 1; i <= 5; i++ {
		fmt.Println("i = ", 100/i)
	}
}
