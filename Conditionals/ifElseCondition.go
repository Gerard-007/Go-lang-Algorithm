package Conditionals

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	price, inStock := 100, true

	if price > 80 {
		fmt.Println("Too Expensive!")
	}
	// _ = inStock

	if price <= 100 && inStock {
		fmt.Println("Buy it!")
	}

	/*
		if price {
			fmt.Println("Something")
		}
		// This would raise compile time error!...
	*/

	if price < 100 {
		fmt.Println("Its cheap")
	} else if price == 100 {
		fmt.Println("On the edge")
	} else {
		fmt.Println("It's Expensive!")
	}

	age := 50

	if age >= 0 && age < 18 {
		fmt.Println("You cannot vote! Please return in d years!\n", 18-age)
	} else if age == 18 {
		fmt.Printf("This is your first vote")
	} else if age > 18 && age <= 100 {
		fmt.Printf("Please vote, it's important")
	} else {
		fmt.Printf("Invalid age!")
	}

	//i, err := strconv.Atoi("45a")

	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println(i)
	//	fmt.Printf("%T", i)
	//}

	if i, err := strconv.Atoi("20"); err == nil {
		fmt.Println("No error, i is : ", i)
	} else {
		fmt.Println(err)
	}

	if args := os.Args; len(args) != 2 {
		fmt.Println("One argument is required!")
	} else if km, err := strconv.Atoi(args[1]); err != nil {
		fmt.Println("The argument must be an integer! Error:", err)
	} else {
		fmt.Printf("%d km in miles is %v\n", km, float64(km)*1.609)
	}
}
