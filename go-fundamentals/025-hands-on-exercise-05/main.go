package main

import "fmt"

func main() {
	// Zero value
	var i uint64 = 0
	fmt.Println("Zero value of i:", i)
	fmt.Printf("i is of type %T\n", i)

	// Short declaration and multiple initializations
	x, y := 42, "Hello, world!"
	fmt.Println("x ", x)
	fmt.Println("and y:", y)
	fmt.Printf("x is of type %T\n", x)

	// Var when specificity is required
	var j int = 10
	fmt.Println("j:", j)
	fmt.Printf("j is of type %T\n", j)

	// Blank identifier
	_, k := "unused", "variable"
	fmt.Println("k:", k)
	fmt.Printf("k is of type %T\n", k)
}

// write a program that uses the following:
// ● var for zero value
// ● short declaration operator
// ● multiple initializations
// ● var when specificity is required
// ● blank identifier
// print items as necessary to make the program interesting
