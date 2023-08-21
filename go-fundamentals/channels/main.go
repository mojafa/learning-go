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

/*
package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	cr := make(<-chan int) // receive
	cs := make(chan<- int) // send

	fmt.Println("-----")
	fmt.Printf("c\t%T\n", c)
	fmt.Printf("cr\t%T\n", cr)
	fmt.Printf("cs\t%T\n", cs)

	// specific to general doesn't assign
	c = cr
	c = cs
}
*/

/*
package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	cr := make(<-chan int) // receive
	cs := make(chan<- int) // send

	fmt.Println("-----")
	fmt.Printf("c\t%T\n", c)
	fmt.Printf("cr\t%T\n", cr)
	fmt.Printf("cs\t%T\n", cs)

	// general to specific converts
	fmt.Println("-----")
	fmt.Printf("c\t%T\n", (<-chan int)(c))
	fmt.Printf("c\t%T\n", (chan<- int)(c))

}
*/
