package operators

import "fmt"

/*
 */

func main() {
	// Arithmetic Operators
	a, b := 4, 2
	c := (a + b) / (a - b) * 2
	fmt.Println(c) // returns 6

	// Assignment Operators
	x, y := 2, 3

	// Increment x
	x += y
	fmt.Println(x, y) // returns 5, 3

	// decrement
	y -= 2
	fmt.Println(y)

	// multiplication
	y *= 10
	fmt.Println(y)

	// division
	y /= 5
	fmt.Println(y)

	// modulus
	x %= 3
	fmt.Println(x)

	// Increment Statement
	z := 1
	z += 1
	z++
	fmt.Println(z)

	// decrement Statement
	v := 1
	v -= 1
	v--
	fmt.Println(v)
}
