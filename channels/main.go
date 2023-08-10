package main

import "fmt"

func doSomething(x int) int {
	return x * 5
}
func main() {
	ch := make(chan int)
	go func() {
		ch <- doSomething(5)
	}()
	fmt.Println(<-ch)
}

/*
func main (){
	c := make(chan int)

	go func (){
	c <- 43
	}()
}



//unsucessful buffer
func main (){
	c := make(chan int, 1)

	go func (){
	c <- 43
	c <- 44
	}()
}


// succesful buffer
func main (){
	c := make(chan int, 2)
	c <- 43
	c <- 44


	fmt.Println(<-c)
	fmt.Println(<-c)
}

*/
