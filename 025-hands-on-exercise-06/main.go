package main

import "fmt"

func main() {
	var a string = "Hello"
	var b int64 = 42
	var c float64 = 3.14
	fmt.Printf("%v is of type %T\n ", a, a)
	fmt.Printf("%v is of type %T\n", b, b)
	fmt.Printf("%v is of type %T\n", c, c)

}

// write a program that uses the following:
// ● for a VARIABLE storing a VALUE of TYPE
//  ○ string
// ○ int
// ○ float64
// ● use print verbs to show
// ○ value
// ○ type
