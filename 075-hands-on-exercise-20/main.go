package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := 0
	y := 450
	// fmt.Println(rand.Intn(y-x)+x, "is a random number between", x, "and", y)
	z := rand.Intn((y - x) + x)
	fmt.Printf("z=%v \n", z)
	//conditionals
	switch {
	case z >= 0 && z <= 100:
		fmt.Printf("z is %v and that is between 0 and 100\n", z)
	case z >= 101 && z <= 200:
		fmt.Printf("z is %v and that is between 101 and 200\n", z)
	case z >= 201 && z <= 250:
		fmt.Printf("z is %v and that is between 201 and 250\n", z)
	default:
		fmt.Println("z is greater than 250")
	}

	// ch1 := make(chan int)
	// ch2 := make(chan int)

	// d1 := time.Duration(rand.Intn(0))
	// d2 := time.Duration(rand.Intn(250))

	// go func() {
	// 	time.Sleep(d1 * time.Millisecond)
	// 	ch1 <- 1
	// }()
	// go func() {
	// 	time.Sleep(d2 * time.Millisecond)
	// 	ch2 <- 1
	// }()

	// select {
	// case <-ch1:
	// 	fmt.Println("ch1")
	// case <-ch2:
	// 	fmt.Println("ch2")
	// }
}

/*
● Modify the previous program to use one of these conditional logic statements
○ a switch statement
○ a select statement
● Which of the above conditional logic statements did you choose and why?
*/
