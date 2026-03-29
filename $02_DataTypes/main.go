package _02_DataTypes

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

/*
	The int, uint, and uintptr types are usually 32 bits wide on 32-bit systems
	and 64 bits wide on 64-bit systems. When you need an integer value, you should
	use int unless you have a specific reason to use a sized or unsigned integer type
*/

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}
