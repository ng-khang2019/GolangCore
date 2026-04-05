package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Print("My favorite number is ", rand.Intn(10), ". Yolo!\n")
	fmt.Printf("Current time is: %s\n", time.Now().Local().Format("2006-01-02"))
	fmt.Printf("Pi constant has the value of %g\n", math.Pi)
	fmt.Printf("Square root of 7 is %.2f\n", math.Sqrt(7))
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("macOS.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("Your current OS is %s\n", os)
	}

}
