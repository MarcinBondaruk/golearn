package main

import (
	"fmt"
	"unsafe"
)

func main() {
	numbers := []int{2, 3}

	// append produces new slice, do not modify
	numbers = append(numbers, 10, 12, 14)

	fmt.Println(numbers)

	newelements := []int{33, 44, 55}
	numbers = append(numbers, newelements...)
	fmt.Println(numbers)

	a := [5]int{1, 2, 3, 4, 5}
	// 1 include 4 exclude so included are indexes 1, 2, 3
	b := a[1:4]
	fmt.Printf("%#v %T\n", a, a)
	fmt.Printf("%#v %T\n", b, b)

	s1 := []int{1, 2, 3, 4, 5, 6}
	s2 := s1[1:3]
	fmt.Println(s2)
	fmt.Println(s1[4:])
	fmt.Println(s1[:3])
	fmt.Println(s1[:])

	s1 = append(s1[:4], 100)
	fmt.Println(s1)

	s1 = append(s1[:4], 200)
	fmt.Println(s1)

	// elements of the slice are stored in backing array, not in a slice itself
	// slice header is data structure SliceHeader
	// 3 fields
	// the address of bbacking array which is a pointer
	// the length, len() returns it
	// the capacity, cap() returns it

	s3 := []int{10, 20, 30, 40, 50}
	s4, s5 := s3[0:2], s3[1:3]

	s4[1] = 600

	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Println(s5)

	arr1 := [4]int{10, 20, 30, 40}
	slice1, slice2 := arr1[0:2], arr1[1:3]
	arr1[1] = 2

	fmt.Println(arr1, slice1, slice2)

	cars := []string{"Ford", "Toyota", "Audi", "Fiat"}
	newCars := []string{}
	newCars = append(newCars, cars[0:2]...)

	cars[0] = "Nissan"

	fmt.Println(cars, newCars)

	nums := []int{10, 20, 30, 40, 50}
	newSlice := nums[0:3]
	fmt.Println(len(newSlice), cap(newSlice))

	newSlice = append(newSlice, 200, 300)
	fmt.Println(len(newSlice), cap(newSlice))

	fmt.Printf("%p\n", &nums)
	fmt.Printf("%p %p\n", &nums, &newSlice) // diffrent slice header addresses

	x := [5]int{1, 2, 3, 4, 5}
	ss := []int{1, 2, 3, 4, 5}

	fmt.Printf("array size in bytes: %d\n", unsafe.Sizeof(x))  // 5 * 8 bytes
	fmt.Printf("slice size in bytes: %d\n", unsafe.Sizeof(ss)) // slice takes less memory because we ask for size of slice header address pointer, length, cap 3* 8 byte = 24
}
