package main

import (
	"fmt"
	"reflect"
)

func main() {
	var i interface{} = "hello"
	fmt.Println(reflect.TypeOf(i))
	fmt.Printf("%T\n", i)

	s := i.(string)
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)
	//
	//f = i.(float64) // panic
	//fmt.Println(f)
}
