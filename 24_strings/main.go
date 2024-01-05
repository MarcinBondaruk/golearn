package main

import "fmt"

func main() {
	var name string = "Marcin"
	country := "Polska"

	fmt.Printf("Your name: %s\nCountry: %s\n", name, country)
	fmt.Println(`He says "Hello"`)
	fmt.Println(`C:\Users\a.txt`)

	var r rune = 'ă'
	s1, s2 := "ma", "m"

	mama := s1 + s2 + string(r)
	fmt.Println(mama)

	s3 := "țară means country in Romanian"

	for i := 0; i < len(s3); i++ {
		fmt.Printf("%d", s3[i])
	}
	fmt.Println()

	for i, r := range s3 {
		fmt.Printf("byte idx: %d rune: %c\n", i, r)
	}

	s4 := "你好 Go!"
	cs4 := []byte(s4)

	fmt.Printf("%#v\n", cs4)

	for i, b := range cs4 {
		fmt.Printf("%d %b\n", i, b)
	}

	rs := []rune(s4)

	fmt.Printf("%#v\n", rs)

	for i, r := range rs {
		fmt.Printf("%d -> %c\n", i, r)
	}
}
