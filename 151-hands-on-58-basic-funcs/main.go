package main

import "fmt"

func main() {
	x := foo()
	fmt.Println(x)

	i, y := bar()
	fmt.Println(i, y)

}

func foo() int {
	return 42
}

func bar() (int, string) {
	return 42, "Jafa"
}

/*
create a func with the identifier foo that returns an int
create a func with the identifier bar that returns an int and a string
call both funcs
print out their results
*/
