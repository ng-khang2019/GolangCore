package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	/*
		Slices are like arrays, but they can be of different lengths.
		Slices are references to underlying arrays. They are made up by
		three parts: the pointer to the underlying array, the length and its capacity
	*/

	// Declaring a slice (empty slice)
	var sliceA []int
	fmt.Println(sliceA)

	// Declare and initialize a slice
	var sliceB = []int{1, 2, 3, 4, 5}
	fmt.Println(sliceB)

	// Quick declaration with values
	sliceC := []int{1, 2, 3, 4, 5}
	fmt.Println(sliceC)

	// Declaring with make() function
	sliceD := make([]int, 5)     // Declare with a specified length
	sliceE := make([]int, 4, 10) // Declare with a specified length and capacity

	// len() and cap() functions
	fmt.Println("sliceD's length:", len(sliceD))
	fmt.Println("sliceC's length and capacity:", len(sliceE), cap(sliceE))

	// Accessing and modifying elements
	sliceE[0] = 1
	sliceE[1] = 2
	sliceE[2] = 3
	sliceE[3] = 4
	fmt.Println("sliceE:", sliceE)

	// Adding elements with append() function
	// Note: append() function returns a new slice, not the original slice
	sliceE = append(sliceE, 5)
	sliceE = append(sliceE, 6, 7, 8)
	fmt.Println("sliceE:", sliceE)

	// Creating a slice from an array
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println("Array of primes:", primes)
	var slicePrimes = primes[1:4]
	fmt.Println("Slice of primes:", slicePrimes)

	// Copying a slice
	mySlice := []int{1, 2, 3, 4, 5}
	copySlice := make([]int, len(mySlice))
	copy(copySlice, mySlice)
	fmt.Println("Original slice:", mySlice)
	fmt.Println("Copied slice:", copySlice)

	// Slicing up a slice
	fmt.Println("Slicing up a slice:")
	s := []int{2, 3, 5, 7, 11, 13, 17, 19}
	fmt.Println(s)
	s = s[1:4]
	fmt.Println(s)
	s = s[2:]
	fmt.Println(s)
	s = s[:3]
	fmt.Println(s)

	// Default values: 0 for lower bound and length of the array for upper bound
	var array = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println(array)
	// These are the same:
	a := array[:10]
	b := array[0:]
	c := array[0:10]
	d := array[:]
	fmt.Println(a, b, c, d)

	// Looping through a slice using range
	// For more range function variants, see the arrays lesson
	fmt.Println("Looping through a slice using range:")
	for index, value := range []int{3, 5, 7, 9, 11} {
		fmt.Print(index, ":", value, "|")
	}
	fmt.Println()

	// Extending and shrinking a slice
	s1 := []int{2, 3, 5, 7, 11, 13}
	printSlice(s1)
	// Slice the slice to give it zero length.
	s1 = s1[:0]
	printSlice(s1)
	// Extend its length.
	s1 = s1[:4]
	printSlice(s1)
	// Drop its first two values.
	s1 = s1[2:]
	printSlice(s1)
	s1 = s1[0:]
	printSlice(s1)

	// Memory leak issue
	// Method 1: Copying a slice helps to avoid unnecessary memory allocation rather than slicing and reassignment
	// New slice will be created and the old slice's underlying array will be processed by Garbage Collector
	// If you reassign the slice with this s3 = s3[:6], then the underlying array will still keep
	// all the value from index 6
	s3 := []int{1, 5, 2, 7, 4, 9, 5, 4, 7, 3, 4, 1, 8}
	copyS3 := make([]int, 6)
	copy(copyS3, s3[:6])
	fmt.Println(copyS3)

	// Method 2: Copy a slice by cloning it using the "slice" package clone function
	s4 := []int{1, 5, 2, 7, 4, 9, 5, 4, 7, 3, 4, 1, 8}
	copyS4 := make([]int, 6)
	copyS4 = slices.Clone(s4[:6]) // Must reassign
	fmt.Println(copyS4)

	// Slices of structs
	// Method 3: Reassign unused elements to default value(0) before slicing and reassigning
	s5 := []Object{{"Apple"}, {"Banana"}, {"Orange"}}
	s5[len(s5)-1] = Object{} // Reassign the last element with empty Object
	s5 = s5[:len(s5)-1]      // Now slice the rest
	fmt.Println(s5)

	// Slice of pointers
	// Method 4: Reassign unused elements to nil before slicing and reassigning
	s6 := []*Object{&Object{"Banana"}, &Object{"Grapes"}, &Object{"Durian"}}
	s6[len(s6)-1] = nil
	s6 = s6[:len(s6)-1]
	fmt.Println(s6)

	// Method 5: Using slices.delete(), applicable for both slices of structs and pointers
	// Note: Must reassign the slice after deletion
	sliceStruct := []Object{{"Bob"}, {"Alex"}, {"Joe"}, {"John"}}
	slicePointer := []*Object{&Object{"Bob"}, &Object{"Alex"}, &Object{"Joe"}, &Object{"John"}}

	sliceStruct = slices.Delete(sliceStruct, 1, 3)
	slicePointer = slices.Delete(slicePointer, 1, 3)

	fmt.Println(sliceStruct)
	fmt.Println(slicePointer)

	// Declare a 2D slice
	twoD := make([][]int, 3)
	for i := range 3 {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := range innerLen {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

type Object struct {
	Name string
}
