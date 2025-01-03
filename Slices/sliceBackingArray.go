package Slices

import "fmt"

func main() {
	s1 := []int{10, 20, 30, 40}
	s3, s4 := s1[0:2], s1[1:3]

	s3[1] = 600
	fmt.Println(s1) // prints [10 600 30 40]
	fmt.Println(s4) // prints [600 30]

	arr1 := [4]int{10, 20, 30, 40}
	slice1, slice2 := arr1[0:2], arr1[1:3]
	arr1[1] = 2
	fmt.Println(arr1)   // prints [10, 2, 30, 40
	fmt.Println(slice1) // prints [10, 2]
	fmt.Println(slice2) // prints [2, 30]
}
