package main

import "fmt"

func main() {
	defer foo()
	bar()
	defer foo()

	bar()
	defer fmt.Println(3)
	defer fmt.Println(2)
	defer fmt.Println(1)
	fmt.Println(0)
}

// func (r receiver) identifier(p parameter(s)) (r return(s)) { code }

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}

/*

“defer” multiple functions in main
show that a deferred func runs after the func containing it exits.
determine the order in which the multiple defer funcs run

*/

// deferred functions run in LIFO order
// last in first out LIFO
