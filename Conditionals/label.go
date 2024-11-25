package Conditionals

import "fmt"

func main() {
	// This doesn't conflict with outer Label
	outer := 19
	_ = outer

	family := [3]string{"Gerard", "Cynthia", "Chibundo"}
	friends := [3]string{"Cynthia", "Nasa", "Ugo"}

outer: // creating a label...
	for index, name := range family {
		for _, friend := range friends {
			if name == friend {
				fmt.Printf("Found a friend %q at index %d\n", friend, index)
				break outer
			}
		}
	}
	fmt.Println("Next instruction after the break...")
}
