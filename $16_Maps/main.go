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
	m5 := map[string]int{"Hailey": 38, "Brian": 49, "Jessica": 50, "Kevin": 51, "Lisa": 52}

	fmt.Println(m1)
	fmt.Println(m2)
	fmt.Println(m3)
	fmt.Println(m4)
	fmt.Println(m5)

	// Map literal
	type Vertex struct {
		Lat, Long float64
	}

	var m = map[string]Vertex{
		"Bell Labs": Vertex{
			40.68433, -74.39967,
		},
		"Google": {37.42202, -122.08408}, // Omit the type name "Vertex" of the value
	}
	fmt.Println(m)

	// Adding, updating and accessing values
	myMap := make(map[string]int)

	myMap["Answer"] = 42
	fmt.Println("The value:", myMap["Answer"])

	myMap["Answer"] = 48
	fmt.Println("The value:", myMap["Answer"])

	myMap["Answer 1"] = 100
	fmt.Println("The value:", myMap["Answer 1"])

	myMap["Question"] = 60
	fmt.Println("The value:", myMap["Question"])

	// Check if a key exists in a map with a two-value assignment
	val, check := myMap["Question"]
	fmt.Println("Value:", val, "Presents?:", check)

	_, exists := myMap["Answer"]
	fmt.Println("Key exists:", exists)

	// Loop through a map with range
	fmt.Println("Looping through a map with range:")
	for key, value := range m5 {
		fmt.Print(key, ":", value, " ")
	}
	fmt.Println()
	for key := range m5 {
		fmt.Print("Key:", key, " ")
	}

	// Compare maps
	n1 := map[string]int{"foo": 1, "bar": 2}
	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n1, n2) {
		fmt.Println("n == n2")
	}

	// Delete a key-value pair
	myMap1 := map[string]int{"Banana": 12, "Durian": 40, "Apple": 30, "Grapes": 20}
	delete(myMap1, "Apple")
	fmt.Println(myMap1)
	fmt.Println("My fruit map length is", len(myMap1))

	// Clear a whole map
	clear(myMap1)
	myMap1["Longan"] = 100
	fmt.Println(myMap1)
}
