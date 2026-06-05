package main

import (
	"fmt"
	"math"
)

const EPS float64 = 1e-9

func equalFloat(a, b float64) bool {
	return math.Abs(a-b) < EPS
}

func validTriangle(a, b, c float64) bool {
	if (a+b >= c) && (a+c >= b) && (b+c >= a) {
		return true
	}
	return false
}

func rightTriangle(a, b, c float64) bool {
	return equalFloat(a*a+b*b, c*c) || equalFloat(b*b+c*c, a*a) || equalFloat(a*a+c*c, b*b)
}

func isoscelesTriangle(a, b, c float64) bool {
	return equalFloat(a, b) || equalFloat(a, c) || equalFloat(b, c)
}

func equilateralTriangle(a, b, c float64) bool {
	return equalFloat(a, b) && equalFloat(b, c)
}

func triangleType(a, b, c float64) string {
	if !validTriangle(a, b, c) {
		return "Not a triangle"
	}
	if rightTriangle(a, b, c) && isoscelesTriangle(a, b, c) {
		return "Right isosceles triangle"
	} else if equilateralTriangle(a, b, c) {
		return "Equilateral triangle"
	} else if isoscelesTriangle(a, b, c) {
		return "Isosceles triangle"
	} else if rightTriangle(a, b, c) {
		return "Right triangle"
	}
	return "Scalene triangle"
}

func main() {
	fmt.Println(triangleType(3, 4, 5))
}
