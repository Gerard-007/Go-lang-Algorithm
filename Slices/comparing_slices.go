package Slices

import (
	"fmt"
)

func main() {
	var n []int
	fmt.Println(n == nil)

	m := []int{}
	fmt.Println(m == nil)

	a, b := []int{1, 2, 3}, []int{1, 2, 3}
	//fmt.Println(a == b)
	//The above code Returns invalid operation can only be compared to nil

	//eq := true
	//fmt.Println(eq)
	//for i, valA := range a {
	//	if valA != b[i] {
	//		eq = false
	//		break
	//	}
	//}

	a = nil
	eq := true
	for i, valA := range a {
		if valA != b[i] {
			eq = false
			break
		}
	}

	if len(a) != len(b) {
		eq = false
	}

	if eq {
		fmt.Println("a and b slices are equal")
	} else {
		fmt.Println("a and b slices are not equal")
	}

	if eq {
		fmt.Println("a and b slices are equal")
	} else {
		fmt.Println("a and b slices are not equal")
	}
}
