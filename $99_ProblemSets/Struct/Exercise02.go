package main

import (
	"fmt"
	"math"
)

type Monomial struct {
	coeff float32
	exp   int
}

func inputMonomial(m *Monomial) {
	fmt.Print("Enter the coefficient and exponent of the monomial: ")
	fmt.Scan(&m.coeff, &m.exp)
}

func printMonomial(m Monomial) {
	if m.coeff != 0 {
		if m.exp == 0 {
			fmt.Println(m.coeff)
		} else {
			fmt.Printf("%.fx^%d\n", m.coeff, m.exp)
		}
	} else {
		fmt.Println(m.coeff)
	}
}

func addition(m1, m2 Monomial) Monomial {
	if m1.exp == m2.exp {
		return Monomial{m1.coeff + m2.coeff, m1.exp}
	}
	fmt.Println("Monomials are not of the same degree")
	return Monomial{}
}

func subtraction(m1, m2 Monomial) Monomial {
	if m1.exp == m2.exp {
		return Monomial{m1.coeff - m2.coeff, m1.exp}
	}
	fmt.Println("Monomials are not of the same degree")
	return Monomial{}
}

func multiplication(m1, m2 Monomial) Monomial {
	return Monomial{m1.coeff * m2.coeff, m1.exp + m2.exp}
}

func division(m1, m2 Monomial) Monomial {
	if m2.coeff != 0 {
		return Monomial{m1.coeff / m2.coeff, m1.exp - m2.exp}
	}
	fmt.Println("Division by zero")
	return Monomial{}
}

func firstGradeDeriative(m Monomial) Monomial {
	return Monomial{m.coeff * float32(m.exp), m.exp - 1}
}

func calculateValue(m Monomial, x float32) float64 {
	return float64(m.coeff) * math.Pow(float64(x), float64(m.exp))
}

func main() {
	var m1, m2 Monomial
	inputMonomial(&m1)
	inputMonomial(&m2)
	fmt.Println("Monomial m1: ")
	printMonomial(m1)
	fmt.Println("Monomial m2: ")
	printMonomial(m2)

	fmt.Print("Addition of two monomials is: ")
	printMonomial(addition(m1, m2))
	fmt.Print("Subtraction of two monomials is: ")
	printMonomial(subtraction(m1, m2))
	fmt.Print("Multiplication of two monomials is: ")
	printMonomial(multiplication(m1, m2))
	fmt.Print("Division of two monomials is: ")
	printMonomial(division(m1, m2))

	fmt.Print("First grade deriative of m1 is: ")
	printMonomial(firstGradeDeriative(m1))
	fmt.Print("Enter x: ")
	var x float32
	fmt.Scan(&x)
	fmt.Println("Value of m1 at", x, "is", calculateValue(m1, x))
}
