package main

import "fmt"

func main() {
	c := make(<-chan int, 2)

	c <- 42
	c <- 43

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println("-----")
	fmt.Printf("%T\n", c)
}

//cannot send to receive-only channel c (variable of type <-chan int)
