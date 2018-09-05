package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

type MyFloat5 float64

func (f MyFloat5) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}

	return float64(f)
}

type Vertex5 struct {
	X, Y float64
}

func (v *Vertex5) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type Sender interface {
	Send()
}

type Receiver interface {
	Receive()
}

type Phone struct {
}

func (Phone) Send() {

}

func (Phone) Receive() {

}

func main() {
	var a Abser
	f := MyFloat5(-math.Sqrt2)
	v := Vertex5{3, 4}

	a = f
	a = &v

	fmt.Println(a.Abs())

	var phone interface{}
	phone = Phone{}

	typeOf, ok := phone.(Sender)

	fmt.Printf("type:%T\n", phone)
	fmt.Println(typeOf)
	fmt.Println(ok)
}
