package main

import (
	"fmt"
)

func zeroVariable() {
	var a = 7
	var b = 3.5

	//a = b //would give error unless we convert to int
	a = int(b)
	fmt.Println(a, b)

	// These returns zero values...
	var value int
	var price float64
	var name string
	var done bool
	fmt.Println(value, price, name, done)
}
