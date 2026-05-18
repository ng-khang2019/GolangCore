package main

import "fmt"

func main() {
	// Declare an array
	var arr1 [3]int

	// Accessing and adding elements
	arr1[0] = 10
	arr1[1] = 20
	arr1[2] = 30
	fmt.Println("Adding elements to array1:")
	fmt.Println(arr1)

	// Declare and initialize an array
	arr2 := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println("\nInitializing array2 with values:")
	fmt.Print(arr2, "\n")

	// Length of the array
	fmt.Println("\nLength of array1 and array2:", len(arr1), len(arr2))
	fmt.Println("Array2 index of 2 -> end:", arr2[2:])
	fmt.Println("Array2 start -> 2:", arr2[:3])
	fmt.Println("Array2 index of 1 -> 4 :", arr2[1:5])

	// Advanced techniques
	// Have the compiler count the number of elements
	arrA := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println("\nArrayA: ", arrA)
	// Specify the elements
	arrB := [...]int{1: 10, 2: 20, 5: 50, 6: 60, 70, 80}
	fmt.Println("ArrayB: ", arrB)

	// Looping through an array using range
	// ignore index, access value only
	fmt.Println("\nLooping through arrayB using range:")
	fmt.Println("Ignoring indexes, accessing values only:")
	for _, value := range arrB {
		fmt.Print(value, " ")
	}
	// access index only
	fmt.Println("\nAccessing indexes only:")
	for index := range arrB {
		fmt.Print(index, " ")
	}
	// access both index and value
	fmt.Println("\nAccessing both index and value:")
	for index, value := range arrB {
		fmt.Print(index, ":", value, "|")
	}
	fmt.Println()

	// Declare a 2D array
	var arr3 [2][3]int
	for i := 0; i < len(arr3); i++ {
		for j := 0; j < len(arr3[i]); j++ {
			arr3[i][j] = i + j
		}
	}
	fmt.Println("\n2D array3:", arr3)

	// Declare and initialize a 2D array
	var arr4 [3][2]int = [3][2]int{
		{1, 2}, {3, 4}, {5, 6},
	}
	fmt.Println("\nInitializing 2D array4 with values:", arr4)

	// Or with quick declaration
	arr5 := [2][3]int{
		{1, 2, 3}, {4, 5, 6},
	}
	fmt.Println("\nInitializing 2D array5 with values:", arr5)
}
