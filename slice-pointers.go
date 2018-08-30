package main

import "fmt"

func main() {
	names := []string{
		"John", "Paul", "George", "Ringo",
	}

	fmt.Println(names)
	fmt.Println(names[:2])
	fmt.Println(names[2:])

	a := names[0:2]
	b := names[1:3]

	fmt.Println(a, b)

	b[0] = "XXX"

	fmt.Println(a, b)
	fmt.Println(names)
}
