package main

import "fmt"

func main() {
	mySlice := make([]int, 0, 5)

	fmt.Println("------------")
	fmt.Println(mySlice)
	fmt.Println("initial len: ", len(mySlice))
	fmt.Println("initial capacity: ", cap(mySlice))
	fmt.Println("-------------")

	for i := 0; i < 8; i++ {
		mySlice = append(mySlice, i)
		fmt.Println("index", i, "Value: ", mySlice[i], "Len:", len(mySlice), "Cap: ", cap(mySlice))

	}
}

