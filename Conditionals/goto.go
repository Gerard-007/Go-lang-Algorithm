package Conditionals

import "fmt"

func main() {
	i := 0

loop:
	if i < 5 {
		fmt.Println(i)
		i++
		goto loop
	}

	//THis would work
	/*
			goto todo
		todo:
			fmt.Println("Something here...")
	*/

	//THis is a goto error
	/*
			goto todo
			x := 5
		todo:
			fmt.Println("Something here...")
	*/
}
