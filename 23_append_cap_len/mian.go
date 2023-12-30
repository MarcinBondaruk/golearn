package main

import "fmt"

func main() {
	var nums []int
	fmt.Printf("%#v\n", nums)

	fmt.Printf("length: %d, cap: %d\n", len(nums), cap(nums))

	nums = append(nums, 1, 2)
	fmt.Printf("length: %d, cap: %d\n", len(nums), cap(nums))

	nums = append(nums, 3)
	fmt.Printf("length: %d, cap: %d\n", len(nums), cap(nums))

	nums = append(nums, 4, 5)
	fmt.Printf("length: %d, cap: %d\n", len(nums), cap(nums))

	nums = append(nums, 4, 5)
	fmt.Printf("length: %d, cap: %d\n", len(nums), cap(nums))

	// doubles backing array size when len reaches cap
}
