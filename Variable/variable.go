package main

import (
	"fmt"
)

func main() {
	/**
	var -> identifier
	age -> variable name
	int -> variable type
	30 -> value
	*/
	var age int = 30
	fmt.Println("Age: ", age)

	var name = "Gerard"
	fmt.Println("Name: ", name)
	// _ = name // this helps us to avoid compilation error due to unused variable

	s := "Learning GoLang"
	fmt.Println(s)

	name1 := "Jude"
	_ = name1
}
