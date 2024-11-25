package Slices

import (
	"fmt"
	"strings"
)

func main() {
	var cities []string
	fmt.Println("cities is equal to nil: ", cities == nil)
	fmt.Printf("cities %#v\n", cities)
	fmt.Println(len(cities))

	// cities[0] = "Lagos" //This is a runtime Error
	numbers := []int{2, 3, 4, 5} // This is a slice
	fmt.Println(numbers)

	//Another way to make a slice using <make([]Type, Length)>
	nums := make([]int, 2)
	fmt.Println(strings.Repeat("#", 20))
	fmt.Printf("%#v\n", nums)

	type names []string
	friends := names{"Dan", "Chuks"}
	fmt.Println(strings.Repeat("#", 20))
	fmt.Println(friends)

	myFriend := friends[0]
	fmt.Println(strings.Repeat("#", 20))
	fmt.Println("My best friend is ", myFriend)

	friends[0] = "Gerard"
	fmt.Println("My best friend is ", friends[0])

	fmt.Println(strings.Repeat("#", 20))
	for index, value := range friends {
		fmt.Printf("index: %v, value: %v\n", index, value)
	}

	var n []int
	n = numbers
	fmt.Println(n)
}
