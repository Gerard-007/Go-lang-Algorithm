package Conditionals

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// While loop in go...
	j := 10
	for j >= 0 {
		fmt.Println(j)
		j--
	}

	// Infinite while loop
	sum := 0
	for {
		sum++
	}
	fmt.Println(sum) // this line is never reached

	// With continue...
	for k := 1; k <= 10; k++ {
		if k%2 != 0 {
			continue
		}
		fmt.Println(k)
	}

	// With break statement
	count := 0
	for l := 0; true; l++ {
		if l%13 == 0 {
			fmt.Printf("%d is divisible by 13 \n", l)
			count++
		}
		if count == 10 {
			break
		}
	}
}
