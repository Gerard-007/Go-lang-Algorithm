package Array

import (
	"fmt"
	"strings"
)

func main() {
	numbers := [3]int{10, 20, 30}
	fmt.Printf("%#v\n", numbers)

	// Appending elements to list...
	numbers[0] = 40
	fmt.Printf("%#v\n", numbers)

	//numbers[5] = 40 // this is an error

	for i, v := range numbers {
		fmt.Println("index:", i, " value:", v)
	}

	fmt.Println(strings.Repeat("#", 20)) //This prints #####...*20
	for i := 0; i < len(numbers); i++ {
		fmt.Println("index:", i, " value:", numbers[i])
	}

	//Array in Array
	fmt.Println(strings.Repeat("#", 20))
	balances := [2][3]int{
		{5, 6, 7},
		[3]int{8, 9, 10},
	}
	fmt.Println(balances)

	fmt.Println(strings.Repeat("#", 20))
	m := [3]int{1, 2, 3}
	n := m
	fmt.Println("n is equals to m", n == m)
	m[0] = -1
	fmt.Println("n is equals to m", n == m)
}
