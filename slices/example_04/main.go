package main

import "fmt"

func main() {

	greeting := []string{
		"Good Morning",
		"Bonjour",
		"Ola",
		"Gutten morgen",
		"Bon matai",
		"Subh divas",
	}

	for i, value := range greeting {
		fmt.Println(i, " - ", value)
	}

	fmt.Print("[1:2] ")
	fmt.Println(greeting[1:2])
}
