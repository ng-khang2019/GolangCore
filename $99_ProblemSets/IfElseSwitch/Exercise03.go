package main

import "fmt"

func linearEquation(a, b int) {
	if a == 0 {
		fmt.Println("Linear equation has no solution")
	} else {
		fmt.Println("x = ", float32(-b)/float32(a))
	}
}

func main() {
	linearEquation(3, 2)
}
