package Slices

import (
	"fmt"
	"unsafe"
)

func main() {
	cars := []string{"Ford", "Honda", "Audi", "Range Rover"}
	newCars := []string{}

	newCars = append(newCars, cars[0:2]...)

	cars[0] = "Nissan"
	fmt.Println(cars)    // Prints ["Nissan", "Honda", "Audi", "Range Rover"]
	fmt.Println(newCars) // Prints ["Ford", "Honda"]

	s1 := []int{10, 20, 30, 40, 50}
	newSlice := s1[0:3]
	fmt.Println(len(newSlice)) // Length: 3
	fmt.Println(cap(newSlice)) // Backing array capacity: 5

	newSlice = s1[2:5]         // [30, 40, 50]
	fmt.Println(len(newSlice)) // Length: 3
	fmt.Println(cap(newSlice)) // Capacity: 3

	// Printing the address of slices...
	fmt.Printf("%p\n", &s1)
	fmt.Printf("%p %p\n", &s1, &newSlice)

	newSlice[0] = 1000
	fmt.Println("s1: ", s1)

	a := [5]int{1, 2, 3, 4, 5}
	s := []int{1, 2, 3, 4, 5}

	fmt.Printf("Array's size in bytes: %d \n", unsafe.Sizeof(a))
	fmt.Printf("Slice's size in bytes: %d \n", unsafe.Sizeof(s))
}
