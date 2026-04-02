package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	toBe     bool       = false
	maxInt   uint64     = 1<<64 - 1
	z        complex128 = cmplx.Sqrt(-5 + 12i)
	maxInt64 int64      = math.MaxInt64
	minInt64 int64      = math.MinInt64
)

/*
	The int, uint, and uintptr types are usually 32 bits wide on 32-bit systems
	and 64 bits wide on 64-bit systems. When you need an integer value, you should
	use int unless you have a specific reason to use a sized or unsigned integer type
*/

func main() {
	fmt.Printf("Type: %T - Value: %v\n", toBe, toBe)
	fmt.Printf("Type: %T - Value: %v\n", maxInt, maxInt)
	fmt.Printf("Type: %T - Value: %v\n", z, z)
	fmt.Printf("Type: %T - Max Value: %v\n", maxInt64, maxInt64)
	fmt.Printf("Type: %T - Max Value: %v\n", minInt64, minInt64)
}
