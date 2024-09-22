package operators

import "fmt"

/*
 */

func main() {
	a, b := 5, 10
	fmt.Println(a == b) // false
	fmt.Println(a != b) // true

	fmt.Println(a > b, a >= b)
	fmt.Println(a < b, a <= b)

	fmt.Println(a < b && b == 10) // returns true
	fmt.Println(a < 1 || a == b)  // returns false
	fmt.Println(!(a > b))         // true
}
