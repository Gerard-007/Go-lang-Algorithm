package main

import (
	"fmt"
	"strconv"
)

/*
 - Type casting
*/

func main() {
	var x = 3
	var y = 3.1

	x = x * int(y)
	fmt.Println(x)

	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", int(y))

	x = int(float64(x) * y)
	fmt.Println(x)

	y = float64(x) * y
	fmt.Println(y)

	var a = 5
	fmt.Printf("%T\n", a)

	var b int64 = 5
	fmt.Printf("%T\n", b)
	a = int(b)

	//STRING TYPE CASTING
	// Convert number to string
	s := string(99)
	fmt.Println(s) // returns the ascii code of 99 -> c

	myStr := fmt.Sprintf("%f", 44.2)
	fmt.Println(myStr)
	fmt.Printf("%T -> %s\n", myStr, myStr) // 44.2 is now a string

	myStr1 := fmt.Sprintf("%d", 34234)
	fmt.Printf("%T -> %s\n", myStr1, myStr1) // 34234 is now a string

	fmt.Println(string(34234)) // returns 薺 an ascii code

	s1 := "3.213" // type string
	fmt.Printf("%T\n", s1)

	// convert string to other types using strconv.ParseType(string)
	var f1, err = strconv.ParseFloat(s1, 64)
	_ = err
	fmt.Println(f1)

	i, err := strconv.Atoi("-50") // converts string to int
	s2 := strconv.Itoa(20)        // converts int to string

	fmt.Printf("i type is %T, i value is %v\n", i, i)
	fmt.Printf("s2 type is %T, s2 value is %q\n", s2, s2)
}
