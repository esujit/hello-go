package main

import "fmt"


/*
Anonymous function
function without a name

function expression
assiging a func to a value


 */
func example_closure() {
	x := 0

	//Anonymous functions
	increment := func() int {
		x++
		return x
	}

	fmt.Println(increment())
	fmt.Println(increment())
}

func wrapper() func() int {
	x := 0

	return func() int {
		x++
		return x
	}
}

func example_closure2() {
	increment := wrapper()
	fmt.Println(increment())
	fmt.Println(increment())
}

func main() {
	example_closure()
	example_closure2()

}
