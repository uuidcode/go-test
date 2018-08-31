package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t T) M2() {
	fmt.Println(t.S)
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}

	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	var i I = &T{"Hello"}
	describe(i)
	i.M()

	var t *T
	describe(t)
	t.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}