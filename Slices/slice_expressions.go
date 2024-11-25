package Slices

import "fmt"

func main() {
	numbers := []int{2, 3}

	numbers = append(numbers, 10)
	//fmt.Println(numbers)

	numbers = append(numbers, 20, 30, 40)
	//fmt.Println(numbers)

	n := []int{100, 200}
	numbers = append(numbers, n...)
	//fmt.Println(numbers)

	src := []int{10, 20, 30}
	dst := make([]int, 2) // this would copy first 2 elements in src [10, 20]
	nn := copy(dst, src)  // prints number of items in dst: 2
	_ = nn
	//fmt.Println(src, dst, nn)

	/*
		SLICE Expression
	*/
	a := [5]int{1, 2, 3, 4, 5}
	// a[start:stop]
	b := a[1:3] // returns [2, 3]
	fmt.Printf("%v %T\n", b, b)

	s1 := []int{1, 2, 3, 4, 5, 6}

	s2 := s1[1:3]
	fmt.Println(s2) // prints [2, 3]

	fmt.Println(s1[2:]) //OR s1[2:len(s1)] prints [3, 4, 5, 6]
	fmt.Println(s1[2:]) //OR s1[0:3] prints [1, 2, 3]
	fmt.Println(s1[:])  //OR s1[0:len(s1)] prints [1, 2, 3, 4, 5, 6]

	// Appending to Slice
	s1 = append(s1[:4], 100)
	fmt.Println(s1) // Prints[1 2 3 4 100]

	s1 = append(s1[:4], 200)
	fmt.Println(s1) // Prints[1 2 3 4 200]
}
