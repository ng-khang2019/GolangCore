package main

import (
	"fmt"
	"math"
)

func linearEquationEx6(a, b int) {
	if a == 0 {
		fmt.Println("Linear equation has no solution")
	} else {
		fmt.Println("x = ", float32(-b)/float32(a))
	}
}

func quadraticEquation(a, b, c int) {
	if a == 0 {
		linearEquationEx6(b, c)
	} else {
		var discriminant = float64(b*b - 4*a*c)
		if discriminant < 0 {
			fmt.Println("Quadratic equation has no real solutions")
		} else if discriminant == 0 {
			fmt.Println("Quadratic equation has one real solution: x = ", float64(-b)/(2*float64(a)))
		} else {
			fmt.Println("Quadratic equation has two real solutions: x1 = ",
				(float64(-b)+math.Sqrt(discriminant))/(2*float64(a)),
				" x2 = ", (float64(-b)-math.Sqrt(discriminant))/(2*float64(a)))
		}
	}
}

func main() {
	quadraticEquation(1, 3, 1)
}
