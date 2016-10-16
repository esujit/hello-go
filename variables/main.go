package main

import "fmt"

var name string = "sujit"                      //package scope
var b, c string = "stored in b", "stored in c" //package scope
const myName string = "Sujit Vaidya"
const y = 747

var x int = 42 //package scope

const (
	PI = 3.14
	Language = "GO"
)

const (
	A = iota //0
	B  //1
	C  //2
)

const (
	D = iota
	E
	F
)

const (
	_ = iota
	A1 = iota * 10 //10
	A2 = iota * 10 //20
)

const (
	_ = iota
	KB = 1 << (iota * 10)
	MB = 1 << (iota * 10)
)

func example_shorthand() {
	//short hand variables can be used only inside functions
	//declare, assign, initialize
	a := 10
	b := "golang"
	c := 4.17
	d := true
	fmt.Println("short hand declerations")
	fmt.Printf("a=%v type=%T\n", a, a)
	fmt.Printf("b=%v type=%T\n", b, b)
	fmt.Printf("c=%v type=%T\n", c, c)
	fmt.Printf("d=%v type=%T\n", d, d)
	fmt.Printf("name=%v type=%T\n", name, name)
	fmt.Printf("myName=%v type=%T\n", myName, myName)
	fmt.Printf("global const y=%v\n", y)
}

func example_zero_val_declaration() {
	var a int
	var b string
	var c float64
	var d bool
	fmt.Println("zero value declarations")
	fmt.Printf("a=%v type=%T\n", a, a)
	fmt.Printf("b=%v type=%T\n", b, b)
	fmt.Printf("c=%v type=%T\n", c, c)
	fmt.Printf("d=%v type=%T\n", d, d)

}

func example_const() {
	fmt.Println("--- const declerations ---")
	fmt.Printf("PI=%v type=%T\n", PI, PI)
	fmt.Printf("Language=%v type=%T\n", Language, Language)
	fmt.Printf("iota A=%v type=%T\n", A, A)
	fmt.Printf("iota B=%v type=%T\n", B, B)
	fmt.Printf("iota C=%v type=%T\n", C, C)
	fmt.Printf("iota D=%v type=%T\n", D, D)
	fmt.Printf("iota E=%v type=%T\n", E, E)
	fmt.Printf("iota F=%v type=%T\n", F, F)
	fmt.Printf("iota A1=%v type=%T\n", A1, A1)
	fmt.Printf("iota A2=%v type=%T\n", A2, A2)

	fmt.Printf("iota KB=%b type=%T\n", KB, KB)
	fmt.Printf("iota KB=%v type=%T\n", KB, KB)
	fmt.Printf("iota MB=%b type=%T\n", MB, MB)
	fmt.Printf("iota MB=%v type=%T\n", MB, MB)

}

func example_runes() {
	//rune literal
	foo := 'a'
	fmt.Println("Rune literal foo:='a' ", foo)
	fmt.Printf("foo type %T\n", foo)
	fmt.Printf("rune(foo) type %T\n", foo)
	fmt.Println("Rune ")
	fmt.Println("Hello") // string is a collection of runes
	fmt.Println([]byte("Hello"))
	/*for i := 500; i <=510; i++ {
		fmt.Println(i, " - ", string(i), " - ", []byte(string(i)))
	}*/

	//Raw string literals
	rawString := `This is raw string`
	fmt.Printf("rawString value %v  -- type %T\n", rawString, rawString)
	//Interpreted string literal
	interpretString := "This char is \x61"
	fmt.Printf("interpreted string value %v -- type %T\n", interpretString, interpretString)
}

func example_scope() {
	/*
		scope:
		universe
		package
		file
		block {}
		keep your scope tight !
	*/
	foo()
	y := 17 //block level order matters
	fmt.Println(y)

}

func foo() {
	fmt.Println(x)
}

func main() {
	example_shorthand()
	example_zero_val_declaration()
	example_scope()
	example_const()
	example_runes()

}
