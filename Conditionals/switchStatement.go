package Conditionals

import (
	"fmt"
	"time"
)

func main() {
	language := "golang"

	switch language {
	case "python":
		fmt.Println("You are learning python...")
	case "Go", "golang":
		fmt.Println("Great! go for Go!")
	default:
		fmt.Println("Any other programming language is ok.")
	}

	n := 5
	switch true {
	case n%2 == 0:

		fmt.Println("Even number!")
	case n%2 != 0:
		fmt.Println("Odd number!")
	default:
		fmt.Println("Never here!")
	}

	hour := time.Now().Hour()
	// fmt.Println(hour)

	switch {
	case hour < 12:
		fmt.Println("Good Morning!")
	case hour < 17:
		fmt.Println("Good Afternoon!")
	default:
		fmt.Println("Good Evening!")
	}
}
