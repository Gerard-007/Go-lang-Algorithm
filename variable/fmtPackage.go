package main

import (
	"fmt"
)

/*
-> Using fmt package
-> %d to print int
-> %f to print float
-> %s or %q for string and quoted strings
-> %v replaced by any type
-> %T print the type
-> %t print the bool
-> %b converts int to binary or base2
-> %x converts int to hexadecimal
*/

func fmtPackcage() {
	name := "Gerard"
	fmt.Println("Hello, i am", name)

	a, b := 4, 6
	fmt.Println("Sum:", a+b, "Mean Value:", (a+b)/2)
	fmt.Printf("My age is %d \n", 34)
	fmt.Printf("x is %d, y is %f \n", 5, 6.8)
	fmt.Printf("He says: \"Hello Go!\"\n")

	figure := "Circle"
	radius := 5
	pi := 3.14159
	_, _ = figure, pi
	fmt.Printf("Radius is %d\n", radius)
	fmt.Printf("Radius is %+d\n", radius)

	fmt.Printf("Pi constant is %f\n", pi)
	fmt.Printf("The diameter of a %s with radius of %d is %f\n", figure, radius, float64(radius)*2*pi)

	fmt.Printf("This is a %q", figure)

	fmt.Printf("The diameter of a %v with radius of %v is %v\n", figure, radius, float64(radius)*2*pi)

	fmt.Printf("figure is of type %T\n", figure)
	fmt.Printf("radius is of type %T\n", radius)
	fmt.Printf("pi is of type %T\n", pi)

	isClosed := true
	fmt.Printf("File closed %t\n", isClosed)

	fmt.Printf("%b \n", 55)
	fmt.Printf("%08b \n", 55)

	fmt.Printf("%x \n", 100)

	x := 3.4
	y := 6.9

	fmt.Printf("x * y = %.2f\n", x*y)
}
