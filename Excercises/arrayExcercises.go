package Excercises

import "fmt"

func main() {
	/*
		1. Using the var keyword, declare an array called cities with 2 elements of type string.
		Don't initialize the array.Print out the cities array and notice the value of its elements.
		2. Declare an array called grades of type [3]float64 and initialize only the first 2 elements
		using an array literal.Print out the grades array and notice the value of its elements.
		3. Declare an array called taskDone using the ellipsis operator.
		The elements are of type bool.Print out taskDone.
		4. Iterate over the array called cities using the classical for loop syntax and the len
		function and print out element by element.
		5. Iterate over grades using the range keyword and print out element by element.
	*/
	var cities = [2]string{"Lagos", "Abuja"}
	fmt.Printf("%#v\n", cities)

	grades := [3]float64{20, 40}
	fmt.Printf("%#v\n", grades)

	taskDone := [...]bool{true}
	fmt.Printf("%#v\n", taskDone)

	for i := 0; i < len(cities); i++ {
		fmt.Printf("%v\n", cities[i])
	}

	for i, v := range grades {
		fmt.Printf("%v, %v\n", i, v)
	}

	/*
		Consider the following array declaration: nums := [...]int{30, -1, -6, 90, -6}
		Write a Go program that counts the number of positive even numbers in the array.
	*/
	nums := [...]int{30, -1, -6, 90, -6}
	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			fmt.Printf("%v\n", cities[i])
		}
	}
	// Or...
	for _, v := range nums {
		if v%2 == 0 && v > 0 {
			fmt.Printf("%v\n", v)
		}
	}

	/*
		There are some errors in the following Go program.
		Try to identify the errors, change the code, and run the program without errors.
	*/
	myArray := [4]float64{1.2, 5.6}
	myArray[2] = 6
	a := 10
	myArray[0] = float64(a)
	myArray[3] = 10.10
	fmt.Println(myArray)
}
