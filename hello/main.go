package main

import "fmt"

func main() {

	fmt.Println("hello")
	foo()
	fmt.Println("done")

	for i := 0; i < 100; i++ {
		if i%20 == 0 {
			fmt.Println(i)
		}
	}
	bar()
}

func foo() {
	fmt.Println("I'm foo")
}

func bar() {
	fmt.Println("and then we exited")
}
