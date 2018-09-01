package main

import "fmt"

type Ball struct {
	Radius   int
	Material string
}

type Football struct {
	Ball
}

func (ball Ball) Bounce() {
	fmt.Printf("bounce %+v", ball)
}

func main() {
	football := Football{}

	fmt.Printf("%+v\n", football)

	football = Football{
		Ball{
			Radius:   100,
			Material: "leather",
		},
	}

	fmt.Printf("%+v\n", football)

	football.Bounce()
}
