package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Print("My favorite number is ", rand.Intn(10), ". Yolo!\n")
	fmt.Printf("Current time is: %s\n", time.Now().Local().Format("2006-01-02"))
	fmt.Printf("Pi constant has the value of %g\n", math.Pi)
	fmt.Printf("Square root of 7 is %.2f", math.Sqrt(7))

}
