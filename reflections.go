package main

import (
	"fmt"
	"reflect"
)

type Element struct {
	left int    "left pixel"
	top  int    "top pixel"
	name string `one:"1" two:"2"`
}

func main() {
	elementType := reflect.TypeOf(Element{})

	{
		field, b := elementType.FieldByName("top")

		if b {
			tag := field.Tag
			fmt.Println(tag)
		}
	}

	{
		field, b := elementType.FieldByName("name")

		if b {
			tag := field.Tag
			fmt.Println(tag)
			value, ok := tag.Lookup("two")

			if ok {
				fmt.Println(value)
			}
		}
	}
}
