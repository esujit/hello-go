package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First	string
	Last 	string
	Age 	int
}



func main() {

	var p1 Person

	bs := []byte(`{"First": "Sujit", "Last": "Vaidya", "Age": 40}`)
	json.Unmarshal(bs, &p1)
	fmt.Println("============")
	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)
	fmt.Printf("%T\n", p1)
}


