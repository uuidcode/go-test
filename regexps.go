package main

import (
	"fmt"
	"regexp"
)

func main() {
	compile, e := regexp.Compile("^([0-9]{3})\\-([0-9]{4})\\-([0-9]{4})$")

	if e == nil {
		fmt.Println(compile.MatchString("010-1234-5678"))
		fmt.Println(compile.MatchString("010-123-5678"))

		submatch := compile.FindAllStringSubmatch("010-1234-5678", -1)

		for i, match := range submatch {
			fmt.Printf("%v, %v\n", i, match)
			fmt.Println(match[0])
			fmt.Println(match[1])
			fmt.Println(match[2])
			fmt.Println(match[3])
		}
	}
}
