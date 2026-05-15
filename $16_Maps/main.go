package main

import (
	"fmt"
	"maps"
)

func main() {
	// Declare and later assign values
	var m1 map[string]int
	m1 = map[string]int{
		"Alice": 30, "Bob": 40,
	}

	// Declare and initialize in one line
	var m2 = map[string]int{"Tron": 32, "Joey": 20}

	// Short declaration with make
	m3 := make(map[string]int)
	m3["Logan"] = 27
	m3["Vanessa"] = 28

	// Short declaration with make and capacity
	m4 := make(map[string]int, 10)
	m4["Walt"] = 37
	m4["Newtron"] = 53

	// Short declaration with literal
	m5 := map[string]int{"Hailey": 38, "Brian": 49}

	fmt.Println(m1)
	fmt.Println(m2)
	fmt.Println(m3)
	fmt.Println(m4)
	fmt.Println(m5)

	// Compare maps
	n1 := map[string]int{"foo": 1, "bar": 2}
	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n1, n2) {
		fmt.Println("n == n2")
	}

	// Delete a key-value pair
	myMap := map[string]int{"Banana": 12, "Durian": 40, "Apple": 30, "Grapes": 20}
	delete(myMap, "Apple")
	fmt.Println(myMap)
	fmt.Println("My fruit map length is", len(myMap))
	clear(myMap)
	myMap["Longan"] = 100
	fmt.Println(myMap)
}
