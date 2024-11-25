package Excercises

import "fmt"

func main() {
	/*
		Using a for loop, an if statement and the
		logical and operator print out all the
		numbers between 1 and 500 that divisible
		both by 7 and 5.
	*/
	//for i := 1; i <= 50; i++ {
	//	if i%7 == 0 {
	//		fmt.Println(i)
	//	}
	//}

	/*
		Change the code from the previous exercise and use
		the continue statement to print out all the
		numbers divisible by 7 between 1 and 50.
	*/
	//for i := 1; i <= 50; i++ {
	//	if i%7 != 0 {
	//		continue
	//	}
	//	fmt.Println(i)
	//}

	/*
		Change the code from the previous exercise and use the
		break statement to print out only the first 3 numbers
		divisible by 7 between 1 and 50.
	*/
	//counter := 0
	//for i := 1; i <= 50; i++ {
	//	if counter == 3 {
	//		break
	//	} else if i%7 == 0 {
	//		fmt.Println(i)
	//		counter++
	//	}
	//}

	birthYear, currentYear := 1990, 2024
	fmt.Printf("%T, %T\n", birthYear, currentYear)

	for i := birthYear; i < currentYear; {
		fmt.Printf("%d ", i)
		i++
	}

	age := -9

	switch {
	case age < 0 || age > 100:
		fmt.Println("Invalid Age")
	case age < 18:
		fmt.Println("You are minor!")
	case age == 18:
		fmt.Println("Congratulations! You've just become major!")
	default:
		fmt.Println("You are major!")
	}

}
