package main

import (
	"fmt"
)

func main() {
	slice_init1()
	slice_init2()
	slice_init3()
	slice_init4()
	slice_inc()
}

func slice_init1() {
	student := []string{}
	fmt.Println(student)
	fmt.Println(student == nil)
}

func slice_init2() {
	//idiomatic way
	student := make([]int, 35)
	fmt.Println(student)
	fmt.Println(student == nil)
}

func slice_init3() {
	//add values only by appending
	var student []string
	fmt.Println(student)
	fmt.Println(student == nil)
}

func slice_init4() {
	student := make([]int, 0)
	fmt.Println(student)
	fmt.Println(student == nil)
}

func slice_inc() {
	mySlice := make([]int, 1)
	mySlice[0] = 7
	fmt.Println(mySlice)
	mySlice[0]++
	fmt.Println(mySlice)
}