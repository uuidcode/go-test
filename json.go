package main

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	Name string
	Body string
	Time int64
}

func main() {
	m := Message{"Alice", "Hello", 1294706395881547000}
	bytes, e := json.Marshal(m)

	if e == nil {
		fmt.Println(string(bytes))
	}

	text := `{"Name":"Bob","Bye":"Hello","Time":1294706395881547000}`
	e = json.Unmarshal([]byte(text), &m)

	if e == nil {
		fmt.Println(m)
	}
}
