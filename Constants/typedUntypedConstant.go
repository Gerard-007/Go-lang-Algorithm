package main

import "fmt"

/*
- Typed and untyped constants
*/

func main() {
	const a float64 = 5.1 // typed constant

	const b = 6.7 // untyped constant

	const c float64 = a * b
	const str = "Hello " + "go!"

	const d = 5 > 10
	fmt.Println(d)

	//const x int = 5
	//const y float64 = 2.2 * x // This gives error due to x is of type int

	const x = 5
	const y = 2.2 * x // This doesnt gives error due to x i untyped
	fmt.Printf("%T\n", y)

	var i int = x     // x changes to int
	var j float64 = x // j changes to float64(x)
	var k byte = x
	fmt.Println(i, j, k)

	const r = 5 // This is int by default
	var rr = r  // rr becomes int by default due to r is int
	fmt.Printf("%T\n", rr)

}
