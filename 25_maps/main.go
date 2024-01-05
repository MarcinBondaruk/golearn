package main

import "fmt"

func main() {
	var m1 map[int]int

	fmt.Printf("%T %#v\n", m1, m1)
	fmt.Println(m1 == nil)

	m2 := map[int]string{1: "Led Zeppelin", 2: "Pink Floyd"}
	m2[10] = "Abba"

	fmt.Println(m2[1])
	fmt.Println(m2[3])

	m := map[int]bool{1: true, 2: false, 3: false}
	delete(m, 2)

	for k, v := range m {
		fmt.Println(k, v)
	}
}
