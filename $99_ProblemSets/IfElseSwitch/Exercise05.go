package main

import "fmt"

func sort(a, b, c int) {
	if a > b {
		a, b = b, a
	}
	if a > c {
		a, c = c, a
	}
	if b > c {
		b, c = c, b
	}
	fmt.Println(a, b, c)
}

func main() {
	sort(7, 9, 8)
}
