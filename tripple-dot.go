package main

import "fmt"

func show(name string, params ...int) {
	fmt.Println(name, params)
}

func remove(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}

func main() {
	params := []int{1, 2, 3, 4}
	show("hello", params...)
	fmt.Println(remove(params, 1))
}
