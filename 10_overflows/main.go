package main

import (
	"fmt"
	"math"
)

func main() {

	// x instead of 266 jumped to 0 because of overflow (runtime)
	var x uint8 = 255
	x++
	fmt.Println(x)

	// a := int8(255 + 1) compile time error

	var b int8 = 127
	fmt.Printf("%d\n", b+1) // -128

	b = -128
	b--
	fmt.Printf("%d\n", b) // 127

	// FLOAT OVERFLOW
	f := float32(math.MaxFloat32)
	fmt.Println(f)

	f *= 1.2
	fmt.Println(f) // overflow to +Inf

	// for big numbers use math/big
}
