package main

import "fmt"

type Vertex2 struct {
	Lat, Long float64
}

var m map[string]Vertex2

var m2 = map[string]Vertex2{
	"Bell Labs": {
		1, 2,
	},
	"Google": {
		3, 4,
	},
}

func main() {
	m = make(map[string]Vertex2)
	m["Bell Labs"] = Vertex2{
		40.68433, -74.39967,
	}

	fmt.Println(m["Bell Labs"])
	fmt.Println(m2)
}
