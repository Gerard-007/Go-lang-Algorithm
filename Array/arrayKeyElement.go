package Array

import "fmt"

func main() {
	grades := [3]int{
		1: 10,
		0: 5,
		2: 7,
	}
	fmt.Println(grades)

	accounts := [3]int{2: 50}
	fmt.Println(accounts)

	names := [...]string{
		5: "Dan",
	}
	fmt.Printf("%#v\n", names)

	states := [...]string{
		5: "Lagos",
		"Anambra",
		1: "Ebonyi",
	}
	fmt.Printf("%#v\n", states)

	weekend := [7]bool{5: true, 6: true}
	fmt.Printf("%#v\n", weekend)
}
