package main

import (
	"fmt"
)

func multiVariable() {
	car, cost := "Camry", 50000
	fmt.Println(car, cost)

	car, year := "Camry", 2024
	_ = year

	var opened = false
	opened, file := true, "a.txt"
	_, _ = opened, file

	var (
		salary    float64
		firstName string
		gender    bool
	)

	fmt.Println(salary, firstName, gender)

	var a, b, c int
	fmt.Println(a, b, c)

	var i, j int
	i, j = 5, 8

	j, i = i, j // swapping variables
	// _, _ = i, j

	fmt.Println(i, j)

	sum := 5 + 2.3
	fmt.Println(sum)

}
