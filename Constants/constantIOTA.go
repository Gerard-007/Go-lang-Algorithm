package main

import "fmt"

/*
- IOTA
*/

func main() {
	const (
		c1 = iota // this becomes 0
		c2 = iota // this becomes 1
		c3 = iota // this becomes 2
	)
	fmt.Println(c1, c2, c3)

	const (
		c11 = iota // this becomes 0
		c22 = iota // this becomes 1
		c33 = iota // this becomes 2
	)
	fmt.Println(c11, c22, c33)

	const (
		North = iota // this becomes 0
		East         // this becomes 1
		South        // this becomes 2
		West         // this becomes 3
	)
	fmt.Println(West)

	const (
		a = iota * 2 // 0
		b            // 2
		c            // 4
	)

	const (
		x = -(iota + 2) // -2
		_               // -3
		y               // -4
		z               // -5
	)
	fmt.Print(x, y, z)
}
