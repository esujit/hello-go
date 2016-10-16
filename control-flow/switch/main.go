package main

import "fmt"

func example1() {
	switch "Sujit" {
	case "Rachana", "Sujit":
		fmt.Println("Hello Rachana OR Sujit")
	case "Subhash":
		fmt.Println("Hello Subhash")
	case "Shobha":
		fmt.Println("hello shobha")

	}

	switch "Subhash" {
	case "Rachana":
		fmt.Println("Hello Rachana")
	case "Subhash":
		fmt.Println("Hello Subhash")
		fallthrough
	case "Shobha":
		fmt.Println("hello shobha")
	case "Sujit":
		fmt.Println("Hello Sujit")
	}

	name := "Sujit"

	switch {
	case len(name) == 5:
		fmt.Println("len 5")
	case len(name) == 6:
		fmt.Println("len 6")
	default:
		fmt.Println("default")
	}
}

type Contact struct {
	name string
	email string
}

func switchOnType(x interface{}) {

	switch x.(type) { //This is an assert; asserting "x is of this type"
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case Contact:
		fmt.Println("contact")
	default:
		fmt.Println("unknown")


	}
}

func example2() {
	switchOnType(7)
	var t = Contact{""}
}

func main() {
	example1()
	example2()

}
