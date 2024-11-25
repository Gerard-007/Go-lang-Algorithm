package Slices

import "fmt"

func main() {
	var nums []int
	fmt.Printf("%#v\n", nums)

	fmt.Printf("length: %d, capacity: %d \n", len(nums), cap(nums))

	nums = append(nums, 1, 2)
	fmt.Printf("length: %d, capacity: %d \n", len(nums), cap(nums))

	nums = append(nums, 3)
	fmt.Printf("length: %d, capacity: %d \n", len(nums), cap(nums))

	nums = append(nums, 4)
	fmt.Printf("length: %d, capacity: %d \n", len(nums), cap(nums))

	nums = append(nums, 100)
	fmt.Printf("length: %d, capacity: %d \n", len(nums), cap(nums))

	nums = append(nums[0:4], 200, 300, 400, 500)
	fmt.Println(nums)
	fmt.Printf("length: %d, capacity: %d \n", len(nums), cap(nums))

	letters := []string{"A", "B", "C", "D", "E", "F"}
	letters = append(letters[0:1], "X", "Y")

	fmt.Println(letters, len(letters), cap(letters))
	// fmt.Println(letters[4])
	fmt.Println(letters[3:6])
}
