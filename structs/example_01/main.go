package main

import "fmt"

type person struct {
	first string
	last string
	age int
}

type DoubleZero struct {
	person
	LicenceToKill bool
}

type overrideField struct {
	person
	first string
}

func (p person) fullName() string {
	return p.first + string(' ') + p.last
}

func (p person) greeting() string {
	return "Greetings to " + p.first
}

func (o overrideField) greeting() string {
	return "Greetings to " + o.first
}

func main() {
	structInit1()
	structInit2()
	embedStruct()
	structOverrideInit()
}

func structInit1() {
	p1 := person{"Sujit", "Vaidya", 32}
	p2 := person{"Rachana", "Karnik", 28}
	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(p2.first, p2.last, p2.age)
}

func structInit2() {
	p1 := person{"Sujit", "Vaidya", 32}
	p2 := person{"Rachana", "Karnik", 28}
	fmt.Println(p1.fullName())
	fmt.Println(p2.fullName())
}

func embedStruct() {
	fmt.Println("embeded Struct example")
	p1 := DoubleZero{
		person: person{
			first: "Sudhakar",
			last: "Deshpande",
			age: 70,
		},
		LicenceToKill: false,
	}

	fmt.Println(p1)


}

func structOverrideInit() {
	fmt.Println("Embed struct field override")
	p1 := overrideField{
		person: person{
			first: "Sujit",
			last: "Vaidya",
			age: 40,
		},
		first: "Rachana",
	}

	fmt.Println("Field.person.first ", p1.person.first, "Field.first ", p1.first)
	fmt.Println(p1.greeting())
	fmt.Println(p1.person.greeting())
	fmt.Println(p1.fullName())
}