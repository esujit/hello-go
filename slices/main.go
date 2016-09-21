package main

import "fmt"

func main() {
	fmt.Println("Slice examples")
	mySlice := []int{1, 2, 3, 4}
	fmt.Printf("%T\n", mySlice)
	fmt.Println(mySlice)
	fmt.Println(mySlice[1:3])
}
