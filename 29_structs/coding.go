package main

import "fmt"

func main() {
	type person struct {
		name      string
		age       int
		favColors []string
	}

	martin := person{
		name:      "Marcin",
		age:       33,
		favColors: []string{"Blue", "Violet"},
	}

	barbara := person{
		name:      "Barbara Brabant",
		age:       32,
		favColors: []string{"Orange", "Brown", "Green"},
	}

	fmt.Printf("%+v\n", martin)
	fmt.Printf("%+v\n", barbara)

	martin.name = "Not Martin"

	barbara.favColors = append(barbara.favColors, "White", "Red")
	fmt.Printf("%+v\n", barbara)

	for _, v := range barbara.favColors {
		fmt.Println(v)
	}
}
