package Array

import (
	"fmt"
)

func main() {
	//var numbers [4]int
	// How to print an array...
	//fmt.Printf("%v\n", numbers)
	//fmt.Printf("%#v\n", numbers)

	//accounts := [3]int{50, 60, 70}
	//for i := 0; i < len(accounts); i++ {
	//	fmt.Println(accounts[i])
	//}

	var a1 = [4]float64{}
	fmt.Printf("%#v\n", a1)

	var a2 = [3]int{-10, 1, 10}
	fmt.Printf("%#v\n", a2)

	var a3 = [3]string{"Gerard", "Gee", "Mix"}
	fmt.Printf("%#v\n", a3)

	var a4 = [3]string{"x", "y"}
	fmt.Printf("%#v\n", a4)

	// ellipsis operator...
	a5 := [...]int{1, 2, 4, 3, 5, 6, 8}
	fmt.Printf("%#v\n", a5)
	fmt.Printf("The length of a5 is %v\n", len(a5))

	a6 := [...]int{
		1,
		2,
		3,
		4,
		5, // This comma is mandatory
	}
	_ = a6
}
