package main

import (
	"strings"
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	First 	string
	Last 	string
	Age 	int
	notExported 	int
}

func main() {
	encode()
	decode()
}

func decode() {
	var p1 Person
	rdr := strings.NewReader(`{"First": "Sujit", "Last": "Vaidya", "Age": 40}`)
	json.NewDecoder(rdr).Decode(&p1)
	fmt.Println(p1)
}

func encode() {
	p1 := Person{
		First: "Rachana",
		Last: "Karnik",
		Age: 30,
		notExported: 1,
	}

	json.NewEncoder(os.Stdout).Encode(p1)
}
