package main

import "fmt"

func main() {
	c := make(chan int)
	// send
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)
	}()

	// receive
	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}

// you can't close a receive only channel
// cr := make (<-chan int)
// cs := make (chan<- int)
// close(cr)
