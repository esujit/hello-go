package main

import "fmt"

func main() {
	map_init1()
	map_init2()
	map_init3()
	map_init4()
	map_init5()
}

func map_init1() {
	fmt.Println("using make to initialize a map")
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map before:", m)

	v, ok := m["k1"]
	fmt.Println("ok?:", ok, v)

	delete(m, "k2")
	fmt.Println("map after:", m)

	v, ok = m["k2"]
	fmt.Println("K2 ok?:", ok, v)
}

func map_init2() {
	fmt.Println("Using map literal to initialize")
	//map literal
	m := map[string]int{
		"k1": 0,
		"k2": 5,
	}

	fmt.Println("map: ", m)
	fmt.Printf("map[%s]=%v\n", "k2", m["k2"])
}

func map_init3() {
	fmt.Println("var ident map[] initializer")
	var myGreeting map[string]string
	fmt.Println(myGreeting == nil)
}

func map_init4() {
	fmt.Println("var ident = map[] initializer")
	var myGreeting = map[string]string{}
	fmt.Println(myGreeting == nil)
}

func map_init5() {
	myGreeting := map[int]string{
		0: "Good Morning",
		1: "Bonjour",
	}

	fmt.Println(myGreeting)

	if val, exists := myGreeting[0]; exists {
		fmt.Println("map[0]=", val)
	}
}