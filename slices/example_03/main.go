package main

import "fmt"

func main() {
	mySlice := []int {
		1,
		2,
		3,
		6,
		11,
	}

	for i, value := range mySlice {
		fmt.Println(i, " - ", value)
	}
}
