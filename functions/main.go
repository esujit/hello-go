package main

import (
	"fmt"
)

func main() {
	fmt.Println("go functions tutorial")
	greet("Sujit", "vaidya")
	greet2("Sujit", "vaidya")
	fmt.Println(greet3("Reyansh", "Vaidya"))
	fmt.Println(greet4("Rachana", "Vaidya"))
	fmt.Println(greet5("Subhash", "Vaidya"))
	n := average(22, 56, 1, 9.0, 5.2)
	fmt.Println(n)
	fmt.Println("==== using slice to pass variadic args ===")
	data := []float64{22, 102, 5, 96, 2.0}
	nslice := average(data...)
	fmt.Println(nslice)
	fmt.Println("==== Function expression ===")
	greeting := func() {
		fmt.Println("Hello from func expression")
	}
	greeting()
	fmt.Printf("%T\n", greeting)
	fmt.Println("=== Closure ===")
	greetClosure := makeGreeter()
	fmt.Println(greetClosure())
	fmt.Printf("%T\n", greetClosure)
	increment := wrapper()
	fmt.Println(increment())
	fmt.Println(increment())
	inc2 := wrapper()
	fmt.Printf("inc2 %v\n", inc2())
	fmt.Printf("inc2 %v\n", inc2())
	fmt.Println("=== callback example ===")
	visit([]int{1, 2, 3, 4}, func(n int) {
		fmt.Println(n)
	})


}

func greet(name string, lname string) {
	fmt.Println(name, lname)
}

func greet2(name, lname string) {
	fmt.Println(name, lname)
}

func greet3(fname, lname string) string {
	return fmt.Sprint(fname, string(' '), lname)
}

func greet4(fname, lname string) (s string) {
	s = fmt.Sprint(fname, string(' '), lname)
	return
}

func greet5(fname, lname string) (string, string) {
	return fmt.Sprint(fname, string(' '), lname), fmt.Sprint(lname, string(' '), fname)
}

func average(sf ...float64) float64 {
	fmt.Println(sf)
	fmt.Printf("%T\n", sf)
	//total := 0.0
	var total float64  // This will initialize to zero value
	for _, v := range sf {
		total += v
	}

	return total / float64(len(sf))
}

func makeGreeter() func() string {
	return func() string {
		return "Hello World from closure"
	}
}

func wrapper() func() int {
	var x int
	return func() int {
		x++
		return x

	}
}

func visit(num []int, callback func(int)) {
	for _, n := range num {
		callback(n)
	}
}