package main

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/gommon/log"
	"io/ioutil"
	"net/http"
)

type Address struct {
	Street  string
	Suit    string
	City    string
	Zipcode string
}

type User struct {
	Id       int
	Name     string
	Username string
	Address  Address
	Phone    string
	Website  string
}

func main() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/users/1")

	if err != nil {
		log.Fatal(err)
	}

	bytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	body := string(bytes[:])

	fmt.Println("body", body)

	var resultMap map[string]interface{}
	err = json.Unmarshal([]byte(body), &resultMap)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resultMap)
	address := resultMap["address"].(map[string]interface{})
	fmt.Println("zipcode", address["zipcode"])

	user := User{}

	err = json.Unmarshal([]byte(body), &user)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("user %+v", user)

	resp.Body.Close()
}
