package main

import "fmt"

func example1() {
	var x [10]string
	fmt.Printf("array x value %v\n", x)
	fmt.Printf("array x type %T\n", x)
	fmt.Println("length of array ", len(x))
	for i:=65; i<=74; i++ {
		x[i-65] = string(i)
	}
	fmt.Println(x)
	fmt.Println(x[2])
	fmt.Println(x[2:4])
}

func main() {
	example1()
}
