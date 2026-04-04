package main

import (
	"fmt"
)

func main() {
	var a, b = 10, 3
	fmt.Printf("Given a = %d, b = %d\n", a, b)
	fmt.Println("Arithmetic Operators:")
	fmt.Println("a + b =", a+b)
	fmt.Println("a - b =", a-b)
	fmt.Println("a * b =", a*b)
	fmt.Println("a / b =", a/b)
	fmt.Println("a / b =", a/b)
	fmt.Println("a / b =", float32(a)/float32(b))
	fmt.Println("a % b =", a%b)
	/* Unlike other languages, Golang only supports postfix increment/decreasement
	 * and it is a statement, not an expression.
	 */
	a++
	fmt.Println("a++ = ", a)
	a--
	fmt.Println("a-- = ", a)

	fmt.Printf("--------------------\n")
	fmt.Println("Relational Operators:")
	fmt.Println("Is a == b?", a == b)
	fmt.Println("Is a != b?", a != b)
	fmt.Println("Is a > b?")
	fmt.Println("Is a - 5 >= b?", a-7 >= b)
	fmt.Println("Is a < b?", a < b)
	fmt.Println("Is a - 6 <= b?", a-8 <= b)

	fmt.Printf("--------------------\n")
	fmt.Println("Logical Operators:")
	fmt.Println("Is a == b && b == 3?", a == b && b == 3)
	fmt.Println("Is a > b || b + 6 =a?", a > b || b+6 == a)
	fmt.Println("Is !(a == b) ?", !(a == b))

	fmt.Printf("--------------------\n")
	fmt.Println("Bitwise Operators:")
	fmt.Println("Given a = 10, b = 3")
	fmt.Printf("With a and b in binary are %08b and %08b\n", a, b)
	fmt.Printf("AND: a & b = %d or %08b in binary\n", a&b, a&b)
	fmt.Printf("OR: a | b = %d or %08b in binary\n", a|b, a|b)
	fmt.Printf("XOR: a ^ b = %d or %08b in binary\n", a^b, a^b)
	fmt.Printf("Left shift: b << 1 = %d or %08b in binary\n", b<<1, b<<1)
	fmt.Printf("Right shift: b >> 1 = %d or %08b in binary\n", b>>1, b>>1)
	fmt.Printf("Bit clear: a &^ b = %d or %08b in binary\n", a&^b, a&^b)

	fmt.Printf("--------------------\n")
	fmt.Println("Assignment Operators:")
	a, b = 8, 3
	var c = a + b
	fmt.Println("Given a = 8, b = 3")
	fmt.Println("c = a + b =", c)
	a += b
	fmt.Println("a += b is", a)
	a -= b
	fmt.Println("a -= b is is", a)
	a *= b
	fmt.Println("a *= b is", a)
	a /= b
	fmt.Println("a /= b is", a)
	a %= b
	fmt.Println("a %= b is", a)
	a <<= b
	fmt.Println("a <<= b is", a)
	a &= b
	fmt.Println("a &= b is", a)

	var val = 42
	ptr := &val
	fmt.Printf("--------------------\n")
	fmt.Println("Miscellaneous Operators:")
	fmt.Println("Value of val:", val)
	fmt.Println("Val memory address:", ptr, "- By declaring \"ptr = &val\"")
	// Use * to dereference a pointer - get the value it points to
	fmt.Println("Value pointed by ptr:", *ptr, "- Using \"*ptr\"")
	// Reassigning a pointer to a new value by using *
	*ptr = 100
	fmt.Println("val new value:", val, "- By doing \"*ptr = 100\"")
}
