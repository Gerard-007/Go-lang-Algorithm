package main

import "fmt"

/*
- Data Type
*/

func main() {
	//INT Type
	var i1 int8 = 127 // this ranges to -128 to 128
	fmt.Printf("%T\n", i1)

	var i2 uint16 = 65535 // this ranges to -65535 to 65535
	fmt.Printf("%T\n", i2)

	//FLOAT Type
	var f1, f2, f3 float64 = 1.1, -0.2, 5.0
	fmt.Println(f1, f2, f3)

	//RUNE Type
	var r rune = 'f'
	fmt.Printf("%T\n", r)
	fmt.Println(r)
	fmt.Printf("%x\n", r)

	//BOOL Type
	var b bool = true
	fmt.Printf("%T\n", b)

	//String Type
	var s string = "Hello Go!"
	fmt.Printf("%T\n", s)

	//Array Type
	var numbers = [4]int{4, 5, -9, 100}
	fmt.Print(numbers)
	fmt.Printf("%T\n", numbers)

	//SLICE Type
	var cities = []string{"Lagos", "Abuja", "Onitsha", "Aba"}
	fmt.Print(cities)
	fmt.Printf("%T\n", cities)

	//MAP Type
	var balances = map[string]float64{
		"BFR": 7000.2,
		"NGN": 300.11,
		"USD": 4000.22,
		"EUR": 5000.32,
	}
	fmt.Print(balances)
	fmt.Printf("%T\n", balances)

	//STRUCT Type
	type Account struct {
		name string
		age  int
	}
	var user Account
	fmt.Printf("%T\n", user)

	//Pointer Type
	var x int = 2
	ptr := &x
	fmt.Printf("ptr is of type %T with the value of %v\n", ptr, ptr)

	// Function Type
	fmt.Printf("%T\n", f)
}

func f() {}
