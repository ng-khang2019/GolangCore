package main

import (
	"fmt"
)

// Define a struct
type person struct {
	name string
	age  int
}

func createPerson(name string) *person {
	p := person{name: name}
	p.age = 17
	return &p
}

// Go also supports methods on structs
func (p person) bornYear() int {
	return 2026 - p.age
}

func main() {

	// Declare a person
	var person1 person = person{"Khang", 24}
	fmt.Println(person1)

	// Short declaration
	person2 := person{"Hoang", 15}
	fmt.Println(person2)

	// Quick define and declare
	dog := struct {
		name            string
		characteristics string
	}{
		"Rex", "Cute",
	}
	fmt.Println(dog)

	// Accessing struct fields
	person1.name = "Nguyen"
	person1.age = 33
	fmt.Println(person1.name, person1.age)

	// Create a person by function
	person3 := createPerson("Nguyen")
	fmt.Println(person3)

	// Function call on struct
	fmt.Println(person3.bornYear())

	// Different types to declare a struct
	type Vertex struct {
		X, Y int
	}
	var (
		v1 = Vertex{1, 2}  // has type Vertex
		v2 = Vertex{X: 1}  // Y:0 is implicit
		v3 = Vertex{}      // X:0 and Y:0
		p  = &Vertex{1, 2} // has type *Vertex
	)
	fmt.Println(v1, v2, v3, p)

	// Slice of structs
	numbers := []struct {
		number int
		name   string
	}{
		{1, "one"},
		{2, "two"},
		{3, "three"},
	}
	fmt.Println(numbers)

}
