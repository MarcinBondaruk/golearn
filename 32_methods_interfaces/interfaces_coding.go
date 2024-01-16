package main

import "fmt"

type vehicle interface {
	License()
	Name()
}

type car struct {
	licenseNo string
	brand     string
}

func (c car) License() {
	fmt.Println(c.licenseNo)
}

func (c car) Name() {
	fmt.Println(c.brand)
}

func main() {
	myCar := car{
		licenseNo: "2137-42069",
		brand:     "Peugot",
	}

	myCar.License()
	myCar.Name()

	var empty interface{}
	fmt.Printf("%T\n", empty)
	empty = 3
	fmt.Printf("%T\n", empty)
	empty = 3.14
	fmt.Printf("%T\n", empty)
	empty = []int{1, 2, 3}
	fmt.Printf("%T\n", empty)
	empty = append(empty.([]int), 10)
	fmt.Printf("%#v\n", empty)
}
