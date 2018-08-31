package main

import (
	"fmt"
	"math"
)

type Vertex4 struct {
	X, Y float64
}

func (v Vertex4) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Abs2(v Vertex4) float64 {
	return v.Abs()
}
func (v *Vertex4) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func Scale(v *Vertex4, f float64) {
	v.Scale(f)
}

func main() {
	v := Vertex4{
		3, 4,
	}

	v.Scale(10)
	fmt.Println(v.Abs())

	v2 := Vertex4{
		3, 4,
	}

	Scale(&v2, 10)
	fmt.Println(Abs2(v))
}
