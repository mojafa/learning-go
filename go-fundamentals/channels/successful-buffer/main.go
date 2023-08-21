package main

import "fmt"

func main() {
	c := make(chan int, 1)

	c <- 42

	fmt.Println(<-c)

	// Path: go-fundamentals/channels/unsuccessful-buffer/main.go
	d := make(chan int, 2)

	d <- 42
	d <- 43

	fmt.Println(<-d)
	fmt.Println(<-d)
}
