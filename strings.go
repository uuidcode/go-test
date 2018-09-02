package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "0123456789"

	fmt.Println(len(text))
	fmt.Println(text[0:])
	fmt.Println(text[4:])
	fmt.Println(text[:4])

	text = "p:user:123"
	fmt.Println(strings.Split(text, ":"))
}
