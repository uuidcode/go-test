package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{Y: 2, X: 1}
	v3 = Vertex{}
	p1 = &Vertex{1, 2}
)

func main() {
	fmt.Println(Vertex{1, 2})

	v := Vertex{1, 2}
	v.X = 4

	fmt.Println(v.X)

	p := &v
	p.X = 1e9

	fmt.Println(v)

	fmt.Println(v1, v2, v3, p1)
}
