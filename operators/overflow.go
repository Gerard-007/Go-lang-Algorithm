package operators

import (
	"fmt"
	"math"
)

/*
- Learning overflow
*/

func main() {
	var x uint8 = 255
	x++ // overflow, x is 0

	// a := int8(255 + 1)
	var b int8 = 127

	fmt.Printf("%d\n", b+1) // returns -128

	b = -128
	b--
	fmt.Println(b) // returns 127

	f := float32(math.MaxFloat32)
	f = f * 1.2
	fmt.Println(f) // f overflows to infinite

	// const i int8 = 300 // this is a compile time error
}
